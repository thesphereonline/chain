package ledger

import (
    "log"
	"crypto/sha256"
    "encoding/hex"
)

type SmartContract struct {
    ID      string
    Code    string
    Creator string
}

func NewSmartContract(code, creator string) *SmartContract {
    contract := &SmartContract{
        Code:    code,
        Creator: creator,
    }
    contract.ID = contract.calculateHash()
    return contract
}

func (sc *SmartContract) calculateHash() string {
    record := sc.Code + sc.Creator
    hash := sha256.New()
    hash.Write([]byte(record))
    hashed := hash.Sum(nil)
    return hex.EncodeToString(hashed)
}

func (sc *SmartContract) Execute() {
    // Simple execution logic (for demonstration purposes)
    log.Printf("Executing smart contract: %s", sc.Code)
    // More advanced logic could involve a virtual machine or interpreter
}
