package tokeninfo

import (
	"context"

	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
)

func (h *Handler) CreateTokenInfo(ctx context.Context) (*npool.TokenInfo, error) {
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := tokeninfo_crud.CreateSet(
			cli.TokenInfo.Create(),
			&tokeninfo_crud.Req{
				EntID:            h.EntID,
				Name:             h.Name,
				Unit:             h.Unit,
				Decimal:          h.Decimal,
				ChainType:        h.ChainType,
				TokenType:        h.TokenType,
				OfficialContract: h.OfficialContract,
				PrivateContract:  h.PrivateContract,
				Remark:           h.Remark,
				Data:             h.Data,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		h.EntID = &info.EntID
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetTokenInfo(ctx)
}
