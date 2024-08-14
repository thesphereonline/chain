package p2p

import (
    "log"
    "net"
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

func (node *Node) Start() {
    listener, err := net.Listen("tcp", node.Address)
    if err != nil {
        log.Fatalf("Failed to start node: %v", err)
    }
    defer listener.Close()

    log.Printf("Node started at %s", node.Address)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
        node.Conn = conn
        go node.handleConnection(conn)
    }
}

func (node *Node) handleConnection(conn net.Conn) {
    buffer := make([]byte, 4096)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            log.Printf("Connection error: %v", err)
            return
        }

        message := string(buffer[:n])
        log.Printf("Received message: %s", message)

        if strings.HasPrefix(message, "NEW_BLOCK|") {
            blockHash := strings.TrimPrefix(message, "NEW_BLOCK|")
            log.Printf("New block received with hash: %s", blockHash)
            // Verify and add the block to the chain
        }
    }
}



func generateNodeID(address string) string {
    // Generate a unique ID for the node based on its address
    return fmt.Sprintf("%x", sha256.Sum256([]byte(address)))
}

func (node *Node) SendMessage(address string, message string) {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        log.Printf("Failed to connect to %s: %v", address, err)
        return
    }
    defer conn.Close()

    _, err = conn.Write([]byte(message))
    if err != nil {
        log.Printf("Failed to send message to %s: %v", address, err)
    }
}

func (node *Node) BroadcastBlock(block *ledger.Block) {
    message := fmt.Sprintf("NEW_BLOCK|%s|%d|%d", block.Hash, block.Index, block.Difficulty)
    for _, neighbor := range node.Neighbors {
        node.SendMessage(neighbor, message)
    }
}


