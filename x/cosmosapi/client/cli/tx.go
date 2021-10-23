package cli

import (
    "fmt"
    "errors"
    "strings"
    "encoding/json"
    "github.com/spf13/cobra"

    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/client/context"
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/x/auth"
    "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
    cosmosapiTxCmd := &cobra.Command{
        Use:                        types.ModuleName,
        Short:                      "Cosmosapi transaction subcommands",
        DisableFlagParsing:         true,
        SuggestionsMinimumDistance: 2,
        RunE:                       client.ValidateCmd,
    }

    cosmosapiTxCmd.AddCommand(client.PostCommands(
        GetCmdCreateTable(cdc),
        GetCmdCreateIndex(cdc),
        GetCmdInsertRow(cdc),
        GetCmdAddAdminAccount(cdc),
    )...)

    return cosmosapiTxCmd
}

// GetCmdCreatePoll is the CLI command for sending a CreatePoll transaction
func GetCmdCreateTable(cdc *codec.Codec) *cobra.Command {
    return &cobra.Command{
        Use:   "create-table [name] [fields]",
        Short: "create a new table",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            cliCtx := context.NewCLIContext().WithCodec(cdc)
            txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

            name := args[0]
            fields := strings.Split(args[1], ",")
            msg := types.NewMsgCreateTable(cliCtx.GetFromAddress(), name, fields)
            err := msg.ValidateBasic()
            if err != nil {
                return err
            }

            return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
        },
    }
}

// GetCmdCreateIndex is the CLI command for sending a CreateIndex transaction
func GetCmdCreateIndex(cdc *codec.Codec) *cobra.Command {
    return &cobra.Command{
        Use:   "create-index [tableName] [field]",
        Short: "create a new index",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            cliCtx := context.NewCLIContext().WithCodec(cdc)
            txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

            tableName := args[0]
            field := args[1]
            msg := types.NewMsgCreateIndex(cliCtx.GetFromAddress(), tableName, field)
            err := msg.ValidateBasic()
            if err != nil {
                return err
            }

            return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
        },
    }
}


func GetCmdInsertRow(cdc *codec.Codec) *cobra.Command {
    return &cobra.Command{
        Use:   "insert-row [tableName] [fields] [values]",
        Short: "create a new row",
        Args:  cobra.ExactArgs(3),
        RunE: func(cmd *cobra.Command, args []string) error {
            cliCtx := context.NewCLIContext().WithCodec(cdc)
            txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

            name := args[0]
            fields := strings.Split(args[1], ",")
            values := strings.Split(args[2], ",")
            rowFields := make(types.RowFields)
            for i, field := range fields {
                if i < len(values) {
                    rowFields[field] = values[i]
                }
            }

            rowFieldsJson, err := json.Marshal(rowFields)
            if err != nil { return err } 

            msg := types.NewMsgInsertRow(cliCtx.GetFromAddress(), name, rowFieldsJson)
            err = msg.ValidateBasic()
            if err != nil {
                return errors.New(fmt.Sprintf("Error %s", err)) 
            }

            return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
        },
    }
}


func GetCmdAddAdminAccount(cdc * codec.Codec) *cobra.Command {
    return &cobra.Command{
        Use:   "add-admin [address]",
        Short: "add an account into admin group",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            cliCtx := context.NewCLIContext().WithCodec(cdc)
            txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

            addr, err := sdk.AccAddressFromBech32(args[0]) // args[0] is the new admin address

            if err != nil {
                return errors.New(fmt.Sprintf("Error %s", err))
            }

            msg := types.NewMsgAddAdminAccount(addr, cliCtx.GetFromAddress())
            err = msg.ValidateBasic()
            if err != nil {
                return errors.New(fmt.Sprintf("Error %s", err))
            }
            return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
        },
    }
}



//////////////////////
//                  //
// helper functions //
//                  //
//////////////////////

