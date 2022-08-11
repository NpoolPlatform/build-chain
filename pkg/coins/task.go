package coins

import (
	"context"
	"errors"

	"github.com/NpoolPlatform/build-chain/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

var (
	DeployTokensHandlers    map[string]handleFunc
	ErrRegisterAlreadyExist = errors.New("handler is already exist,please check")
	ErrHandlerNotExist      = errors.New("handler is not exist")
)

type handleFunc func(ctx context.Context) error

func init() {
	DeployTokensHandlers = make(map[string]handleFunc)
}

func Register(chainType string, fn handleFunc) error {
	if ok := DeployTokensHandlers[chainType]; ok != nil {
		return ErrRegisterAlreadyExist
	}
	DeployTokensHandlers[chainType] = fn
	return nil
}

func Run() {
	chainType := config.GetENV().ChainType
	if ok := DeployTokensHandlers[chainType]; ok == nil {
		logger.Sugar().Error(ErrHandlerNotExist)
		panic(ErrHandlerNotExist)
	}
	fn := DeployTokensHandlers[chainType]
	err := fn(context.Background())
	if err != nil {
		logger.Sugar().Errorf("deploy tokens wrong for %v,%v", chainType, err)
	}
}
