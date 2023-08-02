package farm

// nolint
const (
	// module name
	ModuleName = "farm"

	// StoreKey is the default store key for farm
	StoreKey = ModuleName

	// RouterKey is the message route for farm
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the farm store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the farm querier
	QueryRecord = "farm"

	// RewardCollector is the root string for the reward distribution account address
	RewardCollector = "reward_collector"

	// EscrowCollector is the root string for the reward escrow account address
	EscrowCollector = "escrow_collector"

	// Prefix for farm pool_id
	PrefixFarmPool = "farm"
)

const (
	MsgTypeURLCreatePool = "/irismod.farm.MsgCreatePool"
	MsgTypeURLStake      = "/irismod.farm.MsgStake"
	MsgTypeURLUnstake    = "/irismod.farm.MsgUnstake"
	MsgTypeURLHarvest    = "/irismod.farm.MsgHarvest"
	MsgTypeURLDestroy    = "/irismod.farm.MsgDestroy"
	MsgTypeURLAdjust     = "/irismod.farm.MsgAdjust"
)
