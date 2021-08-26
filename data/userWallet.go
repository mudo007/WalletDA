package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// User account structure for Wallet API
type CryptoBalance struct {
	Currency       string  `json:"currency"`
	Amount         float64 `json:"amount"`
	PriceInDollars float64 `json:"priceInDollars"`
	PriceInEuros   float64 `json:"priceInEuros"`
	RateTimeStamp  string  `json:"timeOfRateUsed"`
	TotalEuros     float64 `json:"totalEuros"`
	TotalDollars   float64 `json:"totalDollars"`
}

type UserBalance struct {
	UserName          string          `json:"userName"`
	TotalEuros        float64         `json:"totalEuros"`
	TotalDollars      float64         `json:"totalDollars"`
	CryptoBalanceList []CryptoBalance `json:"CryptoBalanceList"`
}

//attach a json encoder directly to the interface
func (userBalance *UserBalance) ToJson(writer io.Writer) error {
	newError := json.NewEncoder(writer)
	return newError.Encode(userBalance)
}

//abstraction to return user Balances
func GetUserBalance(userName string, userBalanceList []UserBalance) (UserBalance, error) {
	//search for user in mock array
	userBalance, err := FindUSer(&userBalanceList, userName)
	//carry the error forward
	if err != nil {
		return userBalance, err
	}
	//calculate user balance and return
	return userBalance.CalculateUserBalance(), nil

}

//search the user
func FindUSer(userBalanceList *[]UserBalance, userName string) (UserBalance, error) {
	for _, userBalance := range *userBalanceList {
		if userBalance.UserName == userName {
			return userBalance, nil
		}
	}
	var emptyUserBalance UserBalance
	return emptyUserBalance, userNotFound(userName)
}

type userNotFound string

//custom error function for userBalance
func (e userNotFound) Error() string {
	return fmt.Sprintf("User not found: %s", string(e))
}

//calculate totals
func (userBalance *UserBalance) CalculateUserBalance() UserBalance {
	//accumulators
	var totalEuros, totalDollars float64
	var calculatedUserBalance UserBalance
	var calculatedUserBalanceList []CryptoBalance = userBalance.CryptoBalanceList
	for i := range calculatedUserBalanceList {
		calculatedUserBalanceList[i].TotalEuros = calculatedUserBalanceList[i].Amount * calculatedUserBalanceList[i].PriceInEuros
		totalEuros += calculatedUserBalanceList[i].TotalEuros
		calculatedUserBalanceList[i].TotalDollars = calculatedUserBalanceList[i].Amount * calculatedUserBalanceList[i].PriceInDollars
		totalDollars += calculatedUserBalanceList[i].TotalDollars
	}
	calculatedUserBalance.UserName = userBalance.UserName
	calculatedUserBalance.CryptoBalanceList = calculatedUserBalanceList
	calculatedUserBalance.TotalDollars = totalDollars
	calculatedUserBalance.TotalEuros = totalEuros
	return calculatedUserBalance
}

//------------------Mock Data --------------------------

var BalanceMockData = []UserBalance{
	{
		UserName:     "Holland",
		TotalEuros:   0.0,
		TotalDollars: 0.0,
		CryptoBalanceList: []CryptoBalance{
			{
				Currency:       "btc",
				Amount:         12.5,
				PriceInDollars: 25000,
				PriceInEuros:   18000,
				RateTimeStamp:  "08/09/24 23:34",
				TotalEuros:     0.0,
				TotalDollars:   0.0,
			},
			{
				Currency:       "xrp",
				Amount:         345.987,
				PriceInDollars: 750,
				PriceInEuros:   650,
				RateTimeStamp:  "08/09/24 23:34",
				TotalEuros:     0.0,
				TotalDollars:   0.0,
			},
		},
	},
	{
		UserName:     "Garfield",
		TotalEuros:   0.0,
		TotalDollars: 0.0,
		CryptoBalanceList: []CryptoBalance{
			{
				Currency:       "eth",
				Amount:         12.5,
				PriceInDollars: 650,
				PriceInEuros:   530,
				RateTimeStamp:  "08/09/24 23:34",
				TotalEuros:     0.0,
				TotalDollars:   0.0,
			},
			{
				Currency:       "doge",
				Amount:         345.987,
				PriceInDollars: 750,
				PriceInEuros:   650,
				RateTimeStamp:  "08/09/24 23:34",
				TotalEuros:     0.0,
				TotalDollars:   0.0,
			},
		},
	},
}
