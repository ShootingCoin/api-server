// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GameCoreBetInfo is an auto generated low-level Go binding around an user-defined struct.
type GameCoreBetInfo struct {
	CoinAddress common.Address
	BetAmount   *big.Int
	NftSkinId   [5]*big.Int
}

// GameCoreGameHistory is an auto generated low-level Go binding around an user-defined struct.
type GameCoreGameHistory struct {
	GameId           *big.Int
	User1            common.Address
	User1coinAddress common.Address
	User1GetAmount   *big.Int
	User2            common.Address
	User2coinAddress common.Address
	User2GetAmount   *big.Int
	TimeStamp        *big.Int
}

// GameCoreGameInfo is an auto generated low-level Go binding around an user-defined struct.
type GameCoreGameInfo struct {
	User1        common.Address
	User2        common.Address
	User1BetInfo GameCoreBetInfo
	User2BetInfo GameCoreBetInfo
}

// ShootingCoinManagerMetaData contains all meta data concerning the ShootingCoinManager contract.
var ShootingCoinManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"indexed\":false,\"internalType\":\"structGameCore.BetInfo\",\"name\":\"betInfo\",\"type\":\"tuple\"}],\"name\":\"Entered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"user1BetInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"user2BetInfo\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structGameCore.GameInfo\",\"name\":\"gameInfo\",\"type\":\"tuple\"}],\"name\":\"GameInited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user1coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"user1GetAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"user2GetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint240\",\"name\":\"timeStamp\",\"type\":\"uint240\"}],\"indexed\":false,\"internalType\":\"structGameCore.GameHistory\",\"name\":\"gameHistory\",\"type\":\"tuple\"}],\"name\":\"GameSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"betInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkOnGame\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"despositCoin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"_betInfo\",\"type\":\"tuple\"}],\"name\":\"enterGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecieveAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user1coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"user1GetAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"user2GetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint240\",\"name\":\"timeStamp\",\"type\":\"uint240\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gameInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"user1BetInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"user2BetInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAccount\",\"type\":\"address\"}],\"name\":\"getBetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"getGameInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"user1BetInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"betAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[5]\",\"name\":\"nftSkinId\",\"type\":\"uint256[5]\"}],\"internalType\":\"structGameCore.BetInfo\",\"name\":\"user2BetInfo\",\"type\":\"tuple\"}],\"internalType\":\"structGameCore.GameInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAccount\",\"type\":\"address\"}],\"name\":\"getHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user1coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"user1GetAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2coinAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"user2GetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint240\",\"name\":\"timeStamp\",\"type\":\"uint240\"}],\"internalType\":\"structGameCore.GameHistory[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getShootingNft\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getShootingRole\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roleContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gameFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeRecieveAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOnGame\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"user1GetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"user2GetAmount\",\"type\":\"uint256\"}],\"name\":\"settleGame\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shootingNft\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shootingRole\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user2\",\"type\":\"address\"}],\"name\":\"startGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"}],\"name\":\"updateShootingNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"roleContract\",\"type\":\"address\"}],\"name\":\"updateShootingRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ShootingCoinManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ShootingCoinManagerMetaData.ABI instead.
var ShootingCoinManagerABI = ShootingCoinManagerMetaData.ABI

// ShootingCoinManager is an auto generated Go binding around an Ethereum contract.
type ShootingCoinManager struct {
	ShootingCoinManagerCaller     // Read-only binding to the contract
	ShootingCoinManagerTransactor // Write-only binding to the contract
	ShootingCoinManagerFilterer   // Log filterer for contract events
}

// ShootingCoinManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShootingCoinManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShootingCoinManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShootingCoinManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShootingCoinManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShootingCoinManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShootingCoinManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShootingCoinManagerSession struct {
	Contract     *ShootingCoinManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ShootingCoinManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShootingCoinManagerCallerSession struct {
	Contract *ShootingCoinManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ShootingCoinManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShootingCoinManagerTransactorSession struct {
	Contract     *ShootingCoinManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ShootingCoinManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShootingCoinManagerRaw struct {
	Contract *ShootingCoinManager // Generic contract binding to access the raw methods on
}

// ShootingCoinManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShootingCoinManagerCallerRaw struct {
	Contract *ShootingCoinManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ShootingCoinManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShootingCoinManagerTransactorRaw struct {
	Contract *ShootingCoinManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShootingCoinManager creates a new instance of ShootingCoinManager, bound to a specific deployed contract.
func NewShootingCoinManager(address common.Address, backend bind.ContractBackend) (*ShootingCoinManager, error) {
	contract, err := bindShootingCoinManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManager{ShootingCoinManagerCaller: ShootingCoinManagerCaller{contract: contract}, ShootingCoinManagerTransactor: ShootingCoinManagerTransactor{contract: contract}, ShootingCoinManagerFilterer: ShootingCoinManagerFilterer{contract: contract}}, nil
}

// NewShootingCoinManagerCaller creates a new read-only instance of ShootingCoinManager, bound to a specific deployed contract.
func NewShootingCoinManagerCaller(address common.Address, caller bind.ContractCaller) (*ShootingCoinManagerCaller, error) {
	contract, err := bindShootingCoinManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManagerCaller{contract: contract}, nil
}

// NewShootingCoinManagerTransactor creates a new write-only instance of ShootingCoinManager, bound to a specific deployed contract.
func NewShootingCoinManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ShootingCoinManagerTransactor, error) {
	contract, err := bindShootingCoinManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManagerTransactor{contract: contract}, nil
}

// NewShootingCoinManagerFilterer creates a new log filterer instance of ShootingCoinManager, bound to a specific deployed contract.
func NewShootingCoinManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ShootingCoinManagerFilterer, error) {
	contract, err := bindShootingCoinManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManagerFilterer{contract: contract}, nil
}

// bindShootingCoinManager binds a generic wrapper to an already deployed contract.
func bindShootingCoinManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShootingCoinManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShootingCoinManager *ShootingCoinManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShootingCoinManager.Contract.ShootingCoinManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShootingCoinManager *ShootingCoinManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.ShootingCoinManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShootingCoinManager *ShootingCoinManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.ShootingCoinManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShootingCoinManager *ShootingCoinManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShootingCoinManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShootingCoinManager *ShootingCoinManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShootingCoinManager *ShootingCoinManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.contract.Transact(opts, method, params...)
}

// BetInfo is a free data retrieval call binding the contract method 0xc3b9ed02.
//
// Solidity: function betInfo(address ) view returns(address coinAddress, uint256 betAmount)
func (_ShootingCoinManager *ShootingCoinManagerCaller) BetInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	CoinAddress common.Address
	BetAmount   *big.Int
}, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "betInfo", arg0)

	outstruct := new(struct {
		CoinAddress common.Address
		BetAmount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CoinAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BetAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BetInfo is a free data retrieval call binding the contract method 0xc3b9ed02.
//
// Solidity: function betInfo(address ) view returns(address coinAddress, uint256 betAmount)
func (_ShootingCoinManager *ShootingCoinManagerSession) BetInfo(arg0 common.Address) (struct {
	CoinAddress common.Address
	BetAmount   *big.Int
}, error) {
	return _ShootingCoinManager.Contract.BetInfo(&_ShootingCoinManager.CallOpts, arg0)
}

// BetInfo is a free data retrieval call binding the contract method 0xc3b9ed02.
//
// Solidity: function betInfo(address ) view returns(address coinAddress, uint256 betAmount)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) BetInfo(arg0 common.Address) (struct {
	CoinAddress common.Address
	BetAmount   *big.Int
}, error) {
	return _ShootingCoinManager.Contract.BetInfo(&_ShootingCoinManager.CallOpts, arg0)
}

// CheckOnGame is a free data retrieval call binding the contract method 0x6563c7cd.
//
// Solidity: function checkOnGame(address account) view returns(bool)
func (_ShootingCoinManager *ShootingCoinManagerCaller) CheckOnGame(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "checkOnGame", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOnGame is a free data retrieval call binding the contract method 0x6563c7cd.
//
// Solidity: function checkOnGame(address account) view returns(bool)
func (_ShootingCoinManager *ShootingCoinManagerSession) CheckOnGame(account common.Address) (bool, error) {
	return _ShootingCoinManager.Contract.CheckOnGame(&_ShootingCoinManager.CallOpts, account)
}

// CheckOnGame is a free data retrieval call binding the contract method 0x6563c7cd.
//
// Solidity: function checkOnGame(address account) view returns(bool)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) CheckOnGame(account common.Address) (bool, error) {
	return _ShootingCoinManager.Contract.CheckOnGame(&_ShootingCoinManager.CallOpts, account)
}

// FeeRecieveAddress is a free data retrieval call binding the contract method 0x7999a68c.
//
// Solidity: function feeRecieveAddress() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCaller) FeeRecieveAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "feeRecieveAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecieveAddress is a free data retrieval call binding the contract method 0x7999a68c.
//
// Solidity: function feeRecieveAddress() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerSession) FeeRecieveAddress() (common.Address, error) {
	return _ShootingCoinManager.Contract.FeeRecieveAddress(&_ShootingCoinManager.CallOpts)
}

// FeeRecieveAddress is a free data retrieval call binding the contract method 0x7999a68c.
//
// Solidity: function feeRecieveAddress() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) FeeRecieveAddress() (common.Address, error) {
	return _ShootingCoinManager.Contract.FeeRecieveAddress(&_ShootingCoinManager.CallOpts)
}

// GameHistory is a free data retrieval call binding the contract method 0x3f2ea8bb.
//
// Solidity: function gameHistory(address , uint256 ) view returns(uint256 gameId, address user1, address user1coinAddress, uint256 user1GetAmount, address user2, address user2coinAddress, uint256 user2GetAmount, uint240 timeStamp)
func (_ShootingCoinManager *ShootingCoinManagerCaller) GameHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	GameId           *big.Int
	User1            common.Address
	User1coinAddress common.Address
	User1GetAmount   *big.Int
	User2            common.Address
	User2coinAddress common.Address
	User2GetAmount   *big.Int
	TimeStamp        *big.Int
}, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "gameHistory", arg0, arg1)

	outstruct := new(struct {
		GameId           *big.Int
		User1            common.Address
		User1coinAddress common.Address
		User1GetAmount   *big.Int
		User2            common.Address
		User2coinAddress common.Address
		User2GetAmount   *big.Int
		TimeStamp        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.User1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.User1coinAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.User1GetAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.User2 = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.User2coinAddress = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.User2GetAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TimeStamp = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GameHistory is a free data retrieval call binding the contract method 0x3f2ea8bb.
//
// Solidity: function gameHistory(address , uint256 ) view returns(uint256 gameId, address user1, address user1coinAddress, uint256 user1GetAmount, address user2, address user2coinAddress, uint256 user2GetAmount, uint240 timeStamp)
func (_ShootingCoinManager *ShootingCoinManagerSession) GameHistory(arg0 common.Address, arg1 *big.Int) (struct {
	GameId           *big.Int
	User1            common.Address
	User1coinAddress common.Address
	User1GetAmount   *big.Int
	User2            common.Address
	User2coinAddress common.Address
	User2GetAmount   *big.Int
	TimeStamp        *big.Int
}, error) {
	return _ShootingCoinManager.Contract.GameHistory(&_ShootingCoinManager.CallOpts, arg0, arg1)
}

// GameHistory is a free data retrieval call binding the contract method 0x3f2ea8bb.
//
// Solidity: function gameHistory(address , uint256 ) view returns(uint256 gameId, address user1, address user1coinAddress, uint256 user1GetAmount, address user2, address user2coinAddress, uint256 user2GetAmount, uint240 timeStamp)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) GameHistory(arg0 common.Address, arg1 *big.Int) (struct {
	GameId           *big.Int
	User1            common.Address
	User1coinAddress common.Address
	User1GetAmount   *big.Int
	User2            common.Address
	User2coinAddress common.Address
	User2GetAmount   *big.Int
	TimeStamp        *big.Int
}, error) {
	return _ShootingCoinManager.Contract.GameHistory(&_ShootingCoinManager.CallOpts, arg0, arg1)
}

// GameInfo is a free data retrieval call binding the contract method 0xa6f81668.
//
// Solidity: function gameInfo(uint256 ) view returns(address user1, address user2, (address,uint256,uint256[5]) user1BetInfo, (address,uint256,uint256[5]) user2BetInfo)
func (_ShootingCoinManager *ShootingCoinManagerCaller) GameInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User1        common.Address
	User2        common.Address
	User1BetInfo GameCoreBetInfo
	User2BetInfo GameCoreBetInfo
}, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "gameInfo", arg0)

	outstruct := new(struct {
		User1        common.Address
		User2        common.Address
		User1BetInfo GameCoreBetInfo
		User2BetInfo GameCoreBetInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User1 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.User2 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.User1BetInfo = *abi.ConvertType(out[2], new(GameCoreBetInfo)).(*GameCoreBetInfo)
	outstruct.User2BetInfo = *abi.ConvertType(out[3], new(GameCoreBetInfo)).(*GameCoreBetInfo)

	return *outstruct, err

}

// GameInfo is a free data retrieval call binding the contract method 0xa6f81668.
//
// Solidity: function gameInfo(uint256 ) view returns(address user1, address user2, (address,uint256,uint256[5]) user1BetInfo, (address,uint256,uint256[5]) user2BetInfo)
func (_ShootingCoinManager *ShootingCoinManagerSession) GameInfo(arg0 *big.Int) (struct {
	User1        common.Address
	User2        common.Address
	User1BetInfo GameCoreBetInfo
	User2BetInfo GameCoreBetInfo
}, error) {
	return _ShootingCoinManager.Contract.GameInfo(&_ShootingCoinManager.CallOpts, arg0)
}

// GameInfo is a free data retrieval call binding the contract method 0xa6f81668.
//
// Solidity: function gameInfo(uint256 ) view returns(address user1, address user2, (address,uint256,uint256[5]) user1BetInfo, (address,uint256,uint256[5]) user2BetInfo)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) GameInfo(arg0 *big.Int) (struct {
	User1        common.Address
	User2        common.Address
	User1BetInfo GameCoreBetInfo
	User2BetInfo GameCoreBetInfo
}, error) {
	return _ShootingCoinManager.Contract.GameInfo(&_ShootingCoinManager.CallOpts, arg0)
}

// GetBetInfo is a free data retrieval call binding the contract method 0x32c19466.
//
// Solidity: function getBetInfo(address userAccount) view returns((address,uint256,uint256[5]))
func (_ShootingCoinManager *ShootingCoinManagerCaller) GetBetInfo(opts *bind.CallOpts, userAccount common.Address) (GameCoreBetInfo, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "getBetInfo", userAccount)

	if err != nil {
		return *new(GameCoreBetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GameCoreBetInfo)).(*GameCoreBetInfo)

	return out0, err

}

// GetBetInfo is a free data retrieval call binding the contract method 0x32c19466.
//
// Solidity: function getBetInfo(address userAccount) view returns((address,uint256,uint256[5]))
func (_ShootingCoinManager *ShootingCoinManagerSession) GetBetInfo(userAccount common.Address) (GameCoreBetInfo, error) {
	return _ShootingCoinManager.Contract.GetBetInfo(&_ShootingCoinManager.CallOpts, userAccount)
}

// GetBetInfo is a free data retrieval call binding the contract method 0x32c19466.
//
// Solidity: function getBetInfo(address userAccount) view returns((address,uint256,uint256[5]))
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) GetBetInfo(userAccount common.Address) (GameCoreBetInfo, error) {
	return _ShootingCoinManager.Contract.GetBetInfo(&_ShootingCoinManager.CallOpts, userAccount)
}

// GetGameInfo is a free data retrieval call binding the contract method 0x47e1d550.
//
// Solidity: function getGameInfo(uint256 gameId) view returns((address,address,(address,uint256,uint256[5]),(address,uint256,uint256[5])))
func (_ShootingCoinManager *ShootingCoinManagerCaller) GetGameInfo(opts *bind.CallOpts, gameId *big.Int) (GameCoreGameInfo, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "getGameInfo", gameId)

	if err != nil {
		return *new(GameCoreGameInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GameCoreGameInfo)).(*GameCoreGameInfo)

	return out0, err

}

// GetGameInfo is a free data retrieval call binding the contract method 0x47e1d550.
//
// Solidity: function getGameInfo(uint256 gameId) view returns((address,address,(address,uint256,uint256[5]),(address,uint256,uint256[5])))
func (_ShootingCoinManager *ShootingCoinManagerSession) GetGameInfo(gameId *big.Int) (GameCoreGameInfo, error) {
	return _ShootingCoinManager.Contract.GetGameInfo(&_ShootingCoinManager.CallOpts, gameId)
}

// GetGameInfo is a free data retrieval call binding the contract method 0x47e1d550.
//
// Solidity: function getGameInfo(uint256 gameId) view returns((address,address,(address,uint256,uint256[5]),(address,uint256,uint256[5])))
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) GetGameInfo(gameId *big.Int) (GameCoreGameInfo, error) {
	return _ShootingCoinManager.Contract.GetGameInfo(&_ShootingCoinManager.CallOpts, gameId)
}

// GetHistory is a free data retrieval call binding the contract method 0x80396811.
//
// Solidity: function getHistory(address userAccount) view returns((uint256,address,address,uint256,address,address,uint256,uint240)[])
func (_ShootingCoinManager *ShootingCoinManagerCaller) GetHistory(opts *bind.CallOpts, userAccount common.Address) ([]GameCoreGameHistory, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "getHistory", userAccount)

	if err != nil {
		return *new([]GameCoreGameHistory), err
	}

	out0 := *abi.ConvertType(out[0], new([]GameCoreGameHistory)).(*[]GameCoreGameHistory)

	return out0, err

}

// GetHistory is a free data retrieval call binding the contract method 0x80396811.
//
// Solidity: function getHistory(address userAccount) view returns((uint256,address,address,uint256,address,address,uint256,uint240)[])
func (_ShootingCoinManager *ShootingCoinManagerSession) GetHistory(userAccount common.Address) ([]GameCoreGameHistory, error) {
	return _ShootingCoinManager.Contract.GetHistory(&_ShootingCoinManager.CallOpts, userAccount)
}

// GetHistory is a free data retrieval call binding the contract method 0x80396811.
//
// Solidity: function getHistory(address userAccount) view returns((uint256,address,address,uint256,address,address,uint256,uint240)[])
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) GetHistory(userAccount common.Address) ([]GameCoreGameHistory, error) {
	return _ShootingCoinManager.Contract.GetHistory(&_ShootingCoinManager.CallOpts, userAccount)
}

// GetShootingNft is a free data retrieval call binding the contract method 0x00b95ab2.
//
// Solidity: function getShootingNft() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCaller) GetShootingNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "getShootingNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetShootingNft is a free data retrieval call binding the contract method 0x00b95ab2.
//
// Solidity: function getShootingNft() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerSession) GetShootingNft() (common.Address, error) {
	return _ShootingCoinManager.Contract.GetShootingNft(&_ShootingCoinManager.CallOpts)
}

// GetShootingNft is a free data retrieval call binding the contract method 0x00b95ab2.
//
// Solidity: function getShootingNft() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) GetShootingNft() (common.Address, error) {
	return _ShootingCoinManager.Contract.GetShootingNft(&_ShootingCoinManager.CallOpts)
}

// GetShootingRole is a free data retrieval call binding the contract method 0x48b526e1.
//
// Solidity: function getShootingRole() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCaller) GetShootingRole(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "getShootingRole")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetShootingRole is a free data retrieval call binding the contract method 0x48b526e1.
//
// Solidity: function getShootingRole() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerSession) GetShootingRole() (common.Address, error) {
	return _ShootingCoinManager.Contract.GetShootingRole(&_ShootingCoinManager.CallOpts)
}

// GetShootingRole is a free data retrieval call binding the contract method 0x48b526e1.
//
// Solidity: function getShootingRole() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) GetShootingRole() (common.Address, error) {
	return _ShootingCoinManager.Contract.GetShootingRole(&_ShootingCoinManager.CallOpts)
}

// IsOnGame is a free data retrieval call binding the contract method 0xe2d35f07.
//
// Solidity: function isOnGame(address ) view returns(uint256)
func (_ShootingCoinManager *ShootingCoinManagerCaller) IsOnGame(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "isOnGame", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOnGame is a free data retrieval call binding the contract method 0xe2d35f07.
//
// Solidity: function isOnGame(address ) view returns(uint256)
func (_ShootingCoinManager *ShootingCoinManagerSession) IsOnGame(arg0 common.Address) (*big.Int, error) {
	return _ShootingCoinManager.Contract.IsOnGame(&_ShootingCoinManager.CallOpts, arg0)
}

// IsOnGame is a free data retrieval call binding the contract method 0xe2d35f07.
//
// Solidity: function isOnGame(address ) view returns(uint256)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) IsOnGame(arg0 common.Address) (*big.Int, error) {
	return _ShootingCoinManager.Contract.IsOnGame(&_ShootingCoinManager.CallOpts, arg0)
}

// ShootingNft is a free data retrieval call binding the contract method 0x601a6d89.
//
// Solidity: function shootingNft() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCaller) ShootingNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "shootingNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShootingNft is a free data retrieval call binding the contract method 0x601a6d89.
//
// Solidity: function shootingNft() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerSession) ShootingNft() (common.Address, error) {
	return _ShootingCoinManager.Contract.ShootingNft(&_ShootingCoinManager.CallOpts)
}

// ShootingNft is a free data retrieval call binding the contract method 0x601a6d89.
//
// Solidity: function shootingNft() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) ShootingNft() (common.Address, error) {
	return _ShootingCoinManager.Contract.ShootingNft(&_ShootingCoinManager.CallOpts)
}

// ShootingRole is a free data retrieval call binding the contract method 0xae07a3aa.
//
// Solidity: function shootingRole() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCaller) ShootingRole(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "shootingRole")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShootingRole is a free data retrieval call binding the contract method 0xae07a3aa.
//
// Solidity: function shootingRole() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerSession) ShootingRole() (common.Address, error) {
	return _ShootingCoinManager.Contract.ShootingRole(&_ShootingCoinManager.CallOpts)
}

// ShootingRole is a free data retrieval call binding the contract method 0xae07a3aa.
//
// Solidity: function shootingRole() view returns(address)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) ShootingRole() (common.Address, error) {
	return _ShootingCoinManager.Contract.ShootingRole(&_ShootingCoinManager.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_ShootingCoinManager *ShootingCoinManagerCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ShootingCoinManager.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_ShootingCoinManager *ShootingCoinManagerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _ShootingCoinManager.Contract.Whitelist(&_ShootingCoinManager.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_ShootingCoinManager *ShootingCoinManagerCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _ShootingCoinManager.Contract.Whitelist(&_ShootingCoinManager.CallOpts, arg0)
}

// DespositCoin is a paid mutator transaction binding the contract method 0x229f1db6.
//
// Solidity: function despositCoin(address coinAddress, uint256 amount) payable returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) DespositCoin(opts *bind.TransactOpts, coinAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.Transact(opts, "despositCoin", coinAddress, amount)
}

// DespositCoin is a paid mutator transaction binding the contract method 0x229f1db6.
//
// Solidity: function despositCoin(address coinAddress, uint256 amount) payable returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) DespositCoin(coinAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.DespositCoin(&_ShootingCoinManager.TransactOpts, coinAddress, amount)
}

// DespositCoin is a paid mutator transaction binding the contract method 0x229f1db6.
//
// Solidity: function despositCoin(address coinAddress, uint256 amount) payable returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) DespositCoin(coinAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.DespositCoin(&_ShootingCoinManager.TransactOpts, coinAddress, amount)
}

// EnterGame is a paid mutator transaction binding the contract method 0x11db8595.
//
// Solidity: function enterGame(address account, (address,uint256,uint256[5]) _betInfo) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) EnterGame(opts *bind.TransactOpts, account common.Address, _betInfo GameCoreBetInfo) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.Transact(opts, "enterGame", account, _betInfo)
}

// EnterGame is a paid mutator transaction binding the contract method 0x11db8595.
//
// Solidity: function enterGame(address account, (address,uint256,uint256[5]) _betInfo) returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) EnterGame(account common.Address, _betInfo GameCoreBetInfo) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.EnterGame(&_ShootingCoinManager.TransactOpts, account, _betInfo)
}

// EnterGame is a paid mutator transaction binding the contract method 0x11db8595.
//
// Solidity: function enterGame(address account, (address,uint256,uint256[5]) _betInfo) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) EnterGame(account common.Address, _betInfo GameCoreBetInfo) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.EnterGame(&_ShootingCoinManager.TransactOpts, account, _betInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address roleContract, uint256 _gameFee, address _feeRecieveAddress) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) Initialize(opts *bind.TransactOpts, roleContract common.Address, _gameFee *big.Int, _feeRecieveAddress common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.Transact(opts, "initialize", roleContract, _gameFee, _feeRecieveAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address roleContract, uint256 _gameFee, address _feeRecieveAddress) returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) Initialize(roleContract common.Address, _gameFee *big.Int, _feeRecieveAddress common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.Initialize(&_ShootingCoinManager.TransactOpts, roleContract, _gameFee, _feeRecieveAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc350a1b5.
//
// Solidity: function initialize(address roleContract, uint256 _gameFee, address _feeRecieveAddress) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) Initialize(roleContract common.Address, _gameFee *big.Int, _feeRecieveAddress common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.Initialize(&_ShootingCoinManager.TransactOpts, roleContract, _gameFee, _feeRecieveAddress)
}

// SettleGame is a paid mutator transaction binding the contract method 0x001011ea.
//
// Solidity: function settleGame(uint256 gameId, address user1, address user2, uint256 user1GetAmount, uint256 user2GetAmount) payable returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) SettleGame(opts *bind.TransactOpts, gameId *big.Int, user1 common.Address, user2 common.Address, user1GetAmount *big.Int, user2GetAmount *big.Int) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.Transact(opts, "settleGame", gameId, user1, user2, user1GetAmount, user2GetAmount)
}

// SettleGame is a paid mutator transaction binding the contract method 0x001011ea.
//
// Solidity: function settleGame(uint256 gameId, address user1, address user2, uint256 user1GetAmount, uint256 user2GetAmount) payable returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) SettleGame(gameId *big.Int, user1 common.Address, user2 common.Address, user1GetAmount *big.Int, user2GetAmount *big.Int) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.SettleGame(&_ShootingCoinManager.TransactOpts, gameId, user1, user2, user1GetAmount, user2GetAmount)
}

// SettleGame is a paid mutator transaction binding the contract method 0x001011ea.
//
// Solidity: function settleGame(uint256 gameId, address user1, address user2, uint256 user1GetAmount, uint256 user2GetAmount) payable returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) SettleGame(gameId *big.Int, user1 common.Address, user2 common.Address, user1GetAmount *big.Int, user2GetAmount *big.Int) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.SettleGame(&_ShootingCoinManager.TransactOpts, gameId, user1, user2, user1GetAmount, user2GetAmount)
}

// StartGame is a paid mutator transaction binding the contract method 0xe41d15b0.
//
// Solidity: function startGame(uint256 gameId, address user1, address user2) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) StartGame(opts *bind.TransactOpts, gameId *big.Int, user1 common.Address, user2 common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.Transact(opts, "startGame", gameId, user1, user2)
}

// StartGame is a paid mutator transaction binding the contract method 0xe41d15b0.
//
// Solidity: function startGame(uint256 gameId, address user1, address user2) returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) StartGame(gameId *big.Int, user1 common.Address, user2 common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.StartGame(&_ShootingCoinManager.TransactOpts, gameId, user1, user2)
}

// StartGame is a paid mutator transaction binding the contract method 0xe41d15b0.
//
// Solidity: function startGame(uint256 gameId, address user1, address user2) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) StartGame(gameId *big.Int, user1 common.Address, user2 common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.StartGame(&_ShootingCoinManager.TransactOpts, gameId, user1, user2)
}

// UpdateShootingNft is a paid mutator transaction binding the contract method 0x2f59f6f4.
//
// Solidity: function updateShootingNft(address nftContract) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) UpdateShootingNft(opts *bind.TransactOpts, nftContract common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.Transact(opts, "updateShootingNft", nftContract)
}

// UpdateShootingNft is a paid mutator transaction binding the contract method 0x2f59f6f4.
//
// Solidity: function updateShootingNft(address nftContract) returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) UpdateShootingNft(nftContract common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.UpdateShootingNft(&_ShootingCoinManager.TransactOpts, nftContract)
}

// UpdateShootingNft is a paid mutator transaction binding the contract method 0x2f59f6f4.
//
// Solidity: function updateShootingNft(address nftContract) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) UpdateShootingNft(nftContract common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.UpdateShootingNft(&_ShootingCoinManager.TransactOpts, nftContract)
}

// UpdateShootingRole is a paid mutator transaction binding the contract method 0x660b68ca.
//
// Solidity: function updateShootingRole(address roleContract) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) UpdateShootingRole(opts *bind.TransactOpts, roleContract common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.Transact(opts, "updateShootingRole", roleContract)
}

// UpdateShootingRole is a paid mutator transaction binding the contract method 0x660b68ca.
//
// Solidity: function updateShootingRole(address roleContract) returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) UpdateShootingRole(roleContract common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.UpdateShootingRole(&_ShootingCoinManager.TransactOpts, roleContract)
}

// UpdateShootingRole is a paid mutator transaction binding the contract method 0x660b68ca.
//
// Solidity: function updateShootingRole(address roleContract) returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) UpdateShootingRole(roleContract common.Address) (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.UpdateShootingRole(&_ShootingCoinManager.TransactOpts, roleContract)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShootingCoinManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ShootingCoinManager *ShootingCoinManagerSession) Receive() (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.Receive(&_ShootingCoinManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ShootingCoinManager *ShootingCoinManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _ShootingCoinManager.Contract.Receive(&_ShootingCoinManager.TransactOpts)
}

// ShootingCoinManagerEnteredIterator is returned from FilterEntered and is used to iterate over the raw logs and unpacked data for Entered events raised by the ShootingCoinManager contract.
type ShootingCoinManagerEnteredIterator struct {
	Event *ShootingCoinManagerEntered // Event containing the contract specifics and raw log

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
func (it *ShootingCoinManagerEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShootingCoinManagerEntered)
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
		it.Event = new(ShootingCoinManagerEntered)
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
func (it *ShootingCoinManagerEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShootingCoinManagerEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShootingCoinManagerEntered represents a Entered event raised by the ShootingCoinManager contract.
type ShootingCoinManagerEntered struct {
	User    common.Address
	BetInfo GameCoreBetInfo
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEntered is a free log retrieval operation binding the contract event 0x47b7281b53f7bbfbe15cad82d1b3fb845143f7621aee73273bfd5f18e4437c9e.
//
// Solidity: event Entered(address user, (address,uint256,uint256[5]) betInfo)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) FilterEntered(opts *bind.FilterOpts) (*ShootingCoinManagerEnteredIterator, error) {

	logs, sub, err := _ShootingCoinManager.contract.FilterLogs(opts, "Entered")
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManagerEnteredIterator{contract: _ShootingCoinManager.contract, event: "Entered", logs: logs, sub: sub}, nil
}

// WatchEntered is a free log subscription operation binding the contract event 0x47b7281b53f7bbfbe15cad82d1b3fb845143f7621aee73273bfd5f18e4437c9e.
//
// Solidity: event Entered(address user, (address,uint256,uint256[5]) betInfo)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) WatchEntered(opts *bind.WatchOpts, sink chan<- *ShootingCoinManagerEntered) (event.Subscription, error) {

	logs, sub, err := _ShootingCoinManager.contract.WatchLogs(opts, "Entered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShootingCoinManagerEntered)
				if err := _ShootingCoinManager.contract.UnpackLog(event, "Entered", log); err != nil {
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

// ParseEntered is a log parse operation binding the contract event 0x47b7281b53f7bbfbe15cad82d1b3fb845143f7621aee73273bfd5f18e4437c9e.
//
// Solidity: event Entered(address user, (address,uint256,uint256[5]) betInfo)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) ParseEntered(log types.Log) (*ShootingCoinManagerEntered, error) {
	event := new(ShootingCoinManagerEntered)
	if err := _ShootingCoinManager.contract.UnpackLog(event, "Entered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShootingCoinManagerGameInitedIterator is returned from FilterGameInited and is used to iterate over the raw logs and unpacked data for GameInited events raised by the ShootingCoinManager contract.
type ShootingCoinManagerGameInitedIterator struct {
	Event *ShootingCoinManagerGameInited // Event containing the contract specifics and raw log

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
func (it *ShootingCoinManagerGameInitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShootingCoinManagerGameInited)
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
		it.Event = new(ShootingCoinManagerGameInited)
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
func (it *ShootingCoinManagerGameInitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShootingCoinManagerGameInitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShootingCoinManagerGameInited represents a GameInited event raised by the ShootingCoinManager contract.
type ShootingCoinManagerGameInited struct {
	GameId   *big.Int
	User1    common.Address
	User2    common.Address
	GameInfo GameCoreGameInfo
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGameInited is a free log retrieval operation binding the contract event 0xd48f299fad43f1e5676bfc49e875f38d15ed0a0d99bd81b8e10e319a260c3b49.
//
// Solidity: event GameInited(uint256 indexed gameId, address indexed user1, address indexed user2, (address,address,(address,uint256,uint256[5]),(address,uint256,uint256[5])) gameInfo)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) FilterGameInited(opts *bind.FilterOpts, gameId []*big.Int, user1 []common.Address, user2 []common.Address) (*ShootingCoinManagerGameInitedIterator, error) {

	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}

	logs, sub, err := _ShootingCoinManager.contract.FilterLogs(opts, "GameInited", gameIdRule, user1Rule, user2Rule)
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManagerGameInitedIterator{contract: _ShootingCoinManager.contract, event: "GameInited", logs: logs, sub: sub}, nil
}

// WatchGameInited is a free log subscription operation binding the contract event 0xd48f299fad43f1e5676bfc49e875f38d15ed0a0d99bd81b8e10e319a260c3b49.
//
// Solidity: event GameInited(uint256 indexed gameId, address indexed user1, address indexed user2, (address,address,(address,uint256,uint256[5]),(address,uint256,uint256[5])) gameInfo)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) WatchGameInited(opts *bind.WatchOpts, sink chan<- *ShootingCoinManagerGameInited, gameId []*big.Int, user1 []common.Address, user2 []common.Address) (event.Subscription, error) {

	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}

	logs, sub, err := _ShootingCoinManager.contract.WatchLogs(opts, "GameInited", gameIdRule, user1Rule, user2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShootingCoinManagerGameInited)
				if err := _ShootingCoinManager.contract.UnpackLog(event, "GameInited", log); err != nil {
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

// ParseGameInited is a log parse operation binding the contract event 0xd48f299fad43f1e5676bfc49e875f38d15ed0a0d99bd81b8e10e319a260c3b49.
//
// Solidity: event GameInited(uint256 indexed gameId, address indexed user1, address indexed user2, (address,address,(address,uint256,uint256[5]),(address,uint256,uint256[5])) gameInfo)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) ParseGameInited(log types.Log) (*ShootingCoinManagerGameInited, error) {
	event := new(ShootingCoinManagerGameInited)
	if err := _ShootingCoinManager.contract.UnpackLog(event, "GameInited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShootingCoinManagerGameSettledIterator is returned from FilterGameSettled and is used to iterate over the raw logs and unpacked data for GameSettled events raised by the ShootingCoinManager contract.
type ShootingCoinManagerGameSettledIterator struct {
	Event *ShootingCoinManagerGameSettled // Event containing the contract specifics and raw log

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
func (it *ShootingCoinManagerGameSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShootingCoinManagerGameSettled)
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
		it.Event = new(ShootingCoinManagerGameSettled)
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
func (it *ShootingCoinManagerGameSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShootingCoinManagerGameSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShootingCoinManagerGameSettled represents a GameSettled event raised by the ShootingCoinManager contract.
type ShootingCoinManagerGameSettled struct {
	GameId      *big.Int
	User1       common.Address
	User2       common.Address
	GameHistory GameCoreGameHistory
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGameSettled is a free log retrieval operation binding the contract event 0x8e0f1262e2c81e842cab4abff4244f937fadc8443b9157ec2123403b3cd16eba.
//
// Solidity: event GameSettled(uint256 indexed gameId, address indexed user1, address indexed user2, (uint256,address,address,uint256,address,address,uint256,uint240) gameHistory)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) FilterGameSettled(opts *bind.FilterOpts, gameId []*big.Int, user1 []common.Address, user2 []common.Address) (*ShootingCoinManagerGameSettledIterator, error) {

	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}

	logs, sub, err := _ShootingCoinManager.contract.FilterLogs(opts, "GameSettled", gameIdRule, user1Rule, user2Rule)
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManagerGameSettledIterator{contract: _ShootingCoinManager.contract, event: "GameSettled", logs: logs, sub: sub}, nil
}

// WatchGameSettled is a free log subscription operation binding the contract event 0x8e0f1262e2c81e842cab4abff4244f937fadc8443b9157ec2123403b3cd16eba.
//
// Solidity: event GameSettled(uint256 indexed gameId, address indexed user1, address indexed user2, (uint256,address,address,uint256,address,address,uint256,uint240) gameHistory)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) WatchGameSettled(opts *bind.WatchOpts, sink chan<- *ShootingCoinManagerGameSettled, gameId []*big.Int, user1 []common.Address, user2 []common.Address) (event.Subscription, error) {

	var gameIdRule []interface{}
	for _, gameIdItem := range gameId {
		gameIdRule = append(gameIdRule, gameIdItem)
	}
	var user1Rule []interface{}
	for _, user1Item := range user1 {
		user1Rule = append(user1Rule, user1Item)
	}
	var user2Rule []interface{}
	for _, user2Item := range user2 {
		user2Rule = append(user2Rule, user2Item)
	}

	logs, sub, err := _ShootingCoinManager.contract.WatchLogs(opts, "GameSettled", gameIdRule, user1Rule, user2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShootingCoinManagerGameSettled)
				if err := _ShootingCoinManager.contract.UnpackLog(event, "GameSettled", log); err != nil {
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

// ParseGameSettled is a log parse operation binding the contract event 0x8e0f1262e2c81e842cab4abff4244f937fadc8443b9157ec2123403b3cd16eba.
//
// Solidity: event GameSettled(uint256 indexed gameId, address indexed user1, address indexed user2, (uint256,address,address,uint256,address,address,uint256,uint240) gameHistory)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) ParseGameSettled(log types.Log) (*ShootingCoinManagerGameSettled, error) {
	event := new(ShootingCoinManagerGameSettled)
	if err := _ShootingCoinManager.contract.UnpackLog(event, "GameSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShootingCoinManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ShootingCoinManager contract.
type ShootingCoinManagerInitializedIterator struct {
	Event *ShootingCoinManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ShootingCoinManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShootingCoinManagerInitialized)
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
		it.Event = new(ShootingCoinManagerInitialized)
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
func (it *ShootingCoinManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShootingCoinManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShootingCoinManagerInitialized represents a Initialized event raised by the ShootingCoinManager contract.
type ShootingCoinManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ShootingCoinManagerInitializedIterator, error) {

	logs, sub, err := _ShootingCoinManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ShootingCoinManagerInitializedIterator{contract: _ShootingCoinManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ShootingCoinManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ShootingCoinManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShootingCoinManagerInitialized)
				if err := _ShootingCoinManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ShootingCoinManager *ShootingCoinManagerFilterer) ParseInitialized(log types.Log) (*ShootingCoinManagerInitialized, error) {
	event := new(ShootingCoinManagerInitialized)
	if err := _ShootingCoinManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
