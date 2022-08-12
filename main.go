package main

import (
	"context"
	"math/rand"

	"time"

	deployedcoin_crud "github.com/NpoolPlatform/build-chain/pkg/crud/deployedcoin"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// err := logger.Init(logger.DebugLevel, "./debug.log")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// client, err := rpc.Dial("http://192.168.49.1:8545")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	deployedcoin_crud.AllWithCoinName(context.Background())
}
