package data

import (
	"fmt"
	"time"
)

type Rates struct {
	Currency       string    `json:"currency"`
	PriceInEuros   float64   `json:"priceInEuros"`
	PriceInDollars float64   `json:"priceInDollars"`
	LastUpdated    time.Time `json:"lastUpdated"`
}

//mocked function...
func GetRates(currency string) (Rates, error) {
	rate, err := FindRate(&ratesMockData, currency)
	//carry the error forward
	if err != nil {
		return rate, err
	}
	//add timestamp
	rate.LastUpdated = time.Now()
	return rate, nil

}

//search the user
func FindRate(ratesList *[]Rates, currency string) (Rates, error) {
	for _, element := range *ratesList {
		if element.Currency == currency {
			return element, nil
		}
	}
	var emptyRate Rates
	return emptyRate, rateNotFound(currency)
}

type rateNotFound string

//custom error function for userBalance
func (e rateNotFound) Error() string {
	return fmt.Sprintf("Rate not found: %s", string(e))
}

var ratesMockData = []Rates{
	{
		Currency:       "btc",
		PriceInDollars: 25000,
		PriceInEuros:   18000,
	},
	{
		Currency:       "eth",
		PriceInDollars: 750,
		PriceInEuros:   650,
	},
	{
		Currency:       "ada",
		PriceInDollars: 12,
		PriceInEuros:   8,
	},
	{
		Currency:       "xrp",
		PriceInDollars: 28,
		PriceInEuros:   23,
	},
	{
		Currency:       "doge",
		PriceInDollars: 0.005,
		PriceInEuros:   0.002,
	},
}
