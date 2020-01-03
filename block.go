package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

// Block define block structure
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
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
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Serialize handles serialize a struct
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock handles
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
