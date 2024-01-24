package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
)

func (cli *CLI) getdataamount(address, nodeID string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	var count int
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	bc := NewBlockchain(nodeID) //返回区块链中最后一个区块
	defer bc.db.Close()
	bci := bc.Iterator()
	for {
		block := bci.Next()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, pubKeyHash) == 0 && tx.Form == "store" {
				count += 100
			}
		}
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	fmt.Printf("The address: %s has stored %d MB data \n", address, count)
}
