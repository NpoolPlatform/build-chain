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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	ErrContractNotExist = errors.New("contract data do not exist")
	ErrTrasferFailed    = errors.New("trasfer failed")
	maxRetries          = 5
)

func DeployToken(ctx context.Context, in *npool.TokenInfo) (string, error) {
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

func DeployBaseErc20(ctx context.Context, client *rpc.Client, in *npool.TokenInfo) (common.Address, error) {
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
	coinbaseAuth, err := GetAuth(client)
	if err != nil {
		return err
	}
	ethClient := ethclient.NewClient(client)
	token, err := erc20.NewErc20(contract, ethClient)
	if err != nil {
		return err
	}

	toPri1, toPub1, err := GenPriAndPubKey()
	if err != nil {
		return err
	}

	_, toPub2, err := GenPriAndPubKey()
	if err != nil {
		return err
	}

	chainID, err := ethClient.NetworkID(ctx)
	if err != nil {
		return err
	}

	auth1, err := bind.NewKeyedTransactorWithChainID(toPri1, chainID)
	if err != nil {
		return err
	}

	// faucet gas for transfer token
	_, err = ETHFaucet(toPub1.String(), "1.888")
	if err != nil {
		return err
	}

	var amount int64 = 10000
	err = transferSpy(token, coinbaseAuth, toPub1, amount)
	if err != nil {
		return err
	}

	amount /= 2
	err = transferSpy(token, auth1, toPub2, amount)
	if err != nil {
		return err
	}

	return nil
}

func transferSpy(token *erc20.Erc20, auth *bind.TransactOpts, toAddr common.Address, amount int64) error {
	balance0, err := token.BalanceOf(nil, toAddr)
	if err != nil {
		return err
	}

	_, err = token.Transfer(auth, toAddr, big.NewInt(amount))
	if err != nil {
		return err
	}

	// wait utile to the account
	var balance1 *big.Int
	for i := 0; i <= maxRetries; i++ {
		balance1, err = token.BalanceOf(nil, toAddr)
		if balance1.Int64() != balance0.Int64() || err != nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		return err
	}
	if balance1.Int64() == balance0.Int64() {
		return ErrTrasferFailed
	}
	return nil
}
