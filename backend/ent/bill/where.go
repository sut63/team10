// Code generated by entc, DO NOT EDIT.

package bill

import (
	"time"

	"github.com/b6109868/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Amount applies equality check predicate on the "Amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Date applies equality check predicate on the "Date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// AmountEQ applies the EQ predicate on the "Amount" field.
func AmountEQ(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "Amount" field.
func AmountNEQ(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "Amount" field.
func AmountIn(vs ...string) predicate.Bill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "Amount" field.
func AmountNotIn(vs ...string) predicate.Bill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "Amount" field.
func AmountGT(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "Amount" field.
func AmountGTE(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "Amount" field.
func AmountLT(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "Amount" field.
func AmountLTE(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountContains applies the Contains predicate on the "Amount" field.
func AmountContains(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount), v))
	})
}

// AmountHasPrefix applies the HasPrefix predicate on the "Amount" field.
func AmountHasPrefix(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount), v))
	})
}

// AmountHasSuffix applies the HasSuffix predicate on the "Amount" field.
func AmountHasSuffix(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount), v))
	})
}

// AmountEqualFold applies the EqualFold predicate on the "Amount" field.
func AmountEqualFold(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount), v))
	})
}

// AmountContainsFold applies the ContainsFold predicate on the "Amount" field.
func AmountContainsFold(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount), v))
	})
}

// DateEQ applies the EQ predicate on the "Date" field.
func DateEQ(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "Date" field.
func DateNEQ(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "Date" field.
func DateIn(vs ...time.Time) predicate.Bill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "Date" field.
func DateNotIn(vs ...time.Time) predicate.Bill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Bill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "Date" field.
func DateGT(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "Date" field.
func DateGTE(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "Date" field.
func DateLT(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "Date" field.
func DateLTE(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// HasPaytype applies the HasEdge predicate on the "paytype" edge.
func HasPaytype() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaytypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaytypeTable, PaytypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaytypeWith applies the HasEdge predicate on the "paytype" edge with a given conditions (other predicates).
func HasPaytypeWith(preds ...predicate.Paytype) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaytypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PaytypeTable, PaytypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOfficer applies the HasEdge predicate on the "officer" edge.
func HasOfficer() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OfficerTable, OfficerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficerWith applies the HasEdge predicate on the "officer" edge with a given conditions (other predicates).
func HasOfficerWith(preds ...predicate.Financier) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OfficerTable, OfficerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTreatment applies the HasEdge predicate on the "treatment" edge.
func HasTreatment() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TreatmentTable, TreatmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTreatmentWith applies the HasEdge predicate on the "treatment" edge with a given conditions (other predicates).
func HasTreatmentWith(preds ...predicate.Unpaybill) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TreatmentTable, TreatmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Bill) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Bill) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
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
func Not(p predicate.Bill) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		p(s.Not())
	})
}
