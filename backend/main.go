package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team10/app/controllers"
	_ "github.com/team10/app/docs"
	"github.com/team10/app/ent"

	"github.com/team10/app/ent/bloodtype"
	"github.com/team10/app/ent/doctor"
	"github.com/team10/app/ent/gender"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	_ "github.com/team10/app/ent/registrar"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/typetreatment"
	"github.com/team10/app/ent/user"
	"github.com/team10/app/ent/userstatus"

	//import by patientrights No.3
	//vvv...............................vvv
	"time"
	/*
		"github.com/team10/app/ent/insurance"
		"github.com/team10/app/ent/medicalrecordstaff"
		"github.com/team10/app/ent/patientrecord"
		"github.com/team10/app/ent/patientrightstype"
	*/
	//^^^...............................^^^
	//import by doctorinformation No.6
	//vvv...............................vvv
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/educationlevel"
	"github.com/team10/app/ent/officeroom"
	"github.com/team10/app/ent/prename"
	//^^^...............................^^^
	"github.com/team10/app/ent/unpaybill"
	"github.com/team10/app/ent/paytype"
	"github.com/team10/app/ent/financier"

)

// struct By team 10
//-------------------------------------------------------------------

// Struct By team10 System
//vvv*******************************************************************vvv

// Users defines the struct for the Users
type Users struct {
	User []User
}

// User defines the struct for the User
type User struct {
	Userstatus int
	Email      string
	Password   string
	Images     string
}

//^^^*******************************************************************^^^

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
	User int
}

// Bloodtypes defines the struct for the Bloodtypes
type Bloodtypes struct {
	Bloodtype []Bloodtype
}

// Bloodtype defines the struct for the Bloodtype
type Bloodtype struct {
	bloodtype string
}

// Patientrecords defines the struct for the Patientrecords
type Patientrecords struct {
	Patientrecord []Patientrecord
}

// Patientrecord defines the struct for the Patientrecord
type Patientrecord struct {
	Prename            int
	Name               string
	Gender             int
	Idcardnumber       string
	Age                int
	Bloodtype          int
	Disease            string
	Allergic           string
	Phonenumber        string
	Email              string
	Home               string
	Medicalrecordstaff int
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

// Unpaybills defines the struct for the Unpaybills
type Unpaybills struct {
	Unpaybill []Unpaybill
}

// Unpaybill defines the struct for the Unpaybill
type Unpaybill struct {
	status    string
	treatment int
}

// Bills defines the struct for theBills
type Bills struct {
	Bill []Bill
}

// Bill defines the struct for the Bill
type Bill struct {
	Amount    int
	Payer    string
	Payercontact    string
	Unpaybill int
	Paytype int
	Financier int
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
}

// Struct By Treatment System
//*******************************************************************

// Treatments defines the struct for the Treatments
type Treatments struct {
	Treatment []Treatment
}

// Treatment defines the struct for the Treatment
type Treatment struct {
	Symptom       string
	Treat         string
	Medicine      string
	Typetreatment int
	Doctor        int
	Patientrecord int
}

// Typetreatments defines the struct for the  Typetreatments
type Typetreatments struct {
	Typetreatment []Typetreatment
}

// Typetreatment defines the struct for the  Typetreatment
type Typetreatment struct {
	Typetreatment string
}

// Doctors defines the struct for the Doctors
type Doctors struct {
	Doctor []Doctor
}

// Doctor defines the struct for the Doctor
type Doctor struct {
	Doctorinfo int
	User       int
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

	controllers.NewAbilitypatientrightsController(v1, client)
	controllers.NewInsuranceController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	// Controller By Patientrecord System No.2
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv
	controllers.NewGenderController(v1, client)
	controllers.NewBloodtypeController(v1, client)
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
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	//Controller By Treatment System No.4
	//vvv+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++vvv

	controllers.NewTreatmentController(v1, client)
	controllers.NewTypetreatmentController(v1, client)
	controllers.NewDoctorController(v1, client)
	//^^^+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++^^^

	// Set Postman By Team10
	//vvvvvvvvv-------------------------------------------------------------------vvvvvvvvv
	// Set Postman By Team10 System
	//vvv*******************************************************************vvv
	// Set Userstatus Data
	Userstatus := []string{"Root", "Fin", "Med", "Doc", "Nur", "Reg"}
	for _, reg := range Userstatus {
		client.Userstatus.
			Create().
			SetUserstatus(reg).
			Save(context.Background())
	}
	//##############################################################################################
	// Set User Data
	User := Users{
		User: []User{
			User{1, "B0", "1234", "Images/images1.txt"},
			User{2, "B1", "1234", "Images/images2.txt"},
			User{3, "B2", "1234", "Images/images3.txt"},
			User{3, "B3", "1234", "Images/images4.txt"},
			User{4, "B4", "1234", "Images/images5.txt"},
			User{5, "B5", "1234", "Images/images6.txt"},
			User{6, "B6", "1234", "Images/images7.txt"},
			User{2, "nara_haru", "1234", "Images/images8.txt"},
			User{3, "morani_rode", "1234", "Images/images9.txt"},
			User{3, "faratell_yova", "1234", "Images/images10.txt"},
			User{3, "pulla_visan", "1234", "Images/images11.txt"},
		},
	}

	for _, r := range User.User {
		us, err := client.Userstatus.
			Query().
			Where(userstatus.IDEQ(int(r.Userstatus))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		img, err := ioutil.ReadFile(r.Images)

		if err != nil {
			fmt.Println(err.Error())
		}

		client.User.
			Create().
			SetEdgesOfUserstatus(us).
			SetEmail(r.Email + "@gmail.com").
			SetPassword(r.Password).
			SetImages(string(img)).
			Save(context.Background())
	}
	//##############################################################################################
	// Set Registrar Data
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
		o := fmt.Sprintf("%d", int(a.Operative))
		m := fmt.Sprintf("%d", int(a.MedicalSupplies))
		e := fmt.Sprintf("%d", int(a.Examine))
		var ck string
		ck = o + "-" + m + "-" + e
		client.Abilitypatientrights.
			Create().
			SetOperative(a.Operative).
			SetMedicalSupplies(a.MedicalSupplies).
			SetExamine(a.Examine).
			SetCheck(ck).
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

	// Set Bloodtype Data
	bloodtypes := Bloodtypes{
		Bloodtype: []Bloodtype{
			Bloodtype{"A"},
			Bloodtype{"B"},
			Bloodtype{"O"},
			Bloodtype{"AB"},
		},
	}

	for _, bt := range bloodtypes.Bloodtype {
		client.Bloodtype.
			Create().
			SetBloodtype(bt.bloodtype).
			Save(context.Background())
	}

	// Set Medicalrecordstaff Data
	medicalrecordstaffs := Medicalrecordstaffs{
		Medicalrecordstaff: []Medicalrecordstaff{
			Medicalrecordstaff{"Phonrawin Kudthalaeng", 3},
			Medicalrecordstaff{"Mihayo Shiho", 4},
			Medicalrecordstaff{"Josa Watashi", 9},
			Medicalrecordstaff{"Miyato Karame", 10},
			Medicalrecordstaff{"Watashi Kai", 11},
			Medicalrecordstaff{"Name Surname", 1},
		},
	}

	for _, m := range medicalrecordstaffs.Medicalrecordstaff {
		u, err := client.User.
			Query().
			Where(user.IDEQ(int(m.User))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.Medicalrecordstaff.
			Create().
			SetName(m.Name).
			SetEdgesOfUser(u).
			Save(context.Background())
	}

	// Set Patientrecord Data
	Patientrecords := Patientrecords{
		Patientrecord: []Patientrecord{
			Patientrecord{3, "วิลาฬ ชาญชัย", 1, "1300101198146", 21, 1, "ไม่มี", "tetracyclines", "0957212978", "api1@gmail.com", "บ้านเลขที่ 35/6 ถนนสายไหม อำเภอเมือง ตำบลในเมือง จังหวัดนครราชสีมา 30000", 1},
			Patientrecord{3, "วิชัย ชาญชัย", 1, "1300101198136", 21, 1, "ไม่มี", "penicillin", "0957212976", "api2@gmail.com", "บ้านเลขที่ 35/6 ถนนสายไหม อำเภอเมือง ตำบลในเมือง จังหวัดนครราชสีมา 30000", 1},
			Patientrecord{3, "วิลินา ชาญชัย", 1, "1300101198126", 21, 1, "ไม่มี", "aspirin", "0957212979", "api3@gmail.com", "บ้านเลขที่ 35/6 ถนนสายไหม อำเภอเมือง ตำบลในเมือง จังหวัดนครราชสีมา 30000", 1},
			Patientrecord{3, "ตฤนชัย บังเกิด", 1, "1300101198186", 22, 1, "ไม่มี", "tetracyclines", "0957252978", "api1@gmail.com", "บ้านเลขที่ 35/6 ถนนสายไหม อำเภอเมือง ตำบลในเมือง จังหวัดนครราชสีมา 30000", 1},
			Patientrecord{3, "กรรรชัย โสภา", 1, "1300101198836", 36, 1, "ไม่มี", "penicillin", "0857212976", "api2@gmail.com", "บ้านเลขที่ 35/6 ถนนสายไหม อำเภอเมือง ตำบลในเมือง จังหวัดนครราชสีมา 30000", 1},
			Patientrecord{3, "อลัน พันตา", 1, "1300101198826", 27, 1, "ไม่มี", "aspirin", "0957212179", "api3@gmail.com", "บ้านเลขที่ 35/6 ถนนสายไหม อำเภอเมือง ตำบลในเมือง จังหวัดนครราชสีมา 30000", 1},
		},
	}
	for _, pr := range Patientrecords.Patientrecord {
		p, err := client.Prename.
			Query().
			Where(prename.IDEQ(int(pr.Prename))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		g, err := client.Gender.
			Query().
			Where(gender.IDEQ(int(pr.Gender))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		bt, err := client.Bloodtype.
			Query().
			Where(bloodtype.IDEQ(int(pr.Bloodtype))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		m, err := client.Medicalrecordstaff.
			Query().
			Where(medicalrecordstaff.IDEQ(int(pr.Medicalrecordstaff))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Patientrecord.
			Create().
			SetEdgesOfPrename(p).
			SetName(pr.Name).
			SetEdgesOfGender(g).
			SetIdcardnumber(pr.Idcardnumber).
			SetAge(pr.Age).
			SetEdgesOfBloodtype(bt).
			SetDisease(pr.Disease).
			SetAllergic(pr.Allergic).
			SetPhonenumber(pr.Phonenumber).
			SetEmail(pr.Email).
			SetHome(pr.Home).
			SetEdgesOfMedicalrecordstaff(m).
			SetDate(time.Now().Local()).
			Save(context.Background())
	}

	//^^^*******************************************************************^^^
	// Set Postman By Historytaking System No.5
	//vvv*******************************************************************vvv

	//Set nurse data
	nurses := Nurses{
		Nurse: []Nurse{
			Nurse{"Paonrat Panjainam", "Nurse123456", "พยาบาลวิชาชีพ", 6},
			Nurse{"Name Surname", "Nurse001122", "พยาบาลวิชาชีพ", 1},
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
			SetEdgesOfUser(u).
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

	// Set Postman By Treatment System No.4 Part1
	//vvv*******************************************************************vvv

	//Set Typetreatment data
	typetreatments := Typetreatments{
		Typetreatment: []Typetreatment{
			Typetreatment{"ตรวจผ่าตัด"},
			Typetreatment{"ตรวจทั่วไป"},
			Typetreatment{"ตรวจวินิจฉัย"},
		},
	}
	for _, tm := range typetreatments.Typetreatment {
		client.Typetreatment.
			Create().
			SetTypetreatment(tm.Typetreatment).
			Save(context.Background())
	}
	//^^^*******************************************************************^^^

	// Set Doctorinformation output No.6
	//vvv...................................................................vvv
	Doctorinfos := Doctorinfos{
		Doctorinfo: []Doctorinfo{
			Doctorinfo{"Thanawat", "Srikaewsiew", "0800740864", "BEN10UT1000", 1, 1, 1, 1},
			Doctorinfo{"Paonrat", "Tangtong", "0810740864", "BEN20UT1000", 1, 1, 1, 1},
			Doctorinfo{"Watcharaphong", "Taramol", "0809757643", "BEN20UT1000", 1, 1, 1, 1},
			Doctorinfo{"Payut", "Jundara", "0899654444", "BEN20UT1000", 1, 1, 1, 1},
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

		_, err = client.Doctorinfo.
			Create().
			SetDoctorname(doc.Doctorname).
			SetDoctorsurname(doc.Doctorsurname).
			SetTelephonenumber(doc.Telephonenumber).
			SetLicensenumber(doc.Licensenumber).
			SetEdgesOfDepartment(dp).
			SetEdgesOfEducationlevel(el).
			SetEdgesOfOfficeroom(or).
			SetEdgesOfPrename(pn).
			Save(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	//^^^...................................................................^^^

	// Set Postman By Treatment System No.4 Part2
	//vvv*******************************************************************vvv

	//Set Doctor data
	Doctors := Doctors{
		Doctor: []Doctor{
			Doctor{1, 4},
			Doctor{2, 1},
			Doctor{3, 3},
			Doctor{4, 5},
		},
	}
	for _, d := range Doctors.Doctor {

		di, err := client.Doctorinfo.
			Query().
			Where(doctorinfo.IDEQ(int(d.Doctorinfo))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		u, err := client.User.
			Query().
			Where(user.IDEQ(int(d.User))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.Doctor.
			Create().
			SetEdgesOfDoctorinfo(di).
			SetEdgesOfUser(u).
			Save(context.Background())
	}

	//^^^*******************************************************************^^^

	// Set Postman By Bill System No.1
	//vvv*******************************************************************vvv

	//Set Financier data
	financiers := Financiers{
		Financier: []Financier{
			Financier{"นางไอริณ อารัญย์", 2},
			Financier{"นางดังฝัน พ่วงดี", 1},
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
			SetEdgesOfUser(u).
			Save(context.Background())
	}

	//Set Paytype data
	paytypes := Paytypes{
		Paytype: []Paytype{
			Paytype{"ชำระผ่านธนาคารออนไลน์"},
			Paytype{"ชำระผ่านบัตรเครดิต"},
			Paytype{"เงินสด"},
		},
	}
	for _, pt := range paytypes.Paytype {
		client.Paytype.
			Create().
			SetPaytype(pt.paytype).
			Save(context.Background())
	}

	//Set Treatment data

	treatments := Treatments{
		Treatment: []Treatment{
			Treatment{"ปวดหัว", "ไมเกรน", "ยาแก้ปวด ยาลดความเครียด", 2, 1, 1},
			Treatment{"เจ็บขา", "กล้ามเนื้ออักเสพ", "ยาแก้ปวด ยาแก้อักเสพ", 2, 2, 2},
			Treatment{"มีอาการใจสั่น", "เลือดไปเลี้ยงหัวใจไม่พอ", "ยาบำรุงเลือด", 2, 2, 3},
			Treatment{"มีผดผื่นขึ้นตามตัว", "ภูมิแพ้", "ยาแก้แพ้", 3, 2, 4},
			Treatment{"ผ่าตัดตามนัด", "ผ่าตัดไส้ติ่ง", "ยาฆ่าเชื้อ", 1, 2, 5},
			Treatment{"ปวดคอ", "กล้ามเนื้ออักเสพ", "ยาแก้ปวด ยาแก้อักเสพ", 2, 2, 6},
			Treatment{"หน้ามืด ใจสั่น", "ความดันสูง", "ยาลดความดัน", 3, 2, 1},
			Treatment{"ชาปลายมือปลายเท้า", "ปลายประสาทอักเสบ", "อาหารเสริมวิตามินบี", 3, 2, 2},
			Treatment{"ไข้สูง ไอเป็นเลือด", "ปอดอักเสบ", "ยาแก้อักเสพ", 2, 2, 3},
			Treatment{"ไม่สามารถขยับใบหน้าได้", "กล้ามเนื้ออ่อนแรง", "ไม่มี", 3, 2, 3},
			Treatment{"เป็นลม", "ขาดสารอาหาร", "อาหารเสริม", 2, 2, 3},
			Treatment{"มีอาการหน้ามืดบ่อยครั้ง", "เลือดไปเลี้ยงสมองไม่พอ ออฟฟิสซินโดรม", "ยาคลายกล้ามเนื้อ", 3, 2, 3},
			Treatment{"เจ็บขา", "กล้ามเนื้ออักเสพ", "ยาแก้ปวด ยาแก้อักเสพ", 2, 2, 5},
		},
	}
	for _, t := range treatments.Treatment {
		tt, err := client.Typetreatment.
			Query().
			Where(typetreatment.IDEQ(int(t.Typetreatment))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		d, err := client.Doctor.
			Query().
			Where(doctor.IDEQ(int(t.Doctor))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		m, err := client.Patientrecord.
			Query().
			Where(patientrecord.IDEQ(int(t.Patientrecord))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		times := time.Now().Local()
		_, err = client.Treatment.
			Create().
			SetSymptom(t.Symptom).
			SetTreat(t.Treat).
			SetMedicine(t.Medicine).
			SetEdgesOfTypetreatment(tt).
			SetEdgesOfDoctor(d).
			SetDatetreat(times).
			SetEdgesOfPatientrecord(m).
			Save(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	//Set Unpaybill Data
	unpaybills := Unpaybills{
		Unpaybill: []Unpaybill{
			Unpaybill{"Paid", 1},
			Unpaybill{"Paid", 2},
			Unpaybill{"Paid", 3},
			Unpaybill{"Paid", 4},
			Unpaybill{"Paid", 5},
			Unpaybill{"Paid", 7},
			Unpaybill{"Paid", 8},
			Unpaybill{"Unpay", 9},
			Unpaybill{"Unpay", 10},
			Unpaybill{"Unpay", 11},
			Unpaybill{"Unpay", 12},
			Unpaybill{"Unpay", 13},

		},
	}
	for _, ub := range unpaybills.Unpaybill {
		t, err := client.Treatment.
			Query().
			Where(treatment.IDEQ(ub.treatment)).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.Unpaybill.
			Create().
			SetStatus(ub.status).
			SetEdgesOfTreatment(t).
			Save(context.Background())
	}

	//Set Bill Data for search bill
	bills := Bills{
		Bill: []Bill{
			Bill{159,"นายพงษ์นรินทร์ จันทร์สุข","0912345678",1,3,1},
			Bill{999,"นางอ่านค่ำ นอนเช้า","0912300678",2,2,1},
			Bill{4500,"นายหมองตัน สงสัย","0812345678",3,1,1},
			Bill{963,"นายตาช่ำ แสงแยง","0987654321",4,1,1},
			Bill{2580,"นางสาวญักกิรมุทะ มากโข","0999999999",5,2,1},
			Bill{2452,"นางเห็นใจ ผู้เรียน","0654987123",7,3,1},
			Bill{8090,"นายอ่ำ อึ้ง","0682145369",8,3,1},
		},
	}
	for _, b := range bills.Bill {
		unpay, err := client.Unpaybill.
			Query().
			Where(unpaybill.IDEQ(b.Unpaybill)).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		paytype, err := client.Paytype.
			Query().
			Where(paytype.IDEQ(b.Paytype)).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		f, err := client.Financier.
			Query().
			Where(financier.IDEQ(b.Financier)).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		times := time.Now().Local()
		client.Bill.
		Create().
		SetPayer(b.Payer).
		SetPayercontact(b.Payercontact).
		SetAmount(b.Amount).
		SetDate(times).
		SetEdgesOfPaytype(paytype).
		SetEdgesOfOfficer(f).
		SetEdgesOfUnpaybill(unpay).
		Save(context.Background())	
	}
	//^^^*******************************************************************^^^

	//^^^^^^^^^-------------------------------------------------------------------^^^^^^^^^

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
