// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// MyERC20TokenStudent is an auto generated low-level Go binding around an user-defined struct.
type MyERC20TokenStudent struct {
	Name string
	Age  *big.Int
}

// MyERC20TokenMetaData contains all meta data concerning the MyERC20Token contract.
var MyERC20TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnTuple\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"returnTupleWithStruct\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"Name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"Age\",\"type\":\"uint256\"}],\"internalType\":\"structMyERC20Token.Student\",\"name\":\"s\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"06fdde03": "name()",
		"39009482": "returnTuple()",
		"e354439b": "returnTupleWithStruct()",
		"95d89b41": "symbol()",
		"a9059cbb": "transfer(address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516107bc3803806107bc83398101604081905261002f916101df565b336000908152600360209081526040822086905584516100529291860190610084565b508051610066906001906020840190610084565b50506002805460ff191660ff92909216919091179055506102a09050565b82805461009090610266565b90600052602060002090601f0160209004810192826100b257600085556100f8565b82601f106100cb57805160ff19168380011785556100f8565b828001600101855582156100f8579182015b828111156100f85782518255916020019190600101906100dd565b50610104929150610108565b5090565b5b808211156101045760008155600101610109565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261014457600080fd5b81516001600160401b038082111561015e5761015e61011d565b604051601f8301601f19908116603f011681019082821181831017156101865761018661011d565b816040528381526020925086838588010111156101a257600080fd5b600091505b838210156101c457858201830151818301840152908201906101a7565b838211156101d55760008385830101525b9695505050505050565b600080600080608085870312156101f557600080fd5b845160208601519094506001600160401b038082111561021457600080fd5b61022088838901610133565b94506040870151915060ff8216821461023857600080fd5b60608701519193508082111561024d57600080fd5b5061025a87828801610133565b91505092959194509250565b600181811c9082168061027a57607f821691505b60208210810361029a57634e487b7160e01b600052602260045260246000fd5b50919050565b61050d806102af6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806370a082311161005b57806370a08231146100d357806395d89b4114610101578063a9059cbb14610109578063e354439b1461011e57600080fd5b806306fdde0314610082578063313ce567146100a057806339009482146100bf575b600080fd5b61008a610176565b60405161009791906103a6565b60405180910390f35b6002546100ad9060ff1681565b60405160ff9091168152602001610097565b604080516001808252602082015201610097565b6100f36100e13660046103dc565b60036020526000908152604090205481565b604051908152602001610097565b61008a610204565b61011c6101173660046103f7565b610211565b005b6101686040805180820190915260608152600060208201525060408051608081018252600691810191825265736f7068696160d01b60608201529081526014602082015290600190565b604051610097929190610421565b6000805461018390610458565b80601f01602080910402602001604051908101604052809291908181526020018280546101af90610458565b80156101fc5780601f106101d1576101008083540402835291602001916101fc565b820191906000526020600020905b8154815290600101906020018083116101df57829003601f168201915b505050505081565b6001805461018390610458565b3360009081526003602052604090205481106102695760405162461bcd60e51b81526020600482015260126024820152710c4c2d8c2dcc6ca40dcdee840cadcdeeaced60731b60448201526064015b60405180910390fd5b6001600160a01b03821660009081526003602052604090205461028c82826104a8565b116102c45760405162461bcd60e51b81526020600482015260086024820152676f766572666c6f7760c01b6044820152606401610260565b33600090815260036020526040812080548392906102e39084906104c0565b90915550506001600160a01b038216600090815260036020526040812080548392906103109084906104a8565b90915550506040518181526001600160a01b0383169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561037f57602081850181015186830182015201610363565b81811115610391576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006103b96020830184610359565b9392505050565b80356001600160a01b03811681146103d757600080fd5b919050565b6000602082840312156103ee57600080fd5b6103b9826103c0565b6000806040838503121561040a57600080fd5b610413836103c0565b946020939093013593505050565b604081526000835160408084015261043c6080840182610359565b6020958601516060850152931515929094019190915250919050565b600181811c9082168061046c57607f821691505b60208210810361048c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b600082198211156104bb576104bb610492565b500190565b6000828210156104d2576104d2610492565b50039056fea2646970667358221220ebb8847b3cd89a3d894562a9e121f7d9902e4c606f145ce549d0c0485cf13b5664736f6c634300080e0033",
}

// MyERC20TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MyERC20TokenMetaData.ABI instead.
var MyERC20TokenABI = MyERC20TokenMetaData.ABI

// Deprecated: Use MyERC20TokenMetaData.Sigs instead.
// MyERC20TokenFuncSigs maps the 4-byte function signature to its string representation.
var MyERC20TokenFuncSigs = MyERC20TokenMetaData.Sigs

// MyERC20TokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyERC20TokenMetaData.Bin instead.
var MyERC20TokenBin = MyERC20TokenMetaData.Bin

// DeployMyERC20Token deploys a new Ethereum contract, binding an instance of MyERC20Token to it.
func DeployMyERC20Token(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (common.Address, *types.Transaction, *MyERC20Token, error) {
	parsed, err := MyERC20TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyERC20TokenBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyERC20Token{MyERC20TokenCaller: MyERC20TokenCaller{contract: contract}, MyERC20TokenTransactor: MyERC20TokenTransactor{contract: contract}, MyERC20TokenFilterer: MyERC20TokenFilterer{contract: contract}}, nil
}

// MyERC20Token is an auto generated Go binding around an Ethereum contract.
type MyERC20Token struct {
	MyERC20TokenCaller     // Read-only binding to the contract
	MyERC20TokenTransactor // Write-only binding to the contract
	MyERC20TokenFilterer   // Log filterer for contract events
}

// MyERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyERC20TokenSession struct {
	Contract     *MyERC20Token     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyERC20TokenCallerSession struct {
	Contract *MyERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MyERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyERC20TokenTransactorSession struct {
	Contract     *MyERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MyERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyERC20TokenRaw struct {
	Contract *MyERC20Token // Generic contract binding to access the raw methods on
}

// MyERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyERC20TokenCallerRaw struct {
	Contract *MyERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyERC20TokenTransactorRaw struct {
	Contract *MyERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyERC20Token creates a new instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20Token(address common.Address, backend bind.ContractBackend) (*MyERC20Token, error) {
	contract, err := bindMyERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyERC20Token{MyERC20TokenCaller: MyERC20TokenCaller{contract: contract}, MyERC20TokenTransactor: MyERC20TokenTransactor{contract: contract}, MyERC20TokenFilterer: MyERC20TokenFilterer{contract: contract}}, nil
}

// NewMyERC20TokenCaller creates a new read-only instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*MyERC20TokenCaller, error) {
	contract, err := bindMyERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenCaller{contract: contract}, nil
}

// NewMyERC20TokenTransactor creates a new write-only instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyERC20TokenTransactor, error) {
	contract, err := bindMyERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenTransactor{contract: contract}, nil
}

// NewMyERC20TokenFilterer creates a new log filterer instance of MyERC20Token, bound to a specific deployed contract.
func NewMyERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MyERC20TokenFilterer, error) {
	contract, err := bindMyERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenFilterer{contract: contract}, nil
}

// bindMyERC20Token binds a generic wrapper to an already deployed contract.
func bindMyERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyERC20TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyERC20Token *MyERC20TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyERC20Token.Contract.MyERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyERC20Token *MyERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyERC20Token.Contract.MyERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyERC20Token *MyERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyERC20Token.Contract.MyERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyERC20Token *MyERC20TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyERC20Token *MyERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyERC20Token *MyERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyERC20Token.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyERC20Token.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyERC20Token.Contract.BalanceOf(&_MyERC20Token.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_MyERC20Token *MyERC20TokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyERC20Token.Contract.BalanceOf(&_MyERC20Token.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyERC20Token *MyERC20TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MyERC20Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyERC20Token *MyERC20TokenSession) Decimals() (uint8, error) {
	return _MyERC20Token.Contract.Decimals(&_MyERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MyERC20Token *MyERC20TokenCallerSession) Decimals() (uint8, error) {
	return _MyERC20Token.Contract.Decimals(&_MyERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyERC20Token *MyERC20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyERC20Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyERC20Token *MyERC20TokenSession) Name() (string, error) {
	return _MyERC20Token.Contract.Name(&_MyERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyERC20Token *MyERC20TokenCallerSession) Name() (string, error) {
	return _MyERC20Token.Contract.Name(&_MyERC20Token.CallOpts)
}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyERC20Token *MyERC20TokenCaller) ReturnTuple(opts *bind.CallOpts) (struct {
	A *big.Int
	B bool
}, error) {
	var out []interface{}
	err := _MyERC20Token.contract.Call(opts, &out, "returnTuple")

	outstruct := new(struct {
		A *big.Int
		B bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.B = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyERC20Token *MyERC20TokenSession) ReturnTuple() (struct {
	A *big.Int
	B bool
}, error) {
	return _MyERC20Token.Contract.ReturnTuple(&_MyERC20Token.CallOpts)
}

// ReturnTuple is a free data retrieval call binding the contract method 0x39009482.
//
// Solidity: function returnTuple() pure returns(uint256 a, bool b)
func (_MyERC20Token *MyERC20TokenCallerSession) ReturnTuple() (struct {
	A *big.Int
	B bool
}, error) {
	return _MyERC20Token.Contract.ReturnTuple(&_MyERC20Token.CallOpts)
}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyERC20Token *MyERC20TokenCaller) ReturnTupleWithStruct(opts *bind.CallOpts) (struct {
	S  MyERC20TokenStudent
	Ok bool
}, error) {
	var out []interface{}
	err := _MyERC20Token.contract.Call(opts, &out, "returnTupleWithStruct")

	outstruct := new(struct {
		S  MyERC20TokenStudent
		Ok bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.S = *abi.ConvertType(out[0], new(MyERC20TokenStudent)).(*MyERC20TokenStudent)
	outstruct.Ok = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyERC20Token *MyERC20TokenSession) ReturnTupleWithStruct() (struct {
	S  MyERC20TokenStudent
	Ok bool
}, error) {
	return _MyERC20Token.Contract.ReturnTupleWithStruct(&_MyERC20Token.CallOpts)
}

// ReturnTupleWithStruct is a free data retrieval call binding the contract method 0xe354439b.
//
// Solidity: function returnTupleWithStruct() pure returns((string,uint256) s, bool ok)
func (_MyERC20Token *MyERC20TokenCallerSession) ReturnTupleWithStruct() (struct {
	S  MyERC20TokenStudent
	Ok bool
}, error) {
	return _MyERC20Token.Contract.ReturnTupleWithStruct(&_MyERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyERC20Token *MyERC20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyERC20Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyERC20Token *MyERC20TokenSession) Symbol() (string, error) {
	return _MyERC20Token.Contract.Symbol(&_MyERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyERC20Token *MyERC20TokenCallerSession) Symbol() (string, error) {
	return _MyERC20Token.Contract.Symbol(&_MyERC20Token.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyERC20Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyERC20Token.Contract.Transfer(&_MyERC20Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_MyERC20Token *MyERC20TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyERC20Token.Contract.Transfer(&_MyERC20Token.TransactOpts, _to, _value)
}

// MyERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyERC20Token contract.
type MyERC20TokenTransferIterator struct {
	Event *MyERC20TokenTransfer // Event containing the contract specifics and raw log

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
func (it *MyERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyERC20TokenTransfer)
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
		it.Event = new(MyERC20TokenTransfer)
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
func (it *MyERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyERC20TokenTransfer represents a Transfer event raised by the MyERC20Token contract.
type MyERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyERC20Token *MyERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyERC20TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyERC20Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyERC20TokenTransferIterator{contract: _MyERC20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyERC20Token *MyERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyERC20TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyERC20Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyERC20TokenTransfer)
				if err := _MyERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MyERC20Token *MyERC20TokenFilterer) ParseTransfer(log types.Log) (*MyERC20TokenTransfer, error) {
	event := new(MyERC20TokenTransfer)
	if err := _MyERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
