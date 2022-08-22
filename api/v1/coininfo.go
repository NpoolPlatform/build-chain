package api

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	coininfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/coininfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func coinInfoCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
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
		case coininfo.FieldID:
			fallthrough //nolint
		case coininfo.FieldName:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case coininfo.FieldChainType:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case coininfo.FieldTokenType:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case coininfo.FieldOfficialContract:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case coininfo.FieldPrivateContract:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case coininfo.FieldRemark:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		default:
			return nil, fmt.Errorf("invalid coininfo field")
		}
	}

	return newConds, nil
}

func (s *Server) CreateCoinInfo(ctx context.Context, in *npool.CreateCoinInfoRequest) (*npool.CreateCoinInfoResponse, error) {
	var err error
	info := &npool.CoinInfo{}

	if !in.Force {
		conds := cruder.NewConds()
		conds.WithCond(coininfo.FieldOfficialContract, cruder.EQ, in.Info.OfficialContract)
		_, num, _ := coininfo_crud.Rows(ctx, conds, 0, 0)
		if num != 0 {
			return &npool.CreateCoinInfoResponse{Msg: "it`s exist"}, nil
		}
	}

	contract, err := eth.DeployToken(ctx, in.Info)
	if err != nil {
		return &npool.CreateCoinInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
	}

	in.Info.PrivateContract = contract
	if in.Force {
		conds := cruder.NewConds()
		conds.WithCond(coininfo.FieldOfficialContract, cruder.EQ, in.Info.OfficialContract)
		infos, num, err := coininfo_crud.Rows(ctx, conds, 0, 0)
		if err != nil {
			return &npool.CreateCoinInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}

		if num == 0 {
			info, err = coininfo_crud.Create(ctx, in.Info)
		} else {
			in.Info.ID = infos[0].ID
			info, err = coininfo_crud.Update(ctx, in.Info)
		}

		if err != nil {
			return &npool.CreateCoinInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}
	} else {
		info, err = coininfo_crud.Create(ctx, in.Info)
		if err != nil {
			return &npool.CreateCoinInfoResponse{Msg: err.Error()}, status.Error(codes.Internal, err.Error())
		}
	}

	return &npool.CreateCoinInfoResponse{
		Success: true,
		Msg:     fmt.Sprintf("Contract address: %v", info.PrivateContract),
	}, nil
}

func (s *Server) GetCoinInfos(ctx context.Context, in *npool.GetCoinInfosRequest) (*npool.GetCoinInfosResponse, error) {
	conds, err := coinInfoCondsToConds(in.Conds)
	if err != nil {
		logger.Sugar().Error(err)
		return &npool.GetCoinInfosResponse{}, status.Error(codes.Unavailable, err.Error())
	}

	infos, total, err := coininfo_crud.All(ctx, conds)
	if err != nil {
		logger.Sugar().Error(err)
		return &npool.GetCoinInfosResponse{}, status.Error(codes.Unavailable, err.Error())
	}

	ret := &npool.GetCoinInfosResponse{Total: uint32(total)}
	for _, v := range infos {
		info := &npool.CoinInfoReq{
			ID:              v.ID,
			Name:            v.Name,
			ChainType:       v.ChainType,
			TokenType:       v.TokenType,
			PrivateContract: v.PrivateContract,
			Remark:          v.Remark,
		}
		ret.Infos = append(ret.Infos, info)
	}
	return ret, nil
}
