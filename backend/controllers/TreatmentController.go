package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/doctor"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/typetreatment"
	
)

// TreatmentController defines the struct for the Treatment controller
type TreatmentController struct {
	client *ent.Client
	router gin.IRouter
}

// Treatment defines the struct for the Treatment entity
type Treatment struct {
	Symptom       string
	Treat     string
	Medicine      string
	Datetreat     string
	Typetreatment int
	Doctor        int
	Patientrecord int
}

// CreateTreatment handles POST requests for adding Treatment entities
// @Summary Create Treatment
// @Description Create Treatment
// @ID create-Treatment
// @Accept   json
// @Produce  json
// @Param Treatment body Treatment true "Treatment entity"
// @Success 200 {object} Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Treatments [post]
func (ctl *TreatmentController) CreateTreatment(c *gin.Context) {
	obj := Treatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Treatment binding failed",
		})
		return
	}

	if (int(obj.Doctor) + int(obj.Typetreatment) + int(obj.Patientrecord) + len(obj.Symptom) + len(obj.Treat) + len(obj.Medicine)  ) <= 0 {
		c.JSON(400, gin.H{
			"error": "บันทึกข้อมูลไม่สำเร็จ",
		})
		return
	}

	ttm, err := ctl.client.Typetreatment.
		Query().
		Where(typetreatment.IDEQ(int(obj.Typetreatment))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Typetreatment not found",
		})
		return
	}

	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}

	pr, err := ctl.client.Patientrecord.
		Query().
		Where(patientrecord.IDEQ(int(obj.Patientrecord))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patientrecord not found",
		})
		return
	}

	t := time.Now().Local()
	
	if obj.Symptom == "" {
		c.JSON(400, gin.H{
			"error": "Symptom ไม่ถูกต้อง กรุณากรอกเป็นภาษาไทย",
		})
		return
	}
	 
	if  obj.Treat == ""{
		c.JSON(400, gin.H{
			"error": "Treat ไม่ถูกต้อง กรุณากรอกเป็นภาษาไทย",
		})
		return
	}
	if obj.Medicine == ""{
		c.JSON(400, gin.H{
			"error": "Medicine ไม่ถูกต้อง กรุณากรอกเป็นภาษาไทย",
		})
		return
	}
	tm, err := ctl.client.Treatment.
		Create().
		SetSymptom(obj.Symptom).
		SetTreat(obj.Treat).		
		SetMedicine(obj.Medicine).
		SetDatetreat(t).
		SetEdgesOfTypetreatment(ttm).
		SetEdgesOfDoctor(d).
		SetEdgesOfPatientrecord(pr).
		Save(context.Background())

		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error":  err.Error(),
			})
			return
		}else {
		
		
		
// Create Unpaybill
	ctl.client.Unpaybill.
		Create().
		SetEdgesOfTreatment(tm).
		SetStatus("Unpay").
		Save(context.Background())

		c.JSON(200, gin.H{
			"status": true,
			"data":   tm,
		})
	return
	}
}

// GetTreatment handles GET requests to retrieve a Treatment entity
// @Summary Get a Treatment entity by ID
// @Description get Treatment by ID
// @ID get-Treatment
// @Produce  json
// @Param id path int true "Treatment ID"
// @Success 200 {array} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Treatments/{id} [get]
func (ctl *TreatmentController) GetTreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	tm, err := ctl.client.Treatment.
		Query().
		WithEdgesOfDoctor(func (q *ent.DoctorQuery){
			q.QueryEdgesOfDoctorinfo()
			q.WithEdgesOfDoctorinfo()
		}).
		WithEdgesOfTypetreatment().
		WithEdgesOfPatientrecord().
		Where(treatment.HasEdgesOfPatientrecordWith(patientrecord.IDEQ(int(id)))).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, tm)
}

// ListTreatment handles request to get a list of Treatment entities
// @Summary List Treatment entities
// @Description list Treatment entities
// @ID list-Treatment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Treatments [get]
func (ctl *TreatmentController) ListTreatment(c *gin.Context) {

	Treatment, err := ctl.client.Treatment.
		Query().
		WithEdgesOfDoctor(func (q *ent.DoctorQuery){
			q.QueryEdgesOfDoctorinfo()
			q.WithEdgesOfDoctorinfo()
		}).
		WithEdgesOfTypetreatment().
		WithEdgesOfPatientrecord().
		WithEdgesOfUnpaybills().
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, Treatment)
}

// DeleteTreatment handles DELETE requests to delete a Treatment entity
// @Summary Delete a Treatment entity by ID
// @Description get Treatment by ID
// @ID delete-Treatment
// @Produce  json
// @Param id path int true "Treatment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Treatments/{id} [delete]
func (ctl *TreatmentController) DeleteTreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Treatment.
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

// NewTreatmentController creates and registers handles for the Treatment controller
func NewTreatmentController(router gin.IRouter, client *ent.Client) *TreatmentController {
	uc := &TreatmentController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *TreatmentController) register() {
	Treatments := ctl.router.Group("/Treatments")

	Treatments.GET("", ctl.ListTreatment)
	Treatments.POST("", ctl.CreateTreatment)
	Treatments.GET(":id", ctl.GetTreatment)
	Treatments.DELETE(":id", ctl.DeleteTreatment)
}
