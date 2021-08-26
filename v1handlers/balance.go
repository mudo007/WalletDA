package v1handlers

import (
	"log"
	"net/http"

	"example.com/walletDA/data"
)

type Balance struct {
	logger *log.Logger
}

//placeholder to log to database
func BalanceWithLogger(logger *log.Logger) *Balance {
	return &Balance{logger}
}

func (wallets *Balance) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	//http verbs handling
	if request.Method == http.MethodGet {
		wallets.getWallet(responseWriter, request)
		return
	}

	//catch all
	responseWriter.WriteHeader((http.StatusMethodNotAllowed))
}

func (wallets *Balance) getWallet(responseWriter http.ResponseWriter, request *http.Request) {
	wallets.logger.Println("Running balance handler")

	//read userWallets
	wallet := data.GetUserWallet()
	err := wallet.ToJson(responseWriter)
	if err != nil {
		log.Println("Error encoding json", err)

		//return server error 500
		http.Error(responseWriter, "Error encoding json", http.StatusInternalServerError)
		return
	}
}
