package coininfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertCoinInfo(t *testing.T, actual, expected *npool.CoinInfo) {
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.ChainType, expected.ChainType)
	assert.Equal(t, actual.TokenType, expected.TokenType)
	assert.Equal(t, actual.OfficialContract, expected.OfficialContract)
	assert.Equal(t, actual.Remark, expected.Remark)
	assert.Equal(t, actual.Data, expected.Data)
	assert.Equal(t, actual.PrivateContract, expected.PrivateContract)
}

func TestCURD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	ctx := context.Background()
	coinsInfo := &npool.CoinInfo{
		Name:             fmt.Sprintf("test-%v", uuid.New().String()),
		ChainType:        fmt.Sprintf("test-%v", uuid.New().String()),
		TokenType:        fmt.Sprintf("test-%v", uuid.New().String()),
		OfficialContract: fmt.Sprintf("test-%v", uuid.New().String()),
		PrivateContract:  fmt.Sprintf("test-%v", uuid.New().String()),
		Remark:           fmt.Sprintf("test-%v", uuid.New().String()),
		Data:             []byte{},
	}

	resp1, err := Create(ctx, coinsInfo)
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.ID, uuid.UUID{})
		assertCoinInfo(t, resp1, coinsInfo)
	}

	resp1.OfficialContract = fmt.Sprintf("test-%v", uuid.New().String())
	resp4, err := Update(ctx, resp1)
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.ID, resp1.ID)
		assertCoinInfo(t, resp1, resp4)
	}

	resp5, err := Row(ctx, uuid.MustParse(resp4.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.ID, resp4.ID)
		assertCoinInfo(t, resp5, resp4)
	}

	conds := cruder.NewConds()
	conds.WithCond(coininfo.FieldName, cruder.EQ, coinsInfo.Name)
	resp6, n, err := Rows(ctx, conds, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, n, 1)
		assert.Equal(t, len(resp6), 1)
	}

	resp7, n, err := All(ctx, conds)
	if assert.Nil(t, err) {
		assert.Equal(t, n, 1)
		assert.Equal(t, len(resp7), 1)
	}
}
