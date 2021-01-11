package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/doctor"
	"github.com/team10/app/ent/doctorinfo"
)

// DoctorController defines the struct for the doctor controller
type DoctorController struct {
	client *ent.Client
	router gin.IRouter
}

// Doctor defines the struct for the Bill entity
type Doctor struct {
	Doctorinfo      int	
}

// CreateDoctor handles POST requests for adding doctor entities
// @Summary Create doctor
// @Description Create doctor
// @ID create-doctor
// @Accept   json
// @Produce  json
// @Param doctor body Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors [post]
func (ctl *DoctorController) CreateDoctor(c *gin.Context) {
	obj := Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctor binding failed",
		})
		return
	}
	di, err := ctl.client.Doctorinfo.
		Query().
		Where(doctorinfo.IDEQ(int(obj.Doctorinfo))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "doctorinfo not found",
		})
		return
	}
	
	d, err := ctl.client.Doctor.
		Create().
		SetDoctorinfo(di).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetDoctor handles GET requests to retrieve a doctor entity
// @Summary Get a doctor entity by ID
// @Description get doctor by ID
// @ID get-doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors/{id} [get]
func (ctl *DoctorController) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	d, err := ctl.client.Doctor.
		Query().
		WithDoctorinfo().
		Where(doctor.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, d)
}

// ListDoctor handles request to get a list of doctor entities
// @Summary List doctor entities
// @Description list doctor entities
// @ID list-doctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors [get]
func (ctl *DoctorController) ListDoctor(c *gin.Context) {
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

	d, err := ctl.client.Doctor.
		Query().
		WithDoctorinfo().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, d)
}

// DeleteDoctor handles DELETE requests to delete a doctor entity
// @Summary Delete a doctor entity by ID
// @Description get doctor by ID
// @ID delete-doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors/{id} [delete]
func (ctl *DoctorController) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Doctor.
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

// UpdateDoctor handles PUT requests to update a doctor entity
// @Summary Update a doctor entity by ID
// @Description update doctor by ID
// @ID update-doctor
// @Accept   json
// @Produce  json
// @Param id path int true "Doctor ID"
// @Param doctor body ent.Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctors/{id} [put]
func (ctl *DoctorController) UpdateDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctor binding failed",
		})
		return
	}
	obj.ID = int(id)
	d, err := ctl.client.Doctor.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, d)
}

// NewDoctorController creates and registers handles for the doctor controller
func NewDoctorController(router gin.IRouter, client *ent.Client) *DoctorController {
	uc := &DoctorController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDoctorController registers routes to the main engine
func (ctl *DoctorController) register() {
	doctors := ctl.router.Group("/doctors")

	doctors.GET("", ctl.ListDoctor)

	// CRUD
	doctors.POST("", ctl.CreateDoctor)
	doctors.GET(":id", ctl.GetDoctor)
	doctors.PUT(":id", ctl.UpdateDoctor)
	doctors.DELETE(":id", ctl.DeleteDoctor)
}
