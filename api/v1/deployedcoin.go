package api

import (
	"context"
	"time"

	deployedcoin_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/deployedcoin"
	npool "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetDeployedCoins(ctx context.Context, empty *npool.Empty) (*npool.GetDeployedCoinsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	var infos []*npool.DeployedCoinInfo
	var err error

	infos, err = deployedcoin_crud.AllWithCoinName(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail get DeployedCoins: %v", err)
		return &npool.GetDeployedCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDeployedCoinsResponse{
		Infos: infos,
		Total: uint32(len(infos)),
	}, nil
}
