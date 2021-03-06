// Code generated by entc, DO NOT EDIT.

package patientrights

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
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
func PermissionDate(v time.Time) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionDate), v))
	})
}

// Permission applies equality check predicate on the "Permission" field. It's identical to PermissionEQ.
func Permission(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermission), v))
	})
}

// PermissionArea applies equality check predicate on the "PermissionArea" field. It's identical to PermissionAreaEQ.
func PermissionArea(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionArea), v))
	})
}

// Responsible applies equality check predicate on the "Responsible" field. It's identical to ResponsibleEQ.
func Responsible(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponsible), v))
	})
}

// PermissionDateEQ applies the EQ predicate on the "PermissionDate" field.
func PermissionDateEQ(v time.Time) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateNEQ applies the NEQ predicate on the "PermissionDate" field.
func PermissionDateNEQ(v time.Time) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateIn applies the In predicate on the "PermissionDate" field.
func PermissionDateIn(vs ...time.Time) predicate.Patientrights {
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
func PermissionDateNotIn(vs ...time.Time) predicate.Patientrights {
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
func PermissionDateGT(v time.Time) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateGTE applies the GTE predicate on the "PermissionDate" field.
func PermissionDateGTE(v time.Time) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateLT applies the LT predicate on the "PermissionDate" field.
func PermissionDateLT(v time.Time) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermissionDate), v))
	})
}

// PermissionDateLTE applies the LTE predicate on the "PermissionDate" field.
func PermissionDateLTE(v time.Time) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermissionDate), v))
	})
}

// PermissionEQ applies the EQ predicate on the "Permission" field.
func PermissionEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermission), v))
	})
}

// PermissionNEQ applies the NEQ predicate on the "Permission" field.
func PermissionNEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermission), v))
	})
}

// PermissionIn applies the In predicate on the "Permission" field.
func PermissionIn(vs ...string) predicate.Patientrights {
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
		s.Where(sql.In(s.C(FieldPermission), v...))
	})
}

// PermissionNotIn applies the NotIn predicate on the "Permission" field.
func PermissionNotIn(vs ...string) predicate.Patientrights {
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
		s.Where(sql.NotIn(s.C(FieldPermission), v...))
	})
}

// PermissionGT applies the GT predicate on the "Permission" field.
func PermissionGT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermission), v))
	})
}

// PermissionGTE applies the GTE predicate on the "Permission" field.
func PermissionGTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermission), v))
	})
}

// PermissionLT applies the LT predicate on the "Permission" field.
func PermissionLT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermission), v))
	})
}

// PermissionLTE applies the LTE predicate on the "Permission" field.
func PermissionLTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermission), v))
	})
}

// PermissionContains applies the Contains predicate on the "Permission" field.
func PermissionContains(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPermission), v))
	})
}

// PermissionHasPrefix applies the HasPrefix predicate on the "Permission" field.
func PermissionHasPrefix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPermission), v))
	})
}

// PermissionHasSuffix applies the HasSuffix predicate on the "Permission" field.
func PermissionHasSuffix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPermission), v))
	})
}

// PermissionEqualFold applies the EqualFold predicate on the "Permission" field.
func PermissionEqualFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPermission), v))
	})
}

// PermissionContainsFold applies the ContainsFold predicate on the "Permission" field.
func PermissionContainsFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPermission), v))
	})
}

// PermissionAreaEQ applies the EQ predicate on the "PermissionArea" field.
func PermissionAreaEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaNEQ applies the NEQ predicate on the "PermissionArea" field.
func PermissionAreaNEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaIn applies the In predicate on the "PermissionArea" field.
func PermissionAreaIn(vs ...string) predicate.Patientrights {
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
		s.Where(sql.In(s.C(FieldPermissionArea), v...))
	})
}

// PermissionAreaNotIn applies the NotIn predicate on the "PermissionArea" field.
func PermissionAreaNotIn(vs ...string) predicate.Patientrights {
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
		s.Where(sql.NotIn(s.C(FieldPermissionArea), v...))
	})
}

// PermissionAreaGT applies the GT predicate on the "PermissionArea" field.
func PermissionAreaGT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaGTE applies the GTE predicate on the "PermissionArea" field.
func PermissionAreaGTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaLT applies the LT predicate on the "PermissionArea" field.
func PermissionAreaLT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaLTE applies the LTE predicate on the "PermissionArea" field.
func PermissionAreaLTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaContains applies the Contains predicate on the "PermissionArea" field.
func PermissionAreaContains(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaHasPrefix applies the HasPrefix predicate on the "PermissionArea" field.
func PermissionAreaHasPrefix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaHasSuffix applies the HasSuffix predicate on the "PermissionArea" field.
func PermissionAreaHasSuffix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaEqualFold applies the EqualFold predicate on the "PermissionArea" field.
func PermissionAreaEqualFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPermissionArea), v))
	})
}

// PermissionAreaContainsFold applies the ContainsFold predicate on the "PermissionArea" field.
func PermissionAreaContainsFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPermissionArea), v))
	})
}

// ResponsibleEQ applies the EQ predicate on the "Responsible" field.
func ResponsibleEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResponsible), v))
	})
}

// ResponsibleNEQ applies the NEQ predicate on the "Responsible" field.
func ResponsibleNEQ(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResponsible), v))
	})
}

// ResponsibleIn applies the In predicate on the "Responsible" field.
func ResponsibleIn(vs ...string) predicate.Patientrights {
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
		s.Where(sql.In(s.C(FieldResponsible), v...))
	})
}

// ResponsibleNotIn applies the NotIn predicate on the "Responsible" field.
func ResponsibleNotIn(vs ...string) predicate.Patientrights {
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
		s.Where(sql.NotIn(s.C(FieldResponsible), v...))
	})
}

// ResponsibleGT applies the GT predicate on the "Responsible" field.
func ResponsibleGT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResponsible), v))
	})
}

// ResponsibleGTE applies the GTE predicate on the "Responsible" field.
func ResponsibleGTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResponsible), v))
	})
}

// ResponsibleLT applies the LT predicate on the "Responsible" field.
func ResponsibleLT(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResponsible), v))
	})
}

// ResponsibleLTE applies the LTE predicate on the "Responsible" field.
func ResponsibleLTE(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResponsible), v))
	})
}

// ResponsibleContains applies the Contains predicate on the "Responsible" field.
func ResponsibleContains(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResponsible), v))
	})
}

// ResponsibleHasPrefix applies the HasPrefix predicate on the "Responsible" field.
func ResponsibleHasPrefix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResponsible), v))
	})
}

// ResponsibleHasSuffix applies the HasSuffix predicate on the "Responsible" field.
func ResponsibleHasSuffix(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResponsible), v))
	})
}

// ResponsibleEqualFold applies the EqualFold predicate on the "Responsible" field.
func ResponsibleEqualFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResponsible), v))
	})
}

// ResponsibleContainsFold applies the ContainsFold predicate on the "Responsible" field.
func ResponsibleContainsFold(v string) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResponsible), v))
	})
}

// HasEdgesOfPatientrightsAbilitypatientrights applies the HasEdge predicate on the "EdgesOfPatientrightsAbilitypatientrights" edge.
func HasEdgesOfPatientrightsAbilitypatientrights() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsAbilitypatientrightsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsAbilitypatientrightsTable, EdgesOfPatientrightsAbilitypatientrightsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfPatientrightsAbilitypatientrightsWith applies the HasEdge predicate on the "EdgesOfPatientrightsAbilitypatientrights" edge with a given conditions (other predicates).
func HasEdgesOfPatientrightsAbilitypatientrightsWith(preds ...predicate.Abilitypatientrights) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsAbilitypatientrightsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsAbilitypatientrightsTable, EdgesOfPatientrightsAbilitypatientrightsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgesOfPatientrightsInsurance applies the HasEdge predicate on the "EdgesOfPatientrightsInsurance" edge.
func HasEdgesOfPatientrightsInsurance() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsInsuranceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsInsuranceTable, EdgesOfPatientrightsInsuranceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfPatientrightsInsuranceWith applies the HasEdge predicate on the "EdgesOfPatientrightsInsurance" edge with a given conditions (other predicates).
func HasEdgesOfPatientrightsInsuranceWith(preds ...predicate.Insurance) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsInsuranceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsInsuranceTable, EdgesOfPatientrightsInsuranceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgesOfPatientrightsPatientrecord applies the HasEdge predicate on the "EdgesOfPatientrightsPatientrecord" edge.
func HasEdgesOfPatientrightsPatientrecord() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsPatientrecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsPatientrecordTable, EdgesOfPatientrightsPatientrecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfPatientrightsPatientrecordWith applies the HasEdge predicate on the "EdgesOfPatientrightsPatientrecord" edge with a given conditions (other predicates).
func HasEdgesOfPatientrightsPatientrecordWith(preds ...predicate.Patientrecord) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsPatientrecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsPatientrecordTable, EdgesOfPatientrightsPatientrecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgesOfPatientrightsMedicalrecordstaff applies the HasEdge predicate on the "EdgesOfPatientrightsMedicalrecordstaff" edge.
func HasEdgesOfPatientrightsMedicalrecordstaff() predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsMedicalrecordstaffTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsMedicalrecordstaffTable, EdgesOfPatientrightsMedicalrecordstaffColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfPatientrightsMedicalrecordstaffWith applies the HasEdge predicate on the "EdgesOfPatientrightsMedicalrecordstaff" edge with a given conditions (other predicates).
func HasEdgesOfPatientrightsMedicalrecordstaffWith(preds ...predicate.Medicalrecordstaff) predicate.Patientrights {
	return predicate.Patientrights(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPatientrightsMedicalrecordstaffInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPatientrightsMedicalrecordstaffTable, EdgesOfPatientrightsMedicalrecordstaffColumn),
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
