package mt

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"strings"
)

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

func (m MsgMintMT) Route() string {
	return ModuleName
}

func (m MsgMintMT) Type() string {
	return "mint_mt"
}

func (m MsgMintMT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return sdk.Wrap(err)
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return sdk.Wrapf("missing denom")
	}

	return nil
}

func (m MsgMintMT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgMintMT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgIssueDenom) Route() string {
	return ModuleName
}

func (m MsgIssueDenom) Type() string {
	return "issue_denom"
}

func (m MsgIssueDenom) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return sdk.Wrapf("missing sender address")
	}

	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return sdk.Wrap(err)
	}
	name := strings.TrimSpace(m.Name)
	if len(name) == 0 {
		return sdk.Wrapf("missing name")
	}
	return nil
}

func (m MsgIssueDenom) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgIssueDenom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgEditMT) Route() string {
	return ModuleName
}

func (m MsgEditMT) Type() string {
	return "edit_mt"
}

func (m MsgEditMT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return sdk.Wrap(err)
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return sdk.Wrapf("missing denom")
	}

	tokenID := strings.TrimSpace(m.Id)
	if len(tokenID) == 0 {
		return sdk.Wrapf("missing mtID")
	}
	return nil
}

func (m MsgEditMT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgEditMT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgTransferMT) Route() string {
	return ModuleName
}

func (m MsgTransferMT) Type() string {
	return "transfer_mt"
}

func (m MsgTransferMT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return sdk.Wrap(err)
	}

	if len(m.Recipient) == 0 {
		return sdk.Wrapf("missing recipient address")
	}
	if err := sdk.ValidateAccAddress(m.Recipient); err != nil {
		return sdk.Wrap(err)
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return sdk.Wrapf("missing denom")
	}

	tokenID := strings.TrimSpace(m.Id)
	if len(tokenID) == 0 {
		return sdk.Wrapf("missing mtID")
	}

	if m.Amount <= 0 {
		return sdk.Wrapf("invalid amount")
	}
	return nil
}

func (m MsgTransferMT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgTransferMT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgBurnMT) Route() string {
	return ModuleName
}

func (m MsgBurnMT) Type() string {
	return "burn_mt"
}

func (m MsgBurnMT) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return sdk.Wrapf("missing sender address")
	}
	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return sdk.Wrap(err)
	}

	denom := strings.TrimSpace(m.DenomId)
	if len(denom) == 0 {
		return sdk.Wrapf("missing denom")
	}

	tokenID := strings.TrimSpace(m.Id)
	if len(tokenID) == 0 {
		return sdk.Wrapf("missing ID")
	}
	return nil
}

func (m MsgBurnMT) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgBurnMT) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
}

func (m MsgTransferDenom) Route() string {
	return ModuleName
}

func (m MsgTransferDenom) Type() string {
	return "transfer_denom"
}

func (m MsgTransferDenom) ValidateBasic() error {
	if len(m.Sender) == 0 {
		return sdk.Wrapf("missing sender address")
	}

	if err := sdk.ValidateAccAddress(m.Sender); err != nil {
		return sdk.Wrap(err)
	}
	id := strings.TrimSpace(m.Id)
	if len(id) == 0 {
		return sdk.Wrapf("missing id")
	}

	if len(m.Recipient) == 0 {
		return sdk.Wrapf("missing recipient address")
	}
	if err := sdk.ValidateAccAddress(m.Recipient); err != nil {
		return sdk.Wrap(err)
	}
	return nil
}

func (m MsgTransferDenom) GetSignBytes() []byte {
	bz, err := ModuleCdc.MarshalJSON(&m)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}

func (m MsgTransferDenom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Sender)}
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

func (this Balance) Convert() interface{} {
	return QueryBalanceResp{
		MtId:   this.MtId,
		Amount: this.Amount,
	}
}

type balances []Balance

func (this balances) Convert() interface{} {
	var balances []QueryBalanceResp
	for _, balance := range this {
		balances = append(balances, balance.Convert().(QueryBalanceResp))
	}
	return balances
}
