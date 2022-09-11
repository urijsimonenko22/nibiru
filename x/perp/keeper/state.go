package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/perp/types"
)

func (k Keeper) WhitelistState(ctx sdk.Context) WhitelistState {
	return newWhitelist(ctx, k.storeKey, k.cdc)
}

func (k Keeper) PrepaidBadDebtState(ctx sdk.Context) PrepaidBadDebtState {
	return newPrepaidBadDebtState(ctx, k.storeKey, k.cdc)
}

// getLatestCumulativePremiumFraction returns the last cumulative premium fraction recorded for the
// specific pair.
func (k Keeper) getLatestCumulativePremiumFraction(
	ctx sdk.Context, pair common.AssetPair,
) (sdk.Dec, error) {
	pairMetadata, err := k.PairMetadata.Get(ctx, pair)
	if err != nil {
		k.Logger(ctx).Error(
			err.Error(),
			"pair",
			pair.String(),
		)
		return sdk.Dec{}, types.ErrPairMetadataNotFound
	}
	// this should never fail
	return pairMetadata.CumulativePremiumFractions[len(pairMetadata.CumulativePremiumFractions)-1], nil
}

var whitelistNamespace = []byte{0x3}

type WhitelistState struct {
	whitelists sdk.KVStore
	cdc        codec.BinaryCodec
}

func newWhitelist(ctx sdk.Context, key sdk.StoreKey, cdc codec.BinaryCodec) WhitelistState {
	return WhitelistState{
		whitelists: prefix.NewStore(ctx.KVStore(key), whitelistNamespace),
		cdc:        cdc,
	}
}

func (w WhitelistState) IsWhitelisted(address sdk.AccAddress) bool {
	return w.whitelists.Has(address)
}

func (w WhitelistState) Add(address sdk.AccAddress) {
	w.whitelists.Set(address, []byte{})
}

func (w WhitelistState) Iterate(do func(addr sdk.AccAddress) (stop bool)) {
	iter := w.whitelists.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		if !do(iter.Key()) {
			break
		}
	}
}

var prepaidBadDebtNamespace = []byte{0x4}

type PrepaidBadDebtState struct {
	prepaidBadDebt sdk.KVStore
}

func newPrepaidBadDebtState(ctx sdk.Context, key sdk.StoreKey, _ codec.BinaryCodec) PrepaidBadDebtState {
	return PrepaidBadDebtState{
		prepaidBadDebt: prefix.NewStore(ctx.KVStore(key), prepaidBadDebtNamespace),
	}
}

// Get Fetches the amount of bad debt prepaid by denom. Returns zero if the denom is not found.
func (s PrepaidBadDebtState) Get(denom string) (amount sdk.Int) {
	v := s.prepaidBadDebt.Get([]byte(denom))
	if v == nil {
		return sdk.ZeroInt()
	}

	err := amount.Unmarshal(v)
	if err != nil {
		panic(err)
	}

	return amount
}

// Iterate iterates over known prepaid bad debt.
func (s PrepaidBadDebtState) Iterate(do func(denom string, amount sdk.Int) (stop bool)) {
	iter := s.prepaidBadDebt.Iterator(nil, nil)

	for ; iter.Valid(); iter.Next() {
		amount := sdk.Int{}
		err := amount.Unmarshal(iter.Value())
		if err != nil {
			panic(err)
		}
		if !do(string(iter.Key()), amount) {
			break
		}
	}
}

// Set sets the amount of bad debt prepaid by denom.
func (s PrepaidBadDebtState) Set(denom string, amount sdk.Int) {
	b, err := amount.Marshal()
	if err != nil {
		panic(err)
	}
	s.prepaidBadDebt.Set([]byte(denom), b)
}

// Increment increments the amount of bad debt prepaid by denom.
// Calling this method on a denom that doesn't exist is effectively the same as setting the amount (0 + increment).
func (s PrepaidBadDebtState) Increment(denom string, increment sdk.Int) (amount sdk.Int) {
	amount = s.Get(denom).Add(increment)

	b, err := amount.Marshal()
	if err != nil {
		panic(err)
	}
	s.prepaidBadDebt.Set([]byte(denom), b)

	return amount
}

// Decrement decrements the amount of bad debt prepaid by denom.
// The lowest it can be decremented to is zero. Trying to decrement a prepaid bad
// debt balance to below zero will clip it at zero.
func (s PrepaidBadDebtState) Decrement(denom string, decrement sdk.Int) (amount sdk.Int) {
	amount = sdk.MaxInt(s.Get(denom).Sub(decrement), sdk.ZeroInt())

	b, err := amount.Marshal()
	if err != nil {
		panic(err)
	}
	s.prepaidBadDebt.Set([]byte(denom), b)

	return amount
}
