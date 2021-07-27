package farm

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"strings"
)

const (
	ModuleName = "farm"
)

var (
	_ sdk.Msg = &MsgCreatePool{}
	_ sdk.Msg = &MsgDestroyPool{}
	_ sdk.Msg = &MsgAdjustPool{}
	_ sdk.Msg = &MsgStake{}
	_ sdk.Msg = &MsgUnstake{}
	_ sdk.Msg = &MsgHarvest{}
)

func (msg MsgCreatePool) Route() string {
	return ModuleName
}

func (msg MsgCreatePool) Type() string {
	return "create_pool"
}

func (msg MsgCreatePool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdk.Wrapf("invalid creator")
	}

	feedName := strings.TrimSpace(msg.Name)
	if len(feedName) == 0 {
		return sdk.Wrapf("missing feed name")
	}

	return nil
}

func (msg MsgCreatePool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg MsgDestroyPool) Route() string {
	return ModuleName
}

func (msg MsgDestroyPool) Type() string {
	return "destroy_pool"
}

func (msg MsgDestroyPool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdk.Wrapf("invalid creator")
	}

	PoolName := strings.TrimSpace(msg.PoolName)
	if len(PoolName) == 0 {
		return sdk.Wrapf("missing pool name")
	}

	return nil
}

func (msg MsgDestroyPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgDestroyPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg MsgAdjustPool) Route() string {
	return ModuleName
}

func (msg MsgAdjustPool) Type() string {
	return "adjust_pool"
}

func (msg MsgAdjustPool) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdk.Wrapf("invalid creator")
	}

	PoolName := strings.TrimSpace(msg.PoolName)
	if len(PoolName) == 0 {
		return sdk.Wrapf("missing pool name")
	}

	return nil
}

func (msg MsgAdjustPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgAdjustPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg MsgStake) Route() string {
	return ModuleName
}

func (msg MsgStake) Type() string {
	return "stake"
}

func (msg MsgStake) ValidateBasic() error {

	PoolName := strings.TrimSpace(msg.PoolName)
	if len(PoolName) == 0 {
		return sdk.Wrapf("missing pool name")
	}

	return nil
}

func (msg MsgStake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgStake) GetSigners() []sdk.AccAddress {
	PoolName, err := sdk.AccAddressFromBech32(msg.PoolName)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{PoolName}
}

func (msg MsgUnstake) Route() string {
	return ModuleName
}

func (msg MsgUnstake) Type() string {
	return "unstake"
}

func (msg MsgUnstake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.PoolName); err != nil {
		return sdk.Wrapf("invalid creator")
	}

	PoolName := strings.TrimSpace(msg.PoolName)
	if len(PoolName) == 0 {
		return sdk.Wrapf("missing pool name")
	}

	return nil
}

func (msg MsgUnstake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgUnstake) GetSigners() []sdk.AccAddress {
	PoolName, err := sdk.AccAddressFromBech32(msg.PoolName)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{PoolName}
}

func (msg MsgHarvest) Route() string {
	return ModuleName
}

func (msg MsgHarvest) Type() string {
	return "harvest"
}

func (msg MsgHarvest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.PoolName); err != nil {
		return sdk.Wrapf("invalid creator")
	}

	PoolName := strings.TrimSpace(msg.PoolName)
	if len(PoolName) == 0 {
		return sdk.Wrapf("missing pool name")
	}

	return nil
}

func (msg MsgHarvest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgHarvest) GetSigners() []sdk.AccAddress {
	PoolName, err := sdk.AccAddressFromBech32(msg.PoolName)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{PoolName}
}
