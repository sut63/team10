// Code generated by entc, DO NOT EDIT.

package gender

const (
	// Label holds the string label denoting the gender type in the database.
	Label = "gender"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGenderstatus holds the string denoting the genderstatus field in the database.
	FieldGenderstatus = "genderstatus"

	// EdgePatientrecord holds the string denoting the patientrecord edge name in mutations.
	EdgePatientrecord = "patientrecord"

	// Table holds the table name of the gender in the database.
	Table = "genders"
	// PatientrecordTable is the table the holds the patientrecord relation/edge.
	PatientrecordTable = "patientrecords"
	// PatientrecordInverseTable is the table name for the Patientrecord entity.
	// It exists in this package in order to avoid circular dependency with the "patientrecord" package.
	PatientrecordInverseTable = "patientrecords"
	// PatientrecordColumn is the table column denoting the patientrecord relation/edge.
	PatientrecordColumn = "gender_id"
)

// Columns holds all SQL columns for gender fields.
var Columns = []string{
	FieldID,
	FieldGenderstatus,
}
