package record

// nolint
const (
	// module name
	ModuleName = "record"

	// StoreKey is the default store key for record
	StoreKey = ModuleName

	// RouterKey is the message route for record
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the record store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the record querier
	QueryRecord = "record"
)
