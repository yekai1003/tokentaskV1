package blocks

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"log"
	"strings"
)

var keyJson = `{"address":"9fdda261f18cdbf51efc1e92c97ea5103eb82e40","crypto":{"cipher":"aes-128-ctr","ciphertext":"8640d3fd7b4591fbbc5651139e72cfb18c950a521aede7ed280c921444886c4a","cipherparams":{"iv":"77a6657f862cb60703f0f01b8f77f0d7"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"5246b71667604b618ccc9c99162bb81bea41849c997ceb2314763345e5b8ab56"},"mac":"46952de48f181fea614e089621330acdcf78f7d4f116460b21b65dc83af4a418"},"id":"017a76fe-20de-40fe-a1fd-61f96afd03cf","version":3}`

const (
	contract_addr = "0x498a8114402bF88b5f80B9FBa04E03ED933214BB"
)

var cli *ethclient.Client

//init函数当package被引用时，自动调用
func init() {
	client, err := ethclient.Dial("http://192.168.1.17:8545")
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	cli = client
}

//实现eth的注册调用
func Eth_Register(userid, passwd string) error {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}
	//Register(opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, error)
	auth, err := bind.NewTransactor(strings.NewReader(keyJson), "123")
	if err != nil {
		log.Panic("failed to NewTransactor", err)
	}

	_, err = instance.Register(auth, userid, userid, passwd)
	if err != nil {
		log.Panic("Failed to Register", err)
	}
	return err
}

func Eth_Login(userid, passwd string) (bool, error) {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}
	//Register(opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, error)
	// auth, err := bind.NewTransactor(strings.NewReader(keyJson), "123")
	// if err != nil {
	// 	log.Panic("failed to NewTransactor", err)
	// }
	//func (_Task *TaskCaller) Login(opts *bind.CallOpts, userid string, passwd string) (bool, error)

	opts := bind.CallOpts{
		From: common.HexToAddress("0x9fdda261f18cdbf51efc1e92c97ea5103eb82e40"),
	}
	return instance.Login(&opts, userid, passwd)
}

//实现eth的挖矿调用
func Eth_Mint(userid string, amount int64) error {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}
	//Register(opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, error)
	auth, err := bind.NewTransactor(strings.NewReader(keyJson), "123")
	if err != nil {
		log.Panic("failed to NewTransactor", err)
	}

	_, err = instance.Mint(auth, userid, big.NewInt(amount))
	if err != nil {
		log.Panic("Failed to Mint", err)
	}
	return err
}

//实现eth的任务发布调用
func Eth_Issue(who, pass, desc string, amount int64) error {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}
	//Register(opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, error)
	auth, err := bind.NewTransactor(strings.NewReader(keyJson), "123")
	if err != nil {
		log.Panic("failed to NewTransactor", err)
	}
	//Issue(opts *bind.TransactOpts, issuer string, pass string, desc string, bonus *big.Int)
	_, err = instance.Issue(auth, who, pass, desc, big.NewInt(amount))
	if err != nil {
		log.Panic("Failed to Issue", err)
	}
	return err
}

//实现eth的任务修改调用
func Eth_Update(who, pass, comment string, taskID int64, status uint8) error {

	fmt.Println("who=", who, "pass=", pass, "taskid=", taskID, "status=", status)
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}
	//Register(opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, error)
	auth, err := bind.NewTransactor(strings.NewReader(keyJson), "123")
	if err != nil {
		log.Panic("failed to NewTransactor", err)
	}
	//Take(opts *bind.TransactOpts, who string, pass string, taskID *big.Int)
	//Commit(opts *bind.TransactOpts, who string, pass string, taskID *big.Int)

	//Confirm(opts *bind.TransactOpts, issuer string, pass string, taskID *big.Int, comment string, st uint8)
	if status == 1 {
		_, err = instance.Take(auth, who, pass, big.NewInt(taskID))
		if err != nil {
			log.Panic("Failed to Take", err)
		}
	} else if status == 2 {
		_, err = instance.Commit(auth, who, pass, big.NewInt(taskID))
		if err != nil {
			log.Panic("Failed to Commit", err)
		}
	} else if status == 3 || status == 4 {
		if status == 4 {
			status = 1 //打回处理
		}
		_, err = instance.Confirm(auth, who, pass, big.NewInt(taskID), comment, status)
		if err != nil {
			log.Panic("Failed to Confirm", err)
		}
	}

	return err
}
