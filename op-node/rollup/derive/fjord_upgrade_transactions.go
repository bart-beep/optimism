package derive

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
)

var (
	// Gas Price Oracle Parameters
	deployFjordGasPriceOracleSource       = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Deployment"}
	GasPriceOracleFjordDeployerAddress    = common.HexToAddress("0x4210000000000000000000000000000000000002")
	gasPriceOracleFjordDeploymentBytecode = common.FromHex("0x608060405234801561001057600080fd5b50611928806100206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80636ef25c3a116100b2578063de26c4a111610081578063f45e65d811610066578063f45e65d81461025b578063f820614014610263578063fe173b971461020d57600080fd5b8063de26c4a114610235578063f1c7a58b1461024857600080fd5b80636ef25c3a1461020d5780638e98b10614610213578063960e3a231461021b578063c59859181461022d57600080fd5b806349948e0e11610109578063519b4bd3116100ee578063519b4bd31461019f57806354fd4d50146101a757806368d5dca6146101f057600080fd5b806349948e0e1461016f5780634ef6e2241461018257600080fd5b80630c18c1621461013b57806322b90ab3146101565780632e0f262514610160578063313ce56714610168575b600080fd5b61014361026b565b6040519081526020015b60405180910390f35b61015e61038c565b005b610143600681565b6006610143565b61014361017d36600461132d565b6105af565b60005461018f9060ff1681565b604051901515815260200161014d565b6101436105ec565b6101e36040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b60405161014d91906113fc565b6101f861064d565b60405163ffffffff909116815260200161014d565b48610143565b61015e6106d2565b60005461018f90610100900460ff1681565b6101f86108d4565b61014361024336600461132d565b610935565b61014361025636600461146f565b6109e9565b610143610ad1565b610143610bc4565b6000805460ff1615610304576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f47617350726963654f7261636c653a206f76657268656164282920697320646560448201527f707265636174656400000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103879190611488565b905090565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040f91906114a1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e2073657420697345636f746f6e6520666c6160648201527f6700000000000000000000000000000000000000000000000000000000000000608482015260a4016102fb565b60005460ff1615610582576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a2045636f746f6e6520616c72656164792060448201527f616374697665000000000000000000000000000000000000000000000000000060648201526084016102fb565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60008054610100900460ff16156105cf576105c982610c25565b92915050565b60005460ff16156105e3576105c982610c4f565b6105c982610cf3565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff166368d5dca66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106ae573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038791906114d7565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa158015610731573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075591906114a1565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461080f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e20736574206973466a6f726420666c61670060648201526084016102fb565b600054610100900460ff16156108a6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f47617350726963654f7261636c653a20466a6f726420616c726561647920616360448201527f746976650000000000000000000000000000000000000000000000000000000060648201526084016102fb565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663c59859186040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106ae573d6000803e3d6000fd5b60008061094183610e47565b60005490915060ff16156109555792915050565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d89190611488565b6109e2908261152c565b9392505050565b60008054610100900460ff16610a81576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f47617350726963654f7261636c653a206765744c314665655570706572426f7560448201527f6e64206f6e6c7920737570706f72747320466a6f72640000000000000000000060648201526084016102fb565b6000610a8e60ff84611573565b610a9890846115db565b610aa39060106115db565b610aae9060446115db565b90506000610abd8460446115db565b9050610ac98282610ed7565b949350505050565b6000805460ff1615610b65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a207363616c61722829206973206465707260448201527f656361746564000000000000000000000000000000000000000000000000000060648201526084016102fb565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663f82061406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b600080610c3183610fe1565b51610c3d90604461152c565b9050600083516044610abd919061152c565b600080610c5b83610e47565b90506000610c676105ec565b610c6f6108d4565b610c7a90601061164f565b63ffffffff16610c8a919061167b565b90506000610c96610bc4565b610c9e61064d565b63ffffffff16610cae919061167b565b90506000610cbc828461152c565b610cc6908561167b565b9050610cd46006600a6117d8565b610cdf90601061167b565b610ce990826117e4565b9695505050505050565b600080610cff83610e47565b9050600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d62573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d869190611488565b610d8e6105ec565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ded573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e119190611488565b610e1b908561152c565b610e25919061167b565b610e2f919061167b565b9050610e3d6006600a6117d8565b610ac990826117e4565b80516000908190815b81811015610eca57848181518110610e6a57610e6a6117f8565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016600003610eaa57610ea360048461152c565b9250610eb8565b610eb560108461152c565b92505b80610ec281611827565b915050610e50565b50610ac98261044061152c565b600080610ee2610bc4565b610eea61064d565b63ffffffff16610efa919061167b565b610f026105ec565b610f0a6108d4565b610f1590601061164f565b63ffffffff16610f25919061167b565b610f2f919061152c565b90506000610f5d847ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffea5a861185f565b610f6a86620fbd2661185f565b610f94907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe5f19de6115db565b610f9e91906115db565b90506000811215610fad575060005b610fb96006600261167b565b610fc490600a6117d8565b610fce838361167b565b610fd891906117e4565b95945050505050565b6060611170565b818153600101919050565b600082840393505b838110156109e25782810151828201511860001a1590930292600101610ffb565b825b60208210611068578251611033601f83610fe8565b52602092909201917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09091019060210161101e565b81156109e257825161107d6001840383610fe8565b520160010192915050565b60006001830392505b61010782106110c9576110bb8360ff166110b660fd6110b68760081c60e00189610fe8565b610fe8565b935061010682039150611091565b600782106110f6576110ef8360ff166110b6600785036110b68760081c60e00189610fe8565b90506109e2565b610ac98360ff166110b68560081c8560051b0187610fe8565b61116882820361114c61113c84600081518060001a8160011a60081b178160021a60101b17915050919050565b639e3779b90260131c611fff1690565b8060021b6040510182815160e01c1860e01b8151188152505050565b600101919050565b6180003860405139618000604051016020830180600d8551820103826002015b818110156112a3576000805b50508051604051600082901a600183901a60081b1760029290921a60101b91909117639e3779b9810260111c617ffc16909101805160e081811c878603811890911b909118909152840190818303908484106111f85750611233565b600184019350611fff821161122d578251600081901a600182901a60081b1760029190911a60101b17810361122d5750611233565b5061119c565b8383106112415750506112a3565b6001830392508583111561125f5761125c878788860361101c565b96505b611273600985016003850160038501610ff3565b9150611280878284611088565b965050611298846112938684860161110f565b61110f565b915050809350611190565b50506112b5838384885185010361101c565b925050506040519150618000820180820391508183526020830160005b838110156112ea5782810151828201526020016112d2565b506000920191825250602001604052919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561133f57600080fd5b813567ffffffffffffffff8082111561135757600080fd5b818401915084601f83011261136b57600080fd5b81358181111561137d5761137d6112fe565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156113c3576113c36112fe565b816040528281528760208487010111156113dc57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156114295785810183015185820160400152820161140d565b8181111561143b576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561148157600080fd5b5035919050565b60006020828403121561149a57600080fd5b5051919050565b6000602082840312156114b357600080fd5b815173ffffffffffffffffffffffffffffffffffffffff811681146109e257600080fd5b6000602082840312156114e957600080fd5b815163ffffffff811681146109e257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561153f5761153f6114fd565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261158257611582611544565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83147f8000000000000000000000000000000000000000000000000000000000000000831416156115d6576115d66114fd565b500590565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03841381151615611615576116156114fd565b827f8000000000000000000000000000000000000000000000000000000000000000038412811615611649576116496114fd565b50500190565b600063ffffffff80831681851681830481118215151615611672576116726114fd565b02949350505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156116b3576116b36114fd565b500290565b600181815b8085111561171157817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156116f7576116f76114fd565b8085161561170457918102915b93841c93908002906116bd565b509250929050565b600082611728575060016105c9565b81611735575060006105c9565b816001811461174b576002811461175557611771565b60019150506105c9565b60ff841115611766576117666114fd565b50506001821b6105c9565b5060208310610133831016604e8410600b8410161715611794575081810a6105c9565b61179e83836116b8565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156117d0576117d06114fd565b029392505050565b60006109e28383611719565b6000826117f3576117f3611544565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611858576118586114fd565b5060010190565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6000841360008413858304851182821616156118a0576118a06114fd565b7f800000000000000000000000000000000000000000000000000000000000000060008712868205881281841616156118db576118db6114fd565b600087129250878205871284841616156118f7576118f76114fd565b8785058712818416161561190d5761190d6114fd565b50505092909302939250505056fea164736f6c634300080f000a")

	// Update GasPricePriceOracle Proxy Parameters
	updateFjordGasPriceOracleSource = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Proxy Update"}
	fjordGasPriceOracleAddress      = common.HexToAddress("0xa919894851548179A0750865e7974DA599C0Fac7")

	// Enable Fjord Parameters
	enableFjordSource = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Set Fjord"}
	enableFjordInput  = crypto.Keccak256([]byte("setFjord()"))[:4]
)

// FjordNetworkUpgradeTransactions returns the transactions required to upgrade the Fjord network.
func FjordNetworkUpgradeTransactions() ([]hexutil.Bytes, error) {
	upgradeTxns := make([]hexutil.Bytes, 0, 3)

	// Deploy Gas Price Oracle transaction
	deployGasPriceOracle, err := types.NewTx(&types.DepositTx{
		SourceHash:          deployFjordGasPriceOracleSource.SourceHash(),
		From:                GasPriceOracleFjordDeployerAddress,
		To:                  nil,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 1_450_000,
		IsSystemTransaction: false,
		Data:                gasPriceOracleFjordDeploymentBytecode,
	}).MarshalBinary()

	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, deployGasPriceOracle)

	updateGasPriceOracleProxy, err := types.NewTx(&types.DepositTx{
		SourceHash:          updateFjordGasPriceOracleSource.SourceHash(),
		From:                common.Address{},
		To:                  &predeploys.GasPriceOracleAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 50_000,
		IsSystemTransaction: false,
		Data:                upgradeToCalldata(fjordGasPriceOracleAddress),
	}).MarshalBinary()

	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, updateGasPriceOracleProxy)

	enableFjord, err := types.NewTx(&types.DepositTx{
		SourceHash:          enableFjordSource.SourceHash(),
		From:                L1InfoDepositerAddress,
		To:                  &predeploys.GasPriceOracleAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 80_000,
		IsSystemTransaction: false,
		Data:                enableFjordInput,
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}
	upgradeTxns = append(upgradeTxns, enableFjord)

	return upgradeTxns, nil
}