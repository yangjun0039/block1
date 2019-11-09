package main

import (
	"block/block1/bolt"
	"log"
)

type BlockChainIterator struct{
	db *bolt.DB
	// 游标，用于不断索引
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator{
	return &BlockChainIterator{
		bc.db,
		bc.tail,
	}
}

// 1.返回当前的区块
// 2.指针前移
func (bi *BlockChainIterator) Next() *Block{
	var block Block
	bi.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			log.Panic("bucket为空")
		}
		blockSerialize := bucket.Get(bi.currentHashPointer)
		block = DeSerialize(blockSerialize)
		bi.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}

