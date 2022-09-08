package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/tokeninfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func tokenInfoCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
	newConds := cruder.NewConds()

	for k, v := range conds {
		switch v.Op {
		case cruder.EQ:
		case cruder.GT:
		case cruder.LT:
		case cruder.LIKE:
		default:
			return nil, fmt.Errorf("invalid filter condition op")
		}

		switch k {
		case tokeninfo.FieldID:
			fallthrough //nolint
		case tokeninfo.FieldName:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case tokeninfo.FieldChainType:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case tokeninfo.FieldTokenType:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case tokeninfo.FieldUnit:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case tokeninfo.FieldDecimal:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case tokeninfo.FieldOfficialContract:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case tokeninfo.FieldPrivateContract:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case tokeninfo.FieldRemark:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}

	return newConds, nil
}

func (s *Server) CreateTokenInfo(ctx context.Context, in *npool.CreateTokenInfoRequest) (*npool.CreateTokenInfoResponse, error) {
	var err error
	var info *npool.TokenInfo

	if !in.Force {
		conds := cruder.NewConds()
		conds.WithCond(tokeninfo.FieldOfficialContract, cruder.EQ, in.Info.OfficialContract)
		_, num, _ := tokeninfo_crud.Rows(ctx, conds, 0, 0)
		if num != 0 {
			return &npool.CreateTokenInfoResponse{Msg: "it`s exist"}, nil
		}
	}

	contract, err := eth.DeployToken(ctx, in.Info)
	if err != nil {
		return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
	}

	in.Info.PrivateContract = contract
	if in.Force {
		conds := cruder.NewConds()
		conds.WithCond(tokeninfo.FieldOfficialContract, cruder.EQ, in.Info.OfficialContract)
		_info, err := tokeninfo_crud.RowOnly(ctx, conds)
		if err != nil {
			return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}

		in.Info.ID = _info.ID
		info, err = tokeninfo_crud.Update(ctx, in.Info)
		if err != nil {
			return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}
	} else {
		info, err = tokeninfo_crud.Create(ctx, in.Info)
		if err != nil {
			return &npool.CreateTokenInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}
	}

	return &npool.CreateTokenInfoResponse{
		Info:    info,
		Success: true,
		Msg:     fmt.Sprintf("Contract address: %v", contract),
	}, nil
}

func (s *Server) GetTokenInfos(ctx context.Context, in *npool.GetTokenInfosRequest) (*npool.GetTokenInfosResponse, error) {
	conds, err := tokenInfoCondsToConds(in.Conds)
	if err != nil {
		logger.Sugar().Error(err)
		return &npool.GetTokenInfosResponse{}, status.Error(codes.Unavailable, err.Error())
	}

	infos, total, err := tokeninfo_crud.All(ctx, conds)
	if err != nil {
		logger.Sugar().Error(err)
		return &npool.GetTokenInfosResponse{}, status.Error(codes.Unavailable, err.Error())
	}

	ret := &npool.GetTokenInfosResponse{Total: uint32(total), Infos: infos}
	return ret, nil
}
