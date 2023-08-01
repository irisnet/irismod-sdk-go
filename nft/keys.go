package nft

const (
	// ModuleName is the name of the module
	ModuleName = "nft"

	// StoreKey is the default store key for NFT
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the NFT store.
	QuerierRoute = ModuleName

	// RouterKey is the message route for the NFT module
	RouterKey = ModuleName
)

const (
	MsgTypeURLMintNFT       = "/irismod.nft.MsgMintNFT"
	MsgTypeURLTransferNFT   = "/irismod.nft.MsgTransferNFT"
	MsgTypeURLEditNFT       = "/irismod.nft.MsgEditNFT"
	MsgTypeURLBurnNFT       = "/irismod.nft.MsgBurnNFT"
	MsgTypeURLIssueDenom    = "/irismod.nft.MsgIssueDenom"
	MsgTypeURLTransferDenom = "/irismod.nft.MsgTransferDenom"
)
