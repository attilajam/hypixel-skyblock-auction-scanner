package main

import (
	"encoding/json"
	"net/http"
)

type ItemsRequest struct {
	Items []Item `json:"items"`
}
type Item struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
type PriceMap map[string]int

func findId(items []Item, search string) string {
	for _, item := range items {
		if item.Name == search {
			return item.Id
		}
	}
	return ""
}
func getPrice(name string) int {
	resp, err := http.Get("https://api.hypixel.net/v2/resources/skyblock/items")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var data ItemsRequest
	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	resp, err = http.Get("https://sky.coflnet.com/api/prices/neu")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder = json.NewDecoder(resp.Body)
	var prices PriceMap
	err = decoder.Decode(&prices)
	return prices[findId(data.Items, name)]
}
