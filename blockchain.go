package main

import (
	"block/block1/bolt"
	"fmt"
	"log"
)

// 引入区块链
type BlockChain struct {
	//定义一个区块链数组
	//blocks []*Block
	db   *bolt.DB
	tail []byte //存储最后一个区块的哈希
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"

//定义一个区块链
func NewBlockChain(address string) *BlockChain {
	//genesisBlock := GenesisBlock()
	//return &BlockChain{
	//	[]*Block{genesisBlock},
	//}
	var lastHash []byte

	// 1.打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	//defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败")
	}

	//操作数据库(改写)
	db.Update(func(tx *bolt.Tx) error {
		// 2.找到抽屉bucket(没有就创建)
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			// 没有抽屉，需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("创建抽屉失败")
			}
			// 创建一个创世块，并作为第一个区块加入到区块链中
			genesisBlock := GenesisBlock(address)
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash

		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

//定义一个创世链
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTX(address, "我是创世者")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	db := bc.db
	lastHash := bc.tail

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket为空")
		}
		// 创建新的区块
		block := NewBlock(txs, lastHash)
		// 添加到区块链db中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)
		bc.tail = block.Hash
		return nil
	})
}

// 找到指定地址的所有UTXO
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	// 定义一个map来保存消费过的output，key是这个output的交易id，value是这个交易中索引的数组
	// map[交易id][]int64
	spentOutputs := make(map[string][]int64)
	it := bc.NewIterator()
	for {
		// 1.遍历区块
		block := it.Next()
		// 2.遍历交易
		for _, tx := range block.Transactions {
			fmt.Printf("current txid : %x\n", tx.TXID)
		OUTPUT:
			// 3.遍历output，找到和自己相关的utxo(在添加output之前)
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index: %d\n", i)
				// 在这里做一个过滤，将所有消耗过的outputs和当前的所即将添加output对比一下
				//如果相同则跳过，否则添加
				// 如果当前交易id存在于已经表示的map，那么说明这个交易里面有消耗过的output
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							// 当前准备添加output已经消耗过了，不需要加了
							continue OUTPUT
						}
					}
				}

				// 这个output和我们的地址系统，满足条件，加到返回
				if output.PukKeyHash == address {
					UTXO = append(UTXO, output)
				}
			}

			// 如果当前交易是挖矿交易，那么不做遍历，直接跳过
			if !tx.IsCoinbase() {
				// 4.遍历input，找到自己花费过的utxo的集合(把自己消耗过的标示出来)
				for _, input := range tx.TXInputs {
					// 判断一下当前这个input和目标(李四)是否一致，如果相同，说明这个是李四消耗过的output，就加进来
					if input.Sig == address {
						indexArray := spentOutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
					}
				}
			} else {
				fmt.Println("这是挖矿交易，不做input遍历")
			}

		}
		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历完成，退出！")
			break
		}
	}
	return UTXO
}

// 根据需求找到合理的utxo
func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]int64, float64) {
	// 找到的合理的uyxos集合
	var utxos map[string][]int64
	// 找到的
	var calc float64
	//todo
	spentOutputs := make(map[string][]int64)
	it := bc.NewIterator()
	for {
		// 1.遍历区块
		block := it.Next()
		// 2.遍历交易
		for _, tx := range block.Transactions {
			fmt.Printf("current txid : %x\n", tx.TXID)
		OUTPUT:
			// 3.遍历output，找到和自己相关的utxo(在添加output之前)
			for i, output := range tx.TXOutputs {
				fmt.Printf("current index: %d\n", i)
				// 在这里做一个过滤，将所有消耗过的outputs和当前的所即将添加output对比一下
				//如果相同则跳过，否则添加
				// 如果当前交易id存在于已经表示的map，那么说明这个交易里面有消耗过的output
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							// 当前准备添加output已经消耗过了，不需要加了
							continue OUTPUT
						}
					}
				}

				// 这个output和我们的地址系统，满足条件，加到返回
				if output.PukKeyHash == from {
					if calc < amount {
						utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], int64(i))
						calc += output.Value
						if calc >= amount {
							fmt.Printf("找到了满足的金额：%f\n", calc)
							return utxos, calc
						}
					}

				}
			}

			// 如果当前交易是挖矿交易，那么不做遍历，直接跳过
			if !tx.IsCoinbase() {
				// 4.遍历input，找到自己花费过的utxo的集合(把自己消耗过的标示出来)
				for _, input := range tx.TXInputs {
					// 判断一下当前这个input和目标(李四)是否一致，如果相同，说明这个是李四消耗过的output，就加进来
					if input.Sig == from {
						indexArray := spentOutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
					}
				}
			} else {
				fmt.Println("这是挖矿交易，不做input遍历")
			}

		}
		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历完成，退出！")
			break
		}
	}

	return utxos, calc
}
