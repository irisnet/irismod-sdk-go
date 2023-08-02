package htlc

const (
	// ModuleName is the name of the HTLC module
	ModuleName = "htlc"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the HTLC module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the HTLC module
	RouterKey string = ModuleName

	// DefaultParamspace is the default name for parameter store
	DefaultParamspace = ModuleName
)
