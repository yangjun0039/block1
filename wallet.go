package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"github.com/btcsuite/btcutil/base58"
	"bytes"
)

//这里钱包是一个结构，每一个钱包保存了公钥，私钥

type Wallet struct {
	Private *ecdsa.PrivateKey //私钥
	// PubKey *ecdsa.PublicKey
	// 约定，这里的PubKey不存储原始的公钥，而是存储X和Y拼接的字符串，在校验端重新拆分(参考r,s传递)
	PubKey []byte
}

//创建钱包
func NewWallet() *Wallet {
	//创建曲线
	curve := elliptic.P256()
	//生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//生成公钥
	pubKeyOrigin := privateKey.PublicKey
	//拼接X，Y
	pubKey := append(pubKeyOrigin.X.Bytes(), pubKeyOrigin.Y.Bytes()...)
	return &Wallet{privateKey, pubKey}
}

//生成地址
func (w *Wallet) NewAddress() string{
	pubKey := w.PubKey

	rip160HashValue := HashPubKey(pubKey)

	version := byte(00)
	payload := append([]byte{version}, rip160HashValue...)

	//checksum
	checkCode := CheckSum(payload)

	//25字节
	payload = append(payload, checkCode...)

	//go有一个库叫做btcd，这个是go语言实现的比特币全节点源码
	address := base58.Encode(payload)
	return address
}

func HashPubKey(data []byte) []byte{
	hash := sha256.Sum256(data)
	// 理解为编码器
	hash160hasger := ripemd160.New()
	_,err := hash160hasger.Write(hash[:])
	if err != nil{
		log.Panic(err)
	}
	//返回rip160的哈希结果
	rip160HashValue := hash160hasger.Sum(nil)
	return rip160HashValue
}

func CheckSum(data []byte) []byte {
	//两次hash256
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])
	//前四字节检验码
	checkCode := hash2[:4]
	return checkCode
}

func IsValidAddress(address string) bool{
	//1.解码
	addressByte := base58.Decode(address)
	if len(addressByte) < 4{
		return false
	}
	//2.取数据
	payload := addressByte[:len(addressByte)-4]
	checkSUm1 := addressByte[len(addressByte)-4:]
	//3.做CheckSum
	checkSUm2 := CheckSum(payload)
	//4.比较
	return bytes.Equal(checkSUm1, checkSUm2)
}