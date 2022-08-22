package client

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type BuildChainClientConn struct {
	Host string
	npool.BuildChainClient
}

func NewClientConn(host string) (*BuildChainClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Sugar().Errorf("fail to dial grpc %v: %v", host, err)
		return nil, err
	}

	connState := conn.GetState()
	if connState != connectivity.Idle && connState != connectivity.Ready {
		logger.Sugar().Warnf("conn not available %v: %v", host, connState)
		return nil, err
	}

	cli := npool.NewBuildChainClient(conn)
	return &BuildChainClientConn{
		Host:             host,
		BuildChainClient: cli,
	}, nil
}
