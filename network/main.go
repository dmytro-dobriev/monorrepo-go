// Package network contains functions that deal with stellar network passphrases
// and IDs.
package network

import (
	"github.com/stellar/go/hash"
)

const (
	// PublicNetworkPassphrase is the pass phrase used for every transaction intended for the public stellar network
	PublicNetworkPassphrase = "Public Global Stellar Network ; September 2015"
	// TestNetworkPassphrase is the pass phrase used for every transaction intended for the SDF-run test network
	TestNetworkPassphrase = "Test SDF Network ; September 2015"
)

// ID returns the network ID derived from the provided passphrase.  This value
// also happens to be the raw (i.e. not strkey encoded) secret key for the root
// account of the network.
func ID(passphrase string) [32]byte {
	return hash.Hash([]byte(passphrase))
}
