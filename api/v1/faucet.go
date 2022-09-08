package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/tokeninfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Faucet(ctx context.Context, in *npool.FaucetRequst) (*npool.FaucetResponse, error) {
	ret := &npool.FaucetResponse{}

	conds := cruder.NewConds()
	conds.WithCond(tokeninfo.FieldOfficialContract, cruder.EQ, in.OfficialContract)

	info, err := tokeninfo_crud.RowOnly(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("faucet failed, %v", err)
		return ret, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to query tokeninfo,%v", err))
	}

	var txHash string
	switch info.TokenType {
	case coins.ERC20TOKEN:
		txHash, err = eth.ERC20Faucet(info.PrivateContract, in.To, in.Amount)
	case coins.EthereumChain:
		txHash, err = eth.ETHFaucet(in.To, in.Amount)
	}
	if err != nil {
		logger.Sugar().Errorf("faucet failed, %v", err)
		return ret, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to air-drop,%v", err))
	}

	ret.Msg = fmt.Sprintf("airdrop tx-id:%v", txHash)
	ret.Success = true
	logger.Sugar().Infof("faucet success, %v", ret.Msg)
	return ret, nil
}
