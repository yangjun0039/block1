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

	// 1.创建挖矿交易
	coinbase := NewCoinbaseTX(miner, data)
	// 2.创建一个普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		fmt.Println("无效的交易")
		return
	}
	// 3.添加到区块
	cli.bc.AddBlock([]*Transaction{coinbase, tx})
	fmt.Println("转账成功")
}

func (cli *CLI) NewWallet() {
	//wallet := NewWallet()
	//address := wallet.NewAddress()
	ws := NewWallets()
	address := ws.CreateWallet()
	fmt.Printf("地址: %s\n", address)
	//for address,_ := range ws.WalletsMap{
	//	fmt.Printf("地址: %s\n", address)
	//}

}

func (cli *CLI) ListAddresses(){
	ws := NewWallets()
	addresses := ws.GetAllAddress()

	for _,address := range addresses{
		fmt.Printf("地址：%s\n",address)
	}
}
