// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/educationlevel"
	"github.com/team10/app/ent/officeroom"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/prename"
	"github.com/team10/app/ent/treatment"
)

// DoctorinfoQuery is the builder for querying Doctorinfo entities.
type DoctorinfoQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Doctorinfo
	// eager-loading edges.
	withDepartment     *DepartmentQuery
	withEducationlevel *EducationlevelQuery
	withOfficeroom     *OfficeroomQuery
	withPrename        *PrenameQuery
	withTreatment      *TreatmentQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (dq *DoctorinfoQuery) Where(ps ...predicate.Doctorinfo) *DoctorinfoQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DoctorinfoQuery) Limit(limit int) *DoctorinfoQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DoctorinfoQuery) Offset(offset int) *DoctorinfoQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DoctorinfoQuery) Order(o ...OrderFunc) *DoctorinfoQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDepartment chains the current query on the department edge.
func (dq *DoctorinfoQuery) QueryDepartment() *DepartmentQuery {
	query := &DepartmentQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctorinfo.Table, doctorinfo.FieldID, dq.sqlQuery()),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctorinfo.DepartmentTable, doctorinfo.DepartmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEducationlevel chains the current query on the educationlevel edge.
func (dq *DoctorinfoQuery) QueryEducationlevel() *EducationlevelQuery {
	query := &EducationlevelQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctorinfo.Table, doctorinfo.FieldID, dq.sqlQuery()),
			sqlgraph.To(educationlevel.Table, educationlevel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctorinfo.EducationlevelTable, doctorinfo.EducationlevelColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOfficeroom chains the current query on the officeroom edge.
func (dq *DoctorinfoQuery) QueryOfficeroom() *OfficeroomQuery {
	query := &OfficeroomQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctorinfo.Table, doctorinfo.FieldID, dq.sqlQuery()),
			sqlgraph.To(officeroom.Table, officeroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctorinfo.OfficeroomTable, doctorinfo.OfficeroomColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPrename chains the current query on the prename edge.
func (dq *DoctorinfoQuery) QueryPrename() *PrenameQuery {
	query := &PrenameQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctorinfo.Table, doctorinfo.FieldID, dq.sqlQuery()),
			sqlgraph.To(prename.Table, prename.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doctorinfo.PrenameTable, doctorinfo.PrenameColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTreatment chains the current query on the treatment edge.
func (dq *DoctorinfoQuery) QueryTreatment() *TreatmentQuery {
	query := &TreatmentQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doctorinfo.Table, doctorinfo.FieldID, dq.sqlQuery()),
			sqlgraph.To(treatment.Table, treatment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doctorinfo.TreatmentTable, doctorinfo.TreatmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Doctorinfo entity in the query. Returns *NotFoundError when no doctorinfo was found.
func (dq *DoctorinfoQuery) First(ctx context.Context) (*Doctorinfo, error) {
	ds, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ds) == 0 {
		return nil, &NotFoundError{doctorinfo.Label}
	}
	return ds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DoctorinfoQuery) FirstX(ctx context.Context) *Doctorinfo {
	d, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return d
}

// FirstID returns the first Doctorinfo id in the query. Returns *NotFoundError when no id was found.
func (dq *DoctorinfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{doctorinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (dq *DoctorinfoQuery) FirstXID(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Doctorinfo entity in the query, returns an error if not exactly one entity was returned.
func (dq *DoctorinfoQuery) Only(ctx context.Context) (*Doctorinfo, error) {
	ds, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ds) {
	case 1:
		return ds[0], nil
	case 0:
		return nil, &NotFoundError{doctorinfo.Label}
	default:
		return nil, &NotSingularError{doctorinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DoctorinfoQuery) OnlyX(ctx context.Context) *Doctorinfo {
	d, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// OnlyID returns the only Doctorinfo id in the query, returns an error if not exactly one id was returned.
func (dq *DoctorinfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = &NotSingularError{doctorinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DoctorinfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Doctorinfos.
func (dq *DoctorinfoQuery) All(ctx context.Context) ([]*Doctorinfo, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DoctorinfoQuery) AllX(ctx context.Context) []*Doctorinfo {
	ds, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ds
}

// IDs executes the query and returns a list of Doctorinfo ids.
func (dq *DoctorinfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(doctorinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DoctorinfoQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DoctorinfoQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DoctorinfoQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DoctorinfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DoctorinfoQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DoctorinfoQuery) Clone() *DoctorinfoQuery {
	return &DoctorinfoQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		unique:     append([]string{}, dq.unique...),
		predicates: append([]predicate.Doctorinfo{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

//  WithDepartment tells the query-builder to eager-loads the nodes that are connected to
// the "department" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorinfoQuery) WithDepartment(opts ...func(*DepartmentQuery)) *DoctorinfoQuery {
	query := &DepartmentQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDepartment = query
	return dq
}

//  WithEducationlevel tells the query-builder to eager-loads the nodes that are connected to
// the "educationlevel" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorinfoQuery) WithEducationlevel(opts ...func(*EducationlevelQuery)) *DoctorinfoQuery {
	query := &EducationlevelQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withEducationlevel = query
	return dq
}

//  WithOfficeroom tells the query-builder to eager-loads the nodes that are connected to
// the "officeroom" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorinfoQuery) WithOfficeroom(opts ...func(*OfficeroomQuery)) *DoctorinfoQuery {
	query := &OfficeroomQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withOfficeroom = query
	return dq
}

//  WithPrename tells the query-builder to eager-loads the nodes that are connected to
// the "prename" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorinfoQuery) WithPrename(opts ...func(*PrenameQuery)) *DoctorinfoQuery {
	query := &PrenameQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withPrename = query
	return dq
}

//  WithTreatment tells the query-builder to eager-loads the nodes that are connected to
// the "treatment" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DoctorinfoQuery) WithTreatment(opts ...func(*TreatmentQuery)) *DoctorinfoQuery {
	query := &TreatmentQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withTreatment = query
	return dq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Doctorname string `json:"doctorname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Doctorinfo.Query().
//		GroupBy(doctorinfo.FieldDoctorname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DoctorinfoQuery) GroupBy(field string, fields ...string) *DoctorinfoGroupBy {
	group := &DoctorinfoGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Doctorname string `json:"doctorname,omitempty"`
//	}
//
//	client.Doctorinfo.Query().
//		Select(doctorinfo.FieldDoctorname).
//		Scan(ctx, &v)
//
func (dq *DoctorinfoQuery) Select(field string, fields ...string) *DoctorinfoSelect {
	selector := &DoctorinfoSelect{config: dq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return selector
}

func (dq *DoctorinfoQuery) prepareQuery(ctx context.Context) error {
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DoctorinfoQuery) sqlAll(ctx context.Context) ([]*Doctorinfo, error) {
	var (
		nodes       = []*Doctorinfo{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [5]bool{
			dq.withDepartment != nil,
			dq.withEducationlevel != nil,
			dq.withOfficeroom != nil,
			dq.withPrename != nil,
			dq.withTreatment != nil,
		}
	)
	if dq.withDepartment != nil || dq.withEducationlevel != nil || dq.withOfficeroom != nil || dq.withPrename != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, doctorinfo.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Doctorinfo{config: dq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withDepartment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctorinfo)
		for i := range nodes {
			if fk := nodes[i].department; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "department" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Department = n
			}
		}
	}

	if query := dq.withEducationlevel; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctorinfo)
		for i := range nodes {
			if fk := nodes[i].level; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(educationlevel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "level" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Educationlevel = n
			}
		}
	}

	if query := dq.withOfficeroom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctorinfo)
		for i := range nodes {
			if fk := nodes[i].roomnumber; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(officeroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "roomnumber" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Officeroom = n
			}
		}
	}

	if query := dq.withPrename; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Doctorinfo)
		for i := range nodes {
			if fk := nodes[i].prefix; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(prename.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "prefix" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Prename = n
			}
		}
	}

	if query := dq.withTreatment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Doctorinfo)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Treatment(func(s *sql.Selector) {
			s.Where(sql.InValues(doctorinfo.TreatmentColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.doctorinfo_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "doctorinfo_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctorinfo_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Treatment = append(node.Edges.Treatment, n)
		}
	}

	return nodes, nil
}

func (dq *DoctorinfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DoctorinfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DoctorinfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doctorinfo.Table,
			Columns: doctorinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctorinfo.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DoctorinfoQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(doctorinfo.Table)
	selector := builder.Select(t1.Columns(doctorinfo.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(doctorinfo.Columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DoctorinfoGroupBy is the builder for group-by Doctorinfo entities.
type DoctorinfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DoctorinfoGroupBy) Aggregate(fns ...AggregateFunc) *DoctorinfoGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scan the result into the given value.
func (dgb *DoctorinfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (dgb *DoctorinfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DoctorinfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DoctorinfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgb.sqlQuery().Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DoctorinfoGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DoctorinfoSelect is the builder for select fields of Doctorinfo entities.
type DoctorinfoSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ds *DoctorinfoSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ds.path(ctx)
	if err != nil {
		return err
	}
	ds.sql = query
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DoctorinfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DoctorinfoSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DoctorinfoSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DoctorinfoSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DoctorinfoSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DoctorinfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DoctorinfoSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DoctorinfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DoctorinfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ds *DoctorinfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{doctorinfo.Label}
	default:
		err = fmt.Errorf("ent: DoctorinfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DoctorinfoSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DoctorinfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DoctorinfoSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
