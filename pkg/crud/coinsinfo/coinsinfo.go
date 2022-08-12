package coinsinfo

import (
	"context"
	"errors"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/google/uuid"
)

func toObj(coin *ent.CoinsInfo) *npool.CoinsInfo {
	return &npool.CoinsInfo{
		ID:         coin.ID.String(),
		Name:       coin.Name,
		ChainType:  coin.ChainType,
		TokenType:  coin.TokenType,
		Contract:   coin.Contract,
		Similarity: coin.Similarity,
		Remark:     coin.Remark,
		Data:       coin.Data,
	}
}

func toObjs(coins []*ent.CoinsInfo) []*npool.CoinsInfo {
	_coins := []*npool.CoinsInfo{}
	for _, v := range coins {
		_coins = append(_coins, toObj(v))
	}
	return _coins
}

func Create(ctx context.Context, in *npool.CoinsInfo) (*npool.CoinsInfo, error) {
	var info *ent.CoinsInfo
	var err error

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = tx.CoinsInfo.Create().
			SetName(in.GetName()).
			SetChainType(in.ChainType).
			SetTokenType(in.TokenType).
			SetContract(in.Contract).
			SetSimilarity(in.Similarity).
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

func CheckContractAndCreate(ctx context.Context, in *npool.CoinsInfo) (*npool.CoinsInfo, error) {
	var info *ent.CoinsInfo

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		infos, err := tx.CoinsInfo.Query().Where(coinsinfo.Contract(in.Contract)).All(ctx)
		if err != nil {
			return err
		}
		if len(infos) > 0 {
			info = infos[0]
			return nil
		}
		info, err = tx.CoinsInfo.Create().
			SetName(in.GetName()).
			SetChainType(in.ChainType).
			SetTokenType(in.TokenType).
			SetContract(in.Contract).
			SetSimilarity(in.Similarity).
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

func Update(ctx context.Context, id uuid.UUID, in *npool.CoinsInfo) error {
	var err error
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		_, err = tx.CoinsInfo.UpdateOneID(id).
			SetName(in.GetName()).
			SetChainType(in.ChainType).
			SetTokenType(in.TokenType).
			SetContract(in.Contract).
			SetSimilarity(in.Similarity).
			SetRemark(in.Remark).
			SetData(in.Data).
			Save(ctx)
		return err
	})
	return err
}

func CreateBulk(ctx context.Context, in []*npool.CoinsInfo) ([]*npool.CoinsInfo, error) {
	var err error

	rows := []*ent.CoinsInfo{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.CoinsInfoCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.CoinsInfo.Create().
				SetName(info.GetName()).
				SetChainType(info.ChainType).
				SetTokenType(info.TokenType).
				SetContract(info.Contract).
				SetSimilarity(info.Similarity).
				SetRemark(info.Remark).
				SetData(info.Data)
		}
		rows, err = tx.CoinsInfo.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return toObjs(rows), nil
}

func Row(ctx context.Context, id uuid.UUID) (*npool.CoinsInfo, error) {
	var info *ent.CoinsInfo
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.CoinsInfo.Query().Where(coinsinfo.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return toObj(info), nil
}

func Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.CoinsInfo, int, error) {
	var err error
	rows := []*ent.CoinsInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.CoinsInfo.Query()
		stm, err = queryFromConds(conds, stm)
		if err != nil {
			return err
		}

		rows, err = stm.Order(ent.Desc(coinsinfo.FieldCreatedAt)).Offset(offset).Limit(limit).All(ctx)
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

func All(ctx context.Context, conds cruder.Conds) ([]*npool.CoinsInfo, int, error) {
	var err error
	rows := []*ent.CoinsInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.CoinsInfo.Query()
		stm, err = queryFromConds(conds, stm)
		if err != nil {
			return err
		}

		rows, err = stm.Order(ent.Desc(coinsinfo.FieldCreatedAt)).All(ctx)
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
func queryFromConds(conds cruder.Conds, stm *ent.CoinsInfoQuery) (*ent.CoinsInfoQuery, error) {
	for k, v := range conds {
		switch k {
		case coinsinfo.FieldID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(coinsinfo.ID(id))
		case coinsinfo.FieldName:
			name, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid AppID: %v", err)
			}
			stm = stm.Where(coinsinfo.Name(name))
		case coinsinfo.FieldChainType:
			chainType, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid cointypeid: %v", err)
			}
			stm = stm.Where(coinsinfo.ChainType(chainType))
		case coinsinfo.FieldTokenType:
			tokenType, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid UsedFor: %v", err)
			}
			stm = stm.Where(coinsinfo.TokenType(tokenType))
		case coinsinfo.FieldContract:
			contract, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Message: %v", err)
			}
			stm = stm.Where(coinsinfo.Contract(contract))
		case coinsinfo.FieldSimilarity:
			similarity, err := cruder.AnyTypeInt32(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(coinsinfo.Similarity(similarity))
		case coinsinfo.FieldRemark:
			remark, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(coinsinfo.Remark(remark))
		case coinsinfo.FieldData:
			data, err := AnyTypeBytes(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid Title: %v", err)
			}
			stm = stm.Where(coinsinfo.Data(data))
		default:
			return nil, fmt.Errorf("invalid CoinsInfo field")
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
