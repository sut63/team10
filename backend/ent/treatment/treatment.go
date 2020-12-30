// Code generated by entc, DO NOT EDIT.

package treatment

const (
	// Label holds the string label denoting the treatment type in the database.
	Label = "treatment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTreatment holds the string denoting the treatment field in the database.
	FieldTreatment = "treatment"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"

	// EdgeTypetreatment holds the string denoting the typetreatment edge name in mutations.
	EdgeTypetreatment = "typetreatment"
	// EdgePatientrecord holds the string denoting the patientrecord edge name in mutations.
	EdgePatientrecord = "patientrecord"
	// EdgeDoctorinfo holds the string denoting the doctorinfo edge name in mutations.
	EdgeDoctorinfo = "doctorinfo"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "bills"

	// Table holds the table name of the treatment in the database.
	Table = "treatments"
	// TypetreatmentTable is the table the holds the typetreatment relation/edge.
	TypetreatmentTable = "treatments"
	// TypetreatmentInverseTable is the table name for the Typetreatment entity.
	// It exists in this package in order to avoid circular dependency with the "typetreatment" package.
	TypetreatmentInverseTable = "typetreatments"
	// TypetreatmentColumn is the table column denoting the typetreatment relation/edge.
	TypetreatmentColumn = "typetreatment_id"
	// PatientrecordTable is the table the holds the patientrecord relation/edge.
	PatientrecordTable = "treatments"
	// PatientrecordInverseTable is the table name for the Patientrecord entity.
	// It exists in this package in order to avoid circular dependency with the "patientrecord" package.
	PatientrecordInverseTable = "patientrecords"
	// PatientrecordColumn is the table column denoting the patientrecord relation/edge.
	PatientrecordColumn = "patientrecord_id"
	// DoctorinfoTable is the table the holds the doctorinfo relation/edge.
	DoctorinfoTable = "treatments"
	// DoctorinfoInverseTable is the table name for the Doctorinfo entity.
	// It exists in this package in order to avoid circular dependency with the "doctorinfo" package.
	DoctorinfoInverseTable = "doctorinfos"
	// DoctorinfoColumn is the table column denoting the doctorinfo relation/edge.
	DoctorinfoColumn = "doctorinfo_id"
	// BillsTable is the table the holds the bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the bills relation/edge.
	BillsColumn = "treatment_id"
)

// Columns holds all SQL columns for treatment fields.
var Columns = []string{
	FieldID,
	FieldTreatment,
	FieldDatetime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Treatment type.
var ForeignKeys = []string{
	"doctorinfo_id",
	"patientrecord_id",
	"typetreatment_id",
}

var (
	// TreatmentValidator is a validator for the "treatment" field. It is called by the builders before save.
	TreatmentValidator func(string) error
)
