package main

import (
    "fmt"
    "github.com/thesphereonline/chain/backend/blockchain" // Ensure this path is correct
    "github.com/thesphereonline/chain/backend/blockchain/ledger"
	"github.com/thesphereonline/chain/backend/blockchain/p2p"
)

func main() {
    node1 := p2p.NewNode("localhost:8001")
    node2 := p2p.NewNode("localhost:8002")
    node1.Neighbors = append(node1.Neighbors, node2.Address)

    go node1.Start()
    go node2.Start()

    // Initialize the blockchain
    bc := blockchain.NewBlockchain()

    // Create some sample transactions
    transactions := []ledger.Transaction{
        {ID: "tx1", Amount: 100},
        {ID: "tx2", Amount: 200},
    }

    // Create some sample smart contracts
    contracts := []ledger.SmartContract{
        {ID: "sc1", Code: "contract_code_1"},
        {ID: "sc2", Code: "contract_code_2"},
    }

    // Add a block to the blockchain
    bc.AddBlock(transactions, contracts)

    fmt.Println("Block added successfully.")

    node1.BroadcastBlock(blockchain.Blocks[1])

    select {} // Keep the program running
}
