package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	handler "github.com/NpoolPlatform/build-chain/pkg/mw/tokeninfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Faucet(ctx context.Context, in *npool.FaucetRequst) (*npool.FaucetResponse, error) {
	ret := &npool.FaucetResponse{}
	h, err := handler.NewHandler(ctx,
		handler.WithConds(
			&npool.Conds{
				OfficialContract: &v1.StringVal{
					Op:    cruder.EQ,
					Value: in.OfficialContract,
				},
			}),
	)

	info, err := h.GetTokenInfo(ctx)
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
	logger.Sugar().Infof("faucet success, %v,%v", info.Name, ret.Msg)
	return ret, nil
}
