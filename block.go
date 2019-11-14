package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
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
	//Data []byte //数据
	Transactions []*Transaction //真实的交易数组
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
func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		//Data:       []byte(data),
		Transactions: txs,
	}
	block.MerkelRoot = block.MakeMerkelRoot()

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
/*
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
*/

func (block *Block) MakeMerkelRoot() []byte {
	var info []byte
	for _, tx := range block.Transactions {
		// 将交易的哈希值拼接起来，再整体做哈希
		info = append(info, tx.TXID...)
	}
	hash := sha256.Sum256(info)
	return hash[:]
}

// 序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

// 反序列化
func DeSerialize(data []byte) Block {
	var block Block
	deCoder := gob.NewDecoder(bytes.NewReader(data))
	err := deCoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	return block
}
