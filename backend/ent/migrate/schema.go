// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AbilitypatientrightsColumns holds the columns for the "abilitypatientrights" table.
	AbilitypatientrightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "operative", Type: field.TypeInt},
		{Name: "medical_supplies", Type: field.TypeInt},
		{Name: "examine", Type: field.TypeInt},
		{Name: "check", Type: field.TypeString, Unique: true},
	}
	// AbilitypatientrightsTable holds the schema information for the "abilitypatientrights" table.
	AbilitypatientrightsTable = &schema.Table{
		Name:        "abilitypatientrights",
		Columns:     AbilitypatientrightsColumns,
		PrimaryKey:  []*schema.Column{AbilitypatientrightsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BillsColumns holds the columns for the "bills" table.
	BillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "payer", Type: field.TypeString},
		{Name: "payercontact", Type: field.TypeString, Size: 10},
		{Name: "amount", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "officer_id", Type: field.TypeInt, Nullable: true},
		{Name: "paytype_id", Type: field.TypeInt, Nullable: true},
		{Name: "treatment_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// BillsTable holds the schema information for the "bills" table.
	BillsTable = &schema.Table{
		Name:       "bills",
		Columns:    BillsColumns,
		PrimaryKey: []*schema.Column{BillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bills_financiers_EdgesOfBills",
				Columns: []*schema.Column{BillsColumns[5]},

				RefColumns: []*schema.Column{FinanciersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bills_paytypes_EdgesOfBills",
				Columns: []*schema.Column{BillsColumns[6]},

				RefColumns: []*schema.Column{PaytypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bills_unpaybills_EdgesOfBills",
				Columns: []*schema.Column{BillsColumns[7]},

				RefColumns: []*schema.Column{UnpaybillsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BloodtypesColumns holds the columns for the "bloodtypes" table.
	BloodtypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bloodtype", Type: field.TypeString, Unique: true},
	}
	// BloodtypesTable holds the schema information for the "bloodtypes" table.
	BloodtypesTable = &schema.Table{
		Name:        "bloodtypes",
		Columns:     BloodtypesColumns,
		PrimaryKey:  []*schema.Column{BloodtypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "department", Type: field.TypeString},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:        "departments",
		Columns:     DepartmentsColumns,
		PrimaryKey:  []*schema.Column{DepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DoctorsColumns holds the columns for the "doctors" table.
	DoctorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "doctorinfo_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// DoctorsTable holds the schema information for the "doctors" table.
	DoctorsTable = &schema.Table{
		Name:       "doctors",
		Columns:    DoctorsColumns,
		PrimaryKey: []*schema.Column{DoctorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "doctors_doctorinfos_EdgesOfDoctor",
				Columns: []*schema.Column{DoctorsColumns[1]},

				RefColumns: []*schema.Column{DoctorinfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "doctors_users_EdgesOfDoctor",
				Columns: []*schema.Column{DoctorsColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DoctorinfosColumns holds the columns for the "doctorinfos" table.
	DoctorinfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "doctorname", Type: field.TypeString},
		{Name: "doctorsurname", Type: field.TypeString},
		{Name: "telephonenumber", Type: field.TypeString, Size: 10},
		{Name: "licensenumber", Type: field.TypeString, Size: 11},
		{Name: "department", Type: field.TypeInt, Nullable: true},
		{Name: "level", Type: field.TypeInt, Nullable: true},
		{Name: "roomnumber", Type: field.TypeInt, Nullable: true},
		{Name: "prefix", Type: field.TypeInt, Nullable: true},
	}
	// DoctorinfosTable holds the schema information for the "doctorinfos" table.
	DoctorinfosTable = &schema.Table{
		Name:       "doctorinfos",
		Columns:    DoctorinfosColumns,
		PrimaryKey: []*schema.Column{DoctorinfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "doctorinfos_departments_EdgesOfDepartment2doctorinfo",
				Columns: []*schema.Column{DoctorinfosColumns[5]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "doctorinfos_educationlevels_EdgesOfEducationlevel2doctorinfo",
				Columns: []*schema.Column{DoctorinfosColumns[6]},

				RefColumns: []*schema.Column{EducationlevelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "doctorinfos_officerooms_EdgesOfOfficeroom2doctorinfo",
				Columns: []*schema.Column{DoctorinfosColumns[7]},

				RefColumns: []*schema.Column{OfficeroomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "doctorinfos_prenames_EdgesOfPrename2doctorinfo",
				Columns: []*schema.Column{DoctorinfosColumns[8]},

				RefColumns: []*schema.Column{PrenamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EducationlevelsColumns holds the columns for the "educationlevels" table.
	EducationlevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "level", Type: field.TypeString},
	}
	// EducationlevelsTable holds the schema information for the "educationlevels" table.
	EducationlevelsTable = &schema.Table{
		Name:        "educationlevels",
		Columns:     EducationlevelsColumns,
		PrimaryKey:  []*schema.Column{EducationlevelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FinanciersColumns holds the columns for the "financiers" table.
	FinanciersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// FinanciersTable holds the schema information for the "financiers" table.
	FinanciersTable = &schema.Table{
		Name:       "financiers",
		Columns:    FinanciersColumns,
		PrimaryKey: []*schema.Column{FinanciersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "financiers_users_EdgesOfFinancier",
				Columns: []*schema.Column{FinanciersColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "genderstatus", Type: field.TypeString, Unique: true},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// HistorytakingsColumns holds the columns for the "historytakings" table.
	HistorytakingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hight", Type: field.TypeFloat32},
		{Name: "weight", Type: field.TypeFloat32},
		{Name: "temp", Type: field.TypeFloat32},
		{Name: "pulse", Type: field.TypeInt},
		{Name: "respiration", Type: field.TypeInt},
		{Name: "bp", Type: field.TypeInt},
		{Name: "oxygen", Type: field.TypeString},
		{Name: "symptom", Type: field.TypeString},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "department_id", Type: field.TypeInt, Nullable: true},
		{Name: "nurse_id", Type: field.TypeInt, Nullable: true},
		{Name: "patientrecord_id", Type: field.TypeInt, Nullable: true},
		{Name: "symptomseverity_id", Type: field.TypeInt, Nullable: true},
	}
	// HistorytakingsTable holds the schema information for the "historytakings" table.
	HistorytakingsTable = &schema.Table{
		Name:       "historytakings",
		Columns:    HistorytakingsColumns,
		PrimaryKey: []*schema.Column{HistorytakingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "historytakings_departments_EdgesOfHistorytaking",
				Columns: []*schema.Column{HistorytakingsColumns[10]},

				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "historytakings_nurses_EdgesOfHistorytaking",
				Columns: []*schema.Column{HistorytakingsColumns[11]},

				RefColumns: []*schema.Column{NursesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "historytakings_patientrecords_EdgesOfHistorytaking",
				Columns: []*schema.Column{HistorytakingsColumns[12]},

				RefColumns: []*schema.Column{PatientrecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "historytakings_symptomseverities_EdgesOfHistorytaking",
				Columns: []*schema.Column{HistorytakingsColumns[13]},

				RefColumns: []*schema.Column{SymptomseveritiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InsurancesColumns holds the columns for the "insurances" table.
	InsurancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "insurancecompany", Type: field.TypeString, Unique: true},
	}
	// InsurancesTable holds the schema information for the "insurances" table.
	InsurancesTable = &schema.Table{
		Name:        "insurances",
		Columns:     InsurancesColumns,
		PrimaryKey:  []*schema.Column{InsurancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MedicalrecordstaffsColumns holds the columns for the "medicalrecordstaffs" table.
	MedicalrecordstaffsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// MedicalrecordstaffsTable holds the schema information for the "medicalrecordstaffs" table.
	MedicalrecordstaffsTable = &schema.Table{
		Name:       "medicalrecordstaffs",
		Columns:    MedicalrecordstaffsColumns,
		PrimaryKey: []*schema.Column{MedicalrecordstaffsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "medicalrecordstaffs_users_EdgesOfMedicalrecordstaff",
				Columns: []*schema.Column{MedicalrecordstaffsColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// NursesColumns holds the columns for the "nurses" table.
	NursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "nursinglicense", Type: field.TypeString},
		{Name: "position", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// NursesTable holds the schema information for the "nurses" table.
	NursesTable = &schema.Table{
		Name:       "nurses",
		Columns:    NursesColumns,
		PrimaryKey: []*schema.Column{NursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "nurses_users_EdgesOfNurse",
				Columns: []*schema.Column{NursesColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OfficeroomsColumns holds the columns for the "officerooms" table.
	OfficeroomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "roomnumber", Type: field.TypeString},
	}
	// OfficeroomsTable holds the schema information for the "officerooms" table.
	OfficeroomsTable = &schema.Table{
		Name:        "officerooms",
		Columns:     OfficeroomsColumns,
		PrimaryKey:  []*schema.Column{OfficeroomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PatientrecordsColumns holds the columns for the "patientrecords" table.
	PatientrecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "idcardnumber", Type: field.TypeString, Size: 13},
		{Name: "age", Type: field.TypeInt},
		{Name: "disease", Type: field.TypeString},
		{Name: "allergic", Type: field.TypeString},
		{Name: "phonenumber", Type: field.TypeString, Size: 10},
		{Name: "email", Type: field.TypeString},
		{Name: "home", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "bloodtype_id", Type: field.TypeInt, Nullable: true},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "medicalrecordstaff_id", Type: field.TypeInt, Nullable: true},
		{Name: "prefix_id", Type: field.TypeInt, Nullable: true},
	}
	// PatientrecordsTable holds the schema information for the "patientrecords" table.
	PatientrecordsTable = &schema.Table{
		Name:       "patientrecords",
		Columns:    PatientrecordsColumns,
		PrimaryKey: []*schema.Column{PatientrecordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "patientrecords_bloodtypes_EdgesOfPatientrecord",
				Columns: []*schema.Column{PatientrecordsColumns[10]},

				RefColumns: []*schema.Column{BloodtypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patientrecords_genders_EdgesOfPatientrecord",
				Columns: []*schema.Column{PatientrecordsColumns[11]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patientrecords_medicalrecordstaffs_EdgesOfPatientrecord",
				Columns: []*schema.Column{PatientrecordsColumns[12]},

				RefColumns: []*schema.Column{MedicalrecordstaffsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patientrecords_prenames_EdgesOfPatientrecord",
				Columns: []*schema.Column{PatientrecordsColumns[13]},

				RefColumns: []*schema.Column{PrenamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PatientrightsColumns holds the columns for the "patientrights" table.
	PatientrightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "permission_date", Type: field.TypeTime},
		{Name: "permission", Type: field.TypeString},
		{Name: "permission_area", Type: field.TypeString},
		{Name: "responsible", Type: field.TypeString},
		{Name: "Abilitypatientrights_id", Type: field.TypeInt, Nullable: true},
		{Name: "Insurance_id", Type: field.TypeInt, Nullable: true},
		{Name: "medicalrecordstaff_id", Type: field.TypeInt, Nullable: true},
		{Name: "patientrecord_id", Type: field.TypeInt, Nullable: true},
	}
	// PatientrightsTable holds the schema information for the "patientrights" table.
	PatientrightsTable = &schema.Table{
		Name:       "patientrights",
		Columns:    PatientrightsColumns,
		PrimaryKey: []*schema.Column{PatientrightsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "patientrights_abilitypatientrights_EdgesOfAbilitypatientrightsPatientrights",
				Columns: []*schema.Column{PatientrightsColumns[5]},

				RefColumns: []*schema.Column{AbilitypatientrightsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patientrights_insurances_EdgesOfInsurancePatientrights",
				Columns: []*schema.Column{PatientrightsColumns[6]},

				RefColumns: []*schema.Column{InsurancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patientrights_medicalrecordstaffs_EdgesOfMedicalrecordstaffPatientrights",
				Columns: []*schema.Column{PatientrightsColumns[7]},

				RefColumns: []*schema.Column{MedicalrecordstaffsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "patientrights_patientrecords_EdgesOfPatientrecordPatientrights",
				Columns: []*schema.Column{PatientrightsColumns[8]},

				RefColumns: []*schema.Column{PatientrecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaytypesColumns holds the columns for the "paytypes" table.
	PaytypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "paytype", Type: field.TypeString, Unique: true},
	}
	// PaytypesTable holds the schema information for the "paytypes" table.
	PaytypesTable = &schema.Table{
		Name:        "paytypes",
		Columns:     PaytypesColumns,
		PrimaryKey:  []*schema.Column{PaytypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PrenamesColumns holds the columns for the "prenames" table.
	PrenamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "prefix", Type: field.TypeString, Unique: true},
	}
	// PrenamesTable holds the schema information for the "prenames" table.
	PrenamesTable = &schema.Table{
		Name:        "prenames",
		Columns:     PrenamesColumns,
		PrimaryKey:  []*schema.Column{PrenamesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RegistrarsColumns holds the columns for the "registrars" table.
	RegistrarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "user_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// RegistrarsTable holds the schema information for the "registrars" table.
	RegistrarsTable = &schema.Table{
		Name:       "registrars",
		Columns:    RegistrarsColumns,
		PrimaryKey: []*schema.Column{RegistrarsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "registrars_users_EdgesOfUser2registrar",
				Columns: []*schema.Column{RegistrarsColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SymptomseveritiesColumns holds the columns for the "symptomseverities" table.
	SymptomseveritiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symptomseverity", Type: field.TypeString},
	}
	// SymptomseveritiesTable holds the schema information for the "symptomseverities" table.
	SymptomseveritiesTable = &schema.Table{
		Name:        "symptomseverities",
		Columns:     SymptomseveritiesColumns,
		PrimaryKey:  []*schema.Column{SymptomseveritiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TreatmentsColumns holds the columns for the "treatments" table.
	TreatmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "symptom", Type: field.TypeString},
		{Name: "treat", Type: field.TypeString},
		{Name: "medicine", Type: field.TypeString},
		{Name: "datetreat", Type: field.TypeTime},
		{Name: "doctor_id", Type: field.TypeInt, Nullable: true},
		{Name: "patientrecord_id", Type: field.TypeInt, Nullable: true},
		{Name: "typetreatment_id", Type: field.TypeInt, Nullable: true},
	}
	// TreatmentsTable holds the schema information for the "treatments" table.
	TreatmentsTable = &schema.Table{
		Name:       "treatments",
		Columns:    TreatmentsColumns,
		PrimaryKey: []*schema.Column{TreatmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "treatments_doctors_EdgesOfTreatment",
				Columns: []*schema.Column{TreatmentsColumns[5]},

				RefColumns: []*schema.Column{DoctorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "treatments_patientrecords_EdgesOfTreatment",
				Columns: []*schema.Column{TreatmentsColumns[6]},

				RefColumns: []*schema.Column{PatientrecordsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "treatments_typetreatments_EdgesOfTreatment",
				Columns: []*schema.Column{TreatmentsColumns[7]},

				RefColumns: []*schema.Column{TypetreatmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TypetreatmentsColumns holds the columns for the "typetreatments" table.
	TypetreatmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "typetreatment", Type: field.TypeString},
	}
	// TypetreatmentsTable holds the schema information for the "typetreatments" table.
	TypetreatmentsTable = &schema.Table{
		Name:        "typetreatments",
		Columns:     TypetreatmentsColumns,
		PrimaryKey:  []*schema.Column{TypetreatmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UnpaybillsColumns holds the columns for the "unpaybills" table.
	UnpaybillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString},
		{Name: "treatment_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UnpaybillsTable holds the schema information for the "unpaybills" table.
	UnpaybillsTable = &schema.Table{
		Name:       "unpaybills",
		Columns:    UnpaybillsColumns,
		PrimaryKey: []*schema.Column{UnpaybillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "unpaybills_treatments_EdgesOfUnpaybills",
				Columns: []*schema.Column{UnpaybillsColumns[2]},

				RefColumns: []*schema.Column{TreatmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "images", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
		{Name: "userstatus_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_patientrights_EdgesOfUserPatientrights",
				Columns: []*schema.Column{UsersColumns[4]},

				RefColumns: []*schema.Column{PatientrightsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_userstatuses_EdgesOfUser",
				Columns: []*schema.Column{UsersColumns[5]},

				RefColumns: []*schema.Column{UserstatusesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserstatusesColumns holds the columns for the "userstatuses" table.
	UserstatusesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "userstatus", Type: field.TypeString, Unique: true},
	}
	// UserstatusesTable holds the schema information for the "userstatuses" table.
	UserstatusesTable = &schema.Table{
		Name:        "userstatuses",
		Columns:     UserstatusesColumns,
		PrimaryKey:  []*schema.Column{UserstatusesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AbilitypatientrightsTable,
		BillsTable,
		BloodtypesTable,
		DepartmentsTable,
		DoctorsTable,
		DoctorinfosTable,
		EducationlevelsTable,
		FinanciersTable,
		GendersTable,
		HistorytakingsTable,
		InsurancesTable,
		MedicalrecordstaffsTable,
		NursesTable,
		OfficeroomsTable,
		PatientrecordsTable,
		PatientrightsTable,
		PaytypesTable,
		PrenamesTable,
		RegistrarsTable,
		SymptomseveritiesTable,
		TreatmentsTable,
		TypetreatmentsTable,
		UnpaybillsTable,
		UsersTable,
		UserstatusesTable,
	}
)

func init() {
	BillsTable.ForeignKeys[0].RefTable = FinanciersTable
	BillsTable.ForeignKeys[1].RefTable = PaytypesTable
	BillsTable.ForeignKeys[2].RefTable = UnpaybillsTable
	DoctorsTable.ForeignKeys[0].RefTable = DoctorinfosTable
	DoctorsTable.ForeignKeys[1].RefTable = UsersTable
	DoctorinfosTable.ForeignKeys[0].RefTable = DepartmentsTable
	DoctorinfosTable.ForeignKeys[1].RefTable = EducationlevelsTable
	DoctorinfosTable.ForeignKeys[2].RefTable = OfficeroomsTable
	DoctorinfosTable.ForeignKeys[3].RefTable = PrenamesTable
	FinanciersTable.ForeignKeys[0].RefTable = UsersTable
	HistorytakingsTable.ForeignKeys[0].RefTable = DepartmentsTable
	HistorytakingsTable.ForeignKeys[1].RefTable = NursesTable
	HistorytakingsTable.ForeignKeys[2].RefTable = PatientrecordsTable
	HistorytakingsTable.ForeignKeys[3].RefTable = SymptomseveritiesTable
	MedicalrecordstaffsTable.ForeignKeys[0].RefTable = UsersTable
	NursesTable.ForeignKeys[0].RefTable = UsersTable
	PatientrecordsTable.ForeignKeys[0].RefTable = BloodtypesTable
	PatientrecordsTable.ForeignKeys[1].RefTable = GendersTable
	PatientrecordsTable.ForeignKeys[2].RefTable = MedicalrecordstaffsTable
	PatientrecordsTable.ForeignKeys[3].RefTable = PrenamesTable
	PatientrightsTable.ForeignKeys[0].RefTable = AbilitypatientrightsTable
	PatientrightsTable.ForeignKeys[1].RefTable = InsurancesTable
	PatientrightsTable.ForeignKeys[2].RefTable = MedicalrecordstaffsTable
	PatientrightsTable.ForeignKeys[3].RefTable = PatientrecordsTable
	RegistrarsTable.ForeignKeys[0].RefTable = UsersTable
	TreatmentsTable.ForeignKeys[0].RefTable = DoctorsTable
	TreatmentsTable.ForeignKeys[1].RefTable = PatientrecordsTable
	TreatmentsTable.ForeignKeys[2].RefTable = TypetreatmentsTable
	UnpaybillsTable.ForeignKeys[0].RefTable = TreatmentsTable
	UsersTable.ForeignKeys[0].RefTable = PatientrightsTable
	UsersTable.ForeignKeys[1].RefTable = UserstatusesTable
}
