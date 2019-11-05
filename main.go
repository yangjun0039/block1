package main

import (
	"fmt"
	"crypto/sha256"
)

// 1.定义结构
type Block struct {
	PrevHash []byte    //前区块哈希
	Hash     []byte    //当前区块哈希
	Data     []byte    //数据
}

// 2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block{
	block := Block{
		PrevHash:prevBlockHash,
		Hash:[]byte{},
		Data:[]byte(data),
	}
	block.SetHash()
	return &block
}

// 3.生成哈希
func (block *Block)SetHash(){
 	// 1.拼装数据
	blockInfo := append(block.PrevHash, block.Hash...)
	// 2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

// 4.引入区块链

// 5.添加区块

// 6.重构代码


func main(){
	block := NewBlock("应用广泛",nil)
	fmt.Printf("当前区块哈希:%x\n",block.Hash)
	fmt.Printf("data:%s\n",block.Data)

}