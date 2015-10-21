package main

import (
	"fmt"

	"github.com/blockcypher/gobcy"
)

func main() {
	//Set Blockcy Configuration: Coin/Chain
	btc := gobcy.API{"TESTTOKEN", "btc", "main"}
	//Get Block by height
	block, _ := btc.GetBlock(294322, "")
	fmt.Printf("%+v\n", block)
	//Get latest details on bitcoin blockchain
	//chain, _ := btc.GetChain()
	//fmt.Printf("%+v\n", chain)
	return
}
