package main

import (
	"fmt"
	"os"
)

type CLI struct{
	bc *BlockChain
}

const usage = `
	addBlock --data Data        "add data to blockchain"
	printChain                  "print all bloakchain data"
`

func (cli *CLI)Run(){
	args := os.Args
	if len(args) < 2{
		fmt.Printf(usage)
		return
	}

	cmd := args[1]
	switch cmd {
	case "addblock":
		//添加区块
		fmt.Println("添加区块")
		// 确定命令有效
		if len(args) == 4 && args[2] =="--data"{
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Println("添加区块参数错误")
			fmt.Printf(usage)
		}
	case "printChain":
		//打印区块
		fmt.Println("打印区块")
		cli.PrintBlockChain()

	default:
		fmt.Println("无效的命令")
		fmt.Printf(usage)
	}
}

