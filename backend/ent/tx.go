// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"github.com/facebookincubator/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Abilitypatientrights is the client for interacting with the Abilitypatientrights builders.
	Abilitypatientrights *AbilitypatientrightsClient
	// Bill is the client for interacting with the Bill builders.
	Bill *BillClient
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// Doctorinfo is the client for interacting with the Doctorinfo builders.
	Doctorinfo *DoctorinfoClient
	// Educationlevel is the client for interacting with the Educationlevel builders.
	Educationlevel *EducationlevelClient
	// Financier is the client for interacting with the Financier builders.
	Financier *FinancierClient
	// Gender is the client for interacting with the Gender builders.
	Gender *GenderClient
	// Historytaking is the client for interacting with the Historytaking builders.
	Historytaking *HistorytakingClient
	// Insurance is the client for interacting with the Insurance builders.
	Insurance *InsuranceClient
	// Medicalrecordstaff is the client for interacting with the Medicalrecordstaff builders.
	Medicalrecordstaff *MedicalrecordstaffClient
	// Nurse is the client for interacting with the Nurse builders.
	Nurse *NurseClient
	// Officeroom is the client for interacting with the Officeroom builders.
	Officeroom *OfficeroomClient
	// Patientrecord is the client for interacting with the Patientrecord builders.
	Patientrecord *PatientrecordClient
	// Patientrights is the client for interacting with the Patientrights builders.
	Patientrights *PatientrightsClient
	// Patientrightstype is the client for interacting with the Patientrightstype builders.
	Patientrightstype *PatientrightstypeClient
	// Paytype is the client for interacting with the Paytype builders.
	Paytype *PaytypeClient
	// Prename is the client for interacting with the Prename builders.
	Prename *PrenameClient
	// Symptomseverity is the client for interacting with the Symptomseverity builders.
	Symptomseverity *SymptomseverityClient
	// Treatment is the client for interacting with the Treatment builders.
	Treatment *TreatmentClient
	// Typetreatment is the client for interacting with the Typetreatment builders.
	Typetreatment *TypetreatmentClient
	// User is the client for interacting with the User builders.
	User *UserClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Committer method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollbacker method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Abilitypatientrights = NewAbilitypatientrightsClient(tx.config)
	tx.Bill = NewBillClient(tx.config)
	tx.Department = NewDepartmentClient(tx.config)
	tx.Doctorinfo = NewDoctorinfoClient(tx.config)
	tx.Educationlevel = NewEducationlevelClient(tx.config)
	tx.Financier = NewFinancierClient(tx.config)
	tx.Gender = NewGenderClient(tx.config)
	tx.Historytaking = NewHistorytakingClient(tx.config)
	tx.Insurance = NewInsuranceClient(tx.config)
	tx.Medicalrecordstaff = NewMedicalrecordstaffClient(tx.config)
	tx.Nurse = NewNurseClient(tx.config)
	tx.Officeroom = NewOfficeroomClient(tx.config)
	tx.Patientrecord = NewPatientrecordClient(tx.config)
	tx.Patientrights = NewPatientrightsClient(tx.config)
	tx.Patientrightstype = NewPatientrightstypeClient(tx.config)
	tx.Paytype = NewPaytypeClient(tx.config)
	tx.Prename = NewPrenameClient(tx.config)
	tx.Symptomseverity = NewSymptomseverityClient(tx.config)
	tx.Treatment = NewTreatmentClient(tx.config)
	tx.Typetreatment = NewTypetreatmentClient(tx.config)
	tx.User = NewUserClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Abilitypatientrights.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
