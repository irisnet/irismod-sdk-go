package random

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
)

const (
	ModuleName = "random"

	eventTypeRequestRequestRandom = "request_random"
	attributeKeyRequestID         = "request_id"
	attributeKeyGenerateHeight    = "generate_height"
)

var (
	_ sdk.Msg = &MsgRequestRandom{}
)

// ValidateBasic implements Msg.
func (msg MsgRequestRandom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Consumer); err != nil {
		return sdkerrors.Wrapf(ErrAccAddressFromBech32, "invalid consumer address (%s)", err)
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgRequestRandom) GetSigners() []sdk.AccAddress {
	consumer, err := sdk.AccAddressFromBech32(msg.Consumer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{consumer}
}

func (m Random) Convert() interface{} {
	return QueryRandomResp{
		RequestTxHash: m.RequestTxHash,
		Height:        m.Height,
		Value:         m.Value,
	}
}

type Requests []Request

func (m Requests) Convert() interface{} {
	var res []QueryRandomRequestQueueResp

	for _, request := range m {
		q := QueryRandomRequestQueueResp{
			Height:           request.Height,
			Consumer:         request.Consumer,
			TxHash:           request.TxHash,
			Oracle:           request.Oracle,
			ServiceFeeCap:    request.ServiceFeeCap,
			ServiceContextId: request.ServiceContextId,
		}
		res = append(res, q)
	}
	return res
}
