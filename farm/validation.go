package farm

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// MaxDescriptionLength length of the service and author description
	MaxDescriptionLength = 280
)

// ValidatepPoolId validates the pool id
func ValidatepPoolId(poolId string) (uint64, error) {
	seqStr := strings.TrimPrefix(poolId, PrefixFarmPool+"-")
	seq, err := strconv.ParseUint(seqStr, 10, 64)
	if err != nil || seq == 0 {
		return 0, ErrInvalidPoolId
	}
	return seq, nil
}

// ValidateDescription validates the pool name
func ValidateDescription(description string) error {
	if len(description) > MaxDescriptionLength {
		return sdkerrors.Wrap(ErrInvalidDescription, description)
	}
	return nil
}

// ValidateLpTokenDenom validates the lp token denom
func ValidateLpTokenDenom(denom string) error {
	return sdk.ValidateDenom(denom)
}

// ValidateCoins validates the coin
func ValidateCoins(field string, coins ...sdk.Coin) error {
	if len(coins) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "The %s should be greater than zero", field)
	}
	return sdk.NewCoins(coins...).Validate()
}

// ValidateAddress validates the address
func ValidateAddress(sender string) error {
	_, err := sdk.AccAddressFromBech32(sender)
	return err
}

// ValidateReward validates the coin
func ValidateReward(rewardPerBlock, totalReward sdk.Coins) error {
	if len(rewardPerBlock) != len(totalReward) {
		return sdkerrors.Wrapf(ErrNotMatch, "The length of rewardPerBlock and totalReward must be the same")
	}

	if !rewardPerBlock.DenomsSubsetOf(totalReward) {
		return sdkerrors.Wrapf(ErrInvalidRewardRule, "rewardPerBlock and totalReward token types must be the same")
	}

	for i := range totalReward {
		if !totalReward[i].IsGTE(rewardPerBlock[i]) {
			return sdkerrors.Wrapf(ErrNotMatch, "The totalReward should be greater than or equal to rewardPerBlock")
		}
		//uint64 overflow check
		h := totalReward[i].Amount.Quo(rewardPerBlock[i].Amount)
		if !h.IsInt64() {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Can not convert to int64, overflow")
		}
	}
	return nil
}

func ValidateFund(rewardPerBlock, fundApplied, fundSelfBond []sdk.Coin) error {
	if err := ValidateCoins("FundApplied", fundApplied...); err != nil {
		return err
	}

	if err := sdk.NewCoins(fundSelfBond...).Validate(); err != nil {
		return sdkerrors.Wrapf(err, "The fundSelfBond is invalid coin")
	}

	total := sdk.NewCoins(fundApplied...).Add(fundSelfBond...)
	if len(fundApplied)+len(fundSelfBond) != total.Len() {
		return sdkerrors.Wrapf(ErrInvalidProposal, "the type of The token bond by the user cannot be the same as the one applied for community pool")
	}
	return ValidateReward(rewardPerBlock, total)
}
