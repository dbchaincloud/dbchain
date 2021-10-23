package keeper

import (
    "os"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/x/bank"
    "github.com/tendermint/tendermint/libs/log"
)

var (
    logger = defaultLogger()
)

func defaultLogger() log.Logger {
    return log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("ethan1", "ethan2")
}

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
    CoinKeeper bank.Keeper

    storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

    cdc *codec.Codec // The wire codec for binary encoding/decoding.
}


// NewKeeper creates new instances of the cosmosapi Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
    return Keeper{
        CoinKeeper: coinKeeper,
        storeKey:   storeKey,
        cdc:        cdc,
    }
}


//////////////////////
//                  //
// helper functions //
//                  //
//////////////////////

