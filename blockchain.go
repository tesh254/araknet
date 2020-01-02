package main

// Blockchain define blockchain structure
type Blockchain struct {
	blocks []*Block
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
