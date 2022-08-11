package api

import (
	"context"
	"time"

	crud "github.com/NpoolPlatform/build-chain/pkg/crud/deployedcoin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetDeployedCoins(ctx context.Context, in *npool.GetDeployedCoinsRequest) (*npool.GetDeployedCoinsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	infos, total, err := crud.Rows(ctx, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get DeployedCoins: %v", err)
		return &npool.GetDeployedCoinsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDeployedCoinsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}
