package main

import (
	"fmt"
	//import my BlockCy repo
	"github.com/acityinohio/blockcy"
)

func main() {
	//Set Blockcy Configuration: Coin/Chain
	blockcy.Config.Coin, blockcy.Config.Chain = "btc", "main"
	//Get Block by height
	block, _ := blockcy.GetBlock(294322, "")
	fmt.Printf("%+v\n", block)
	//Get latest details on bitcoin blockchain
	//chain, _ := blockcy.GetChain()
	//fmt.Printf("%+v\n", chain)
	return
}
