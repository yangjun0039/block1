package main

import (
	"io/ioutil"
	"bytes"
	"encoding/gob"
	"log"
	"crypto/elliptic"
)

//定义一个Wallets结构，它保存所有的wallet以及他的地址
type Wallets struct {
	//map[地址]钱包
	WalletsMap map[string]*Wallet
}

//创建方法
func NewWallets() *Wallets {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)
	ws.loadFile()
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()
	//var wallets Wallets
	//wallets.WalletsMap = make(map[string]*Wallet)
	ws.WalletsMap[address] = wallet
	ws.saveToFile()
	return address
}

//保存方法，把新建的wallet添CreateWallet加进去
func (ws *Wallets) saveToFile() {
	var buffer bytes.Buffer

	// panic: gob: type not registered for interface: elliptic.p256Curve
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile("wallet.dat", buffer.Bytes(), 0600)
}

//读取文件方法，把所有的wallet读出来
func (ws *Wallets) loadFile(){
	// 读取内容
	content,err := ioutil.ReadFile("wallet.dat")
	if err != nil{
		log.Panic(err)
	}
	// 解码
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(content))
	var wsLocal Wallets
	err = decoder.Decode(&wsLocal)
	if err != nil{
		log.Panic(err)
	}
	//ws = &wsLocal
	//对结构来说，里面有map的，要指定赋值，不能直接赋值
	ws.WalletsMap = wsLocal.WalletsMap

}

func (ws *Wallets) GetAllAddress() []string{
	var addresses []string
	// 遍历钱包，将所有key取出来返回
	for address := range ws.WalletsMap{
		addresses = append(addresses, address)
	}
	return addresses
}
