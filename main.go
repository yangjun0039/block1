package main

import (
)




// 4.引入区块链

// 5.添加区块

// 6.重构代码


func main(){
	bc := NewBlockChain()
	bc.AddBlock("我是第二块")
	bc.AddBlock("我是第三块")

	//for i,block := range bc.blocks {
	//	fmt.Printf("=============当前区块高度===========:%x\n", i)
	//	//block := NewBlock("应用广泛", nil)
	//	fmt.Printf("前区块哈希:%x\n", block.PrevHash)
	//	fmt.Printf("当前区块哈希:%x\n", block.Hash)
	//	fmt.Printf("data:%s\n", block.Data)
	//}
}