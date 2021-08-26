package v1handlers

import (
	"log"
	"net/http"

	"example.com/walletDA/data"
	"github.com/gorilla/mux"
)

type Balance struct {
	logger *log.Logger
}

//placeholder to log to database
func BalanceWithLogger(logger *log.Logger) *Balance {
	return &Balance{logger}
}

func (wallets *Balance) GetBalance(responseWriter http.ResponseWriter, request *http.Request) {
	wallets.logger.Println("Running balance handler")
	//extract userName from querystring
	vars := mux.Vars(request)
	userName := vars["userName"]

	//read userWallets
	userBalance, errUser := data.GetUserBalance(userName, data.BalanceMockData)
	if errUser != nil {
		//return server error 404
		http.Error(responseWriter, "User Not Found", http.StatusNotFound)
		return
	}
	err := userBalance.ToJson(responseWriter)
	if err != nil {
		wallets.logger.Println("Error encoding json", err)

		//return server error 500
		http.Error(responseWriter, "Error encoding json", http.StatusInternalServerError)
		return
	}
}
