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

// BusinessABI is the input ABI used to generate the binding from.
const BusinessABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"}],\"name\":\"getState\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stashNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newMember\",\"type\":\"address\"}],\"name\":\"registerMemberApi\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCreditHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_typeState\",\"type\":\"int8\"}],\"name\":\"reCreateStash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"debitIdx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"name\":\"txRef\",\"type\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"bytes32\"},{\"name\":\"receiver\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"int256\"},{\"name\":\"note\",\"type\":\"string\"},{\"name\":\"txType\",\"type\":\"int8\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"credits\",\"outputs\":[{\"name\":\"txRef\",\"type\":\"bytes32\"},{\"name\":\"stashName\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"int256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_stash\",\"type\":\"address\"}],\"name\":\"loadStashRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_txRef\",\"type\":\"bytes32\"},{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"int256\"}],\"name\":\"debit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditIdx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_typeState\",\"type\":\"int8\"}],\"name\":\"createStash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_stashState\",\"type\":\"int8\"}],\"name\":\"setState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDebitHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_stashName\",\"type\":\"bytes32\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"debits\",\"outputs\":[{\"name\":\"txRef\",\"type\":\"bytes32\"},{\"name\":\"stashName\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"int256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_txRef\",\"type\":\"bytes32\"},{\"name\":\"_sender\",\"type\":\"bytes32\"},{\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"int256\"},{\"name\":\"_note\",\"type\":\"string\"},{\"name\":\"_txType\",\"type\":\"int8\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"sender_bal\",\"type\":\"int256\"},{\"name\":\"receiver_bal\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_txRef\",\"type\":\"bytes32\"},{\"name\":\"_stashName\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"int256\"}],\"name\":\"credit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"name\":\"string_tobytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferIdx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTransferHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegistedAccEthLength\",\"outputs\":[{\"name\":\"\",\"type\":\"int16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_listAcc\",\"type\":\"address[]\"}],\"name\":\"registerAccETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMemberApiIdxLenght\",\"outputs\":[{\"name\":\"\",\"type\":\"int16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerAllStash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStashNamesLenght\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"stashRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wallet_type\",\"type\":\"int8\"}],\"name\":\"event_createStash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"old_wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"new_wallet_address\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wallet_type\",\"type\":\"int8\"}],\"name\":\"event_reCreateStash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"stashState\",\"type\":\"int8\"},{\"indexed\":false,\"name\":\"oldState\",\"type\":\"int256\"}],\"name\":\"event_setState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txRef\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"event_debit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txRef\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"wallet_code\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"event_credit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txRef\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"note\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"txType\",\"type\":\"int8\"},{\"indexed\":false,\"name\":\"sender_bal\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"receiver_bal\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"event_transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"listAcc\",\"type\":\"address[]\"}],\"name\":\"event_registerAccETH\",\"type\":\"event\"}]"

// Business is an auto generated Go binding around an Ethereum contract.
type Business struct {
	BusinessCaller     // Read-only binding to the contract
	BusinessTransactor // Write-only binding to the contract
	BusinessFilterer   // Log filterer for contract events
}

// BusinessCaller is an auto generated read-only Go binding around an Ethereum contract.
type BusinessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BusinessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BusinessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BusinessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BusinessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BusinessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BusinessSession struct {
	Contract     *Business         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BusinessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BusinessCallerSession struct {
	Contract *BusinessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BusinessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BusinessTransactorSession struct {
	Contract     *BusinessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BusinessRaw is an auto generated low-level Go binding around an Ethereum contract.
type BusinessRaw struct {
	Contract *Business // Generic contract binding to access the raw methods on
}

// BusinessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BusinessCallerRaw struct {
	Contract *BusinessCaller // Generic read-only contract binding to access the raw methods on
}

// BusinessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BusinessTransactorRaw struct {
	Contract *BusinessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBusiness creates a new instance of Business, bound to a specific deployed contract.
func NewBusiness(address common.Address, backend bind.ContractBackend) (*Business, error) {
	contract, err := bindBusiness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Business{BusinessCaller: BusinessCaller{contract: contract}, BusinessTransactor: BusinessTransactor{contract: contract}, BusinessFilterer: BusinessFilterer{contract: contract}}, nil
}

// NewBusinessCaller creates a new read-only instance of Business, bound to a specific deployed contract.
func NewBusinessCaller(address common.Address, caller bind.ContractCaller) (*BusinessCaller, error) {
	contract, err := bindBusiness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BusinessCaller{contract: contract}, nil
}

// NewBusinessTransactor creates a new write-only instance of Business, bound to a specific deployed contract.
func NewBusinessTransactor(address common.Address, transactor bind.ContractTransactor) (*BusinessTransactor, error) {
	contract, err := bindBusiness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BusinessTransactor{contract: contract}, nil
}

// NewBusinessFilterer creates a new log filterer instance of Business, bound to a specific deployed contract.
func NewBusinessFilterer(address common.Address, filterer bind.ContractFilterer) (*BusinessFilterer, error) {
	contract, err := bindBusiness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BusinessFilterer{contract: contract}, nil
}

// bindBusiness binds a generic wrapper to an already deployed contract.
func bindBusiness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BusinessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Business *BusinessRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Business.Contract.BusinessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Business *BusinessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Business.Contract.BusinessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Business *BusinessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Business.Contract.BusinessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Business *BusinessCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Business.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Business *BusinessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Business.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Business *BusinessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Business.Contract.contract.Transact(opts, method, params...)
}

// CreditIdx is a free data retrieval call binding the contract method 0x735677c8.
//
// Solidity: function creditIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessCaller) CreditIdx(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "creditIdx", arg0)
	return *ret0, err
}

// CreditIdx is a free data retrieval call binding the contract method 0x735677c8.
//
// Solidity: function creditIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessSession) CreditIdx(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.CreditIdx(&_Business.CallOpts, arg0)
}

// CreditIdx is a free data retrieval call binding the contract method 0x735677c8.
//
// Solidity: function creditIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessCallerSession) CreditIdx(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.CreditIdx(&_Business.CallOpts, arg0)
}

// Credits is a free data retrieval call binding the contract method 0x42997913.
//
// Solidity: function credits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Business *BusinessCaller) Credits(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _Business.contract.Call(opts, out, "credits", arg0)
	return *ret, err
}

// Credits is a free data retrieval call binding the contract method 0x42997913.
//
// Solidity: function credits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Business *BusinessSession) Credits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Business.Contract.Credits(&_Business.CallOpts, arg0)
}

// Credits is a free data retrieval call binding the contract method 0x42997913.
//
// Solidity: function credits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Business *BusinessCallerSession) Credits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Business.Contract.Credits(&_Business.CallOpts, arg0)
}

// DebitIdx is a free data retrieval call binding the contract method 0x203aa89c.
//
// Solidity: function debitIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessCaller) DebitIdx(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "debitIdx", arg0)
	return *ret0, err
}

// DebitIdx is a free data retrieval call binding the contract method 0x203aa89c.
//
// Solidity: function debitIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessSession) DebitIdx(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.DebitIdx(&_Business.CallOpts, arg0)
}

// DebitIdx is a free data retrieval call binding the contract method 0x203aa89c.
//
// Solidity: function debitIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessCallerSession) DebitIdx(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.DebitIdx(&_Business.CallOpts, arg0)
}

// Debits is a free data retrieval call binding the contract method 0xa2cffc96.
//
// Solidity: function debits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Business *BusinessCaller) Debits(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _Business.contract.Call(opts, out, "debits", arg0)
	return *ret, err
}

// Debits is a free data retrieval call binding the contract method 0xa2cffc96.
//
// Solidity: function debits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Business *BusinessSession) Debits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Business.Contract.Debits(&_Business.CallOpts, arg0)
}

// Debits is a free data retrieval call binding the contract method 0xa2cffc96.
//
// Solidity: function debits(bytes32 ) constant returns(bytes32 txRef, bytes32 stashName, int256 amount, uint256 timestamp)
func (_Business *BusinessCallerSession) Debits(arg0 [32]byte) (struct {
	TxRef     [32]byte
	StashName [32]byte
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Business.Contract.Debits(&_Business.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _stashName) constant returns(int256)
func (_Business *BusinessCaller) GetBalance(opts *bind.CallOpts, _stashName [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getBalance", _stashName)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _stashName) constant returns(int256)
func (_Business *BusinessSession) GetBalance(_stashName [32]byte) (*big.Int, error) {
	return _Business.Contract.GetBalance(&_Business.CallOpts, _stashName)
}

// GetBalance is a free data retrieval call binding the contract method 0x8e739461.
//
// Solidity: function getBalance(bytes32 _stashName) constant returns(int256)
func (_Business *BusinessCallerSession) GetBalance(_stashName [32]byte) (*big.Int, error) {
	return _Business.Contract.GetBalance(&_Business.CallOpts, _stashName)
}

// GetCreditHistoryLength is a free data retrieval call binding the contract method 0x1d74f2e2.
//
// Solidity: function getCreditHistoryLength() constant returns(uint256)
func (_Business *BusinessCaller) GetCreditHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getCreditHistoryLength")
	return *ret0, err
}

// GetCreditHistoryLength is a free data retrieval call binding the contract method 0x1d74f2e2.
//
// Solidity: function getCreditHistoryLength() constant returns(uint256)
func (_Business *BusinessSession) GetCreditHistoryLength() (*big.Int, error) {
	return _Business.Contract.GetCreditHistoryLength(&_Business.CallOpts)
}

// GetCreditHistoryLength is a free data retrieval call binding the contract method 0x1d74f2e2.
//
// Solidity: function getCreditHistoryLength() constant returns(uint256)
func (_Business *BusinessCallerSession) GetCreditHistoryLength() (*big.Int, error) {
	return _Business.Contract.GetCreditHistoryLength(&_Business.CallOpts)
}

// GetDebitHistoryLength is a free data retrieval call binding the contract method 0x8d7a194a.
//
// Solidity: function getDebitHistoryLength() constant returns(uint256)
func (_Business *BusinessCaller) GetDebitHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getDebitHistoryLength")
	return *ret0, err
}

// GetDebitHistoryLength is a free data retrieval call binding the contract method 0x8d7a194a.
//
// Solidity: function getDebitHistoryLength() constant returns(uint256)
func (_Business *BusinessSession) GetDebitHistoryLength() (*big.Int, error) {
	return _Business.Contract.GetDebitHistoryLength(&_Business.CallOpts)
}

// GetDebitHistoryLength is a free data retrieval call binding the contract method 0x8d7a194a.
//
// Solidity: function getDebitHistoryLength() constant returns(uint256)
func (_Business *BusinessCallerSession) GetDebitHistoryLength() (*big.Int, error) {
	return _Business.Contract.GetDebitHistoryLength(&_Business.CallOpts)
}

// GetMemberApiIdxLenght is a free data retrieval call binding the contract method 0xe45ab733.
//
// Solidity: function getMemberApiIdxLenght() constant returns(int16)
func (_Business *BusinessCaller) GetMemberApiIdxLenght(opts *bind.CallOpts) (int16, error) {
	var (
		ret0 = new(int16)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getMemberApiIdxLenght")
	return *ret0, err
}

// GetMemberApiIdxLenght is a free data retrieval call binding the contract method 0xe45ab733.
//
// Solidity: function getMemberApiIdxLenght() constant returns(int16)
func (_Business *BusinessSession) GetMemberApiIdxLenght() (int16, error) {
	return _Business.Contract.GetMemberApiIdxLenght(&_Business.CallOpts)
}

// GetMemberApiIdxLenght is a free data retrieval call binding the contract method 0xe45ab733.
//
// Solidity: function getMemberApiIdxLenght() constant returns(int16)
func (_Business *BusinessCallerSession) GetMemberApiIdxLenght() (int16, error) {
	return _Business.Contract.GetMemberApiIdxLenght(&_Business.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Business *BusinessCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Business *BusinessSession) GetOwner() (common.Address, error) {
	return _Business.Contract.GetOwner(&_Business.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Business *BusinessCallerSession) GetOwner() (common.Address, error) {
	return _Business.Contract.GetOwner(&_Business.CallOpts)
}

// GetRegistedAccEthLength is a free data retrieval call binding the contract method 0xdf646612.
//
// Solidity: function getRegistedAccEthLength() constant returns(int16)
func (_Business *BusinessCaller) GetRegistedAccEthLength(opts *bind.CallOpts) (int16, error) {
	var (
		ret0 = new(int16)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getRegistedAccEthLength")
	return *ret0, err
}

// GetRegistedAccEthLength is a free data retrieval call binding the contract method 0xdf646612.
//
// Solidity: function getRegistedAccEthLength() constant returns(int16)
func (_Business *BusinessSession) GetRegistedAccEthLength() (int16, error) {
	return _Business.Contract.GetRegistedAccEthLength(&_Business.CallOpts)
}

// GetRegistedAccEthLength is a free data retrieval call binding the contract method 0xdf646612.
//
// Solidity: function getRegistedAccEthLength() constant returns(int16)
func (_Business *BusinessCallerSession) GetRegistedAccEthLength() (int16, error) {
	return _Business.Contract.GetRegistedAccEthLength(&_Business.CallOpts)
}

// GetStashNamesLenght is a free data retrieval call binding the contract method 0xee9a5812.
//
// Solidity: function getStashNamesLenght() constant returns(int256)
func (_Business *BusinessCaller) GetStashNamesLenght(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getStashNamesLenght")
	return *ret0, err
}

// GetStashNamesLenght is a free data retrieval call binding the contract method 0xee9a5812.
//
// Solidity: function getStashNamesLenght() constant returns(int256)
func (_Business *BusinessSession) GetStashNamesLenght() (*big.Int, error) {
	return _Business.Contract.GetStashNamesLenght(&_Business.CallOpts)
}

// GetStashNamesLenght is a free data retrieval call binding the contract method 0xee9a5812.
//
// Solidity: function getStashNamesLenght() constant returns(int256)
func (_Business *BusinessCallerSession) GetStashNamesLenght() (*big.Int, error) {
	return _Business.Contract.GetStashNamesLenght(&_Business.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 _stashName) constant returns(int8)
func (_Business *BusinessCaller) GetState(opts *bind.CallOpts, _stashName [32]byte) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getState", _stashName)
	return *ret0, err
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 _stashName) constant returns(int8)
func (_Business *BusinessSession) GetState(_stashName [32]byte) (int8, error) {
	return _Business.Contract.GetState(&_Business.CallOpts, _stashName)
}

// GetState is a free data retrieval call binding the contract method 0x09648a9d.
//
// Solidity: function getState(bytes32 _stashName) constant returns(int8)
func (_Business *BusinessCallerSession) GetState(_stashName [32]byte) (int8, error) {
	return _Business.Contract.GetState(&_Business.CallOpts, _stashName)
}

// GetTransferHistoryLength is a free data retrieval call binding the contract method 0xd236e1ba.
//
// Solidity: function getTransferHistoryLength() constant returns(uint256)
func (_Business *BusinessCaller) GetTransferHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getTransferHistoryLength")
	return *ret0, err
}

// GetTransferHistoryLength is a free data retrieval call binding the contract method 0xd236e1ba.
//
// Solidity: function getTransferHistoryLength() constant returns(uint256)
func (_Business *BusinessSession) GetTransferHistoryLength() (*big.Int, error) {
	return _Business.Contract.GetTransferHistoryLength(&_Business.CallOpts)
}

// GetTransferHistoryLength is a free data retrieval call binding the contract method 0xd236e1ba.
//
// Solidity: function getTransferHistoryLength() constant returns(uint256)
func (_Business *BusinessCallerSession) GetTransferHistoryLength() (*big.Int, error) {
	return _Business.Contract.GetTransferHistoryLength(&_Business.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _stashName) constant returns(int8)
func (_Business *BusinessCaller) GetType(opts *bind.CallOpts, _stashName [32]byte) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "getType", _stashName)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _stashName) constant returns(int8)
func (_Business *BusinessSession) GetType(_stashName [32]byte) (int8, error) {
	return _Business.Contract.GetType(&_Business.CallOpts, _stashName)
}

// GetType is a free data retrieval call binding the contract method 0x7f2aeea4.
//
// Solidity: function getType(bytes32 _stashName) constant returns(int8)
func (_Business *BusinessCallerSession) GetType(_stashName [32]byte) (int8, error) {
	return _Business.Contract.GetType(&_Business.CallOpts, _stashName)
}

// StashNames is a free data retrieval call binding the contract method 0x0ae94616.
//
// Solidity: function stashNames(uint256 ) constant returns(bytes32)
func (_Business *BusinessCaller) StashNames(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "stashNames", arg0)
	return *ret0, err
}

// StashNames is a free data retrieval call binding the contract method 0x0ae94616.
//
// Solidity: function stashNames(uint256 ) constant returns(bytes32)
func (_Business *BusinessSession) StashNames(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.StashNames(&_Business.CallOpts, arg0)
}

// StashNames is a free data retrieval call binding the contract method 0x0ae94616.
//
// Solidity: function stashNames(uint256 ) constant returns(bytes32)
func (_Business *BusinessCallerSession) StashNames(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.StashNames(&_Business.CallOpts, arg0)
}

// StashRegistry is a free data retrieval call binding the contract method 0xfb6cb61f.
//
// Solidity: function stashRegistry(bytes32 ) constant returns(address)
func (_Business *BusinessCaller) StashRegistry(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "stashRegistry", arg0)
	return *ret0, err
}

// StashRegistry is a free data retrieval call binding the contract method 0xfb6cb61f.
//
// Solidity: function stashRegistry(bytes32 ) constant returns(address)
func (_Business *BusinessSession) StashRegistry(arg0 [32]byte) (common.Address, error) {
	return _Business.Contract.StashRegistry(&_Business.CallOpts, arg0)
}

// StashRegistry is a free data retrieval call binding the contract method 0xfb6cb61f.
//
// Solidity: function stashRegistry(bytes32 ) constant returns(address)
func (_Business *BusinessCallerSession) StashRegistry(arg0 [32]byte) (common.Address, error) {
	return _Business.Contract.StashRegistry(&_Business.CallOpts, arg0)
}

// StringTobytes is a free data retrieval call binding the contract method 0xbc6d0577.
//
// Solidity: function string_tobytes(string s) constant returns(bytes)
func (_Business *BusinessCaller) StringTobytes(opts *bind.CallOpts, s string) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "string_tobytes", s)
	return *ret0, err
}

// StringTobytes is a free data retrieval call binding the contract method 0xbc6d0577.
//
// Solidity: function string_tobytes(string s) constant returns(bytes)
func (_Business *BusinessSession) StringTobytes(s string) ([]byte, error) {
	return _Business.Contract.StringTobytes(&_Business.CallOpts, s)
}

// StringTobytes is a free data retrieval call binding the contract method 0xbc6d0577.
//
// Solidity: function string_tobytes(string s) constant returns(bytes)
func (_Business *BusinessCallerSession) StringTobytes(s string) ([]byte, error) {
	return _Business.Contract.StringTobytes(&_Business.CallOpts, s)
}

// TransferIdx is a free data retrieval call binding the contract method 0xc83fe142.
//
// Solidity: function transferIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessCaller) TransferIdx(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Business.contract.Call(opts, out, "transferIdx", arg0)
	return *ret0, err
}

// TransferIdx is a free data retrieval call binding the contract method 0xc83fe142.
//
// Solidity: function transferIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessSession) TransferIdx(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.TransferIdx(&_Business.CallOpts, arg0)
}

// TransferIdx is a free data retrieval call binding the contract method 0xc83fe142.
//
// Solidity: function transferIdx(uint256 ) constant returns(bytes32)
func (_Business *BusinessCallerSession) TransferIdx(arg0 *big.Int) ([32]byte, error) {
	return _Business.Contract.TransferIdx(&_Business.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 txRef, bytes32 sender, bytes32 receiver, int256 amount, string note, int8 txType, uint256 timestamp)
func (_Business *BusinessCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _Business.contract.Call(opts, out, "transfers", arg0)
	return *ret, err
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 txRef, bytes32 sender, bytes32 receiver, int256 amount, string note, int8 txType, uint256 timestamp)
func (_Business *BusinessSession) Transfers(arg0 [32]byte) (struct {
	TxRef     [32]byte
	Sender    [32]byte
	Receiver  [32]byte
	Amount    *big.Int
	Note      string
	TxType    int8
	Timestamp *big.Int
}, error) {
	return _Business.Contract.Transfers(&_Business.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) constant returns(bytes32 txRef, bytes32 sender, bytes32 receiver, int256 amount, string note, int8 txType, uint256 timestamp)
func (_Business *BusinessCallerSession) Transfers(arg0 [32]byte) (struct {
	TxRef     [32]byte
	Sender    [32]byte
	Receiver  [32]byte
	Amount    *big.Int
	Note      string
	TxType    int8
	Timestamp *big.Int
}, error) {
	return _Business.Contract.Transfers(&_Business.CallOpts, arg0)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Business *BusinessTransactor) ChangeOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "changeOwner", _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Business *BusinessSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Business.Contract.ChangeOwner(&_Business.TransactOpts, _newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _newOwner) returns()
func (_Business *BusinessTransactorSession) ChangeOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Business.Contract.ChangeOwner(&_Business.TransactOpts, _newOwner)
}

// ChangeOwnerAllStash is a paid mutator transaction binding the contract method 0xeae5180e.
//
// Solidity: function changeOwnerAllStash(address _newOwner) returns()
func (_Business *BusinessTransactor) ChangeOwnerAllStash(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "changeOwnerAllStash", _newOwner)
}

// ChangeOwnerAllStash is a paid mutator transaction binding the contract method 0xeae5180e.
//
// Solidity: function changeOwnerAllStash(address _newOwner) returns()
func (_Business *BusinessSession) ChangeOwnerAllStash(_newOwner common.Address) (*types.Transaction, error) {
	return _Business.Contract.ChangeOwnerAllStash(&_Business.TransactOpts, _newOwner)
}

// ChangeOwnerAllStash is a paid mutator transaction binding the contract method 0xeae5180e.
//
// Solidity: function changeOwnerAllStash(address _newOwner) returns()
func (_Business *BusinessTransactorSession) ChangeOwnerAllStash(_newOwner common.Address) (*types.Transaction, error) {
	return _Business.Contract.ChangeOwnerAllStash(&_Business.TransactOpts, _newOwner)
}

// CreateStash is a paid mutator transaction binding the contract method 0x7e53308a.
//
// Solidity: function createStash(bytes32 _stashName, int8 _typeState) returns()
func (_Business *BusinessTransactor) CreateStash(opts *bind.TransactOpts, _stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "createStash", _stashName, _typeState)
}

// CreateStash is a paid mutator transaction binding the contract method 0x7e53308a.
//
// Solidity: function createStash(bytes32 _stashName, int8 _typeState) returns()
func (_Business *BusinessSession) CreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Business.Contract.CreateStash(&_Business.TransactOpts, _stashName, _typeState)
}

// CreateStash is a paid mutator transaction binding the contract method 0x7e53308a.
//
// Solidity: function createStash(bytes32 _stashName, int8 _typeState) returns()
func (_Business *BusinessTransactorSession) CreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Business.Contract.CreateStash(&_Business.TransactOpts, _stashName, _typeState)
}

// Credit is a paid mutator transaction binding the contract method 0xb91f176e.
//
// Solidity: function credit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Business *BusinessTransactor) Credit(opts *bind.TransactOpts, _txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "credit", _txRef, _stashName, _amount)
}

// Credit is a paid mutator transaction binding the contract method 0xb91f176e.
//
// Solidity: function credit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Business *BusinessSession) Credit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Business.Contract.Credit(&_Business.TransactOpts, _txRef, _stashName, _amount)
}

// Credit is a paid mutator transaction binding the contract method 0xb91f176e.
//
// Solidity: function credit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Business *BusinessTransactorSession) Credit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Business.Contract.Credit(&_Business.TransactOpts, _txRef, _stashName, _amount)
}

// Debit is a paid mutator transaction binding the contract method 0x4ea679e7.
//
// Solidity: function debit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Business *BusinessTransactor) Debit(opts *bind.TransactOpts, _txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "debit", _txRef, _stashName, _amount)
}

// Debit is a paid mutator transaction binding the contract method 0x4ea679e7.
//
// Solidity: function debit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Business *BusinessSession) Debit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Business.Contract.Debit(&_Business.TransactOpts, _txRef, _stashName, _amount)
}

// Debit is a paid mutator transaction binding the contract method 0x4ea679e7.
//
// Solidity: function debit(bytes32 _txRef, bytes32 _stashName, int256 _amount) returns()
func (_Business *BusinessTransactorSession) Debit(_txRef [32]byte, _stashName [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Business.Contract.Debit(&_Business.TransactOpts, _txRef, _stashName, _amount)
}

// LoadStashRegistry is a paid mutator transaction binding the contract method 0x4e3d7de0.
//
// Solidity: function loadStashRegistry(bytes32 _stashName, address _stash) returns()
func (_Business *BusinessTransactor) LoadStashRegistry(opts *bind.TransactOpts, _stashName [32]byte, _stash common.Address) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "loadStashRegistry", _stashName, _stash)
}

// LoadStashRegistry is a paid mutator transaction binding the contract method 0x4e3d7de0.
//
// Solidity: function loadStashRegistry(bytes32 _stashName, address _stash) returns()
func (_Business *BusinessSession) LoadStashRegistry(_stashName [32]byte, _stash common.Address) (*types.Transaction, error) {
	return _Business.Contract.LoadStashRegistry(&_Business.TransactOpts, _stashName, _stash)
}

// LoadStashRegistry is a paid mutator transaction binding the contract method 0x4e3d7de0.
//
// Solidity: function loadStashRegistry(bytes32 _stashName, address _stash) returns()
func (_Business *BusinessTransactorSession) LoadStashRegistry(_stashName [32]byte, _stash common.Address) (*types.Transaction, error) {
	return _Business.Contract.LoadStashRegistry(&_Business.TransactOpts, _stashName, _stash)
}

// ReCreateStash is a paid mutator transaction binding the contract method 0x1e022e75.
//
// Solidity: function reCreateStash(bytes32 _stashName, int8 _typeState) returns()
func (_Business *BusinessTransactor) ReCreateStash(opts *bind.TransactOpts, _stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "reCreateStash", _stashName, _typeState)
}

// ReCreateStash is a paid mutator transaction binding the contract method 0x1e022e75.
//
// Solidity: function reCreateStash(bytes32 _stashName, int8 _typeState) returns()
func (_Business *BusinessSession) ReCreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Business.Contract.ReCreateStash(&_Business.TransactOpts, _stashName, _typeState)
}

// ReCreateStash is a paid mutator transaction binding the contract method 0x1e022e75.
//
// Solidity: function reCreateStash(bytes32 _stashName, int8 _typeState) returns()
func (_Business *BusinessTransactorSession) ReCreateStash(_stashName [32]byte, _typeState int8) (*types.Transaction, error) {
	return _Business.Contract.ReCreateStash(&_Business.TransactOpts, _stashName, _typeState)
}

// RegisterAccETH is a paid mutator transaction binding the contract method 0xe1b8ff83.
//
// Solidity: function registerAccETH(address[] _listAcc) returns()
func (_Business *BusinessTransactor) RegisterAccETH(opts *bind.TransactOpts, _listAcc []common.Address) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "registerAccETH", _listAcc)
}

// RegisterAccETH is a paid mutator transaction binding the contract method 0xe1b8ff83.
//
// Solidity: function registerAccETH(address[] _listAcc) returns()
func (_Business *BusinessSession) RegisterAccETH(_listAcc []common.Address) (*types.Transaction, error) {
	return _Business.Contract.RegisterAccETH(&_Business.TransactOpts, _listAcc)
}

// RegisterAccETH is a paid mutator transaction binding the contract method 0xe1b8ff83.
//
// Solidity: function registerAccETH(address[] _listAcc) returns()
func (_Business *BusinessTransactorSession) RegisterAccETH(_listAcc []common.Address) (*types.Transaction, error) {
	return _Business.Contract.RegisterAccETH(&_Business.TransactOpts, _listAcc)
}

// RegisterMemberApi is a paid mutator transaction binding the contract method 0x0d3b2b6f.
//
// Solidity: function registerMemberApi(address _newMember) returns()
func (_Business *BusinessTransactor) RegisterMemberApi(opts *bind.TransactOpts, _newMember common.Address) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "registerMemberApi", _newMember)
}

// RegisterMemberApi is a paid mutator transaction binding the contract method 0x0d3b2b6f.
//
// Solidity: function registerMemberApi(address _newMember) returns()
func (_Business *BusinessSession) RegisterMemberApi(_newMember common.Address) (*types.Transaction, error) {
	return _Business.Contract.RegisterMemberApi(&_Business.TransactOpts, _newMember)
}

// RegisterMemberApi is a paid mutator transaction binding the contract method 0x0d3b2b6f.
//
// Solidity: function registerMemberApi(address _newMember) returns()
func (_Business *BusinessTransactorSession) RegisterMemberApi(_newMember common.Address) (*types.Transaction, error) {
	return _Business.Contract.RegisterMemberApi(&_Business.TransactOpts, _newMember)
}

// SetState is a paid mutator transaction binding the contract method 0x8b356a4d.
//
// Solidity: function setState(bytes32 _stashName, int8 _stashState) returns()
func (_Business *BusinessTransactor) SetState(opts *bind.TransactOpts, _stashName [32]byte, _stashState int8) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "setState", _stashName, _stashState)
}

// SetState is a paid mutator transaction binding the contract method 0x8b356a4d.
//
// Solidity: function setState(bytes32 _stashName, int8 _stashState) returns()
func (_Business *BusinessSession) SetState(_stashName [32]byte, _stashState int8) (*types.Transaction, error) {
	return _Business.Contract.SetState(&_Business.TransactOpts, _stashName, _stashState)
}

// SetState is a paid mutator transaction binding the contract method 0x8b356a4d.
//
// Solidity: function setState(bytes32 _stashName, int8 _stashState) returns()
func (_Business *BusinessTransactorSession) SetState(_stashName [32]byte, _stashState int8) (*types.Transaction, error) {
	return _Business.Contract.SetState(&_Business.TransactOpts, _stashName, _stashState)
}

// Transfer is a paid mutator transaction binding the contract method 0xa4642866.
//
// Solidity: function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int256 _amount, string _note, int8 _txType) returns(int256 sender_bal, int256 receiver_bal)
func (_Business *BusinessTransactor) Transfer(opts *bind.TransactOpts, _txRef [32]byte, _sender [32]byte, _receiver [32]byte, _amount *big.Int, _note string, _txType int8) (*types.Transaction, error) {
	return _Business.contract.Transact(opts, "transfer", _txRef, _sender, _receiver, _amount, _note, _txType)
}

// Transfer is a paid mutator transaction binding the contract method 0xa4642866.
//
// Solidity: function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int256 _amount, string _note, int8 _txType) returns(int256 sender_bal, int256 receiver_bal)
func (_Business *BusinessSession) Transfer(_txRef [32]byte, _sender [32]byte, _receiver [32]byte, _amount *big.Int, _note string, _txType int8) (*types.Transaction, error) {
	return _Business.Contract.Transfer(&_Business.TransactOpts, _txRef, _sender, _receiver, _amount, _note, _txType)
}

// Transfer is a paid mutator transaction binding the contract method 0xa4642866.
//
// Solidity: function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int256 _amount, string _note, int8 _txType) returns(int256 sender_bal, int256 receiver_bal)
func (_Business *BusinessTransactorSession) Transfer(_txRef [32]byte, _sender [32]byte, _receiver [32]byte, _amount *big.Int, _note string, _txType int8) (*types.Transaction, error) {
	return _Business.Contract.Transfer(&_Business.TransactOpts, _txRef, _sender, _receiver, _amount, _note, _txType)
}

// BusinessEventCreateStashIterator is returned from FilterEventCreateStash and is used to iterate over the raw logs and unpacked data for EventCreateStash events raised by the Business contract.
type BusinessEventCreateStashIterator struct {
	Event *BusinessEventCreateStash // Event containing the contract specifics and raw log

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
func (it *BusinessEventCreateStashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BusinessEventCreateStash)
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
		it.Event = new(BusinessEventCreateStash)
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
func (it *BusinessEventCreateStashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BusinessEventCreateStashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BusinessEventCreateStash represents a EventCreateStash event raised by the Business contract.
type BusinessEventCreateStash struct {
	WalletCode    [32]byte
	WalletAddress common.Address
	WalletType    int8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventCreateStash is a free log retrieval operation binding the contract event 0x0cb60d28f723451a1bd65a935f5e1d7b4e708818cf6cc6c30ad742fd67531b00.
//
// Solidity: event event_createStash(bytes32 indexed wallet_code, address wallet_address, int8 wallet_type)
func (_Business *BusinessFilterer) FilterEventCreateStash(opts *bind.FilterOpts, wallet_code [][32]byte) (*BusinessEventCreateStashIterator, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.FilterLogs(opts, "event_createStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &BusinessEventCreateStashIterator{contract: _Business.contract, event: "event_createStash", logs: logs, sub: sub}, nil
}

// WatchEventCreateStash is a free log subscription operation binding the contract event 0x0cb60d28f723451a1bd65a935f5e1d7b4e708818cf6cc6c30ad742fd67531b00.
//
// Solidity: event event_createStash(bytes32 indexed wallet_code, address wallet_address, int8 wallet_type)
func (_Business *BusinessFilterer) WatchEventCreateStash(opts *bind.WatchOpts, sink chan<- *BusinessEventCreateStash, wallet_code [][32]byte) (event.Subscription, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.WatchLogs(opts, "event_createStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BusinessEventCreateStash)
				if err := _Business.contract.UnpackLog(event, "event_createStash", log); err != nil {
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

// BusinessEventCreditIterator is returned from FilterEventCredit and is used to iterate over the raw logs and unpacked data for EventCredit events raised by the Business contract.
type BusinessEventCreditIterator struct {
	Event *BusinessEventCredit // Event containing the contract specifics and raw log

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
func (it *BusinessEventCreditIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BusinessEventCredit)
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
		it.Event = new(BusinessEventCredit)
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
func (it *BusinessEventCreditIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BusinessEventCreditIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BusinessEventCredit represents a EventCredit event raised by the Business contract.
type BusinessEventCredit struct {
	TxRef      [32]byte
	WalletCode [32]byte
	Amount     *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventCredit is a free log retrieval operation binding the contract event 0x2bda1f3aff897895cbea7f906b2e41ef99d39163a7b36af5782b8b718ed7e456.
//
// Solidity: event event_credit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Business *BusinessFilterer) FilterEventCredit(opts *bind.FilterOpts, txRef [][32]byte, wallet_code [][32]byte) (*BusinessEventCreditIterator, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.FilterLogs(opts, "event_credit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &BusinessEventCreditIterator{contract: _Business.contract, event: "event_credit", logs: logs, sub: sub}, nil
}

// WatchEventCredit is a free log subscription operation binding the contract event 0x2bda1f3aff897895cbea7f906b2e41ef99d39163a7b36af5782b8b718ed7e456.
//
// Solidity: event event_credit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Business *BusinessFilterer) WatchEventCredit(opts *bind.WatchOpts, sink chan<- *BusinessEventCredit, txRef [][32]byte, wallet_code [][32]byte) (event.Subscription, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.WatchLogs(opts, "event_credit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BusinessEventCredit)
				if err := _Business.contract.UnpackLog(event, "event_credit", log); err != nil {
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

// BusinessEventDebitIterator is returned from FilterEventDebit and is used to iterate over the raw logs and unpacked data for EventDebit events raised by the Business contract.
type BusinessEventDebitIterator struct {
	Event *BusinessEventDebit // Event containing the contract specifics and raw log

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
func (it *BusinessEventDebitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BusinessEventDebit)
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
		it.Event = new(BusinessEventDebit)
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
func (it *BusinessEventDebitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BusinessEventDebitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BusinessEventDebit represents a EventDebit event raised by the Business contract.
type BusinessEventDebit struct {
	TxRef      [32]byte
	WalletCode [32]byte
	Amount     *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventDebit is a free log retrieval operation binding the contract event 0x3e888fb580f7ba504d1d06b6af38e430c683d120af553208a48629dbd58bb0f8.
//
// Solidity: event event_debit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Business *BusinessFilterer) FilterEventDebit(opts *bind.FilterOpts, txRef [][32]byte, wallet_code [][32]byte) (*BusinessEventDebitIterator, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.FilterLogs(opts, "event_debit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &BusinessEventDebitIterator{contract: _Business.contract, event: "event_debit", logs: logs, sub: sub}, nil
}

// WatchEventDebit is a free log subscription operation binding the contract event 0x3e888fb580f7ba504d1d06b6af38e430c683d120af553208a48629dbd58bb0f8.
//
// Solidity: event event_debit(bytes32 indexed txRef, bytes32 indexed wallet_code, int256 amount, uint256 timestamp)
func (_Business *BusinessFilterer) WatchEventDebit(opts *bind.WatchOpts, sink chan<- *BusinessEventDebit, txRef [][32]byte, wallet_code [][32]byte) (event.Subscription, error) {

	var txRefRule []interface{}
	for _, txRefItem := range txRef {
		txRefRule = append(txRefRule, txRefItem)
	}
	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.WatchLogs(opts, "event_debit", txRefRule, wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BusinessEventDebit)
				if err := _Business.contract.UnpackLog(event, "event_debit", log); err != nil {
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

// BusinessEventReCreateStashIterator is returned from FilterEventReCreateStash and is used to iterate over the raw logs and unpacked data for EventReCreateStash events raised by the Business contract.
type BusinessEventReCreateStashIterator struct {
	Event *BusinessEventReCreateStash // Event containing the contract specifics and raw log

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
func (it *BusinessEventReCreateStashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BusinessEventReCreateStash)
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
		it.Event = new(BusinessEventReCreateStash)
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
func (it *BusinessEventReCreateStashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BusinessEventReCreateStashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BusinessEventReCreateStash represents a EventReCreateStash event raised by the Business contract.
type BusinessEventReCreateStash struct {
	WalletCode       [32]byte
	OldWalletAddress common.Address
	NewWalletAddress common.Address
	WalletType       int8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEventReCreateStash is a free log retrieval operation binding the contract event 0x8c8e3419f9e71366ada9d23ae8a8dba5229a20e91deb3cab00f419ea5b0480e8.
//
// Solidity: event event_reCreateStash(bytes32 indexed wallet_code, address old_wallet_address, address new_wallet_address, int8 wallet_type)
func (_Business *BusinessFilterer) FilterEventReCreateStash(opts *bind.FilterOpts, wallet_code [][32]byte) (*BusinessEventReCreateStashIterator, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.FilterLogs(opts, "event_reCreateStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &BusinessEventReCreateStashIterator{contract: _Business.contract, event: "event_reCreateStash", logs: logs, sub: sub}, nil
}

// WatchEventReCreateStash is a free log subscription operation binding the contract event 0x8c8e3419f9e71366ada9d23ae8a8dba5229a20e91deb3cab00f419ea5b0480e8.
//
// Solidity: event event_reCreateStash(bytes32 indexed wallet_code, address old_wallet_address, address new_wallet_address, int8 wallet_type)
func (_Business *BusinessFilterer) WatchEventReCreateStash(opts *bind.WatchOpts, sink chan<- *BusinessEventReCreateStash, wallet_code [][32]byte) (event.Subscription, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.WatchLogs(opts, "event_reCreateStash", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BusinessEventReCreateStash)
				if err := _Business.contract.UnpackLog(event, "event_reCreateStash", log); err != nil {
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

// BusinessEventRegisterAccETHIterator is returned from FilterEventRegisterAccETH and is used to iterate over the raw logs and unpacked data for EventRegisterAccETH events raised by the Business contract.
type BusinessEventRegisterAccETHIterator struct {
	Event *BusinessEventRegisterAccETH // Event containing the contract specifics and raw log

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
func (it *BusinessEventRegisterAccETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BusinessEventRegisterAccETH)
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
		it.Event = new(BusinessEventRegisterAccETH)
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
func (it *BusinessEventRegisterAccETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BusinessEventRegisterAccETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BusinessEventRegisterAccETH represents a EventRegisterAccETH event raised by the Business contract.
type BusinessEventRegisterAccETH struct {
	ListAcc []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventRegisterAccETH is a free log retrieval operation binding the contract event 0x3b0686a3c64982476da002e36939997a11d9ef748420920b111c3ba5d2e9784f.
//
// Solidity: event event_registerAccETH(address[] listAcc)
func (_Business *BusinessFilterer) FilterEventRegisterAccETH(opts *bind.FilterOpts) (*BusinessEventRegisterAccETHIterator, error) {

	logs, sub, err := _Business.contract.FilterLogs(opts, "event_registerAccETH")
	if err != nil {
		return nil, err
	}
	return &BusinessEventRegisterAccETHIterator{contract: _Business.contract, event: "event_registerAccETH", logs: logs, sub: sub}, nil
}

// WatchEventRegisterAccETH is a free log subscription operation binding the contract event 0x3b0686a3c64982476da002e36939997a11d9ef748420920b111c3ba5d2e9784f.
//
// Solidity: event event_registerAccETH(address[] listAcc)
func (_Business *BusinessFilterer) WatchEventRegisterAccETH(opts *bind.WatchOpts, sink chan<- *BusinessEventRegisterAccETH) (event.Subscription, error) {

	logs, sub, err := _Business.contract.WatchLogs(opts, "event_registerAccETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BusinessEventRegisterAccETH)
				if err := _Business.contract.UnpackLog(event, "event_registerAccETH", log); err != nil {
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

// BusinessEventSetStateIterator is returned from FilterEventSetState and is used to iterate over the raw logs and unpacked data for EventSetState events raised by the Business contract.
type BusinessEventSetStateIterator struct {
	Event *BusinessEventSetState // Event containing the contract specifics and raw log

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
func (it *BusinessEventSetStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BusinessEventSetState)
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
		it.Event = new(BusinessEventSetState)
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
func (it *BusinessEventSetStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BusinessEventSetStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BusinessEventSetState represents a EventSetState event raised by the Business contract.
type BusinessEventSetState struct {
	WalletCode [32]byte
	StashState int8
	OldState   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEventSetState is a free log retrieval operation binding the contract event 0x787d29ad64e9e02e1fb29fb54868408884d3969457c61da4943ec8d57e8ac011.
//
// Solidity: event event_setState(bytes32 indexed wallet_code, int8 stashState, int256 oldState)
func (_Business *BusinessFilterer) FilterEventSetState(opts *bind.FilterOpts, wallet_code [][32]byte) (*BusinessEventSetStateIterator, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.FilterLogs(opts, "event_setState", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return &BusinessEventSetStateIterator{contract: _Business.contract, event: "event_setState", logs: logs, sub: sub}, nil
}

// WatchEventSetState is a free log subscription operation binding the contract event 0x787d29ad64e9e02e1fb29fb54868408884d3969457c61da4943ec8d57e8ac011.
//
// Solidity: event event_setState(bytes32 indexed wallet_code, int8 stashState, int256 oldState)
func (_Business *BusinessFilterer) WatchEventSetState(opts *bind.WatchOpts, sink chan<- *BusinessEventSetState, wallet_code [][32]byte) (event.Subscription, error) {

	var wallet_codeRule []interface{}
	for _, wallet_codeItem := range wallet_code {
		wallet_codeRule = append(wallet_codeRule, wallet_codeItem)
	}

	logs, sub, err := _Business.contract.WatchLogs(opts, "event_setState", wallet_codeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BusinessEventSetState)
				if err := _Business.contract.UnpackLog(event, "event_setState", log); err != nil {
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

// BusinessEventTransferIterator is returned from FilterEventTransfer and is used to iterate over the raw logs and unpacked data for EventTransfer events raised by the Business contract.
type BusinessEventTransferIterator struct {
	Event *BusinessEventTransfer // Event containing the contract specifics and raw log

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
func (it *BusinessEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BusinessEventTransfer)
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
		it.Event = new(BusinessEventTransfer)
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
func (it *BusinessEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BusinessEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BusinessEventTransfer represents a EventTransfer event raised by the Business contract.
type BusinessEventTransfer struct {
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
func (_Business *BusinessFilterer) FilterEventTransfer(opts *bind.FilterOpts, txRef [][32]byte, sender [][32]byte, receiver [][32]byte) (*BusinessEventTransferIterator, error) {

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

	logs, sub, err := _Business.contract.FilterLogs(opts, "event_transfer", txRefRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &BusinessEventTransferIterator{contract: _Business.contract, event: "event_transfer", logs: logs, sub: sub}, nil
}

// WatchEventTransfer is a free log subscription operation binding the contract event 0x20244c6c280a8d1e1be518a3fc3acb8c46f303ad98dd9ec4c01baf9e2999f82e.
//
// Solidity: event event_transfer(bytes32 indexed txRef, bytes32 indexed sender, bytes32 indexed receiver, int256 amount, string note, int8 txType, int256 sender_bal, int256 receiver_bal, uint256 timestamp)
func (_Business *BusinessFilterer) WatchEventTransfer(opts *bind.WatchOpts, sink chan<- *BusinessEventTransfer, txRef [][32]byte, sender [][32]byte, receiver [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Business.contract.WatchLogs(opts, "event_transfer", txRefRule, senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BusinessEventTransfer)
				if err := _Business.contract.UnpackLog(event, "event_transfer", log); err != nil {
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
