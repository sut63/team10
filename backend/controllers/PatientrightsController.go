package controllers

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/abilitypatientrights"
	"github.com/team10/app/ent/insurance"
	"github.com/team10/app/ent/medicalrecordstaff"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/patientrights"
)

// PatientrightsController defines the struct for the Patientrights controller
type PatientrightsController struct {
	client *ent.Client
	router gin.IRouter
}

// Patientrights defines the struct for the Patientrights
type Patientrights struct {
	//PermissionDate string

	//Abilitypatientrights int

	Patientrecord        int
	Insurance            int
	Medicalrecordstaff   int
	Abilitypatientrights int
	Permission           string
	PermissionArea       string
	Responsible          string
}

// CreatePatientrights handles POST requests for adding Patientrights entities
// @Summary Create Patientrights
// @Description Create Patientrights
// @ID create-Patientrights
// @Accept   json
// @Produce  json
// @Param Patientrights body Patientrights true "Patientrights entity"
// @Success 200 {object} Patientrights
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightss [post]
func (ctl *PatientrightsController) CreatePatientrights(c *gin.Context) {
	obj := Patientrights{}

	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่สามารถสร้าง สิทธิ์ได้",
		})
		return
	}



	Patientrecord, err := ctl.client.Patientrecord.
		Query().
		Where(patientrecord.IDEQ(int(obj.Patientrecord))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่พบผู้ป่วย",
		})
		return
	}

	Medicalrecordstaff, err := ctl.client.Medicalrecordstaff.
		Query().
		Where(medicalrecordstaff.IDEQ(int(obj.Medicalrecordstaff))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่สามารถระบุพนักงานเวชระเบียงได้",
		})
		return
	}

	Insurance, err := ctl.client.Insurance.
		Query().
		Where(insurance.IDEQ(int(obj.Insurance))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่พบประกัน",
		})
		return
	}

	Abilitypatientrights, err := ctl.client.Abilitypatientrights.
		Query().
		Where(abilitypatientrights.IDEQ(int(obj.Abilitypatientrights))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ไม่พบความสามารถสิทธิ์",
		})
		return
	}

	settime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	t, err := time.Parse(time.RFC3339, settime)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "ระบบบันทึกเวลาผิดพลาด",
		})
		return
	}

	u, err := ctl.client.Patientrights.
		Create().
		SetPermissionDate(t).
		SetPermission(obj.Permission).
		SetPermissionArea(obj.PermissionArea).
		SetResponsible(obj.Responsible).
		SetEdgesOfPatientrightsAbilitypatientrights(Abilitypatientrights).
		SetEdgesOfPatientrightsPatientrecord(Patientrecord).
		SetEdgesOfPatientrightsMedicalrecordstaff(Medicalrecordstaff).
		SetEdgesOfPatientrightsInsurance(Insurance).
		Save(context.Background())

	if err != nil {
		e := strings.Split(err.Error(), ":")
		c.JSON(400, gin.H{
			"status": false,
			"error":  e[2],
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   u,
	})
}

// GetPatientrights handles GET requests to retrieve a patientrights entity
// @Summary Get a Patientrights entity by ID
// @Description get Patientrights by ID
// @ID get-patientrights
// @Produce  json
// @Param id path int true "Patientrights ID"
// @Success 200 {object} ent.Patientrights
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightss/{id} [get]
func (ctl *PatientrightsController) GetPatientrights(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Patientrights.
		Query().
		WithEdgesOfPatientrightsInsurance().
		WithEdgesOfPatientrightsAbilitypatientrights().
		WithEdgesOfPatientrightsPatientrecord().
		WithEdgesOfPatientrightsMedicalrecordstaff().
		Where(patientrights.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// SearchPatientrights handles request to get a search of patientrights entities
// @Summary Search patientrights entities
// @Description search patientrights entities
// @Produce json
// @Param name  query string false "Name"
// @Success 200 {array} ent.Patientrights
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrights [get]
func (  ctl *PatientrightsController) SearchPatientrights(c *gin.Context) {
	

	
	nameQuery := c.Request.URL.Query().Get("name")
	
	
	
	patientrightss, err := ctl.client.Patientrights.
		Query().
		WithEdgesOfPatientrightsInsurance().
		WithEdgesOfPatientrightsAbilitypatientrights().
		WithEdgesOfPatientrightsPatientrecord().
		WithEdgesOfPatientrightsMedicalrecordstaff().
		Where(patientrights.
			HasEdgesOfPatientrightsPatientrecordWith(patientrecord.
				NameContains(nameQuery))).

		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patientrightss)
}


// DeletePatientrights handles DELETE requests to delete a patientrights entity
// @Summary Delete a patientrights entity by ID
// @Description get patientrights by ID
// @ID delete-patientrights
// @Produce  json
// @Param id path int true "Patientrights ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightss/{id} [delete]
func (ctl *PatientrightsController) DeletePatientrights(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Patientrights.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// ListPatientrights handles request to get a list of patientrights entities
// @Summary List patientrights entities
// @Description list patientrights entities
// @ID list-patientrights
// @Produce json
// @Success 200 {array} ent.Patientrights
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightss [get]
func (ctl *PatientrightsController) ListPatientrights(c *gin.Context) {


	patientrightss, err := ctl.client.Patientrights.
		Query().
		WithEdgesOfPatientrightsInsurance().
		WithEdgesOfPatientrightsAbilitypatientrights().
		WithEdgesOfPatientrightsPatientrecord().
		WithEdgesOfPatientrightsMedicalrecordstaff().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patientrightss)
}

// UpdatePatientrights handles PUT requests to update a patientrights entity
// @Summary Update a patientrights entity by ID
// @Description update patientrights by ID
// @ID update-patientrights
// @Accept   json
// @Produce  json
// @Param id path int true "Patientrights ID"
// @Param patientrights body ent.Patientrights true "Patientrights entity"
// @Success 200 {array} ent.Patientrights
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientrightss/{id} [put]
func (ctl *PatientrightsController) UpdatePatientrights(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Patientrights{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patientrights binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Patientrights.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewPatientrightsController creates and registers handles for the patientrights controller
func NewPatientrightsController(router gin.IRouter, client *ent.Client) *PatientrightsController {
	uc := &PatientrightsController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitPatientrightsController registers routes to the main engine
func (ctl *PatientrightsController) register() {
	patientrightss := ctl.router.Group("/patientrightss")
	patientrights := ctl.router.Group("/patientrights")

	
	patientrights.GET("", ctl.SearchPatientrights)
	patientrightss.GET("", ctl.ListPatientrights)
	// CRUD
	patientrightss.POST("", ctl.CreatePatientrights)
	patientrightss.GET(":id", ctl.GetPatientrights)
	patientrightss.PUT(":id", ctl.UpdatePatientrights)
	patientrightss.DELETE(":id", ctl.DeletePatientrights)
}