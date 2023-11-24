package tokeninfo

import (
	"context"
	"fmt"

	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	ent_tokeninfo "github.com/NpoolPlatform/build-chain/pkg/db/ent/tokeninfo"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
)

func (h *Handler) UpdateTokenInfo(ctx context.Context) (*npool.TokenInfo, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			TokenInfo.
			Query().
			Where(
				ent_tokeninfo.ID(*h.ID),
			).
			Only(_ctx)
		if err != nil {
			return err
		}

		info, err = tokeninfo_crud.UpdateSet(
			info.Update(),
			&tokeninfo_crud.Req{
				PrivateContract: h.PrivateContract,
				Remark:          h.Remark,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.EntID = &info.EntID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetTokenInfo(ctx)
}
