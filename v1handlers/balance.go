package v1handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Balance struct {
	logger *log.Logger
}

//placeholder to log to database
func BalanceLogSQL(logger *log.Logger) *Balance {
	return &Balance{logger}
}

func (handler *Balance) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	handler.logger.Println("Running balance handler")

	// read the body
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println("Error reading body", err)

		http.Error(responseWriter, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// write the mock response
	fmt.Fprintf(responseWriter, "Hello %s", body)
}
