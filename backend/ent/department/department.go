// Code generated by entc, DO NOT EDIT.

package department

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDepartment holds the string denoting the department field in the database.
	FieldDepartment = "department"

	// EdgeDepartment2doctorinfo holds the string denoting the department2doctorinfo edge name in mutations.
	EdgeDepartment2doctorinfo = "department2doctorinfo"
	// EdgeHistorytaking holds the string denoting the historytaking edge name in mutations.
	EdgeHistorytaking = "historytaking"

	// Table holds the table name of the department in the database.
	Table = "departments"
	// Department2doctorinfoTable is the table the holds the department2doctorinfo relation/edge.
	Department2doctorinfoTable = "doctorinfos"
	// Department2doctorinfoInverseTable is the table name for the Doctorinfo entity.
	// It exists in this package in order to avoid circular dependency with the "doctorinfo" package.
	Department2doctorinfoInverseTable = "doctorinfos"
	// Department2doctorinfoColumn is the table column denoting the department2doctorinfo relation/edge.
	Department2doctorinfoColumn = "department"
	// HistorytakingTable is the table the holds the historytaking relation/edge.
	HistorytakingTable = "historytakings"
	// HistorytakingInverseTable is the table name for the Historytaking entity.
	// It exists in this package in order to avoid circular dependency with the "historytaking" package.
	HistorytakingInverseTable = "historytakings"
	// HistorytakingColumn is the table column denoting the historytaking relation/edge.
	HistorytakingColumn = "department_id"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldDepartment,
}

var (
	// DepartmentValidator is a validator for the "department" field. It is called by the builders before save.
	DepartmentValidator func(string) error
)