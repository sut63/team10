// Code generated by entc, DO NOT EDIT.

package bloodtype

const (
	// Label holds the string label denoting the bloodtype type in the database.
	Label = "bloodtype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBloodtype holds the string denoting the bloodtype field in the database.
	FieldBloodtype = "bloodtype"

	// EdgeEdgesOfPatientrecord holds the string denoting the edgesofpatientrecord edge name in mutations.
	EdgeEdgesOfPatientrecord = "EdgesOfPatientrecord"

	// Table holds the table name of the bloodtype in the database.
	Table = "bloodtypes"
	// EdgesOfPatientrecordTable is the table the holds the EdgesOfPatientrecord relation/edge.
	EdgesOfPatientrecordTable = "patientrecords"
	// EdgesOfPatientrecordInverseTable is the table name for the Patientrecord entity.
	// It exists in this package in order to avoid circular dependency with the "patientrecord" package.
	EdgesOfPatientrecordInverseTable = "patientrecords"
	// EdgesOfPatientrecordColumn is the table column denoting the EdgesOfPatientrecord relation/edge.
	EdgesOfPatientrecordColumn = "bloodtype_id"
)

// Columns holds all SQL columns for bloodtype fields.
var Columns = []string{
	FieldID,
	FieldBloodtype,
}