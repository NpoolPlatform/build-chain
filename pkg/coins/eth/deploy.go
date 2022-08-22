package eth

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"time"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/build-chain/pkg/coins/eth/erc20"
	npool "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	ErrContractNotExist       = errors.New("contract data do not exist")
	ErrTrasferFailed          = errors.New("trasfer failed")
	testAmount          int64 = 60
	toAddr                    = common.HexToAddress("0xBcE9e4a7aa5eF6998439618771D4754596045b76")
	maxRetries                = 5
)

func DeployToken(ctx context.Context, in *npool.CoinInfo) (string, error) {
	client, err := Client()
	if err != nil {
		return "", err
	}
	defer client.Close()

	err = UnlockCoinbase(client)
	if err != nil {
		return "", err
	}

	// TODO: support other erc20 token
	contract, err := DeployBaseErc20(ctx, client, in)
	if err != nil {
		return "", err
	}

	for i := 0; i <= maxRetries; i++ {
		ok, err := hasContractCode(ctx, client, contract)
		if ok || err != nil {
			continue
		}
		// to prevent to be ban ip
		time.Sleep(time.Second)
	}
	if err != nil {
		return "", err
	}

	err = TransferSpy(ctx, client, contract)
	if err != nil {
		return "", err
	}

	return contract.String(), nil
}

func hasContractCode(ctx context.Context, client *rpc.Client, contract common.Address) (bool, error) {
	ret, err := ethclient.NewClient(client).CodeAt(ctx, contract, nil)
	if err != nil {
		return false, err
	}
	if len(ret) == 0 {
		return false, nil
	}
	return true, nil
}

func DeployBaseErc20(ctx context.Context, client *rpc.Client, in *npool.CoinInfo) (common.Address, error) {
	contract := &coins.Contract{}
	contractAddr := common.Address{}
	err := json.Unmarshal(in.Data, contract)
	if err != nil {
		return contractAddr, err
	}

	contractAddr, err = DeployContract(client, contract.CreateCode)
	if err != nil {
		return contractAddr, err
	}

	return contractAddr, nil
}

func TransferSpy(ctx context.Context, client *rpc.Client, contract common.Address) error {
	auth, err := GetAuth(client)
	if err != nil {
		return err
	}

	token, err := erc20.NewErc20(contract, ethclient.NewClient(client))
	if err != nil {
		return err
	}

	_, err = token.Transfer(auth, toAddr, big.NewInt(testAmount))
	if err != nil {
		return err
	}

	// wait utile to the account
	var balance *big.Int
	for i := 0; i <= maxRetries; i++ {
		balance, err = token.BalanceOf(nil, toAddr)
		if balance.Int64() != 0 || err != nil {
			continue
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		return err
	}
	if balance.Int64() == 0 {
		return ErrTrasferFailed
	}
	return nil
}
