package main

import (
	"fmt"

	"github.com/acityinohio/blockcy"
)

func main() {
	blockcy.Config.Token = "e212e91ac4d218cbc18f7eb3975122e3"
	//START OMIT
	blockcy.Config.Coin, blockcy.Config.Chain = "bcy", "test"
	//Generate some addresses
	Wallet1, _ := blockcy.GenAddrPair()
	Wallet2, _ := blockcy.GenAddrPair()
	Wallet3, _ := blockcy.GenAddrPair()
	//Generate some blockcypher test coin
	Wallet1.Faucet(1e6)

	//Post a Payment Forwarding Request
	payment, err := blockcy.PostPayment(blockcy.Payment{Destination: Wallet3.Address,
		ProcessAddr:    Wallet2.Address,
		ProcessPercent: 0.2,
		CallbackUrl:    ""})
	//Send some bcy to the payment forwarding address
	micro, err := blockcy.SendMicro(blockcy.Micro{Private: Wallet1.Private,
		ToAddr: payment.InputAddr, Value: 100000})

	//Print out addresses
	fmt.Printf("Wallet1 Address: %s\nWallet2 Address: %s\nWallet3 Address: %s\nPayment Forwarding Address: %s\n",
		Wallet1.Address, Wallet2.Address, Wallet3.Address, payment.InputAddr)
	fmt.Printf("Microtransaction hash: %s\n", micro.Hash)
	//END OMIT
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
