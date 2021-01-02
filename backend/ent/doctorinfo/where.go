// Code generated by entc, DO NOT EDIT.

package doctorinfo

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Doctorname applies equality check predicate on the "doctorname" field. It's identical to DoctornameEQ.
func Doctorname(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorname), v))
	})
}

// Doctorsurname applies equality check predicate on the "doctorsurname" field. It's identical to DoctorsurnameEQ.
func Doctorsurname(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorsurname), v))
	})
}

// Telephonenumber applies equality check predicate on the "telephonenumber" field. It's identical to TelephonenumberEQ.
func Telephonenumber(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephonenumber), v))
	})
}

// Licensenumber applies equality check predicate on the "licensenumber" field. It's identical to LicensenumberEQ.
func Licensenumber(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLicensenumber), v))
	})
}

// DoctornameEQ applies the EQ predicate on the "doctorname" field.
func DoctornameEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorname), v))
	})
}

// DoctornameNEQ applies the NEQ predicate on the "doctorname" field.
func DoctornameNEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorname), v))
	})
}

// DoctornameIn applies the In predicate on the "doctorname" field.
func DoctornameIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorname), v...))
	})
}

// DoctornameNotIn applies the NotIn predicate on the "doctorname" field.
func DoctornameNotIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorname), v...))
	})
}

// DoctornameGT applies the GT predicate on the "doctorname" field.
func DoctornameGT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorname), v))
	})
}

// DoctornameGTE applies the GTE predicate on the "doctorname" field.
func DoctornameGTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorname), v))
	})
}

// DoctornameLT applies the LT predicate on the "doctorname" field.
func DoctornameLT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorname), v))
	})
}

// DoctornameLTE applies the LTE predicate on the "doctorname" field.
func DoctornameLTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorname), v))
	})
}

// DoctornameContains applies the Contains predicate on the "doctorname" field.
func DoctornameContains(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorname), v))
	})
}

// DoctornameHasPrefix applies the HasPrefix predicate on the "doctorname" field.
func DoctornameHasPrefix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorname), v))
	})
}

// DoctornameHasSuffix applies the HasSuffix predicate on the "doctorname" field.
func DoctornameHasSuffix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorname), v))
	})
}

// DoctornameEqualFold applies the EqualFold predicate on the "doctorname" field.
func DoctornameEqualFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorname), v))
	})
}

// DoctornameContainsFold applies the ContainsFold predicate on the "doctorname" field.
func DoctornameContainsFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorname), v))
	})
}

// DoctorsurnameEQ applies the EQ predicate on the "doctorsurname" field.
func DoctorsurnameEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameNEQ applies the NEQ predicate on the "doctorsurname" field.
func DoctorsurnameNEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameIn applies the In predicate on the "doctorsurname" field.
func DoctorsurnameIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorsurname), v...))
	})
}

// DoctorsurnameNotIn applies the NotIn predicate on the "doctorsurname" field.
func DoctorsurnameNotIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorsurname), v...))
	})
}

// DoctorsurnameGT applies the GT predicate on the "doctorsurname" field.
func DoctorsurnameGT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameGTE applies the GTE predicate on the "doctorsurname" field.
func DoctorsurnameGTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameLT applies the LT predicate on the "doctorsurname" field.
func DoctorsurnameLT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameLTE applies the LTE predicate on the "doctorsurname" field.
func DoctorsurnameLTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameContains applies the Contains predicate on the "doctorsurname" field.
func DoctorsurnameContains(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameHasPrefix applies the HasPrefix predicate on the "doctorsurname" field.
func DoctorsurnameHasPrefix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameHasSuffix applies the HasSuffix predicate on the "doctorsurname" field.
func DoctorsurnameHasSuffix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameEqualFold applies the EqualFold predicate on the "doctorsurname" field.
func DoctorsurnameEqualFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorsurname), v))
	})
}

// DoctorsurnameContainsFold applies the ContainsFold predicate on the "doctorsurname" field.
func DoctorsurnameContainsFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorsurname), v))
	})
}

// TelephonenumberEQ applies the EQ predicate on the "telephonenumber" field.
func TelephonenumberEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberNEQ applies the NEQ predicate on the "telephonenumber" field.
func TelephonenumberNEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberIn applies the In predicate on the "telephonenumber" field.
func TelephonenumberIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTelephonenumber), v...))
	})
}

// TelephonenumberNotIn applies the NotIn predicate on the "telephonenumber" field.
func TelephonenumberNotIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTelephonenumber), v...))
	})
}

// TelephonenumberGT applies the GT predicate on the "telephonenumber" field.
func TelephonenumberGT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberGTE applies the GTE predicate on the "telephonenumber" field.
func TelephonenumberGTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberLT applies the LT predicate on the "telephonenumber" field.
func TelephonenumberLT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberLTE applies the LTE predicate on the "telephonenumber" field.
func TelephonenumberLTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberContains applies the Contains predicate on the "telephonenumber" field.
func TelephonenumberContains(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberHasPrefix applies the HasPrefix predicate on the "telephonenumber" field.
func TelephonenumberHasPrefix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberHasSuffix applies the HasSuffix predicate on the "telephonenumber" field.
func TelephonenumberHasSuffix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberEqualFold applies the EqualFold predicate on the "telephonenumber" field.
func TelephonenumberEqualFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTelephonenumber), v))
	})
}

// TelephonenumberContainsFold applies the ContainsFold predicate on the "telephonenumber" field.
func TelephonenumberContainsFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTelephonenumber), v))
	})
}

// LicensenumberEQ applies the EQ predicate on the "licensenumber" field.
func LicensenumberEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberNEQ applies the NEQ predicate on the "licensenumber" field.
func LicensenumberNEQ(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberIn applies the In predicate on the "licensenumber" field.
func LicensenumberIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLicensenumber), v...))
	})
}

// LicensenumberNotIn applies the NotIn predicate on the "licensenumber" field.
func LicensenumberNotIn(vs ...string) predicate.Doctorinfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Doctorinfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLicensenumber), v...))
	})
}

// LicensenumberGT applies the GT predicate on the "licensenumber" field.
func LicensenumberGT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberGTE applies the GTE predicate on the "licensenumber" field.
func LicensenumberGTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberLT applies the LT predicate on the "licensenumber" field.
func LicensenumberLT(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberLTE applies the LTE predicate on the "licensenumber" field.
func LicensenumberLTE(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberContains applies the Contains predicate on the "licensenumber" field.
func LicensenumberContains(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberHasPrefix applies the HasPrefix predicate on the "licensenumber" field.
func LicensenumberHasPrefix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberHasSuffix applies the HasSuffix predicate on the "licensenumber" field.
func LicensenumberHasSuffix(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberEqualFold applies the EqualFold predicate on the "licensenumber" field.
func LicensenumberEqualFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLicensenumber), v))
	})
}

// LicensenumberContainsFold applies the ContainsFold predicate on the "licensenumber" field.
func LicensenumberContainsFold(v string) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLicensenumber), v))
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEducationlevel applies the HasEdge predicate on the "educationlevel" edge.
func HasEducationlevel() predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EducationlevelTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EducationlevelTable, EducationlevelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEducationlevelWith applies the HasEdge predicate on the "educationlevel" edge with a given conditions (other predicates).
func HasEducationlevelWith(preds ...predicate.Educationlevel) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EducationlevelInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EducationlevelTable, EducationlevelColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOfficeroom applies the HasEdge predicate on the "officeroom" edge.
func HasOfficeroom() predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficeroomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OfficeroomTable, OfficeroomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficeroomWith applies the HasEdge predicate on the "officeroom" edge with a given conditions (other predicates).
func HasOfficeroomWith(preds ...predicate.Officeroom) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficeroomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OfficeroomTable, OfficeroomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrename applies the HasEdge predicate on the "prename" edge.
func HasPrename() predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrenameTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrenameTable, PrenameColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrenameWith applies the HasEdge predicate on the "prename" edge with a given conditions (other predicates).
func HasPrenameWith(preds ...predicate.Prename) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrenameInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrenameTable, PrenameColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTreatment applies the HasEdge predicate on the "treatment" edge.
func HasTreatment() predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TreatmentTable, TreatmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTreatmentWith applies the HasEdge predicate on the "treatment" edge with a given conditions (other predicates).
func HasTreatmentWith(preds ...predicate.Treatment) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TreatmentTable, TreatmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Doctorinfo) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Doctorinfo) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
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
func Not(p predicate.Doctorinfo) predicate.Doctorinfo {
	return predicate.Doctorinfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
