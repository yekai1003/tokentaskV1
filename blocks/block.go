package blocks

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"log"
	"strings"
)

type TaskInfoData struct {
	Task_ID string `json:"task_id"`
	Issuer  string `json:"issuer"`
	Worker  string `json:"task_user"`
	Bonus   int64  `json:"bonus"`
	Desc    string `json:"desc"`
	Comment string `json:"comment"`
	Status  uint8  `json:"status"`
}

var keyJson = `{"address":"1997f640f547af13b6acebaa7d176d14746e0c47","crypto":{"cipher":"aes-128-ctr","ciphertext":"753465dba334025383082cfe19e6b454ef03f5a412e4bf818444fda9ff99d3a1","cipherparams":{"iv":"ef419164186fc62367d797da8429884b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"83d4ff2bd2589d371cb2d89993502f607ffe9b00451ac0854bf32dc83d054718"},"mac":"f548202fa3be51cda79c7dc1daf07995a89a150ad6aaf0c05b79c92176c216e1"},"id":"e46cac91-f6e3-4b3e-8499-06a4d1b63207","version":3}`

const (
	contract_addr = "0xB17232bdf3245D75DBEf8301388cdc52B5e9ee3B"
)

var cli *ethclient.Client

//init函数当package被引用时，自动调用
func init() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
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

func QueryTask() ([]TaskInfoData, error) {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}
	tasks, err := instance.TaskQueryAll(nil)
	if err != nil {
		log.Panic("Failed to TaskQueryAll", err)
	}
	var taskinfodatas []TaskInfoData
	var taskinfo TaskInfoData
	var num int64 = 0
	for _, v := range tasks {
		taskinfo.Task_ID = strconv.FormatInt(num, 10)
		taskinfo.Bonus = v.Bonus.Int64()
		taskinfo.Comment = v.Comment
		taskinfo.Desc = v.Desc
		taskinfo.Issuer = v.Issuer
		taskinfo.Worker = v.Worker
		taskinfo.Status = v.Status
		taskinfodatas = append(taskinfodatas, taskinfo)
		num++
	}

	return taskinfodatas, err
}
