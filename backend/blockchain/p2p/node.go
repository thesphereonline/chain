package p2p

import (
    "fmt"
    "net"
    "bufio"
    "strings"
    "github.com/thesphereonline/chain/backend/blockchain/ledger"
)

type Node struct {
    ID        string
    Address   string
    Neighbors []string
    Conn      net.Conn
}

func NewNode(address string) *Node {
    return &Node{
        ID:        generateNodeID(address),
        Address:   address,
        Neighbors: []string{},
    }
}

// Start starts the node and listens for incoming connections.
func (node *Node) Start() {
    listener, err := net.Listen("tcp", node.Address)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()

    fmt.Printf("Node started at %s\n", node.Address)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        go node.handleConnection(conn)
    }
}

// handleConnection handles incoming connections to the node.
func (node *Node) handleConnection(conn net.Conn) {
    defer conn.Close()
    message, _ := bufio.NewReader(conn).ReadString('\n')
    node.processMessage(strings.TrimSpace(message))
}

// processMessage processes incoming messages.
func (node *Node) processMessage(message string) {
    // Process the message (e.g., block propagation, transaction, etc.)
    fmt.Printf("Received message: %s\n", message)
}

// BroadcastBlock broadcasts a block to all neighbors.
func (node *Node) BroadcastBlock(block *ledger.Block) {
    message := fmt.Sprintf("NEW_BLOCK|%s|%d|%d", block.Hash, block.Index, block.Difficulty)
    for _, neighbor := range node.Neighbors {
        node.SendMessage(neighbor, message)
    }
}

// SendMessage sends a message to a neighbor node.
func (node *Node) SendMessage(address string, message string) {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    fmt.Fprintf(conn, "%s\n", message)
}
