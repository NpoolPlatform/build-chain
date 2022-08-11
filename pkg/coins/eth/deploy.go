package eth

import (
	"context"
	"errors"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	crud "github.com/NpoolPlatform/build-chain/pkg/crud/deployedcoin"
	res "github.com/NpoolPlatform/build-chain/resource"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	proto "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/ethereum/go-ethereum/rpc"
)

var ErrContractNotExist = errors.New("contract data do not exist")

func init() {
	err := coins.Register(coins.EthereumChain, DeployTokens)
	if err != nil {
		panic(err)
	}
}

func DeployTokens(ctx context.Context) error {
	client, err := Client()
	if err != nil {
		return err
	}
	defer client.Close()
	err = UnlockCoinbase(client)
	if err != nil {
		return err
	}
	return deployErc20(ctx, client)
}

func deployErc20(ctx context.Context, client *rpc.Client) error {
	constract := res.GetContract()
	if constract == nil {
		return ErrContractNotExist
	}
	fmt.Println(constract.ERC20Coins[0].ConstuctData)
	successNum := 0
	for _, erc20coin := range constract.ERC20Coins {
		contractAddr, err := DeployContract(client, erc20coin.ConstuctData)
		if err != nil {
			logger.Sugar().Errorf("%v-%v,%v", erc20coin.Type, erc20coin.Name, err)
			continue
		}
		_, err = crud.Create(ctx, &proto.DeployedCoin{
			// Name:     erc20coin.Name,
			// Type:     erc20coin.Type,
			Contract: contractAddr.String(),
		})
		if err != nil {
			logger.Sugar().Errorf("%v-%v,insert to database faild,%v", erc20coin.Type, erc20coin.Name, err)
		}
		successNum++
	}
	logger.Sugar().Infof("%v/%v", successNum, len(constract.ERC20Coins))
	return nil
}
