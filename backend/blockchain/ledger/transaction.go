package ledger

import (
    "crypto/sha256"
    "encoding/hex"
)

type Transaction struct {
    ID        string
    Sender    string
    Receiver  string
    Amount    int
    TokenType string  // New field for token type (e.g., "SPHERE", "ETH")
}

func NewTokenTransaction(sender, receiver, tokenType string, amount int) *Transaction {
    tx := &Transaction{
        Sender:    sender,
        Receiver:  receiver,
        Amount:    amount,
        TokenType: tokenType,
    }
    tx.ID = tx.calculateHash()
    return tx
}


func NewTransaction(sender, receiver string, amount int) *Transaction {
    tx := &Transaction{
        Sender:   sender,
        Receiver: receiver,
        Amount:   amount,
    }
    tx.ID = tx.calculateHash()
    return tx
}

func (tx *Transaction) calculateHash() string {
    record := tx.Sender + tx.Receiver + string(tx.Amount)
    hash := sha256.New()
    hash.Write([]byte(record))
    hashed := hash.Sum(nil)
    return hex.EncodeToString(hashed)
}
