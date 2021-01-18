// Code generated by entc, DO NOT EDIT.

package nurse

const (
	// Label holds the string label denoting the nurse type in the database.
	Label = "nurse"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNursinglicense holds the string denoting the nursinglicense field in the database.
	FieldNursinglicense = "nursinglicense"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"

	// EdgeEdgesOfHistorytaking holds the string denoting the edgesofhistorytaking edge name in mutations.
	EdgeEdgesOfHistorytaking = "EdgesOfHistorytaking"
	// EdgeEdgesOfUser holds the string denoting the edgesofuser edge name in mutations.
	EdgeEdgesOfUser = "EdgesOfUser"

	// Table holds the table name of the nurse in the database.
	Table = "nurses"
	// EdgesOfHistorytakingTable is the table the holds the EdgesOfHistorytaking relation/edge.
	EdgesOfHistorytakingTable = "historytakings"
	// EdgesOfHistorytakingInverseTable is the table name for the Historytaking entity.
	// It exists in this package in order to avoid circular dependency with the "historytaking" package.
	EdgesOfHistorytakingInverseTable = "historytakings"
	// EdgesOfHistorytakingColumn is the table column denoting the EdgesOfHistorytaking relation/edge.
	EdgesOfHistorytakingColumn = "nurse_id"
	// EdgesOfUserTable is the table the holds the EdgesOfUser relation/edge.
	EdgesOfUserTable = "nurses"
	// EdgesOfUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	EdgesOfUserInverseTable = "users"
	// EdgesOfUserColumn is the table column denoting the EdgesOfUser relation/edge.
	EdgesOfUserColumn = "user_id"
)

// Columns holds all SQL columns for nurse fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNursinglicense,
	FieldPosition,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Nurse type.
var ForeignKeys = []string{
	"user_id",
}
