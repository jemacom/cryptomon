package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 0, "TCP port for the HTTP server to listen on")

	flag.Parse()

	http.HandleFunc("/v1/ranking", rankingHandler)
	log.Println("Ranking server is running...")

	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {

	limit, err := parseLimitParam(r)
	if err != nil {
		return
	}
	ranking, err := getRanking(limit)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(ranking)
}

func parseLimitParam(r *http.Request) (int, error) {
	limitParam, ok := r.URL.Query()["limit"]

	if !ok || len(limitParam) < 1 {
		return -1, errors.New("URL Param 'limit' is missing")
	}

	l, err := strconv.Atoi(limitParam[0])
	if err != nil {
		return -1, err
	}

	return l, nil
}
