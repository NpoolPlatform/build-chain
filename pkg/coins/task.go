package coins

import (
	"context"

	handler "github.com/NpoolPlatform/build-chain/pkg/mw/tokeninfo"
)

func Init() {
	// TODO: should be remove,will replace
	h, err := handler.NewHandler(
		context.Background(),
		handler.WithEntID(toStrPinter("5768cdea-8cd3-11ee-9cb5-cfd8fd9e4312"), true),
		handler.WithName(toStrPinter("Ethereum"), false),
		handler.WithUnit(toStrPinter("ETH"), false),
		handler.WithDecimal(toStrPinter("18"), false),
		handler.WithChainType(toStrPinter(EthereumChain), false),
		handler.WithTokenType(toStrPinter(EthereumChain), false),
		handler.WithOfficialContract(toStrPinter("ethereum"), false),
		handler.WithPrivateContract(toStrPinter("ethereum"), false),
	)
	if err != nil {
		panic(err)
	}

	_, err = h.CreateTokenInfo(context.Background())
	if err != nil {
		panic(err)
	}
}

func toStrPinter(s string) *string {
	return &s
}
