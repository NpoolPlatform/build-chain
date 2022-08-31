package main

import (
	"github.com/NpoolPlatform/build-chain/pkg/coins"
)

func main() {
	// eth.CrawlContractInfo("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	coins.Gen(&coins.GenTaskInfo{
		Host:      "127.0.0.1:50491",
		Package:   "erc20",
		ChainType: "ethereum",
		TokenType: "erc20",
		ListName:  "erc20data",
	})
}
