package main

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/lucazulian/cryptocomparego"
)

func getRanking(limit int) (map[int]string, error) {
	ctx := context.TODO()

	client := cryptocomparego.NewClient(nil)
	coinList, _, err := client.Coin.List(ctx)
	if err != nil {
		fmt.Printf("Something bad happened: %s\n", err)
		return nil, err
	}

	rankingList := make(map[int]string)
	for _, coin := range coinList {
		order, err := strconv.Atoi(coin.SortOrder)
		if err != nil {
			return nil, err
		}
		if order <= limit {
			rankingList[order] = coin.Name
		}
	}

	keys := make([]int, 0, len(rankingList))

	for k := range rankingList {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	sortedRankingList := make(map[int]string)
	for _, k := range keys {
		sortedRankingList[k] = rankingList[k]
	}

	return sortedRankingList, nil
}
