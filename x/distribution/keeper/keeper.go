package keeper

import (
	"context"

	"github.com/alexanderbez/novestingyield/x/distribution/types"

	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	dstrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	dstrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

type Keeper struct {
	dstrkeeper.Keeper

	authKeeper dstrtypes.AccountKeeper
}

// NewKeeper creates a new distribution Keeper instance, embedding or wrapping
// a real distribution keeper.
//
// Typically, we don't need to do this but since the original distribution keeper
// doesn't expose private fields that we need access to, we "hijack" them by
// re-creating a constructor to capture those private fields before passing them
// onto the real constructor.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	ak dstrtypes.AccountKeeper,
	bk dstrtypes.BankKeeper,
	sk dstrtypes.StakingKeeper,
	feeCollectorName, authority string,
) Keeper {
	return Keeper{
		Keeper:     dstrkeeper.NewKeeper(cdc, storeService, ak, bk, sk, feeCollectorName, authority),
		authKeeper: ak,
	}
}

// WithdrawDelegationRewards override the original function to "extend" functionality,
// by prohibiting withdrawal of rewards if the delegation account is a vesting
// account.
//
// Specifically, we prohibit withdrawal of rewards if the delegation account is
// has ANY tokens left that are still escrowed/vesting.
//
// Note, this is just an example of how a chain can prohibit withdrawal of rewards
// for vesting accounts. A more sophisticated implementation could allow for
// withdrawal up to a certain threshold or perhaps based on a function of total
// vested so far. Regardless, the example shows how any such mechanism can be
// implemented.
func (k Keeper) WithdrawDelegationRewards(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error) {
	delegator := k.authKeeper.GetAccount(ctx, delAddr)

	vacc, ok := delegator.(vesting.VestingAccount)
	if ok {
		sdkCtx := sdk.UnwrapSDKContext(ctx)

		if !vacc.GetVestingCoins(sdkCtx.BlockTime()).IsZero() {
			return nil, types.ErrForbiddenWithdrawal.Wrapf("vesting account %s cannot withdraw rewards", delAddr)
		}
	}

	return k.WithdrawDelegationRewards(ctx, delAddr, valAddr)
}
