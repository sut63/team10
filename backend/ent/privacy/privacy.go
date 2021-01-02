// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"context"
	"errors"
	"fmt"

	"github.com/theuo/app/ent"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = errors.New("ent/privacy: allow rule")

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = errors.New("ent/privacy: deny rule")

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = errors.New("ent/privacy: skip rule")
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

type decisionCtxKey struct{}

// DecisionContext creates a decision context.
func DecisionContext(parent context.Context, decision error) context.Context {
	if decision == nil || errors.Is(decision, Skip) {
		return parent
	}
	return context.WithValue(parent, decisionCtxKey{}, decision)
}

func decisionFromContext(ctx context.Context) (error, bool) {
	decision, ok := ctx.Value(decisionCtxKey{}).(error)
	if ok && errors.Is(decision, Allow) {
		decision = nil
	}
	return decision, ok
}

type (
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy []QueryRule

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule interface {
		EvalQuery(context.Context, ent.Query) error
	}
)

// EvalQuery evaluates a query against a query policy.
func (policy QueryPolicy) EvalQuery(ctx context.Context, q ent.Query) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalQuery(ctx, q); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

type (
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy []MutationRule

	// MutationRule defines the interface deciding whether a
	// mutation is allowed and optionally modify it.
	MutationRule interface {
		EvalMutation(context.Context, ent.Mutation) error
	}
)

// EvalMutation evaluates a mutation against a mutation policy.
func (policy MutationPolicy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if decision, ok := decisionFromContext(ctx); ok {
		return decision
	}
	for _, rule := range policy {
		switch decision := rule.EvalMutation(ctx, m); {
		case decision == nil || errors.Is(decision, Skip):
		case errors.Is(decision, Allow):
			return nil
		default:
			return decision
		}
	}
	return nil
}

// MutationRuleFunc type is an adapter to allow the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query policy.
func (policy Policy) EvalQuery(ctx context.Context, q ent.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutation policy.
func (policy Policy) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is the interface that groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AbilitypatientrightsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AbilitypatientrightsQueryRuleFunc func(context.Context, *ent.AbilitypatientrightsQuery) error

// EvalQuery return f(ctx, q).
func (f AbilitypatientrightsQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AbilitypatientrightsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AbilitypatientrightsQuery", q)
}

// The AbilitypatientrightsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AbilitypatientrightsMutationRuleFunc func(context.Context, *ent.AbilitypatientrightsMutation) error

// EvalMutation calls f(ctx, m).
func (f AbilitypatientrightsMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AbilitypatientrightsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AbilitypatientrightsMutation", m)
}

// The BillQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BillQueryRuleFunc func(context.Context, *ent.BillQuery) error

// EvalQuery return f(ctx, q).
func (f BillQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BillQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BillQuery", q)
}

// The BillMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BillMutationRuleFunc func(context.Context, *ent.BillMutation) error

// EvalMutation calls f(ctx, m).
func (f BillMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BillMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BillMutation", m)
}

// The DepartmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DepartmentQueryRuleFunc func(context.Context, *ent.DepartmentQuery) error

// EvalQuery return f(ctx, q).
func (f DepartmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DepartmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DepartmentQuery", q)
}

// The DepartmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DepartmentMutationRuleFunc func(context.Context, *ent.DepartmentMutation) error

// EvalMutation calls f(ctx, m).
func (f DepartmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DepartmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DepartmentMutation", m)
}

// The DoctorinfoQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DoctorinfoQueryRuleFunc func(context.Context, *ent.DoctorinfoQuery) error

// EvalQuery return f(ctx, q).
func (f DoctorinfoQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DoctorinfoQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DoctorinfoQuery", q)
}

// The DoctorinfoMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DoctorinfoMutationRuleFunc func(context.Context, *ent.DoctorinfoMutation) error

// EvalMutation calls f(ctx, m).
func (f DoctorinfoMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DoctorinfoMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DoctorinfoMutation", m)
}

// The EducationlevelQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EducationlevelQueryRuleFunc func(context.Context, *ent.EducationlevelQuery) error

// EvalQuery return f(ctx, q).
func (f EducationlevelQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EducationlevelQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.EducationlevelQuery", q)
}

// The EducationlevelMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EducationlevelMutationRuleFunc func(context.Context, *ent.EducationlevelMutation) error

// EvalMutation calls f(ctx, m).
func (f EducationlevelMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.EducationlevelMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.EducationlevelMutation", m)
}

// The FinancierQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FinancierQueryRuleFunc func(context.Context, *ent.FinancierQuery) error

// EvalQuery return f(ctx, q).
func (f FinancierQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FinancierQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FinancierQuery", q)
}

// The FinancierMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FinancierMutationRuleFunc func(context.Context, *ent.FinancierMutation) error

// EvalMutation calls f(ctx, m).
func (f FinancierMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FinancierMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FinancierMutation", m)
}

// The GenderQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GenderQueryRuleFunc func(context.Context, *ent.GenderQuery) error

// EvalQuery return f(ctx, q).
func (f GenderQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GenderQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GenderQuery", q)
}

// The GenderMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GenderMutationRuleFunc func(context.Context, *ent.GenderMutation) error

// EvalMutation calls f(ctx, m).
func (f GenderMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GenderMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GenderMutation", m)
}

// The HistorytakingQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type HistorytakingQueryRuleFunc func(context.Context, *ent.HistorytakingQuery) error

// EvalQuery return f(ctx, q).
func (f HistorytakingQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.HistorytakingQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.HistorytakingQuery", q)
}

// The HistorytakingMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type HistorytakingMutationRuleFunc func(context.Context, *ent.HistorytakingMutation) error

// EvalMutation calls f(ctx, m).
func (f HistorytakingMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.HistorytakingMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.HistorytakingMutation", m)
}

// The InsuranceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InsuranceQueryRuleFunc func(context.Context, *ent.InsuranceQuery) error

// EvalQuery return f(ctx, q).
func (f InsuranceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.InsuranceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.InsuranceQuery", q)
}

// The InsuranceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InsuranceMutationRuleFunc func(context.Context, *ent.InsuranceMutation) error

// EvalMutation calls f(ctx, m).
func (f InsuranceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.InsuranceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.InsuranceMutation", m)
}

// The MedicalrecordstaffQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MedicalrecordstaffQueryRuleFunc func(context.Context, *ent.MedicalrecordstaffQuery) error

// EvalQuery return f(ctx, q).
func (f MedicalrecordstaffQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MedicalrecordstaffQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MedicalrecordstaffQuery", q)
}

// The MedicalrecordstaffMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MedicalrecordstaffMutationRuleFunc func(context.Context, *ent.MedicalrecordstaffMutation) error

// EvalMutation calls f(ctx, m).
func (f MedicalrecordstaffMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MedicalrecordstaffMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MedicalrecordstaffMutation", m)
}

// The NurseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type NurseQueryRuleFunc func(context.Context, *ent.NurseQuery) error

// EvalQuery return f(ctx, q).
func (f NurseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NurseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.NurseQuery", q)
}

// The NurseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type NurseMutationRuleFunc func(context.Context, *ent.NurseMutation) error

// EvalMutation calls f(ctx, m).
func (f NurseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.NurseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.NurseMutation", m)
}

// The OfficeroomQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OfficeroomQueryRuleFunc func(context.Context, *ent.OfficeroomQuery) error

// EvalQuery return f(ctx, q).
func (f OfficeroomQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OfficeroomQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OfficeroomQuery", q)
}

// The OfficeroomMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OfficeroomMutationRuleFunc func(context.Context, *ent.OfficeroomMutation) error

// EvalMutation calls f(ctx, m).
func (f OfficeroomMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OfficeroomMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OfficeroomMutation", m)
}

// The PatientrecordQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PatientrecordQueryRuleFunc func(context.Context, *ent.PatientrecordQuery) error

// EvalQuery return f(ctx, q).
func (f PatientrecordQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PatientrecordQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PatientrecordQuery", q)
}

// The PatientrecordMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PatientrecordMutationRuleFunc func(context.Context, *ent.PatientrecordMutation) error

// EvalMutation calls f(ctx, m).
func (f PatientrecordMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PatientrecordMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PatientrecordMutation", m)
}

// The PatientrightsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PatientrightsQueryRuleFunc func(context.Context, *ent.PatientrightsQuery) error

// EvalQuery return f(ctx, q).
func (f PatientrightsQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PatientrightsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PatientrightsQuery", q)
}

// The PatientrightsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PatientrightsMutationRuleFunc func(context.Context, *ent.PatientrightsMutation) error

// EvalMutation calls f(ctx, m).
func (f PatientrightsMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PatientrightsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PatientrightsMutation", m)
}

// The PatientrightstypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PatientrightstypeQueryRuleFunc func(context.Context, *ent.PatientrightstypeQuery) error

// EvalQuery return f(ctx, q).
func (f PatientrightstypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PatientrightstypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PatientrightstypeQuery", q)
}

// The PatientrightstypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PatientrightstypeMutationRuleFunc func(context.Context, *ent.PatientrightstypeMutation) error

// EvalMutation calls f(ctx, m).
func (f PatientrightstypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PatientrightstypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PatientrightstypeMutation", m)
}

// The PaytypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PaytypeQueryRuleFunc func(context.Context, *ent.PaytypeQuery) error

// EvalQuery return f(ctx, q).
func (f PaytypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PaytypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PaytypeQuery", q)
}

// The PaytypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PaytypeMutationRuleFunc func(context.Context, *ent.PaytypeMutation) error

// EvalMutation calls f(ctx, m).
func (f PaytypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PaytypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PaytypeMutation", m)
}

// The PrenameQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PrenameQueryRuleFunc func(context.Context, *ent.PrenameQuery) error

// EvalQuery return f(ctx, q).
func (f PrenameQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PrenameQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PrenameQuery", q)
}

// The PrenameMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PrenameMutationRuleFunc func(context.Context, *ent.PrenameMutation) error

// EvalMutation calls f(ctx, m).
func (f PrenameMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PrenameMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PrenameMutation", m)
}

// The SymptomseverityQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SymptomseverityQueryRuleFunc func(context.Context, *ent.SymptomseverityQuery) error

// EvalQuery return f(ctx, q).
func (f SymptomseverityQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SymptomseverityQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SymptomseverityQuery", q)
}

// The SymptomseverityMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SymptomseverityMutationRuleFunc func(context.Context, *ent.SymptomseverityMutation) error

// EvalMutation calls f(ctx, m).
func (f SymptomseverityMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SymptomseverityMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SymptomseverityMutation", m)
}

// The TreatmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TreatmentQueryRuleFunc func(context.Context, *ent.TreatmentQuery) error

// EvalQuery return f(ctx, q).
func (f TreatmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TreatmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TreatmentQuery", q)
}

// The TreatmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TreatmentMutationRuleFunc func(context.Context, *ent.TreatmentMutation) error

// EvalMutation calls f(ctx, m).
func (f TreatmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TreatmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TreatmentMutation", m)
}

// The TypetreatmentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TypetreatmentQueryRuleFunc func(context.Context, *ent.TypetreatmentQuery) error

// EvalQuery return f(ctx, q).
func (f TypetreatmentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TypetreatmentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TypetreatmentQuery", q)
}

// The TypetreatmentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TypetreatmentMutationRuleFunc func(context.Context, *ent.TypetreatmentMutation) error

// EvalMutation calls f(ctx, m).
func (f TypetreatmentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TypetreatmentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TypetreatmentMutation", m)
}

// The UnpaybillQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UnpaybillQueryRuleFunc func(context.Context, *ent.UnpaybillQuery) error

// EvalQuery return f(ctx, q).
func (f UnpaybillQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UnpaybillQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UnpaybillQuery", q)
}

// The UnpaybillMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UnpaybillMutationRuleFunc func(context.Context, *ent.UnpaybillMutation) error

// EvalMutation calls f(ctx, m).
func (f UnpaybillMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UnpaybillMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UnpaybillMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}
