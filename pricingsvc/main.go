package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 0, "TCP port for the HTTP server to listen on")

	flag.Parse()
	http.HandleFunc("/v1/pricing", pricesHandler)
	log.Println("Pricing server is running on http://localhost:8081/v1/pricing...")

	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func pricesHandler(w http.ResponseWriter, r *http.Request) {
	limitParam, ok := r.URL.Query()["limit"]

	if !ok || len(limitParam) < 1 {
		log.Println("Url Param 'limit' is missing")
		return
	}

	limit, err := strconv.Atoi(limitParam[0])
	if err != nil {
		log.Println(err)
	}
	prices, err := getCurrentPrices(limit, "USD")
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(prices)
}
