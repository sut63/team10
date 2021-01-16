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
	"github.com/team10/app/ent/insurance"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/patientrightstype"
	"github.com/team10/app/ent/predicate"
)

// PatientrightsQuery is the builder for querying Patientrights entities.
type PatientrightsQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Patientrights
	// eager-loading edges.
	withEdgesOfPatientrightsPatientrightstype  *PatientrightstypeQuery
	withEdgesOfPatientrightsInsurance          *InsuranceQuery
	withEdgesOfPatientrightsPatientrecord      *PatientrecordQuery
	withEdgesOfPatientrightsMedicalrecordstaff *MedicalrecordstaffQuery
	withFKs                                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (pq *PatientrightsQuery) Where(ps ...predicate.Patientrights) *PatientrightsQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PatientrightsQuery) Limit(limit int) *PatientrightsQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PatientrightsQuery) Offset(offset int) *PatientrightsQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PatientrightsQuery) Order(o ...OrderFunc) *PatientrightsQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryEdgesOfPatientrightsPatientrightstype chains the current query on the EdgesOfPatientrightsPatientrightstype edge.
func (pq *PatientrightsQuery) QueryEdgesOfPatientrightsPatientrightstype() *PatientrightstypeQuery {
	query := &PatientrightstypeQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrights.Table, patientrights.FieldID, pq.sqlQuery()),
			sqlgraph.To(patientrightstype.Table, patientrightstype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrights.EdgesOfPatientrightsPatientrightstypeTable, patientrights.EdgesOfPatientrightsPatientrightstypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfPatientrightsInsurance chains the current query on the EdgesOfPatientrightsInsurance edge.
func (pq *PatientrightsQuery) QueryEdgesOfPatientrightsInsurance() *InsuranceQuery {
	query := &InsuranceQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrights.Table, patientrights.FieldID, pq.sqlQuery()),
			sqlgraph.To(insurance.Table, insurance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrights.EdgesOfPatientrightsInsuranceTable, patientrights.EdgesOfPatientrightsInsuranceColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfPatientrightsPatientrecord chains the current query on the EdgesOfPatientrightsPatientrecord edge.
func (pq *PatientrightsQuery) QueryEdgesOfPatientrightsPatientrecord() *PatientrecordQuery {
	query := &PatientrecordQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrights.Table, patientrights.FieldID, pq.sqlQuery()),
			sqlgraph.To(patientrecord.Table, patientrecord.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrights.EdgesOfPatientrightsPatientrecordTable, patientrights.EdgesOfPatientrightsPatientrecordColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfPatientrightsMedicalrecordstaff chains the current query on the EdgesOfPatientrightsMedicalrecordstaff edge.
func (pq *PatientrightsQuery) QueryEdgesOfPatientrightsMedicalrecordstaff() *MedicalrecordstaffQuery {
	query := &MedicalrecordstaffQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrights.Table, patientrights.FieldID, pq.sqlQuery()),
			sqlgraph.To(medicalrecordstaff.Table, medicalrecordstaff.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrights.EdgesOfPatientrightsMedicalrecordstaffTable, patientrights.EdgesOfPatientrightsMedicalrecordstaffColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Patientrights entity in the query. Returns *NotFoundError when no patientrights was found.
func (pq *PatientrightsQuery) First(ctx context.Context) (*Patientrights, error) {
	pas, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pas) == 0 {
		return nil, &NotFoundError{patientrights.Label}
	}
	return pas[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PatientrightsQuery) FirstX(ctx context.Context) *Patientrights {
	pa, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pa
}

// FirstID returns the first Patientrights id in the query. Returns *NotFoundError when no id was found.
func (pq *PatientrightsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{patientrights.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PatientrightsQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Patientrights entity in the query, returns an error if not exactly one entity was returned.
func (pq *PatientrightsQuery) Only(ctx context.Context) (*Patientrights, error) {
	pas, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pas) {
	case 1:
		return pas[0], nil
	case 0:
		return nil, &NotFoundError{patientrights.Label}
	default:
		return nil, &NotSingularError{patientrights.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PatientrightsQuery) OnlyX(ctx context.Context) *Patientrights {
	pa, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// OnlyID returns the only Patientrights id in the query, returns an error if not exactly one id was returned.
func (pq *PatientrightsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = &NotSingularError{patientrights.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PatientrightsQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PatientrightsSlice.
func (pq *PatientrightsQuery) All(ctx context.Context) ([]*Patientrights, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PatientrightsQuery) AllX(ctx context.Context) []*Patientrights {
	pas, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pas
}

// IDs executes the query and returns a list of Patientrights ids.
func (pq *PatientrightsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(patientrights.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PatientrightsQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PatientrightsQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PatientrightsQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PatientrightsQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PatientrightsQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PatientrightsQuery) Clone() *PatientrightsQuery {
	return &PatientrightsQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Patientrights{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

//  WithEdgesOfPatientrightsPatientrightstype tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfPatientrightsPatientrightstype" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrightsQuery) WithEdgesOfPatientrightsPatientrightstype(opts ...func(*PatientrightstypeQuery)) *PatientrightsQuery {
	query := &PatientrightstypeQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfPatientrightsPatientrightstype = query
	return pq
}

//  WithEdgesOfPatientrightsInsurance tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfPatientrightsInsurance" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrightsQuery) WithEdgesOfPatientrightsInsurance(opts ...func(*InsuranceQuery)) *PatientrightsQuery {
	query := &InsuranceQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfPatientrightsInsurance = query
	return pq
}

//  WithEdgesOfPatientrightsPatientrecord tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfPatientrightsPatientrecord" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrightsQuery) WithEdgesOfPatientrightsPatientrecord(opts ...func(*PatientrecordQuery)) *PatientrightsQuery {
	query := &PatientrecordQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfPatientrightsPatientrecord = query
	return pq
}

//  WithEdgesOfPatientrightsMedicalrecordstaff tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfPatientrightsMedicalrecordstaff" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrightsQuery) WithEdgesOfPatientrightsMedicalrecordstaff(opts ...func(*MedicalrecordstaffQuery)) *PatientrightsQuery {
	query := &MedicalrecordstaffQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfPatientrightsMedicalrecordstaff = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PermissionDate time.Time `json:"PermissionDate,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Patientrights.Query().
//		GroupBy(patientrights.FieldPermissionDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PatientrightsQuery) GroupBy(field string, fields ...string) *PatientrightsGroupBy {
	group := &PatientrightsGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		PermissionDate time.Time `json:"PermissionDate,omitempty"`
//	}
//
//	client.Patientrights.Query().
//		Select(patientrights.FieldPermissionDate).
//		Scan(ctx, &v)
//
func (pq *PatientrightsQuery) Select(field string, fields ...string) *PatientrightsSelect {
	selector := &PatientrightsSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return selector
}

func (pq *PatientrightsQuery) prepareQuery(ctx context.Context) error {
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PatientrightsQuery) sqlAll(ctx context.Context) ([]*Patientrights, error) {
	var (
		nodes       = []*Patientrights{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [4]bool{
			pq.withEdgesOfPatientrightsPatientrightstype != nil,
			pq.withEdgesOfPatientrightsInsurance != nil,
			pq.withEdgesOfPatientrightsPatientrecord != nil,
			pq.withEdgesOfPatientrightsMedicalrecordstaff != nil,
		}
	)
	if pq.withEdgesOfPatientrightsPatientrightstype != nil || pq.withEdgesOfPatientrightsInsurance != nil || pq.withEdgesOfPatientrightsPatientrecord != nil || pq.withEdgesOfPatientrightsMedicalrecordstaff != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, patientrights.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Patientrights{config: pq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withEdgesOfPatientrightsPatientrightstype; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrights)
		for i := range nodes {
			if fk := nodes[i].Patientrightstype_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patientrightstype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "Patientrightstype_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfPatientrightsPatientrightstype = n
			}
		}
	}

	if query := pq.withEdgesOfPatientrightsInsurance; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrights)
		for i := range nodes {
			if fk := nodes[i].Insurance_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(insurance.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "Insurance_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfPatientrightsInsurance = n
			}
		}
	}

	if query := pq.withEdgesOfPatientrightsPatientrecord; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrights)
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
				nodes[i].Edges.EdgesOfPatientrightsPatientrecord = n
			}
		}
	}

	if query := pq.withEdgesOfPatientrightsMedicalrecordstaff; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrights)
		for i := range nodes {
			if fk := nodes[i].medicalrecordstaff_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(medicalrecordstaff.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "medicalrecordstaff_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfPatientrightsMedicalrecordstaff = n
			}
		}
	}

	return nodes, nil
}

func (pq *PatientrightsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PatientrightsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PatientrightsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientrights.Table,
			Columns: patientrights.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrights.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PatientrightsQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(patientrights.Table)
	selector := builder.Select(t1.Columns(patientrights.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(patientrights.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PatientrightsGroupBy is the builder for group-by Patientrights entities.
type PatientrightsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PatientrightsGroupBy) Aggregate(fns ...AggregateFunc) *PatientrightsGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PatientrightsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrightsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrightsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PatientrightsGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PatientrightsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PatientrightsGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PatientrightsSelect is the builder for select fields of Patientrights entities.
type PatientrightsSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PatientrightsSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ps.path(ctx)
	if err != nil {
		return err
	}
	ps.sql = query
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PatientrightsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PatientrightsSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PatientrightsSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PatientrightsSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PatientrightsSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PatientrightsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PatientrightsSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrightsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PatientrightsSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ps *PatientrightsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrights.Label}
	default:
		err = fmt.Errorf("ent: PatientrightsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PatientrightsSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PatientrightsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PatientrightsSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}
