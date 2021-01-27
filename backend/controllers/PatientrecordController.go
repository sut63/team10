package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/bloodtype"
	"github.com/team10/app/ent/gender"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/prename"
)

// PatientrecordController defines the struct for the patientrecord controller
type PatientrecordController struct {
	client *ent.Client
	router gin.IRouter
}

// Patientrecord defines the struct for the Patientrecord entity
type Patientrecord struct {
	Gender             int
	Medicalrecordstaff int
	Prename            int
	Bloodtype          int
	Name               string
	Idcardnumber       string
	Age                string
	Disease            string
	Allergic           string
	Phonenumber        string
	Email              string
	Home               string
	Date               string
}

// CreatePatientrecord handles POST requests for adding patientrecord entities
// @Summary Create patientrecord
// @Description Create patientrecord
// @ID create-patientrecord
// @Accept   json
// @Produce  json
// @Param patientrecord body Patientrecord true "Patientrecord entity"
// @Success 200 {object} Patientrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecord [post]
func (ctl *PatientrecordController) CreatePatientrecord(c *gin.Context) {
	obj := Patientrecord{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Patientrecord binding failed",
		})
		return
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Gender saving failed",
		})
		return
	}

	m, err := ctl.client.Medicalrecordstaff.
		Query().
		Where(medicalrecordstaff.IDEQ(int(obj.Medicalrecordstaff))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Medicalrecordstaff saving failed",
		})
		return
	}

	p, err := ctl.client.Prename.
		Query().
		Where(prename.IDEQ(int(obj.Prename))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Prename saving failed",
		})
		return
	}

	bt, err := ctl.client.Bloodtype.
		Query().
		Where(bloodtype.IDEQ(int(obj.Bloodtype))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bloodtype saving failed",
		})
		return
	}

	timess := time.Now().Local()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Time saving failed",
		})
		return
	}

	a, err := strconv.Atoi(obj.Age)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Age saving failed",
		})
		return
	}

	pr, err := ctl.client.Patientrecord.
		Create().
		SetEdgesOfGender(g).
		SetEdgesOfMedicalrecordstaff(m).
		SetEdgesOfPrename(p).
		SetName(obj.Name).
		SetIdcardnumber(obj.Idcardnumber).
		SetAge(a).
		SetEdgesOfBloodtype(bt).
		SetDisease(obj.Disease).
		SetAllergic(obj.Allergic).
		SetPhonenumber(obj.Phonenumber).
		SetEmail(obj.Email).
		SetHome(obj.Home).
		SetDate(timess).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   pr,
	})
}

// GetPatientrecord handles GET requests to retrieve a patientrecord entity
// @Summary Get a patientrecord entity by ID
// @Description get patientrecord by ID
// @ID get-patientrecord
// @Produce  json
// @Param id path int true "Patientrecord ID"
// @Success 200 {object} ent.Patientrecord
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecord/{id} [get]
func (ctl *PatientrecordController) GetPatientrecord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pr, err := ctl.client.Patientrecord.
		Query().
		WithEdgesOfGender().
		WithEdgesOfMedicalrecordstaff().
		WithEdgesOfBloodtype().
		WithEdgesOfPrename().
		WithEdgesOfHistorytaking().
		WithEdgesOfPatientrecordPatientrights().
		Where(patientrecord.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pr)
}

// ListPatientrecord handles request to get a list of patientrecord entities
// @Summary List patientrecord entities
// @Description list patientrecord entities
// @ID list-patientrecord
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patientrecord
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrecord [get]
func (ctl *PatientrecordController) ListPatientrecord(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	patientrecord, err := ctl.client.Patientrecord.
		Query().
		WithEdgesOfGender().
		WithEdgesOfMedicalrecordstaff().
		WithEdgesOfBloodtype().
		WithEdgesOfPrename().
		WithEdgesOfPatientrecordPatientrights().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patientrecord)
}

// NewPatientrecordController creates and registers handles for the patientrecord controller
func NewPatientrecordController(router gin.IRouter, client *ent.Client) *PatientrecordController {
	prc := &PatientrecordController{
		client: client,
		router: router,
	}
	prc.register()
	return prc
}

// InitPatientrecordController registers routes to the main engine
func (ctl *PatientrecordController) register() {
	patientrecord := ctl.router.Group("/patientrecord")

	patientrecord.GET("", ctl.ListPatientrecord)

	// CRUD
	patientrecord.POST("", ctl.CreatePatientrecord)
	patientrecord.GET(":id", ctl.GetPatientrecord)
}
