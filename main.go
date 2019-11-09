package main

import (
	"fmt"
)




// 4.引入区块链

// 5.添加区块

// 6.重构代码


func main(){
	bc := NewBlockChain()
	bc.AddBlock("111111")
	bc.AddBlock("222222")

	fmt.Println("------------------------------")
	fmt.Println("------------------------------")
	fmt.Println("------------------------------")

	// 调用迭代器
	it := bc.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("前区块哈希:%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希:%x\n", block.Hash)
		fmt.Printf("data:%s\n", block.Data)
		fmt.Printf("***********************\n\n")
		if len(block.PrevHash) == 0{
			fmt.Println("区块链遍历结束")
			break
		}
	}

}