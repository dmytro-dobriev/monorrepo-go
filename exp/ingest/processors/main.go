package processors

import (
	"github.com/stellar/go/xdr"
)

// RootProcessor is useful when a pipeline needs to be split stream into
// multiple branches right away. This processor is a no-op - just passes the data
// to all children.
type RootProcessor struct {
	noStateProcessor
}

// CSVPrinter prints ledger entries to a file or stdout (when Filename is empty).
// Can be used both for processing state and ledgers.
type CSVPrinter struct {
	noStateProcessor

	Filename string
}

// StatusLogger prints number of processed entries every N
// entries.
type StatusLogger struct {
	noStateProcessor

	N int
}

// EntryTypeFilter is a pipeline.StateProcessor that filters out all
// entries that are not of type `Type`.
type EntryTypeFilter struct {
	noStateProcessor

	Type xdr.LedgerEntryType
}

type noStateProcessor struct{}

func (n *noStateProcessor) Reset() {
	// No internal state
}

// LedgerFilter is a pipeline.LedgerProcessor which ignores transactions occuring in ledgers
// before a given threshold
type LedgerFilter struct {
	noStateProcessor
	IgnoreLedgersBefore uint32
}
