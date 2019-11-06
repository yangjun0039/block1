package main

// 引入区块链
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}

//定义一个区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		[]*Block{genesisBlock},
	}
}

//定义一个创世链
func GenesisBlock() *Block {
	return NewBlock("我是创世者", []byte{})
}

//添加区块
func (bc *BlockChain) AddBlock(data string){
	// 最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	// 创建新的区块
	block := NewBlock(data, prevHash)
	// 添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}
