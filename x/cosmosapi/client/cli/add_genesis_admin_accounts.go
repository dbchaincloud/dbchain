package cli

import (
    "fmt"
    "bytes"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/tendermint/tendermint/libs/cli"
    "github.com/cosmos/cosmos-sdk/client/keys"

    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/server"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/x/genutil"

    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)

const (
    flagClientHome   = "home-client"
)

func AddGenesisAdminAccountCmd(ctx *server.Context, cdc *codec.Codec,
    defaultNodeHome, defaultClientHome string) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "add-genesis-admin-account [address_or_key_name]",
        Short: "Add a genesis admin account to genesis.json",
        Args:  cobra.ExactArgs(1),
        RunE: func(_ *cobra.Command, args []string) error {
            config := ctx.Config
            config.SetRoot(viper.GetString(cli.HomeFlag))
            addr, err := sdk.AccAddressFromBech32(args[0])
            if err != nil {
                kb, err := keys.NewKeyBaseFromDir(viper.GetString(flagClientHome))
                if err != nil {
                    return err
                }

                info, err := kb.Get(args[0])
                if err != nil {
                    return err
                }

                addr = info.GetAddress()
            }

            // retrieve the app state
            genFile := config.GenesisFile()
            appState, genDoc, err := genutil.GenesisStateFromGenFile(cdc, genFile)
            if err != nil {
                return err
            }

            var genesisState types.GenesisState

            cdc.MustUnmarshalJSON(appState[types.ModuleName], &genesisState)

            for _, address := range genesisState.AdminAddresses {
                if bytes.Compare(address, addr) == 0 {
                    return fmt.Errorf("cannot add account at existing address %v", addr)
                }
            }

            adminAddresses := append(genesisState.AdminAddresses, addr)

            genesisStateBz := cdc.MustMarshalJSON(types.NewGenesisState(adminAddresses))
            appState[types.ModuleName] = genesisStateBz

            appStateJSON, err := cdc.MarshalJSON(appState)
            if err != nil {
                return err
            }

            // export app state
            genDoc.AppState = appStateJSON

            return genutil.ExportGenesisFile(genDoc, genFile)
        },
    }
    return cmd
}


