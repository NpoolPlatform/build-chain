// Code generated by entc, DO NOT EDIT.

package deployedcoin

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deployedcoin type in the database.
	Label = "deployed_coin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCoinID holds the string denoting the coin_id field in the database.
	FieldCoinID = "coin_id"
	// FieldContract holds the string denoting the contract field in the database.
	FieldContract = "contract"
	// Table holds the table name of the deployedcoin in the database.
	Table = "deployed_coins"
)

// Columns holds all SQL columns for deployedcoin fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCoinID,
	FieldContract,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/build-chain/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultCoinID holds the default value on creation for the "coin_id" field.
	DefaultCoinID func() uuid.UUID
	// DefaultContract holds the default value on creation for the "contract" field.
	DefaultContract string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
