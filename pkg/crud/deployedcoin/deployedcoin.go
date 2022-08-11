package deployedcoin

import (
	"context"

	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/deployedcoin"
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

func Row(ctx context.Context, id uuid.UUID) (*npool.DeployedCoin, error) {
	var info *ent.DeployedCoin
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.DeployedCoin.Query().Where(deployedcoin.ID(id)).Only(_ctx)
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
