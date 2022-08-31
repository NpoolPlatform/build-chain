package main

import (
	"math/rand"
	"time"

	cli "github.com/urfave/cli/v2"

	"github.com/NpoolPlatform/build-chain/pkg/coins/eth"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// TODO: support other tokens
var crawlTaskInfo = &eth.CrawlTaskInfo{}

var crawlCmd = &cli.Command{
	Name:    "crawl",
	Aliases: []string{"c"},
	Usage:   "Crawl contract to private chain",
	Before: func(ctx *cli.Context) error {
		return nil
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "host",
			Aliases:     []string{"ho"},
			Usage:       "required, host to build-chain server",
			Required:    true,
			Value:       "",
			Destination: &crawlTaskInfo.Host,
		},
		&cli.BoolFlag{
			Name:        "force",
			Aliases:     []string{"f"},
			Usage:       "optional, force to redeploy contract",
			Required:    false,
			Value:       false,
			Destination: &crawlTaskInfo.Force,
		},
		&cli.StringFlag{
			Name:        "contract",
			Aliases:     []string{"co"},
			Usage:       "optional , appoint the official contract which deployed to build-chain server",
			Required:    false,
			Value:       "",
			Destination: &crawlTaskInfo.Contract,
		},
		&cli.StringFlag{
			Name:        "chainType",
			Aliases:     []string{"ch"},
			Usage:       "optional <ethereum>, appoint the chain type",
			Required:    false,
			Value:       "ethereum",
			Destination: &crawlTaskInfo.ChainType,
		},
		&cli.StringFlag{
			Name:        "tokenType",
			Aliases:     []string{"to"},
			Usage:       "optional <erc20>, appoint the token type",
			Required:    false,
			Value:       "erc20",
			Destination: &crawlTaskInfo.TokenType,
		},
		&cli.IntFlag{
			Name:        "offset",
			Aliases:     []string{"o"},
			Usage:       "optional, appoint the top token offset,use with limit",
			Required:    false,
			Value:       -1,
			Destination: &crawlTaskInfo.Offset,
		},
		&cli.IntFlag{
			Name:        "limit",
			Aliases:     []string{"l"},
			Usage:       "optional, appoint the top token limit,use with offset",
			Required:    false,
			Value:       -1,
			Destination: &crawlTaskInfo.Limit,
		},
	},
	Action: func(c *cli.Context) error {
		eth.Crawl(crawlTaskInfo)
		return nil
	},
}
