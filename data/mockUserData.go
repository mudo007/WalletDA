package data

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
