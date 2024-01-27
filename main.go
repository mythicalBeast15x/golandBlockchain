package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// calculateHash creates a SHA256 hash of the block
func (b *Block) calculateHash() string {
	record := fmt.Sprintf("%d%s%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}

// createBlock is a function to create a new block
func createBlock(index int, data string, prevHash string) *Block {
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}
	block.Hash = block.calculateHash()
	return block
}

func main() {
	// Creating a new block
	start := time.Now()
	createBlock(0, "Genesis Block", "")
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Mining Time:", elapsed)
	//fmt.Println("Genesis Block:", genesisBlock)
}
