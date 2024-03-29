package data

import "time"

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
				RateTimeStamp:  time.Now(),
				TotalEuros:     0.0,
				TotalDollars:   0.0,
			},
			{
				Currency:       "xrp",
				Amount:         345.987,
				PriceInDollars: 750,
				PriceInEuros:   650,
				RateTimeStamp:  time.Now(),
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
				RateTimeStamp:  time.Now(),
				TotalEuros:     0.0,
				TotalDollars:   0.0,
			},
			{
				Currency:       "doge",
				Amount:         345.987,
				PriceInDollars: 750,
				PriceInEuros:   650,
				RateTimeStamp:  time.Now(),
				TotalEuros:     0.0,
				TotalDollars:   0.0,
			},
		},
	},
}
