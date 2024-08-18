package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Index     int    `json:"index"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevHash"`
}

type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{}}
}

func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := &Block{
		Index:     len(b.Blocks) + 1,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = b.calculateHash(newBlock)
	b.Blocks = append(b.Blocks, newBlock)
}

func (b *Blockchain) calculateHash(block *Block) string {
	blockData, _ := json.Marshal(block)
	hash := sha256.Sum256(blockData)
	return fmt.Sprintf("%x", hash)
}
