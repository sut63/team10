// Code generated by entc, DO NOT EDIT.

package doctor

const (
	// Label holds the string label denoting the doctor type in the database.
	Label = "doctor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// EdgeEdgesOfDoctorinfo holds the string denoting the edgesofdoctorinfo edge name in mutations.
	EdgeEdgesOfDoctorinfo = "EdgesOfDoctorinfo"
	// EdgeEdgesOfUser holds the string denoting the edgesofuser edge name in mutations.
	EdgeEdgesOfUser = "EdgesOfUser"
	// EdgeEdgesOfTreatment holds the string denoting the edgesoftreatment edge name in mutations.
	EdgeEdgesOfTreatment = "EdgesOfTreatment"

	// Table holds the table name of the doctor in the database.
	Table = "doctors"
	// EdgesOfDoctorinfoTable is the table the holds the EdgesOfDoctorinfo relation/edge.
	EdgesOfDoctorinfoTable = "doctors"
	// EdgesOfDoctorinfoInverseTable is the table name for the Doctorinfo entity.
	// It exists in this package in order to avoid circular dependency with the "doctorinfo" package.
	EdgesOfDoctorinfoInverseTable = "doctorinfos"
	// EdgesOfDoctorinfoColumn is the table column denoting the EdgesOfDoctorinfo relation/edge.
	EdgesOfDoctorinfoColumn = "doctorinfo_id"
	// EdgesOfUserTable is the table the holds the EdgesOfUser relation/edge.
	EdgesOfUserTable = "doctors"
	// EdgesOfUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	EdgesOfUserInverseTable = "users"
	// EdgesOfUserColumn is the table column denoting the EdgesOfUser relation/edge.
	EdgesOfUserColumn = "user_id"
	// EdgesOfTreatmentTable is the table the holds the EdgesOfTreatment relation/edge.
	EdgesOfTreatmentTable = "treatments"
	// EdgesOfTreatmentInverseTable is the table name for the Treatment entity.
	// It exists in this package in order to avoid circular dependency with the "treatment" package.
	EdgesOfTreatmentInverseTable = "treatments"
	// EdgesOfTreatmentColumn is the table column denoting the EdgesOfTreatment relation/edge.
	EdgesOfTreatmentColumn = "doctor_id"
)

// Columns holds all SQL columns for doctor fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Doctor type.
var ForeignKeys = []string{
	"doctorinfo_id",
	"user_id",
}
