package tokeninfo

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/constant"
	tokeninfo_crud "github.com/NpoolPlatform/build-chain/pkg/crud/v1/tokeninfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/build-chain/v1"
	"github.com/google/uuid"
)

type Handler struct {
	ID               *uint32
	EntID            *uuid.UUID
	Name             *string
	Unit             *string
	Decimal          *string
	ChainType        *string
	TokenType        *string
	OfficialContract *string
	PrivateContract  *string
	Remark           *string
	Data             []byte
	Conds            *tokeninfo_crud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" && must {
			return fmt.Errorf("invalid name")
		}
		h.Name = name
		return nil
	}
}
func WithUnit(unit *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if unit == nil {
			if must {
				return fmt.Errorf("invalid unit")
			}
			return nil
		}
		if *unit == "" && must {
			return fmt.Errorf("invalid unit")
		}
		h.Unit = unit
		return nil
	}
}
func WithDecimal(decimal *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if decimal == nil {
			if must {
				return fmt.Errorf("invalid decimal")
			}
			return nil
		}
		if *decimal == "" && must {
			return fmt.Errorf("invalid decimal")
		}
		h.Decimal = decimal
		return nil
	}
}
func WithChainType(chain_type *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if chain_type == nil {
			if must {
				return fmt.Errorf("invalid chain_type")
			}
			return nil
		}
		if *chain_type == "" && must {
			return fmt.Errorf("invalid chain_type")
		}
		h.ChainType = chain_type
		return nil
	}
}
func WithTokenType(token_type *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if token_type == nil {
			if must {
				return fmt.Errorf("invalid token_type")
			}
			return nil
		}
		if *token_type == "" && must {
			return fmt.Errorf("invalid token_type")
		}
		h.TokenType = token_type
		return nil
	}
}
func WithOfficialContract(official_contract *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if official_contract == nil {
			if must {
				return fmt.Errorf("invalid official_contract")
			}
			return nil
		}
		if *official_contract == "" && must {
			return fmt.Errorf("invalid official_contract")
		}
		h.OfficialContract = official_contract
		return nil
	}
}
func WithPrivateContract(private_contract *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if private_contract == nil {
			if must {
				return fmt.Errorf("invalid private_contract")
			}
			return nil
		}
		if *private_contract == "" && must {
			return fmt.Errorf("invalid private_contract")
		}
		h.PrivateContract = private_contract
		return nil
	}
}
func WithRemark(remark *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if remark == nil {
			if must {
				return fmt.Errorf("invalid remark")
			}
			return nil
		}
		if *remark == "" && must {
			return fmt.Errorf("invalid remark")
		}
		h.Remark = remark
		return nil
	}
}
func WithData(data []byte, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if data == nil {
			if must {
				return fmt.Errorf("invalid data")
			}
			return nil
		}
		if len(data) == 0 && must {
			return fmt.Errorf("invalid data")
		}
		h.Data = data
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &tokeninfo_crud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op:  conds.GetEntID().GetOp(),
				Val: id,
			}
		}
		if conds.Unit != nil {
			if conds.Unit.GetValue() != "" {
				h.Conds.Unit = &cruder.Cond{
					Op:  conds.GetUnit().GetOp(),
					Val: conds.GetUnit().GetValue(),
				}
			}
		}
		if conds.Decimal != nil {
			if conds.Decimal.GetValue() != "" {
				h.Conds.Decimal = &cruder.Cond{
					Op:  conds.GetDecimal().GetOp(),
					Val: conds.GetDecimal().GetValue(),
				}
			}
		}
		if conds.ChainType != nil {
			if conds.ChainType.GetValue() != "" {
				h.Conds.ChainType = &cruder.Cond{
					Op:  conds.GetChainType().GetOp(),
					Val: conds.GetChainType().GetValue(),
				}
			}
		}
		if conds.TokenType != nil {
			if conds.TokenType.GetValue() != "" {
				h.Conds.TokenType = &cruder.Cond{
					Op:  conds.GetTokenType().GetOp(),
					Val: conds.GetTokenType().GetValue(),
				}
			}
		}
		if conds.OfficialContract != nil {
			if conds.OfficialContract.GetValue() != "" {
				h.Conds.OfficialContract = &cruder.Cond{
					Op:  conds.GetOfficialContract().GetOp(),
					Val: conds.GetOfficialContract().GetValue(),
				}
			}
		}
		if conds.PrivateContract != nil {
			if conds.PrivateContract.GetValue() != "" {
				h.Conds.PrivateContract = &cruder.Cond{
					Op:  conds.GetPrivateContract().GetOp(),
					Val: conds.GetPrivateContract().GetValue(),
				}
			}
		}
		if conds.Remark != nil {
			if conds.Remark.GetValue() != "" {
				h.Conds.Remark = &cruder.Cond{
					Op:  conds.GetRemark().GetOp(),
					Val: conds.GetRemark().GetValue(),
				}
			}
		}
		if conds.Name != nil {
			if conds.Name.GetValue() != "" {
				h.Conds.Name = &cruder.Cond{
					Op:  conds.GetName().GetOp(),
					Val: conds.GetName().GetValue(),
				}
			}
		}
		if conds.EntIDs != nil {
			if len(conds.EntIDs.GetValue()) != 0 {
				h.Conds.Data = &cruder.Cond{
					Op:  conds.GetEntIDs().GetOp(),
					Val: conds.GetEntIDs().GetValue(),
				}
			}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
