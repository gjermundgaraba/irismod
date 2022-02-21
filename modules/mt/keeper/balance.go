package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/mt/types"
)

// AddBalance adds amounts to an account
func (k Keeper) AddBalance(ctx sdk.Context,
	denomID, mtID string,
	amount uint64,
	addr sdk.AccAddress) {

	balance := k.getBalance(ctx, denomID, mtID, addr)
	balance += amount

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalAmount(k.cdc, balance)
	store.Set(types.KeyBalance(addr, denomID, mtID), bz)
}

func (k Keeper) subBalance(ctx sdk.Context,
	denomID, mtID string,
	amount uint64,
	addr sdk.AccAddress) {

	store := ctx.KVStore(k.storeKey)
	balance := k.getBalance(ctx, denomID, mtID, addr)
	balance -= amount

	bz := types.MustMarshalAmount(k.cdc, balance)
	store.Set(types.KeyBalance(addr, denomID, mtID), bz)
}

func (k Keeper) getBalance(ctx sdk.Context,
	denomID, mtID string,
	addr sdk.AccAddress) uint64 {

	store := ctx.KVStore(k.storeKey)

	amount := store.Get(types.KeyBalance(addr, denomID, mtID))
	if len(amount) == 0 {
		return 0
	}

	return types.MustUnMarshalAmount(k.cdc, amount)
}

func (k Keeper) transfer(ctx sdk.Context,
	denomID, mtID string,
	amount uint64,
	from, to sdk.AccAddress) {

	k.subBalance(ctx, denomID, mtID, amount, from)

	k.AddBalance(ctx, denomID, mtID, amount, to)
}

// GetDenomSupply returns the number of Mts by the specified denom ID
func (k Keeper) GetDenomSupply(ctx sdk.Context, denomID string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeySupply(denomID, ""))
	if len(bz) == 0 {
		return 0
	}
	return types.MustUnMarshalSupply(k.cdc, bz)
}

// GetMTSupply returns the supply of a specified MT
func (k Keeper) GetMTSupply(ctx sdk.Context, denomID, mtID string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeySupply(denomID, mtID))
	if len(bz) == 0 {
		return 0
	}
	return types.MustUnMarshalSupply(k.cdc, bz)
}

// increaseDenomSupply increase total supply (count of MTs) of a denom
func (k Keeper) increaseDenomSupply(ctx sdk.Context, denomID string) {
	supply := k.GetDenomSupply(ctx, denomID)
	supply++

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeySupply(denomID, ""), bz)
}

// increaseMTSupply increase total supply of an MT
func (k Keeper) increaseMTSupply(ctx sdk.Context, denomID, mtID string, amount uint64) {
	supply := k.GetMTSupply(ctx, denomID, mtID)
	supply += amount

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeySupply(denomID, mtID), bz)
}

// decreaseMTSupply decrease total supply of an MT
func (k Keeper) decreaseMTSupply(ctx sdk.Context, denomID, mtID string, amount uint64) {
	supply := k.GetMTSupply(ctx, denomID, mtID)
	supply -= amount

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalSupply(k.cdc, supply)
	store.Set(types.KeySupply(denomID, mtID), bz)
}
