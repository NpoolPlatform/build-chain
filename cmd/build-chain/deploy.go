package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	bc_client "github.com/NpoolPlatform/build-chain/pkg/client/v1"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	proto "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"github.com/gocarina/gocsv"
	cli "github.com/urfave/cli/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DeployInfo struct {
	FilePath  string
	Force     bool
	BC_Server string
}

var deployInfo = &DeployInfo{}

var deployCmd = &cli.Command{
	Name:    "deploy",
	Aliases: []string{"r"},
	Usage:   "deploy smart contract from file",
	Before: func(ctx *cli.Context) error {
		err := logger.Init(logger.DebugLevel, "./")
		if err != nil {
			panic(fmt.Errorf("fail to init logger: %v", err))
		}
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Aliases:     []string{"fi"},
			Usage:       "must , read contract info from file",
			Required:    true,
			Value:       "",
			Destination: &deployInfo.FilePath,
		},
		&cli.BoolFlag{
			Name:        "force",
			Aliases:     []string{"fo"},
			Usage:       "optional (default: false), if already exist force deploy",
			Required:    false,
			Value:       false,
			Destination: &deployInfo.Force,
		},
		&cli.StringFlag{
			Name:        "bc-server",
			Aliases:     []string{"bc-server"},
			Usage:       "must , build-chain server address",
			Required:    true,
			Value:       "",
			Destination: &deployInfo.BC_Server,
		},
	},
	Action: func(c *cli.Context) error {
		return Deploy(context.Background(), deployInfo)
	},
}

func Deploy(ctx context.Context, deployInfo *DeployInfo) error {
	bcConn, err := bc_client.NewClientConn(ctx, deployInfo.BC_Server)
	if err != nil {
		return err
	}
	defer bcConn.Close()

	tokens, err := ReadTokenInfos(deployInfo)
	if err != nil {
		return err
	}
	for _, v := range tokens {
		err = CreateTokenInfo(ctx, bcConn, v, deployInfo.Force)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadTokenInfos(deployInfo *DeployInfo) ([]*proto.TokenInfoReq, error) {
	in, err := os.Open(deployInfo.FilePath)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	tokenInfos := []*proto.TokenInfoReq{}

	if err := gocsv.UnmarshalFile(in, &tokenInfos); err != nil {
		return nil, err
	}

	return tokenInfos, nil
}

func CreateTokenInfo(ctx context.Context, bcConn *bc_client.BuildChainClientConn, token *proto.TokenInfoReq, force bool) (err error) {
	if !force {
		conds := &proto.Conds{
			OfficialContract: &basetypes.StringVal{Op: cruder.EQ, Value: *token.OfficialContract},
		}
		resp, err := bcConn.GetTokenInfos(context.Background(), &proto.GetTokenInfosRequest{Conds: conds})
		if err == nil && len(resp.Infos) != 0 {
			return nil
		}
	}

	retry := true
	for i := 0; i < 3 && retry; i++ {
		retry = false
		_, err := bcConn.CreateTokenInfo(ctx, &proto.CreateTokenInfoRequest{
			Force: force,
			Info:  token,
		})
		if err != nil &&
			strings.Contains(err.Error(), "replacement transaction underpriced") &&
			strings.Contains(err.Error(), " max fee per gas less than block base fee") {
			i--
			retry = true
		}
	}

	if err != nil {
		return fmt.Errorf("failed create token %v, %v", token.Name, err)
	}
	return nil
}
