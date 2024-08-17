package main

import (
    "fmt"
    "github.com/thesphereonline/chain/backend/blockchain"
    "github.com/thesphereonline/chain/backend/blockchain/ledger"
)

func main() {
    bc := blockchain.NewBlockchain()
    transactions := []shared.Transaction{
        {ID: "tx1", Amount: 100},
        {ID: "tx2", Amount: 200},
    }
    contracts := []shared.SmartContract{
        {ID: "sc1", Code: "contract_code_1"},
        {ID: "sc2", Code: "contract_code_2"},
    }
    bc.AddBlock(transactions, contracts)
    fmt.Println("Block added successfully.")
}
