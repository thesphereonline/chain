package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesphereonline/chain/backend/blockchain"
)

func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	blockchain := blockchain.NewBlockchain()
	blockchain.CreateGenesisBlock()
	json.NewEncoder(w).Encode(blockchain)
}

func AddBlock(w http.ResponseWriter, r *http.Request) {
	blockchain := blockchain.NewBlockchain()
	blockchain.CreateGenesisBlock()
	data := r.URL.Query().Get("data")
	blockchain.AddBlock(data)
	json.NewEncoder(w).Encode(blockchain)
}