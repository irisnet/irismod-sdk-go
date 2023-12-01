package mt

type IssueClassRequest struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

type TransferClassRequest struct {
	ID        string `json:"id"`
	Recipient string `json:"recipient"`
}

type MintMTRequest struct {
	ClassID   string `json:"class_id"`
	Amount    uint64 `json:"amount"`
	Data      []byte `json:"data"`
	Recipient string `json:"recipient"`
}

type AddMTRequest struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	Amount    uint64 `json:"amount"`
	Recipient string `json:"recipient"`
}

type EditMTRequest struct {
	ID      string `json:"id"`
	ClassID string `json:"class_id"`
	Data    []byte `json:"data"`
}

type TransferMTRequest struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	Amount    uint64 `json:"amount"`
	Recipient string `json:"recipient"`
}

type BurnMTRequest struct {
	ID      string `json:"id"`
	ClassID string `json:"class_id"`
	Amount  uint64 `json:"amount"`
}

// QueryMTResp defines a multi token
// BaseMT non fungible token definition
type QueryMTResp struct {
	ID     string `json:"id"`
	Supply uint64 `json:"supply"`
	Data   []byte `json:"data"`
}

type QueryClassResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Data  []byte `json:"data"`
	Owner string `json:"owner"`
}

type QueryBalanceResp struct {
	MtId   string `json:"mt_id"`
	Amount uint64 `json:"amount"`
}

type QueryClassesResp struct {
	Classes    []QueryClassResp `json:"classes" yaml:"classes"`
	Pagination *PageResponse    `json:"pagination"`
}

type QueryMtsResp struct {
	Mts        []QueryMTResp `json:"mts" yaml:"mts"`
	Pagination *PageResponse `json:"pagination"`
}

type QueryBalancesResp struct {
	Balances   []QueryBalanceResp `json:"balances" yaml:"balances"`
	Pagination *PageResponse      `json:"pagination"`
}

type PageResponse struct {
	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	NextKey []byte `json:"next_key"`
	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total uint64 `json:"total"`
}
