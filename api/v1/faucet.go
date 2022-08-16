package api

import (
	"context"
	"fmt"
	"math/big"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Faucet(ctx context.Context, in *npool.FaucetRequst) (*npool.FaucetResponse, error) {
	ret := &npool.FaucetResponse{}
	// TODO: support muti-chaintype and muti-tokentype
	if ok := common.IsHexAddress(in.Contract); !ok {
		return ret, status.Error(codes.InvalidArgument, "contract address invalid")
	}

	if ok := common.IsHexAddress(in.To); !ok {
		return ret, status.Error(codes.InvalidArgument, "to address invalid")
	}

	amount, ok := big.NewFloat(0).SetString(in.Amount)
	if !ok {
		return ret, status.Error(codes.InvalidArgument, "amount invalid, cannot parse to big.float")
	}

	tx, err := eth.ERC20Faucet(common.HexToAddress(in.Contract), common.HexToAddress(in.To), amount)
	if err != nil {
		return ret, status.Error(codes.InvalidArgument, fmt.Sprintf("faild to air-drop,%v", err))
	}

	ret.Msg = fmt.Sprintf("airdrop tx-id:%v", tx.Hash())
	ret.Success = true
	return ret, nil
}
