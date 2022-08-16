// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/coinsinfo"
	"github.com/NpoolPlatform/build-chain/pkg/db/ent/deployedcoin"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coinsinfo.Table,
			Columns: coinsinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinsinfo.FieldID,
			},
		},
		Type: "CoinsInfo",
		Fields: map[string]*sqlgraph.FieldSpec{
			coinsinfo.FieldCreatedAt:  {Type: field.TypeUint32, Column: coinsinfo.FieldCreatedAt},
			coinsinfo.FieldUpdatedAt:  {Type: field.TypeUint32, Column: coinsinfo.FieldUpdatedAt},
			coinsinfo.FieldDeletedAt:  {Type: field.TypeUint32, Column: coinsinfo.FieldDeletedAt},
			coinsinfo.FieldName:       {Type: field.TypeString, Column: coinsinfo.FieldName},
			coinsinfo.FieldChainType:  {Type: field.TypeString, Column: coinsinfo.FieldChainType},
			coinsinfo.FieldTokenType:  {Type: field.TypeString, Column: coinsinfo.FieldTokenType},
			coinsinfo.FieldContract:   {Type: field.TypeString, Column: coinsinfo.FieldContract},
			coinsinfo.FieldSimilarity: {Type: field.TypeInt32, Column: coinsinfo.FieldSimilarity},
			coinsinfo.FieldRemark:     {Type: field.TypeString, Column: coinsinfo.FieldRemark},
			coinsinfo.FieldData:       {Type: field.TypeBytes, Column: coinsinfo.FieldData},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   deployedcoin.Table,
			Columns: deployedcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deployedcoin.FieldID,
			},
		},
		Type: "DeployedCoin",
		Fields: map[string]*sqlgraph.FieldSpec{
			deployedcoin.FieldCreatedAt: {Type: field.TypeUint32, Column: deployedcoin.FieldCreatedAt},
			deployedcoin.FieldUpdatedAt: {Type: field.TypeUint32, Column: deployedcoin.FieldUpdatedAt},
			deployedcoin.FieldDeletedAt: {Type: field.TypeUint32, Column: deployedcoin.FieldDeletedAt},
			deployedcoin.FieldCoinID:    {Type: field.TypeUUID, Column: deployedcoin.FieldCoinID},
			deployedcoin.FieldContract:  {Type: field.TypeString, Column: deployedcoin.FieldContract},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (ciq *CoinsInfoQuery) addPredicate(pred func(s *sql.Selector)) {
	ciq.predicates = append(ciq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CoinsInfoQuery builder.
func (ciq *CoinsInfoQuery) Filter() *CoinsInfoFilter {
	return &CoinsInfoFilter{config: ciq.config, predicateAdder: ciq}
}

// addPredicate implements the predicateAdder interface.
func (m *CoinsInfoMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CoinsInfoMutation builder.
func (m *CoinsInfoMutation) Filter() *CoinsInfoFilter {
	return &CoinsInfoFilter{config: m.config, predicateAdder: m}
}

// CoinsInfoFilter provides a generic filtering capability at runtime for CoinsInfoQuery.
type CoinsInfoFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CoinsInfoFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CoinsInfoFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(coinsinfo.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CoinsInfoFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(coinsinfo.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CoinsInfoFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(coinsinfo.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CoinsInfoFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(coinsinfo.FieldDeletedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *CoinsInfoFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(coinsinfo.FieldName))
}

// WhereChainType applies the entql string predicate on the chain_type field.
func (f *CoinsInfoFilter) WhereChainType(p entql.StringP) {
	f.Where(p.Field(coinsinfo.FieldChainType))
}

// WhereTokenType applies the entql string predicate on the token_type field.
func (f *CoinsInfoFilter) WhereTokenType(p entql.StringP) {
	f.Where(p.Field(coinsinfo.FieldTokenType))
}

// WhereContract applies the entql string predicate on the contract field.
func (f *CoinsInfoFilter) WhereContract(p entql.StringP) {
	f.Where(p.Field(coinsinfo.FieldContract))
}

// WhereSimilarity applies the entql int32 predicate on the similarity field.
func (f *CoinsInfoFilter) WhereSimilarity(p entql.Int32P) {
	f.Where(p.Field(coinsinfo.FieldSimilarity))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *CoinsInfoFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(coinsinfo.FieldRemark))
}

// WhereData applies the entql []byte predicate on the data field.
func (f *CoinsInfoFilter) WhereData(p entql.BytesP) {
	f.Where(p.Field(coinsinfo.FieldData))
}

// addPredicate implements the predicateAdder interface.
func (dcq *DeployedCoinQuery) addPredicate(pred func(s *sql.Selector)) {
	dcq.predicates = append(dcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DeployedCoinQuery builder.
func (dcq *DeployedCoinQuery) Filter() *DeployedCoinFilter {
	return &DeployedCoinFilter{config: dcq.config, predicateAdder: dcq}
}

// addPredicate implements the predicateAdder interface.
func (m *DeployedCoinMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DeployedCoinMutation builder.
func (m *DeployedCoinMutation) Filter() *DeployedCoinFilter {
	return &DeployedCoinFilter{config: m.config, predicateAdder: m}
}

// DeployedCoinFilter provides a generic filtering capability at runtime for DeployedCoinQuery.
type DeployedCoinFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DeployedCoinFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *DeployedCoinFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(deployedcoin.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *DeployedCoinFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(deployedcoin.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *DeployedCoinFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(deployedcoin.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *DeployedCoinFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(deployedcoin.FieldDeletedAt))
}

// WhereCoinID applies the entql [16]byte predicate on the coin_id field.
func (f *DeployedCoinFilter) WhereCoinID(p entql.ValueP) {
	f.Where(p.Field(deployedcoin.FieldCoinID))
}

// WhereContract applies the entql string predicate on the contract field.
func (f *DeployedCoinFilter) WhereContract(p entql.StringP) {
	f.Where(p.Field(deployedcoin.FieldContract))
}