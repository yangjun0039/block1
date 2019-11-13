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

func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)
	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("%s余额为%f", address, total)
}

func (cli *CLI) Send(from, to string, amount float64, miner, data string) {
	fmt.Printf("from: %s\n", from)
	fmt.Printf("to: %s\n", to)
	fmt.Printf("amount: %f\n", amount)
	fmt.Printf("miner: %s\n", miner)
	fmt.Printf("data: %s\n", data)
	// todo
}
