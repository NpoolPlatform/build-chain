package tokeninfo

import (
	"context"
	"errors"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/tokeninfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/google/uuid"
)

func toObj(coin *ent.TokenInfo) *npool.TokenInfo {
	if coin == nil {
		return nil
	}
	return &npool.TokenInfo{
		ID:               coin.ID.String(),
		Name:             coin.Name,
		ChainType:        coin.ChainType,
		TokenType:        coin.TokenType,
		Unit:             coin.Unit,
		Decimal:          coin.Decimal,
		OfficialContract: coin.OfficialContract,
		PrivateContract:  coin.PrivateContract,
		Remark:           coin.Remark,
		Data:             coin.Data,
	}
}

func toObjs(coins []*ent.TokenInfo) []*npool.TokenInfo {
	_coins := []*npool.TokenInfo{}
	for _, v := range coins {
		_coins = append(_coins, toObj(v))
	}
	return _coins
}

func Create(ctx context.Context, in *npool.TokenInfo) (*npool.TokenInfo, error) {
	var info *ent.TokenInfo
	var err error

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = tx.TokenInfo.Create().
			SetName(in.GetName()).
			SetChainType(in.ChainType).
			SetTokenType(in.TokenType).
			SetUnit(in.Unit).
			SetDecimal(in.Decimal).
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

func Update(ctx context.Context, in *npool.TokenInfo) (*npool.TokenInfo, error) {
	var ret *ent.TokenInfo
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return nil, err
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		ret, err = tx.TokenInfo.UpdateOneID(id).
			SetName(in.GetName()).
			SetChainType(in.ChainType).
			SetTokenType(in.TokenType).
			SetUnit(in.Unit).
			SetDecimal(in.Decimal).
			SetOfficialContract(in.OfficialContract).
			SetPrivateContract(in.PrivateContract).
			SetRemark(in.Remark).
			SetData(in.Data).
			Save(ctx)
		return err
	})
	return toObj(ret), err
}

func Row(ctx context.Context, id uuid.UUID) (*npool.TokenInfo, error) {
	var info *ent.TokenInfo
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.TokenInfo.Query().Where(tokeninfo.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return toObj(info), nil
}

func Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.TokenInfo, int, error) {
	var err error
	rows := []*ent.TokenInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.TokenInfo.Query()
		stm, err = queryFromConds(conds, stm)
		if err != nil {
			return err
		}

		rows, err = stm.Order(ent.Desc(tokeninfo.FieldCreatedAt)).Offset(offset).Limit(limit).All(ctx)
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

func All(ctx context.Context, conds cruder.Conds) ([]*npool.TokenInfo, int, error) {
	var err error
	rows := []*ent.TokenInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.TokenInfo.Query()
		stm, err = queryFromConds(conds, stm)
		if err != nil {
			return err
		}

		rows, err = stm.Order(ent.Desc(tokeninfo.FieldCreatedAt)).All(ctx)
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
func queryFromConds(conds cruder.Conds, stm *ent.TokenInfoQuery) (*ent.TokenInfoQuery, error) {
	for k, v := range conds {
		switch k {
		case tokeninfo.FieldID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(tokeninfo.ID(id))
		case tokeninfo.FieldName:
			name, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid AppID: %v", err)
			}
			stm = stm.Where(tokeninfo.Name(name))
		case tokeninfo.FieldChainType:
			chainType, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid cointypeid: %v", err)
			}
			stm = stm.Where(tokeninfo.ChainType(chainType))
		case tokeninfo.FieldTokenType:
			tokenType, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid UsedFor: %v", err)
			}
			stm = stm.Where(tokeninfo.TokenType(tokenType))
		case tokeninfo.FieldUnit:
			unit, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid cointypeid: %v", err)
			}
			stm = stm.Where(tokeninfo.Unit(unit))
		case tokeninfo.FieldDecimal:
			decimal, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid UsedFor: %v", err)
			}
			stm = stm.Where(tokeninfo.Decimal(decimal))
		case tokeninfo.FieldOfficialContract:
			contract, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Message: %v", err)
			}
			stm = stm.Where(tokeninfo.OfficialContract(contract))
		case tokeninfo.FieldPrivateContract:
			contract, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(tokeninfo.PrivateContract(contract))
		case tokeninfo.FieldRemark:
			remark, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(tokeninfo.Remark(remark))
		case tokeninfo.FieldData:
			data, err := AnyTypeBytes(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(tokeninfo.Data(data))
		default:
			return nil, fmt.Errorf("invalid TokenInfo field")
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
