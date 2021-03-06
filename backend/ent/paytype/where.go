// Code generated by entc, DO NOT EDIT.

package paytype

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Paytype applies equality check predicate on the "Paytype" field. It's identical to PaytypeEQ.
func Paytype(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaytype), v))
	})
}

// PaytypeEQ applies the EQ predicate on the "Paytype" field.
func PaytypeEQ(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaytype), v))
	})
}

// PaytypeNEQ applies the NEQ predicate on the "Paytype" field.
func PaytypeNEQ(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaytype), v))
	})
}

// PaytypeIn applies the In predicate on the "Paytype" field.
func PaytypeIn(vs ...string) predicate.Paytype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Paytype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaytype), v...))
	})
}

// PaytypeNotIn applies the NotIn predicate on the "Paytype" field.
func PaytypeNotIn(vs ...string) predicate.Paytype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Paytype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaytype), v...))
	})
}

// PaytypeGT applies the GT predicate on the "Paytype" field.
func PaytypeGT(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaytype), v))
	})
}

// PaytypeGTE applies the GTE predicate on the "Paytype" field.
func PaytypeGTE(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaytype), v))
	})
}

// PaytypeLT applies the LT predicate on the "Paytype" field.
func PaytypeLT(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaytype), v))
	})
}

// PaytypeLTE applies the LTE predicate on the "Paytype" field.
func PaytypeLTE(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaytype), v))
	})
}

// PaytypeContains applies the Contains predicate on the "Paytype" field.
func PaytypeContains(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaytype), v))
	})
}

// PaytypeHasPrefix applies the HasPrefix predicate on the "Paytype" field.
func PaytypeHasPrefix(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaytype), v))
	})
}

// PaytypeHasSuffix applies the HasSuffix predicate on the "Paytype" field.
func PaytypeHasSuffix(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaytype), v))
	})
}

// PaytypeEqualFold applies the EqualFold predicate on the "Paytype" field.
func PaytypeEqualFold(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaytype), v))
	})
}

// PaytypeContainsFold applies the ContainsFold predicate on the "Paytype" field.
func PaytypeContainsFold(v string) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaytype), v))
	})
}

// HasEdgesOfBills applies the HasEdge predicate on the "EdgesOfBills" edge.
func HasEdgesOfBills() predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfBillsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgesOfBillsTable, EdgesOfBillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfBillsWith applies the HasEdge predicate on the "EdgesOfBills" edge with a given conditions (other predicates).
func HasEdgesOfBillsWith(preds ...predicate.Bill) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfBillsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgesOfBillsTable, EdgesOfBillsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Paytype) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Paytype) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Paytype) predicate.Paytype {
	return predicate.Paytype(func(s *sql.Selector) {
		p(s.Not())
	})
}
