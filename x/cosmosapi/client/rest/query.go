package rest

import (
    "fmt"
    "net/http"

    "github.com/cosmos/cosmos-sdk/client/context"

    "github.com/cosmos/cosmos-sdk/types/rest"

    "github.com/gorilla/mux"
)

func showTablesHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/tables", storeName), nil)
        if err != nil {
            rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
            return
        }
        rest.PostProcessResponse(w, cliCtx, res)
    }
}

func showTableHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/tables/%s", storeName, vars["name"]), nil)
        if err != nil {
            rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
            return
        }
        rest.PostProcessResponse(w, cliCtx, res)
    }
}

func showRowHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/find/%s/%s", storeName, vars["name"], vars["id"]), nil)
        if err != nil {
            rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
            return
        }
        rest.PostProcessResponse(w, cliCtx, res)
    }
}

