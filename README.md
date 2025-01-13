# Hypixel Skyblock Auction Scanner
This program scans the Hypixel Skyblock Auction API to find the newest auctions that were just listed, giving the user a chance to search and find items that may have been listed for much less than the typical auction price. If used correctly, you can use this program to auction flip and make lots of "coins" quickly.

Installation:
```
git clone https://github.com/attilajam/hypixel-skyblock-auction-scanner.git
cd hypixel-skyblock-auction-scanner
go run main.go itemmatching.go
```
This will start the web UI on `localhost:8080`. It may take some time to search through all of the data if you go directly to this link, so if you want to search for a specific item, go to `localhost:8080/query=God+Potion` for example.
