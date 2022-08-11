package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/build-chain/pkg/db/mixin"
	"github.com/google/uuid"
)

// DeployedCoins holds the schema definition for the DeployedCoins entity.
type CoinsInfo struct {
	ent.Schema
}

func (CoinsInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the DeployedCoins.
func (CoinsInfo) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
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
			String("contract").
			Unique(),
		field.
			Int32("similarity").
			Optional().
			Default(0),
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
func (CoinsInfo) Edges() []ent.Edge {
	return nil
}
