package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/walletDA/v1handlers"
	"github.com/gorilla/mux"
)

func main() {
	//initiate handlers and loggers
	logger := log.New(os.Stdout, "v1_api ", log.LstdFlags)
	balanceHandler := v1handlers.BalanceWithLogger((logger))

	//Create new Servermux with gorilla
	serverMux := mux.NewRouter()
	//configure the response content-type globally
	serverMux.Use(commonMiddleware)

	//define get router
	getRouter := serverMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/balance/{userName}", balanceHandler.GetBalance)
	//TODO: GET endpoint for history

	//TODO: PUT endpoints for withdraw/deposit

	//create custom server to use options
	//good keep-alive timeout for those trading bots IdleTimeout
	walletServer := &http.Server{
		Addr:         ":9090",
		Handler:      serverMux,
		IdleTimeout:  600 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//run our server non-blocking
	go func() {
		logger.Println("Starting server on port 9090")

		err := walletServer.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	osChannel := make(chan os.Signal, 1)
	signal.Notify(osChannel, os.Interrupt)
	signal.Notify(osChannel, os.Kill)

	// Block until a signal is received.
	signalCaptured := <-osChannel
	log.Println("Got signal:", signalCaptured)

	//TODO: cancel function
	shutDownContext, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		logger.Printf("Time up, shutting down!\n")
	}
	walletServer.Shutdown(shutDownContext)

}

//awlays add response type  application/json
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
