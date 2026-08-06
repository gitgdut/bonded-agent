// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// BondedExecutorPlan is an auto generated low-level Go binding around an user-defined struct.
type BondedExecutorPlan struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}

// BondedExecutorMetaData contains all meta data concerning the BondedExecutor contract.
var BondedExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_tUSDC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelPlan\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executePlan\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"calldata_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"lockedBond\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"openPlan\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"plan\",\"type\":\"tuple\",\"internalType\":\"structBondedExecutor.Plan\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"guaranteedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"failureCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"calldataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executed\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"calldata_\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"plans\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"guaranteedOutput\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"failureCompensation\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"calldataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"serviceFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setServiceFee\",\"inputs\":[{\"name\":\"_feeBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tUSDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractMockUSDC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usedNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BondReleased\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanExecuted\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"actualOutput\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paidToUser\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanFailed\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"refundedMON\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"compensationPaid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlanOpened\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"guaranteedOutput\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxCompensation\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ServiceFeePaid\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"swapOutput\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"userReceived\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ServiceFeeUpdated\",\"inputs\":[{\"name\":\"oldFeeBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newFeeBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShortfallPaid\",\"inputs\":[{\"name\":\"planId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"guaranteed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shortfall\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyExecuted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BondTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CalldataMismatch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"actual\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FeeTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBond\",\"inputs\":[{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NonceAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPlanUser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PlanExpired\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"now_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SwapFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051612a0d380380612a0d833981810160405281019061003191906100c9565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506100f4565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100988261006f565b9050919050565b6100a88161008e565b81146100b2575f5ffd5b50565b5f815190506100c38161009f565b92915050565b5f602082840312156100de576100dd61006b565b5b5f6100eb848285016100b5565b91505092915050565b6080516128bb6101525f395f81816104cc0152818161072b01528181610bdf01528181610d1301528181610df001528181610ef501528181611126015281816112b3015281816115bf0152818161174d015261191301526128bb5ff3fe608060405260043610610089575f3560e01c8063bde4592c11610058578063bde4592c14610169578063c708aa4014610191578063dbf4bcab146101bb578063e6d687b9146101f7578063f30e6f7d1461021f57610090565b8063529c5514146100945780635cdf76f8146100be57806394157cc5146100e6578063aa4f26531461012257610090565b3661009057005b5f5ffd5b34801561009f575f5ffd5b506100a861023b565b6040516100b59190611a9d565b60405180910390f35b3480156100c9575f5ffd5b506100e460048036038101906100df9190611ae8565b610240565b005b3480156100f1575f5ffd5b5061010c60048036038101906101079190611b6d565b6102c2565b6040516101199190611bd7565b60405180910390f35b34801561012d575f5ffd5b5061014860048036038101906101439190611c23565b6102f7565b6040516101609c9b9a99989796959493929190611c6c565b60405180910390f35b348015610174575f5ffd5b5061018f600480360381019061018a9190611c23565b6103bc565b005b34801561019c575f5ffd5b506101a56104ca565b6040516101b29190611d7f565b60405180910390f35b3480156101c6575f5ffd5b506101e160048036038101906101dc9190611d98565b6104ee565b6040516101ee9190611a9d565b60405180910390f35b348015610202575f5ffd5b5061021d60048036038101906102189190611e47565b610503565b005b61023960048036038101906102349190611eba565b610a03565b005b5f5481565b606481111561027b576040517fcd4e616700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f549050815f819055507e3b413cf14a67407425bd0b5c065b2de08876554d8489ad7dd4aa95604d280c81836040516102b6929190611f17565b60405180910390a15050565b6002602052825f5260405f20602052815f5260405f20602052805f5260405f205f92509250509054906101000a900460ff1681565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015490806003015490806004015490806005015490806006015490806007015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600801549080600901549080600a01549080600b015f9054906101000a900460ff1690508c565b5f60015f8381526020019081526020015f209050806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610458576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600b015f9054906101000a900460ff16156104a0576040517f0dc1019700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600181600b015f6101000a81548160ff0219169083151502179055506104c68282611893565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6003602052805f5260405f205f915090505481565b60025f845f0160208101906105189190611d98565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8460200160208101906105659190611d98565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f84610140013581526020019081526020015f205f9054906101000a900460ff16156105f6576040517f1fb09b8000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660015f8681526020019081526020015f206001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461068e576040517f0dc1019700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82610100013582826040516106a4929190611f7a565b60405180910390201461070c5782610100013582826040516106c7929190611f7a565b60405180910390206040517f02aed649000000000000000000000000000000000000000000000000000000008152600401610703929190611f92565b60405180910390fd5b5f8360a001359050808460c001351115610728578360c0013590505b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161078693929190611fb9565b6020604051808303815f875af11580156107a2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107c69190612018565b9050806107ff576040517f83e6cc6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8460015f8881526020019081526020015f20818161081d91906123e3565b9050505f60015f8881526020019081526020015f20600b015f6101000a81548160ff021916908315150217905550600160025f875f0160208101906108629190611d98565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8760200160208101906108af9190611d98565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f87610140013581526020019081526020015f205f6101000a81548160ff0219169083151502179055508160035f8760200160208101906109299190611d98565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610970919061241e565b925050819055503373ffffffffffffffffffffffffffffffffffffffff16855f0160208101906109a09190611d98565b73ffffffffffffffffffffffffffffffffffffffff16877f069f545d83672fd15e7535eeca92d2f955aefb37ebdb2d8c50e7fa4dd4b81e808860800135868a61012001356040516109f393929190612451565b60405180910390a4505050505050565b5f60015f8581526020019081526020015f209050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a9e576040517f43edff1b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600b015f9054906101000a900460ff1615610ae6576040517f0dc1019700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060090154421115610b35578060090154426040517f0d937c9c000000000000000000000000000000000000000000000000000000008152600401610b2c929190611f17565b60405180910390fd5b80600801548383604051610b4a929190611f7a565b604051809103902014610bb15780600801548383604051610b6c929190611f7a565b60405180910390206040517f02aed649000000000000000000000000000000000000000000000000000000008152600401610ba8929190611f92565b60405180910390fd5b80600201543414610bc0575f5ffd5b600181600b015f6101000a81548160ff0219169083151502179055505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610c369190612486565b602060405180830381865afa158015610c51573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c7591906124b3565b90505f826007015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1683600201548686604051610cc7929190611f7a565b5f6040518083038185875af1925050503d805f8114610d01576040519150601f19603f3d011682016040523d82523d5f602084013e610d06565b606091505b505090508015611472575f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610d6a9190612486565b602060405180830381865afa158015610d85573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610da991906124b3565b90505f8382610db891906124de565b90505f6127105f5483610dcb9190612511565b610dd5919061257f565b90505f8183610de491906124de565b90505f811115610eeb577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb885f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610e6b9291906125af565b6020604051808303815f875af1158015610e87573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610eab9190612018565b610eea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee190612630565b60405180910390fd5b5b5f821115610ff1577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb886001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401610f719291906125af565b6020604051808303815f875af1158015610f8d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fb19190612018565b610ff0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe790612698565b60405180910390fd5b5b86600401548110611081576110068a88611893565b897fdccbf075436929bdfcd8b760ecd6de8cdad7478265c0800a3cd7fbbaf094ea5d84848460405161103a93929190612451565b60405180910390a2897f0dcff691f0b07d64e21612cd5009ae109a7d58655d4a8b29c4607421b1593ff6845f6040516110749291906126ef565b60405180910390a2611469565b5f81886004015461109291906124de565b90505f886005015482116110a657816110ac565b88600501545b90508060035f8b6001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461111d91906124de565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8a5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016111a19291906125af565b6020604051808303815f875af11580156111bd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111e19190612018565b611220576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161121790612760565b60405180910390fd5b5f818a6005015461123191906124de565b90505f8111156113af578060035f8c6001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112aa91906124de565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8b6001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161132f9291906125af565b6020604051808303815f875af115801561134b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061136f9190612018565b6113ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a5906127c8565b60405180910390fd5b5b8c7f58c01d07a40937a0d19358bc0fb9f5a36464b764dfed12a021bded6c832c059b8b6004015486856040516113e793929190612451565b60405180910390a28c7fdccbf075436929bdfcd8b760ecd6de8cdad7478265c0800a3cd7fbbaf094ea5d87878760405161142393929190612451565b60405180910390a28c7f0dcff691f0b07d64e21612cd5009ae109a7d58655d4a8b29c4607421b1593ff6878460405161145d929190611f17565b60405180910390a25050505b5050505061188b565b5f835f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1684600201546040516114bd90612809565b5f6040518083038185875af1925050503d805f81146114f7576040519150601f19603f3d011682016040523d82523d5f602084013e6114fc565b606091505b5050905080611537576040517f81ceff3000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f846006015490505f8111156116ba578060035f876001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546115b691906124de565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb865f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161163a9291906125af565b6020604051808303815f875af1158015611656573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061167a9190612018565b6116b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116b090612867565b60405180910390fd5b5b5f8186600501546116cb91906124de565b90505f811115611849578060035f886001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461174491906124de565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb876001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016117c99291906125af565b6020604051808303815f875af11580156117e5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118099190612018565b611848576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161183f906127c8565b60405180910390fd5b5b887f522e9517e2a616874167480af7dd8afa4922df6fb13bb4d2b2f2fd1bd3ea97de87600201548460405161187f929190611f17565b60405180910390a25050505b505050505050565b5f816005015490508060035f846001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461190a91906124de565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb836001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161198f9291906125af565b6020604051808303815f875af11580156119ab573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119cf9190612018565b611a0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a05906127c8565b60405180910390fd5b816001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16837f9623c6ff501754597055f382769436c3175bd84a326e54017b43b355dcdffcbc83604051611a789190611a9d565b60405180910390a3505050565b5f819050919050565b611a9781611a85565b82525050565b5f602082019050611ab05f830184611a8e565b92915050565b5f5ffd5b5f5ffd5b611ac781611a85565b8114611ad1575f5ffd5b50565b5f81359050611ae281611abe565b92915050565b5f60208284031215611afd57611afc611ab6565b5b5f611b0a84828501611ad4565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611b3c82611b13565b9050919050565b611b4c81611b32565b8114611b56575f5ffd5b50565b5f81359050611b6781611b43565b92915050565b5f5f5f60608486031215611b8457611b83611ab6565b5b5f611b9186828701611b59565b9350506020611ba286828701611b59565b9250506040611bb386828701611ad4565b9150509250925092565b5f8115159050919050565b611bd181611bbd565b82525050565b5f602082019050611bea5f830184611bc8565b92915050565b5f819050919050565b611c0281611bf0565b8114611c0c575f5ffd5b50565b5f81359050611c1d81611bf9565b92915050565b5f60208284031215611c3857611c37611ab6565b5b5f611c4584828501611c0f565b91505092915050565b611c5781611b32565b82525050565b611c6681611bf0565b82525050565b5f61018082019050611c805f83018f611c4e565b611c8d602083018e611c4e565b611c9a604083018d611a8e565b611ca7606083018c611a8e565b611cb4608083018b611a8e565b611cc160a083018a611a8e565b611cce60c0830189611a8e565b611cdb60e0830188611c4e565b611ce9610100830187611c5d565b611cf7610120830186611a8e565b611d05610140830185611a8e565b611d13610160830184611bc8565b9d9c50505050505050505050505050565b5f819050919050565b5f611d47611d42611d3d84611b13565b611d24565b611b13565b9050919050565b5f611d5882611d2d565b9050919050565b5f611d6982611d4e565b9050919050565b611d7981611d5f565b82525050565b5f602082019050611d925f830184611d70565b92915050565b5f60208284031215611dad57611dac611ab6565b5b5f611dba84828501611b59565b91505092915050565b5f5ffd5b5f6101808284031215611ddd57611ddc611dc3565b5b81905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112611e0757611e06611de6565b5b8235905067ffffffffffffffff811115611e2457611e23611dea565b5b602083019150836001820283011115611e4057611e3f611dee565b5b9250929050565b5f5f5f5f6101c08587031215611e6057611e5f611ab6565b5b5f611e6d87828801611c0f565b9450506020611e7e87828801611dc7565b9350506101a085013567ffffffffffffffff811115611ea057611e9f611aba565b5b611eac87828801611df2565b925092505092959194509250565b5f5f5f60408486031215611ed157611ed0611ab6565b5b5f611ede86828701611c0f565b935050602084013567ffffffffffffffff811115611eff57611efe611aba565b5b611f0b86828701611df2565b92509250509250925092565b5f604082019050611f2a5f830185611a8e565b611f376020830184611a8e565b9392505050565b5f81905092915050565b828183375f83830152505050565b5f611f618385611f3e565b9350611f6e838584611f48565b82840190509392505050565b5f611f86828486611f56565b91508190509392505050565b5f604082019050611fa55f830185611c5d565b611fb26020830184611c5d565b9392505050565b5f606082019050611fcc5f830186611c4e565b611fd96020830185611c4e565b611fe66040830184611a8e565b949350505050565b611ff781611bbd565b8114612001575f5ffd5b50565b5f8151905061201281611fee565b92915050565b5f6020828403121561202d5761202c611ab6565b5b5f61203a84828501612004565b91505092915050565b5f813561204f81611b43565b80915050919050565b5f815f1b9050919050565b5f73ffffffffffffffffffffffffffffffffffffffff61208284612058565b9350801983169250808416831791505092915050565b5f6120a282611d4e565b9050919050565b5f819050919050565b6120bb82612098565b6120ce6120c7826120a9565b8354612063565b8255505050565b5f81356120e181611abe565b80915050919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61211584612058565b9350801983169250808416831791505092915050565b5f61214561214061213b84611a85565b611d24565b611a85565b9050919050565b5f819050919050565b61215e8261212b565b61217161216a8261214c565b83546120ea565b8255505050565b5f813561218481611bf9565b80915050919050565b5f61219782611bf0565b9050919050565b5f815f1c9050919050565b5f6121b38261219e565b9050919050565b6121c38261218d565b6121d66121cf826121a9565b83546120ea565b8255505050565b5f81356121e981611fee565b80915050919050565b5f60ff6121fe84612058565b9350801983169250808416831791505092915050565b5f61221e82611bbd565b9050919050565b5f819050919050565b61223782612214565b61224a61224382612225565b83546121f2565b8255505050565b5f81015f83018061226181612043565b905061226d81846120b2565b50505060018101602083018061228281612043565b905061228e81846120b2565b5050506002810160408301806122a3816120d5565b90506122af8184612155565b5050506003810160608301806122c4816120d5565b90506122d08184612155565b5050506004810160808301806122e5816120d5565b90506122f18184612155565b5050506005810160a0830180612306816120d5565b90506123128184612155565b5050506006810160c0830180612327816120d5565b90506123338184612155565b5050506007810160e083018061234881612043565b905061235481846120b2565b5050506008810161010083018061236a81612178565b905061237681846121ba565b5050506009810161012083018061238c816120d5565b90506123988184612155565b505050600a81016101408301806123ae816120d5565b90506123ba8184612155565b505050600b81016101608301806123d0816121dd565b90506123dc818461222e565b5050505050565b6123ed8282612251565b5050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61242882611a85565b915061243383611a85565b925082820190508082111561244b5761244a6123f1565b5b92915050565b5f6060820190506124645f830186611a8e565b6124716020830185611a8e565b61247e6040830184611a8e565b949350505050565b5f6020820190506124995f830184611c4e565b92915050565b5f815190506124ad81611abe565b92915050565b5f602082840312156124c8576124c7611ab6565b5b5f6124d58482850161249f565b91505092915050565b5f6124e882611a85565b91506124f383611a85565b925082820390508181111561250b5761250a6123f1565b5b92915050565b5f61251b82611a85565b915061252683611a85565b925082820261253481611a85565b9150828204841483151761254b5761254a6123f1565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61258982611a85565b915061259483611a85565b9250826125a4576125a3612552565b5b828204905092915050565b5f6040820190506125c25f830185611c4e565b6125cf6020830184611a8e565b9392505050565b5f82825260208201905092915050565b7f6f7574707574207472616e73666572206661696c6564000000000000000000005f82015250565b5f61261a6016836125d6565b9150612625826125e6565b602082019050919050565b5f6020820190508181035f8301526126478161260e565b9050919050565b7f666565207472616e73666572206661696c6564000000000000000000000000005f82015250565b5f6126826013836125d6565b915061268d8261264e565b602082019050919050565b5f6020820190508181035f8301526126af81612676565b9050919050565b5f819050919050565b5f6126d96126d46126cf846126b6565b611d24565b611a85565b9050919050565b6126e9816126bf565b82525050565b5f6040820190506127025f830185611a8e565b61270f60208301846126e0565b9392505050565b7f636f6d70656e736174696f6e207472616e73666572206661696c6564000000005f82015250565b5f61274a601c836125d6565b915061275582612716565b602082019050919050565b5f6020820190508181035f8301526127778161273e565b9050919050565b7f626f6e642072656c65617365206661696c6564000000000000000000000000005f82015250565b5f6127b26013836125d6565b91506127bd8261277e565b602082019050919050565b5f6020820190508181035f8301526127df816127a6565b9050919050565b50565b5f6127f45f83611f3e565b91506127ff826127e6565b5f82019050919050565b5f612813826127e9565b9150819050919050565b7f6661696c75726520636f6d70656e736174696f6e206661696c656400000000005f82015250565b5f612851601b836125d6565b915061285c8261281d565b602082019050919050565b5f6020820190508181035f83015261287e81612845565b905091905056fea264697066735822122091850a29b9b069b9f5395ce5c1a4210b94c808feef875aadb5e044d0944ee08a64736f6c63430008210033",
}

// BondedExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BondedExecutorMetaData.ABI instead.
var BondedExecutorABI = BondedExecutorMetaData.ABI

// BondedExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BondedExecutorMetaData.Bin instead.
var BondedExecutorBin = BondedExecutorMetaData.Bin

// DeployBondedExecutor deploys a new Ethereum contract, binding an instance of BondedExecutor to it.
func DeployBondedExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, _tUSDC common.Address) (common.Address, *types.Transaction, *BondedExecutor, error) {
	parsed, err := BondedExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BondedExecutorBin), backend, _tUSDC)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BondedExecutor{BondedExecutorCaller: BondedExecutorCaller{contract: contract}, BondedExecutorTransactor: BondedExecutorTransactor{contract: contract}, BondedExecutorFilterer: BondedExecutorFilterer{contract: contract}}, nil
}

// BondedExecutor is an auto generated Go binding around an Ethereum contract.
type BondedExecutor struct {
	BondedExecutorCaller     // Read-only binding to the contract
	BondedExecutorTransactor // Write-only binding to the contract
	BondedExecutorFilterer   // Log filterer for contract events
}

// BondedExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondedExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondedExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondedExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondedExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondedExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondedExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondedExecutorSession struct {
	Contract     *BondedExecutor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondedExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondedExecutorCallerSession struct {
	Contract *BondedExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BondedExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondedExecutorTransactorSession struct {
	Contract     *BondedExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BondedExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondedExecutorRaw struct {
	Contract *BondedExecutor // Generic contract binding to access the raw methods on
}

// BondedExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondedExecutorCallerRaw struct {
	Contract *BondedExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// BondedExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondedExecutorTransactorRaw struct {
	Contract *BondedExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondedExecutor creates a new instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutor(address common.Address, backend bind.ContractBackend) (*BondedExecutor, error) {
	contract, err := bindBondedExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondedExecutor{BondedExecutorCaller: BondedExecutorCaller{contract: contract}, BondedExecutorTransactor: BondedExecutorTransactor{contract: contract}, BondedExecutorFilterer: BondedExecutorFilterer{contract: contract}}, nil
}

// NewBondedExecutorCaller creates a new read-only instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutorCaller(address common.Address, caller bind.ContractCaller) (*BondedExecutorCaller, error) {
	contract, err := bindBondedExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorCaller{contract: contract}, nil
}

// NewBondedExecutorTransactor creates a new write-only instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*BondedExecutorTransactor, error) {
	contract, err := bindBondedExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorTransactor{contract: contract}, nil
}

// NewBondedExecutorFilterer creates a new log filterer instance of BondedExecutor, bound to a specific deployed contract.
func NewBondedExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*BondedExecutorFilterer, error) {
	contract, err := bindBondedExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorFilterer{contract: contract}, nil
}

// bindBondedExecutor binds a generic wrapper to an already deployed contract.
func bindBondedExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BondedExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondedExecutor *BondedExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondedExecutor.Contract.BondedExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondedExecutor *BondedExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondedExecutor.Contract.BondedExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondedExecutor *BondedExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondedExecutor.Contract.BondedExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondedExecutor *BondedExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondedExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondedExecutor *BondedExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondedExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondedExecutor *BondedExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondedExecutor.Contract.contract.Transact(opts, method, params...)
}

// LockedBond is a free data retrieval call binding the contract method 0xdbf4bcab.
//
// Solidity: function lockedBond(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorCaller) LockedBond(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "lockedBond", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBond is a free data retrieval call binding the contract method 0xdbf4bcab.
//
// Solidity: function lockedBond(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorSession) LockedBond(arg0 common.Address) (*big.Int, error) {
	return _BondedExecutor.Contract.LockedBond(&_BondedExecutor.CallOpts, arg0)
}

// LockedBond is a free data retrieval call binding the contract method 0xdbf4bcab.
//
// Solidity: function lockedBond(address ) view returns(uint256)
func (_BondedExecutor *BondedExecutorCallerSession) LockedBond(arg0 common.Address) (*big.Int, error) {
	return _BondedExecutor.Contract.LockedBond(&_BondedExecutor.CallOpts, arg0)
}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(address user, address operator, uint256 inputAmount, uint256 expectedOutput, uint256 guaranteedOutput, uint256 maxCompensation, uint256 failureCompensation, address target, bytes32 calldataHash, uint256 deadline, uint256 nonce, bool executed)
func (_BondedExecutor *BondedExecutorCaller) Plans(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "plans", arg0)

	outstruct := new(struct {
		User                common.Address
		Operator            common.Address
		InputAmount         *big.Int
		ExpectedOutput      *big.Int
		GuaranteedOutput    *big.Int
		MaxCompensation     *big.Int
		FailureCompensation *big.Int
		Target              common.Address
		CalldataHash        [32]byte
		Deadline            *big.Int
		Nonce               *big.Int
		Executed            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.InputAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ExpectedOutput = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.GuaranteedOutput = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxCompensation = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.FailureCompensation = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Target = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.CalldataHash = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.Deadline = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Nonce = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[11], new(bool)).(*bool)

	return *outstruct, err

}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(address user, address operator, uint256 inputAmount, uint256 expectedOutput, uint256 guaranteedOutput, uint256 maxCompensation, uint256 failureCompensation, address target, bytes32 calldataHash, uint256 deadline, uint256 nonce, bool executed)
func (_BondedExecutor *BondedExecutorSession) Plans(arg0 [32]byte) (struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}, error) {
	return _BondedExecutor.Contract.Plans(&_BondedExecutor.CallOpts, arg0)
}

// Plans is a free data retrieval call binding the contract method 0xaa4f2653.
//
// Solidity: function plans(bytes32 ) view returns(address user, address operator, uint256 inputAmount, uint256 expectedOutput, uint256 guaranteedOutput, uint256 maxCompensation, uint256 failureCompensation, address target, bytes32 calldataHash, uint256 deadline, uint256 nonce, bool executed)
func (_BondedExecutor *BondedExecutorCallerSession) Plans(arg0 [32]byte) (struct {
	User                common.Address
	Operator            common.Address
	InputAmount         *big.Int
	ExpectedOutput      *big.Int
	GuaranteedOutput    *big.Int
	MaxCompensation     *big.Int
	FailureCompensation *big.Int
	Target              common.Address
	CalldataHash        [32]byte
	Deadline            *big.Int
	Nonce               *big.Int
	Executed            bool
}, error) {
	return _BondedExecutor.Contract.Plans(&_BondedExecutor.CallOpts, arg0)
}

// ServiceFeeBps is a free data retrieval call binding the contract method 0x529c5514.
//
// Solidity: function serviceFeeBps() view returns(uint256)
func (_BondedExecutor *BondedExecutorCaller) ServiceFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "serviceFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ServiceFeeBps is a free data retrieval call binding the contract method 0x529c5514.
//
// Solidity: function serviceFeeBps() view returns(uint256)
func (_BondedExecutor *BondedExecutorSession) ServiceFeeBps() (*big.Int, error) {
	return _BondedExecutor.Contract.ServiceFeeBps(&_BondedExecutor.CallOpts)
}

// ServiceFeeBps is a free data retrieval call binding the contract method 0x529c5514.
//
// Solidity: function serviceFeeBps() view returns(uint256)
func (_BondedExecutor *BondedExecutorCallerSession) ServiceFeeBps() (*big.Int, error) {
	return _BondedExecutor.Contract.ServiceFeeBps(&_BondedExecutor.CallOpts)
}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_BondedExecutor *BondedExecutorCaller) TUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "tUSDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_BondedExecutor *BondedExecutorSession) TUSDC() (common.Address, error) {
	return _BondedExecutor.Contract.TUSDC(&_BondedExecutor.CallOpts)
}

// TUSDC is a free data retrieval call binding the contract method 0xc708aa40.
//
// Solidity: function tUSDC() view returns(address)
func (_BondedExecutor *BondedExecutorCallerSession) TUSDC() (common.Address, error) {
	return _BondedExecutor.Contract.TUSDC(&_BondedExecutor.CallOpts)
}

// UsedNonces is a free data retrieval call binding the contract method 0x94157cc5.
//
// Solidity: function usedNonces(address , address , uint256 ) view returns(bool)
func (_BondedExecutor *BondedExecutorCaller) UsedNonces(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _BondedExecutor.contract.Call(opts, &out, "usedNonces", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNonces is a free data retrieval call binding the contract method 0x94157cc5.
//
// Solidity: function usedNonces(address , address , uint256 ) view returns(bool)
func (_BondedExecutor *BondedExecutorSession) UsedNonces(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _BondedExecutor.Contract.UsedNonces(&_BondedExecutor.CallOpts, arg0, arg1, arg2)
}

// UsedNonces is a free data retrieval call binding the contract method 0x94157cc5.
//
// Solidity: function usedNonces(address , address , uint256 ) view returns(bool)
func (_BondedExecutor *BondedExecutorCallerSession) UsedNonces(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (bool, error) {
	return _BondedExecutor.Contract.UsedNonces(&_BondedExecutor.CallOpts, arg0, arg1, arg2)
}

// CancelPlan is a paid mutator transaction binding the contract method 0xbde4592c.
//
// Solidity: function cancelPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorTransactor) CancelPlan(opts *bind.TransactOpts, planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "cancelPlan", planId)
}

// CancelPlan is a paid mutator transaction binding the contract method 0xbde4592c.
//
// Solidity: function cancelPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorSession) CancelPlan(planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.CancelPlan(&_BondedExecutor.TransactOpts, planId)
}

// CancelPlan is a paid mutator transaction binding the contract method 0xbde4592c.
//
// Solidity: function cancelPlan(bytes32 planId) returns()
func (_BondedExecutor *BondedExecutorTransactorSession) CancelPlan(planId [32]byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.CancelPlan(&_BondedExecutor.TransactOpts, planId)
}

// ExecutePlan is a paid mutator transaction binding the contract method 0xf30e6f7d.
//
// Solidity: function executePlan(bytes32 planId, bytes calldata_) payable returns()
func (_BondedExecutor *BondedExecutorTransactor) ExecutePlan(opts *bind.TransactOpts, planId [32]byte, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "executePlan", planId, calldata_)
}

// ExecutePlan is a paid mutator transaction binding the contract method 0xf30e6f7d.
//
// Solidity: function executePlan(bytes32 planId, bytes calldata_) payable returns()
func (_BondedExecutor *BondedExecutorSession) ExecutePlan(planId [32]byte, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.ExecutePlan(&_BondedExecutor.TransactOpts, planId, calldata_)
}

// ExecutePlan is a paid mutator transaction binding the contract method 0xf30e6f7d.
//
// Solidity: function executePlan(bytes32 planId, bytes calldata_) payable returns()
func (_BondedExecutor *BondedExecutorTransactorSession) ExecutePlan(planId [32]byte, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.ExecutePlan(&_BondedExecutor.TransactOpts, planId, calldata_)
}

// OpenPlan is a paid mutator transaction binding the contract method 0xe6d687b9.
//
// Solidity: function openPlan(bytes32 planId, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,uint256,uint256,bool) plan, bytes calldata_) returns()
func (_BondedExecutor *BondedExecutorTransactor) OpenPlan(opts *bind.TransactOpts, planId [32]byte, plan BondedExecutorPlan, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "openPlan", planId, plan, calldata_)
}

// OpenPlan is a paid mutator transaction binding the contract method 0xe6d687b9.
//
// Solidity: function openPlan(bytes32 planId, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,uint256,uint256,bool) plan, bytes calldata_) returns()
func (_BondedExecutor *BondedExecutorSession) OpenPlan(planId [32]byte, plan BondedExecutorPlan, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.OpenPlan(&_BondedExecutor.TransactOpts, planId, plan, calldata_)
}

// OpenPlan is a paid mutator transaction binding the contract method 0xe6d687b9.
//
// Solidity: function openPlan(bytes32 planId, (address,address,uint256,uint256,uint256,uint256,uint256,address,bytes32,uint256,uint256,bool) plan, bytes calldata_) returns()
func (_BondedExecutor *BondedExecutorTransactorSession) OpenPlan(planId [32]byte, plan BondedExecutorPlan, calldata_ []byte) (*types.Transaction, error) {
	return _BondedExecutor.Contract.OpenPlan(&_BondedExecutor.TransactOpts, planId, plan, calldata_)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 _feeBps) returns()
func (_BondedExecutor *BondedExecutorTransactor) SetServiceFee(opts *bind.TransactOpts, _feeBps *big.Int) (*types.Transaction, error) {
	return _BondedExecutor.contract.Transact(opts, "setServiceFee", _feeBps)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 _feeBps) returns()
func (_BondedExecutor *BondedExecutorSession) SetServiceFee(_feeBps *big.Int) (*types.Transaction, error) {
	return _BondedExecutor.Contract.SetServiceFee(&_BondedExecutor.TransactOpts, _feeBps)
}

// SetServiceFee is a paid mutator transaction binding the contract method 0x5cdf76f8.
//
// Solidity: function setServiceFee(uint256 _feeBps) returns()
func (_BondedExecutor *BondedExecutorTransactorSession) SetServiceFee(_feeBps *big.Int) (*types.Transaction, error) {
	return _BondedExecutor.Contract.SetServiceFee(&_BondedExecutor.TransactOpts, _feeBps)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BondedExecutor *BondedExecutorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondedExecutor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BondedExecutor *BondedExecutorSession) Receive() (*types.Transaction, error) {
	return _BondedExecutor.Contract.Receive(&_BondedExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BondedExecutor *BondedExecutorTransactorSession) Receive() (*types.Transaction, error) {
	return _BondedExecutor.Contract.Receive(&_BondedExecutor.TransactOpts)
}

// BondedExecutorBondReleasedIterator is returned from FilterBondReleased and is used to iterate over the raw logs and unpacked data for BondReleased events raised by the BondedExecutor contract.
type BondedExecutorBondReleasedIterator struct {
	Event *BondedExecutorBondReleased // Event containing the contract specifics and raw log

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
func (it *BondedExecutorBondReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorBondReleased)
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
		it.Event = new(BondedExecutorBondReleased)
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
func (it *BondedExecutorBondReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorBondReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorBondReleased represents a BondReleased event raised by the BondedExecutor contract.
type BondedExecutorBondReleased struct {
	PlanId   [32]byte
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBondReleased is a free log retrieval operation binding the contract event 0x9623c6ff501754597055f382769436c3175bd84a326e54017b43b355dcdffcbc.
//
// Solidity: event BondReleased(bytes32 indexed planId, address indexed operator, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) FilterBondReleased(opts *bind.FilterOpts, planId [][32]byte, operator []common.Address) (*BondedExecutorBondReleasedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "BondReleased", planIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorBondReleasedIterator{contract: _BondedExecutor.contract, event: "BondReleased", logs: logs, sub: sub}, nil
}

// WatchBondReleased is a free log subscription operation binding the contract event 0x9623c6ff501754597055f382769436c3175bd84a326e54017b43b355dcdffcbc.
//
// Solidity: event BondReleased(bytes32 indexed planId, address indexed operator, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) WatchBondReleased(opts *bind.WatchOpts, sink chan<- *BondedExecutorBondReleased, planId [][32]byte, operator []common.Address) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "BondReleased", planIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorBondReleased)
				if err := _BondedExecutor.contract.UnpackLog(event, "BondReleased", log); err != nil {
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

// ParseBondReleased is a log parse operation binding the contract event 0x9623c6ff501754597055f382769436c3175bd84a326e54017b43b355dcdffcbc.
//
// Solidity: event BondReleased(bytes32 indexed planId, address indexed operator, uint256 amount)
func (_BondedExecutor *BondedExecutorFilterer) ParseBondReleased(log types.Log) (*BondedExecutorBondReleased, error) {
	event := new(BondedExecutorBondReleased)
	if err := _BondedExecutor.contract.UnpackLog(event, "BondReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorPlanExecutedIterator is returned from FilterPlanExecuted and is used to iterate over the raw logs and unpacked data for PlanExecuted events raised by the BondedExecutor contract.
type BondedExecutorPlanExecutedIterator struct {
	Event *BondedExecutorPlanExecuted // Event containing the contract specifics and raw log

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
func (it *BondedExecutorPlanExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorPlanExecuted)
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
		it.Event = new(BondedExecutorPlanExecuted)
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
func (it *BondedExecutorPlanExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorPlanExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorPlanExecuted represents a PlanExecuted event raised by the BondedExecutor contract.
type BondedExecutorPlanExecuted struct {
	PlanId       [32]byte
	ActualOutput *big.Int
	PaidToUser   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPlanExecuted is a free log retrieval operation binding the contract event 0x0dcff691f0b07d64e21612cd5009ae109a7d58655d4a8b29c4607421b1593ff6.
//
// Solidity: event PlanExecuted(bytes32 indexed planId, uint256 actualOutput, uint256 paidToUser)
func (_BondedExecutor *BondedExecutorFilterer) FilterPlanExecuted(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorPlanExecutedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "PlanExecuted", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorPlanExecutedIterator{contract: _BondedExecutor.contract, event: "PlanExecuted", logs: logs, sub: sub}, nil
}

// WatchPlanExecuted is a free log subscription operation binding the contract event 0x0dcff691f0b07d64e21612cd5009ae109a7d58655d4a8b29c4607421b1593ff6.
//
// Solidity: event PlanExecuted(bytes32 indexed planId, uint256 actualOutput, uint256 paidToUser)
func (_BondedExecutor *BondedExecutorFilterer) WatchPlanExecuted(opts *bind.WatchOpts, sink chan<- *BondedExecutorPlanExecuted, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "PlanExecuted", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorPlanExecuted)
				if err := _BondedExecutor.contract.UnpackLog(event, "PlanExecuted", log); err != nil {
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

// ParsePlanExecuted is a log parse operation binding the contract event 0x0dcff691f0b07d64e21612cd5009ae109a7d58655d4a8b29c4607421b1593ff6.
//
// Solidity: event PlanExecuted(bytes32 indexed planId, uint256 actualOutput, uint256 paidToUser)
func (_BondedExecutor *BondedExecutorFilterer) ParsePlanExecuted(log types.Log) (*BondedExecutorPlanExecuted, error) {
	event := new(BondedExecutorPlanExecuted)
	if err := _BondedExecutor.contract.UnpackLog(event, "PlanExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorPlanFailedIterator is returned from FilterPlanFailed and is used to iterate over the raw logs and unpacked data for PlanFailed events raised by the BondedExecutor contract.
type BondedExecutorPlanFailedIterator struct {
	Event *BondedExecutorPlanFailed // Event containing the contract specifics and raw log

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
func (it *BondedExecutorPlanFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorPlanFailed)
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
		it.Event = new(BondedExecutorPlanFailed)
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
func (it *BondedExecutorPlanFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorPlanFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorPlanFailed represents a PlanFailed event raised by the BondedExecutor contract.
type BondedExecutorPlanFailed struct {
	PlanId           [32]byte
	RefundedMON      *big.Int
	CompensationPaid *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPlanFailed is a free log retrieval operation binding the contract event 0x522e9517e2a616874167480af7dd8afa4922df6fb13bb4d2b2f2fd1bd3ea97de.
//
// Solidity: event PlanFailed(bytes32 indexed planId, uint256 refundedMON, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) FilterPlanFailed(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorPlanFailedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "PlanFailed", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorPlanFailedIterator{contract: _BondedExecutor.contract, event: "PlanFailed", logs: logs, sub: sub}, nil
}

// WatchPlanFailed is a free log subscription operation binding the contract event 0x522e9517e2a616874167480af7dd8afa4922df6fb13bb4d2b2f2fd1bd3ea97de.
//
// Solidity: event PlanFailed(bytes32 indexed planId, uint256 refundedMON, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) WatchPlanFailed(opts *bind.WatchOpts, sink chan<- *BondedExecutorPlanFailed, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "PlanFailed", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorPlanFailed)
				if err := _BondedExecutor.contract.UnpackLog(event, "PlanFailed", log); err != nil {
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

// ParsePlanFailed is a log parse operation binding the contract event 0x522e9517e2a616874167480af7dd8afa4922df6fb13bb4d2b2f2fd1bd3ea97de.
//
// Solidity: event PlanFailed(bytes32 indexed planId, uint256 refundedMON, uint256 compensationPaid)
func (_BondedExecutor *BondedExecutorFilterer) ParsePlanFailed(log types.Log) (*BondedExecutorPlanFailed, error) {
	event := new(BondedExecutorPlanFailed)
	if err := _BondedExecutor.contract.UnpackLog(event, "PlanFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorPlanOpenedIterator is returned from FilterPlanOpened and is used to iterate over the raw logs and unpacked data for PlanOpened events raised by the BondedExecutor contract.
type BondedExecutorPlanOpenedIterator struct {
	Event *BondedExecutorPlanOpened // Event containing the contract specifics and raw log

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
func (it *BondedExecutorPlanOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorPlanOpened)
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
		it.Event = new(BondedExecutorPlanOpened)
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
func (it *BondedExecutorPlanOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorPlanOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorPlanOpened represents a PlanOpened event raised by the BondedExecutor contract.
type BondedExecutorPlanOpened struct {
	PlanId           [32]byte
	User             common.Address
	Operator         common.Address
	GuaranteedOutput *big.Int
	MaxCompensation  *big.Int
	Deadline         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPlanOpened is a free log retrieval operation binding the contract event 0x069f545d83672fd15e7535eeca92d2f955aefb37ebdb2d8c50e7fa4dd4b81e80.
//
// Solidity: event PlanOpened(bytes32 indexed planId, address indexed user, address indexed operator, uint256 guaranteedOutput, uint256 maxCompensation, uint256 deadline)
func (_BondedExecutor *BondedExecutorFilterer) FilterPlanOpened(opts *bind.FilterOpts, planId [][32]byte, user []common.Address, operator []common.Address) (*BondedExecutorPlanOpenedIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "PlanOpened", planIdRule, userRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorPlanOpenedIterator{contract: _BondedExecutor.contract, event: "PlanOpened", logs: logs, sub: sub}, nil
}

// WatchPlanOpened is a free log subscription operation binding the contract event 0x069f545d83672fd15e7535eeca92d2f955aefb37ebdb2d8c50e7fa4dd4b81e80.
//
// Solidity: event PlanOpened(bytes32 indexed planId, address indexed user, address indexed operator, uint256 guaranteedOutput, uint256 maxCompensation, uint256 deadline)
func (_BondedExecutor *BondedExecutorFilterer) WatchPlanOpened(opts *bind.WatchOpts, sink chan<- *BondedExecutorPlanOpened, planId [][32]byte, user []common.Address, operator []common.Address) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "PlanOpened", planIdRule, userRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorPlanOpened)
				if err := _BondedExecutor.contract.UnpackLog(event, "PlanOpened", log); err != nil {
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

// ParsePlanOpened is a log parse operation binding the contract event 0x069f545d83672fd15e7535eeca92d2f955aefb37ebdb2d8c50e7fa4dd4b81e80.
//
// Solidity: event PlanOpened(bytes32 indexed planId, address indexed user, address indexed operator, uint256 guaranteedOutput, uint256 maxCompensation, uint256 deadline)
func (_BondedExecutor *BondedExecutorFilterer) ParsePlanOpened(log types.Log) (*BondedExecutorPlanOpened, error) {
	event := new(BondedExecutorPlanOpened)
	if err := _BondedExecutor.contract.UnpackLog(event, "PlanOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorServiceFeePaidIterator is returned from FilterServiceFeePaid and is used to iterate over the raw logs and unpacked data for ServiceFeePaid events raised by the BondedExecutor contract.
type BondedExecutorServiceFeePaidIterator struct {
	Event *BondedExecutorServiceFeePaid // Event containing the contract specifics and raw log

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
func (it *BondedExecutorServiceFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorServiceFeePaid)
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
		it.Event = new(BondedExecutorServiceFeePaid)
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
func (it *BondedExecutorServiceFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorServiceFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorServiceFeePaid represents a ServiceFeePaid event raised by the BondedExecutor contract.
type BondedExecutorServiceFeePaid struct {
	PlanId       [32]byte
	SwapOutput   *big.Int
	Fee          *big.Int
	UserReceived *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterServiceFeePaid is a free log retrieval operation binding the contract event 0xdccbf075436929bdfcd8b760ecd6de8cdad7478265c0800a3cd7fbbaf094ea5d.
//
// Solidity: event ServiceFeePaid(bytes32 indexed planId, uint256 swapOutput, uint256 fee, uint256 userReceived)
func (_BondedExecutor *BondedExecutorFilterer) FilterServiceFeePaid(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorServiceFeePaidIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "ServiceFeePaid", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorServiceFeePaidIterator{contract: _BondedExecutor.contract, event: "ServiceFeePaid", logs: logs, sub: sub}, nil
}

// WatchServiceFeePaid is a free log subscription operation binding the contract event 0xdccbf075436929bdfcd8b760ecd6de8cdad7478265c0800a3cd7fbbaf094ea5d.
//
// Solidity: event ServiceFeePaid(bytes32 indexed planId, uint256 swapOutput, uint256 fee, uint256 userReceived)
func (_BondedExecutor *BondedExecutorFilterer) WatchServiceFeePaid(opts *bind.WatchOpts, sink chan<- *BondedExecutorServiceFeePaid, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "ServiceFeePaid", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorServiceFeePaid)
				if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeePaid", log); err != nil {
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

// ParseServiceFeePaid is a log parse operation binding the contract event 0xdccbf075436929bdfcd8b760ecd6de8cdad7478265c0800a3cd7fbbaf094ea5d.
//
// Solidity: event ServiceFeePaid(bytes32 indexed planId, uint256 swapOutput, uint256 fee, uint256 userReceived)
func (_BondedExecutor *BondedExecutorFilterer) ParseServiceFeePaid(log types.Log) (*BondedExecutorServiceFeePaid, error) {
	event := new(BondedExecutorServiceFeePaid)
	if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorServiceFeeUpdatedIterator is returned from FilterServiceFeeUpdated and is used to iterate over the raw logs and unpacked data for ServiceFeeUpdated events raised by the BondedExecutor contract.
type BondedExecutorServiceFeeUpdatedIterator struct {
	Event *BondedExecutorServiceFeeUpdated // Event containing the contract specifics and raw log

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
func (it *BondedExecutorServiceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorServiceFeeUpdated)
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
		it.Event = new(BondedExecutorServiceFeeUpdated)
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
func (it *BondedExecutorServiceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorServiceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorServiceFeeUpdated represents a ServiceFeeUpdated event raised by the BondedExecutor contract.
type BondedExecutorServiceFeeUpdated struct {
	OldFeeBps *big.Int
	NewFeeBps *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterServiceFeeUpdated is a free log retrieval operation binding the contract event 0x003b413cf14a67407425bd0b5c065b2de08876554d8489ad7dd4aa95604d280c.
//
// Solidity: event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps)
func (_BondedExecutor *BondedExecutorFilterer) FilterServiceFeeUpdated(opts *bind.FilterOpts) (*BondedExecutorServiceFeeUpdatedIterator, error) {

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "ServiceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &BondedExecutorServiceFeeUpdatedIterator{contract: _BondedExecutor.contract, event: "ServiceFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceFeeUpdated is a free log subscription operation binding the contract event 0x003b413cf14a67407425bd0b5c065b2de08876554d8489ad7dd4aa95604d280c.
//
// Solidity: event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps)
func (_BondedExecutor *BondedExecutorFilterer) WatchServiceFeeUpdated(opts *bind.WatchOpts, sink chan<- *BondedExecutorServiceFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "ServiceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorServiceFeeUpdated)
				if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
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

// ParseServiceFeeUpdated is a log parse operation binding the contract event 0x003b413cf14a67407425bd0b5c065b2de08876554d8489ad7dd4aa95604d280c.
//
// Solidity: event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps)
func (_BondedExecutor *BondedExecutorFilterer) ParseServiceFeeUpdated(log types.Log) (*BondedExecutorServiceFeeUpdated, error) {
	event := new(BondedExecutorServiceFeeUpdated)
	if err := _BondedExecutor.contract.UnpackLog(event, "ServiceFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondedExecutorShortfallPaidIterator is returned from FilterShortfallPaid and is used to iterate over the raw logs and unpacked data for ShortfallPaid events raised by the BondedExecutor contract.
type BondedExecutorShortfallPaidIterator struct {
	Event *BondedExecutorShortfallPaid // Event containing the contract specifics and raw log

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
func (it *BondedExecutorShortfallPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondedExecutorShortfallPaid)
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
		it.Event = new(BondedExecutorShortfallPaid)
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
func (it *BondedExecutorShortfallPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondedExecutorShortfallPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondedExecutorShortfallPaid represents a ShortfallPaid event raised by the BondedExecutor contract.
type BondedExecutorShortfallPaid struct {
	PlanId     [32]byte
	Guaranteed *big.Int
	Actual     *big.Int
	Shortfall  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterShortfallPaid is a free log retrieval operation binding the contract event 0x58c01d07a40937a0d19358bc0fb9f5a36464b764dfed12a021bded6c832c059b.
//
// Solidity: event ShortfallPaid(bytes32 indexed planId, uint256 guaranteed, uint256 actual, uint256 shortfall)
func (_BondedExecutor *BondedExecutorFilterer) FilterShortfallPaid(opts *bind.FilterOpts, planId [][32]byte) (*BondedExecutorShortfallPaidIterator, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.FilterLogs(opts, "ShortfallPaid", planIdRule)
	if err != nil {
		return nil, err
	}
	return &BondedExecutorShortfallPaidIterator{contract: _BondedExecutor.contract, event: "ShortfallPaid", logs: logs, sub: sub}, nil
}

// WatchShortfallPaid is a free log subscription operation binding the contract event 0x58c01d07a40937a0d19358bc0fb9f5a36464b764dfed12a021bded6c832c059b.
//
// Solidity: event ShortfallPaid(bytes32 indexed planId, uint256 guaranteed, uint256 actual, uint256 shortfall)
func (_BondedExecutor *BondedExecutorFilterer) WatchShortfallPaid(opts *bind.WatchOpts, sink chan<- *BondedExecutorShortfallPaid, planId [][32]byte) (event.Subscription, error) {

	var planIdRule []interface{}
	for _, planIdItem := range planId {
		planIdRule = append(planIdRule, planIdItem)
	}

	logs, sub, err := _BondedExecutor.contract.WatchLogs(opts, "ShortfallPaid", planIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondedExecutorShortfallPaid)
				if err := _BondedExecutor.contract.UnpackLog(event, "ShortfallPaid", log); err != nil {
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

// ParseShortfallPaid is a log parse operation binding the contract event 0x58c01d07a40937a0d19358bc0fb9f5a36464b764dfed12a021bded6c832c059b.
//
// Solidity: event ShortfallPaid(bytes32 indexed planId, uint256 guaranteed, uint256 actual, uint256 shortfall)
func (_BondedExecutor *BondedExecutorFilterer) ParseShortfallPaid(log types.Log) (*BondedExecutorShortfallPaid, error) {
	event := new(BondedExecutorShortfallPaid)
	if err := _BondedExecutor.contract.UnpackLog(event, "ShortfallPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
