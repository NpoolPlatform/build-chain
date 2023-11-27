package tokeninfo

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"github.com/stretchr/testify/assert"
)

func init() {
	//nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
}

var ret = &npool.TokenInfo{
	ChainType:        "test_data",
	TokenType:        "test_data",
	OfficialContract: "test_data",
	Name:             "test_data",
	Unit:             "test_data",
	Decimal:          "test_data",
	PrivateContract:  "test_data",
	Remark:           "test_data",
	Data:             []byte("test_data"),
}

func create(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithChainType(&ret.ChainType, false),
		WithTokenType(&ret.TokenType, false),
		WithOfficialContract(&ret.OfficialContract, false),
		WithName(&ret.Name, true),
		WithUnit(&ret.Unit, false),
		WithDecimal(&ret.Decimal, false),
		WithPrivateContract(&ret.PrivateContract, false),
		WithRemark(&ret.Remark, false),
		WithData(ret.Data, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateTokenInfo(context.Background())
		if assert.Nil(t, err) && assert.NotNil(t, info) {
			ret.EntID = info.EntID
			ret.ID = info.ID
			ret.Name = info.Name
			ret.PrivateContract = info.PrivateContract
			assert.Equal(t, ret.String(), info.String())
		}
	}
}

func get(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetTokenInfo(context.Background())
		if assert.Nil(t, err) && assert.NotNil(t, info) {
			assert.Equal(t, ret.String(), info.String())
		}
	}
}

func update(t *testing.T) {
	ret.PrivateContract = "I don't know"
	ret.Remark = "Somebody want to kill me"

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithPrivateContract(&ret.PrivateContract, false),
		WithRemark(&ret.Remark, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateTokenInfo(context.Background())
		if assert.Nil(t, err) && assert.NotNil(t, info) {
			ret.PrivateContract = info.PrivateContract
			ret.Remark = info.Remark
			assert.Equal(t, ret.String(), info.String())
		}
	}
}

func getConds(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			EntID:           &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			Name:            &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
			PrivateContract: &basetypes.StringVal{Op: cruder.EQ, Value: ret.PrivateContract},
			TokenType:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.TokenType},
		}),
		WithOffset(0),
		WithLimit(2),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetTokenInfos(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, uint32(1), total)
			if assert.Equal(t, 1, len(infos)) {
				assert.Equal(t, infos[0].String(), ret.String())
			}
		}
	}
}

func TestOpLog(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("get", get)
	t.Run("update", update)
	t.Run("getConds", getConds)
}
