package main

import (
	"block/block1/bolt"
	"log"
)

// 引入区块链
type BlockChain struct {
	//定义一个区块链数组
	//blocks []*Block
	db   *bolt.DB
	tail []byte //存储最后一个区块的哈希
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

//定义一个区块链
func NewBlockChain() *BlockChain {
	//genesisBlock := GenesisBlock()
	//return &BlockChain{
	//	[]*Block{genesisBlock},
	//}
	var lastHash []byte

	// 1.打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}

	//操作数据库(改写)
	db.Update(func(tx *bolt.Tx) error {
		// 2.找到抽屉bucket(没有就创建)
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			// 没有抽屉，需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建抽屉失败")
			}
			// 创建一个创世块，并作为第一个区块加入到区块链中
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

//定义一个创世链
func GenesisBlock() *Block {
	return NewBlock("我是创世者", []byte{})
}

//添加区块
func (bc *BlockChain) AddBlock(data string) {
//	// 最后一个区块
//	lastBlock := bc.blocks[len(bc.blocks)-1]
//	prevHash := lastBlock.Hash
//	// 创建新的区块
//	block := NewBlock(data, prevHash)
//	// 添加到区块链数组中
//	bc.blocks = append(bc.blocks, block)
}
