package main

import (
	"fmt"

	"github.com/blockcypher/gobcy"
)

func main() {
	//START OMIT
	bcy := gobcy.API{"TESTTOKEN", "bcy", "test"}
	//Generate some addresses
	Pair1, _ := bcy.GenAddrKeychain()
	Pair2, _ := bcy.GenAddrKeychain()
	Pair3, _ := bcy.GenAddrKeychain()
	//Generate some blockcypher test coin
	bcy.Faucet(Pair1, 1e6)

	//Post a Payment Forwarding Request
	payment, err := bcy.CreatePayFwd(gobcy.PayFwd{Destination: Pair3.Address,
		ProcessAddr:    Pair2.Address,
		ProcessPercent: 0.2,
		CallbackURL:    ""})
	//Send some bcy to the payment forwarding address
	micro, err := bcy.SendMicro(gobcy.MicroTX{Priv: Pair1.Private,
		ToAddr: payment.InputAddr, Value: 100000})

	//Print out addresses
	fmt.Printf("Pair1 Address: %s\nPair2 Address: %s\nPair3 Address: %s\nPayment Forwarding Address: %s\n",
		Pair1.Address, Pair2.Address, Pair3.Address, payment.InputAddr)
	fmt.Printf("Microtransaction hash: %s\n", micro.Hash)
	//END OMIT
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
