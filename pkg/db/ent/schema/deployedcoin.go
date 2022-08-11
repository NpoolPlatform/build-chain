package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/build-chain/pkg/db/mixin"
	"github.com/google/uuid"
)

// DeployedCoins holds the schema definition for the DeployedCoins entity.
type DeployedCoin struct {
	ent.Schema
}

func (DeployedCoin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the DeployedCoins.
func (DeployedCoin) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("coin_id", uuid.UUID{}).
			Default(uuid.New),
		field.
			String("contract").
			Optional().
			Default(""),
	}
}

// Edges of the DeployedCoins.
func (DeployedCoin) Edges() []ent.Edge {
	return nil
}
