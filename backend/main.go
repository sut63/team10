package main

import (
	"context"
	"fmt"
	"log"
	

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team10/app/controllers"
	_ "github.com/team10/app/docs"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/abilitypatientrights"
	"github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/user"

	//import by patientrights No.3
	//vvv...............................vvv
	/*
	"time"
	"github.com/team10/app/ent/insurance"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/patientrightstype"
	*/
	//^^^...............................^^^
	//import by doctorinformation No.6
	//vvv...............................vvv
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/educationlevel"
	"github.com/team10/app/ent/officeroom"
	"github.com/team10/app/ent/prename"
	//^^^...............................^^^
)

// struct By team 10
//-------------------------------------------------------------------
// Struct By Doctorinformation System No.6
//vvv*******************************************************************vvv

// Officerooms defines the struct for the Officerooms
type Officerooms struct {
	Officeroom []Officeroom
}

// Officeroom defines the struct for the Officeroom
type Officeroom struct {
	Roomnumber int
}

// Educationlevels defines the struct for the Educationlevels
type Educationlevels struct {
	Educationlevel []Educationlevel
}

// Educationlevel defines the struct for the Educationlevel
type Educationlevel struct {
	Level string
}

// Prenames defines the struct for the Prenames
type Prenames struct {
	Prename []Prename
}

// Prename defines the struct for the Prename
type Prename struct {
	Prefix string
}

// Departments defines the struct for the Departments
type Departments struct {
	Department []Department
}

// Department defines the struct for the Department
type Department struct {
	Department string
}

//^^^*******************************************************************^^^

// Struct By Patientrights System No.3
//*******************************************************************

// Patientrightstypes defines the struct for the Patientrightstypes
type Patientrightstypes struct {
	Patientrightstype []Patientrightstype
}

// Patientrightstype defines the struct for the Patientrightstype
type Patientrightstype struct {
	Permission           string
	PermissionArea       string
	Responsible          string
	Abilitypatientrights int
}

// Abilitypatientrightss defines the struct for the Abilitypatientrightss
type Abilitypatientrightss struct {
	Abilitypatientrights []Abilitypatientrights
}

// Abilitypatientrights defines the struct for the Abilitypatientrights
type Abilitypatientrights struct {
	Operative       int
	MedicalSupplies int
	Examine         int
}

// Patientrightss defines the struct for the Patientrightss
type Patientrightss struct {
	Patientrights []Patientrights
}

// Patientrights defines the struct for the Patientrights
type Patientrights struct {
	Patientrightstype  int
	Patientrecord      int
	Insurance          int
	Medicalrecordstaff int
}

// Struct By Patientrecord System  No.2

//*******************************************************************

// Genders defines the struct for the Genders
type Genders struct {
	Gender []Gender
}

// Gender defines the struct for the Gender
type Gender struct {
	Genderstatus string
}

// Medicalrecordstaffs defines the struct for the Medicalrecordstaffs
type Medicalrecordstaffs struct {
	Medicalrecordstaff []Medicalrecordstaff
}

// Medicalrecordstaff defines the struct for the Medicalrecordstaff
type Medicalrecordstaff struct {
	Name string
}

//*******************************************************************
// Users defines the struct for the Users

// Paytypes defines the struct for the Paytypes
type Paytypes struct {
	Paytype []Paytype
}

// Paytype defines the struct for the Paytype
type Paytype struct {
	paytype string
}

// Financiers defines the struct for the Financiers
type Financiers struct {
	Financier []Financier
}

// Financier defines the struct for the Financier
type Financier struct {
	name string
	user int
}

//*******************************************************************

// Struct By Historytaking System  No.5
//*******************************************************************

// Nurses defines the struct for the Nurses
type Nurses struct {
	Nurse []Nurse
}

// Nurse defines the struct for the Nurse
type Nurse struct {
	Name           string
	Nursinglicense string
	Position       string
	User           int
}

// Symptomseveritys defines the struct for the Symptomseveritys
type Symptomseveritys struct {
	Symptomseverity []Symptomseverity
}

// Symptomseverity defines the struct for the Symptomseverity
type Symptomseverity struct {
	Symptomseverity string
}

//^^^*******************************************************************^^^

//^^^^^^^^^-------------------------------------------------------------------^^^^^^^^^

//vvv::::::::::::::::::::::::::::::::::::::::::::::::vvv

// Doctorinfos defines the struct for the Doctorinfos
type Doctorinfos struct {
	Doctorinfo []Doctorinfo
}

// Doctorinfo defines the struct for the Doctorinfo
type Doctorinfo struct {
	Doctorname      string
	Doctorsurname   string
	Telephonenumber string
	Licensenumber   string
	Department      int
	Educationlevel  int
	Officeroom      int
	Prename         int
	User            int
	Registrar       int
}

// Struct By Treatment System
//*******************************************************************

// Treatments defines the struct for the Treatments
type Treatments struct {
	Treatment []Treatment
}

// Treatment defines the struct for the Treatment
type Treatment struct {
	Treatment string
	Datetreat string
}

// Typetreatments defines the struct for the  Typetreatments
type Typetreatments struct {
	Typetreatment [] Typetreatment
}

//  Typetreatment defines the struct for the  Typetreatment
type  Typetreatment struct {
	Typetreatment string
}

//^^^::::::::::::::::::::::::::::::::::::::::::::::::^^^

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {

	// Set router By Team10
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")

	// Controller By Team 10 System
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv
	controllers.NewUserController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	// Controller By Patientrights System  No.3
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv
	controllers.NewPatientrightsController(v1, client)
	controllers.NewPatientrightstypeController(v1, client)
	controllers.NewAbilitypatientrightsController(v1, client)
	controllers.NewInsuranceController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	// Controller By Patientrecord System No.2
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv
	controllers.NewGenderController(v1, client)
	controllers.NewMedicalrecordstaffController(v1, client)
	controllers.NewPatientrecordController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	//Controller By Bill System No.1
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv
	controllers.NewBillController(v1, client)
	controllers.NewFinancierController(v1, client)
	controllers.NewPaytypeController(v1, client)
	controllers.NewUnpaybillController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	//Controller By Historytaking System No.5
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv
	controllers.NewHistorytakingController(v1, client)
	controllers.NewNurseController(v1, client)
	controllers.NewSymptomseverityController(v1, client)
	controllers.NewDepartmentController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	//Controller By Doctorinformation System No.6
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv
	controllers.NewEducationlevelController(v1, client)
	controllers.NewDoctorinfoController(v1, client)
	controllers.NewOfficeroomController(v1, client)
	controllers.NewPrenameController(v1, client)
	controllers.NewRegistrarController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	// Set Postman By Team10
	//vvvvvvvvv-------------------------------------------------------------------vvvvvvvvv
	// Set Postman By Team10 System
	//vvv*******************************************************************vvv
	User := []string{"Khatadet_khianchainat", "nara_haru", "morani_rode", "faratell_yova", "pulla_visan", "omaha_yad",}
	for _, r := range User {
		client.User.
			Create().
			SetEmail(r + "@gmail.com").
			SetPassword("123456").
			Save(context.Background())
	}
	//##############################################################################################
	Registrar := []string{"Khatadet_khianchainat", "nara_haru", "morani_rode", "faratell_yova", "pulla_visan", "omaha_yad"}
	for _, reg := range Registrar {
		client.Registrar.
			Create().
			SetName(reg).
			Save(context.Background())
	}
	//^^^*******************************************************************^^^

	// Set Postman By Doctorinfomation System No.6
	//vvv*******************************************************************vvv
	// Set Educationlevel Data
	Educationlevel := []string{"Bachelor's degree", "Master's degree", "Doctoral Degree"}

	for _, el := range Educationlevel {
		client.Educationlevel.
			Create().
			SetLevel(el).
			Save(context.Background())
	}
	// Set Officeroom Data
	Officeroom := []string{"B10", "B11", "B12", "B13", "B14"}

	for _, or := range Officeroom {
		client.Officeroom.
			Create().
			SetRoomnumber(or).
			Save(context.Background())
	}

	// Set Prename Data
	Prename := []string{"นายเเพทย์", "เเพทย์หญิง", "นาย", "นาง", "นางสาว", "เด็กชาย", "เด็กหญิง"}

	for _, pn := range Prename {
		client.Prename.
			Create().
			SetPrefix(pn).
			Save(context.Background())
	}
	//^^^*******************************************************************^^^

	// Set Postman By Patientrights System No.3
	//vvv*******************************************************************vvv

	// Set insurance Data
	Insurance := []string{"เมืองไทยประกันภัย", "ไทยสมุทรประกันชีวิต", "อื่น ๆ ", "กรมบัญชีกลาง", "AIA"}
	for _, r := range Insurance {
		client.Insurance.
			Create().
			SetInsurancecompany(r).
			Save(context.Background())
	}

	// Set Abilitypatientrights Data
	Abilitypatientrights := Abilitypatientrightss{
		Abilitypatientrights: []Abilitypatientrights{
			Abilitypatientrights{100, 100, 100},
			Abilitypatientrights{50, 100, 100},
			Abilitypatientrights{50, 100, 50},
		},
	}

	for _, a := range Abilitypatientrights.Abilitypatientrights {
		client.Abilitypatientrights.
			Create().
			SetOperative(a.Operative).
			SetMedicalSupplies(a.MedicalSupplies).
			SetExamine(a.Examine).
			Save(context.Background())
	}

	// Set Patientrightstypes Data
	patientrightstypes := Patientrightstypes{
		Patientrightstype: []Patientrightstype{
			Patientrightstype{"จ่ายตรง", "ทั่วประเทศ", "ราชการ", 1},
			Patientrightstype{"ผู้สูงอายุ", "ทั่วประเทศ", "นาย แหวง", 2},
			Patientrightstype{"สุขภาพถ้วนหน้า", "โคราช", "ราชการ", 2},
			Patientrightstype{"อุบัติเหตุสบายใจ", "ทั่วประเทศ", "นาย มา", 2},
		},
	}

	for _, p := range patientrightstypes.Patientrightstype {

		a, err := client.Abilitypatientrights.
			Query().
			Where(abilitypatientrights.IDEQ(int(p.Abilitypatientrights))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Patientrightstype.
			Create().
			SetPermission(p.Permission).
			SetPermissionArea(p.PermissionArea).
			SetResponsible(p.Responsible).
			SetPatientrightstypeAbilitypatientrights(a).
			Save(context.Background())
	}
	//^^^*******************************************************************^^^

	// Set Postman By Patientrecord System No.2
	//vvv*******************************************************************vvv

	// Set Gender Data
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGenderstatus(g.Genderstatus).
			Save(context.Background())
	}

	// Set Medicalrecordstaff Data
	medicalrecordstaffs := Medicalrecordstaffs{
		Medicalrecordstaff: []Medicalrecordstaff{
			Medicalrecordstaff{"Phonrawin Kudthalaeng"},
			Medicalrecordstaff{"Shin Sura"},
		},
	}

	for _, m := range medicalrecordstaffs.Medicalrecordstaff {
		client.Medicalrecordstaff.
			Create().
			SetName(m.Name).
			Save(context.Background())
	}
	//^^^*******************************************************************^^^
	// Set Postman By Historytaking System No.5
	//vvv*******************************************************************vvv

	//Set nurse data
	nurses := Nurses{
		Nurse: []Nurse{
			Nurse{"Paonrat Panjainam", "Nurse123456", "พยาบาลวิชาชีพ", 2},
			Nurse{"Name Surname", "Nurse001122", "พยาบาลวิชาชีพ", 3},
		},
	}

	for _, n := range nurses.Nurse {
		u, err := client.User.
			Query().
			Where(user.IDEQ(int(n.User))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Nurse.
			Create().
			SetUser(u).
			SetName(n.Name).
			SetNursinglicense(n.Nursinglicense).
			SetPosition(n.Position).
			Save(context.Background())
	}

	//Set Symptomseverity data
	symptomseveritys := Symptomseveritys{
		Symptomseverity: []Symptomseverity{
			Symptomseverity{"ฉุกเฉิน"},
			Symptomseverity{"ฉุกเฉินรอได้"},
			Symptomseverity{"ปานกลาง"},
		},
	}

	for _, ss := range symptomseveritys.Symptomseverity {
		client.Symptomseverity.
			Create().
			SetSymptomseverity(ss.Symptomseverity).
			Save(context.Background())
	}
	//Set Department data
	departments := Departments{
		Department: []Department{
			Department{"ตาคอหู"},
			Department{"กระดูก"},
			Department{"อายุรกรรม"},
		},
	}

	for _, d := range departments.Department {
		client.Department.
			Create().
			SetDepartment(d.Department).
			Save(context.Background())
	}

	//^^^*******************************************************************^^^
	// Set Postman By Bill System No.1
	//vvv*******************************************************************vvv

	//Set Financier data
	financiers := Financiers{
		Financier: []Financier{
			Financier{"Nutchaporn Klinrod",4},
			Financier{"Name Surname",5},
		},
	}
	for _, f := range financiers.Financier {
		u, err := client.User.
			Query().
			Where(user.IDEQ(f.user)).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.Financier.
			Create().
			SetName(f.name).
			SetUser(u).
			Save(context.Background())
	}


	//Set Paytype data
	paytypes := Paytypes{
		Paytype: []Paytype{
			Paytype{"Online Banking"},
			Paytype{"Credit / Debit Card"},
			Paytype{"Cash"},
		},
	}

	for _, pt := range paytypes.Paytype {
		client.Paytype.
			Create().
			SetPaytype(pt.paytype).
			Save(context.Background())
	}
	//^^^*******************************************************************^^^
	//^^^^^^^^^-------------------------------------------------------------------^^^^^^^^^

	//Set Postman Output
	//vvv*******************************************************************vvv

	

	// Set Doctorinformation output No.6
	//vvv...................................................................vvv
	Doctorinfos := Doctorinfos{
		Doctorinfo: []Doctorinfo{
			Doctorinfo{"Thanawat", "Srikaewsiew", "0800740864", "BEN10UT100", 1, 1, 1, 1, 1, 1},
			Doctorinfo{"Paonrat", "Tangtong", "0810740864", "BEN20UT100", 1, 1, 1, 1, 1, 1},
		},
	}
	for _, doc := range Doctorinfos.Doctorinfo {

		dp, err := client.Department.
			Query().
			Where(department.IDEQ(int(doc.Department))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		el, err := client.Educationlevel.
			Query().
			Where(educationlevel.IDEQ(int(doc.Educationlevel))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		or, err := client.Officeroom.
			Query().
			Where(officeroom.IDEQ(int(doc.Officeroom))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		pn, err := client.Prename.
			Query().
			Where(prename.IDEQ(int(doc.Prename))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		us, err := client.User.
			Query().
			Where(user.IDEQ(int(doc.User))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		rg, err := client.Registrar.
			Query().
			Where(registrar.IDEQ(int(doc.Registrar))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Doctorinfo.
			Create().
			SetDoctorname(doc.Doctorname).
			SetDoctorsurname(doc.Doctorsurname).
			SetTelephonenumber(doc.Telephonenumber).
			SetLicensenumber(doc.Licensenumber).
			SetDepartment(dp).
			SetEducationlevel(el).
			SetOfficeroom(or).
			SetPrename(pn).
			SetUser(us).
			SetRegistrar(rg).
			Save(context.Background())
	}

	//^^^...................................................................^^^

	// Set Patientrightstypes output No.3
	//vvv...................................................................vvv
/*
	patientrightss := Patientrightss{
		Patientrights: []Patientrights{
			Patientrights{1, 1, 1, 1},
			Patientrights{1, 1, 1, 1},
			Patientrights{1, 1, 1, 1},
			Patientrights{1, 1, 1, 1},
		},
	}

	for _, p := range patientrightss.Patientrights {

		Patientrightstype, err := client.Patientrightstype.
			Query().
			Where(patientrightstype.IDEQ(int(p.Patientrightstype))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		Insurance, err := client.Insurance.
			Query().
			Where(insurance.IDEQ(int(p.Insurance))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		Patientrecord, err := client.Patientrecord.
			Query().
			Where(patientrecord.IDEQ(int(p.Patientrecord))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		Medicalrecordstaff, err := client.Medicalrecordstaff.
			Query().
			Where(medicalrecordstaff.IDEQ(int(p.Medicalrecordstaff))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t := time.Now().Local()
		client.Patientrights.
			Create().
			SetPermissionDate(t).
			SetPatientrightsPatientrightstype(Patientrightstype).
			SetPatientrightsPatientrecord(Patientrecord).
			SetPatientrightsMedicalrecordstaff(Medicalrecordstaff).
			SetPatientrightsInsurance(Insurance).
			Save(context.Background())

	}
	*/
	//^^^...................................................................^^^
	
	// Set Postman By Treatment System
	//vvv*******************************************************************vvv

	//Set Typetreatment data
	typetreatments := Typetreatments{
		Typetreatment: []Typetreatment{
			Typetreatment{"ตรวจผ่าตัด"},
			Typetreatment{"ตรวจทั่วไป"},
			Typetreatment{"ตรวจวินิจฉัย"},
		},
	}
	for _, tm := range typetreatment.Typetreatment {
		client.Typetreatment.
			Create().
			SetTypetreatment(tm.Typetreatment).
			Save(context.Background())
	}

	//^^^*******************************************************************^^^

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
