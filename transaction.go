package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"log"
)

// 1.定义交易结构
type Transaction struct {
	TXID      []byte     //交易ID
	TXInputs  []TXInput  //交易输入数组
	TXOutputs []TXOutput //交易输出数组

}

// 定义交易输入
type TXInput struct {
	TXid  []byte //引用的交易output索引值id
	Index int64  //引用的交易output索引值
	//Sig   string //解锁脚本，用地址来模拟
	Signature []byte //真正的数字签名，由r,s拼成的[]byte
	PubKey    []byte // 约定，这里的PubKey不存储原始的公钥，而是存储X和Y拼接的字符串，在校验端重新拆分，注意：是公钥，不是哈希，也不是地址
}

// 定义交易输出
type TXOutput struct {
	Value float64 //转账金额
	//PukKeyHash string  //锁定脚本，用地址来模拟
	PukKeyHash []byte //收款方的公钥哈希，注意：是哈希，不是公钥，不是地址
}

// 由于现在存储的字段是公钥的哈希，所以无法直接创建TXOutput
// 为了能够得到公钥哈希，需要处理一下，写一个Lock函数
func (output *TXOutput) Lock(address string) {
	//1.解码
	addressByte := base58.Decode(address) //25字节
	//2.截取出公钥哈希：（去除version，去除校验码）
	len := len(addressByte)
	pubKeyHash := addressByte[1 : len-4] // 去除version(1字节)，去除校验码(4字节)
	output.PukKeyHash = pubKeyHash
}

//给TXOutput提供一个创建的方法，否则无法调用Lock
func NewTXOutput(value float64, address string) *TXOutput {
	output := TXOutput{
		Value: value,
	}
	output.Lock(address)
	return &output
}

//设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

const reward = 12.5

// 实现一个函数，判断当前交易是否为挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	// 1.交易input只有一个
	// 2.交易id为空
	// 3.交易的index为-1
	if len(tx.TXInputs) == 1 {
		input := tx.TXInputs[0]
		if bytes.Equal(input.TXid, []byte{}) && input.Index == -1 {
			return true
		}
	}
	return false
}

// 2.提供创建交易方法(挖矿交易)
func NewCoinbaseTX(address string, data string) *Transaction {
	//挖矿交易特点
	//1.只有一个input
	//2.无需引用交易id
	//3.无需引用index
	//矿工由于挖矿时无需指定签名，所以这个PubKey字段可以由矿工自由填写数据，一般填写矿池的名字
	//签名先填写为空，后面创建完整交易后，最后做一次签名即可
	input := TXInput{[]byte{}, -1, nil, []byte(data)}
	//output := TXOutput{reward, address}
	//新的创建方法
	output := NewTXOutput(reward, address)

	//对于挖矿交易来说，只有一个input和一个output
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{*output}}
	tx.SetHash()
	return &tx
}

// 创建普通挖矿交易

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	// 1.创建交易后要进行数字签名->所以需要私钥->打开钱包"NewWallets()"
	ws := NewWallets()
	// 2.找到自己钱包，根据地址返回自己的wallet
	wallet := ws.WalletsMap[from]
	if wallet == nil {
		fmt.Println("没有找到该地址的钱包，交易创建失败")
		return nil
	}
	// 3.得到对应的公钥私钥
	pubKey := wallet.PubKey
	//privateKey := wallet.Private

	// 传递公钥的哈希，而不是地址
	pubKeyHash := HashPubKey(pubKey)

	// 1.找到最合理UTXO集合 map[string][]uint64
	utxos, resValue := bc.FindNeedUTXOs(pubKeyHash, amount)

	if resValue < amount {
		fmt.Println("与额不足，交易失败")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput
	// 2.创建交易输入，将UTXO逐一转成inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), nil, pubKey}
			inputs = append(inputs, input)
		}
	}

	// 创建交易输出
	//output := TXOutput{amount, to}
	output := NewTXOutput(amount, to)
	outputs = append(outputs, *output)

	if resValue > amount {
		//找零
		//outputs = append(outputs, TXOutput{resValue - amount, from})
		output = NewTXOutput(resValue-amount, from)
		outputs = append(outputs, *output)
	}
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return &tx
}
