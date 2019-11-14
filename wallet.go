package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
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
