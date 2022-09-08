package coins

import (
	"context"

	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	proto "github.com/NpoolPlatform/message/npool/build-chain"
)

func Init() {
	// TODO: should be remove,will replace
	tokeninfo_crud.Create(context.Background(), &proto.TokenInfo{
		Name:             "Ethereum",
		Unit:             "ETH",
		Decimal:          "18",
		ChainType:        EthereumChain,
		TokenType:        EthereumChain,
		OfficialContract: "ethereum",
		PrivateContract:  "ethereum",
	})
}
