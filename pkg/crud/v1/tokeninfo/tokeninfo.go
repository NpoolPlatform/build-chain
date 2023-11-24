package tokeninfo

import (
	"fmt"

	"github.com/NpoolPlatform/build-chain/pkg/db/ent"
	enttokeninfo "github.com/NpoolPlatform/build-chain/pkg/db/ent/tokeninfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	EntID            *uuid.UUID
	ChainType        *string
	Name             *string
	Unit             *string
	Decimal          *string
	TokenType        *string
	OfficialContract *string
	PrivateContract  *string
	Remark           *string
	Data             []byte
}

func CreateSet(c *ent.TokenInfoCreate, req *Req) *ent.TokenInfoCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.ChainType != nil {
		c.SetChainType(*req.ChainType)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Unit != nil {
		c.SetUnit(*req.Unit)
	}
	if req.Decimal != nil {
		c.SetDecimal(*req.Decimal)
	}
	if req.TokenType != nil {
		c.SetTokenType(*req.TokenType)
	}
	if req.OfficialContract != nil {
		c.SetOfficialContract(*req.OfficialContract)
	}
	if req.PrivateContract != nil {
		c.SetPrivateContract(*req.PrivateContract)
	}
	if req.Remark != nil {
		c.SetRemark(*req.Remark)
	}
	if req.Data != nil {
		c.SetData(req.Data)
	}
	return c
}

func UpdateSet(u *ent.TokenInfoUpdateOne, req *Req) (*ent.TokenInfoUpdateOne, error) {
	if req.PrivateContract != nil {
		u = u.SetPrivateContract(*req.PrivateContract)
	}
	if req.Remark != nil {
		u = u.SetRemark(*req.Remark)
	}
	return u, nil
}

type Conds struct {
	EntID            *cruder.Cond
	ChainType        *cruder.Cond
	Name             *cruder.Cond
	Unit             *cruder.Cond
	Decimal          *cruder.Cond
	TokenType        *cruder.Cond
	OfficialContract *cruder.Cond
	PrivateContract  *cruder.Cond
	Remark           *cruder.Cond
	Data             *cruder.Cond
}

func SetQueryConds(q *ent.TokenInfoQuery, conds *Conds) (*ent.TokenInfoQuery, error) {
	if conds.EntID != nil {
		ent_id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ent_id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.EntID(ent_id))
		default:
			return nil, fmt.Errorf("invalid ent_id field")
		}
	}
	if conds.ChainType != nil {
		chain_type, ok := conds.ChainType.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid chain_type")
		}
		switch conds.ChainType.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.ChainType(chain_type))
		default:
			return nil, fmt.Errorf("invalid chain_type field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.Name(name))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.Unit != nil {
		unit, ok := conds.Unit.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid unit")
		}
		switch conds.Unit.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.Unit(unit))
		default:
			return nil, fmt.Errorf("invalid unit field")
		}
	}
	if conds.Decimal != nil {
		decimal, ok := conds.Decimal.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid decimal")
		}
		switch conds.Decimal.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.Decimal(decimal))
		default:
			return nil, fmt.Errorf("invalid decimal field")
		}
	}
	if conds.TokenType != nil {
		token_type, ok := conds.TokenType.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid token_type")
		}
		switch conds.TokenType.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.TokenType(token_type))
		default:
			return nil, fmt.Errorf("invalid token_type field")
		}
	}
	if conds.OfficialContract != nil {
		official_contract, ok := conds.OfficialContract.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid official_contract")
		}
		switch conds.OfficialContract.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.OfficialContract(official_contract))
		default:
			return nil, fmt.Errorf("invalid official_contract field")
		}
	}
	if conds.PrivateContract != nil {
		private_contract, ok := conds.PrivateContract.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid private_contract")
		}
		switch conds.PrivateContract.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.PrivateContract(private_contract))
		default:
			return nil, fmt.Errorf("invalid private_contract field")
		}
	}
	if conds.Remark != nil {
		remark, ok := conds.Remark.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid remark")
		}
		switch conds.Remark.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.Remark(remark))
		default:
			return nil, fmt.Errorf("invalid remark field")
		}
	}
	if conds.Data != nil {
		data, ok := conds.Data.Val.([]byte)
		if !ok {
			return nil, fmt.Errorf("invalid data")
		}
		switch conds.Data.Op {
		case cruder.EQ:
			q.Where(enttokeninfo.Data(data))
		default:
			return nil, fmt.Errorf("invalid data field")
		}
	}
	return q, nil
}
