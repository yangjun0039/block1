package main

import (
	"crypto/sha256"
	"bytes"
	"encoding/binary"
	"time"
)

// 1.定义结构
type Block struct {
	Version    uint64 //版本号
	PrevHash   []byte //前区块哈希
	MerkelRoot []byte //Merkel根
	TimeStamp  uint64 //
	Difficulty uint64 //难度值
	Nonce      uint64 //随机数

	Hash []byte //当前区块哈希，正常比特币区块中没有当前区块的哈希，此处为了方便
	Data []byte //数据
}

//实现一个辅助函数，功能是将uint64转成[]byte
func Unit64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

// 2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}


	//block.SetHash()
	// 创建一个pow对象
	pow := NewProofOfWork(&block)

	// 查找随机数，不停的进行哈希运算
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	// 根据挖矿结果对区块数据进行更新


	return &block
}

// 3.生成哈希
func (block *Block) SetHash() {
	// 1.拼装数据
	//blockInfo := append(block.PrevHash, block.Hash...)
	temp := [][]byte{
		Unit64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Unit64ToByte(block.TimeStamp),
		Unit64ToByte(block.Difficulty),
		Unit64ToByte(block.Nonce),
		block.Hash,
		block.Data,
	}
	// 将二维的切片数组连接起来，返回一个一维的切片
	blockInfo := bytes.Join(temp,[]byte{})
	// 2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

func (block *Block) toByte()  []byte{
	return []byte{}
}