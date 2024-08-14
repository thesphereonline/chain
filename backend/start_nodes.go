package main

import (
    "github.com/thesphereonline/chain/backend/blockchain/ledger"
    "github.com/thesphereonline/chain/backend/blockchain/p2p"
)

func main() {
    node1 := p2p.NewNode("localhost:8001")
    node2 := p2p.NewNode("localhost:8002")
    node1.Neighbors = append(node1.Neighbors, node2.Address)

    go node1.Start()
    go node2.Start()

    // Simulate adding a block
    blockchain := ledger.NewBlockchain()
    tx := ledger.NewTransaction("Alice", "Bob", 50)
    blockchain.AddBlock([]ledger.Transaction{*tx})

    node1.BroadcastBlock(blockchain.Blocks[1])

    select {} // Keep the program running
}
