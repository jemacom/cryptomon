package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type coin struct {
	Rank     int
	Symbol   string
	PriceUSD float64
}

func main() {
	port := flag.Int("port", 0, "TCP port for the HTTP server to listen on")
	pricingServerURL := flag.String("pricingServerURL", "", "URL to the 'princing' instance")
	rankingServerURL := flag.String("rankingServerURL", "", "URL to the 'ranking' instance")

	flag.Parse()

	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		listCoinsHandler(w, r, *pricingServerURL, *rankingServerURL)
	})
	log.Println("API server is running...")

	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func listCoinsHandler(w http.ResponseWriter, r *http.Request, pricingServerURL string, rankingServerURL string) {
	limitParam, ok := r.URL.Query()["limit"]

	if !ok || len(limitParam) < 1 {
		log.Println("URL Param 'limit' is missing")
		return
	}

	var coins []coin
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")

	// The limit parameter shouldn't be negative or 0
	// if so we return empty price list
	l, err := strconv.Atoi(limitParam[0])
	if err != nil {
		return
	}
	if l <= 0 {
		if err = enc.Encode(coin{}); err != nil {
			panic(err)
		}
		return
	}

	// get ranking
	ranking, err := getRanking(rankingServerURL, limitParam[0])
	if err != nil {
		return
	}
	// get pricing
	pricing, err := getPricing(pricingServerURL, limitParam[0])
	if err != nil {
		return
	}

	keys := make([]int, 1, len(ranking))
	for k := range ranking {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for k := range keys {
		if ranking[k] != "" {
			c := coin{Rank: k, Symbol: ranking[k], PriceUSD: pricing[ranking[k]]}
			coins = append(coins, c)
		}
	}
	if err := enc.Encode(coins); err != nil {
		panic(err)
	}
}

func getRanking(rankingServerURL string, limit string) (map[int]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s?limit=%s", rankingServerURL, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := make(map[int]string)
	json.NewDecoder(resp.Body).Decode(&r)

	return r, nil
}

func getPricing(pricingServerURL string, limit string) (map[string]float64, error) {
	resp, err := http.Get(fmt.Sprintf("%s?limit=%s", pricingServerURL, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	p := make(map[string]float64)
	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}
