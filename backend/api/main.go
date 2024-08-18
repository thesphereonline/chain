package main

import (
	"log"
	"net/http"

	"github.com/thesphereonline/chain/backend/api"
)

func main() {
	http.HandleFunc("/blockchain", api.GetBlockchain)
	http.HandleFunc("/addblock", api.AddBlock)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
