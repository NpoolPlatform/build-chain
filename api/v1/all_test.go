package api

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/build-chain/pkg/client/v1"
	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
	proto "github.com/NpoolPlatform/message/npool/build-chain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	host := "127.0.0.1:50491"
	eth.Crawl(&eth.CrawlTaskInfo{
		Host:      host,
		Offset:    1,
		Limit:     100,
		Force:     true,
		ChainType: "ethereum",
		TokenType: "erc20",
	})

	conn, err := client.NewClientConn(context.Background(), host)
	assert.Nil(t, err)

	resp, err := conn.GetTokenInfos(context.Background(), &proto.GetTokenInfosRequest{})
	assert.Nil(t, err)
	assert.Equal(t, resp.Total, uint32(len(resp.Infos)))

	amount := "1.11"
	to := "0xAfBacba7435dAB4b68E446bb4e8744174b592b57"
	for _, info := range resp.Infos {
		faucetResp, err := conn.Faucet(context.Background(), &proto.FaucetRequst{
			To:               to,
			Amount:           amount,
			OfficialContract: info.OfficialContract,
		})
		assert.Nil(t, err)
		assert.Equal(t, faucetResp.Success, true)
	}

	time.Sleep(3 * time.Second)
	for _, info := range resp.Infos {
		ret, err := eth.ERC20Balance(common.HexToAddress(info.PrivateContract), common.HexToAddress(to))
		assert.Nil(t, err)
		assert.Equal(t, ret, amount)
	}
}
