package main

import (
	"math/rand"
	"time"

	cli "github.com/urfave/cli/v2"

	"github.com/NpoolPlatform/build-chain/pkg/coins"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// TODO: support other tokens
var genTaskInfo = &coins.GenTaskInfo{}

var genCmd = &cli.Command{
	Name:    "gen",
	Aliases: []string{"g"},
	Usage:   "Generate code for sphinx-plugin",
	Before: func(ctx *cli.Context) error {
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "buildChain-server",
			Aliases:     []string{"bc-server"},
			Usage:       "required, host to build-chain server",
			Required:    true,
			Value:       "",
			Destination: &genTaskInfo.Host,
		},
		&cli.StringFlag{
			Name:        "package",
			Aliases:     []string{"p"},
			Usage:       "optional, go file package name",
			Required:    false,
			Value:       "gen",
			Destination: &genTaskInfo.Package,
		},
		&cli.StringFlag{
			Name:        "chaintype",
			Aliases:     []string{"c"},
			Usage:       "optional, chain type",
			Required:    false,
			Value:       "ethereum",
			Destination: &genTaskInfo.ChainType,
		},
		&cli.StringFlag{
			Name:        "tokentype",
			Aliases:     []string{"t"},
			Usage:       "optional, token type",
			Required:    false,
			Value:       "erc20",
			Destination: &genTaskInfo.TokenType,
		},
		&cli.StringFlag{
			Name:        "listname",
			Aliases:     []string{"l"},
			Usage:       "optional, go file package name",
			Required:    false,
			Value:       "tokenList",
			Destination: &genTaskInfo.ListName,
		},
		&cli.StringFlag{
			Name:        "out",
			Aliases:     []string{"o"},
			Usage:       "optional, out file path",
			Required:    false,
			Value:       "data.go",
			Destination: &genTaskInfo.Out,
		},
	},
	Action: func(c *cli.Context) error {
		coins.Gen(genTaskInfo)
		return nil
	},
}
