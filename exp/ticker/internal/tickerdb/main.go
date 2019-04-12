package tickerdb

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stellar/go/support/db"
)

// TickerSession provides helper methods for making queries against `DB`.
type TickerSession struct {
	db.Session
}

// Asset represents an entry on the assets table
type Asset struct {
	ID                          int32     `db:"id"`
	Code                        string    `db:"code"`
	IssuerAccount               string    `db:"issuer_account"`
	Type                        string    `db:"type"`
	NumAccounts                 int32     `db:"num_accounts"`
	AuthRequired                bool      `db:"auth_required"`
	AuthRevocable               bool      `db:"auth_revocable"`
	Amount                      float64   `db:"amount"`
	AssetControlledByDomain     bool      `db:"asset_controlled_by_domain"`
	AnchorAssetCode             string    `db:"anchor_asset_code"`
	AnchorAssetType             string    `db:"anchor_asset_type"`
	IsValid                     bool      `db:"is_valid"`
	ValidationError             string    `db:"validation_error"`
	LastValid                   time.Time `db:"last_valid"`
	LastChecked                 time.Time `db:"last_checked"`
	DisplayDecimals             int       `db:"display_decimals"`
	Name                        string    `db:"name"`
	Desc                        string    `db:"description"`
	Conditions                  string    `db:"conditions"`
	IsAssetAnchored             bool      `db:"is_asset_anchored"`
	FixedNumber                 int       `db:"fixed_number"`
	MaxNumber                   int       `db:"max_number"`
	IsUnlimited                 bool      `db:"is_unlimited"`
	RedemptionInstructions      string    `db:"redemption_instructions"`
	CollateralAddresses         string    `db:"collateral_addresses"`
	CollateralAddressSignatures string    `db:"collateral_address_signatures"`
	Countries                   string    `db:"countries"`
	Status                      string    `db:"status"`
	IssuerID                    int32     `db:"issuer_id"`
}

// Issuer represents an entry on the issuers table
type Issuer struct {
	ID               int32  `db:"id"`
	PublicKey        string `db:"public_key"`
	Name             string `db:"name"`
	URL              string `db:"url"`
	TOMLURL          string `db:"toml_url"`
	FederationServer string `db:"federation_server"`
	AuthServer       string `db:"auth_server"`
	TransferServer   string `db:"transfer_server"`
	WebAuthEndpoint  string `db:"web_auth_endpoint"`
	DepositServer    string `db:"deposit_server"`
	OrgTwitter       string `db:"org_twitter"`
}

// Trade represents an entry on the trades table
type Trade struct {
	ID              int32     `db:"id"`
	HorizonID       string    `db:"horizon_id"`
	LedgerCloseTime time.Time `db:"ledger_close_time"`
	OfferID         string    `db:"offer_id"`
	BaseOfferID     string    `db:"base_offer_id"`
	BaseAccount     string    `db:"base_account"`
	BaseAmount      float64   `db:"base_amount"`
	BaseAssetID     int32     `db:"base_asset_id"`
	CounterOfferID  string    `db:"counter_offer_id"`
	CounterAccount  string    `db:"counter_account"`
	CounterAmount   float64   `db:"counter_amount"`
	CounterAssetID  int32     `db:"counter_asset_id"`
	BaseIsSeller    bool      `db:"base_is_seller"`
	Price           float64   `db:"price"`
}

// CreateSession returns a new TickerSession that connects to the given db settings
func CreateSession(driverName, dataSourceName string) (session TickerSession, err error) {
	dbconn, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return
	}

	session.DB = dbconn
	return
}