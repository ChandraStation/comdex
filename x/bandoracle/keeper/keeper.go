package keeper

import (
	"context"
	"fmt"
	assetkeeper "github.com/comdex-official/comdex/x/asset/keeper"
	"github.com/comdex-official/comdex/x/bandoracle/expected"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/comdex-official/comdex/x/bandoracle/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/spm/ibckeeper"
)

type (
	Keeper struct {
		*ibckeeper.Keeper
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		scoped      expected.ScopedKeeper
		assetKeeper assetkeeper.Keeper
	}
)

func (k *Keeper) QueryMarkets(ctx context.Context, request *types.QueryMarketsRequest) (*types.QueryMarketsResponse, error) {
	panic("implement me")
}

func (k *Keeper) QueryMarket(ctx context.Context, request *types.QueryMarketRequest) (*types.QueryMarketResponse, error) {
	panic("implement me")
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper ibckeeper.ChannelKeeper,
	portKeeper ibckeeper.PortKeeper,
	scopedKeeper ibckeeper.ScopedKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		Keeper: ibckeeper.NewKeeper(
			types.PortKey,
			storeKey,
			channelKeeper,
			portKeeper,
			scopedKeeper,
		),
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}
