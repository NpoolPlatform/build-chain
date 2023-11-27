package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	handler "github.com/NpoolPlatform/build-chain/pkg/mw/tokeninfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateTokenInfo(ctx context.Context, in *npool.CreateTokenInfoRequest) (*npool.CreateTokenInfoResponse, error) {
	var err error
	var info *npool.TokenInfo

	conds := &npool.Conds{
		OfficialContract: &v1.StringVal{Op: cruder.EQ, Value: *in.Info.OfficialContract},
	}

	h, err := handler.NewHandler(ctx,
		handler.WithConds(conds),
		handler.WithName(in.Info.Name, true),
		handler.WithUnit(in.Info.Unit, true),
		handler.WithDecimal(in.Info.Decimal, true),
		handler.WithChainType(in.Info.ChainType, true),
		handler.WithTokenType(in.Info.TokenType, true),
		handler.WithOfficialContract(in.Info.OfficialContract, true),
		handler.WithPrivateContract(in.Info.PrivateContract, false),
		handler.WithRemark(in.Info.Remark, false),
		handler.WithData(in.Info.Data, true),
	)

	_info, _ := h.GetTokenInfo(ctx)
	if _info != nil && !in.Force {
		logger.Sugar().Infof("create tokeninfo,it is exist, %v", in.Info.Name)
		return &npool.CreateTokenInfoResponse{Info: _info, Success: true, Msg: "it`s exist"}, nil
	}

	contract, err := eth.DeployToken(ctx, in.Info)
	if err != nil {
		logger.Sugar().Errorf("create tokeninfo failed,%v, contract official name: %v", err, in.Info.Name)
		return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
	}

	in.Info.PrivateContract = &contract
	if _info != nil {
		h.ID = &_info.ID
		info, err = h.UpdateTokenInfo(ctx)
		if err != nil {
			logger.Sugar().Errorf("create tokeninfo failed%v, contract official name: %v", err, in.Info.Name)
			return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}
	} else {
		info, err = h.CreateTokenInfo(ctx)
		if err != nil {
			logger.Sugar().Errorf("create tokeninfo failed%v, contract official name: %v", err, in.Info.Name)
			return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}
	}

	msg := fmt.Sprintf("contract official name: %v", in.Info.Name)
	logger.Sugar().Infof("create tokeninfo success, %v", msg)
	return &npool.CreateTokenInfoResponse{
		Info:    info,
		Success: true,
		Msg:     msg,
	}, nil
}

func (s *Server) GetTokenInfos(ctx context.Context, in *npool.GetTokenInfosRequest) (*npool.GetTokenInfosResponse, error) {
	h, err := handler.NewHandler(ctx,
		handler.WithConds(in.Conds),
	)

	infos, total, err := h.GetTokenInfos(ctx)
	if err != nil {
		logger.Sugar().Errorf("get tokeninfos failed, %v", err)
		return &npool.GetTokenInfosResponse{}, status.Error(codes.Unavailable, err.Error())
	}

	logger.Sugar().Infof("get tokeninfos success,total %v", total)
	ret := &npool.GetTokenInfosResponse{Total: uint32(total), Infos: infos}
	return ret, nil
}
