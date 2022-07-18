package record

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

// expose Record module api for user
type Client interface {
	sdk.Module

	CreateRecord(request CreateRecordRequest, baseTx sdk.BaseTx) (CreateRecordResp, sdk.Error)
	QueryRecord(request QueryRecordReq) (QueryRecordResp, sdk.Error)
}

type CreateRecordRequest struct {
	Contents []Content
}

type QueryRecordReq struct {
	RecordID string `json:"record_id"`
	Prove    bool   `json:"prove"`
	Height   int64  `json:"height"`
}

type QueryRecordResp struct {
	Record Data           `json:"record"`
	Proof  sdk.ProofValue `json:"proof"`
	Height int64          `json:"height"`
}

type Data struct {
	TxHash   string    `json:"tx_hash" yaml:"tx_hash"`
	Contents []Content `json:"contents" yaml:"contents"`
	Creator  string    `json:"creator" yaml:"creator"`
}

type CreateRecordResp struct {
	RecordId  string `json:"record_id"`
	Hash      string `json:"hash"`
	GasWanted int64  `json:"gas_wanted"`
	GasUsed   int64  `json:"gas_used"`
	Height    int64  `json:"height"`
}
