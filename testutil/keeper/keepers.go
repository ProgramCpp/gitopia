package keeper

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gitopia/gitopia/app"
	"github.com/gitopia/gitopia/app/keepers"
	"github.com/gitopia/gitopia/x/gitopia/keeper"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func AppKeepers(t testing.TB) (keepers.AppKeepers, sdk.Context) {
	appKeepers := keepers.AppKeepers{}
	logger := log.NewNopLogger()

	appKeepers.GenerateKeys()

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	keys := appKeepers.GetKVStoreKey()
	mkeys := appKeepers.GetMemoryStoreKey()
	tkeys := appKeepers.GetTransientStoreKey()
	for _, key := range keys {
		stateStore.MountStoreWithDB(key, storetypes.StoreTypeIAVL, nil)
	}

	for _, key := range mkeys {
		stateStore.MountStoreWithDB(key, storetypes.StoreTypeMemory, db)
	}

	for _, key := range tkeys {
		stateStore.MountStoreWithDB(key, storetypes.StoreTypeTransient, nil)
	}

	require.NoError(t, stateStore.LoadLatestVersion())

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger).WithBlockTime(time.Now())

	registry := codectypes.NewInterfaceRegistry()
	encConfig := app.MakeEncodingConfig()
	appCodec := encConfig.Codec

	amino := codec.NewLegacyAmino()

	appKeepers.ParamsKeeper = paramskeeper.NewKeeper(
		appCodec, amino, keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey])

	appKeepers.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		appKeepers.ParamsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		"gitopia",
	)

	authzkeeper := authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		nil,
		appKeepers.AccountKeeper,
	)
	appKeepers.AuthzKeeper = &authzkeeper

	appKeepers.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.ParamsKeeper.Subspace(banktypes.ModuleName),
		nil,
	)
	appKeepers.BankKeeper.SetParams(ctx, banktypes.DefaultParams())

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.ParamsKeeper.Subspace(stakingtypes.ModuleName),
	)

	mintKeeper := mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey],
		appKeepers.ParamsKeeper.Subspace(minttypes.ModuleName),
		stakingKeeper, appKeepers.AccountKeeper,
		appKeepers.BankKeeper, types.MinterAccountName)

	appKeepers.GitopiaKeeper = *keeper.NewKeeper(
		codec.NewProtoCodec(registry),
		keys[types.StoreKey],
		keys[types.MemStoreKey],
		types.MinterAccountName,
		authtypes.FeeCollectorName,
		appKeepers.AccountKeeper,
		appKeepers.AuthzKeeper,
		appKeepers.BankKeeper,
		mintKeeper,
	)

	return appKeepers, ctx
}
