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

	http.HandleFunc("/v1/ranking", rankingHandler)
	log.Println("Ranking server is running on http://localhost:8082/v1/ranking...")

	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {
	limitParam, ok := r.URL.Query()["limit"]

	if !ok || len(limitParam) < 1 {
		log.Println("Url Param 'limit' is missing")
		return
	}

	limit, err := strconv.Atoi(limitParam[0])
	if err != nil {
		log.Println(err)
	}

	ranking, err := getRanking(limit)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(ranking)
}
