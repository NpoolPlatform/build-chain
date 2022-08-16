package deployedcoin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/build-chain/pkg/db/ent/deployedcoin"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertDeployedCoin(t *testing.T, actual, expected *npool.DeployedCoin) {
	assert.Equal(t, actual.CoinID, expected.CoinID)
	assert.Equal(t, actual.Contract, expected.Contract)
}

func TestCURD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	ctx := context.Background()
	deployedCoin := &npool.DeployedCoin{
		CoinID:   uuid.New().String(),
		Contract: fmt.Sprintf("test-%v", uuid.New().String()),
	}

	resp1, err := Create(ctx, deployedCoin)
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.ID, uuid.UUID{})
		assertDeployedCoin(t, resp1, deployedCoin)
	}

	resp2, err := Row(ctx, uuid.MustParse(resp1.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.ID, resp1.ID)
		assertDeployedCoin(t, resp2, resp1)
	}

	conds := cruder.NewConds()
	conds.WithCond(deployedcoin.FieldContract, cruder.EQ, deployedCoin.Contract)
	resp3, n, err := Rows(ctx, conds, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, n, 1)
		assert.Equal(t, len(resp3), 1)
	}

	resp4, n, err := All(ctx, conds)
	if assert.Nil(t, err) {
		assert.Equal(t, n, 1)
		assert.Equal(t, len(resp4), 1)
	}

}
