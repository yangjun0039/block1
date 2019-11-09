package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

// 定义一个工作量证明的结构ProofOfWork
type ProofOfWork struct {
	// block
	block *Block
	// 目标值
	// 一个非常大的数，他有很多丰富的方法：比较，赋值方法
	targat *big.Int
}

// 提供创建POW的函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	// 指定难度值，现在是一个string类型，需要进行转换
	taggetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	// 引入的辅助变量，目的是将上面的难度值转成big.int
	tmpInt := big.Int{}
	// 将难度值赋值给big.int,指定16进制的格式
	tmpInt.SetString(taggetStr, 16)
	pow.targat = &tmpInt
	return &pow
}

// 提供不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	var nonce uint64
	block := pow.block
	var hash [32]byte
	for {
		// 拼装数据(区块数据还有不断变换的随机数)
		temp := [][]byte{
			Unit64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Unit64ToByte(block.TimeStamp),
			Unit64ToByte(block.Difficulty),
			Unit64ToByte(nonce),
			block.Hash,
			block.Data,
		}
		// 将二维的切片数组连接起来，返回一个一维的切片
		blockInfo := bytes.Join(temp, []byte{})

		// 做哈希运算
		hash = sha256.Sum256(blockInfo)

		// 与pow中target进行比较
		tmpInt := big.Int{}
		// 将得到的hash数组转成一个big.int
		tmpInt.SetBytes(hash[:])
		// 比较当前哈希值与目标哈希值，如果当前哈希值小于目标哈希值就证明找到了，否则继续找
		if tmpInt.Cmp(pow.targat) == -1 {
			//找到了
			fmt.Printf("挖矿成功，hash: %x, nonce: %d\n", hash, nonce)
			break
		} else {
			nonce ++
		}

	}

	return hash[:], nonce
}

// 提供一个校验函数
