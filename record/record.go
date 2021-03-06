package record

import (
	"encoding/hex"

	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type recordClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return recordClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (r recordClient) Name() string {
	return ModuleName
}

func (r recordClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (r recordClient) CreateRecord(request CreateRecordRequest, baseTx sdk.BaseTx) (CreateRecordResp, sdk.Error) {
	creator, err := r.QueryAddress(baseTx.From, baseTx.Password)
	createRecordResp := CreateRecordResp{}
	if err != nil {
		return createRecordResp, sdk.Wrap(err)
	}

	msg := &MsgCreateRecord{
		Contents: request.Contents,
		Creator:  creator.String(),
	}

	res, err := r.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return createRecordResp, err
	}
	if res.Events != nil {
		recordID, er := res.Events.GetValue(eventTypeCreateRecord, attributeKeyRecordID)
		if er != nil {
			return createRecordResp, sdk.Wrap(er)
		}
		createRecordResp.RecordId = recordID
	}
	createRecordResp.Hash = res.Hash
	createRecordResp.Height = res.Height
	createRecordResp.GasUsed = res.GasUsed
	createRecordResp.GasWanted = res.GasWanted

	return createRecordResp, nil
}

func (r recordClient) QueryRecord(request QueryRecordReq) (QueryRecordResp, sdk.Error) {
	rID, err := hex.DecodeString(request.RecordID)
	if err != nil {
		return QueryRecordResp{}, sdk.Wrapf("invalid record id, must be hex encoded string,but got %s", request.RecordID)
	}

	recordKey := GetRecordKey(rID)

	res, err := r.QueryStore(recordKey, ModuleName, request.Height, request.Prove)
	if err != nil {
		return QueryRecordResp{}, sdk.Wrap(err)
	}

	var record Record
	if err := r.Marshaler.UnmarshalBinaryBare(res.Value, &record); err != nil {
		return QueryRecordResp{}, sdk.Wrap(err)
	}

	result := record.Convert().(QueryRecordResp)

	var proof []byte
	if request.Prove {
		proof = r.MustMarshalJSON(res.ProofOps)
	}

	result.Proof = sdk.ProofValue{
		Proof: proof,
		Path:  []string{ModuleName, string(recordKey)},
		Value: res.Value,
	}
	result.Height = res.Height
	return result, nil
}
