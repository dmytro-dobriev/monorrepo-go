package paths

import (
	"github.com/stellar/go/xdr"
)

// Query is a query for paths
type Query struct {
	DestinationAsset    xdr.Asset
	DestinationAmount   xdr.Int64
	SourceAssets        []xdr.Asset
	SourceAssetBalances []xdr.Int64
	// if ValidateSourceBalance is true then we won't consider payment paths
	// which require a source asset amount which exceeds the balance present in `SourceAssetBalances`
	ValidateSourceBalance bool
	SourceAccount         *xdr.AccountId
}

// Path is the result returned by a path finder and is tied to the DestinationAmount used in the input query
type Path struct {
	Path              []xdr.Asset
	Source            xdr.Asset
	SourceAmount      xdr.Int64
	Destination       xdr.Asset
	DestinationAmount xdr.Int64
}

// Finder finds paths.
type Finder interface {
	// Returns path for a Query of a maximum length `maxLength`
	Find(q Query, maxLength uint) ([]Path, error)
	// FindFixedPaths return a list of payment paths each of which
	// start by spending `amountToSpend` of `sourceAsset` and end
	// with delivering a postive amount of `destinationAsset`
	FindFixedPaths(
		sourceAsset xdr.Asset,
		amountToSpend xdr.Int64,
		destinationAssets []xdr.Asset,
		maxLength uint,
	) ([]Path, error)
}
