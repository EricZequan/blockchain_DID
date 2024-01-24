package main

import (
	"fmt"
	"log"
)

//address是创建区块链的地址， nodeID是创建区块链的ID标识
func (cli *CLI) createBlockchain(address, nodeID string) {
	//检查address是否在钱包的已创建地址内
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	//创建一个区块链并且保存在本地数据库内
	bc := CreateBlockchain(address, nodeID)
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
