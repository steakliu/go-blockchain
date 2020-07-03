package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//区块结构
type Block struct {
	Index         int64
	Timestamp     int64
	PrevBlockHash string
	Hash          string
	Data          string
}

//将区块计算为hash
func calculateHash(block Block) string {
	blockData := string(block.Index) + string(block.Timestamp) + block.PrevBlockHash + block.Data
	hashBlockData := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashBlockData[:])
}

//生成一个新区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

//创世区块,区块链的第一个区块,创世区块没有父区块,也就没有父区块的hash
func GenerateGenesisBlock() Block {
	preBlock := Block{} //虚拟的父区块(实际不存在)
	preBlock.Index = -1
	preBlock.Hash = "" //父区块为空
	return GenerateNewBlock(preBlock, "Genesis Block")
}
