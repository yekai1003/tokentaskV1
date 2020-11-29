package bcos

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"

	"log"
)

type TaskInfoData struct {
	Task_ID string `json:"task_id"`
	Issuer  string `json:"issuer"`
	Worker  string `json:"task_user"`
	Bonus   int64  `json:"bonus"`
	Desc    string `json:"task_name"`
	Comment string `json:"comment"`
	Status  uint8  `json:"task_status"`
}

const (
	contract_addr = "0xfcf507b2078d6ec9e01d8aac9bcdba30938bef5f"
)

var cli *client.Client

//init函数当package被引用时，自动调用
func init() {

	//os.Chdir("./dist")
	confs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		fmt.Println("Failed to ParseConfigFile", err)
		return
	}
	fmt.Println(confs[0])
	clii, err := client.Dial(&confs[0])
	if err != nil {
		fmt.Println("Failed to Dial", err)
		return
	}
	cli = clii
}

//实现eth的注册调用
func Bcos_Register(userid, passwd string) error {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}

	_, _, err = instance.Register(cli.GetTransactOpts(), userid, userid, passwd)
	if err != nil {
		log.Panic("Failed to Register", err)
	}
	return err
}

func Bcos_Login(userid, passwd string) (bool, error) {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}

	return instance.Login(cli.GetCallOpts(), userid, passwd)
}

//实现eth的挖矿调用
func Bcos_Mint(userid string, amount int64) error {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}

	_, _, err = instance.Mint(cli.GetTransactOpts(), userid, big.NewInt(amount))
	if err != nil {
		log.Panic("Failed to Mint", err)
	}
	return err
}

//实现eth的任务发布调用
func Bcos_Issue(who, pass, desc string, amount int64) error {
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}

	//Issue(opts *bind.TransactOpts, issuer string, pass string, desc string, bonus *big.Int)
	_, _, err = instance.Issue(cli.GetTransactOpts(), who, pass, desc, big.NewInt(amount))
	if err != nil {
		log.Panic("Failed to Issue", err)
	}
	return err
}

//实现eth的任务修改调用
func Bcos_Update(who, pass, comment string, taskID int64, status uint8) error {

	fmt.Println("who=", who, "pass=", pass, "taskid=", taskID, "status=", status)
	instance, err := NewTask(common.HexToAddress(contract_addr), cli)
	if err != nil {
		log.Panic("Failed to NewTask", err)
	}

	//Confirm(opts *bind.TransactOpts, issuer string, pass string, taskID *big.Int, comment string, st uint8)
	if status == 1 {
		_, _, err = instance.Take(cli.GetTransactOpts(), who, pass, big.NewInt(taskID))
		if err != nil {
			log.Panic("Failed to Take", err)
		}
	} else if status == 2 {
		_, _, err = instance.Commit(cli.GetTransactOpts(), who, pass, big.NewInt(taskID))
		if err != nil {
			log.Panic("Failed to Commit", err)
		}
	} else if status == 3 || status == 4 {
		if status == 4 {
			status = 1 //打回处理
		}
		_, _, err = instance.Confirm(cli.GetTransactOpts(), who, pass, big.NewInt(taskID), comment, status)
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
	issuers_, err := instance.TaskQryAllIssuer(nil)
	if err != nil {
		fmt.Println("failed to query TaskQryAllIssuer", err)
		return nil, err
	}

	workers_, err := instance.TaskQryAllWorker(nil)
	if err != nil {
		fmt.Println("failed to query TaskQryAllWorker", err)
		return nil, err
	}
	desc_, err := instance.TaskQryAllDesc(nil)
	if err != nil {
		fmt.Println("failed to query TaskQryAllDesc", err)
		return nil, err
	}
	comment_, err := instance.TaskQryAllComment(nil)
	if err != nil {
		fmt.Println("failed to query TaskQryAllComment", err)
		return nil, err
	}
	bonus_, err := instance.TaskQryAllBonus(nil)
	if err != nil {
		fmt.Println("failed to query TaskQryAllComment", err)
		return nil, err
	}
	status_, err := instance.TaskQryAllStatus(nil)
	if err != nil {
		fmt.Println("failed to query all TaskQryAllStatus", err)
		return nil, err
	}

	var taskinfodatas []TaskInfoData
	var taskinfo TaskInfoData
	//var num int64 = 0
	fmt.Printf("task count[%d]\n", len(issuers_))
	for k, v := range issuers_ {
		taskinfo.Task_ID = strconv.FormatInt(int64(k), 10)
		taskinfo.Bonus = bonus_[k].Int64()
		taskinfo.Comment = comment_[k]
		taskinfo.Desc = desc_[k]
		taskinfo.Issuer = v
		taskinfo.Worker = workers_[k]
		taskinfo.Status = status_[k]
		taskinfodatas = append(taskinfodatas, taskinfo)

	}

	return taskinfodatas, err
}
