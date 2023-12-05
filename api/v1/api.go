package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedBuildChainServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterBuildChainServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterBuildChainHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
