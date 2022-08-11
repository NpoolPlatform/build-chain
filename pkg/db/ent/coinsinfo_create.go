// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/google/uuid"
)

// CoinsInfoCreate is the builder for creating a CoinsInfo entity.
type CoinsInfoCreate struct {
	config
	mutation *CoinsInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cic *CoinsInfoCreate) SetCreatedAt(u uint32) *CoinsInfoCreate {
	cic.mutation.SetCreatedAt(u)
	return cic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableCreatedAt(u *uint32) *CoinsInfoCreate {
	if u != nil {
		cic.SetCreatedAt(*u)
	}
	return cic
}

// SetUpdatedAt sets the "updated_at" field.
func (cic *CoinsInfoCreate) SetUpdatedAt(u uint32) *CoinsInfoCreate {
	cic.mutation.SetUpdatedAt(u)
	return cic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableUpdatedAt(u *uint32) *CoinsInfoCreate {
	if u != nil {
		cic.SetUpdatedAt(*u)
	}
	return cic
}

// SetDeletedAt sets the "deleted_at" field.
func (cic *CoinsInfoCreate) SetDeletedAt(u uint32) *CoinsInfoCreate {
	cic.mutation.SetDeletedAt(u)
	return cic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableDeletedAt(u *uint32) *CoinsInfoCreate {
	if u != nil {
		cic.SetDeletedAt(*u)
	}
	return cic
}

// SetName sets the "name" field.
func (cic *CoinsInfoCreate) SetName(s string) *CoinsInfoCreate {
	cic.mutation.SetName(s)
	return cic
}

// SetChainType sets the "chain_type" field.
func (cic *CoinsInfoCreate) SetChainType(s string) *CoinsInfoCreate {
	cic.mutation.SetChainType(s)
	return cic
}

// SetNillableChainType sets the "chain_type" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableChainType(s *string) *CoinsInfoCreate {
	if s != nil {
		cic.SetChainType(*s)
	}
	return cic
}

// SetTokenType sets the "token_type" field.
func (cic *CoinsInfoCreate) SetTokenType(s string) *CoinsInfoCreate {
	cic.mutation.SetTokenType(s)
	return cic
}

// SetNillableTokenType sets the "token_type" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableTokenType(s *string) *CoinsInfoCreate {
	if s != nil {
		cic.SetTokenType(*s)
	}
	return cic
}

// SetContract sets the "contract" field.
func (cic *CoinsInfoCreate) SetContract(s string) *CoinsInfoCreate {
	cic.mutation.SetContract(s)
	return cic
}

// SetSimilarity sets the "similarity" field.
func (cic *CoinsInfoCreate) SetSimilarity(i int32) *CoinsInfoCreate {
	cic.mutation.SetSimilarity(i)
	return cic
}

// SetNillableSimilarity sets the "similarity" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableSimilarity(i *int32) *CoinsInfoCreate {
	if i != nil {
		cic.SetSimilarity(*i)
	}
	return cic
}

// SetRemark sets the "remark" field.
func (cic *CoinsInfoCreate) SetRemark(s string) *CoinsInfoCreate {
	cic.mutation.SetRemark(s)
	return cic
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableRemark(s *string) *CoinsInfoCreate {
	if s != nil {
		cic.SetRemark(*s)
	}
	return cic
}

// SetData sets the "data" field.
func (cic *CoinsInfoCreate) SetData(b []byte) *CoinsInfoCreate {
	cic.mutation.SetData(b)
	return cic
}

// SetID sets the "id" field.
func (cic *CoinsInfoCreate) SetID(u uuid.UUID) *CoinsInfoCreate {
	cic.mutation.SetID(u)
	return cic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cic *CoinsInfoCreate) SetNillableID(u *uuid.UUID) *CoinsInfoCreate {
	if u != nil {
		cic.SetID(*u)
	}
	return cic
}

// Mutation returns the CoinsInfoMutation object of the builder.
func (cic *CoinsInfoCreate) Mutation() *CoinsInfoMutation {
	return cic.mutation
}

// Save creates the CoinsInfo in the database.
func (cic *CoinsInfoCreate) Save(ctx context.Context) (*CoinsInfo, error) {
	var (
		err  error
		node *CoinsInfo
	)
	if err := cic.defaults(); err != nil {
		return nil, err
	}
	if len(cic.hooks) == 0 {
		if err = cic.check(); err != nil {
			return nil, err
		}
		node, err = cic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinsInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cic.check(); err != nil {
				return nil, err
			}
			cic.mutation = mutation
			if node, err = cic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cic.hooks) - 1; i >= 0; i-- {
			if cic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cic *CoinsInfoCreate) SaveX(ctx context.Context) *CoinsInfo {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cic *CoinsInfoCreate) Exec(ctx context.Context) error {
	_, err := cic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cic *CoinsInfoCreate) ExecX(ctx context.Context) {
	if err := cic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cic *CoinsInfoCreate) defaults() error {
	if _, ok := cic.mutation.CreatedAt(); !ok {
		if coinsinfo.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinsinfo.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := coinsinfo.DefaultCreatedAt()
		cic.mutation.SetCreatedAt(v)
	}
	if _, ok := cic.mutation.UpdatedAt(); !ok {
		if coinsinfo.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinsinfo.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinsinfo.DefaultUpdatedAt()
		cic.mutation.SetUpdatedAt(v)
	}
	if _, ok := cic.mutation.DeletedAt(); !ok {
		if coinsinfo.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized coinsinfo.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := coinsinfo.DefaultDeletedAt()
		cic.mutation.SetDeletedAt(v)
	}
	if _, ok := cic.mutation.ChainType(); !ok {
		v := coinsinfo.DefaultChainType
		cic.mutation.SetChainType(v)
	}
	if _, ok := cic.mutation.TokenType(); !ok {
		v := coinsinfo.DefaultTokenType
		cic.mutation.SetTokenType(v)
	}
	if _, ok := cic.mutation.Similarity(); !ok {
		v := coinsinfo.DefaultSimilarity
		cic.mutation.SetSimilarity(v)
	}
	if _, ok := cic.mutation.Remark(); !ok {
		v := coinsinfo.DefaultRemark
		cic.mutation.SetRemark(v)
	}
	if _, ok := cic.mutation.ID(); !ok {
		if coinsinfo.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized coinsinfo.DefaultID (forgotten import ent/runtime?)")
		}
		v := coinsinfo.DefaultID()
		cic.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cic *CoinsInfoCreate) check() error {
	if _, ok := cic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CoinsInfo.created_at"`)}
	}
	if _, ok := cic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CoinsInfo.updated_at"`)}
	}
	if _, ok := cic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "CoinsInfo.deleted_at"`)}
	}
	if _, ok := cic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CoinsInfo.name"`)}
	}
	if _, ok := cic.mutation.Contract(); !ok {
		return &ValidationError{Name: "contract", err: errors.New(`ent: missing required field "CoinsInfo.contract"`)}
	}
	return nil
}

func (cic *CoinsInfoCreate) sqlSave(ctx context.Context) (*CoinsInfo, error) {
	_node, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cic *CoinsInfoCreate) createSpec() (*CoinsInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &CoinsInfo{config: cic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coinsinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinsinfo.FieldID,
			},
		}
	)
	_spec.OnConflict = cic.conflict
	if id, ok := cic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cic.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cic.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cic.mutation.ChainType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldChainType,
		})
		_node.ChainType = value
	}
	if value, ok := cic.mutation.TokenType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldTokenType,
		})
		_node.TokenType = value
	}
	if value, ok := cic.mutation.Contract(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldContract,
		})
		_node.Contract = value
	}
	if value, ok := cic.mutation.Similarity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coinsinfo.FieldSimilarity,
		})
		_node.Similarity = value
	}
	if value, ok := cic.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := cic.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: coinsinfo.FieldData,
		})
		_node.Data = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CoinsInfo.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinsInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cic *CoinsInfoCreate) OnConflict(opts ...sql.ConflictOption) *CoinsInfoUpsertOne {
	cic.conflict = opts
	return &CoinsInfoUpsertOne{
		create: cic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CoinsInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cic *CoinsInfoCreate) OnConflictColumns(columns ...string) *CoinsInfoUpsertOne {
	cic.conflict = append(cic.conflict, sql.ConflictColumns(columns...))
	return &CoinsInfoUpsertOne{
		create: cic,
	}
}

type (
	// CoinsInfoUpsertOne is the builder for "upsert"-ing
	//  one CoinsInfo node.
	CoinsInfoUpsertOne struct {
		create *CoinsInfoCreate
	}

	// CoinsInfoUpsert is the "OnConflict" setter.
	CoinsInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CoinsInfoUpsert) SetCreatedAt(v uint32) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateCreatedAt() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinsInfoUpsert) AddCreatedAt(v uint32) *CoinsInfoUpsert {
	u.Add(coinsinfo.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinsInfoUpsert) SetUpdatedAt(v uint32) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateUpdatedAt() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinsInfoUpsert) AddUpdatedAt(v uint32) *CoinsInfoUpsert {
	u.Add(coinsinfo.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinsInfoUpsert) SetDeletedAt(v uint32) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateDeletedAt() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinsInfoUpsert) AddDeletedAt(v uint32) *CoinsInfoUpsert {
	u.Add(coinsinfo.FieldDeletedAt, v)
	return u
}

// SetName sets the "name" field.
func (u *CoinsInfoUpsert) SetName(v string) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateName() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldName)
	return u
}

// SetChainType sets the "chain_type" field.
func (u *CoinsInfoUpsert) SetChainType(v string) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldChainType, v)
	return u
}

// UpdateChainType sets the "chain_type" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateChainType() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldChainType)
	return u
}

// ClearChainType clears the value of the "chain_type" field.
func (u *CoinsInfoUpsert) ClearChainType() *CoinsInfoUpsert {
	u.SetNull(coinsinfo.FieldChainType)
	return u
}

// SetTokenType sets the "token_type" field.
func (u *CoinsInfoUpsert) SetTokenType(v string) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldTokenType, v)
	return u
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateTokenType() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldTokenType)
	return u
}

// ClearTokenType clears the value of the "token_type" field.
func (u *CoinsInfoUpsert) ClearTokenType() *CoinsInfoUpsert {
	u.SetNull(coinsinfo.FieldTokenType)
	return u
}

// SetContract sets the "contract" field.
func (u *CoinsInfoUpsert) SetContract(v string) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldContract, v)
	return u
}

// UpdateContract sets the "contract" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateContract() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldContract)
	return u
}

// SetSimilarity sets the "similarity" field.
func (u *CoinsInfoUpsert) SetSimilarity(v int32) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldSimilarity, v)
	return u
}

// UpdateSimilarity sets the "similarity" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateSimilarity() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldSimilarity)
	return u
}

// AddSimilarity adds v to the "similarity" field.
func (u *CoinsInfoUpsert) AddSimilarity(v int32) *CoinsInfoUpsert {
	u.Add(coinsinfo.FieldSimilarity, v)
	return u
}

// ClearSimilarity clears the value of the "similarity" field.
func (u *CoinsInfoUpsert) ClearSimilarity() *CoinsInfoUpsert {
	u.SetNull(coinsinfo.FieldSimilarity)
	return u
}

// SetRemark sets the "remark" field.
func (u *CoinsInfoUpsert) SetRemark(v string) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateRemark() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *CoinsInfoUpsert) ClearRemark() *CoinsInfoUpsert {
	u.SetNull(coinsinfo.FieldRemark)
	return u
}

// SetData sets the "data" field.
func (u *CoinsInfoUpsert) SetData(v []byte) *CoinsInfoUpsert {
	u.Set(coinsinfo.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *CoinsInfoUpsert) UpdateData() *CoinsInfoUpsert {
	u.SetExcluded(coinsinfo.FieldData)
	return u
}

// ClearData clears the value of the "data" field.
func (u *CoinsInfoUpsert) ClearData() *CoinsInfoUpsert {
	u.SetNull(coinsinfo.FieldData)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CoinsInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coinsinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinsInfoUpsertOne) UpdateNewValues() *CoinsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(coinsinfo.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CoinsInfo.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CoinsInfoUpsertOne) Ignore() *CoinsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinsInfoUpsertOne) DoNothing() *CoinsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinsInfoCreate.OnConflict
// documentation for more info.
func (u *CoinsInfoUpsertOne) Update(set func(*CoinsInfoUpsert)) *CoinsInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinsInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinsInfoUpsertOne) SetCreatedAt(v uint32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinsInfoUpsertOne) AddCreatedAt(v uint32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateCreatedAt() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinsInfoUpsertOne) SetUpdatedAt(v uint32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinsInfoUpsertOne) AddUpdatedAt(v uint32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateUpdatedAt() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinsInfoUpsertOne) SetDeletedAt(v uint32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinsInfoUpsertOne) AddDeletedAt(v uint32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateDeletedAt() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *CoinsInfoUpsertOne) SetName(v string) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateName() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateName()
	})
}

// SetChainType sets the "chain_type" field.
func (u *CoinsInfoUpsertOne) SetChainType(v string) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetChainType(v)
	})
}

// UpdateChainType sets the "chain_type" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateChainType() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateChainType()
	})
}

// ClearChainType clears the value of the "chain_type" field.
func (u *CoinsInfoUpsertOne) ClearChainType() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearChainType()
	})
}

// SetTokenType sets the "token_type" field.
func (u *CoinsInfoUpsertOne) SetTokenType(v string) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetTokenType(v)
	})
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateTokenType() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateTokenType()
	})
}

// ClearTokenType clears the value of the "token_type" field.
func (u *CoinsInfoUpsertOne) ClearTokenType() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearTokenType()
	})
}

// SetContract sets the "contract" field.
func (u *CoinsInfoUpsertOne) SetContract(v string) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetContract(v)
	})
}

// UpdateContract sets the "contract" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateContract() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateContract()
	})
}

// SetSimilarity sets the "similarity" field.
func (u *CoinsInfoUpsertOne) SetSimilarity(v int32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetSimilarity(v)
	})
}

// AddSimilarity adds v to the "similarity" field.
func (u *CoinsInfoUpsertOne) AddSimilarity(v int32) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddSimilarity(v)
	})
}

// UpdateSimilarity sets the "similarity" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateSimilarity() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateSimilarity()
	})
}

// ClearSimilarity clears the value of the "similarity" field.
func (u *CoinsInfoUpsertOne) ClearSimilarity() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearSimilarity()
	})
}

// SetRemark sets the "remark" field.
func (u *CoinsInfoUpsertOne) SetRemark(v string) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateRemark() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *CoinsInfoUpsertOne) ClearRemark() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearRemark()
	})
}

// SetData sets the "data" field.
func (u *CoinsInfoUpsertOne) SetData(v []byte) *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *CoinsInfoUpsertOne) UpdateData() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *CoinsInfoUpsertOne) ClearData() *CoinsInfoUpsertOne {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearData()
	})
}

// Exec executes the query.
func (u *CoinsInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinsInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinsInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CoinsInfoUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CoinsInfoUpsertOne.ID is not supported by MySQL driver. Use CoinsInfoUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CoinsInfoUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CoinsInfoCreateBulk is the builder for creating many CoinsInfo entities in bulk.
type CoinsInfoCreateBulk struct {
	config
	builders []*CoinsInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the CoinsInfo entities in the database.
func (cicb *CoinsInfoCreateBulk) Save(ctx context.Context) ([]*CoinsInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cicb.builders))
	nodes := make([]*CoinsInfo, len(cicb.builders))
	mutators := make([]Mutator, len(cicb.builders))
	for i := range cicb.builders {
		func(i int, root context.Context) {
			builder := cicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CoinsInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cicb *CoinsInfoCreateBulk) SaveX(ctx context.Context) []*CoinsInfo {
	v, err := cicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cicb *CoinsInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := cicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicb *CoinsInfoCreateBulk) ExecX(ctx context.Context) {
	if err := cicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CoinsInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinsInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cicb *CoinsInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *CoinsInfoUpsertBulk {
	cicb.conflict = opts
	return &CoinsInfoUpsertBulk{
		create: cicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CoinsInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cicb *CoinsInfoCreateBulk) OnConflictColumns(columns ...string) *CoinsInfoUpsertBulk {
	cicb.conflict = append(cicb.conflict, sql.ConflictColumns(columns...))
	return &CoinsInfoUpsertBulk{
		create: cicb,
	}
}

// CoinsInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of CoinsInfo nodes.
type CoinsInfoUpsertBulk struct {
	create *CoinsInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CoinsInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coinsinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinsInfoUpsertBulk) UpdateNewValues() *CoinsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(coinsinfo.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CoinsInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CoinsInfoUpsertBulk) Ignore() *CoinsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinsInfoUpsertBulk) DoNothing() *CoinsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinsInfoCreateBulk.OnConflict
// documentation for more info.
func (u *CoinsInfoUpsertBulk) Update(set func(*CoinsInfoUpsert)) *CoinsInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinsInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CoinsInfoUpsertBulk) SetCreatedAt(v uint32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CoinsInfoUpsertBulk) AddCreatedAt(v uint32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateCreatedAt() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CoinsInfoUpsertBulk) SetUpdatedAt(v uint32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CoinsInfoUpsertBulk) AddUpdatedAt(v uint32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateUpdatedAt() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CoinsInfoUpsertBulk) SetDeletedAt(v uint32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CoinsInfoUpsertBulk) AddDeletedAt(v uint32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateDeletedAt() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *CoinsInfoUpsertBulk) SetName(v string) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateName() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateName()
	})
}

// SetChainType sets the "chain_type" field.
func (u *CoinsInfoUpsertBulk) SetChainType(v string) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetChainType(v)
	})
}

// UpdateChainType sets the "chain_type" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateChainType() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateChainType()
	})
}

// ClearChainType clears the value of the "chain_type" field.
func (u *CoinsInfoUpsertBulk) ClearChainType() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearChainType()
	})
}

// SetTokenType sets the "token_type" field.
func (u *CoinsInfoUpsertBulk) SetTokenType(v string) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetTokenType(v)
	})
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateTokenType() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateTokenType()
	})
}

// ClearTokenType clears the value of the "token_type" field.
func (u *CoinsInfoUpsertBulk) ClearTokenType() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearTokenType()
	})
}

// SetContract sets the "contract" field.
func (u *CoinsInfoUpsertBulk) SetContract(v string) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetContract(v)
	})
}

// UpdateContract sets the "contract" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateContract() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateContract()
	})
}

// SetSimilarity sets the "similarity" field.
func (u *CoinsInfoUpsertBulk) SetSimilarity(v int32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetSimilarity(v)
	})
}

// AddSimilarity adds v to the "similarity" field.
func (u *CoinsInfoUpsertBulk) AddSimilarity(v int32) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.AddSimilarity(v)
	})
}

// UpdateSimilarity sets the "similarity" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateSimilarity() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateSimilarity()
	})
}

// ClearSimilarity clears the value of the "similarity" field.
func (u *CoinsInfoUpsertBulk) ClearSimilarity() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearSimilarity()
	})
}

// SetRemark sets the "remark" field.
func (u *CoinsInfoUpsertBulk) SetRemark(v string) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateRemark() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *CoinsInfoUpsertBulk) ClearRemark() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearRemark()
	})
}

// SetData sets the "data" field.
func (u *CoinsInfoUpsertBulk) SetData(v []byte) *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *CoinsInfoUpsertBulk) UpdateData() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.UpdateData()
	})
}

// ClearData clears the value of the "data" field.
func (u *CoinsInfoUpsertBulk) ClearData() *CoinsInfoUpsertBulk {
	return u.Update(func(s *CoinsInfoUpsert) {
		s.ClearData()
	})
}

// Exec executes the query.
func (u *CoinsInfoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CoinsInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinsInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinsInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
