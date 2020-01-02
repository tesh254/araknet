package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block define block structure
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Blockchain define blockchain structure
type Blockchain struct {
	blocks []*Block
}

// SetHash handles setting the hash of a single block
func (b *Block) SetHash() {
	// Define block timestamp
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// Define combined block data
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// Generate hash from block data
	hash := sha256.Sum256(headers)

	// return hash
	b.Hash = hash[:]
}

// NewBlock handles creation of a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	// Define block based on the Block struct
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	// Generate block hash
	block.SetHash()
	// Return created block
	return block
}

// AddBlock handles adding a block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	// Fetch the last added block
	prevBlock := bc.blocks[len(bc.blocks)-1]
	// Call method to create new block
	newBlock := NewBlock(data, prevBlock.Hash)
	// Update blocks in the blockchain
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock handles creating a genesis block
func NewGenesisBlock() *Block {
	// Return generated block
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain handles create new blockchain
func NewBlockchain() *Blockchain {
	// return blockchain with genesis block
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 btc to Wachira")
	bc.AddBlock("Send 2 more btc to Wachira")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
