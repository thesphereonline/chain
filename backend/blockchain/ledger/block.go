package ledger

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
	"strings"
)

type Block struct {
    Index        int
    Timestamp    time.Time
    Transactions []Transaction
    PreviousHash string
    Hash         string
    Nonce        int
    Difficulty   int
}


func (b *Block) CalculateHash() string {
    record := string(b.Index) + b.Timestamp.String() + b.PreviousHash + string(b.Nonce)
    for _, tx := range b.Transactions {
        record += tx.ID + tx.Sender + tx.Receiver + string(tx.Amount)
    }
    hash := sha256.New()
    hash.Write([]byte(record))
    hashed := hash.Sum(nil)
    return hex.EncodeToString(hashed)
}

func (b *Block) MineBlock(difficulty int) {
    target := strings.Repeat("0", difficulty)
    for {
        b.Nonce++
        hash := b.CalculateHash()
        if hash[:difficulty] == target {
            b.Hash = hash
            break
        }
    }
}

