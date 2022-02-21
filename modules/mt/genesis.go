package mt

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/irisnet/irismod/modules/mt/keeper"
	"github.com/irisnet/irismod/modules/mt/types"
)

// InitGenesis stores the MT genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	// ---------- init infos ---------- //

	// set denom sequence
	k.SetDenomSequence(ctx, uint64(len(data.Collections)+1))

	var mtSequence uint64 = 1
	for _, c := range data.Collections {
		// store denom
		k.SetDenom(ctx, *c.Denom)

		for _, m := range c.Mts {
			// store mt
			k.SetMT(ctx, c.Denom.Id, m)
			mtSequence++
		}
	}

	// set mt sequence
	k.SetMTSequence(ctx, mtSequence)

	// ---------- init balances ---------- //

	for _, o := range data.Owners {
		addr, err := sdk.AccAddressFromBech32(o.Address)
		if err != nil {
			panic(sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", addr))
		}

		for _, d := range o.Denoms {
			for _, b := range d.Balances {
				k.AddBalance(ctx, d.DenomId, b.MtId, b.Amount, addr)
			}
		}
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	//return types.NewGenesisState(k.GetCollections(ctx))
	return nil
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.Collection{}, []types.Owner{})
}
