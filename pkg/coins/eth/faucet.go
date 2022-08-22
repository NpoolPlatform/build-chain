package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth/erc20"
	"github.com/ethereum/go-ethereum/common"
)

const Ten = 10

func ERC20Faucet(contract, to common.Address, amount string) (*types.Transaction, error) {
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

	err = UnlockCoinbase(client)
	if err != nil {
		return nil, err
	}

	deci, err := token.Decimals(nil)
	if err != nil {
		return nil, err
	}
	decimal.DivisionPrecision = int(deci)
	_decimal, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, err
	}

	_value := _decimal.Mul(decimal.NewFromInt(Ten).Pow(decimal.NewFromInt(int64(deci)))).BigInt()

	tx, err := token.Transfer(auth, to, _value)
	return tx, err
}

func ERC20Balance(contract, acc common.Address) (string, error) {
	client, err := Client()
	if err != nil {
		return "", err
	}
	defer client.Close()

	ethcli := ethclient.NewClient(client)
	token, err := erc20.NewErc20(contract, ethcli)
	if err != nil {
		return "", err
	}

	ret, err := token.BalanceOf(&bind.CallOpts{Pending: true}, acc)
	if err != nil {
		return "", err
	}
	deci, err := token.Decimals(nil)
	if err != nil {
		return "", err
	}
	decimal.DivisionPrecision = int(deci)
	value := decimal.NewFromBigInt(ret, 0).Div(decimal.NewFromInt(Ten).Pow(decimal.NewFromInt(int64(deci)))).String()
	return value, nil
}
