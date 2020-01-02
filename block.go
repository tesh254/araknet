package main

import (
	"bytes"
	"crypto/sha256"
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
