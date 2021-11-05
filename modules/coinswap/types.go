package coinswap

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
)

const (
	ModuleName = "coinswap"

	eventTypeTransfer = "transfer"
	eventTypeSwap     = "swap"

	attributeKeyAmount = "amount"
)

var (
	_ sdk.Msg = &MsgAddLiquidity{}
	_ sdk.Msg = &MsgRemoveLiquidity{}
	_ sdk.Msg = &MsgSwapOrder{}
)

type totalSupply = func() (sdk.Coins, error)

// ValidateBasic implements Msg.
func (msg MsgAddLiquidity) ValidateBasic() error {
	if !(msg.MaxToken.IsValid() && msg.MaxToken.IsPositive()) {
		return sdkerrors.Wrapf(ErrTokenCount, "invalid MaxToken: %s", msg.MaxToken.String())
	}

	if !msg.ExactStandardAmt.IsPositive() {
		return sdkerrors.Wrapf(ErrTokenCount, "standard token amount must be positive")
	}

	if msg.MinLiquidity.IsNegative() {
		return sdkerrors.Wrapf(ErrTokenCount, "minimum liquidity can not be negative")
	}

	if msg.Deadline <= 0 {
		return sdkerrors.Wrapf(ErrDeadline, "deadline %d must be greater than 0", msg.Deadline)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(ErrAccAddressFromBech32, err.Error())
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgAddLiquidity) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implements Msg.
func (msg MsgRemoveLiquidity) ValidateBasic() error {
	if msg.MinToken.IsNegative() {
		return sdkerrors.Wrapf(ErrTokenCount, "minimum token amount can not be negative")
	}
	if !msg.WithdrawLiquidity.IsValid() || !msg.WithdrawLiquidity.IsPositive() {
		return sdkerrors.Wrapf(ErrTokenCount, "invalid withdrawLiquidity (%s)", msg.WithdrawLiquidity.String())
	}
	if msg.MinStandardAmt.IsNegative() {
		return sdkerrors.Wrapf(ErrTokenCount, "minimum standard token amount %s can not be negative", msg.MinStandardAmt.String())
	}
	if msg.Deadline <= 0 {
		return sdkerrors.Wrapf(ErrDeadline, "deadline %d must be greater than 0", msg.Deadline)
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgRemoveLiquidity) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Sender)}
}

// ValidateBasic implements Msg.
func (msg MsgSwapOrder) ValidateBasic() error {
	if !(msg.Input.Coin.IsValid() && msg.Input.Coin.IsPositive()) {
		return sdkerrors.Wrapf(ErrTokenCount, "invalid input (%s)", msg.Input.Coin.String())
	}

	if _, err := sdk.AccAddressFromBech32(msg.Input.Address); err != nil {
		return sdkerrors.Wrapf(ErrAccAddressFromBech32, err.Error())
	}

	if !(msg.Output.Coin.IsValid() && msg.Output.Coin.IsPositive()) {
		return sdkerrors.Wrapf(ErrTokenCount, "invalid output (%s)", msg.Output.Coin.String())
	}

	if _, err := sdk.AccAddressFromBech32(msg.Output.Address); err != nil {
		return sdkerrors.Wrapf(ErrAccAddressFromBech32, err.Error())
	}

	if msg.Input.Coin.Denom == msg.Output.Coin.Denom {
		return sdkerrors.Wrapf(ErrDenom, "invalid swap")
	}

	if msg.Deadline <= 0 {
		return sdkerrors.Wrapf(ErrDeadline, "deadline %d must be greater than 0", msg.Deadline)
	}
	return nil
}

// GetSigners implements Msg.
func (msg MsgSwapOrder) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Input.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (m QueryLiquidityPoolResponse) Convert() interface{} {
	return &QueryPoolResponse{
		Pool: _loadPoolInfo(m.Pool),
	}
}

func (m QueryLiquidityPoolsResponse) Convert() interface{} {

	return &QueryAllPoolsResponse{
		Pagination: m.Pagination,
		Pools:      _loadPools(m.Pools),
	}
}

func _loadPoolInfo(info PoolInfo) PoolInfo {
	return PoolInfo{
		Id:            info.Id,
		EscrowAddress: info.EscrowAddress,
		Standard:      info.Standard,
		Token:         info.Token,
		Lpt:           info.Lpt,
		Fee:           info.Fee,
	}
}
func _loadPools(pools []PoolInfo) (ret []PoolInfo) {
	for _, pool := range pools {
		ret = append(ret, _loadPoolInfo(pool))
	}
	return
}
