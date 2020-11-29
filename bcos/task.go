// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bcos

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskQryAllDesc\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskQryAllIssuer\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"username\",\"type\":\"string\"},{\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskQryAllWorker\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskQryAllComment\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"who\",\"type\":\"string\"},{\"name\":\"pass\",\"type\":\"string\"},{\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"take\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"issuer\",\"type\":\"string\"},{\"name\":\"pass\",\"type\":\"string\"},{\"name\":\"desc\",\"type\":\"string\"},{\"name\":\"bonus\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"who\",\"type\":\"string\"},{\"name\":\"pass\",\"type\":\"string\"},{\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskQryAllStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taskQryAllBonus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"taskQueryOne\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"issuer\",\"type\":\"string\"},{\"name\":\"pass\",\"type\":\"string\"},{\"name\":\"taskID\",\"type\":\"uint256\"},{\"name\":\"comment\",\"type\":\"string\"},{\"name\":\"st\",\"type\":\"uint8\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_sym\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Task is an auto generated Go binding around a Solidity contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around a Solidity contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around a Solidity contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around a Solidity contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string who) constant returns(uint256)
func (_Task *TaskCaller) BalanceOf(opts *bind.CallOpts, who string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string who) constant returns(uint256)
func (_Task *TaskSession) BalanceOf(who string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string who) constant returns(uint256)
func (_Task *TaskCallerSession) BalanceOf(who string) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, who)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string userid, string passwd) constant returns(bool)
func (_Task *TaskCaller) Login(opts *bind.CallOpts, userid string, passwd string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "login", userid, passwd)
	return *ret0, err
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string userid, string passwd) constant returns(bool)
func (_Task *TaskSession) Login(userid string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, userid, passwd)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string userid, string passwd) constant returns(bool)
func (_Task *TaskCallerSession) Login(userid string, passwd string) (bool, error) {
	return _Task.Contract.Login(&_Task.CallOpts, userid, passwd)
}

// TaskQryAllBonus is a free data retrieval call binding the contract method 0x9987f845.
//
// Solidity: function taskQryAllBonus() constant returns(uint256[])
func (_Task *TaskCaller) TaskQryAllBonus(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQryAllBonus")
	return *ret0, err
}

// TaskQryAllBonus is a free data retrieval call binding the contract method 0x9987f845.
//
// Solidity: function taskQryAllBonus() constant returns(uint256[])
func (_Task *TaskSession) TaskQryAllBonus() ([]*big.Int, error) {
	return _Task.Contract.TaskQryAllBonus(&_Task.CallOpts)
}

// TaskQryAllBonus is a free data retrieval call binding the contract method 0x9987f845.
//
// Solidity: function taskQryAllBonus() constant returns(uint256[])
func (_Task *TaskCallerSession) TaskQryAllBonus() ([]*big.Int, error) {
	return _Task.Contract.TaskQryAllBonus(&_Task.CallOpts)
}

// TaskQryAllComment is a free data retrieval call binding the contract method 0x583a845e.
//
// Solidity: function taskQryAllComment() constant returns(string[])
func (_Task *TaskCaller) TaskQryAllComment(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQryAllComment")
	return *ret0, err
}

// TaskQryAllComment is a free data retrieval call binding the contract method 0x583a845e.
//
// Solidity: function taskQryAllComment() constant returns(string[])
func (_Task *TaskSession) TaskQryAllComment() ([]string, error) {
	return _Task.Contract.TaskQryAllComment(&_Task.CallOpts)
}

// TaskQryAllComment is a free data retrieval call binding the contract method 0x583a845e.
//
// Solidity: function taskQryAllComment() constant returns(string[])
func (_Task *TaskCallerSession) TaskQryAllComment() ([]string, error) {
	return _Task.Contract.TaskQryAllComment(&_Task.CallOpts)
}

// TaskQryAllDesc is a free data retrieval call binding the contract method 0x061d5b65.
//
// Solidity: function taskQryAllDesc() constant returns(string[])
func (_Task *TaskCaller) TaskQryAllDesc(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQryAllDesc")
	return *ret0, err
}

// TaskQryAllDesc is a free data retrieval call binding the contract method 0x061d5b65.
//
// Solidity: function taskQryAllDesc() constant returns(string[])
func (_Task *TaskSession) TaskQryAllDesc() ([]string, error) {
	return _Task.Contract.TaskQryAllDesc(&_Task.CallOpts)
}

// TaskQryAllDesc is a free data retrieval call binding the contract method 0x061d5b65.
//
// Solidity: function taskQryAllDesc() constant returns(string[])
func (_Task *TaskCallerSession) TaskQryAllDesc() ([]string, error) {
	return _Task.Contract.TaskQryAllDesc(&_Task.CallOpts)
}

// TaskQryAllIssuer is a free data retrieval call binding the contract method 0x32e49838.
//
// Solidity: function taskQryAllIssuer() constant returns(string[])
func (_Task *TaskCaller) TaskQryAllIssuer(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQryAllIssuer")
	return *ret0, err
}

// TaskQryAllIssuer is a free data retrieval call binding the contract method 0x32e49838.
//
// Solidity: function taskQryAllIssuer() constant returns(string[])
func (_Task *TaskSession) TaskQryAllIssuer() ([]string, error) {
	return _Task.Contract.TaskQryAllIssuer(&_Task.CallOpts)
}

// TaskQryAllIssuer is a free data retrieval call binding the contract method 0x32e49838.
//
// Solidity: function taskQryAllIssuer() constant returns(string[])
func (_Task *TaskCallerSession) TaskQryAllIssuer() ([]string, error) {
	return _Task.Contract.TaskQryAllIssuer(&_Task.CallOpts)
}

// TaskQryAllStatus is a free data retrieval call binding the contract method 0x936ed972.
//
// Solidity: function taskQryAllStatus() constant returns(uint8[])
func (_Task *TaskCaller) TaskQryAllStatus(opts *bind.CallOpts) ([]uint8, error) {
	var (
		ret0 = new([]uint8)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQryAllStatus")
	return *ret0, err
}

// TaskQryAllStatus is a free data retrieval call binding the contract method 0x936ed972.
//
// Solidity: function taskQryAllStatus() constant returns(uint8[])
func (_Task *TaskSession) TaskQryAllStatus() ([]uint8, error) {
	return _Task.Contract.TaskQryAllStatus(&_Task.CallOpts)
}

// TaskQryAllStatus is a free data retrieval call binding the contract method 0x936ed972.
//
// Solidity: function taskQryAllStatus() constant returns(uint8[])
func (_Task *TaskCallerSession) TaskQryAllStatus() ([]uint8, error) {
	return _Task.Contract.TaskQryAllStatus(&_Task.CallOpts)
}

// TaskQryAllWorker is a free data retrieval call binding the contract method 0x51db743f.
//
// Solidity: function taskQryAllWorker() constant returns(string[])
func (_Task *TaskCaller) TaskQryAllWorker(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQryAllWorker")
	return *ret0, err
}

// TaskQryAllWorker is a free data retrieval call binding the contract method 0x51db743f.
//
// Solidity: function taskQryAllWorker() constant returns(string[])
func (_Task *TaskSession) TaskQryAllWorker() ([]string, error) {
	return _Task.Contract.TaskQryAllWorker(&_Task.CallOpts)
}

// TaskQryAllWorker is a free data retrieval call binding the contract method 0x51db743f.
//
// Solidity: function taskQryAllWorker() constant returns(string[])
func (_Task *TaskCallerSession) TaskQryAllWorker() ([]string, error) {
	return _Task.Contract.TaskQryAllWorker(&_Task.CallOpts)
}

// TaskQueryOne is a free data retrieval call binding the contract method 0xb8cf744a.
//
// Solidity: function taskQueryOne(uint256 taskID) constant returns(string, string, string, uint256, uint8)
func (_Task *TaskCaller) TaskQueryOne(opts *bind.CallOpts, taskID *big.Int) (string, string, string, *big.Int, uint8, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(*big.Int)
		ret4 = new(uint8)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Task.contract.Call(opts, out, "taskQueryOne", taskID)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// TaskQueryOne is a free data retrieval call binding the contract method 0xb8cf744a.
//
// Solidity: function taskQueryOne(uint256 taskID) constant returns(string, string, string, uint256, uint8)
func (_Task *TaskSession) TaskQueryOne(taskID *big.Int) (string, string, string, *big.Int, uint8, error) {
	return _Task.Contract.TaskQueryOne(&_Task.CallOpts, taskID)
}

// TaskQueryOne is a free data retrieval call binding the contract method 0xb8cf744a.
//
// Solidity: function taskQueryOne(uint256 taskID) constant returns(string, string, string, uint256, uint8)
func (_Task *TaskCallerSession) TaskQueryOne(taskID *big.Int) (string, string, string, *big.Int, uint8, error) {
	return _Task.Contract.TaskQueryOne(&_Task.CallOpts, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactor) Commit(opts *bind.TransactOpts, who string, pass string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "commit", who, pass, taskID)
}

func (_Task *TaskTransactor) AsyncCommit(handler func(*types.Receipt, error), opts *bind.TransactOpts, who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "commit", who, pass, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string who, string pass, uint256 taskID) returns()
func (_Task *TaskSession) Commit(who string, pass string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, who, pass, taskID)
}

func (_Task *TaskSession) AsyncCommit(handler func(*types.Receipt, error), who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncCommit(handler, &_Task.TransactOpts, who, pass, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Commit(who string, pass string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, who, pass, taskID)
}

func (_Task *TaskTransactorSession) AsyncCommit(handler func(*types.Receipt, error), who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncCommit(handler, &_Task.TransactOpts, who, pass, taskID)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string pass, uint256 taskID, string comment, uint8 st) returns()
func (_Task *TaskTransactor) Confirm(opts *bind.TransactOpts, issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "confirm", issuer, pass, taskID, comment, st)
}

func (_Task *TaskTransactor) AsyncConfirm(handler func(*types.Receipt, error), opts *bind.TransactOpts, issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "confirm", issuer, pass, taskID, comment, st)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string pass, uint256 taskID, string comment, uint8 st) returns()
func (_Task *TaskSession) Confirm(issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, pass, taskID, comment, st)
}

func (_Task *TaskSession) AsyncConfirm(handler func(*types.Receipt, error), issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, error) {
	return _Task.Contract.AsyncConfirm(handler, &_Task.TransactOpts, issuer, pass, taskID, comment, st)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string pass, uint256 taskID, string comment, uint8 st) returns()
func (_Task *TaskTransactorSession) Confirm(issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, pass, taskID, comment, st)
}

func (_Task *TaskTransactorSession) AsyncConfirm(handler func(*types.Receipt, error), issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, error) {
	return _Task.Contract.AsyncConfirm(handler, &_Task.TransactOpts, issuer, pass, taskID, comment, st)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string pass, string desc, uint256 bonus) returns()
func (_Task *TaskTransactor) Issue(opts *bind.TransactOpts, issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "issue", issuer, pass, desc, bonus)
}

func (_Task *TaskTransactor) AsyncIssue(handler func(*types.Receipt, error), opts *bind.TransactOpts, issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "issue", issuer, pass, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string pass, string desc, uint256 bonus) returns()
func (_Task *TaskSession) Issue(issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, pass, desc, bonus)
}

func (_Task *TaskSession) AsyncIssue(handler func(*types.Receipt, error), issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncIssue(handler, &_Task.TransactOpts, issuer, pass, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string pass, string desc, uint256 bonus) returns()
func (_Task *TaskTransactorSession) Issue(issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, pass, desc, bonus)
}

func (_Task *TaskTransactorSession) AsyncIssue(handler func(*types.Receipt, error), issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncIssue(handler, &_Task.TransactOpts, issuer, pass, desc, bonus)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns()
func (_Task *TaskTransactor) Mint(opts *bind.TransactOpts, to string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "mint", to, value)
}

func (_Task *TaskTransactor) AsyncMint(handler func(*types.Receipt, error), opts *bind.TransactOpts, to string, value *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns()
func (_Task *TaskSession) Mint(to string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, to, value)
}

func (_Task *TaskSession) AsyncMint(handler func(*types.Receipt, error), to string, value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncMint(handler, &_Task.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns()
func (_Task *TaskTransactorSession) Mint(to string, value *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, to, value)
}

func (_Task *TaskTransactorSession) AsyncMint(handler func(*types.Receipt, error), to string, value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncMint(handler, &_Task.TransactOpts, to, value)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string userid, string username, string passwd) returns()
func (_Task *TaskTransactor) Register(opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "register", userid, username, passwd)
}

func (_Task *TaskTransactor) AsyncRegister(handler func(*types.Receipt, error), opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "register", userid, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string userid, string username, string passwd) returns()
func (_Task *TaskSession) Register(userid string, username string, passwd string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, userid, username, passwd)
}

func (_Task *TaskSession) AsyncRegister(handler func(*types.Receipt, error), userid string, username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.AsyncRegister(handler, &_Task.TransactOpts, userid, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string userid, string username, string passwd) returns()
func (_Task *TaskTransactorSession) Register(userid string, username string, passwd string) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, userid, username, passwd)
}

func (_Task *TaskTransactorSession) AsyncRegister(handler func(*types.Receipt, error), userid string, username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.AsyncRegister(handler, &_Task.TransactOpts, userid, username, passwd)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactor) Take(opts *bind.TransactOpts, who string, pass string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.contract.Transact(opts, "take", who, pass, taskID)
}

func (_Task *TaskTransactor) AsyncTake(handler func(*types.Receipt, error), opts *bind.TransactOpts, who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.AsyncTransact(opts, handler, "take", who, pass, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string who, string pass, uint256 taskID) returns()
func (_Task *TaskSession) Take(who string, pass string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, who, pass, taskID)
}

func (_Task *TaskSession) AsyncTake(handler func(*types.Receipt, error), who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTake(handler, &_Task.TransactOpts, who, pass, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Take(who string, pass string, taskID *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, who, pass, taskID)
}

func (_Task *TaskTransactorSession) AsyncTake(handler func(*types.Receipt, error), who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.AsyncTake(handler, &_Task.TransactOpts, who, pass, taskID)
}
