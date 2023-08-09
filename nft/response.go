package nft

type QueryOwnerResp struct {
	OwnerResp  *OwnerResp    `json:"owner"`
	Pagination *PageResponse `json:"pagination"`
}

type OwnerResp struct {
	Address string `json:"address" yaml:"address"`
	IDCs    []IDC  `json:"idcs" yaml:"idcs"`
}

// IDC defines a set of nft ids that belong to a specific
type IDC struct {
	Class    string   `json:"class" yaml:"class"`
	TokenIDs []string `json:"token_ids" yaml:"token_ids"`
}

type PageResponse struct {
	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	NextKey []byte `json:"next_key"`
	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total uint64 `json:"total"`
}

// QueryNFTResp defines the response type for querying a single NFT
type QueryNFTResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	Owner   string `json:"owner"`
	URIHash string `json:"uri_hash"`
}

type QueryClassResp struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Schema           string `json:"schema"`
	Symbol           string `json:"symbol"`
	Creator          string `json:"creator"`
	Description      string `json:"description"`
	Uri              string `json:"uri"`
	UriHash          string `json:"uri_hash"`
	Data             string `json:"data"`
	MintRestricted   bool   `json:"mint_restricted"`
	UpdateRestricted bool   `json:"update_restricted"`
}

type QueryCollectionResp struct {
	Class      *QueryClassResp `json:"class" yaml:"class"`
	NFTs       []QueryNFTResp  `json:"nfts" yaml:"nfts"`
	Pagination *PageResponse   `json:"pagination"`
}

type QueryClassesResp struct {
	Classes    []QueryClassResp `json:"classes" yaml:"classes"`
	Pagination *PageResponse    `json:"pagination"`
}
