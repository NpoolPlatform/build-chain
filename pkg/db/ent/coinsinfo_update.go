// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/predicate"
)

// CoinsInfoUpdate is the builder for updating CoinsInfo entities.
type CoinsInfoUpdate struct {
	config
	hooks    []Hook
	mutation *CoinsInfoMutation
}

// Where appends a list predicates to the CoinsInfoUpdate builder.
func (ciu *CoinsInfoUpdate) Where(ps ...predicate.CoinsInfo) *CoinsInfoUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetCreatedAt sets the "created_at" field.
func (ciu *CoinsInfoUpdate) SetCreatedAt(u uint32) *CoinsInfoUpdate {
	ciu.mutation.ResetCreatedAt()
	ciu.mutation.SetCreatedAt(u)
	return ciu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciu *CoinsInfoUpdate) SetNillableCreatedAt(u *uint32) *CoinsInfoUpdate {
	if u != nil {
		ciu.SetCreatedAt(*u)
	}
	return ciu
}

// AddCreatedAt adds u to the "created_at" field.
func (ciu *CoinsInfoUpdate) AddCreatedAt(u int32) *CoinsInfoUpdate {
	ciu.mutation.AddCreatedAt(u)
	return ciu
}

// SetUpdatedAt sets the "updated_at" field.
func (ciu *CoinsInfoUpdate) SetUpdatedAt(u uint32) *CoinsInfoUpdate {
	ciu.mutation.ResetUpdatedAt()
	ciu.mutation.SetUpdatedAt(u)
	return ciu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ciu *CoinsInfoUpdate) AddUpdatedAt(u int32) *CoinsInfoUpdate {
	ciu.mutation.AddUpdatedAt(u)
	return ciu
}

// SetDeletedAt sets the "deleted_at" field.
func (ciu *CoinsInfoUpdate) SetDeletedAt(u uint32) *CoinsInfoUpdate {
	ciu.mutation.ResetDeletedAt()
	ciu.mutation.SetDeletedAt(u)
	return ciu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciu *CoinsInfoUpdate) SetNillableDeletedAt(u *uint32) *CoinsInfoUpdate {
	if u != nil {
		ciu.SetDeletedAt(*u)
	}
	return ciu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ciu *CoinsInfoUpdate) AddDeletedAt(u int32) *CoinsInfoUpdate {
	ciu.mutation.AddDeletedAt(u)
	return ciu
}

// SetName sets the "name" field.
func (ciu *CoinsInfoUpdate) SetName(s string) *CoinsInfoUpdate {
	ciu.mutation.SetName(s)
	return ciu
}

// SetChainType sets the "chain_type" field.
func (ciu *CoinsInfoUpdate) SetChainType(s string) *CoinsInfoUpdate {
	ciu.mutation.SetChainType(s)
	return ciu
}

// SetNillableChainType sets the "chain_type" field if the given value is not nil.
func (ciu *CoinsInfoUpdate) SetNillableChainType(s *string) *CoinsInfoUpdate {
	if s != nil {
		ciu.SetChainType(*s)
	}
	return ciu
}

// ClearChainType clears the value of the "chain_type" field.
func (ciu *CoinsInfoUpdate) ClearChainType() *CoinsInfoUpdate {
	ciu.mutation.ClearChainType()
	return ciu
}

// SetTokenType sets the "token_type" field.
func (ciu *CoinsInfoUpdate) SetTokenType(s string) *CoinsInfoUpdate {
	ciu.mutation.SetTokenType(s)
	return ciu
}

// SetNillableTokenType sets the "token_type" field if the given value is not nil.
func (ciu *CoinsInfoUpdate) SetNillableTokenType(s *string) *CoinsInfoUpdate {
	if s != nil {
		ciu.SetTokenType(*s)
	}
	return ciu
}

// ClearTokenType clears the value of the "token_type" field.
func (ciu *CoinsInfoUpdate) ClearTokenType() *CoinsInfoUpdate {
	ciu.mutation.ClearTokenType()
	return ciu
}

// SetContract sets the "contract" field.
func (ciu *CoinsInfoUpdate) SetContract(s string) *CoinsInfoUpdate {
	ciu.mutation.SetContract(s)
	return ciu
}

// SetSimilarity sets the "similarity" field.
func (ciu *CoinsInfoUpdate) SetSimilarity(i int32) *CoinsInfoUpdate {
	ciu.mutation.ResetSimilarity()
	ciu.mutation.SetSimilarity(i)
	return ciu
}

// SetNillableSimilarity sets the "similarity" field if the given value is not nil.
func (ciu *CoinsInfoUpdate) SetNillableSimilarity(i *int32) *CoinsInfoUpdate {
	if i != nil {
		ciu.SetSimilarity(*i)
	}
	return ciu
}

// AddSimilarity adds i to the "similarity" field.
func (ciu *CoinsInfoUpdate) AddSimilarity(i int32) *CoinsInfoUpdate {
	ciu.mutation.AddSimilarity(i)
	return ciu
}

// ClearSimilarity clears the value of the "similarity" field.
func (ciu *CoinsInfoUpdate) ClearSimilarity() *CoinsInfoUpdate {
	ciu.mutation.ClearSimilarity()
	return ciu
}

// SetRemark sets the "remark" field.
func (ciu *CoinsInfoUpdate) SetRemark(s string) *CoinsInfoUpdate {
	ciu.mutation.SetRemark(s)
	return ciu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ciu *CoinsInfoUpdate) SetNillableRemark(s *string) *CoinsInfoUpdate {
	if s != nil {
		ciu.SetRemark(*s)
	}
	return ciu
}

// ClearRemark clears the value of the "remark" field.
func (ciu *CoinsInfoUpdate) ClearRemark() *CoinsInfoUpdate {
	ciu.mutation.ClearRemark()
	return ciu
}

// SetData sets the "data" field.
func (ciu *CoinsInfoUpdate) SetData(b []byte) *CoinsInfoUpdate {
	ciu.mutation.SetData(b)
	return ciu
}

// ClearData clears the value of the "data" field.
func (ciu *CoinsInfoUpdate) ClearData() *CoinsInfoUpdate {
	ciu.mutation.ClearData()
	return ciu
}

// Mutation returns the CoinsInfoMutation object of the builder.
func (ciu *CoinsInfoUpdate) Mutation() *CoinsInfoMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CoinsInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ciu.defaults(); err != nil {
		return 0, err
	}
	if len(ciu.hooks) == 0 {
		affected, err = ciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinsInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ciu.mutation = mutation
			affected, err = ciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ciu.hooks) - 1; i >= 0; i-- {
			if ciu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CoinsInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CoinsInfoUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CoinsInfoUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciu *CoinsInfoUpdate) defaults() error {
	if _, ok := ciu.mutation.UpdatedAt(); !ok {
		if coinsinfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinsinfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinsinfo.UpdateDefaultUpdatedAt()
		ciu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ciu *CoinsInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinsinfo.Table,
			Columns: coinsinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinsinfo.FieldID,
			},
		},
	}
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldCreatedAt,
		})
	}
	if value, ok := ciu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldCreatedAt,
		})
	}
	if value, ok := ciu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldDeletedAt,
		})
	}
	if value, ok := ciu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldDeletedAt,
		})
	}
	if value, ok := ciu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldName,
		})
	}
	if value, ok := ciu.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldChainType,
		})
	}
	if ciu.mutation.ChainTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinsinfo.FieldChainType,
		})
	}
	if value, ok := ciu.mutation.TokenType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldTokenType,
		})
	}
	if ciu.mutation.TokenTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinsinfo.FieldTokenType,
		})
	}
	if value, ok := ciu.mutation.Contract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldContract,
		})
	}
	if value, ok := ciu.mutation.Similarity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coinsinfo.FieldSimilarity,
		})
	}
	if value, ok := ciu.mutation.AddedSimilarity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coinsinfo.FieldSimilarity,
		})
	}
	if ciu.mutation.SimilarityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: coinsinfo.FieldSimilarity,
		})
	}
	if value, ok := ciu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldRemark,
		})
	}
	if ciu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinsinfo.FieldRemark,
		})
	}
	if value, ok := ciu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: coinsinfo.FieldData,
		})
	}
	if ciu.mutation.DataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: coinsinfo.FieldData,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinsinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CoinsInfoUpdateOne is the builder for updating a single CoinsInfo entity.
type CoinsInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CoinsInfoMutation
}

// SetCreatedAt sets the "created_at" field.
func (ciuo *CoinsInfoUpdateOne) SetCreatedAt(u uint32) *CoinsInfoUpdateOne {
	ciuo.mutation.ResetCreatedAt()
	ciuo.mutation.SetCreatedAt(u)
	return ciuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciuo *CoinsInfoUpdateOne) SetNillableCreatedAt(u *uint32) *CoinsInfoUpdateOne {
	if u != nil {
		ciuo.SetCreatedAt(*u)
	}
	return ciuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ciuo *CoinsInfoUpdateOne) AddCreatedAt(u int32) *CoinsInfoUpdateOne {
	ciuo.mutation.AddCreatedAt(u)
	return ciuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuo *CoinsInfoUpdateOne) SetUpdatedAt(u uint32) *CoinsInfoUpdateOne {
	ciuo.mutation.ResetUpdatedAt()
	ciuo.mutation.SetUpdatedAt(u)
	return ciuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ciuo *CoinsInfoUpdateOne) AddUpdatedAt(u int32) *CoinsInfoUpdateOne {
	ciuo.mutation.AddUpdatedAt(u)
	return ciuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ciuo *CoinsInfoUpdateOne) SetDeletedAt(u uint32) *CoinsInfoUpdateOne {
	ciuo.mutation.ResetDeletedAt()
	ciuo.mutation.SetDeletedAt(u)
	return ciuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciuo *CoinsInfoUpdateOne) SetNillableDeletedAt(u *uint32) *CoinsInfoUpdateOne {
	if u != nil {
		ciuo.SetDeletedAt(*u)
	}
	return ciuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ciuo *CoinsInfoUpdateOne) AddDeletedAt(u int32) *CoinsInfoUpdateOne {
	ciuo.mutation.AddDeletedAt(u)
	return ciuo
}

// SetName sets the "name" field.
func (ciuo *CoinsInfoUpdateOne) SetName(s string) *CoinsInfoUpdateOne {
	ciuo.mutation.SetName(s)
	return ciuo
}

// SetChainType sets the "chain_type" field.
func (ciuo *CoinsInfoUpdateOne) SetChainType(s string) *CoinsInfoUpdateOne {
	ciuo.mutation.SetChainType(s)
	return ciuo
}

// SetNillableChainType sets the "chain_type" field if the given value is not nil.
func (ciuo *CoinsInfoUpdateOne) SetNillableChainType(s *string) *CoinsInfoUpdateOne {
	if s != nil {
		ciuo.SetChainType(*s)
	}
	return ciuo
}

// ClearChainType clears the value of the "chain_type" field.
func (ciuo *CoinsInfoUpdateOne) ClearChainType() *CoinsInfoUpdateOne {
	ciuo.mutation.ClearChainType()
	return ciuo
}

// SetTokenType sets the "token_type" field.
func (ciuo *CoinsInfoUpdateOne) SetTokenType(s string) *CoinsInfoUpdateOne {
	ciuo.mutation.SetTokenType(s)
	return ciuo
}

// SetNillableTokenType sets the "token_type" field if the given value is not nil.
func (ciuo *CoinsInfoUpdateOne) SetNillableTokenType(s *string) *CoinsInfoUpdateOne {
	if s != nil {
		ciuo.SetTokenType(*s)
	}
	return ciuo
}

// ClearTokenType clears the value of the "token_type" field.
func (ciuo *CoinsInfoUpdateOne) ClearTokenType() *CoinsInfoUpdateOne {
	ciuo.mutation.ClearTokenType()
	return ciuo
}

// SetContract sets the "contract" field.
func (ciuo *CoinsInfoUpdateOne) SetContract(s string) *CoinsInfoUpdateOne {
	ciuo.mutation.SetContract(s)
	return ciuo
}

// SetSimilarity sets the "similarity" field.
func (ciuo *CoinsInfoUpdateOne) SetSimilarity(i int32) *CoinsInfoUpdateOne {
	ciuo.mutation.ResetSimilarity()
	ciuo.mutation.SetSimilarity(i)
	return ciuo
}

// SetNillableSimilarity sets the "similarity" field if the given value is not nil.
func (ciuo *CoinsInfoUpdateOne) SetNillableSimilarity(i *int32) *CoinsInfoUpdateOne {
	if i != nil {
		ciuo.SetSimilarity(*i)
	}
	return ciuo
}

// AddSimilarity adds i to the "similarity" field.
func (ciuo *CoinsInfoUpdateOne) AddSimilarity(i int32) *CoinsInfoUpdateOne {
	ciuo.mutation.AddSimilarity(i)
	return ciuo
}

// ClearSimilarity clears the value of the "similarity" field.
func (ciuo *CoinsInfoUpdateOne) ClearSimilarity() *CoinsInfoUpdateOne {
	ciuo.mutation.ClearSimilarity()
	return ciuo
}

// SetRemark sets the "remark" field.
func (ciuo *CoinsInfoUpdateOne) SetRemark(s string) *CoinsInfoUpdateOne {
	ciuo.mutation.SetRemark(s)
	return ciuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ciuo *CoinsInfoUpdateOne) SetNillableRemark(s *string) *CoinsInfoUpdateOne {
	if s != nil {
		ciuo.SetRemark(*s)
	}
	return ciuo
}

// ClearRemark clears the value of the "remark" field.
func (ciuo *CoinsInfoUpdateOne) ClearRemark() *CoinsInfoUpdateOne {
	ciuo.mutation.ClearRemark()
	return ciuo
}

// SetData sets the "data" field.
func (ciuo *CoinsInfoUpdateOne) SetData(b []byte) *CoinsInfoUpdateOne {
	ciuo.mutation.SetData(b)
	return ciuo
}

// ClearData clears the value of the "data" field.
func (ciuo *CoinsInfoUpdateOne) ClearData() *CoinsInfoUpdateOne {
	ciuo.mutation.ClearData()
	return ciuo
}

// Mutation returns the CoinsInfoMutation object of the builder.
func (ciuo *CoinsInfoUpdateOne) Mutation() *CoinsInfoMutation {
	return ciuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CoinsInfoUpdateOne) Select(field string, fields ...string) *CoinsInfoUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CoinsInfo entity.
func (ciuo *CoinsInfoUpdateOne) Save(ctx context.Context) (*CoinsInfo, error) {
	var (
		err  error
		node *CoinsInfo
	)
	if err := ciuo.defaults(); err != nil {
		return nil, err
	}
	if len(ciuo.hooks) == 0 {
		node, err = ciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinsInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ciuo.mutation = mutation
			node, err = ciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ciuo.hooks) - 1; i >= 0; i-- {
			if ciuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ciuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CoinsInfoUpdateOne) SaveX(ctx context.Context) *CoinsInfo {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CoinsInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CoinsInfoUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciuo *CoinsInfoUpdateOne) defaults() error {
	if _, ok := ciuo.mutation.UpdatedAt(); !ok {
		if coinsinfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinsinfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinsinfo.UpdateDefaultUpdatedAt()
		ciuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ciuo *CoinsInfoUpdateOne) sqlSave(ctx context.Context) (_node *CoinsInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinsinfo.Table,
			Columns: coinsinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinsinfo.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinsInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinsinfo.FieldID)
		for _, f := range fields {
			if !coinsinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coinsinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldCreatedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldCreatedAt,
		})
	}
	if value, ok := ciuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldDeletedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinsinfo.FieldDeletedAt,
		})
	}
	if value, ok := ciuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldName,
		})
	}
	if value, ok := ciuo.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldChainType,
		})
	}
	if ciuo.mutation.ChainTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinsinfo.FieldChainType,
		})
	}
	if value, ok := ciuo.mutation.TokenType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldTokenType,
		})
	}
	if ciuo.mutation.TokenTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinsinfo.FieldTokenType,
		})
	}
	if value, ok := ciuo.mutation.Contract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldContract,
		})
	}
	if value, ok := ciuo.mutation.Similarity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coinsinfo.FieldSimilarity,
		})
	}
	if value, ok := ciuo.mutation.AddedSimilarity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coinsinfo.FieldSimilarity,
		})
	}
	if ciuo.mutation.SimilarityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: coinsinfo.FieldSimilarity,
		})
	}
	if value, ok := ciuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinsinfo.FieldRemark,
		})
	}
	if ciuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinsinfo.FieldRemark,
		})
	}
	if value, ok := ciuo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: coinsinfo.FieldData,
		})
	}
	if ciuo.mutation.DataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: coinsinfo.FieldData,
		})
	}
	_node = &CoinsInfo{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinsinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
