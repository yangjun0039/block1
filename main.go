package main

// 4.引入区块链

// 5.添加区块

// 6.重构代码

func main() {
	address := "1LfgYNozXixcV2E7MZ8gRStxXQehyQ1cWt"
	bc := NewBlockChain(address)
	cli := CLI{bc}
	cli.Run()

	// 调用迭代器
	//it := bc.NewIterator()
	//for {
	//	block := it.Next()
	//	fmt.Printf("前区块哈希:%x\n", block.PrevHash)
	//	fmt.Printf("当前区块哈希:%x\n", block.Hash)
	//	fmt.Printf("data:%s\n", block.Data)
	//	fmt.Printf("***********************\n\n")
	//	if len(block.PrevHash) == 0{
	//		fmt.Println("区块链遍历结束")
	//		break
	//	}
	//}

}
