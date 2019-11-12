package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
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
	TXid  []byte //引用的交易id
	Index int64  //引用的交易output索引值
	Sig   string //解锁脚本，用地址来模拟
}

// 定义交易输出
type TXOutput struct {
	Value      float64 //转账金额
	PukKeyHash string  //锁定脚本，用地址来模拟
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
	//矿工由于挖矿时无需指定签名，所以这个sig字段可以由矿工自由填写数据，一般填写矿池的名字
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}

	//对于挖矿交易来说，只有一个input和一个output
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()
	return &tx
}

// 创建普通挖矿交易

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	// 1.找到最合理UTXO集合 map[string][]uint64
	utxos, resValue := bc.FindNeedUTXOs(from, amount)

	if resValue < amount {
		fmt.Println("与额不足，交易失败")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput
	// 2.创建交易输入，将UTXO逐一转成inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}

	// 创建交易输出
	output := TXOutput{amount, to}
	outputs = append(outputs, output)

	if resValue > amount {
		//找零
		outputs = append(outputs, TXOutput{resValue - amount, from})
	}
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return &tx
}
