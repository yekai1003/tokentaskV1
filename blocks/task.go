// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blocks

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	Issuer   string
	Worker   string
	Bonus    *big.Int
	Desc     string
	Tasktime *big.Int
	Comment  string
	Status   uint8
}

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_sym\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"who\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"who\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pass\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pass\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"st\",\"type\":\"uint8\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pass\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"passwd\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"who\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pass\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"take\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskQueryAll\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tasktime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structTaskInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"taskQueryOne\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"issuer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"worker\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bonus\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tasktime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structTaskInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Task is an auto generated Go binding around an Ethereum contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
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
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
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

// TaskQueryAll is a free data retrieval call binding the contract method 0x251a4e4d.
//
// Solidity: function taskQueryAll() constant returns([]TaskInfo)
func (_Task *TaskCaller) TaskQueryAll(opts *bind.CallOpts) ([]TaskInfo, error) {
	var (
		ret0 = new([]TaskInfo)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQueryAll")
	return *ret0, err
}

// TaskQueryAll is a free data retrieval call binding the contract method 0x251a4e4d.
//
// Solidity: function taskQueryAll() constant returns([]TaskInfo)
func (_Task *TaskSession) TaskQueryAll() ([]TaskInfo, error) {
	return _Task.Contract.TaskQueryAll(&_Task.CallOpts)
}

// TaskQueryAll is a free data retrieval call binding the contract method 0x251a4e4d.
//
// Solidity: function taskQueryAll() constant returns([]TaskInfo)
func (_Task *TaskCallerSession) TaskQueryAll() ([]TaskInfo, error) {
	return _Task.Contract.TaskQueryAll(&_Task.CallOpts)
}

// TaskQueryOne is a free data retrieval call binding the contract method 0xb8cf744a.
//
// Solidity: function taskQueryOne(uint256 taskID) constant returns(TaskInfo)
func (_Task *TaskCaller) TaskQueryOne(opts *bind.CallOpts, taskID *big.Int) (TaskInfo, error) {
	var (
		ret0 = new(TaskInfo)
	)
	out := ret0
	err := _Task.contract.Call(opts, out, "taskQueryOne", taskID)
	return *ret0, err
}

// TaskQueryOne is a free data retrieval call binding the contract method 0xb8cf744a.
//
// Solidity: function taskQueryOne(uint256 taskID) constant returns(TaskInfo)
func (_Task *TaskSession) TaskQueryOne(taskID *big.Int) (TaskInfo, error) {
	return _Task.Contract.TaskQueryOne(&_Task.CallOpts, taskID)
}

// TaskQueryOne is a free data retrieval call binding the contract method 0xb8cf744a.
//
// Solidity: function taskQueryOne(uint256 taskID) constant returns(TaskInfo)
func (_Task *TaskCallerSession) TaskQueryOne(taskID *big.Int) (TaskInfo, error) {
	return _Task.Contract.TaskQueryOne(&_Task.CallOpts, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactor) Commit(opts *bind.TransactOpts, who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "commit", who, pass, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string who, string pass, uint256 taskID) returns()
func (_Task *TaskSession) Commit(who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, who, pass, taskID)
}

// Commit is a paid mutator transaction binding the contract method 0x89dc99d2.
//
// Solidity: function commit(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Commit(who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Commit(&_Task.TransactOpts, who, pass, taskID)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string pass, uint256 taskID, string comment, uint8 st) returns()
func (_Task *TaskTransactor) Confirm(opts *bind.TransactOpts, issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "confirm", issuer, pass, taskID, comment, st)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string pass, uint256 taskID, string comment, uint8 st) returns()
func (_Task *TaskSession) Confirm(issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, pass, taskID, comment, st)
}

// Confirm is a paid mutator transaction binding the contract method 0xfc18cb62.
//
// Solidity: function confirm(string issuer, string pass, uint256 taskID, string comment, uint8 st) returns()
func (_Task *TaskTransactorSession) Confirm(issuer string, pass string, taskID *big.Int, comment string, st uint8) (*types.Transaction, error) {
	return _Task.Contract.Confirm(&_Task.TransactOpts, issuer, pass, taskID, comment, st)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string pass, string desc, uint256 bonus) returns()
func (_Task *TaskTransactor) Issue(opts *bind.TransactOpts, issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "issue", issuer, pass, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string pass, string desc, uint256 bonus) returns()
func (_Task *TaskSession) Issue(issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, pass, desc, bonus)
}

// Issue is a paid mutator transaction binding the contract method 0x679a5e14.
//
// Solidity: function issue(string issuer, string pass, string desc, uint256 bonus) returns()
func (_Task *TaskTransactorSession) Issue(issuer string, pass string, desc string, bonus *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Issue(&_Task.TransactOpts, issuer, pass, desc, bonus)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns()
func (_Task *TaskTransactor) Mint(opts *bind.TransactOpts, to string, value *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns()
func (_Task *TaskSession) Mint(to string, value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns()
func (_Task *TaskTransactorSession) Mint(to string, value *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Mint(&_Task.TransactOpts, to, value)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string userid, string username, string passwd) returns()
func (_Task *TaskTransactor) Register(opts *bind.TransactOpts, userid string, username string, passwd string) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "register", userid, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string userid, string username, string passwd) returns()
func (_Task *TaskSession) Register(userid string, username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, userid, username, passwd)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string userid, string username, string passwd) returns()
func (_Task *TaskTransactorSession) Register(userid string, username string, passwd string) (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts, userid, username, passwd)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactor) Take(opts *bind.TransactOpts, who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "take", who, pass, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string who, string pass, uint256 taskID) returns()
func (_Task *TaskSession) Take(who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, who, pass, taskID)
}

// Take is a paid mutator transaction binding the contract method 0x5b4c0b54.
//
// Solidity: function take(string who, string pass, uint256 taskID) returns()
func (_Task *TaskTransactorSession) Take(who string, pass string, taskID *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Take(&_Task.TransactOpts, who, pass, taskID)
}
