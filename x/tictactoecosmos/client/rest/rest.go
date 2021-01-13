package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers tictactoecosmos-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/tictactoecosmos/start", listStartHandler(cliCtx, "tictactoecosmos")).Methods("GET")
	r.HandleFunc("/tictactoecosmos/start", createStartHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/tictactoecosmos/move", listMoveHandler(cliCtx, "tictactoecosmos")).Methods("GET")
	r.HandleFunc("/tictactoecosmos/move", createMoveHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/tictactoecosmos/game", listGameHandler(cliCtx, "tictactoecosmos")).Methods("GET")
	r.HandleFunc("/tictactoecosmos/game", createGameHandler(cliCtx)).Methods("POST")
}
