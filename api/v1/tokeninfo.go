package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func tokenInfoCondsToConds(conds *npool.Conds) *tokeninfo_crud.Conds {
	newConds := &tokeninfo_crud.Conds{}
	if conds == nil {
		return newConds
	}
	if conds.ID != nil {
		newConds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
	}
	if conds.Name != nil {
		newConds.Name = &cruder.Cond{Op: conds.GetName().GetOp(), Val: conds.GetName().GetValue()}
	}
	if conds.ChainType != nil {
		newConds.ChainType = &cruder.Cond{Op: conds.GetChainType().GetOp(), Val: conds.GetChainType().GetValue()}
	}
	if conds.TokenType != nil {
		newConds.TokenType = &cruder.Cond{Op: conds.GetTokenType().GetOp(), Val: conds.GetTokenType().GetValue()}
	}
	if conds.Unit != nil {
		newConds.Unit = &cruder.Cond{Op: conds.GetUnit().GetOp(), Val: conds.GetUnit().GetValue()}
	}
	if conds.Decimal != nil {
		newConds.Decimal = &cruder.Cond{Op: conds.GetDecimal().GetOp(), Val: conds.GetDecimal().GetValue()}
	}
	if conds.OfficialContract != nil {
		newConds.OfficialContract = &cruder.Cond{Op: conds.GetOfficialContract().GetOp(), Val: conds.GetOfficialContract().GetValue()}
	}
	if conds.PrivateContract != nil {
		newConds.PrivateContract = &cruder.Cond{Op: conds.GetPrivateContract().GetOp(), Val: conds.GetPrivateContract().GetValue()}
	}
	if conds.Remark != nil {
		newConds.Remark = &cruder.Cond{Op: conds.GetRemark().GetOp(), Val: conds.GetRemark().GetValue()}
	}
	if conds.Data != nil {
		newConds.Data = &cruder.Cond{Op: conds.GetData().GetOp(), Val: conds.GetData().GetValue()}
	}
	return newConds
}

func (s *Server) CreateTokenInfo(ctx context.Context, in *npool.CreateTokenInfoRequest) (*npool.CreateTokenInfoResponse, error) {
	var err error
	var _info *npool.TokenInfo
	var info *npool.TokenInfo

	conds := &tokeninfo_crud.Conds{
		OfficialContract: &cruder.Cond{Op: cruder.EQ, Val: in.Info.OfficialContract},
	}
	_info, _ = tokeninfo_crud.RowOnly(ctx, conds)
	if _info != nil && !in.Force {
		logger.Sugar().Infof("create tokeninfo,it is exist, %v", in.Info.Name)
		return &npool.CreateTokenInfoResponse{Info: _info, Success: true, Msg: "it`s exist"}, nil
	}

	contract, err := eth.DeployToken(ctx, in.Info)
	if err != nil {
		logger.Sugar().Errorf("create tokeninfo failed,%v, contract official name: %v", err, in.Info.Name)
		return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
	}

	in.Info.PrivateContract = contract
	if _info != nil {
		in.Info.ID = _info.ID
		info, err = tokeninfo_crud.Update(ctx, in.Info)
		if err != nil {
			logger.Sugar().Errorf("create tokeninfo failed%v, contract official name: %v", err, in.Info.Name)
			return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}
	} else {
		info, err = tokeninfo_crud.Create(ctx, in.Info)
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
	conds := tokenInfoCondsToConds(in.Conds)
	infos, total, err := tokeninfo_crud.All(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("get tokeninfos failed, %v", err)
		return &npool.GetTokenInfosResponse{}, status.Error(codes.Unavailable, err.Error())
	}

	logger.Sugar().Infof("get tokeninfos success,total %v", total)
	ret := &npool.GetTokenInfosResponse{Total: uint32(total), Infos: infos}
	return ret, nil
}
