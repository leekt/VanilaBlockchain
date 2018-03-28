package cli

import (
	"ScamList/core"
	"fmt"
)

func (cli *CLI) CreateWallet(nodeID string) {
	wallets, _ := core.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Println("Your new address :", address)
}
