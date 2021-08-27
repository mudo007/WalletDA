package v1handlers

import (
	"log"
	"net/http"

	"example.com/walletDA/data"
)

// resultado := data.ExecuteTransaction(sqlParams, "btc", "Bob", 12.5)

// log.Println(resultado.Error())
type Transaction struct {
	logger *log.Logger
}

//placeholder to log to database
func TransactionWithLogger(logger *log.Logger) *Transaction {
	return &Transaction{logger}
}

func (transaction *Transaction) WithdrawFunds(responseWriter http.ResponseWriter, request *http.Request) {
	transaction.logger.Println("Running withdraw handler")

	//read json from body
	t := data.TransactionBody{}

	err := t.FromJSON(request.Body)
	if err != nil {
		transaction.logger.Println("[ERROR] deserializing product", err)
		http.Error(responseWriter, "Error processing transactiont", http.StatusBadRequest)
		return
	}

	//execute the withdraw
	//yes, it is ugly to copy the desposit function just to invert the amount signal...
	errTransaction := data.ExecuteTransaction(t.Currency, t.UserName, t.Amount*(-1.0))
	switch errTransaction.Error() {
	case "Insufficient funds":
		http.Error(responseWriter, "Insufficient funds", http.StatusForbidden)
		break
	case "Success":
		break

	}

}

func (transaction *Transaction) DepositFunds(responseWriter http.ResponseWriter, request *http.Request) {
	transaction.logger.Println("Running deposit endpoint")

	//read json from body
	t := data.TransactionBody{}

	err := t.FromJSON(request.Body)
	if err != nil {
		transaction.logger.Println("[ERROR] deserializing product", err)
		http.Error(responseWriter, "Error processing transaction", http.StatusBadRequest)
		return
	}

	//execute the deposit
	errTransaction := data.ExecuteTransaction(t.Currency, t.UserName, t.Amount)
	switch errTransaction.Error() {
	case "Insufficient funds":
		http.Error(responseWriter, "Insufficient funds", http.StatusForbidden)
		break
	case "Success":
		break

	}

}
