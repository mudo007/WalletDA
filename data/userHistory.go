package data

import (
	"encoding/json"
	"io"
	"time"
)

// User account structure for Wallet API
type TransactionHistory struct {
	Currency  string    `json:"currency"`
	Amount    float64   `json:"amount"`
	TimeStamp time.Time `json:"timeStamp"`
}

type HistoryBodyQuery struct {
	UserName  string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func (t *HistoryBodyQuery) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}

//attach a json encoder directly to the interface
func (transactionHistory *TransactionHistory) ToJson(writer io.Writer) error {
	newError := json.NewEncoder(writer)
	return newError.Encode(transactionHistory)
}

// Products is a collection of Product
type Transactions []*TransactionHistory

func (p *Transactions) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
