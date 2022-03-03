package mt

import sdk "github.com/irisnet/core-sdk-go/types"

const (
	ModuleName = "mt"
)

var (
	_ sdk.Msg = &MsgIssueDenom{}
	_ sdk.Msg = &MsgMintMT{}
	_ sdk.Msg = &MsgEditMT{}
	_ sdk.Msg = &MsgTransferMT{}
	_ sdk.Msg = &MsgBurnMT{}
	_ sdk.Msg = &MsgTransferDenom{}
)

func (m *MsgMintMT) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMintMT) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMintMT) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMintMT) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgMintMT) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgIssueDenom) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgIssueDenom) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgIssueDenom) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgIssueDenom) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgIssueDenom) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgEditMT) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgEditMT) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgEditMT) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgEditMT) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgEditMT) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferMT) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferMT) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferMT) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferMT) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferMT) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgBurnMT) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgBurnMT) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgBurnMT) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgBurnMT) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgBurnMT) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferDenom) Route() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferDenom) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferDenom) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferDenom) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgTransferDenom) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

func (this Denom) Convert() interface{} {
	return QueryDenomResp{
		ID:    this.Id,
		Name:  this.Name,
		Data:  this.Data,
		Owner: this.Owner,
	}
}

type denoms []Denom

func (this denoms) Convert() interface{} {
	var denoms []QueryDenomResp
	for _, denom := range this {
		denoms = append(denoms, denom.Convert().(QueryDenomResp))
	}
	return denoms
}

func (this MT) Convert() interface{} {
	return QueryMTResp{
		ID:     this.Id,
		Supply: this.Supply,
		Data:   this.Data,
	}
}

type mts []MT

func (this mts) Convert() interface{} {
	var mts []QueryMTResp
	for _, mt := range this {
		mts = append(mts, mt.Convert().(QueryMTResp))
	}
	return mts
}
