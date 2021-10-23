package types

import (
    "github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
    RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
    cdc.RegisterConcrete(MsgCreateTable{}, "cosmosapi/CreateTable", nil)
    cdc.RegisterConcrete(MsgCreateIndex{}, "cosmosapi/CreateIndex", nil)
    cdc.RegisterConcrete(MsgInsertRow{}, "cosmosapi/InsertRow", nil)
    cdc.RegisterConcrete(MsgAddAdminAccount{}, "cosmosapi/AddAdminAccount", nil)
}

