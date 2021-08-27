package data

import (
	"encoding/json"
	"io"
)

type TransactionBody struct {
	UserName string  `json:"name"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

func (t *TransactionBody) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(t)
}
