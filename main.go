package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"text/template"
)

type Request struct {
	Auctions []Auction `json:"auctions"`
}

type Auction struct {
	UUID         string `json:"uuid"`
	Item         string `json:"item_name"`
	Price        int    `json:"starting_bid"`
	AveragePrice int
	Profit       int
	Rarity       string `json:"tier"`
	Bin          bool   `json:"bin"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		itemParam := r.URL.Query().Get("query")
		if r.Method == "POST" {
			// Get the input text from the form
			itemParam = r.FormValue("item")
			return
		}
		t, err := template.ParseFiles("table.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp, err := http.Get("https://api.hypixel.net/v2/skyblock/auctions")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var data Request
		err = decoder.Decode(&data)
		if err != nil {
			panic(err)
		}

		// Create a data structure to pass to template

		var filteredAuctions []Auction
		var mu sync.Mutex // Add mutex to protect concurrent access
		var wg sync.WaitGroup

		wg.Add(len(data.Auctions))
		for _, auction := range data.Auctions {
			go func(auction Auction) {
				defer wg.Done()
				if auction.Bin && (auction.Item == itemParam || itemParam == "") {
					mu.Lock()
					auction.AveragePrice = getPrice(auction.Item)
					filteredAuctions = append(filteredAuctions, auction)
					mu.Unlock()
				}
			}(auction)
		}
		wg.Wait()

		err = t.Execute(w, filteredAuctions)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(err.Error()))
		}
	})

	l, err := net.Listen("tcp", ":8080")
	if err == nil {
		fmt.Println("Listening on port 8080")
	}
	http.Serve(l, nil)

}
