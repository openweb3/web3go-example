package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/signers"
)

func main() {
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	client := web3go.MustNewClientWithOption("http://localhost:7545", *new(web3go.ClientOption).WithLooger(os.Stdout).WithSignerManager(sm))

	clientForContract, signFunc := client.ToClientForContract()

	opt := &bind.TransactOpts{
		From:   sm.List()[0].Address(),
		Signer: signFunc,
	}

	// Deploy
	addr, tx, token, err := DeployMyERC20Token(opt, clientForContract, big.NewInt(1000), "abc", 18, "ABC")
	if err != nil {
		panic(err)
	}
	fmt.Printf("DeployMyERC20Token, addr: %v, tx: %v, token: %v\n", addr, tx, token)

	// Read
	balance, err := token.BalanceOf(nil, opt.From)
	if err != nil {
		panic(err)
	}
	fmt.Printf("balance of %v: %v\n", opt.From, balance)

	// Write
	tx, err = token.Transfer(opt, common.HexToAddress(""), big.NewInt(5))
	if err != nil {
		panic(err)
	}
	fmt.Printf("sent 5 done, txhash: %v\n", tx.Hash())

	balance, err = token.BalanceOf(nil, opt.From)
	if err != nil {
		panic(err)
	}
	fmt.Printf("after sent, balance of %v: %v\n", opt.From, balance)
}
