package coininfo

import (
	"context"
	"errors"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/google/uuid"
)

func toObj(coin *ent.CoinInfo) *npool.CoinInfo {
	if coin == nil {
		return nil
	}
	return &npool.CoinInfo{
		ID:               coin.ID.String(),
		Name:             coin.Name,
		ChainType:        coin.ChainType,
		TokenType:        coin.TokenType,
		OfficialContract: coin.OfficialContract,
		PrivateContract:  coin.PrivateContract,
		Remark:           coin.Remark,
		Data:             coin.Data,
	}
}

func toObjs(coins []*ent.CoinInfo) []*npool.CoinInfo {
	_coins := []*npool.CoinInfo{}
	for _, v := range coins {
		_coins = append(_coins, toObj(v))
	}
	return _coins
}

func Create(ctx context.Context, in *npool.CoinInfo) (*npool.CoinInfo, error) {
	var info *ent.CoinInfo
	var err error

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = tx.CoinInfo.Create().
			SetName(in.GetName()).
			SetChainType(in.ChainType).
			SetTokenType(in.TokenType).
			SetOfficialContract(in.OfficialContract).
			SetPrivateContract(in.PrivateContract).
			SetRemark(in.Remark).
			SetData(in.Data).
			Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return toObj(info), nil
}

func Update(ctx context.Context, in *npool.CoinInfo) (*npool.CoinInfo, error) {
	var ret *ent.CoinInfo
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return nil, err
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		ret, err = tx.CoinInfo.UpdateOneID(id).
			SetName(in.GetName()).
			SetChainType(in.ChainType).
			SetTokenType(in.TokenType).
			SetOfficialContract(in.OfficialContract).
			SetPrivateContract(in.PrivateContract).
			SetRemark(in.Remark).
			SetData(in.Data).
			Save(ctx)
		return err
	})
	return toObj(ret), err
}

func Row(ctx context.Context, id uuid.UUID) (*npool.CoinInfo, error) {
	var info *ent.CoinInfo
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.CoinInfo.Query().Where(coininfo.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return toObj(info), nil
}

func Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.CoinInfo, int, error) {
	var err error
	rows := []*ent.CoinInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.CoinInfo.Query()
		stm, err = queryFromConds(conds, stm)
		if err != nil {
			return err
		}

		rows, err = stm.Order(ent.Desc(coininfo.FieldCreatedAt)).Offset(offset).Limit(limit).All(ctx)
		if err != nil {
			return err
		}

		total, err = stm.Count(ctx)
		return err
	})
	if err != nil {
		return nil, 0, err
	}
	return toObjs(rows), total, nil
}

func All(ctx context.Context, conds cruder.Conds) ([]*npool.CoinInfo, int, error) {
	var err error
	rows := []*ent.CoinInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.CoinInfo.Query()
		stm, err = queryFromConds(conds, stm)
		if err != nil {
			return err
		}

		rows, err = stm.Order(ent.Desc(coininfo.FieldCreatedAt)).All(ctx)
		if err != nil {
			return err
		}

		total, err = stm.Count(ctx)
		return err
	})
	if err != nil {
		return nil, 0, err
	}
	return toObjs(rows), total, nil
}

//nolint
func queryFromConds(conds cruder.Conds, stm *ent.CoinInfoQuery) (*ent.CoinInfoQuery, error) {
	for k, v := range conds {
		switch k {
		case coininfo.FieldID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(coininfo.ID(id))
		case coininfo.FieldName:
			name, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid AppID: %v", err)
			}
			stm = stm.Where(coininfo.Name(name))
		case coininfo.FieldChainType:
			chainType, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid cointypeid: %v", err)
			}
			stm = stm.Where(coininfo.ChainType(chainType))
		case coininfo.FieldTokenType:
			tokenType, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid UsedFor: %v", err)
			}
			stm = stm.Where(coininfo.TokenType(tokenType))
		case coininfo.FieldOfficialContract:
			contract, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Message: %v", err)
			}
			stm = stm.Where(coininfo.OfficialContract(contract))
		case coininfo.FieldPrivateContract:
			contract, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(coininfo.PrivateContract(contract))
		case coininfo.FieldRemark:
			remark, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(coininfo.Remark(remark))
		case coininfo.FieldData:
			data, err := AnyTypeBytes(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(coininfo.Data(data))
		default:
			return nil, fmt.Errorf("invalid CoinInfo field")
		}
	}

	return stm, nil
}

func AnyTypeBytes(v cruder.Any) ([]byte, error) {
	if _, ok := v.([]byte); !ok {
		return nil, errors.New("can not parse to []byte")
	}
	return v.([]byte), nil
}
