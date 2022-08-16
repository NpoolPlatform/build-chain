package eth

import (
	"context"
	"encoding/json"
	"errors"

	coinsinfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/coinsinfo"
	deployedcoin_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/deployedcoin"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/deployedcoin"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	proto "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	return DeployBaseErc20(ctx, client, false)
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

func DeployBaseErc20(ctx context.Context, client *rpc.Client, spy bool) error {
	contract := &coins.Contract{}
	conds := cruder.NewConds()
	dCoinConds := cruder.NewConds()
	conds.WithCond(coinsinfo.FieldSimilarity, cruder.EQ, SimBaseERC20)

	infos, _, err := coinsinfo_crud.All(ctx, conds)
	if err != nil {
		return err
	}
	for _, info := range infos {
		if info.Data == nil || len(info.Data) == 0 {
			continue
		}
		dCoinConds.WithCond(deployedcoin.FieldCoinID, cruder.EQ, info.ID)
		dCoins, _, err := deployedcoin_crud.All(ctx, dCoinConds)
		if dCoins != nil && len(dCoins) > 0 {
			for _, coin := range dCoins {
				if ok, _ := hasContractCode(ctx, client, common.HexToAddress(coin.Contract)); ok {
					continue
				}
			}
		}

		err = json.Unmarshal(info.Data, contract)
		if err != nil {
			continue
		}

		contractAddr, err := DeployContract(client, contract.CreateCode)
		if err != nil {
			continue
		}

		_, err = deployedcoin_crud.Create(ctx, &proto.DeployedCoin{
			CoinID:   info.ID,
			Contract: contractAddr.String(),
		})
		if err != nil {
			continue
		}
	}
	return nil
}
