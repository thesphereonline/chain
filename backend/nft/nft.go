package nft

import (
	"math/big"
)

type NFT struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Price    *big.Int `json:"price"`
	Metadata string   `json:"metadata"`
}

type NFTMarketplace struct {
	NFTs []*NFT `json:"nfts"`
}

func NewNFTMarketplace() *NFTMarketplace {
	return &NFTMarketplace{NFTs: []*NFT{}}
}

func (m *NFTMarketplace) AddNFT(nft *NFT) {
	m.NFTs = append(m.NFTs, nft)
}

func (m *NFTMarketplace) BuyNFT(id string, buyer string) {
	for _, nft := range m.NFTs {
		if nft.ID == id {
			nft.Owner = buyer
			return
		}
	}
}

func (m *NFTMarketplace) SellNFT(id string, seller string, price *big.Int) {
	for _, nft := range m.NFTs {
		if nft.ID == id {
			nft.Owner = seller
			nft.Price = price
			return
		}
	}
}

func (m *NFTMarketplace) GetNFT(id string) *NFT {
	for _, nft := range m.NFTs {
		if nft.ID == id {
			return nft
		}
	}
	return nil
}
