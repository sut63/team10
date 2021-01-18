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
	"github.com/team10/app/ent/gender"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/patientrights"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/prename"
	"github.com/team10/app/ent/treatment"
)

// PatientrecordQuery is the builder for querying Patientrecord entities.
type PatientrecordQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Patientrecord
	// eager-loading edges.
	withEdgesOfGender                     *GenderQuery
	withEdgesOfMedicalrecordstaff         *MedicalrecordstaffQuery
	withEdgesOfPrename                    *PrenameQuery
	withEdgesOfHistorytaking              *HistorytakingQuery
	withEdgesOfTreatment                  *TreatmentQuery
	withEdgesOfPatientrecordPatientrights *PatientrightsQuery
	withFKs                               bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (pq *PatientrecordQuery) Where(ps ...predicate.Patientrecord) *PatientrecordQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PatientrecordQuery) Limit(limit int) *PatientrecordQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PatientrecordQuery) Offset(offset int) *PatientrecordQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PatientrecordQuery) Order(o ...OrderFunc) *PatientrecordQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryEdgesOfGender chains the current query on the EdgesOfGender edge.
func (pq *PatientrecordQuery) QueryEdgesOfGender() *GenderQuery {
	query := &GenderQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrecord.Table, patientrecord.FieldID, pq.sqlQuery()),
			sqlgraph.To(gender.Table, gender.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrecord.EdgesOfGenderTable, patientrecord.EdgesOfGenderColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfMedicalrecordstaff chains the current query on the EdgesOfMedicalrecordstaff edge.
func (pq *PatientrecordQuery) QueryEdgesOfMedicalrecordstaff() *MedicalrecordstaffQuery {
	query := &MedicalrecordstaffQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrecord.Table, patientrecord.FieldID, pq.sqlQuery()),
			sqlgraph.To(medicalrecordstaff.Table, medicalrecordstaff.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrecord.EdgesOfMedicalrecordstaffTable, patientrecord.EdgesOfMedicalrecordstaffColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfPrename chains the current query on the EdgesOfPrename edge.
func (pq *PatientrecordQuery) QueryEdgesOfPrename() *PrenameQuery {
	query := &PrenameQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrecord.Table, patientrecord.FieldID, pq.sqlQuery()),
			sqlgraph.To(prename.Table, prename.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, patientrecord.EdgesOfPrenameTable, patientrecord.EdgesOfPrenameColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfHistorytaking chains the current query on the EdgesOfHistorytaking edge.
func (pq *PatientrecordQuery) QueryEdgesOfHistorytaking() *HistorytakingQuery {
	query := &HistorytakingQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrecord.Table, patientrecord.FieldID, pq.sqlQuery()),
			sqlgraph.To(historytaking.Table, historytaking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patientrecord.EdgesOfHistorytakingTable, patientrecord.EdgesOfHistorytakingColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfTreatment chains the current query on the EdgesOfTreatment edge.
func (pq *PatientrecordQuery) QueryEdgesOfTreatment() *TreatmentQuery {
	query := &TreatmentQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrecord.Table, patientrecord.FieldID, pq.sqlQuery()),
			sqlgraph.To(treatment.Table, treatment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patientrecord.EdgesOfTreatmentTable, patientrecord.EdgesOfTreatmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEdgesOfPatientrecordPatientrights chains the current query on the EdgesOfPatientrecordPatientrights edge.
func (pq *PatientrecordQuery) QueryEdgesOfPatientrecordPatientrights() *PatientrightsQuery {
	query := &PatientrightsQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(patientrecord.Table, patientrecord.FieldID, pq.sqlQuery()),
			sqlgraph.To(patientrights.Table, patientrights.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, patientrecord.EdgesOfPatientrecordPatientrightsTable, patientrecord.EdgesOfPatientrecordPatientrightsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Patientrecord entity in the query. Returns *NotFoundError when no patientrecord was found.
func (pq *PatientrecordQuery) First(ctx context.Context) (*Patientrecord, error) {
	pas, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pas) == 0 {
		return nil, &NotFoundError{patientrecord.Label}
	}
	return pas[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PatientrecordQuery) FirstX(ctx context.Context) *Patientrecord {
	pa, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pa
}

// FirstID returns the first Patientrecord id in the query. Returns *NotFoundError when no id was found.
func (pq *PatientrecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{patientrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PatientrecordQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Patientrecord entity in the query, returns an error if not exactly one entity was returned.
func (pq *PatientrecordQuery) Only(ctx context.Context) (*Patientrecord, error) {
	pas, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pas) {
	case 1:
		return pas[0], nil
	case 0:
		return nil, &NotFoundError{patientrecord.Label}
	default:
		return nil, &NotSingularError{patientrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PatientrecordQuery) OnlyX(ctx context.Context) *Patientrecord {
	pa, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// OnlyID returns the only Patientrecord id in the query, returns an error if not exactly one id was returned.
func (pq *PatientrecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = &NotSingularError{patientrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PatientrecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Patientrecords.
func (pq *PatientrecordQuery) All(ctx context.Context) ([]*Patientrecord, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PatientrecordQuery) AllX(ctx context.Context) []*Patientrecord {
	pas, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pas
}

// IDs executes the query and returns a list of Patientrecord ids.
func (pq *PatientrecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(patientrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PatientrecordQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PatientrecordQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PatientrecordQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PatientrecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PatientrecordQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PatientrecordQuery) Clone() *PatientrecordQuery {
	return &PatientrecordQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Patientrecord{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

//  WithEdgesOfGender tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfGender" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrecordQuery) WithEdgesOfGender(opts ...func(*GenderQuery)) *PatientrecordQuery {
	query := &GenderQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfGender = query
	return pq
}

//  WithEdgesOfMedicalrecordstaff tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfMedicalrecordstaff" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrecordQuery) WithEdgesOfMedicalrecordstaff(opts ...func(*MedicalrecordstaffQuery)) *PatientrecordQuery {
	query := &MedicalrecordstaffQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfMedicalrecordstaff = query
	return pq
}

//  WithEdgesOfPrename tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfPrename" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrecordQuery) WithEdgesOfPrename(opts ...func(*PrenameQuery)) *PatientrecordQuery {
	query := &PrenameQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfPrename = query
	return pq
}

//  WithEdgesOfHistorytaking tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfHistorytaking" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrecordQuery) WithEdgesOfHistorytaking(opts ...func(*HistorytakingQuery)) *PatientrecordQuery {
	query := &HistorytakingQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfHistorytaking = query
	return pq
}

//  WithEdgesOfTreatment tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfTreatment" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrecordQuery) WithEdgesOfTreatment(opts ...func(*TreatmentQuery)) *PatientrecordQuery {
	query := &TreatmentQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfTreatment = query
	return pq
}

//  WithEdgesOfPatientrecordPatientrights tells the query-builder to eager-loads the nodes that are connected to
// the "EdgesOfPatientrecordPatientrights" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PatientrecordQuery) WithEdgesOfPatientrecordPatientrights(opts ...func(*PatientrightsQuery)) *PatientrecordQuery {
	query := &PatientrightsQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withEdgesOfPatientrecordPatientrights = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Patientrecord.Query().
//		GroupBy(patientrecord.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PatientrecordQuery) GroupBy(field string, fields ...string) *PatientrecordGroupBy {
	group := &PatientrecordGroupBy{config: pq.config}
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
//		Name string `json:"Name,omitempty"`
//	}
//
//	client.Patientrecord.Query().
//		Select(patientrecord.FieldName).
//		Scan(ctx, &v)
//
func (pq *PatientrecordQuery) Select(field string, fields ...string) *PatientrecordSelect {
	selector := &PatientrecordSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return selector
}

func (pq *PatientrecordQuery) prepareQuery(ctx context.Context) error {
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PatientrecordQuery) sqlAll(ctx context.Context) ([]*Patientrecord, error) {
	var (
		nodes       = []*Patientrecord{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [6]bool{
			pq.withEdgesOfGender != nil,
			pq.withEdgesOfMedicalrecordstaff != nil,
			pq.withEdgesOfPrename != nil,
			pq.withEdgesOfHistorytaking != nil,
			pq.withEdgesOfTreatment != nil,
			pq.withEdgesOfPatientrecordPatientrights != nil,
		}
	)
	if pq.withEdgesOfGender != nil || pq.withEdgesOfMedicalrecordstaff != nil || pq.withEdgesOfPrename != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, patientrecord.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Patientrecord{config: pq.config}
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

	if query := pq.withEdgesOfGender; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrecord)
		for i := range nodes {
			if fk := nodes[i].gender_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(gender.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "gender_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfGender = n
			}
		}
	}

	if query := pq.withEdgesOfMedicalrecordstaff; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrecord)
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
				nodes[i].Edges.EdgesOfMedicalrecordstaff = n
			}
		}
	}

	if query := pq.withEdgesOfPrename; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Patientrecord)
		for i := range nodes {
			if fk := nodes[i].prefix_id; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "prefix_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EdgesOfPrename = n
			}
		}
	}

	if query := pq.withEdgesOfHistorytaking; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Patientrecord)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Historytaking(func(s *sql.Selector) {
			s.Where(sql.InValues(patientrecord.EdgesOfHistorytakingColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.patientrecord_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "patientrecord_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patientrecord_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.EdgesOfHistorytaking = append(node.Edges.EdgesOfHistorytaking, n)
		}
	}

	if query := pq.withEdgesOfTreatment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Patientrecord)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Treatment(func(s *sql.Selector) {
			s.Where(sql.InValues(patientrecord.EdgesOfTreatmentColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.patientrecord_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "patientrecord_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patientrecord_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.EdgesOfTreatment = append(node.Edges.EdgesOfTreatment, n)
		}
	}

	if query := pq.withEdgesOfPatientrecordPatientrights; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Patientrecord)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Patientrights(func(s *sql.Selector) {
			s.Where(sql.InValues(patientrecord.EdgesOfPatientrecordPatientrightsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.patientrecord_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "patientrecord_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patientrecord_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.EdgesOfPatientrecordPatientrights = append(node.Edges.EdgesOfPatientrecordPatientrights, n)
		}
	}

	return nodes, nil
}

func (pq *PatientrecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PatientrecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PatientrecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientrecord.Table,
			Columns: patientrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrecord.FieldID,
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

func (pq *PatientrecordQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(patientrecord.Table)
	selector := builder.Select(t1.Columns(patientrecord.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(patientrecord.Columns...)...)
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

// PatientrecordGroupBy is the builder for group-by Patientrecord entities.
type PatientrecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PatientrecordGroupBy) Aggregate(fns ...AggregateFunc) *PatientrecordGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PatientrecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PatientrecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pgb *PatientrecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PatientrecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PatientrecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PatientrecordGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PatientrecordSelect is the builder for select fields of Patientrecord entities.
type PatientrecordSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PatientrecordSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ps.path(ctx)
	if err != nil {
		return err
	}
	ps.sql = query
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PatientrecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PatientrecordSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PatientrecordSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PatientrecordSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PatientrecordSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PatientrecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PatientrecordSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PatientrecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PatientrecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ps *PatientrecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{patientrecord.Label}
	default:
		err = fmt.Errorf("ent: PatientrecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PatientrecordSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PatientrecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PatientrecordSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}
