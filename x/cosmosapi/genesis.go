package cosmosapi

import (
    "errors"
    sdk "github.com/cosmos/cosmos-sdk/types"
    abci "github.com/tendermint/tendermint/abci/types"
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)

func ValidateGenesis(data GenesisState) error {
    adminAddresses := data.AdminAddresses
    if len(adminAddresses) < 1 {
        return errors.New("At least one admin address is needed")
    }

    //for _, address := range data {
    // TODO: validate address
    //}
    return nil
}

func DefaultGenesisState() GenesisState {
    return types.GenesisState{} 
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
    keeper.CreateGenesisAdminGroup(ctx, data)
    return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
// TODO: update the following after implementing k.GetPollsIterator(ctx)
//	iterator := k.GetNamesIterator(ctx)
//	for ; iterator.Valid(); iterator.Next() {
//
//		name := string(iterator.Key())
//		whois := k.GetWhois(ctx, name)
//		records = append(records, whois)
//
//	}
    return types.GenesisState{}
}
