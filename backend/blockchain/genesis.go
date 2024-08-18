package blockchain

import "time"

func (b *Blockchain) CreateGenesisBlock() {
	genesisBlock := &Block{
		Index:     1,
		Timestamp: time.Now().Unix(),
		Data:      "Genesis Block",
		Hash:      b.calculateHash(&Block{Index: 1, Timestamp: time.Now().Unix(), Data: "Genesis Block"}),
		PrevHash:  "",
	}
	b.Blocks = append(b.Blocks, genesisBlock)
}
