package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

//创建一个新的区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock() //创世区块
	blockChain := BlockChain{}
	blockChain.AppendBlock(&genesisBlock)
	return &blockChain
}

//添加区块
func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 { //如果为0 代表是创世区块
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

/**
 */
func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1] //父区块
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

//校验添加的区块是否合法
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	//判断当前区块的区块头所含的父级区块hash是否和当前区块hash相同
	if oldBlock.Hash != newBlock.PrevBlockHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index %d\n", block.Index)
		fmt.Printf("PrevBlockHash %s\n", block.PrevBlockHash)
		fmt.Printf("Hash %s\n", block.Hash)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("TimeStamp %d\n", block.Timestamp)
		fmt.Println()
	}
}
