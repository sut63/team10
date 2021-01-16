// Code generated by entc, DO NOT EDIT.

package userstatus

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Userstatus applies equality check predicate on the "Userstatus" field. It's identical to UserstatusEQ.
func Userstatus(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserstatus), v))
	})
}

// UserstatusEQ applies the EQ predicate on the "Userstatus" field.
func UserstatusEQ(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserstatus), v))
	})
}

// UserstatusNEQ applies the NEQ predicate on the "Userstatus" field.
func UserstatusNEQ(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserstatus), v))
	})
}

// UserstatusIn applies the In predicate on the "Userstatus" field.
func UserstatusIn(vs ...string) predicate.Userstatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Userstatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserstatus), v...))
	})
}

// UserstatusNotIn applies the NotIn predicate on the "Userstatus" field.
func UserstatusNotIn(vs ...string) predicate.Userstatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Userstatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserstatus), v...))
	})
}

// UserstatusGT applies the GT predicate on the "Userstatus" field.
func UserstatusGT(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserstatus), v))
	})
}

// UserstatusGTE applies the GTE predicate on the "Userstatus" field.
func UserstatusGTE(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserstatus), v))
	})
}

// UserstatusLT applies the LT predicate on the "Userstatus" field.
func UserstatusLT(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserstatus), v))
	})
}

// UserstatusLTE applies the LTE predicate on the "Userstatus" field.
func UserstatusLTE(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserstatus), v))
	})
}

// UserstatusContains applies the Contains predicate on the "Userstatus" field.
func UserstatusContains(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserstatus), v))
	})
}

// UserstatusHasPrefix applies the HasPrefix predicate on the "Userstatus" field.
func UserstatusHasPrefix(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserstatus), v))
	})
}

// UserstatusHasSuffix applies the HasSuffix predicate on the "Userstatus" field.
func UserstatusHasSuffix(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserstatus), v))
	})
}

// UserstatusEqualFold applies the EqualFold predicate on the "Userstatus" field.
func UserstatusEqualFold(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserstatus), v))
	})
}

// UserstatusContainsFold applies the ContainsFold predicate on the "Userstatus" field.
func UserstatusContainsFold(v string) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserstatus), v))
	})
}

// HasEdgesOfUser applies the HasEdge predicate on the "EdgesOfUser" edge.
func HasEdgesOfUser() predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgesOfUserTable, EdgesOfUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfUserWith applies the HasEdge predicate on the "EdgesOfUser" edge with a given conditions (other predicates).
func HasEdgesOfUserWith(preds ...predicate.User) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgesOfUserTable, EdgesOfUserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Userstatus) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Userstatus) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
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
func Not(p predicate.Userstatus) predicate.Userstatus {
	return predicate.Userstatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
