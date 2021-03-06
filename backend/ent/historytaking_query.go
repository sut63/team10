// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/symptomseverity"
)

// HistorytakingQuery is the builder for querying Historytaking entities.
type HistorytakingQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Historytaking
	// eager-loading edges.
	withEdgesOfNurse           *NurseQuery
	withEdgesOfDepartment      *DepartmentQuery
	withEdgesOfSymptomseverity *SymptomseverityQuery
	withEdgesOfPatientrecord   *PatientrecordQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (hq *HistorytakingQuery) Where(ps ...predicate.Historytaking) *HistorytakingQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit adds a limit step to the query.
func (hq *HistorytakingQuery) Limit(limit int) *HistorytakingQuery {
	hq.limit = &limit
	return hq
}

// Offset adds an offset step to the query.
func (hq *HistorytakingQuery) Offset(offset int) *HistorytakingQuery {
	hq.offset = &offset
	return hq
}

// Order adds an order step to the query.
func (hq *HistorytakingQuery) Order(o ...OrderFunc) *HistorytakingQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// QueryEdgesOfNurse chains the current query on the EdgesOfNurse edge.
func (hq *HistorytakingQuery) QueryEdgesOfNurse() *NurseQuery {
	query := &NurseQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(historytaking.Table, historytaking.FieldID, hq.sqlQuery()),
			sqlgraph.To(nurse.Table, nurse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, historytaking.EdgesOfNurseTable, historytaking.EdgesOfNurseColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfDepartment chains the current query on the EdgesOfDepartment edge.
func (hq *HistorytakingQuery) QueryEdgesOfDepartment() *DepartmentQuery {
	query := &DepartmentQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(historytaking.Table, historytaking.FieldID, hq.sqlQuery()),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, historytaking.EdgesOfDepartmentTable, historytaking.EdgesOfDepartmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfSymptomseverity chains the current query on the EdgesOfSymptomseverity edge.
func (hq *HistorytakingQuery) QueryEdgesOfSymptomseverity() *SymptomseverityQuery {
	query := &SymptomseverityQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(historytaking.Table, historytaking.FieldID, hq.sqlQuery()),
			sqlgraph.To(symptomseverity.Table, symptomseverity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, historytaking.EdgesOfSymptomseverityTable, historytaking.EdgesOfSymptomseverityColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfPatientrecord chains the current query on the EdgesOfPatientrecord edge.
func (hq *HistorytakingQuery) QueryEdgesOfPatientrecord() *PatientrecordQuery {
	query := &PatientrecordQuery{config: hq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(historytaking.Table, historytaking.FieldID, hq.sqlQuery()),
			sqlgraph.To(patientrecord.Table, patientrecord.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, historytaking.EdgesOfPatientrecordTable, historytaking.EdgesOfPatientrecordColumn),
		)
		fromU = sqlgraph.SetNeighbors(hq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Historytaking entity in the query. Returns *NotFoundError when no historytaking was found.
func (hq *HistorytakingQuery) First(ctx context.Context) (*Historytaking, error) {
	hs, err := hq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(hs) == 0 {
		return nil, &NotFoundError{historytaking.Label}
	}
	return hs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HistorytakingQuery) FirstX(ctx context.Context) *Historytaking {
	h, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return h
}

// FirstID returns the first Historytaking id in the query. Returns *NotFoundError when no id was found.
func (hq *HistorytakingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{historytaking.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (hq *HistorytakingQuery) FirstXID(ctx context.Context) int {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Historytaking entity in the query, returns an error if not exactly one entity was returned.
func (hq *HistorytakingQuery) Only(ctx context.Context) (*Historytaking, error) {
	hs, err := hq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(hs) {
	case 1:
		return hs[0], nil
	case 0:
		return nil, &NotFoundError{historytaking.Label}
	default:
		return nil, &NotSingularError{historytaking.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HistorytakingQuery) OnlyX(ctx context.Context) *Historytaking {
	h, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return h
}

// OnlyID returns the only Historytaking id in the query, returns an error if not exactly one id was returned.
func (hq *HistorytakingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = &NotSingularError{historytaking.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HistorytakingQuery) OnlyIDX(ctx context.Context) int {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Historytakings.
func (hq *HistorytakingQuery) All(ctx context.Context) ([]*Historytaking, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hq *HistorytakingQuery) AllX(ctx context.Context) []*Historytaking {
	hs, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return hs
}

// IDs executes the query and returns a list of Historytaking ids.
func (hq *HistorytakingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := hq.Select(historytaking.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HistorytakingQuery) IDsX(ctx context.Context) []int {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HistorytakingQuery) Count(ctx context.Context) (int, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HistorytakingQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HistorytakingQuery) Exist(ctx context.Context) (bool, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HistorytakingQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HistorytakingQuery) Clone() *HistorytakingQuery {
	return &HistorytakingQuery{
		config:     hq.config,
		limit:      hq.limit,
		offset:     hq.offset,
		order:      append([]OrderFunc{}, hq.order...),
		unique:     append([]string{}, hq.unique...),
		predicates: append([]predicate.Historytaking{}, hq.predicates...),
		// clone intermediate query.
		sql:  hq.sql.Clone(),
		path: hq.path,
	}
}

//  WithEdgesOfNurse tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfNurse" edge. The optional arguments used to configure the query builder of the edge.
func (hq *HistorytakingQuery) WithEdgesOfNurse(opts ...func(*NurseQuery)) *HistorytakingQuery {
	query := &NurseQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withEdgesOfNurse = query
	return hq
}

//  WithEdgesOfDepartment tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfDepartment" edge. The optional arguments used to configure the query builder of the edge.
func (hq *HistorytakingQuery) WithEdgesOfDepartment(opts ...func(*DepartmentQuery)) *HistorytakingQuery {
	query := &DepartmentQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withEdgesOfDepartment = query
	return hq
}

//  WithEdgesOfSymptomseverity tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfSymptomseverity" edge. The optional arguments used to configure the query builder of the edge.
func (hq *HistorytakingQuery) WithEdgesOfSymptomseverity(opts ...func(*SymptomseverityQuery)) *HistorytakingQuery {
	query := &SymptomseverityQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withEdgesOfSymptomseverity = query
	return hq
}

//  WithEdgesOfPatientrecord tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfPatientrecord" edge. The optional arguments used to configure the query builder of the edge.
func (hq *HistorytakingQuery) WithEdgesOfPatientrecord(opts ...func(*PatientrecordQuery)) *HistorytakingQuery {
	query := &PatientrecordQuery{config: hq.config}
	for _, opt := range opts {
		opt(query)
	}
	hq.withEdgesOfPatientrecord = query
	return hq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Hight float32 `json:"hight,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Historytaking.Query().
//		GroupBy(historytaking.FieldHight).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (hq *HistorytakingQuery) GroupBy(field string, fields ...string) *HistorytakingGroupBy {
	group := &HistorytakingGroupBy{config: hq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Hight float32 `json:"hight,omitempty"`
//	}
//
//	client.Historytaking.Query().
//		Select(historytaking.FieldHight).
//		Scan(ctx, &v)
//
func (hq *HistorytakingQuery) Select(field string, fields ...string) *HistorytakingSelect {
	selector := &HistorytakingSelect{config: hq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hq.sqlQuery(), nil
	}
	return selector
}

func (hq *HistorytakingQuery) prepareQuery(ctx context.Context) error {
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HistorytakingQuery) sqlAll(ctx context.Context) ([]*Historytaking, error) {
	var (
		nodes       = []*Historytaking{}
		withFKs     = hq.withFKs
		_spec       = hq.querySpec()
		loadedTypes = [4]bool{
			hq.withEdgesOfNurse != nil,
			hq.withEdgesOfDepartment != nil,
			hq.withEdgesOfSymptomseverity != nil,
			hq.withEdgesOfPatientrecord != nil,
		}
	)
	if hq.withEdgesOfNurse != nil || hq.withEdgesOfDepartment != nil || hq.withEdgesOfSymptomseverity != nil || hq.withEdgesOfPatientrecord != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, historytaking.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Historytaking{config: hq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := hq.withEdgesOfNurse; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Historytaking)
		for i := range nodes {
			if fk := nodes[i].nurse_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(nurse.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "nurse_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfNurse = n
			}
		}
	}

	if query := hq.withEdgesOfDepartment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Historytaking)
		for i := range nodes {
			if fk := nodes[i].department_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(department.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "department_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfDepartment = n
			}
		}
	}

	if query := hq.withEdgesOfSymptomseverity; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Historytaking)
		for i := range nodes {
			if fk := nodes[i].symptomseverity_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(symptomseverity.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "symptomseverity_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfSymptomseverity = n
			}
		}
	}

	if query := hq.withEdgesOfPatientrecord; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Historytaking)
		for i := range nodes {
			if fk := nodes[i].patientrecord_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patientrecord.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patientrecord_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfPatientrecord = n
			}
		}
	}

	return nodes, nil
}

func (hq *HistorytakingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HistorytakingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (hq *HistorytakingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   historytaking.Table,
			Columns: historytaking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: historytaking.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HistorytakingQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(historytaking.Table)
	selector := builder.Select(t1.Columns(historytaking.Columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(historytaking.Columns...)...)
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HistorytakingGroupBy is the builder for group-by Historytaking entities.
type HistorytakingGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HistorytakingGroupBy) Aggregate(fns ...AggregateFunc) *HistorytakingGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the group-by query and scan the result into the given value.
func (hgb *HistorytakingGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hgb.path(ctx)
	if err != nil {
		return err
	}
	hgb.sql = query
	return hgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := hgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistorytakingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) StringsX(ctx context.Context) []string {
	v, err := hgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) StringX(ctx context.Context) string {
	v, err := hgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistorytakingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) IntsX(ctx context.Context) []int {
	v, err := hgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) IntX(ctx context.Context) int {
	v, err := hgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistorytakingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := hgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) Float64X(ctx context.Context) float64 {
	v, err := hgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(hgb.fields) > 1 {
		return nil, errors.New("ent: HistorytakingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := hgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := hgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (hgb *HistorytakingGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hgb *HistorytakingGroupBy) BoolX(ctx context.Context) bool {
	v, err := hgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hgb *HistorytakingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hgb.sqlQuery().Query()
	if err := hgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hgb *HistorytakingGroupBy) sqlQuery() *sql.Selector {
	selector := hgb.sql
	columns := make([]string, 0, len(hgb.fields)+len(hgb.fns))
	columns = append(columns, hgb.fields...)
	for _, fn := range hgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(hgb.fields...)
}

// HistorytakingSelect is the builder for select fields of Historytaking entities.
type HistorytakingSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (hs *HistorytakingSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := hs.path(ctx)
	if err != nil {
		return err
	}
	hs.sql = query
	return hs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hs *HistorytakingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := hs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorytakingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hs *HistorytakingSelect) StringsX(ctx context.Context) []string {
	v, err := hs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hs *HistorytakingSelect) StringX(ctx context.Context) string {
	v, err := hs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorytakingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hs *HistorytakingSelect) IntsX(ctx context.Context) []int {
	v, err := hs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hs *HistorytakingSelect) IntX(ctx context.Context) int {
	v, err := hs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorytakingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hs *HistorytakingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := hs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hs *HistorytakingSelect) Float64X(ctx context.Context) float64 {
	v, err := hs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(hs.fields) > 1 {
		return nil, errors.New("ent: HistorytakingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := hs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hs *HistorytakingSelect) BoolsX(ctx context.Context) []bool {
	v, err := hs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (hs *HistorytakingSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{historytaking.Label}
	default:
		err = fmt.Errorf("ent: HistorytakingSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hs *HistorytakingSelect) BoolX(ctx context.Context) bool {
	v, err := hs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hs *HistorytakingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hs.sqlQuery().Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hs *HistorytakingSelect) sqlQuery() sql.Querier {
	selector := hs.sql
	selector.Select(selector.Columns(hs.fields...)...)
	return selector
}
