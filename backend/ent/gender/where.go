// Code generated by entc, DO NOT EDIT.

package gender

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Genderstatus applies equality check predicate on the "Genderstatus" field. It's identical to GenderstatusEQ.
func Genderstatus(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusEQ applies the EQ predicate on the "Genderstatus" field.
func GenderstatusEQ(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusNEQ applies the NEQ predicate on the "Genderstatus" field.
func GenderstatusNEQ(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusIn applies the In predicate on the "Genderstatus" field.
func GenderstatusIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGenderstatus), v...))
	})
}

// GenderstatusNotIn applies the NotIn predicate on the "Genderstatus" field.
func GenderstatusNotIn(vs ...string) predicate.Gender {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Gender(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGenderstatus), v...))
	})
}

// GenderstatusGT applies the GT predicate on the "Genderstatus" field.
func GenderstatusGT(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusGTE applies the GTE predicate on the "Genderstatus" field.
func GenderstatusGTE(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusLT applies the LT predicate on the "Genderstatus" field.
func GenderstatusLT(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusLTE applies the LTE predicate on the "Genderstatus" field.
func GenderstatusLTE(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusContains applies the Contains predicate on the "Genderstatus" field.
func GenderstatusContains(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusHasPrefix applies the HasPrefix predicate on the "Genderstatus" field.
func GenderstatusHasPrefix(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusHasSuffix applies the HasSuffix predicate on the "Genderstatus" field.
func GenderstatusHasSuffix(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusEqualFold applies the EqualFold predicate on the "Genderstatus" field.
func GenderstatusEqualFold(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGenderstatus), v))
	})
}

// GenderstatusContainsFold applies the ContainsFold predicate on the "Genderstatus" field.
func GenderstatusContainsFold(v string) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGenderstatus), v))
	})
}

// HasEdgesOfPatientrecord applies the HasEdge predicate on the "EdgesOfPatientrecord" edge.
func HasEdgesOfPatientrecord() predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgesOfPatientrecordTable, EdgesOfPatientrecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfPatientrecordWith applies the HasEdge predicate on the "EdgesOfPatientrecord" edge with a given conditions (other predicates).
func HasEdgesOfPatientrecordWith(preds ...predicate.Patientrecord) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgesOfPatientrecordTable, EdgesOfPatientrecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
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
func Not(p predicate.Gender) predicate.Gender {
	return predicate.Gender(func(s *sql.Selector) {
		p(s.Not())
	})
}
