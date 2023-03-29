// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// MantleMintableERC20FactoryMetaData contains all meta data concerning the MantleMintableERC20Factory contract.
var MantleMintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"MantleMintableERC20Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"}],\"name\":\"StandardL2TokenCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createMantleMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createStandardL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b5060405161243138038061243183398101604081905261003091610050565b6001608081905260a052600060c0526001600160a01b031660e052610080565b60006020828403121561006257600080fd5b81516001600160a01b038116811461007957600080fd5b9392505050565b60805160a05160c05160e0516123726100bf6000396000818160ce015261019e0152600061030d015260006102e2015260006102b701526123726000f3fe60806040523480156200001157600080fd5b5060043610620000525760003560e01c80631758e772146200005757806354fd4d501462000098578063896f93d114620000b1578063ee9a31a214620000c8575b600080fd5b6200006e62000068366004620005ae565b620000f0565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b620000a2620002af565b6040516200008f9190620006c4565b6200006e620000c2366004620005ae565b6200035a565b6200006e7f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff84166200019a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f4d616e746c654d696e7461626c654552433230466163746f72793a206d75737460448201527f2070726f766964652072656d6f746520746f6b656e2061646472657373000000606482015260840160405180910390fd5b60007f0000000000000000000000000000000000000000000000000000000000000000858585604051620001ce90620004be565b620001dd9493929190620006e0565b604051809103906000f080158015620001fa573d6000803e3d6000fd5b5090508073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf60405160405180910390a360405133815273ffffffffffffffffffffffffffffffffffffffff80871691908316907f085f39fc12fc0212e9e5f09958bd8d4d731a3c2a38a3741e7629a2b5dd9d7fbb9060200160405180910390a3949350505050565b6060620002dc7f000000000000000000000000000000000000000000000000000000000000000062000371565b620003077f000000000000000000000000000000000000000000000000000000000000000062000371565b620003327f000000000000000000000000000000000000000000000000000000000000000062000371565b60405160200162000346939291906200073a565b604051602081830303815290604052905090565b600062000369848484620000f0565b949350505050565b606081600003620003b557505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115620003e55780620003cc81620007e5565b9150620003dd9050600a836200084f565b9150620003b9565b60008167ffffffffffffffff811115620004035762000403620004cc565b6040519080825280601f01601f1916602001820160405280156200042e576020820181803683370190505b5090505b841562000369576200044660018362000866565b915062000455600a8662000880565b6200046290603062000897565b60f81b8183815181106200047a576200047a620008b2565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350620004b6600a866200084f565b945062000432565b611a8480620008e283390190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200050d57600080fd5b813567ffffffffffffffff808211156200052b576200052b620004cc565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620005745762000574620004cc565b816040528381528660208588010111156200058e57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215620005c457600080fd5b833573ffffffffffffffffffffffffffffffffffffffff81168114620005e957600080fd5b9250602084013567ffffffffffffffff808211156200060757600080fd5b6200061587838801620004fb565b935060408601359150808211156200062c57600080fd5b506200063b86828701620004fb565b9150509250925092565b60005b838110156200066257818101518382015260200162000648565b8381111562000672576000848401525b50505050565b600081518084526200069281602086016020860162000645565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620006d9602083018462000678565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250608060408301526200071b608083018562000678565b82810360608401526200072f818562000678565b979650505050505050565b600084516200074e81846020890162000645565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516200078c816001850160208a0162000645565b60019201918201528351620007a981600284016020880162000645565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620008195762000819620007b6565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008262000861576200086162000820565b500490565b6000828210156200087b576200087b620007b6565b500390565b60008262000892576200089262000820565b500690565b60008219821115620008ad57620008ad620007b6565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe6101206040523480156200001257600080fd5b5060405162001a8438038062001a8483398101604081905262000035916200016d565b6001600080848460036200004a83826200028c565b5060046200005982826200028c565b50505060809290925260a05260c05250506001600160a01b0390811660e052166101005262000358565b80516001600160a01b03811681146200009b57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000c857600080fd5b81516001600160401b0380821115620000e557620000e5620000a0565b604051601f8301601f19908116603f01168101908282118183101715620001105762000110620000a0565b816040528381526020925086838588010111156200012d57600080fd5b600091505b8382101562000151578582018301518183018401529082019062000132565b83821115620001635760008385830101525b9695505050505050565b600080600080608085870312156200018457600080fd5b6200018f8562000083565b93506200019f6020860162000083565b60408601519093506001600160401b0380821115620001bd57600080fd5b620001cb88838901620000b6565b93506060870151915080821115620001e257600080fd5b50620001f187828801620000b6565b91505092959194509250565b600181811c908216806200021257607f821691505b6020821081036200023357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028757600081815260208120601f850160051c81016020861015620002625750805b601f850160051c820191505b8181101562000283578281556001016200026e565b5050505b505050565b81516001600160401b03811115620002a857620002a8620000a0565b620002c081620002b98454620001fd565b8462000239565b602080601f831160018114620002f85760008415620002df5750858301515b600019600386901b1c1916600185901b17855562000283565b600085815260208120601f198616915b82811015620003295788860151825594840194600190910190840162000308565b5085821015620003485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e051610100516116cb620003b9600039600081816102f50152818161038a015281816105cf01526107a90152600081816101a9015261031b015260006107380152600061070f015260006106e601526116cb6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e1461033f578063e78cea92146102f3578063ee9a31a21461038557600080fd5b8063ae1f6aaf146102f3578063c01e1bd614610319578063d6c0b2c41461031957600080fd5b80639dc29fac116100bd5780639dc29fac146102ba578063a457c2d7146102cd578063a9059cbb146102e057600080fd5b806370a082311461027c57806395d89b41146102b257600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461024c57806340c10f191461025f57806354fd4d501461027457600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004611307565b6103ac565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f861049d565b60405161019b919061137c565b61018f6102133660046113f6565b61052f565b6002545b60405190815260200161019b565b61018f610238366004611420565b610547565b6040516012815260200161019b565b61018f61025a3660046113f6565b61056b565b61027261026d3660046113f6565b6105b7565b005b6101f86106df565b61021c61028a36600461145c565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f8610782565b6102726102c83660046113f6565b610791565b61018f6102db3660046113f6565b6108a8565b61018f6102ee3660046113f6565b610979565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c61034d366004611477565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061046557507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b8061049457507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104ac906114aa565b80601f01602080910402602001604051908101604052809291908181526020018280546104d8906114aa565b80156105255780601f106104fa57610100808354040283529160200191610525565b820191906000526020600020905b81548152906001019060200180831161050857829003601f168201915b5050505050905090565b60003361053d818585610987565b5060019392505050565b600033610555858285610b3b565b610560858585610c12565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061053d90829086906105b290879061152c565b610987565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4d616e746c654d696e7461626c6545524332303a206f6e6c792062726964676560448201527f2063616e206d696e7420616e64206275726e000000000000000000000000000060648201526084015b60405180910390fd5b61068b8282610ec5565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516106d391815260200190565b60405180910390a25050565b606061070a7f0000000000000000000000000000000000000000000000000000000000000000610fe5565b6107337f0000000000000000000000000000000000000000000000000000000000000000610fe5565b61075c7f0000000000000000000000000000000000000000000000000000000000000000610fe5565b60405160200161076e93929190611544565b604051602081830303815290604052905090565b6060600480546104ac906114aa565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610856576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4d616e746c654d696e7461626c6545524332303a206f6e6c792062726964676560448201527f2063616e206d696e7420616e64206275726e00000000000000000000000000006064820152608401610678565b6108608282611122565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516106d391815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561096c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610678565b6105608286868403610987565b60003361053d818585610c12565b73ffffffffffffffffffffffffffffffffffffffff8316610a29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff8216610acc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c0c5781811015610bff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610678565b610c0c8484848403610987565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610cb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff8216610d58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610e0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610e5290849061152c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610eb891815260200190565b60405180910390a3610c0c565b73ffffffffffffffffffffffffffffffffffffffff8216610f42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610678565b8060026000828254610f54919061152c565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610f8e90849061152c565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b60608160000361102857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611052578061103c816115ba565b915061104b9050600a83611621565b915061102c565b60008167ffffffffffffffff81111561106d5761106d611635565b6040519080825280601f01601f191660200182016040528015611097576020820181803683370190505b5090505b841561111a576110ac600183611664565b91506110b9600a8661167b565b6110c490603061152c565b60f81b8183815181106110d9576110d961168f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611113600a86611621565b945061109b565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff82166111c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff82166000908152602081905260409020548181101561127b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604081208383039055600280548492906112b7908490611664565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610b2e565b60006020828403121561131957600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461134957600080fd5b9392505050565b60005b8381101561136b578181015183820152602001611353565b83811115610c0c5750506000910152565b602081526000825180602084015261139b816040850160208701611350565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146113f157600080fd5b919050565b6000806040838503121561140957600080fd5b611412836113cd565b946020939093013593505050565b60008060006060848603121561143557600080fd5b61143e846113cd565b925061144c602085016113cd565b9150604084013590509250925092565b60006020828403121561146e57600080fd5b611349826113cd565b6000806040838503121561148a57600080fd5b611493836113cd565b91506114a1602084016113cd565b90509250929050565b600181811c908216806114be57607f821691505b6020821081036114f7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561153f5761153f6114fd565b500190565b60008451611556818460208901611350565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611592816001850160208a01611350565b600192019182015283516115ad816002840160208801611350565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115eb576115eb6114fd565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082611630576116306115f2565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015611676576116766114fd565b500390565b60008261168a5761168a6115f2565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a",
}

// MantleMintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use MantleMintableERC20FactoryMetaData.ABI instead.
var MantleMintableERC20FactoryABI = MantleMintableERC20FactoryMetaData.ABI

// MantleMintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MantleMintableERC20FactoryMetaData.Bin instead.
var MantleMintableERC20FactoryBin = MantleMintableERC20FactoryMetaData.Bin

// DeployMantleMintableERC20Factory deploys a new Ethereum contract, binding an instance of MantleMintableERC20Factory to it.
func DeployMantleMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *MantleMintableERC20Factory, error) {
	parsed, err := MantleMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MantleMintableERC20FactoryBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MantleMintableERC20Factory{MantleMintableERC20FactoryCaller: MantleMintableERC20FactoryCaller{contract: contract}, MantleMintableERC20FactoryTransactor: MantleMintableERC20FactoryTransactor{contract: contract}, MantleMintableERC20FactoryFilterer: MantleMintableERC20FactoryFilterer{contract: contract}}, nil
}

// MantleMintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type MantleMintableERC20Factory struct {
	MantleMintableERC20FactoryCaller     // Read-only binding to the contract
	MantleMintableERC20FactoryTransactor // Write-only binding to the contract
	MantleMintableERC20FactoryFilterer   // Log filterer for contract events
}

// MantleMintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MantleMintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantleMintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MantleMintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantleMintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MantleMintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MantleMintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MantleMintableERC20FactorySession struct {
	Contract     *MantleMintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MantleMintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MantleMintableERC20FactoryCallerSession struct {
	Contract *MantleMintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// MantleMintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MantleMintableERC20FactoryTransactorSession struct {
	Contract     *MantleMintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// MantleMintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MantleMintableERC20FactoryRaw struct {
	Contract *MantleMintableERC20Factory // Generic contract binding to access the raw methods on
}

// MantleMintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MantleMintableERC20FactoryCallerRaw struct {
	Contract *MantleMintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// MantleMintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MantleMintableERC20FactoryTransactorRaw struct {
	Contract *MantleMintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMantleMintableERC20Factory creates a new instance of MantleMintableERC20Factory, bound to a specific deployed contract.
func NewMantleMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*MantleMintableERC20Factory, error) {
	contract, err := bindMantleMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MantleMintableERC20Factory{MantleMintableERC20FactoryCaller: MantleMintableERC20FactoryCaller{contract: contract}, MantleMintableERC20FactoryTransactor: MantleMintableERC20FactoryTransactor{contract: contract}, MantleMintableERC20FactoryFilterer: MantleMintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewMantleMintableERC20FactoryCaller creates a new read-only instance of MantleMintableERC20Factory, bound to a specific deployed contract.
func NewMantleMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*MantleMintableERC20FactoryCaller, error) {
	contract, err := bindMantleMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MantleMintableERC20FactoryCaller{contract: contract}, nil
}

// NewMantleMintableERC20FactoryTransactor creates a new write-only instance of MantleMintableERC20Factory, bound to a specific deployed contract.
func NewMantleMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*MantleMintableERC20FactoryTransactor, error) {
	contract, err := bindMantleMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MantleMintableERC20FactoryTransactor{contract: contract}, nil
}

// NewMantleMintableERC20FactoryFilterer creates a new log filterer instance of MantleMintableERC20Factory, bound to a specific deployed contract.
func NewMantleMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*MantleMintableERC20FactoryFilterer, error) {
	contract, err := bindMantleMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MantleMintableERC20FactoryFilterer{contract: contract}, nil
}

// bindMantleMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindMantleMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MantleMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantleMintableERC20Factory.Contract.MantleMintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.MantleMintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.MantleMintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MantleMintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MantleMintableERC20Factory.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactorySession) BRIDGE() (common.Address, error) {
	return _MantleMintableERC20Factory.Contract.BRIDGE(&_MantleMintableERC20Factory.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryCallerSession) BRIDGE() (common.Address, error) {
	return _MantleMintableERC20Factory.Contract.BRIDGE(&_MantleMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MantleMintableERC20Factory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MantleMintableERC20Factory *MantleMintableERC20FactorySession) Version() (string, error) {
	return _MantleMintableERC20Factory.Contract.Version(&_MantleMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryCallerSession) Version() (string, error) {
	return _MantleMintableERC20Factory.Contract.Version(&_MantleMintableERC20Factory.CallOpts)
}

// CreateMantleMintableERC20 is a paid mutator transaction binding the contract method 0x1758e772.
//
// Solidity: function createMantleMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryTransactor) CreateMantleMintableERC20(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.contract.Transact(opts, "createMantleMintableERC20", _remoteToken, _name, _symbol)
}

// CreateMantleMintableERC20 is a paid mutator transaction binding the contract method 0x1758e772.
//
// Solidity: function createMantleMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactorySession) CreateMantleMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.CreateMantleMintableERC20(&_MantleMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateMantleMintableERC20 is a paid mutator transaction binding the contract method 0x1758e772.
//
// Solidity: function createMantleMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryTransactorSession) CreateMantleMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.CreateMantleMintableERC20(&_MantleMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryTransactor) CreateStandardL2Token(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.contract.Transact(opts, "createStandardL2Token", _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactorySession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.CreateStandardL2Token(&_MantleMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateStandardL2Token is a paid mutator transaction binding the contract method 0x896f93d1.
//
// Solidity: function createStandardL2Token(address _remoteToken, string _name, string _symbol) returns(address)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryTransactorSession) CreateStandardL2Token(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MantleMintableERC20Factory.Contract.CreateStandardL2Token(&_MantleMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// MantleMintableERC20FactoryMantleMintableERC20CreatedIterator is returned from FilterMantleMintableERC20Created and is used to iterate over the raw logs and unpacked data for MantleMintableERC20Created events raised by the MantleMintableERC20Factory contract.
type MantleMintableERC20FactoryMantleMintableERC20CreatedIterator struct {
	Event *MantleMintableERC20FactoryMantleMintableERC20Created // Event containing the contract specifics and raw log

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
func (it *MantleMintableERC20FactoryMantleMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantleMintableERC20FactoryMantleMintableERC20Created)
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
		it.Event = new(MantleMintableERC20FactoryMantleMintableERC20Created)
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
func (it *MantleMintableERC20FactoryMantleMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantleMintableERC20FactoryMantleMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantleMintableERC20FactoryMantleMintableERC20Created represents a MantleMintableERC20Created event raised by the MantleMintableERC20Factory contract.
type MantleMintableERC20FactoryMantleMintableERC20Created struct {
	LocalToken  common.Address
	RemoteToken common.Address
	Deployer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMantleMintableERC20Created is a free log retrieval operation binding the contract event 0x085f39fc12fc0212e9e5f09958bd8d4d731a3c2a38a3741e7629a2b5dd9d7fbb.
//
// Solidity: event MantleMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryFilterer) FilterMantleMintableERC20Created(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address) (*MantleMintableERC20FactoryMantleMintableERC20CreatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _MantleMintableERC20Factory.contract.FilterLogs(opts, "MantleMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &MantleMintableERC20FactoryMantleMintableERC20CreatedIterator{contract: _MantleMintableERC20Factory.contract, event: "MantleMintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchMantleMintableERC20Created is a free log subscription operation binding the contract event 0x085f39fc12fc0212e9e5f09958bd8d4d731a3c2a38a3741e7629a2b5dd9d7fbb.
//
// Solidity: event MantleMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryFilterer) WatchMantleMintableERC20Created(opts *bind.WatchOpts, sink chan<- *MantleMintableERC20FactoryMantleMintableERC20Created, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _MantleMintableERC20Factory.contract.WatchLogs(opts, "MantleMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantleMintableERC20FactoryMantleMintableERC20Created)
				if err := _MantleMintableERC20Factory.contract.UnpackLog(event, "MantleMintableERC20Created", log); err != nil {
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

// ParseMantleMintableERC20Created is a log parse operation binding the contract event 0x085f39fc12fc0212e9e5f09958bd8d4d731a3c2a38a3741e7629a2b5dd9d7fbb.
//
// Solidity: event MantleMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryFilterer) ParseMantleMintableERC20Created(log types.Log) (*MantleMintableERC20FactoryMantleMintableERC20Created, error) {
	event := new(MantleMintableERC20FactoryMantleMintableERC20Created)
	if err := _MantleMintableERC20Factory.contract.UnpackLog(event, "MantleMintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MantleMintableERC20FactoryStandardL2TokenCreatedIterator is returned from FilterStandardL2TokenCreated and is used to iterate over the raw logs and unpacked data for StandardL2TokenCreated events raised by the MantleMintableERC20Factory contract.
type MantleMintableERC20FactoryStandardL2TokenCreatedIterator struct {
	Event *MantleMintableERC20FactoryStandardL2TokenCreated // Event containing the contract specifics and raw log

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
func (it *MantleMintableERC20FactoryStandardL2TokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MantleMintableERC20FactoryStandardL2TokenCreated)
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
		it.Event = new(MantleMintableERC20FactoryStandardL2TokenCreated)
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
func (it *MantleMintableERC20FactoryStandardL2TokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MantleMintableERC20FactoryStandardL2TokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MantleMintableERC20FactoryStandardL2TokenCreated represents a StandardL2TokenCreated event raised by the MantleMintableERC20Factory contract.
type MantleMintableERC20FactoryStandardL2TokenCreated struct {
	RemoteToken common.Address
	LocalToken  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStandardL2TokenCreated is a free log retrieval operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryFilterer) FilterStandardL2TokenCreated(opts *bind.FilterOpts, remoteToken []common.Address, localToken []common.Address) (*MantleMintableERC20FactoryStandardL2TokenCreatedIterator, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _MantleMintableERC20Factory.contract.FilterLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return &MantleMintableERC20FactoryStandardL2TokenCreatedIterator{contract: _MantleMintableERC20Factory.contract, event: "StandardL2TokenCreated", logs: logs, sub: sub}, nil
}

// WatchStandardL2TokenCreated is a free log subscription operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryFilterer) WatchStandardL2TokenCreated(opts *bind.WatchOpts, sink chan<- *MantleMintableERC20FactoryStandardL2TokenCreated, remoteToken []common.Address, localToken []common.Address) (event.Subscription, error) {

	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}

	logs, sub, err := _MantleMintableERC20Factory.contract.WatchLogs(opts, "StandardL2TokenCreated", remoteTokenRule, localTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MantleMintableERC20FactoryStandardL2TokenCreated)
				if err := _MantleMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
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

// ParseStandardL2TokenCreated is a log parse operation binding the contract event 0xceeb8e7d520d7f3b65fc11a262b91066940193b05d4f93df07cfdced0eb551cf.
//
// Solidity: event StandardL2TokenCreated(address indexed remoteToken, address indexed localToken)
func (_MantleMintableERC20Factory *MantleMintableERC20FactoryFilterer) ParseStandardL2TokenCreated(log types.Log) (*MantleMintableERC20FactoryStandardL2TokenCreated, error) {
	event := new(MantleMintableERC20FactoryStandardL2TokenCreated)
	if err := _MantleMintableERC20Factory.contract.UnpackLog(event, "StandardL2TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
