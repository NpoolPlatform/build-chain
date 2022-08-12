// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/google/uuid"
)

// CoinsInfo is the model entity for the CoinsInfo schema.
type CoinsInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ChainType holds the value of the "chain_type" field.
	ChainType string `json:"chain_type,omitempty"`
	// TokenType holds the value of the "token_type" field.
	TokenType string `json:"token_type,omitempty"`
	// Contract holds the value of the "contract" field.
	Contract string `json:"contract,omitempty"`
	// Similarity holds the value of the "similarity" field.
	Similarity int32 `json:"similarity,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// Data holds the value of the "data" field.
	Data []byte `json:"data,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinsInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinsinfo.FieldData:
			values[i] = new([]byte)
		case coinsinfo.FieldCreatedAt, coinsinfo.FieldUpdatedAt, coinsinfo.FieldDeletedAt, coinsinfo.FieldSimilarity:
			values[i] = new(sql.NullInt64)
		case coinsinfo.FieldName, coinsinfo.FieldChainType, coinsinfo.FieldTokenType, coinsinfo.FieldContract, coinsinfo.FieldRemark:
			values[i] = new(sql.NullString)
		case coinsinfo.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinsInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinsInfo fields.
func (ci *CoinsInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinsinfo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ci.ID = *value
			}
		case coinsinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ci.CreatedAt = uint32(value.Int64)
			}
		case coinsinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ci.UpdatedAt = uint32(value.Int64)
			}
		case coinsinfo.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ci.DeletedAt = uint32(value.Int64)
			}
		case coinsinfo.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ci.Name = value.String
			}
		case coinsinfo.FieldChainType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_type", values[i])
			} else if value.Valid {
				ci.ChainType = value.String
			}
		case coinsinfo.FieldTokenType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_type", values[i])
			} else if value.Valid {
				ci.TokenType = value.String
			}
		case coinsinfo.FieldContract:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract", values[i])
			} else if value.Valid {
				ci.Contract = value.String
			}
		case coinsinfo.FieldSimilarity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field similarity", values[i])
			} else if value.Valid {
				ci.Similarity = int32(value.Int64)
			}
		case coinsinfo.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				ci.Remark = value.String
			}
		case coinsinfo.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil {
				ci.Data = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinsInfo.
// Note that you need to call CoinsInfo.Unwrap() before calling this method if this CoinsInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ci *CoinsInfo) Update() *CoinsInfoUpdateOne {
	return (&CoinsInfoClient{config: ci.config}).UpdateOne(ci)
}

// Unwrap unwraps the CoinsInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ci *CoinsInfo) Unwrap() *CoinsInfo {
	_tx, ok := ci.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinsInfo is not a transactional entity")
	}
	ci.config.driver = _tx.drv
	return ci
}

// String implements the fmt.Stringer.
func (ci *CoinsInfo) String() string {
	var builder strings.Builder
	builder.WriteString("CoinsInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ci.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ci.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ci.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ci.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ci.Name)
	builder.WriteString(", ")
	builder.WriteString("chain_type=")
	builder.WriteString(ci.ChainType)
	builder.WriteString(", ")
	builder.WriteString("token_type=")
	builder.WriteString(ci.TokenType)
	builder.WriteString(", ")
	builder.WriteString("contract=")
	builder.WriteString(ci.Contract)
	builder.WriteString(", ")
	builder.WriteString("similarity=")
	builder.WriteString(fmt.Sprintf("%v", ci.Similarity))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(ci.Remark)
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", ci.Data))
	builder.WriteByte(')')
	return builder.String()
}

// CoinsInfos is a parsable slice of CoinsInfo.
type CoinsInfos []*CoinsInfo

func (ci CoinsInfos) config(cfg config) {
	for _i := range ci {
		ci[_i].config = cfg
	}
}
