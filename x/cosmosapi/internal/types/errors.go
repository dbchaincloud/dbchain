package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultCodespace is the Module Name
const (
    DefaultCodespace sdk.CodespaceType = ModuleName

    CodeWrongBallot sdk.CodeType = 101
)

// ErrNameDoesNotExist is the error for name not existing
func ErrWrongBallot(codespace sdk.CodespaceType) sdk.Error {
    return sdk.NewError(codespace, CodeWrongBallot, "Wrong ballot")
}
