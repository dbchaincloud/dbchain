package cosmosapi

import (
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/keeper"
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)

const (
    ModuleName = types.ModuleName
    RouterKey  = types.RouterKey
    StoreKey   = types.StoreKey
)

var (
    NewKeeper        = keeper.NewKeeper
    NewQuerier       = keeper.NewQuerier
    ModuleCdc        = types.ModuleCdc
    RegisterCodec    = types.RegisterCodec
)

type (
    Keeper          = keeper.Keeper
    MsgCreateTable  = types.MsgCreateTable
    MsgCreateIndex  = types.MsgCreateIndex
    MsgInsertRow    = types.MsgInsertRow

    MsgAddAdminAccount = types.MsgAddAdminAccount

    Table           = types.Table

    GenesisState    = types.GenesisState
)
