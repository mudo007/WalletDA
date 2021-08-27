package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// User account structure for Wallet API
type CryptoBalance struct {
	Currency       string    `json:"currency"`
	Amount         float64   `json:"amount"`
	PriceInDollars float64   `json:"priceInDollars"`
	PriceInEuros   float64   `json:"priceInEuros"`
	RateTimeStamp  time.Time `json:"timeOfRateUsed"`
	TotalEuros     float64   `json:"totalEuros"`
	TotalDollars   float64   `json:"totalDollars"`
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
	result, errCalc := userBalance.CalculateUserBalance()
	if errCalc != nil {
		return userBalance, errCalc
	}
	return result, nil

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
func (userBalance *UserBalance) CalculateUserBalance() (UserBalance, error) {
	//accumulators
	var totalEuros, totalDollars float64
	var calculatedUserBalance UserBalance
	var calculatedUserBalanceList []CryptoBalance = userBalance.CryptoBalanceList
	for i := range calculatedUserBalanceList {
		rate, err := GetRates(calculatedUserBalanceList[i].Currency)
		if err != nil {
			return calculatedUserBalance, err
		}

		calculatedUserBalanceList[i].PriceInEuros = rate.PriceInEuros
		calculatedUserBalanceList[i].PriceInDollars = rate.PriceInDollars
		calculatedUserBalanceList[i].RateTimeStamp = rate.LastUpdated

		calculatedUserBalanceList[i].TotalEuros = calculatedUserBalanceList[i].Amount * calculatedUserBalanceList[i].PriceInEuros
		totalEuros += calculatedUserBalanceList[i].TotalEuros
		calculatedUserBalanceList[i].TotalDollars = calculatedUserBalanceList[i].Amount * calculatedUserBalanceList[i].PriceInDollars
		totalDollars += calculatedUserBalanceList[i].TotalDollars
	}
	calculatedUserBalance.UserName = userBalance.UserName
	calculatedUserBalance.CryptoBalanceList = calculatedUserBalanceList
	calculatedUserBalance.TotalDollars = totalDollars
	calculatedUserBalance.TotalEuros = totalEuros
	return calculatedUserBalance, nil
}
