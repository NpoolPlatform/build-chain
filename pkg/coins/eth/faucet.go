package eth

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth/erc20"
	"github.com/ethereum/go-ethereum/common"
)

const (
	Ten        = 10
	EthDecimal = 18
)

func ERC20Faucet(_contract, _to, amount string) (string, error) {
	if ok := common.IsHexAddress(_contract); !ok {
		return "", errors.New("contract address invalid")
	}
	contract := common.HexToAddress(_contract)

	if ok := common.IsHexAddress(_to); !ok {
		return "", errors.New("to address invalid")
	}
	to := common.HexToAddress(_to)

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

	auth, err := GetAuth(client)
	if err != nil {
		return "", err
	}

	deci, err := token.Decimals(nil)
	if err != nil {
		return "", err
	}
	decimal.DivisionPrecision = int(deci)
	_decimal, err := decimal.NewFromString(amount)
	if err != nil {
		return "", err
	}

	_value := _decimal.Mul(decimal.NewFromInt(Ten).Pow(decimal.NewFromInt(int64(deci)))).BigInt()

	tx, err := token.Transfer(auth, to, _value)
	return tx.Hash().String(), err
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

func ETHFaucet(_to, amountStr string) (string, error) {
	if ok := common.IsHexAddress(_to); !ok {
		return "", errors.New("to address invalid")
	}
	to := common.HexToAddress(_to)

	client, err := Client()
	if err != nil {
		return "", err
	}
	defer client.Close()

	cli := ethclient.NewClient(client)
	ctx := context.Background()
	auth, err := GetAuth(client)
	if err != nil {
		return "", err
	}

	var (
		nonce    uint64
		gasPrice *big.Int
		gasLimit uint64 = 21_000
	)

	nonce, err = cli.PendingNonceAt(ctx, auth.From)
	if err != nil {
		return "", err
	}

	gasPrice, err = cli.SuggestGasPrice(ctx)
	if err != nil {
		return "", err
	}

	amount, ok := big.NewFloat(0).SetString(amountStr)
	if !ok {
		return "", fmt.Errorf("parse amount %v failed", amountStr)
	}

	amount.Mul(amount, big.NewFloat(math.Pow10(EthDecimal)))

	amountBig, ok := big.NewInt(0).SetString(amount.Text('f', 0), Ten)
	if !ok {
		return "", errors.New("invalid eth amount")
	}

	if amountBig.Cmp(common.Big0) <= 0 {
		return "", errors.New("invalid eth amount")
	}

	// build tx
	tx := types.NewTransaction(
		nonce,
		to,
		amountBig,
		gasLimit,
		big.NewInt(gasPrice.Int64()),
		nil,
	)

	tx, err = auth.Signer(auth.From, tx)
	if err != nil {
		return "", err
	}

	err = cli.SendTransaction(ctx, tx)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), err
}
