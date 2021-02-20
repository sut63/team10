// Code generated by entc, DO NOT EDIT.

package bill

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
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
func Amount(v int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Payercontact applies equality check predicate on the "Payercontact" field. It's identical to PayercontactEQ.
func Payercontact(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayercontact), v))
	})
}

// Note applies equality check predicate on the "Note" field. It's identical to NoteEQ.
func Note(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// Date applies equality check predicate on the "Date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// AmountEQ applies the EQ predicate on the "Amount" field.
func AmountEQ(v int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "Amount" field.
func AmountNEQ(v int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "Amount" field.
func AmountIn(vs ...int) predicate.Bill {
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
func AmountNotIn(vs ...int) predicate.Bill {
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
func AmountGT(v int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "Amount" field.
func AmountGTE(v int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "Amount" field.
func AmountLT(v int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "Amount" field.
func AmountLTE(v int) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// PayercontactEQ applies the EQ predicate on the "Payercontact" field.
func PayercontactEQ(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPayercontact), v))
	})
}

// PayercontactNEQ applies the NEQ predicate on the "Payercontact" field.
func PayercontactNEQ(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPayercontact), v))
	})
}

// PayercontactIn applies the In predicate on the "Payercontact" field.
func PayercontactIn(vs ...string) predicate.Bill {
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
		s.Where(sql.In(s.C(FieldPayercontact), v...))
	})
}

// PayercontactNotIn applies the NotIn predicate on the "Payercontact" field.
func PayercontactNotIn(vs ...string) predicate.Bill {
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
		s.Where(sql.NotIn(s.C(FieldPayercontact), v...))
	})
}

// PayercontactGT applies the GT predicate on the "Payercontact" field.
func PayercontactGT(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPayercontact), v))
	})
}

// PayercontactGTE applies the GTE predicate on the "Payercontact" field.
func PayercontactGTE(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPayercontact), v))
	})
}

// PayercontactLT applies the LT predicate on the "Payercontact" field.
func PayercontactLT(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPayercontact), v))
	})
}

// PayercontactLTE applies the LTE predicate on the "Payercontact" field.
func PayercontactLTE(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPayercontact), v))
	})
}

// PayercontactContains applies the Contains predicate on the "Payercontact" field.
func PayercontactContains(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPayercontact), v))
	})
}

// PayercontactHasPrefix applies the HasPrefix predicate on the "Payercontact" field.
func PayercontactHasPrefix(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPayercontact), v))
	})
}

// PayercontactHasSuffix applies the HasSuffix predicate on the "Payercontact" field.
func PayercontactHasSuffix(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPayercontact), v))
	})
}

// PayercontactEqualFold applies the EqualFold predicate on the "Payercontact" field.
func PayercontactEqualFold(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPayercontact), v))
	})
}

// PayercontactContainsFold applies the ContainsFold predicate on the "Payercontact" field.
func PayercontactContainsFold(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPayercontact), v))
	})
}

// NoteEQ applies the EQ predicate on the "Note" field.
func NoteEQ(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "Note" field.
func NoteNEQ(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "Note" field.
func NoteIn(vs ...string) predicate.Bill {
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
		s.Where(sql.In(s.C(FieldNote), v...))
	})
}

// NoteNotIn applies the NotIn predicate on the "Note" field.
func NoteNotIn(vs ...string) predicate.Bill {
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
		s.Where(sql.NotIn(s.C(FieldNote), v...))
	})
}

// NoteGT applies the GT predicate on the "Note" field.
func NoteGT(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "Note" field.
func NoteGTE(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "Note" field.
func NoteLT(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "Note" field.
func NoteLTE(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "Note" field.
func NoteContains(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "Note" field.
func NoteHasPrefix(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "Note" field.
func NoteHasSuffix(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "Note" field.
func NoteEqualFold(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "Note" field.
func NoteContainsFold(v string) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
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

// HasEdgesOfPaytype applies the HasEdge predicate on the "EdgesOfPaytype" edge.
func HasEdgesOfPaytype() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPaytypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPaytypeTable, EdgesOfPaytypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfPaytypeWith applies the HasEdge predicate on the "EdgesOfPaytype" edge with a given conditions (other predicates).
func HasEdgesOfPaytypeWith(preds ...predicate.Paytype) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfPaytypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfPaytypeTable, EdgesOfPaytypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgesOfOfficer applies the HasEdge predicate on the "EdgesOfOfficer" edge.
func HasEdgesOfOfficer() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfOfficerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfOfficerTable, EdgesOfOfficerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfOfficerWith applies the HasEdge predicate on the "EdgesOfOfficer" edge with a given conditions (other predicates).
func HasEdgesOfOfficerWith(preds ...predicate.Financier) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfOfficerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EdgesOfOfficerTable, EdgesOfOfficerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdgesOfUnpaybill applies the HasEdge predicate on the "EdgesOfUnpaybill" edge.
func HasEdgesOfUnpaybill() predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfUnpaybillTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EdgesOfUnpaybillTable, EdgesOfUnpaybillColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesOfUnpaybillWith applies the HasEdge predicate on the "EdgesOfUnpaybill" edge with a given conditions (other predicates).
func HasEdgesOfUnpaybillWith(preds ...predicate.Unpaybill) predicate.Bill {
	return predicate.Bill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EdgesOfUnpaybillInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EdgesOfUnpaybillTable, EdgesOfUnpaybillColumn),
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
