package data

import (
	"encoding/json"
	"io"
)

// User account structure for Wallet API
type UserWallet struct {
	ID            string         `json:"id"`
	Username      string         `json:"userName"`
	CryptoBalance []CryptoAmount `json:"cryptoBalance"`
}

type CryptoAmount struct {
	Cryptocurrency string  `json:"cryptoCurrency"`
	Amount         float64 `json:"amount"`
}

//attach a json encoder directly to the interface
type UserWallets []*UserWallet

func (userWallets *UserWallets) ToJson(writer io.Writer) error {
	newError := json.NewEncoder(writer)
	return newError.Encode(userWallets)
}

//abstraction to return UserWallets
func GetUserWallet() UserWallets {
	return WalletEntries
}

//mock data for local testing
var WalletEntries = []*UserWallet{
	{
		ID:       "1fc96ed7-b55a-48d7-855d-4868ba326353",
		Username: "Holland",
		CryptoBalance: []CryptoAmount{
			{
				Cryptocurrency: "btc",
				Amount:         12.34,
			},
			{
				Cryptocurrency: "xrp",
				Amount:         345.49,
			},
			{
				Cryptocurrency: "doge",
				Amount:         0.000008976,
			},
		},
	},
	{
		ID:       "5dd5fabe-d82d-4e40-a233-3e214189bfd3",
		Username: "Garfield",
		CryptoBalance: []CryptoAmount{
			{
				Cryptocurrency: "eth",
				Amount:         812.34,
			},

			{
				Cryptocurrency: "doge",
				Amount:         89768.09,
			},
		},
	},
	{
		ID:       "fe547020-6941-40a5-9c68-0655a0793c54",
		Username: "McGuire",
		CryptoBalance: []CryptoAmount{
			{
				Cryptocurrency: "btc",
				Amount:         0.013,
			},
		},
	},
}
