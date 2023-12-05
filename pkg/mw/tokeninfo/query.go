package tokeninfo

import (
	"context"
	"fmt"

	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	"github.com/NpoolPlatform/build-chain/pkg/db"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	ent_tokeninfo "github.com/NpoolPlatform/build-chain/pkg/db/ent/tokeninfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
)

func (h *Handler) GetTokenInfo(ctx context.Context) (*npool.TokenInfo, error) {
	if h.ID == nil && h.EntID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	h.Offset = 0
	h.Limit = 2
	h.Conds = &tokeninfo_crud.Conds{
		EntID: &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
	}
	if h.ID != nil {
		h.Conds.ID = &cruder.Cond{Op: cruder.EQ, Val: *h.ID}
	}
	if h.EntID != nil {
		h.Conds.EntID = &cruder.Cond{Op: cruder.EQ, Val: *h.EntID}
	}

	infos, _, err := h.GetTokenInfos(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
}

func (h *Handler) GetTokenInfos(ctx context.Context) ([]*npool.TokenInfo, uint32, error) {
	infos := []*npool.TokenInfo{}
	total := uint32(0)

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := tokeninfo_crud.SetQueryConds(
			cli.TokenInfo.Query(),
			h.Conds,
		)
		if err != nil {
			return err
		}

		_total, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(_total)
		return stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Select(
				ent_tokeninfo.FieldID,
				ent_tokeninfo.FieldEntID,
				ent_tokeninfo.FieldName,
				ent_tokeninfo.FieldUnit,
				ent_tokeninfo.FieldDecimal,
				ent_tokeninfo.FieldChainType,
				ent_tokeninfo.FieldTokenType,
				ent_tokeninfo.FieldOfficialContract,
				ent_tokeninfo.FieldPrivateContract,
				ent_tokeninfo.FieldRemark,
				ent_tokeninfo.FieldData,
				ent_tokeninfo.FieldCreatedAt,
				ent_tokeninfo.FieldUpdatedAt,
			).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}
