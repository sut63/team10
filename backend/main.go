package main
 
import (
   "context"
   "fmt"
   "log"
  
  
 
   "github.com/team10/app/controllers"
   _ "github.com/team10/app/docs"
   "github.com/team10/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"

   "github.com/team10/app/ent/abilitypatientrights"


  


)
// struct By team 10
//-------------------------------------------------------------------


// Struct By B6103866
//*******************************************************************

// Patientrightstypes defines the struct for the Patientrightstypes
type Patientrightstypes struct {
	Patientrightstype []Patientrightstype
}

// Patientrightstype defines the struct for the Patientrightstype
type Patientrightstype struct {
	Permission          string
    PermissionArea      string
    Responsible         string
    Abilitypatientrights    int
}

// Abilitypatientrightss defines the struct for the Abilitypatientrightss
type Abilitypatientrightss struct {
	Abilitypatientrights []Abilitypatientrights
}

// Abilitypatientrights defines the struct for the Abilitypatientrights
type Abilitypatientrights struct {
	Operative       string
    MedicalSupplies string
    Examine         string
}

// Patientrightss defines the struct for the Patientrightss
type Patientrightss struct {
	Patientrights []Patientrights
}

// Patientrights defines the struct for the Patientrights
type Patientrights struct {
	
	Patientrightstype    int
	Patientrecord        int
	Insurance            int
	Medicalrecordstaff   int
}
//*******************************************************************

type Paytypes struct {
	Paytype []Paytype
}
type Paytype struct {
	paytype string
}

type Financiers struct {
	Financier []Financier
}
type Financier struct {
	name string
}



//-------------------------------------------------------------------




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
   // Controller By B6103866 
   //+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
   controllers.NewPatientrightsController(v1, client)
   controllers.NewPatientrightstypeController(v1, client)
   controllers.NewAbilitypatientrightsController(v1, client)
   controllers.NewInsuranceController(v1, client)
   //+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

   // Controller By 
   //+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
   controllers.NewMedicalrecordstaffController(v1, client)
   controllers.NewPatientrecordController(v1, client)
<<<<<<< HEAD
   //+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++


=======
   //Controller By Bill System
   controllers.NewBillController(v1, client)
	controllers.NewFinancierController(v1, client)
	controllers.NewPaytypeController(v1, client)
	controllers.NewUnpaybillController(v1, client)
>>>>>>> 4cdc72865dccead9f8786a5d9d1a419934b8c1fe
   // Set Postman By Team10
   //-------------------------------------------------------------------

   // Set medicalrecordstaff Data
   	Patientrecord := []string{"Khatadet khianchainat","nara haru","morani rode","faratell yova","pulla visan","omaha yad",}
	for _, r := range Patientrecord {
		client.Patientrecord.
			Create().
			SetName(r).
			Save(context.Background())
	}

	// Set medicalrecordstaff Data
	
	for i := 1; i < 5; i++ {
		client.Medicalrecordstaff.
			Create().
			
			Save(context.Background())
	}


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
			Abilitypatientrights{"100", "100","100"},
			Abilitypatientrights{"50", "100","100"},
			Abilitypatientrights{"50", "100","50"},
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
		Patientrightstype : []Patientrightstype{
			Patientrightstype{"จ่ายตรง", "ทั่วประเทศ","ราชการ", 1},
			Patientrightstype{"ผู้สูงอายุ", "ทั่วประเทศ","นาย แหวง" , 2},
			Patientrightstype{"สุขภาพถ้วนหน้า", "โคราช","ราชการ" ,2},
			Patientrightstype{"อุบัติเหตุสบายใจ", "ทั่วประเทศ","นาย มา", 2},
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
	//Set Financier data

	financiers := Financiers{
		Financier: []Financier{
			Financier{"Nutchaporn Klinrod"},
			Financier{"Name Surname"},
		},
	}
	for _, f := range financiers.Financier {
		client.Financier.
			Create().
			SetName(f.name).
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
			SetPaytype(pt.paytype)
	}
   //-------------------------------------------------------------------

   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}
 
