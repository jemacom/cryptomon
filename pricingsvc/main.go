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
	http.HandleFunc("/v1/pricing", pricingHandler)
	log.Println("Pricing server is running...")

	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func pricingHandler(w http.ResponseWriter, r *http.Request) {
	prices, err := getCurrentPrices("USD")
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(prices)
}
