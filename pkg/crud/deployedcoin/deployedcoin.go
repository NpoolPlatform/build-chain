package deployedcoin

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/deployedcoin"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"

	"github.com/google/uuid"
)

func toObj(coin *ent.DeployedCoin) *npool.DeployedCoin {
	return &npool.DeployedCoin{
		ID:       coin.ID.String(),
		CoinID:   coin.CoinID.String(),
		Contract: coin.Contract,
	}
}

func toObjs(coins []*ent.DeployedCoin) []*npool.DeployedCoin {
	_coins := []*npool.DeployedCoin{}
	for _, v := range coins {
		_coins = append(_coins, toObj(v))
	}
	return _coins
}

func Create(ctx context.Context, in *npool.DeployedCoin) (*npool.DeployedCoin, error) {
	var info *ent.DeployedCoin
	var err error

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err = tx.DeployedCoin.Create().
			SetCoinID(uuid.MustParse(in.GetCoinID())).
			SetContract(in.GetContract()).
			Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return toObj(info), nil
}

func CreateBulk(ctx context.Context, in []*npool.DeployedCoin) ([]*npool.DeployedCoin, error) {
	var err error

	rows := []*ent.DeployedCoin{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.DeployedCoinCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.DeployedCoin.Create().
				SetCoinID(uuid.MustParse(info.GetCoinID())).
				SetContract(info.GetContract())
		}
		rows, err = tx.DeployedCoin.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return toObjs(rows), nil
}

func Row(ctx context.Context, conds cruder.Conds) (*npool.DeployedCoin, error) {
	var info *ent.DeployedCoin
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.DeployedCoin.Query()
		stm, err = queryFromConds(conds, stm)
		if err != nil {
			return err
		}
		info, err = stm.Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return toObj(info), nil
}

func Rows(ctx context.Context, offset, limit int) ([]*npool.DeployedCoin, int, error) {
	var err error
	rows := []*ent.DeployedCoin{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.DeployedCoin.Query()
		rows, err = stm.Order(ent.Desc(deployedcoin.FieldCreatedAt)).Offset(offset).Limit(limit).All(ctx)
		total, err = stm.Count(ctx)
		return err
	})
	if err != nil {
		return nil, 0, err
	}
	return toObjs(rows), total, nil
}

func All(ctx context.Context) ([]*npool.DeployedCoin, int, error) {
	var err error
	rows := []*ent.DeployedCoin{}
	var total int

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm := tx.DeployedCoin.Query()
		rows, err = stm.Order(ent.Desc(deployedcoin.FieldCreatedAt)).All(ctx)
		total, err = stm.Count(ctx)
		return err
	})
	if err != nil {
		return nil, 0, err
	}
	return toObjs(rows), total, nil
}

//nolint
func queryFromConds(conds cruder.Conds, stm *ent.DeployedCoinQuery) (*ent.DeployedCoinQuery, error) {
	for k, v := range conds {
		switch k {
		case deployedcoin.FieldID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(deployedcoin.ID(id))
		case deployedcoin.FieldCoinID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(deployedcoin.CoinID(id))
		case deployedcoin.FieldContract:
			contract, err := cruder.AnyTypeString(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(deployedcoin.Contract(contract))
		default:
			return nil, fmt.Errorf("invalid CoinsInfo field")
		}
	}

	return stm, nil
}

func AllWithCoinName(ctx context.Context) ([]*npool.DeployedCoinInfo, error) {
	coinsInfo := []*npool.DeployedCoinInfo{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		err := tx.DeployedCoin.Query().
			Order(ent.Desc(deployedcoin.FieldCreatedAt)).
			Select(
				deployedcoin.FieldID,
				deployedcoin.FieldCoinID,
				deployedcoin.FieldContract,
			).
			Modify(func(s *sql.Selector) {
				t := sql.Table(coinsinfo.Table)
				s.LeftJoin(t).On(
					s.C(deployedcoin.FieldCoinID),
					t.C(coinsinfo.FieldID),
				).AppendSelect(
					t.C(coinsinfo.FieldName),
					t.C(coinsinfo.FieldChainType),
					t.C(coinsinfo.FieldTokenType),
				)
			}).Scan(ctx, &coinsInfo)
		return err
	})

	return coinsInfo, err
}
