package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	//cli.bc.AddBlock(data) todo
}

func (cli *CLI) PrintBlockChain() {
	it := cli.bc.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("版本号:%x\n", block.Version)
		fmt.Printf("前区块哈希:%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希:%x\n", block.Hash)
		fmt.Printf("默克尔根:%x\n", block.MerkelRoot)
		fmt.Printf("时间戳:%d\n", block.TimeStamp)
		fmt.Printf("难度值:%d\n", block.Difficulty)
		fmt.Printf("随机数:%x\n", block.Nonce)
		fmt.Printf("data:%s\n", block.Transactions[0].TXInputs[0].Sig)
		fmt.Printf("***********************\n\n")
		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束")
			break
		}
	}
}
