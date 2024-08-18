package api

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/thesphereonline/chain/backend/nft"
)

func GetNFTs(w http.ResponseWriter, r *http.Request) {
	marketplace := nft.NewNFTMarketplace()
	json.NewEncoder(w).Encode(marketplace.NFTs)
}

func AddNFT(w http.ResponseWriter, r *http.Request) {
	marketplace := nft.NewNFTMarketplace()
	nft := &nft.NFT{
		ID:        r.URL.Query().Get("id"),
		Name:      r.URL.Query().Get("name"),
		Owner:     r.URL.Query().Get("owner"),
		Price:     big.NewInt(0),
		Metadata:  r.URL.Query().Get("metadata"),
	}
	marketplace.AddNFT(nft)
	json.NewEncoder(w).Encode(marketplace.NFTs)
}

func BuyNFT(w http.ResponseWriter, r *http.Request) {
	marketplace := nft.NewNFTMarketplace()
	id := r.URL.Query().Get("id")
	buyer := r.URL.Query().Get("buyer")
	marketplace.BuyNFT(id, buyer)
	json.NewEncoder(w).Encode(marketplace.NFTs)
}

func SellNFT(w http.ResponseWriter, r *http.Request) {
	marketplace := nft.NewNFTMarketplace()
	id := r.URL.Query().Get("id")
	seller := r.URL.Query().Get("seller")
	price := big.NewInt(0)
	marketplace.SellNFT(id, seller, price)
	json.NewEncoder(w).Encode(marketplace.NFTs)
}