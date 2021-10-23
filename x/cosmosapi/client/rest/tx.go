package rest

import (
    "net/http"

    "github.com/cosmos/cosmos-sdk/client/context"
    "github.com/yzhanginwa/cosmos-api/x/cosmosapi/internal/types"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/rest"
    "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type createTableReq struct  {
    BaseReq rest.BaseReq   `json:"base_req"`
    Owner   string         `json:"owner"`
    Name    string         `json:"title"`
    Fields  []string       `json:"fields"`
}

func createTableHandler(cliCtx context.CLIContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req createTableReq

        if !rest.ReadRESTReq(w, r, cliCtx.Codec,  &req)  {
            rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
            return
        }

        baseReq := req.BaseReq.Sanitize()
        if !baseReq.ValidateBasic(w) {
            return
        }

        addr, err := sdk.AccAddressFromBech32(req.Owner)
        if err != nil {
            rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
            return
        }

        msg := types.NewMsgCreateTable(addr, req.Name, req.Fields)
        err = msg.ValidateBasic()
        if err != nil {
            rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
            return
        }

        utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
    }
}

