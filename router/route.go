package route

import (
	"banking/handler"
	"github.com/gorilla/mux"
)

func NewRouter(accountHandler handler.AccountHandler, transferHandler handler.TransferHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/accounts", accountHandler.ListAccount).Methods("GET")
	router.HandleFunc("/account", accountHandler.InsertAccount).Methods("POST")
	router.HandleFunc("/transfer", transferHandler.ListTransferByAccount).Methods("GET")
	router.HandleFunc("/transfer", transferHandler.Transfer).Methods("POST")

	return router
}
