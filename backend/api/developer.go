package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesphereonline/chain/backend/blockchain"
	"github.com/thesphereonline/chain/backend/nft"
	"github.com/thesphereonline/chain/backend/token"
)

func GetBlockchainData(w http.ResponseWriter, r *http.Request) {
	blockchain := blockchain.NewBlockchain()
	json.NewEncoder(w).Encode(blockchain)
}

func GetNFTData(w http.ResponseWriter, r *http.Request) {
	marketplace := nft.NewNFTMarketplace()
	json.NewEncoder(w).Encode(marketplace.NFTs)
}

func GetTokenData(w http.ResponseWriter, r *http.Request) {
	token := token.NewToken()
	json.NewEncoder(w).Encode(token)
}