// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newMember\",\"type\":\"address\"}],\"name\":\"registerMemberApi\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCreditHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_typeState\",\"type\":\"int8\"}],\"name\":\"reCreateStash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"debitIdx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"name\":\"txRef\",\"type\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"bytes32\"},{\"name\":\"receiver\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"int256\"},{\"name\":\"note\",\"type\":\"string\"},{\"name\":\"txType\",\"type\":\"int8\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"credits\",\"outputs\":[{\"name\":\"txRef\",\"type\":\"bytes32\"},{\"name\":\"stashName\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"int256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_stash\",\"type\":\"address\"}],\"name\":\"loadStashRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_txRef\",\"type\":\"bytes32\"},{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"int256\"}],\"name\":\"debit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditIdx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_typeState\",\"type\":\"int8\"}],\"name\":\"createStash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_stashState\",\"type\":\"int8\"}],\"name\":\"setState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDebitHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"debits\",\"outputs\":[{\"name\":\"txRef\",\"type\":\"bytes32\"},{\"name\":\"stashName\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"int256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_txRef\",\"type\":\"bytes32\"},{\"name\":\"_sender\",\"type\":\"bytes32\"},{\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"int256\"},{\"name\":\"_note\",\"type\":\"string\"},{\"name\":\"_txType\",\"type\":\"int8\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"sender_bal\",\"type\":\"int256\"},{\"name\":\"receiver_bal\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_txRef\",\"type\":\"bytes32\"},{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"int256\"}],\"name\":\"credit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"name\":\"string_tobytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferIdx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransferHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegistedAccEthLength\",\"outputs\":[{\"name\":\"\",\"type\":\"int16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_listAcc\",\"type\":\"address[]\"}],\"name\":\"registerAccETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMemberApiIdxLenght\",\"outputs\":[{\"name\":\"\",\"type\":\"int16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerAllStash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStashNamesLenght\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stashRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wallet_type\",\"type\":\"int8\"}],\"name\":\"event_createStash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"old_wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"new_wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wallet_type\",\"type\":\"int8\"}],\"name\":\"event_reCreateStash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"stashState\",\"type\":\"int8\"},{\"indexed\":false,\"name\":\"oldState\",\"type\":\"int256\"}],\"name\":\"event_setState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txRef\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"event_debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txRef\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"event_credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txRef\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"note\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"txType\",\"type\":\"int8\"},{\"indexed\":false,\"name\":\"sender_bal\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"receiver_bal\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"event_transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"listAcc\",\"type\":\"address[]\"}],\"name\":\"event_registerAccETH\",\"type\":\"event\"}]"

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// CreditIdx is a free data retrieval call binding the contract method 0x735677c8.
//
// Solidity: function creditIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCaller) CreditIdx(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "creditIdx", arg0)
	return *ret0, err
}

// CreditIdx is a free data retrieval call binding the contract method 0x735677c8.
//
// Solidity: function creditIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsSession) CreditIdx(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.CreditIdx(&_Contracts.CallOpts, arg0)
}

// CreditIdx is a free data retrieval call binding the contract method 0x735677c8.
//
// Solidity: function creditIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCallerSession) CreditIdx(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.CreditIdx(&_Contracts.CallOpts, arg0)
}

// Credits is a free data retrieval call binding the contract method 0x42997913.
//
// Solidity: function credits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Contracts *ContractsCaller) Credits(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		TxRef     [32]byte
		StashName [32]byte
		Amount    *big.Int
		Timestamp *big.Int
	})
	out := ret
	err := _Contracts.contract.Call(opts, out, "credits", arg0)
	return *ret, err
}

// Credits is a free data retrieval call binding the contract method 0x42997913.
//
// Solidity: function credits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Contracts *ContractsSession) Credits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.Credits(&_Contracts.CallOpts, arg0)
}

// Credits is a free data retrieval call binding the contract method 0x42997913.
//
// Solidity: function credits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Contracts *ContractsCallerSession) Credits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.Credits(&_Contracts.CallOpts, arg0)
}

// DebitIdx is a free data retrieval call binding the contract method 0x203aa89c.
//
// Solidity: function debitIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCaller) DebitIdx(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "debitIdx", arg0)
	return *ret0, err
}

// DebitIdx is a free data retrieval call binding the contract method 0x203aa89c.
//
// Solidity: function debitIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsSession) DebitIdx(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.DebitIdx(&_Contracts.CallOpts, arg0)
}

// DebitIdx is a free data retrieval call binding the contract method 0x203aa89c.
//
// Solidity: function debitIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCallerSession) DebitIdx(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.DebitIdx(&_Contracts.CallOpts, arg0)
}

// Debits is a free data retrieval call binding the contract method 0xa2cffc96.
//
// Solidity: function debits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Contracts *ContractsCaller) Debits(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		TxRef     [32]byte
		StashName [32]byte
		Amount    *big.Int
		Timestamp *big.Int
	})
	out := ret
	err := _Contracts.contract.Call(opts, out, "debits", arg0)
	return *ret, err
}

// Debits is a free data retrieval call binding the contract method 0xa2cffc96.
//
// Solidity: function debits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Contracts *ContractsSession) Debits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.Debits(&_Contracts.CallOpts, arg0)
}

// Debits is a free data retrieval call binding the contract method 0xa2cffc96.
//
// Solidity: function debits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Contracts *ContractsCallerSession) Debits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.Debits(&_Contracts.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _stashName) constant returns(int256)
func (_Contracts *ContractsCaller) GetBalance(opts *bind.CallOpts, _stashName [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getBalance", _stashName)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _stashName) constant returns(int256)
func (_Contracts *ContractsSession) GetBalance(_stashName [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts, _stashName)
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _stashName) constant returns(int256)
func (_Contracts *ContractsCallerSession) GetBalance(_stashName [32]byte) (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts, _stashName)
}

// GetCreditHistoryLength is a free data retrieval call binding the contract method 0x1d74f2e2.
//
// Solidity: function getCreditHistoryLength() constant returns(uint256)
func (_Contracts *ContractsCaller) GetCreditHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getCreditHistoryLength")
	return *ret0, err
}

// GetCreditHistoryLength is a free data retrieval call binding the contract method 0x1d74f2e2.
//
// Solidity: function getCreditHistoryLength() constant returns(uint256)
func (_Contracts *ContractsSession) GetCreditHistoryLength() (*big.Int, error) {
	return _Contracts.Contract.GetCreditHistoryLength(&_Contracts.CallOpts)
}

// GetCreditHistoryLength is a free data retrieval call binding the contract method 0x1d74f2e2.
//
// Solidity: function getCreditHistoryLength() constant returns(uint256)
func (_Contracts *ContractsCallerSession) GetCreditHistoryLength() (*big.Int, error) {
	return _Contracts.Contract.GetCreditHistoryLength(&_Contracts.CallOpts)
}

// GetDebitHistoryLength is a free data retrieval call binding the contract method 0x8d7a194a.
//
// Solidity: function getDebitHistoryLength() constant returns(uint256)
func (_Contracts *ContractsCaller) GetDebitHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getDebitHistoryLength")
	return *ret0, err
}

// GetDebitHistoryLength is a free data retrieval call binding the contract method 0x8d7a194a.
//
// Solidity: function getDebitHistoryLength() constant returns(uint256)
func (_Contracts *ContractsSession) GetDebitHistoryLength() (*big.Int, error) {
	return _Contracts.Contract.GetDebitHistoryLength(&_Contracts.CallOpts)
}

// GetDebitHistoryLength is a free data retrieval call binding the contract method 0x8d7a194a.
//
// Solidity: function getDebitHistoryLength() constant returns(uint256)
func (_Contracts *ContractsCallerSession) GetDebitHistoryLength() (*big.Int, error) {
	return _Contracts.Contract.GetDebitHistoryLength(&_Contracts.CallOpts)
}

// GetMemberApiIdxLenght is a free data retrieval call binding the contract method 0xe45ab733.
//
// Solidity: function getMemberApiIdxLenght() constant returns(int16)
func (_Contracts *ContractsCaller) GetMemberApiIdxLenght(opts *bind.CallOpts) (int16, error) {
	var (
		ret0 = new(int16)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getMemberApiIdxLenght")
	return *ret0, err
}

// GetMemberApiIdxLenght is a free data retrieval call binding the contract method 0xe45ab733.
//
// Solidity: function getMemberApiIdxLenght() constant returns(int16)
func (_Contracts *ContractsSession) GetMemberApiIdxLenght() (int16, error) {
	return _Contracts.Contract.GetMemberApiIdxLenght(&_Contracts.CallOpts)
}

// GetMemberApiIdxLenght is a free data retrieval call binding the contract method 0xe45ab733.
//
// Solidity: function getMemberApiIdxLenght() constant returns(int16)
func (_Contracts *ContractsCallerSession) GetMemberApiIdxLenght() (int16, error) {
	return _Contracts.Contract.GetMemberApiIdxLenght(&_Contracts.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Contracts *ContractsCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Contracts *ContractsSession) GetOwner() (common.Address, error) {
	return _Contracts.Contract.GetOwner(&_Contracts.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Contracts *ContractsCallerSession) GetOwner() (common.Address, error) {
	return _Contracts.Contract.GetOwner(&_Contracts.CallOpts)
}

// GetRegistedAccEthLength is a free data retrieval call binding the contract method 0xdf646612.
//
// Solidity: function getRegistedAccEthLength() constant returns(int16)
func (_Contracts *ContractsCaller) GetRegistedAccEthLength(opts *bind.CallOpts) (int16, error) {
	var (
		ret0 = new(int16)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getRegistedAccEthLength")
	return *ret0, err
}

// GetRegistedAccEthLength is a free data retrieval call binding the contract method 0xdf646612.
//
// Solidity: function getRegistedAccEthLength() constant returns(int16)
func (_Contracts *ContractsSession) GetRegistedAccEthLength() (int16, error) {
	return _Contracts.Contract.GetRegistedAccEthLength(&_Contracts.CallOpts)
}

// GetRegistedAccEthLength is a free data retrieval call binding the contract method 0xdf646612.
//
// Solidity: function getRegistedAccEthLength() constant returns(int16)
func (_Contracts *ContractsCallerSession) GetRegistedAccEthLength() (int16, error) {
	return _Contracts.Contract.GetRegistedAccEthLength(&_Contracts.CallOpts)
}

// GetStashNamesLenght is a free data retrieval call binding the contract method 0xee9a5812.
//
// Solidity: function getStashNamesLenght() constant returns(int256)
func (_Contracts *ContractsCaller) GetStashNamesLenght(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getStashNamesLenght")
	return *ret0, err
}

// GetStashNamesLenght is a free data retrieval call binding the contract method 0xee9a5812.
//
// Solidity: function getStashNamesLenght() constant returns(int256)
func (_Contracts *ContractsSession) GetStashNamesLenght() (*big.Int, error) {
	return _Contracts.Contract.GetStashNamesLenght(&_Contracts.CallOpts)
}

// GetStashNamesLenght is a free data retrieval call binding the contract method 0xee9a5812.
//
// Solidity: function getStashNamesLenght() constant returns(int256)
func (_Contracts *ContractsCallerSession) GetStashNamesLenght() (*big.Int, error) {
	return _Contracts.Contract.GetStashNamesLenght(&_Contracts.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 _stashName) constant returns(int8)
func (_Contracts *ContractsCaller) GetState(opts *bind.CallOpts, _stashName [32]byte) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getState", _stashName)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 _stashName) constant returns(int8)
func (_Contracts *ContractsSession) GetState(_stashName [32]byte) (int8, error) {
	return _Contracts.Contract.GetState(&_Contracts.CallOpts, _stashName)
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 _stashName) constant returns(int8)
func (_Contracts *ContractsCallerSession) GetState(_stashName [32]byte) (int8, error) {
	return _Contracts.Contract.GetState(&_Contracts.CallOpts, _stashName)
}

// GetTransferHistoryLength is a free data retrieval call binding the contract method 0xd236e1ba.
//
// Solidity: function getTransferHistoryLength() constant returns(uint256)
func (_Contracts *ContractsCaller) GetTransferHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getTransferHistoryLength")
	return *ret0, err
}

// GetTransferHistoryLength is a free data retrieval call binding the contract method 0xd236e1ba.
//
// Solidity: function getTransferHistoryLength() constant returns(uint256)
func (_Contracts *ContractsSession) GetTransferHistoryLength() (*big.Int, error) {
	return _Contracts.Contract.GetTransferHistoryLength(&_Contracts.CallOpts)
}

// GetTransferHistoryLength is a free data retrieval call binding the contract method 0xd236e1ba.
//
// Solidity: function getTransferHistoryLength() constant returns(uint256)
func (_Contracts *ContractsCallerSession) GetTransferHistoryLength() (*big.Int, error) {
	return _Contracts.Contract.GetTransferHistoryLength(&_Contracts.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _stashName) constant returns(int8)
func (_Contracts *ContractsCaller) GetType(opts *bind.CallOpts, _stashName [32]byte) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getType", _stashName)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _stashName) constant returns(int8)
func (_Contracts *ContractsSession) GetType(_stashName [32]byte) (int8, error) {
	return _Contracts.Contract.GetType(&_Contracts.CallOpts, _stashName)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _stashName) constant returns(int8)
func (_Contracts *ContractsCallerSession) GetType(_stashName [32]byte) (int8, error) {
	return _Contracts.Contract.GetType(&_Contracts.CallOpts, _stashName)
}

// StashNames is a free data retrieval call binding the contract method 0x0ae94616.
//
// Solidity: function stashNames(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCaller) StashNames(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "stashNames", arg0)
	return *ret0, err
}

// StashNames is a free data retrieval call binding the contract method 0x0ae94616.
//
// Solidity: function stashNames(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsSession) StashNames(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.StashNames(&_Contracts.CallOpts, arg0)
}

// StashNames is a free data retrieval call binding the contract method 0x0ae94616.
//
// Solidity: function stashNames(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCallerSession) StashNames(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.StashNames(&_Contracts.CallOpts, arg0)
}

// StashRegistry is a free data retrieval call binding the contract method 0xfb6cb61f.
//
// Solidity: function stashRegistry(bytes32 ) constant returns(address)
func (_Contracts *ContractsCaller) StashRegistry(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "stashRegistry", arg0)
	return *ret0, err
}

// StashRegistry is a free data retrieval call binding the contract method 0xfb6cb61f.
//
// Solidity: function stashRegistry(bytes32 ) constant returns(address)
func (_Contracts *ContractsSession) StashRegistry(arg0 [32]byte) (common.Address, error) {
	return _Contracts.Contract.StashRegistry(&_Contracts.CallOpts, arg0)
}

// StashRegistry is a free data retrieval call binding the contract method 0xfb6cb61f.
//
// Solidity: function stashRegistry(bytes32 ) constant returns(address)
func (_Contracts *ContractsCallerSession) StashRegistry(arg0 [32]byte) (common.Address, error) {
	return _Contracts.Contract.StashRegistry(&_Contracts.CallOpts, arg0)
}

// StringTobytes is a free data retrieval call binding the contract method 0xbc6d0577.
//
// Solidity: function string_tobytes(string s) constant returns(bytes)
func (_Contracts *ContractsCaller) StringTobytes(opts *bind.CallOpts, s string) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "string_tobytes", s)
	return *ret0, err
}

// StringTobytes is a free data retrieval call binding the contract method 0xbc6d0577.
//
// Solidity: function string_tobytes(string s) constant returns(bytes)
func (_Contracts *ContractsSession) StringTobytes(s string) ([]byte, error) {
	return _Contracts.Contract.StringTobytes(&_Contracts.CallOpts, s)
}

// StringTobytes is a free data retrieval call binding the contract method 0xbc6d0577.
//
// Solidity: function string_tobytes(string s) constant returns(bytes)
func (_Contracts *ContractsCallerSession) StringTobytes(s string) ([]byte, error) {
	return _Contracts.Contract.StringTobytes(&_Contracts.CallOpts, s)
}

// TransferIdx is a free data retrieval call binding the contract method 0xc83fe142.
//
// Solidity: function transferIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCaller) TransferIdx(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "transferIdx", arg0)
	return *ret0, err
}

// TransferIdx is a free data retrieval call binding the contract method 0xc83fe142.
//
// Solidity: function transferIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsSession) TransferIdx(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.TransferIdx(&_Contracts.CallOpts, arg0)
}

// TransferIdx is a free data retrieval call binding the contract method 0xc83fe142.
//
// Solidity: function transferIdx(uint256 ) constant returns(bytes32)
func (_Contracts *ContractsCallerSession) TransferIdx(arg0 *big.Int) ([32]byte, error) {
	return _Contracts.Contract.TransferIdx(&_Contracts.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 txRef, bytes32 sender, bytes32 receiver, int256 amount, string note, int8 txType, uint256 timestamp)
func (_Contracts *ContractsCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TxRef     [32]byte
	Sender    [32]byte
	Receiver  [32]byte
	Amount    *big.Int
	Note      string
	TxType    int8
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		TxRef     [32]byte
		Sender    [32]byte
		Receiver  [32]byte
		Amount    *big.Int
		Note      string
		TxType    int8
		Timestamp *big.Int
	})
	out := ret
	err := _Contracts.contract.Call(opts, out, "transfers", arg0)
	return *ret, err
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 txRef, bytes32 sender, bytes32 receiver, int256 amount, string note, int8 txType, uint256 timestamp)
func (_Contracts *ContractsSession) Transfers(arg0 [32]byte) (struct {
	TxRef     [32]byte
	Sender    [32]byte
	Receiver  [32]byte
	Amount    *big.Int
	Note      string
	TxType    int8
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.Transfers(&_Contracts.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 txRef, bytes32 sender, bytes32 receiver, int256 amount, string note, int8 txType, uint256 timestamp)
func (_Contracts *ContractsCallerSession) Transfers(arg0 [32]byte) (struct {
	TxRef     [32]byte
	Sender    [32]byte
	Receiver  [32]byte
	Amount    *big.Int
	Note      string
	TxType    int8
	Timestamp *big.Int
}, error) {
	return _Contracts.Contract.Transfers(&_Contracts.CallOpts, arg0)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Contracts *ContractsTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Contracts *ContractsSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOwner(&_Contracts.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Contracts *ContractsTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOwner(&_Contracts.TransactOpts, _newOwner)
}

// ChangeOwnerAllStash is a paid mutator transaction binding the contract method 0xeae5180e.
//
// Solidity: function changeOwnerAllStash(address _newOwner) returns()
func (_Contracts *ContractsTransactor) ChangeOwnerAllStash(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeOwnerAllStash", _newOwner)
}

// ChangeOwnerAllStash is a paid mutator transaction binding the contract method 0xeae5180e.
//
// Solidity: function changeOwnerAllStash(address _newOwner) returns()
func (_Contracts *ContractsSession) ChangeOwnerAllStash(_newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOwnerAllStash(&_Contracts.TransactOpts, _newOwner)
}

// ChangeOwnerAllStash is a paid mutator transaction binding the contract method 0xeae5180e.
//
// Solidity: function changeOwnerAllStash(address _newOwner) returns()
func (_Contracts *ContractsTransactorSession) ChangeOwnerAllStash(_newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeOwnerAllStash(&_Contracts.TransactOpts, _newOwner)
}

// CreateStash is a paid mutator transaction binding the contract method 0x7e53308a.
//
// Solidity: function createStash(bytes32 _stashName, int8 _typeState) returns()
func (_Contracts *ContractsTransactor) CreateStash(opts *bind.TransactOpts, _stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createStash", _stashName, _typeState)
}

// CreateStash is a paid mutator transaction binding the contract method 0x7e53308a.
//
// Solidity: function createStash(bytes32 _stashName, int8 _typeState) returns()
func (_Contracts *ContractsSession) CreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Contracts.Contract.CreateStash(&_Contracts.TransactOpts, _stashName, _typeState)
}

// CreateStash is a paid mutator transaction binding the contract method 0x7e53308a.
//
// Solidity: function createStash(bytes32 _stashName, int8 _typeState) returns()
func (_Contracts *ContractsTransactorSession) CreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Contracts.Contract.CreateStash(&_Contracts.TransactOpts, _stashName, _typeState)
}

// Credit is a paid mutator transaction binding the contract method 0xb91f176e.
//
// Solidity: function credit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Contracts *ContractsTransactor) Credit(opts *bind.TransactOpts, _txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "credit", _txRef, _stashName, _amount)
}

// Credit is a paid mutator transaction binding the contract method 0xb91f176e.
//
// Solidity: function credit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Contracts *ContractsSession) Credit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Credit(&_Contracts.TransactOpts, _txRef, _stashName, _amount)
}

// Credit is a paid mutator transaction binding the contract method 0xb91f176e.
//
// Solidity: function credit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Contracts *ContractsTransactorSession) Credit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Credit(&_Contracts.TransactOpts, _txRef, _stashName, _amount)
}

// Debit is a paid mutator transaction binding the contract method 0x4ea679e7.
//
// Solidity: function debit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Contracts *ContractsTransactor) Debit(opts *bind.TransactOpts, _txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "debit", _txRef, _stashName, _amount)
}

// Debit is a paid mutator transaction binding the contract method 0x4ea679e7.
//
// Solidity: function debit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Contracts *ContractsSession) Debit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Debit(&_Contracts.TransactOpts, _txRef, _stashName, _amount)
}

// Debit is a paid mutator transaction binding the contract method 0x4ea679e7.
//
// Solidity: function debit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Contracts *ContractsTransactorSession) Debit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Debit(&_Contracts.TransactOpts, _txRef, _stashName, _amount)
}

// LoadStashRegistry is a paid mutator transaction binding the contract method 0x4e3d7de0.
//
// Solidity: function loadStashRegistry(bytes32 _stashName, address _stash) returns()
func (_Contracts *ContractsTransactor) LoadStashRegistry(opts *bind.TransactOpts, _stashName [32]byte, _stash common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "loadStashRegistry", _stashName, _stash)
}

// LoadStashRegistry is a paid mutator transaction binding the contract method 0x4e3d7de0.
//
// Solidity: function loadStashRegistry(bytes32 _stashName, address _stash) returns()
func (_Contracts *ContractsSession) LoadStashRegistry(_stashName [32]byte, _stash common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.LoadStashRegistry(&_Contracts.TransactOpts, _stashName, _stash)
}

// LoadStashRegistry is a paid mutator transaction binding the contract method 0x4e3d7de0.
//
// Solidity: function loadStashRegistry(bytes32 _stashName, address _stash) returns()
func (_Contracts *ContractsTransactorSession) LoadStashRegistry(_stashName [32]byte, _stash common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.LoadStashRegistry(&_Contracts.TransactOpts, _stashName, _stash)
}

// ReCreateStash is a paid mutator transaction binding the contract method 0x1e022e75.
//
// Solidity: function reCreateStash(bytes32 _stashName, int8 _typeState) returns()
func (_Contracts *ContractsTransactor) ReCreateStash(opts *bind.TransactOpts, _stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "reCreateStash", _stashName, _typeState)
}

// ReCreateStash is a paid mutator transaction binding the contract method 0x1e022e75.
//
// Solidity: function reCreateStash(bytes32 _stashName, int8 _typeState) returns()
func (_Contracts *ContractsSession) ReCreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Contracts.Contract.ReCreateStash(&_Contracts.TransactOpts, _stashName, _typeState)
}

// ReCreateStash is a paid mutator transaction binding the contract method 0x1e022e75.
//
// Solidity: function reCreateStash(bytes32 _stashName, int8 _typeState) returns()
func (_Contracts *ContractsTransactorSession) ReCreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Contracts.Contract.ReCreateStash(&_Contracts.TransactOpts, _stashName, _typeState)
}

// RegisterAccETH is a paid mutator transaction binding the contract method 0xe1b8ff83.
//
// Solidity: function registerAccETH(address[] _listAcc) returns()
func (_Contracts *ContractsTransactor) RegisterAccETH(opts *bind.TransactOpts, _listAcc []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerAccETH", _listAcc)
}

// RegisterAccETH is a paid mutator transaction binding the contract method 0xe1b8ff83.
//
// Solidity: function registerAccETH(address[] _listAcc) returns()
func (_Contracts *ContractsSession) RegisterAccETH(_listAcc []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterAccETH(&_Contracts.TransactOpts, _listAcc)
}

// RegisterAccETH is a paid mutator transaction binding the contract method 0xe1b8ff83.
//
// Solidity: function registerAccETH(address[] _listAcc) returns()
func (_Contracts *ContractsTransactorSession) RegisterAccETH(_listAcc []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterAccETH(&_Contracts.TransactOpts, _listAcc)
}

// RegisterMemberApi is a paid mutator transaction binding the contract method 0x0d3b2b6f.
//
// Solidity: function registerMemberApi(address _newMember) returns()
func (_Contracts *ContractsTransactor) RegisterMemberApi(opts *bind.TransactOpts, _newMember common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerMemberApi", _newMember)
}

// RegisterMemberApi is a paid mutator transaction binding the contract method 0x0d3b2b6f.
//
// Solidity: function registerMemberApi(address _newMember) returns()
func (_Contracts *ContractsSession) RegisterMemberApi(_newMember common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterMemberApi(&_Contracts.TransactOpts, _newMember)
}

// RegisterMemberApi is a paid mutator transaction binding the contract method 0x0d3b2b6f.
//
// Solidity: function registerMemberApi(address _newMember) returns()
func (_Contracts *ContractsTransactorSession) RegisterMemberApi(_newMember common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterMemberApi(&_Contracts.TransactOpts, _newMember)
}

// SetState is a paid mutator transaction binding the contract method 0x8b356a4d.
//
// Solidity: function setState(bytes32 _stashName, int8 _stashState) returns()
func (_Contracts *ContractsTransactor) SetState(opts *bind.TransactOpts, _stashName [32]byte, _stashState int8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setState", _stashName, _stashState)
}

// SetState is a paid mutator transaction binding the contract method 0x8b356a4d.
//
// Solidity: function setState(bytes32 _stashName, int8 _stashState) returns()
func (_Contracts *ContractsSession) SetState(_stashName [32]byte, _stashState int8) (*types.Transaction, error) {
	return _Contracts.Contract.SetState(&_Contracts.TransactOpts, _stashName, _stashState)
}

// SetState is a paid mutator transaction binding the contract method 0x8b356a4d.
//
// Solidity: function setState(bytes32 _stashName, int8 _stashState) returns()
func (_Contracts *ContractsTransactorSession) SetState(_stashName [32]byte, _stashState int8) (*types.Transaction, error) {
	return _Contracts.Contract.SetState(&_Contracts.TransactOpts, _stashName, _stashState)
}

// Transfer is a paid mutator transaction binding the contract method 0xa4642866.
//
// Solidity: function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int256 _amount, string _note, int8 _txType) returns(int256 sender_bal, int256 receiver_bal)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, _txRef [32]byte, _sender [32]byte, _receiver [32]byte, _amount *big.Int, _note string, _txType int8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", _txRef, _sender, _receiver, _amount, _note, _txType)
}

// Transfer is a paid mutator transaction binding the contract method 0xa4642866.
//
// Solidity: function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int256 _amount, string _note, int8 _txType) returns(int256 sender_bal, int256 receiver_bal)
func (_Contracts *ContractsSession) Transfer(_txRef [32]byte, _sender [32]byte, _receiver [32]byte, _amount *big.Int, _note string, _txType int8) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, _txRef, _sender, _receiver, _amount, _note, _txType)
}

// Transfer is a paid mutator transaction binding the contract method 0xa4642866.
//
// Solidity: function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int256 _amount, string _note, int8 _txType) returns(int256 sender_bal, int256 receiver_bal)
func (_Contracts *ContractsTransactorSession) Transfer(_txRef [32]byte, _sender [32]byte, _receiver [32]byte, _amount *big.Int, _note string, _txType int8) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, _txRef, _sender, _receiver, _amount, _note, _txType)
}

// ContractsEventCreateStashIterator is returned from FilterEventCreateStash and is used to iterate over the raw logs and unpacked data for EventCreateStash events raised by the Contracts contract.
type ContractsEventCreateStashIterator struct {
	Event *ContractsEventCreateStash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEventCreateStashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEventCreateStash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEventCreateStash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEventCreateStashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEventCreateStashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEventCreateStash represents a EventCreateStash event raised by the Contracts contract.
type ContractsEventCreateStash struct {
	WalletCode    [32]byte
	WalletAddress common.Address
	WalletType    int8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventCreateStash is a free log retrieval operation binding the contract event 0x0cb60d28f723451a1bd65a935f5e1d7b4e708818cf6cc6c30ad742fd67531b00.
//
// Solidity: event event_createStash(bytes32 indexed wallet_code, address wallet_address, int8 wallet_type)
func (_Contracts *ContractsFilterer) FilterEventCreateStash(opts *bind.FilterOpts, wallet_code [][32]byte) (*ContractsEventCreateStashIterator, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "event_createStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsEventCreateStashIterator{contract: _Contracts.contract, event: "event_createStash", logs: logs, sub: sub}, nil
}

// WatchEventCreateStash is a free log subscription operation binding the contract event 0x0cb60d28f723451a1bd65a935f5e1d7b4e708818cf6cc6c30ad742fd67531b00.
//
// Solidity: event event_createStash(bytes32 indexed wallet_code, address wallet_address, int8 wallet_type)
func (_Contracts *ContractsFilterer) WatchEventCreateStash(opts *bind.WatchOpts, sink chan<- *ContractsEventCreateStash, wallet_code [][32]byte) (event.Subscription, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "event_createStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEventCreateStash)
				if err := _Contracts.contract.UnpackLog(event, "event_createStash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ContractsEventCreditIterator is returned from FilterEventCredit and is used to iterate over the raw logs and unpacked data for EventCredit events raised by the Contracts contract.
type ContractsEventCreditIterator struct {
	Event *ContractsEventCredit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEventCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEventCredit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEventCredit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEventCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEventCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEventCredit represents a EventCredit event raised by the Contracts contract.
type ContractsEventCredit struct {
	TxRef      [32]byte
	WalletCode [32]byte
	Amount     *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventCredit is a free log retrieval operation binding the contract event 0x2bda1f3aff897895cbea7f906b2e41ef99d39163a7b36af5782b8b718ed7e456.
//
// Solidity: event event_credit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Contracts *ContractsFilterer) FilterEventCredit(opts *bind.FilterOpts, txRef [][32]byte, wallet_code [][32]byte) (*ContractsEventCreditIterator, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "event_credit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsEventCreditIterator{contract: _Contracts.contract, event: "event_credit", logs: logs, sub: sub}, nil
}

// WatchEventCredit is a free log subscription operation binding the contract event 0x2bda1f3aff897895cbea7f906b2e41ef99d39163a7b36af5782b8b718ed7e456.
//
// Solidity: event event_credit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Contracts *ContractsFilterer) WatchEventCredit(opts *bind.WatchOpts, sink chan<- *ContractsEventCredit, txRef [][32]byte, wallet_code [][32]byte) (event.Subscription, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "event_credit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEventCredit)
				if err := _Contracts.contract.UnpackLog(event, "event_credit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ContractsEventDebitIterator is returned from FilterEventDebit and is used to iterate over the raw logs and unpacked data for EventDebit events raised by the Contracts contract.
type ContractsEventDebitIterator struct {
	Event *ContractsEventDebit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEventDebitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEventDebit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEventDebit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEventDebitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEventDebitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEventDebit represents a EventDebit event raised by the Contracts contract.
type ContractsEventDebit struct {
	TxRef      [32]byte
	WalletCode [32]byte
	Amount     *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventDebit is a free log retrieval operation binding the contract event 0x3e888fb580f7ba504d1d06b6af38e430c683d120af553208a48629dbd58bb0f8.
//
// Solidity: event event_debit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Contracts *ContractsFilterer) FilterEventDebit(opts *bind.FilterOpts, txRef [][32]byte, wallet_code [][32]byte) (*ContractsEventDebitIterator, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "event_debit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsEventDebitIterator{contract: _Contracts.contract, event: "event_debit", logs: logs, sub: sub}, nil
}

// WatchEventDebit is a free log subscription operation binding the contract event 0x3e888fb580f7ba504d1d06b6af38e430c683d120af553208a48629dbd58bb0f8.
//
// Solidity: event event_debit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Contracts *ContractsFilterer) WatchEventDebit(opts *bind.WatchOpts, sink chan<- *ContractsEventDebit, txRef [][32]byte, wallet_code [][32]byte) (event.Subscription, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "event_debit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEventDebit)
				if err := _Contracts.contract.UnpackLog(event, "event_debit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ContractsEventReCreateStashIterator is returned from FilterEventReCreateStash and is used to iterate over the raw logs and unpacked data for EventReCreateStash events raised by the Contracts contract.
type ContractsEventReCreateStashIterator struct {
	Event *ContractsEventReCreateStash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEventReCreateStashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEventReCreateStash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEventReCreateStash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEventReCreateStashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEventReCreateStashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEventReCreateStash represents a EventReCreateStash event raised by the Contracts contract.
type ContractsEventReCreateStash struct {
	WalletCode       [32]byte
	OldWalletAddress common.Address
	NewWalletAddress common.Address
	WalletType       int8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEventReCreateStash is a free log retrieval operation binding the contract event 0x8c8e3419f9e71366ada9d23ae8a8dba5229a20e91deb3cab00f419ea5b0480e8.
//
// Solidity: event event_reCreateStash(bytes32 indexed wallet_code, address old_wallet_address, address new_wallet_address, int8 wallet_type)
func (_Contracts *ContractsFilterer) FilterEventReCreateStash(opts *bind.FilterOpts, wallet_code [][32]byte) (*ContractsEventReCreateStashIterator, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "event_reCreateStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsEventReCreateStashIterator{contract: _Contracts.contract, event: "event_reCreateStash", logs: logs, sub: sub}, nil
}

// WatchEventReCreateStash is a free log subscription operation binding the contract event 0x8c8e3419f9e71366ada9d23ae8a8dba5229a20e91deb3cab00f419ea5b0480e8.
//
// Solidity: event event_reCreateStash(bytes32 indexed wallet_code, address old_wallet_address, address new_wallet_address, int8 wallet_type)
func (_Contracts *ContractsFilterer) WatchEventReCreateStash(opts *bind.WatchOpts, sink chan<- *ContractsEventReCreateStash, wallet_code [][32]byte) (event.Subscription, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "event_reCreateStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEventReCreateStash)
				if err := _Contracts.contract.UnpackLog(event, "event_reCreateStash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ContractsEventRegisterAccETHIterator is returned from FilterEventRegisterAccETH and is used to iterate over the raw logs and unpacked data for EventRegisterAccETH events raised by the Contracts contract.
type ContractsEventRegisterAccETHIterator struct {
	Event *ContractsEventRegisterAccETH // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEventRegisterAccETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEventRegisterAccETH)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEventRegisterAccETH)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEventRegisterAccETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEventRegisterAccETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEventRegisterAccETH represents a EventRegisterAccETH event raised by the Contracts contract.
type ContractsEventRegisterAccETH struct {
	ListAcc []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventRegisterAccETH is a free log retrieval operation binding the contract event 0x3b0686a3c64982476da002e36939997a11d9ef748420920b111c3ba5d2e9784f.
//
// Solidity: event event_registerAccETH(address[] listAcc)
func (_Contracts *ContractsFilterer) FilterEventRegisterAccETH(opts *bind.FilterOpts) (*ContractsEventRegisterAccETHIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "event_registerAccETH")
	if err != nil {
		return nil, err
	}
	return &ContractsEventRegisterAccETHIterator{contract: _Contracts.contract, event: "event_registerAccETH", logs: logs, sub: sub}, nil
}

// WatchEventRegisterAccETH is a free log subscription operation binding the contract event 0x3b0686a3c64982476da002e36939997a11d9ef748420920b111c3ba5d2e9784f.
//
// Solidity: event event_registerAccETH(address[] listAcc)
func (_Contracts *ContractsFilterer) WatchEventRegisterAccETH(opts *bind.WatchOpts, sink chan<- *ContractsEventRegisterAccETH) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "event_registerAccETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEventRegisterAccETH)
				if err := _Contracts.contract.UnpackLog(event, "event_registerAccETH", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ContractsEventSetStateIterator is returned from FilterEventSetState and is used to iterate over the raw logs and unpacked data for EventSetState events raised by the Contracts contract.
type ContractsEventSetStateIterator struct {
	Event *ContractsEventSetState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEventSetStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEventSetState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEventSetState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEventSetStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEventSetStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEventSetState represents a EventSetState event raised by the Contracts contract.
type ContractsEventSetState struct {
	WalletCode [32]byte
	StashState int8
	OldState   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventSetState is a free log retrieval operation binding the contract event 0x787d29ad64e9e02e1fb29fb54868408884d3969457c61da4943ec8d57e8ac011.
//
// Solidity: event event_setState(bytes32 indexed wallet_code, int8 stashState, int256 oldState)
func (_Contracts *ContractsFilterer) FilterEventSetState(opts *bind.FilterOpts, wallet_code [][32]byte) (*ContractsEventSetStateIterator, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "event_setState", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsEventSetStateIterator{contract: _Contracts.contract, event: "event_setState", logs: logs, sub: sub}, nil
}

// WatchEventSetState is a free log subscription operation binding the contract event 0x787d29ad64e9e02e1fb29fb54868408884d3969457c61da4943ec8d57e8ac011.
//
// Solidity: event event_setState(bytes32 indexed wallet_code, int8 stashState, int256 oldState)
func (_Contracts *ContractsFilterer) WatchEventSetState(opts *bind.WatchOpts, sink chan<- *ContractsEventSetState, wallet_code [][32]byte) (event.Subscription, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "event_setState", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEventSetState)
				if err := _Contracts.contract.UnpackLog(event, "event_setState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ContractsEventTransferIterator is returned from FilterEventTransfer and is used to iterate over the raw logs and unpacked data for EventTransfer events raised by the Contracts contract.
type ContractsEventTransferIterator struct {
	Event *ContractsEventTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsEventTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsEventTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsEventTransfer represents a EventTransfer event raised by the Contracts contract.
type ContractsEventTransfer struct {
	TxRef       [32]byte
	Sender      [32]byte
	Receiver    [32]byte
	Amount      *big.Int
	Note        string
	TxType      int8
	SenderBal   *big.Int
	ReceiverBal *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventTransfer is a free log retrieval operation binding the contract event 0x20244c6c280a8d1e1be518a3fc3acb8c46f303ad98dd9ec4c01baf9e2999f82e.
//
// Solidity: event event_transfer(bytes32 indexed txRef, bytes32 indexed sender, bytes32 indexed receiver, int256 amount, string note, int8 txType, int256 sender_bal, int256 receiver_bal, uint256 timestamp)
func (_Contracts *ContractsFilterer) FilterEventTransfer(opts *bind.FilterOpts, txRef [][32]byte, sender [][32]byte, receiver [][32]byte) (*ContractsEventTransferIterator, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "event_transfer", txRefRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractsEventTransferIterator{contract: _Contracts.contract, event: "event_transfer", logs: logs, sub: sub}, nil
}

// WatchEventTransfer is a free log subscription operation binding the contract event 0x20244c6c280a8d1e1be518a3fc3acb8c46f303ad98dd9ec4c01baf9e2999f82e.
//
// Solidity: event event_transfer(bytes32 indexed txRef, bytes32 indexed sender, bytes32 indexed receiver, int256 amount, string note, int8 txType, int256 sender_bal, int256 receiver_bal, uint256 timestamp)
func (_Contracts *ContractsFilterer) WatchEventTransfer(opts *bind.WatchOpts, sink chan<- *ContractsEventTransfer, txRef [][32]byte, sender [][32]byte, receiver [][32]byte) (event.Subscription, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "event_transfer", txRefRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsEventTransfer)
				if err := _Contracts.contract.UnpackLog(event, "event_transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
