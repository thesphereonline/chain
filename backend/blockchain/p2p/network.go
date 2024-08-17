package p2p

import (
	"fmt"
	"crypto/sha256"
)

func generateNodeID(address string) string {
    // Generate a unique ID for the node based on its address
    return fmt.Sprintf("%x", sha256.Sum256([]byte(address)))
}
