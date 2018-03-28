package cli

import (
	"ScamList/core"
	"fmt"
	"log"
)

func (cli *CLI) GetBalance(address, nodeID string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR : Address is not valid")
	}

	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	defer bc.Db.Close()

	balance := 0
	pubKeyHash := core.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of %s : %d\n", address, balance)
}
