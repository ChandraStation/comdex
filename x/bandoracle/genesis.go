package bandoracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/comdex-official/comdex/x/bandoracle/keeper"
	"github.com/comdex-official/comdex/x/bandoracle/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, state *types.GenesisState) {

	for _, item := range state.Markets {
		k.SetMarket(ctx, item)
	}

	k.SetParams(ctx, state.Params)
	k.SetPort(ctx, state.PortId)
	//portId:= k.IBCPort(ctx)

	if !k.IsBound(ctx, state.PortId) {

		err := k.BindPort(ctx, state.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(
		k.GetMarkets(ctx),
		k.GetParams(ctx),
		k.GetPort(ctx),
	)
}