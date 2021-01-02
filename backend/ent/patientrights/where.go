// Code generated by entc, DO NOT EDIT.

package patientrights

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/theuo/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PermissionDate applies equality check predicate on the "PermissionDate" field. It's identical to PermissionDateEQ.
func PermissionDate(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateEQ applies the EQ predicate on the "PermissionDate" field.
func PermissionDateEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateNEQ applies the NEQ predicate on the "PermissionDate" field.
func PermissionDateNEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateIn applies the In predicate on the "PermissionDate" field.
func PermissionDateIn(vs ...string) predicate.Patientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPermissionDate), v...))
	})
}

// PermissionDateNotIn applies the NotIn predicate on the "PermissionDate" field.
func PermissionDateNotIn(vs ...string) predicate.Patientrights {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Patientrights(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPermissionDate), v...))
	})
}

// PermissionDateGT applies the GT predicate on the "PermissionDate" field.
func PermissionDateGT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateGTE applies the GTE predicate on the "PermissionDate" field.
func PermissionDateGTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateLT applies the LT predicate on the "PermissionDate" field.
func PermissionDateLT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateLTE applies the LTE predicate on the "PermissionDate" field.
func PermissionDateLTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateContains applies the Contains predicate on the "PermissionDate" field.
func PermissionDateContains(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateHasPrefix applies the HasPrefix predicate on the "PermissionDate" field.
func PermissionDateHasPrefix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateHasSuffix applies the HasSuffix predicate on the "PermissionDate" field.
func PermissionDateHasSuffix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateEqualFold applies the EqualFold predicate on the "PermissionDate" field.
func PermissionDateEqualFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateContainsFold applies the ContainsFold predicate on the "PermissionDate" field.
func PermissionDateContainsFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPermissionDate), v))
	})
}

// HasPatientrightsPatientrightstype applies the HasEdge predicate on the "PatientrightsPatientrightstype" edge.
func HasPatientrightsPatientrightstype() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsPatientrightstypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsPatientrightstypeTable, PatientrightsPatientrightstypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientrightsPatientrightstypeWith applies the HasEdge predicate on the "PatientrightsPatientrightstype" edge with a given conditions (other predicates).
func HasPatientrightsPatientrightstypeWith(preds ...predicate.Patientrightstype) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsPatientrightstypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsPatientrightstypeTable, PatientrightsPatientrightstypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientrightsInsurance applies the HasEdge predicate on the "PatientrightsInsurance" edge.
func HasPatientrightsInsurance() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsInsuranceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsInsuranceTable, PatientrightsInsuranceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientrightsInsuranceWith applies the HasEdge predicate on the "PatientrightsInsurance" edge with a given conditions (other predicates).
func HasPatientrightsInsuranceWith(preds ...predicate.Insurance) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsInsuranceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsInsuranceTable, PatientrightsInsuranceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientrightsPatientrecord applies the HasEdge predicate on the "PatientrightsPatientrecord" edge.
func HasPatientrightsPatientrecord() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsPatientrecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsPatientrecordTable, PatientrightsPatientrecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientrightsPatientrecordWith applies the HasEdge predicate on the "PatientrightsPatientrecord" edge with a given conditions (other predicates).
func HasPatientrightsPatientrecordWith(preds ...predicate.Patientrecord) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsPatientrecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsPatientrecordTable, PatientrightsPatientrecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientrightsMedicalrecordstaff applies the HasEdge predicate on the "PatientrightsMedicalrecordstaff" edge.
func HasPatientrightsMedicalrecordstaff() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsMedicalrecordstaffTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsMedicalrecordstaffTable, PatientrightsMedicalrecordstaffColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientrightsMedicalrecordstaffWith applies the HasEdge predicate on the "PatientrightsMedicalrecordstaff" edge with a given conditions (other predicates).
func HasPatientrightsMedicalrecordstaffWith(preds ...predicate.Medicalrecordstaff) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrightsMedicalrecordstaffInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrightsMedicalrecordstaffTable, PatientrightsMedicalrecordstaffColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Patientrights) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Patientrights) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
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
func Not(p predicate.Patientrights) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		p(s.Not())
	})
}