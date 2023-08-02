package mt

const (
	// ModuleName is the name of the module
	ModuleName = "mt"

	// StoreKey is the default store key for MT
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the MT store.
	QuerierRoute = ModuleName

	// RouterKey is the message route for the MT module
	RouterKey = ModuleName

	// KeyNextDenomSequence is the key used to store the next denom sequence in the keeper
	KeyNextDenomSequence = "nextDenomSequence"

	// KeyNextMTSequence is the key used to store the next MT sequence in the keeper
	KeyNextMTSequence = "nextMTSequence"
)

const (
	MsgTypeURLMintMT        = "/irismod.mt.MsgMintMT"
	MsgTypeURLBurnMT        = "/irismod.mt.MsgBurnMT"
	MsgTypeURLTransferMT    = "/irismod.mt.MsgTransferMT"
	MsgTypeURLIssueDenom    = "/irismod.mt.MsgIssueDenom"
	MsgTypeURLTransferDenom = "/irismod.mt.MsgTransferDenom"
)
