// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinsInfoQuery is the builder for querying CoinsInfo entities.
type CoinsInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CoinsInfo
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinsInfoQuery builder.
func (ciq *CoinsInfoQuery) Where(ps ...predicate.CoinsInfo) *CoinsInfoQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit adds a limit step to the query.
func (ciq *CoinsInfoQuery) Limit(limit int) *CoinsInfoQuery {
	ciq.limit = &limit
	return ciq
}

// Offset adds an offset step to the query.
func (ciq *CoinsInfoQuery) Offset(offset int) *CoinsInfoQuery {
	ciq.offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *CoinsInfoQuery) Unique(unique bool) *CoinsInfoQuery {
	ciq.unique = &unique
	return ciq
}

// Order adds an order step to the query.
func (ciq *CoinsInfoQuery) Order(o ...OrderFunc) *CoinsInfoQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// First returns the first CoinsInfo entity from the query.
// Returns a *NotFoundError when no CoinsInfo was found.
func (ciq *CoinsInfoQuery) First(ctx context.Context) (*CoinsInfo, error) {
	nodes, err := ciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coinsinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CoinsInfoQuery) FirstX(ctx context.Context) *CoinsInfo {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoinsInfo ID from the query.
// Returns a *NotFoundError when no CoinsInfo ID was found.
func (ciq *CoinsInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coinsinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *CoinsInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoinsInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CoinsInfo entity is found.
// Returns a *NotFoundError when no CoinsInfo entities are found.
func (ciq *CoinsInfoQuery) Only(ctx context.Context) (*CoinsInfo, error) {
	nodes, err := ciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coinsinfo.Label}
	default:
		return nil, &NotSingularError{coinsinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CoinsInfoQuery) OnlyX(ctx context.Context) *CoinsInfo {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoinsInfo ID in the query.
// Returns a *NotSingularError when more than one CoinsInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciq *CoinsInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coinsinfo.Label}
	default:
		err = &NotSingularError{coinsinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CoinsInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoinsInfos.
func (ciq *CoinsInfoQuery) All(ctx context.Context) ([]*CoinsInfo, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CoinsInfoQuery) AllX(ctx context.Context) []*CoinsInfo {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoinsInfo IDs.
func (ciq *CoinsInfoQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := ciq.Select(coinsinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CoinsInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CoinsInfoQuery) Count(ctx context.Context) (int, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CoinsInfoQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CoinsInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CoinsInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinsInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CoinsInfoQuery) Clone() *CoinsInfoQuery {
	if ciq == nil {
		return nil
	}
	return &CoinsInfoQuery{
		config:     ciq.config,
		limit:      ciq.limit,
		offset:     ciq.offset,
		order:      append([]OrderFunc{}, ciq.order...),
		predicates: append([]predicate.CoinsInfo{}, ciq.predicates...),
		// clone intermediate query.
		sql:    ciq.sql.Clone(),
		path:   ciq.path,
		unique: ciq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CoinsInfo.Query().
//		GroupBy(coinsinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ciq *CoinsInfoQuery) GroupBy(field string, fields ...string) *CoinsInfoGroupBy {
	grbuild := &CoinsInfoGroupBy{config: ciq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(ctx), nil
	}
	grbuild.label = coinsinfo.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.CoinsInfo.Query().
//		Select(coinsinfo.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ciq *CoinsInfoQuery) Select(fields ...string) *CoinsInfoSelect {
	ciq.fields = append(ciq.fields, fields...)
	selbuild := &CoinsInfoSelect{CoinsInfoQuery: ciq}
	selbuild.label = coinsinfo.Label
	selbuild.flds, selbuild.scan = &ciq.fields, selbuild.Scan
	return selbuild
}

func (ciq *CoinsInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ciq.fields {
		if !coinsinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	if coinsinfo.Policy == nil {
		return errors.New("ent: uninitialized coinsinfo.Policy (forgotten import ent/runtime?)")
	}
	if err := coinsinfo.Policy.EvalQuery(ctx, ciq); err != nil {
		return err
	}
	return nil
}

func (ciq *CoinsInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CoinsInfo, error) {
	var (
		nodes = []*CoinsInfo{}
		_spec = ciq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CoinsInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CoinsInfo{config: ciq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ciq.modifiers) > 0 {
		_spec.Modifiers = ciq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ciq *CoinsInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	if len(ciq.modifiers) > 0 {
		_spec.Modifiers = ciq.modifiers
	}
	_spec.Node.Columns = ciq.fields
	if len(ciq.fields) > 0 {
		_spec.Unique = ciq.unique != nil && *ciq.unique
	}
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CoinsInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ciq *CoinsInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinsinfo.Table,
			Columns: coinsinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinsinfo.FieldID,
			},
		},
		From:   ciq.sql,
		Unique: true,
	}
	if unique := ciq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ciq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinsinfo.FieldID)
		for i := range fields {
			if fields[i] != coinsinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CoinsInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(coinsinfo.Table)
	columns := ciq.fields
	if len(columns) == 0 {
		columns = coinsinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciq.unique != nil && *ciq.unique {
		selector.Distinct()
	}
	for _, m := range ciq.modifiers {
		m(selector)
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ciq *CoinsInfoQuery) Modify(modifiers ...func(s *sql.Selector)) *CoinsInfoSelect {
	ciq.modifiers = append(ciq.modifiers, modifiers...)
	return ciq.Select()
}

// CoinsInfoGroupBy is the group-by builder for CoinsInfo entities.
type CoinsInfoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CoinsInfoGroupBy) Aggregate(fns ...AggregateFunc) *CoinsInfoGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the group-by query and scans the result into the given value.
func (cigb *CoinsInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cigb.path(ctx)
	if err != nil {
		return err
	}
	cigb.sql = query
	return cigb.sqlScan(ctx, v)
}

func (cigb *CoinsInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cigb.fields {
		if !coinsinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cigb *CoinsInfoGroupBy) sqlQuery() *sql.Selector {
	selector := cigb.sql.Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cigb.fields)+len(cigb.fns))
		for _, f := range cigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cigb.fields...)...)
}

// CoinsInfoSelect is the builder for selecting fields of CoinsInfo entities.
type CoinsInfoSelect struct {
	*CoinsInfoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cis *CoinsInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	cis.sql = cis.CoinsInfoQuery.sqlQuery(ctx)
	return cis.sqlScan(ctx, v)
}

func (cis *CoinsInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cis.sql.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cis *CoinsInfoSelect) Modify(modifiers ...func(s *sql.Selector)) *CoinsInfoSelect {
	cis.modifiers = append(cis.modifiers, modifiers...)
	return cis
}