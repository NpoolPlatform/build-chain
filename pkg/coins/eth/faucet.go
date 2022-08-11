package eth

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth/erc20"
	"github.com/ethereum/go-ethereum/common"
)

const Ten = 10

func ERC20Faucet(contract, to common.Address, value *big.Float) (*types.Transaction, error) {
	client, err := Client()
	if err != nil {
		return nil, err
	}
	defer client.Close()
	ethcli := ethclient.NewClient(client)
	token, err := erc20.NewErc20(contract, ethcli)

	if err != nil {
		return nil, err
	}

	auth, err := GetAuth(client)
	if err != nil {
		return nil, err
	}

	deci, err := token.Decimals(nil)
	if err != nil {
		return nil, err
	}
	_value := new(big.Int)
	big.NewFloat(0).Mul(value, big.NewFloat(math.Pow10(int(deci)))).Int(_value)

	return token.Transfer(auth, to, _value)
}

func ERC20Balance(contract, acc common.Address) (*big.Int, error) {
	client, err := Client()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ethcli := ethclient.NewClient(client)
	token, err := erc20.NewErc20(contract, ethcli)
	if err != nil {
		return nil, err
	}

	return token.BalanceOf(&bind.CallOpts{Pending: true}, acc)
}
