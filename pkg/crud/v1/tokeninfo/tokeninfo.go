package tokeninfo

import (
	"context"
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

type Conds struct {
	ID               *cruder.Cond
	Name             *cruder.Cond
	ChainType        *cruder.Cond
	TokenType        *cruder.Cond
	Unit             *cruder.Cond
	Decimal          *cruder.Cond
	OfficialContract *cruder.Cond
	PrivateContract  *cruder.Cond
	Remark           *cruder.Cond
	Data             *cruder.Cond
}

func Rows(ctx context.Context, conds *Conds, offset, limit int) ([]*npool.TokenInfo, int, error) {
	var err error
	rows := []*ent.TokenInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.TokenInfo.Query()
		stm, err = queryFromConds(stm, conds)
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

func RowOnly(ctx context.Context, conds *Conds) (*npool.TokenInfo, error) {
	var err error
	row := &ent.TokenInfo{}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.TokenInfo.Query()
		stm, err = queryFromConds(stm, conds)
		if err != nil {
			return err
		}

		row, err = stm.Only(ctx)
		return err
	})

	if err != nil {
		return nil, err
	}
	return toObj(row), nil
}

func All(ctx context.Context, conds *Conds) ([]*npool.TokenInfo, int, error) {
	var err error
	rows := []*ent.TokenInfo{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.TokenInfo.Query()
		stm, err = queryFromConds(stm, conds)
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
func queryFromConds(q *ent.TokenInfoQuery, conds *Conds) (*ent.TokenInfoQuery, error) {
	q.Where(tokeninfo.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.ID(id))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.Name(name))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.ChainType != nil {
		_type, ok := conds.ChainType.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid chaintype")
		}
		switch conds.ChainType.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.ChainType(_type))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.TokenType != nil {
		_type, ok := conds.TokenType.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid tokentype")
		}
		switch conds.TokenType.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.TokenType(_type))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.Unit != nil {
		_type, ok := conds.Unit.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid unit")
		}
		switch conds.Unit.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.Unit(_type))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.Decimal != nil {
		_type, ok := conds.Decimal.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid decimal")
		}
		switch conds.Decimal.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.Decimal(_type))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.OfficialContract != nil {
		contract, ok := conds.OfficialContract.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid officialcontract")
		}
		switch conds.OfficialContract.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.OfficialContract(contract))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.PrivateContract != nil {
		contract, ok := conds.PrivateContract.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid privatecontract")
		}
		switch conds.PrivateContract.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.PrivateContract(contract))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	if conds.Data != nil {
		data, ok := conds.Data.Val.([]byte)
		if !ok {
			return nil, fmt.Errorf("invalid data")
		}
		switch conds.Data.Op {
		case cruder.EQ:
			q = q.Where(tokeninfo.Data(data))
		default:
			return nil, fmt.Errorf("invalid tokeninfo field")
		}
	}
	return q, nil
}
