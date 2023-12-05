package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/build-chain/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

type TokenInfo struct {
	ent.Schema
}

func (TokenInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		mixin.TimeMixin{},
	}
}

// Fields of the DeployedCoins.
func (TokenInfo) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name"),
		field.
			String("chain_type").
			Optional().
			Default(""),
		field.
			String("token_type").
			Optional().
			Default(""),
		field.
			String("unit").
			Optional().
			Default(""),
		field.
			String("decimal").
			Optional().
			Default(""),
		field.
			String("official_contract").
			Unique(),
		field.
			String("private_contract").
			Unique(),
		field.
			String("remark").
			Optional().
			Default(""),
		field.
			Bytes("data").
			Optional(),
	}
}

// Edges of the DeployedCoins.
func (TokenInfo) Edges() []ent.Edge {
	return nil
}
