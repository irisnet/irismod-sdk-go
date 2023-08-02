package service

const (
	// ModuleName is the name of the service module
	ModuleName = "service"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the service module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the service module
	RouterKey string = ModuleName

	// DefaultParamspace is the default name for parameter store
	DefaultParamspace = ModuleName

	// DepositAccName is the root string for the service deposit account address
	DepositAccName = "service_deposit_account"

	// RequestAccName is the root string for the service request account address
	RequestAccName = "service_request_account"

	// FeeCollectorName is the root string for the service fee collector address used for testing
	FeeCollectorName = "service_fee_collector"
)
