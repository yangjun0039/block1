package main

import (
	"fmt"
	"os"
)

type CLI struct{
	bc *BlockChain
}

const usage = `
	addBlock --data Data         "添加区块"
	printChain                   "打印区块链"
	printChainR                  "反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余额"
`

func (cli *CLI)Run(){
	args := os.Args
	if len(args) < 2{
		fmt.Printf(usage)
		return
	}

	cmd := args[1]
	switch cmd {
	case "addBlock":
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
	case "getBalance":
		//获取余额
		fmt.Println("获取余额")
		if len(args) == 4 && args[2] =="--address"{
			address := args[3]
			cli.GetBalance(address )
		}
	default:
		fmt.Println("无效的命令")
		fmt.Printf(usage)
	}
}

