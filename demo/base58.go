package main

import (
	"github.com/btcsuite/btcutil/base58"
	"fmt"
)

func main(){

	address := "1LfgYNozXixcV2E7MZ8gRStxXQehyQ1cWt"
	addressByte := base58.Decode(address) //25字节
	//2.截取出公钥哈希：（去除version，去除校验码）
	length := len(addressByte)

	fmt.Println("length=", length)
	fmt.Println("addressByte=", addressByte)
}
