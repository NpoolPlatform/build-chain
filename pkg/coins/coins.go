package coins

// chain name define
const (
	EthereumChain = "ethereum"
	TronChain     = "tron"
	ERC20TOKEN    = "erc20"
)

type Contract struct {
	ABI         string `json:"abi"`
	Code        string `json:"code"`
	Bytecode    string `json:"bytecode"`
	CreateCode  string `json:"create_code"`
	CreateArgs  string `json:"create_args"`
	SwarmSource string `json:"swarm_source"`
}
