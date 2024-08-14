package ledger

import "time"

type Blockchain struct {
    Blocks []*Block
}

func NewBlockchain() *Blockchain {
    return &Blockchain{
        Blocks: []*Block{genesisBlock()},
    }
}

func (bc *Blockchain) AddBlock(transactions []Transaction, contracts []SmartContract) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := &Block{
        Index:        len(bc.Blocks),
        Timestamp:    time.Now(),
        Transactions: transactions,
        PreviousHash: prevBlock.Hash,
        Difficulty:   bc.getDifficulty(),
        Nonce:        0,
    }
    
    for _, contract := range contracts {
        contract.Execute()
    }
    
    newBlock.MineBlock(newBlock.Difficulty)
    newBlock.Hash = newBlock.CalculateHash()
    bc.Blocks = append(bc.Blocks, newBlock)
}


func genesisBlock() *Block {
    genesis := &Block{
        Index:     0,
        Timestamp: time.Now(),
        Nonce:     0,
    }
    genesis.Hash = genesis.CalculateHash()
    return genesis
}

func (bc *Blockchain) AddReceivedBlock(block *Block) bool {
    if block.PreviousHash == bc.Blocks[len(bc.Blocks)-1].Hash {
        bc.Blocks = append(bc.Blocks, block)
        return true
    }
    return false
}

func (bc *Blockchain) AddBlock(transactions []Transaction) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := &Block{
        Index:        len(bc.Blocks),
        Timestamp:    time.Now(),
        Transactions: transactions,
        PreviousHash: prevBlock.Hash,
        Difficulty:   bc.getDifficulty(),
        Nonce:        0,
    }
    newBlock.MineBlock(newBlock.Difficulty)
    newBlock.Hash = newBlock.CalculateHash()
    bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) getDifficulty() int {
    // Simple difficulty adjustment logic (for example, every 10 blocks)
    if len(bc.Blocks)%10 == 0 {
        return bc.Blocks[len(bc.Blocks)-1].Difficulty + 1
    }
    return bc.Blocks[len(bc.Blocks)-1].Difficulty
}


