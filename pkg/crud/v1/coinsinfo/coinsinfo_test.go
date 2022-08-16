package coinsinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func assertCoinsInfo(t *testing.T, actual, expected *npool.CoinsInfo) {
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.ChainType, expected.ChainType)
	assert.Equal(t, actual.TokenType, expected.TokenType)
	assert.Equal(t, actual.Contract, expected.Contract)
	assert.Equal(t, actual.Remark, expected.Remark)
	assert.Equal(t, actual.Data, expected.Data)
	assert.Equal(t, actual.Similarity, expected.Similarity)
}

func TestCURD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	ctx := context.Background()
	coinsInfo := &npool.CoinsInfo{
		Name:       fmt.Sprintf("test-%v", uuid.New().String()),
		ChainType:  fmt.Sprintf("test-%v", uuid.New().String()),
		TokenType:  fmt.Sprintf("test-%v", uuid.New().String()),
		Contract:   fmt.Sprintf("test-%v", uuid.New().String()),
		Similarity: 10023,
		Remark:     fmt.Sprintf("test-%v", uuid.New().String()),
		Data:       []byte{},
	}

	resp1, err := Create(ctx, coinsInfo)
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp1.ID, uuid.UUID{})
		assertCoinsInfo(t, resp1, coinsInfo)
	}

	resp2, err := CheckContractAndCreate(ctx, coinsInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.ID, resp1.ID)
		assertCoinsInfo(t, resp2, coinsInfo)
	}

	coinsInfo.Contract = fmt.Sprintf("test-%v", uuid.New().String())
	resp3, err := CheckContractAndCreate(ctx, coinsInfo)
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp3.ID, resp2.ID)
		assertCoinsInfo(t, resp3, coinsInfo)
	}

	resp3.Contract = fmt.Sprintf("test-%v", uuid.New().String())
	resp4, err := Update(ctx, resp3)
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.ID, resp3.ID)
		assertCoinsInfo(t, resp3, resp4)
	}

	resp5, err := Row(ctx, uuid.MustParse(resp4.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.ID, resp4.ID)
		assertCoinsInfo(t, resp5, resp4)
	}

	conds := cruder.NewConds()
	conds.WithCond(coinsinfo.FieldName, cruder.EQ, coinsInfo.Name)
	resp6, n, err := Rows(ctx, conds, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, n, 2)
		assert.Equal(t, len(resp6), 1)
	}

	resp7, n, err := All(ctx, conds)
	if assert.Nil(t, err) {
		assert.Equal(t, n, 2)
		assert.Equal(t, len(resp7), 2)
	}

}
