package rest

import (
    "fmt"
    "github.com/cosmos/cosmos-sdk/client/context"
    "github.com/gorilla/mux"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
    r.HandleFunc(fmt.Sprintf("/%s/tables", storeName), createTableHandler(cliCtx)).Methods("POST")
    r.HandleFunc(fmt.Sprintf("/%s/tables/{%s}", storeName, "name"), showTableHandler(cliCtx, storeName)).Methods("GET")
    r.HandleFunc(fmt.Sprintf("/%s/tables", storeName), showTablesHandler(cliCtx, storeName)).Methods("GET")
    r.HandleFunc(fmt.Sprintf("/%s/find/{%s}/{%s}", storeName, "name", "id"), showRowHandler(cliCtx, storeName)).Methods("GET")
}
