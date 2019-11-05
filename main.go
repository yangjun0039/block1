package main

import "fmt"

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
		Hash:[]byte{},                      //todo
		Data:[]byte(data),
	}
	return &block
}

// 3.生成哈希

// 4.引入区块链

// 5.添加区块

// 6.重构代码


func main(){
	block := NewBlock("应用广泛",nil)
	fmt.Printf("data:%s",block.Data)
}