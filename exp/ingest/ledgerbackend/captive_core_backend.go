package ledgerbackend

import (
	"io"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/stellar/go/network"
	"github.com/stellar/go/support/historyarchive"
	"github.com/stellar/go/support/log"
	"github.com/stellar/go/xdr"
)

// Ensure CaptiveStellarCore implements LedgerBackend
var _ LedgerBackend = (*CaptiveStellarCore)(nil)

// This is a not-very-complete or well-organized sketch of code be used to
// stream LedgerCloseMeta data from a "captive" stellar-core: one running as a
// subprocess and replaying portions of history against an in-memory ledger.
//
// A captive stellar-core still needs (and allocates, in os.TempDir()) a
// temporary directory to run in: one in which its config file is stored, along
// with temporary files it downloads and decompresses, and its bucket
// state. Only the ledger will be in-memory (and we might even switch this to
// SQLite + large buffers in the future if the in-memory ledger gets too big.)
//
// Feel free to reorganize this to fit better. It's preliminary!

// TODO: switch from history URLs to history archive interface provided from support package, to permit mocking

const (
	ledgersPerCheckpoint = 64
	readAheadBufferSize  = 2
)

func roundDownToFirstReplayAfterCheckpointStart(ledger uint32) uint32 {
	v := (ledger / ledgersPerCheckpoint) * ledgersPerCheckpoint
	if v == 0 {
		return 1
	}
	// All other checkpoints start at the next multiple of 64
	return v
}

type metaResult struct {
	*xdr.LedgerCloseMeta
	err error
}

type CaptiveStellarCore struct {
	networkPassphrase string
	historyURLs       []string

	// read-ahead buffer
	stop  chan struct{}
	wait  sync.WaitGroup
	metaC chan metaResult

	stellarCoreRunner stellarCoreRunnerInterface
	cachedMeta        *LedgerCloseMeta

	// Defines if the blocking mode (off by default) is on or off. In blocking mode,
	// calling GetLedger blocks until the requested ledger is available. This is useful
	// for scenarios when Horizon consumes ledgers faster than Stellar-Core produces them
	// and using `time.Sleep` when ledger is not available can actually slow entire
	// ingestion process.
	blocking bool

	nextLedger uint32  // next ledger expected, error w/ restart if not seen
	lastLedger *uint32 // end of current segment if offline, nil if online
}

// NewCaptive returns a new CaptiveStellarCore that is not running. Will lazily start a subprocess
// to feed it a block of streaming metadata when user calls .GetLedger(), and will kill
// and restart the subprocess if subsequent calls to .GetLedger() are discontiguous.
//
// Platform-specific pipe setup logic is in the .start() methods.
func NewCaptive(executablePath, networkPassphrase string, historyURLs []string) *CaptiveStellarCore {
	return &CaptiveStellarCore{
		networkPassphrase: networkPassphrase,
		historyURLs:       historyURLs,
		nextLedger:        0,
		stellarCoreRunner: &stellarCoreRunner{
			executablePath:    executablePath,
			networkPassphrase: networkPassphrase,
			historyURLs:       historyURLs,
		},
	}
}

// Returns the sequence number of an LCM, returning an error if the LCM is of
// an unknown version.
func peekLedgerSequence(xlcm *xdr.LedgerCloseMeta) (uint32, error) {
	v0, ok := xlcm.GetV0()
	if !ok {
		err := errors.New("unexpected XDR LedgerCloseMeta version")
		return 0, err
	}
	return uint32(v0.LedgerHeader.Header.LedgerSeq), nil
}

// Note: the xdr.LedgerCloseMeta structure is _not_ the same as
// the ledgerbackend.LedgerCloseMeta structure; the latter should
// probably migrate to the former eventually.
func (c *CaptiveStellarCore) copyLedgerCloseMeta(xlcm *xdr.LedgerCloseMeta, lcm *LedgerCloseMeta) error {
	v0, ok := xlcm.GetV0()
	if !ok {
		return errors.New("unexpected XDR LedgerCloseMeta version")
	}
	lcm.LedgerHeader = v0.LedgerHeader
	envelopes := make(map[xdr.Hash]xdr.TransactionEnvelope)
	for _, tx := range v0.TxSet.Txs {
		hash, err := network.HashTransactionInEnvelope(tx, c.networkPassphrase)
		if err != nil {
			return errors.Wrap(err, "error hashing tx in LedgerCloseMeta")
		}
		envelopes[hash] = tx
	}
	for _, trm := range v0.TxProcessing {
		txe, ok := envelopes[trm.Result.TransactionHash]
		if !ok {
			return errors.New("unknown tx hash in LedgerCloseMeta")
		}
		lcm.TransactionEnvelope = append(lcm.TransactionEnvelope, txe)
		lcm.TransactionResult = append(lcm.TransactionResult, trm.Result)
		lcm.TransactionMeta = append(lcm.TransactionMeta, trm.TxApplyProcessing)
		lcm.TransactionFeeChanges = append(lcm.TransactionFeeChanges, trm.FeeProcessing)
	}
	for _, urm := range v0.UpgradesProcessing {
		lcm.UpgradesMeta = append(lcm.UpgradesMeta, urm.Changes)
	}
	return nil
}

func (c *CaptiveStellarCore) getLatestCheckpointSequence() (uint32, error) {
	archive, err := historyarchive.Connect(
		c.historyURLs[0],
		historyarchive.ConnectOptions{},
	)
	if err != nil {
		return 0, errors.Wrap(err, "error connecting to history archive")
	}
	has, err := archive.GetRootHAS()
	if err != nil {
		return 0, errors.Wrap(err, "error getting root HAS")
	}

	return has.CurrentLedger, nil
}

func (c *CaptiveStellarCore) openOfflineReplaySubprocess(from, to uint32) error {
	err := c.Close()
	if err != nil {
		return errors.Wrap(err, "error closing existing session")
	}

	latestCheckpointSequence, err := c.getLatestCheckpointSequence()
	if err != nil {
		return errors.Wrap(err, "error getting latest checkpoint sequence")
	}
	if from > latestCheckpointSequence {
		return errors.Errorf(
			"sequence: %d is greater than max available in history archives: %d",
			from,
			latestCheckpointSequence,
		)
	}
	if to > latestCheckpointSequence {
		to = latestCheckpointSequence
	}

	err = c.stellarCoreRunner.catchup(from, to)
	if err != nil {
		return errors.Wrap(err, "error running stellar-core")
	}

	// The next ledger should be the first ledger of the checkpoint containing
	// the requested ledger
	c.nextLedger = roundDownToFirstReplayAfterCheckpointStart(from)
	c.lastLedger = &to
	c.blocking = true

	// read-ahead buffer
	c.metaC = make(chan metaResult, readAheadBufferSize)
	c.stop = make(chan struct{})
	c.wait.Add(1)
	go c.sendLedgerMeta(to)
	return nil
}

func (c *CaptiveStellarCore) openOnlineReplaySubprocess(from uint32) error {
	// Check if existing session works for this request
	if c.lastLedger == nil && c.nextLedger != 0 && c.nextLedger <= from {
		// Use existing session, GetLedger will fast-forward
		return nil
	}

	err := c.Close()
	if err != nil {
		return errors.Wrap(err, "error closing existing session")
	}

	latestCheckpointSequence, err := c.getLatestCheckpointSequence()
	if err != nil {
		return errors.Wrap(err, "error getting latest checkpoint sequence")
	}

	// We don't allow starting the online mode starting with more than two
	// checkpoints from now. Such requests are likely buggy.
	// We should allow only one checkpoint here but sometimes there are up to a
	// minute delays when updating root HAS by stellar-core.
	maxLedger := latestCheckpointSequence + 2*64
	if from > maxLedger {
		return errors.Errorf(
			"trying to start online mode too far (latest checkpoint=%d), only two checkpoints in the future allowed",
			latestCheckpointSequence,
		)
	}

	err = c.stellarCoreRunner.runFrom(from)
	if err != nil {
		return errors.Wrap(err, "error running stellar-core")
	}

	// The next ledger should be the ledger actually requested because
	// we run `catchup X/0` command in the online mode.
	c.nextLedger = from
	c.lastLedger = nil
	c.blocking = false

	// read-ahead buffer
	c.metaC = make(chan metaResult, readAheadBufferSize)
	c.stop = make(chan struct{})
	c.wait.Add(1)
	go c.sendLedgerMeta(0)
	return nil
}

// sendLedgerMeta reads from the captive core pipe, decodes the ledger metadata
// and sends it to the metadata buffered channel
func (c *CaptiveStellarCore) sendLedgerMeta(untilSequence uint32) {
	defer c.wait.Done()
	printBufferOccupation := time.NewTicker(5 * time.Second)
	defer printBufferOccupation.Stop()
	for {
		select {
		case <-c.stop:
			return
		case <-printBufferOccupation.C:
			log.Debug("captive core read-ahead buffer occupation:", len(c.metaC))
		default:
		}
		meta, err := c.readLedgerMetaFromPipe()
		if err != nil {
			select {
			case <-c.stop:
			case c.metaC <- metaResult{nil, err}:
			}
			return
		}
		select {
		case <-c.stop:
			return
		case c.metaC <- metaResult{meta, nil}:
		}

		if untilSequence != 0 {
			seq, err := peekLedgerSequence(meta)
			if err != nil {
				select {
				case <-c.stop:
				case c.metaC <- metaResult{nil, err}:
				}
				return
			}
			if seq >= untilSequence {
				// we are done
				return
			}
		}
	}
}

func (c *CaptiveStellarCore) readLedgerMetaFromPipe() (*xdr.LedgerCloseMeta, error) {
	metaPipe := c.stellarCoreRunner.getMetaPipe()
	if metaPipe == nil {
		return nil, errors.New("missing metadata pipe")
	}
	var xlcm xdr.LedgerCloseMeta
	_, e0 := xdr.UnmarshalFramed(metaPipe, &xlcm)
	if e0 != nil {
		if e0 == io.EOF {
			return nil, errors.Wrap(e0, "got EOF from subprocess")
		} else {
			return nil, errors.Wrap(e0, "unmarshalling framed LedgerCloseMeta")
		}
	}
	return &xlcm, nil
}

// PrepareRange prepares the given range (including from and to) to be loaded.
// Some backends (like captive stellar-core) need to initalize data to be
// able to stream ledgers.
// Set `to` to 0 to stream starting from `from` but without limits.
func (c *CaptiveStellarCore) PrepareRange(from uint32, to uint32) error {
	if c.nextLedger != 0 && c.nextLedger >= from {
		// Range already prepared
		return nil
	}

	var err error
	if to == 0 {
		err = c.openOnlineReplaySubprocess(from)
	} else {
		err = c.openOfflineReplaySubprocess(from, to)
	}
	if err != nil {
		return errors.Wrap(err, "opening subprocess")
	}

	metaPipe := c.stellarCoreRunner.getMetaPipe()
	if metaPipe == nil {
		return errors.New("missing metadata pipe")
	}

	for {
		// Wait for the first ledger
		if len(c.metaC) > 0 {
			break
		}
		time.Sleep(time.Second)
	}

	return nil
}

// GetLedger returns true when ledger is found and it's LedgerCloseMeta.
// Call `PrepareRange` first to instruct the backend which ledgers to fetch.
//
// We assume that we'll be called repeatedly asking for ledgers in ascending
// order, so when asked for ledger 23 we start a subprocess doing catchup
// "100023/100000", which should replay 23, 24, 25, ... 100023. The wrinkle in
// this is that core will actually replay from the _checkpoint before_
// the implicit start ledger, so we might need to skip a few ledgers until
// we hit the one requested (this routine does so transparently if needed).
func (c *CaptiveStellarCore) GetLedger(sequence uint32) (bool, LedgerCloseMeta, error) {
	if c.cachedMeta != nil && sequence == uint32(c.cachedMeta.LedgerHeader.Header.LedgerSeq) {
		// GetLedger can be called multiple times using the same sequence, ex. to create
		// change and transaction readers. If we have this ledger buffered, let's return it.
		return true, *c.cachedMeta, nil
	}

	if c.isClosed() {
		return false, LedgerCloseMeta{}, errors.New("session is closed, call PrepareRange first")
	}

	// Now loop along the range until we find the ledger we want.
	var errOut error
loop:
	for {
		if !c.blocking {
			if len(c.metaC) == 0 {
				return false, LedgerCloseMeta{}, nil
			}
		}

		metaResult := <-c.metaC
		if metaResult.err != nil {
			errOut = metaResult.err
			break loop
		}

		seq, e1 := peekLedgerSequence(metaResult.LedgerCloseMeta)
		if e1 != nil {
			errOut = e1
			break
		}
		if seq != c.nextLedger {
			// We got something unexpected; close and reset
			errOut = errors.Errorf("unexpected ledger (expected=%d actual=%d)", c.nextLedger, seq)
			break
		}
		c.nextLedger++
		if seq == sequence {
			// Found the requested seq
			var lcm LedgerCloseMeta
			e2 := c.copyLedgerCloseMeta(metaResult.LedgerCloseMeta, &lcm)
			if e2 != nil {
				errOut = e2
				break
			}

			c.cachedMeta = &lcm

			// If we got the _last_ ledger in a segment, close before returning.
			if c.lastLedger != nil && *c.lastLedger == seq {
				if err := c.Close(); err != nil {
					return false, LedgerCloseMeta{}, errors.Wrap(err, "error closing session")
				}
			}
			return true, lcm, nil
		}
	}
	// All paths above that break out of the loop (instead of return)
	// set e to non-nil: there was an error and we should close and
	// reset state before retuning an error to our caller.
	c.Close()
	return false, LedgerCloseMeta{}, errOut
}

// GetLatestLedgerSequence returns the sequence of the latest ledger available
// in the backend.
// Will return error if not in session (start with `PrepareRange`).
func (c *CaptiveStellarCore) GetLatestLedgerSequence() (uint32, error) {
	if c.isClosed() {
		return 0, errors.New("stellar-core must be opened to return latest available sequence")
	}

	if c.lastLedger == nil {
		// TODO Get latest buffered ledger when XDR buffer is ready
		if len(c.metaC) > 0 {
			return c.nextLedger, nil
		} else {
			return c.nextLedger - 1, nil
		}
	} else {
		return *c.lastLedger, nil
	}
}

func (c *CaptiveStellarCore) isClosed() bool {
	return c.nextLedger == 0
}

// Close closes existing stellar-core process and streaming sessions.
func (c *CaptiveStellarCore) Close() error {
	if c.isClosed() {
		return nil
	}
	c.nextLedger = 0
	c.lastLedger = nil

	if c.stop != nil {
		close(c.stop)
		// discard pending data in case the goroutine is blocked writing to the channel
		select {
		case <-c.metaC:
		default:
		}
		// Do not close the communication channel until we know
		// the goroutine is done
		c.wait.Wait()
		close(c.metaC)
	}

	err := c.stellarCoreRunner.close()
	if err != nil {
		return errors.Wrap(err, "error closing stellar-core subprocess")
	}
	return nil
}
