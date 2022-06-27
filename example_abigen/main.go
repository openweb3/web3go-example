package main

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/signers"
)

func main() {
	mnemonic := "crisp shove million stem shiver side hospital split play lottery join vintage"
	sm := signers.MustNewSignerManagerByMnemonic(mnemonic, 10, nil)
	client, err := web3go.NewClientWithOption("http://localhost:7545", *new(web3go.ClientOption).WithLooger(os.Stdout).WithSignerManager(sm))
	if err != nil {
		panic(err)
	}

	chainId, err := client.Eth.ChainId()
	if err != nil {
		panic(err)
	}

	signFunc := func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		for _, s := range sm.List() {
			if s.Address() == a {
				return s.SignTransaction(t, new(big.Int).SetUint64(*chainId))
			}
		}
		return nil, errors.New("not found signer")
	}

	opt := &bind.TransactOpts{
		From:   sm.List()[0].Address(),
		Signer: signFunc,
	}

	// Deploy
	addr, tx, token, err := DeployMyERC20Token(opt, client.ToClientForContract(), big.NewInt(1000), "abc", 18, "ABC")
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
