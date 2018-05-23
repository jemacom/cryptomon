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

func main() {
	port := flag.Int("port", 0, "TCP port for the HTTP server to listen on")
	pricingServerURL := flag.String("pricingServerURL", "", "URL to the 'princing' instance")
	rankingServerURL := flag.String("rankingServerURL", "", "URL to the 'ranking' instance")

	flag.Parse()

	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		topCoinsHandler(w, r, *pricingServerURL, *rankingServerURL)
	})
	log.Println("API server is running on http://localhost:8080/v1...")

	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func topCoinsHandler(w http.ResponseWriter, r *http.Request, pricingServerURL string, rankingServerURL string) {
	limitParam, ok := r.URL.Query()["limit"]

	if !ok || len(limitParam) < 1 {
		log.Println("Url Param 'limit' is missing")
		return
	}

	// get ranking
	resp, err := http.Get(fmt.Sprintf("%s?limit=%s", rankingServerURL, limitParam[0]))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	ranking := make(map[int]string)
	json.NewDecoder(resp.Body).Decode(&ranking)

	// get pricing
	resp, err = http.Get(fmt.Sprintf("%s?limit=%s", pricingServerURL, limitParam[0]))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	pricing := make(map[string]float64)
	json.NewDecoder(resp.Body).Decode(&pricing)

	keys := make([]int, 1, len(ranking))

	for k := range ranking {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var coins []coin

	fmt.Println()
	fmt.Println("Rank,   Symbol,   Price USD")
	for k := range keys {
		if ranking[k] != "" {
			c := coin{Rank: k, Symbol: ranking[k], PriceUSD: pricing[ranking[k]]}
			coins = append(coins, c)
		}
	}
	json.NewEncoder(w).Encode(coins)
}
