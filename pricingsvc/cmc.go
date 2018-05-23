package main

import (
	"fmt"

	cmc "github.com/coincircle/go-coinmarketcap"
)

func getCurrentPrices(limit int, currency string) (map[string]float64, error) {
	currentPrices := make(map[string]float64)

	tickers, err := cmc.Tickers(&cmc.TickersOptions{
		Convert: currency,
	})
	if err != nil {
		return nil, err
	}

	for _, ticker := range tickers {
		currentPrices[ticker.Symbol] = ticker.Quotes["USD"].Price
	}

	fmt.Printf("%v", currentPrices)

	return currentPrices, nil
}
