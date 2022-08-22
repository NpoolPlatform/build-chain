// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/predicate"
)

// CoinInfoUpdate is the builder for updating CoinInfo entities.
type CoinInfoUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinInfoUpdate builder.
func (ciu *CoinInfoUpdate) Where(ps ...predicate.CoinInfo) *CoinInfoUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetCreatedAt sets the "created_at" field.
func (ciu *CoinInfoUpdate) SetCreatedAt(u uint32) *CoinInfoUpdate {
	ciu.mutation.ResetCreatedAt()
	ciu.mutation.SetCreatedAt(u)
	return ciu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableCreatedAt(u *uint32) *CoinInfoUpdate {
	if u != nil {
		ciu.SetCreatedAt(*u)
	}
	return ciu
}

// AddCreatedAt adds u to the "created_at" field.
func (ciu *CoinInfoUpdate) AddCreatedAt(u int32) *CoinInfoUpdate {
	ciu.mutation.AddCreatedAt(u)
	return ciu
}

// SetUpdatedAt sets the "updated_at" field.
func (ciu *CoinInfoUpdate) SetUpdatedAt(u uint32) *CoinInfoUpdate {
	ciu.mutation.ResetUpdatedAt()
	ciu.mutation.SetUpdatedAt(u)
	return ciu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ciu *CoinInfoUpdate) AddUpdatedAt(u int32) *CoinInfoUpdate {
	ciu.mutation.AddUpdatedAt(u)
	return ciu
}

// SetDeletedAt sets the "deleted_at" field.
func (ciu *CoinInfoUpdate) SetDeletedAt(u uint32) *CoinInfoUpdate {
	ciu.mutation.ResetDeletedAt()
	ciu.mutation.SetDeletedAt(u)
	return ciu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableDeletedAt(u *uint32) *CoinInfoUpdate {
	if u != nil {
		ciu.SetDeletedAt(*u)
	}
	return ciu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ciu *CoinInfoUpdate) AddDeletedAt(u int32) *CoinInfoUpdate {
	ciu.mutation.AddDeletedAt(u)
	return ciu
}

// SetName sets the "name" field.
func (ciu *CoinInfoUpdate) SetName(s string) *CoinInfoUpdate {
	ciu.mutation.SetName(s)
	return ciu
}

// SetChainType sets the "chain_type" field.
func (ciu *CoinInfoUpdate) SetChainType(s string) *CoinInfoUpdate {
	ciu.mutation.SetChainType(s)
	return ciu
}

// SetNillableChainType sets the "chain_type" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableChainType(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetChainType(*s)
	}
	return ciu
}

// ClearChainType clears the value of the "chain_type" field.
func (ciu *CoinInfoUpdate) ClearChainType() *CoinInfoUpdate {
	ciu.mutation.ClearChainType()
	return ciu
}

// SetTokenType sets the "token_type" field.
func (ciu *CoinInfoUpdate) SetTokenType(s string) *CoinInfoUpdate {
	ciu.mutation.SetTokenType(s)
	return ciu
}

// SetNillableTokenType sets the "token_type" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableTokenType(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetTokenType(*s)
	}
	return ciu
}

// ClearTokenType clears the value of the "token_type" field.
func (ciu *CoinInfoUpdate) ClearTokenType() *CoinInfoUpdate {
	ciu.mutation.ClearTokenType()
	return ciu
}

// SetOfficialContract sets the "official_contract" field.
func (ciu *CoinInfoUpdate) SetOfficialContract(s string) *CoinInfoUpdate {
	ciu.mutation.SetOfficialContract(s)
	return ciu
}

// SetPrivateContract sets the "private_contract" field.
func (ciu *CoinInfoUpdate) SetPrivateContract(s string) *CoinInfoUpdate {
	ciu.mutation.SetPrivateContract(s)
	return ciu
}

// SetRemark sets the "remark" field.
func (ciu *CoinInfoUpdate) SetRemark(s string) *CoinInfoUpdate {
	ciu.mutation.SetRemark(s)
	return ciu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ciu *CoinInfoUpdate) SetNillableRemark(s *string) *CoinInfoUpdate {
	if s != nil {
		ciu.SetRemark(*s)
	}
	return ciu
}

// ClearRemark clears the value of the "remark" field.
func (ciu *CoinInfoUpdate) ClearRemark() *CoinInfoUpdate {
	ciu.mutation.ClearRemark()
	return ciu
}

// SetData sets the "data" field.
func (ciu *CoinInfoUpdate) SetData(b []byte) *CoinInfoUpdate {
	ciu.mutation.SetData(b)
	return ciu
}

// ClearData clears the value of the "data" field.
func (ciu *CoinInfoUpdate) ClearData() *CoinInfoUpdate {
	ciu.mutation.ClearData()
	return ciu
}

// Mutation returns the CoinInfoMutation object of the builder.
func (ciu *CoinInfoUpdate) Mutation() *CoinInfoMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CoinInfoUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*CoinInfoMutation)
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
func (ciu *CoinInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CoinInfoUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CoinInfoUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciu *CoinInfoUpdate) defaults() error {
	if _, ok := ciu.mutation.UpdatedAt(); !ok {
		if coininfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coininfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coininfo.UpdateDefaultUpdatedAt()
		ciu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ciu *CoinInfoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinInfoUpdate {
	ciu.modifiers = append(ciu.modifiers, modifiers...)
	return ciu
}

func (ciu *CoinInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coininfo.Table,
			Columns: coininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coininfo.FieldID,
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
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
	}
	if value, ok := ciu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
	}
	if value, ok := ciu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldName,
		})
	}
	if value, ok := ciu.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldChainType,
		})
	}
	if ciu.mutation.ChainTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coininfo.FieldChainType,
		})
	}
	if value, ok := ciu.mutation.TokenType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldTokenType,
		})
	}
	if ciu.mutation.TokenTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coininfo.FieldTokenType,
		})
	}
	if value, ok := ciu.mutation.OfficialContract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldOfficialContract,
		})
	}
	if value, ok := ciu.mutation.PrivateContract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldPrivateContract,
		})
	}
	if value, ok := ciu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldRemark,
		})
	}
	if ciu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coininfo.FieldRemark,
		})
	}
	if value, ok := ciu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: coininfo.FieldData,
		})
	}
	if ciu.mutation.DataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: coininfo.FieldData,
		})
	}
	_spec.Modifiers = ciu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinInfoUpdateOne is the builder for updating a single CoinInfo entity.
type CoinInfoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ciuo *CoinInfoUpdateOne) SetCreatedAt(u uint32) *CoinInfoUpdateOne {
	ciuo.mutation.ResetCreatedAt()
	ciuo.mutation.SetCreatedAt(u)
	return ciuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableCreatedAt(u *uint32) *CoinInfoUpdateOne {
	if u != nil {
		ciuo.SetCreatedAt(*u)
	}
	return ciuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ciuo *CoinInfoUpdateOne) AddCreatedAt(u int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddCreatedAt(u)
	return ciuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuo *CoinInfoUpdateOne) SetUpdatedAt(u uint32) *CoinInfoUpdateOne {
	ciuo.mutation.ResetUpdatedAt()
	ciuo.mutation.SetUpdatedAt(u)
	return ciuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ciuo *CoinInfoUpdateOne) AddUpdatedAt(u int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddUpdatedAt(u)
	return ciuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ciuo *CoinInfoUpdateOne) SetDeletedAt(u uint32) *CoinInfoUpdateOne {
	ciuo.mutation.ResetDeletedAt()
	ciuo.mutation.SetDeletedAt(u)
	return ciuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableDeletedAt(u *uint32) *CoinInfoUpdateOne {
	if u != nil {
		ciuo.SetDeletedAt(*u)
	}
	return ciuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ciuo *CoinInfoUpdateOne) AddDeletedAt(u int32) *CoinInfoUpdateOne {
	ciuo.mutation.AddDeletedAt(u)
	return ciuo
}

// SetName sets the "name" field.
func (ciuo *CoinInfoUpdateOne) SetName(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetName(s)
	return ciuo
}

// SetChainType sets the "chain_type" field.
func (ciuo *CoinInfoUpdateOne) SetChainType(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetChainType(s)
	return ciuo
}

// SetNillableChainType sets the "chain_type" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableChainType(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetChainType(*s)
	}
	return ciuo
}

// ClearChainType clears the value of the "chain_type" field.
func (ciuo *CoinInfoUpdateOne) ClearChainType() *CoinInfoUpdateOne {
	ciuo.mutation.ClearChainType()
	return ciuo
}

// SetTokenType sets the "token_type" field.
func (ciuo *CoinInfoUpdateOne) SetTokenType(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetTokenType(s)
	return ciuo
}

// SetNillableTokenType sets the "token_type" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableTokenType(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetTokenType(*s)
	}
	return ciuo
}

// ClearTokenType clears the value of the "token_type" field.
func (ciuo *CoinInfoUpdateOne) ClearTokenType() *CoinInfoUpdateOne {
	ciuo.mutation.ClearTokenType()
	return ciuo
}

// SetOfficialContract sets the "official_contract" field.
func (ciuo *CoinInfoUpdateOne) SetOfficialContract(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetOfficialContract(s)
	return ciuo
}

// SetPrivateContract sets the "private_contract" field.
func (ciuo *CoinInfoUpdateOne) SetPrivateContract(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetPrivateContract(s)
	return ciuo
}

// SetRemark sets the "remark" field.
func (ciuo *CoinInfoUpdateOne) SetRemark(s string) *CoinInfoUpdateOne {
	ciuo.mutation.SetRemark(s)
	return ciuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ciuo *CoinInfoUpdateOne) SetNillableRemark(s *string) *CoinInfoUpdateOne {
	if s != nil {
		ciuo.SetRemark(*s)
	}
	return ciuo
}

// ClearRemark clears the value of the "remark" field.
func (ciuo *CoinInfoUpdateOne) ClearRemark() *CoinInfoUpdateOne {
	ciuo.mutation.ClearRemark()
	return ciuo
}

// SetData sets the "data" field.
func (ciuo *CoinInfoUpdateOne) SetData(b []byte) *CoinInfoUpdateOne {
	ciuo.mutation.SetData(b)
	return ciuo
}

// ClearData clears the value of the "data" field.
func (ciuo *CoinInfoUpdateOne) ClearData() *CoinInfoUpdateOne {
	ciuo.mutation.ClearData()
	return ciuo
}

// Mutation returns the CoinInfoMutation object of the builder.
func (ciuo *CoinInfoUpdateOne) Mutation() *CoinInfoMutation {
	return ciuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CoinInfoUpdateOne) Select(field string, fields ...string) *CoinInfoUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CoinInfo entity.
func (ciuo *CoinInfoUpdateOne) Save(ctx context.Context) (*CoinInfo, error) {
	var (
		err  error
		node *CoinInfo
	)
	if err := ciuo.defaults(); err != nil {
		return nil, err
	}
	if len(ciuo.hooks) == 0 {
		node, err = ciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinInfoMutation)
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
		v, err := mut.Mutate(ctx, ciuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CoinInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CoinInfoUpdateOne) SaveX(ctx context.Context) *CoinInfo {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CoinInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CoinInfoUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciuo *CoinInfoUpdateOne) defaults() error {
	if _, ok := ciuo.mutation.UpdatedAt(); !ok {
		if coininfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coininfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coininfo.UpdateDefaultUpdatedAt()
		ciuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ciuo *CoinInfoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinInfoUpdateOne {
	ciuo.modifiers = append(ciuo.modifiers, modifiers...)
	return ciuo
}

func (ciuo *CoinInfoUpdateOne) sqlSave(ctx context.Context) (_node *CoinInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coininfo.Table,
			Columns: coininfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coininfo.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coininfo.FieldID)
		for _, f := range fields {
			if !coininfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coininfo.FieldID {
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
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldCreatedAt,
		})
	}
	if value, ok := ciuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldUpdatedAt,
		})
	}
	if value, ok := ciuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
	}
	if value, ok := ciuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coininfo.FieldDeletedAt,
		})
	}
	if value, ok := ciuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldName,
		})
	}
	if value, ok := ciuo.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldChainType,
		})
	}
	if ciuo.mutation.ChainTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coininfo.FieldChainType,
		})
	}
	if value, ok := ciuo.mutation.TokenType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldTokenType,
		})
	}
	if ciuo.mutation.TokenTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coininfo.FieldTokenType,
		})
	}
	if value, ok := ciuo.mutation.OfficialContract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldOfficialContract,
		})
	}
	if value, ok := ciuo.mutation.PrivateContract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldPrivateContract,
		})
	}
	if value, ok := ciuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coininfo.FieldRemark,
		})
	}
	if ciuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coininfo.FieldRemark,
		})
	}
	if value, ok := ciuo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: coininfo.FieldData,
		})
	}
	if ciuo.mutation.DataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: coininfo.FieldData,
		})
	}
	_spec.Modifiers = ciuo.modifiers
	_node = &CoinInfo{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}