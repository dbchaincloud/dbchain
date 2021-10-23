package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
    AdminAddresses []sdk.AccAddress `json:"admin_addresses"`
}

func NewGenesisState(adminAddresses []sdk.AccAddress) GenesisState {
    return GenesisState{AdminAddresses: adminAddresses}
}

