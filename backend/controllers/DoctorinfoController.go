package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/educationlevel"
	"github.com/team10/app/ent/officeroom"
	"github.com/team10/app/ent/prename"
	
)

// DoctorinfoController defines the struct for the doctorinfo controller
type DoctorinfoController struct {
	client *ent.Client
	router gin.IRouter
}

// Doctorinfo defines the struct for the Bill entity
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

// CreateDoctorinfo handles POST requests for adding doctorinfo entities
// @Summary Create doctorinfo
// @Description Create doctorinfo
// @ID create-doctorinfo
// @Accept   json
// @Produce  json
// @Param doctorinfo body Doctorinfo true "Doctorinfo entity"
// @Success 200 {object} Doctorinfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorinfos [post]
func (ctl *DoctorinfoController) CreateDoctorinfo(c *gin.Context) {
	obj := Doctorinfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctorinfo binding failed",
		})
		return
	}
	dp, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "department not found",
		})
		return
	}

	el, err := ctl.client.Educationlevel.
		Query().
		Where(educationlevel.IDEQ(int(obj.Educationlevel))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "educationlevel not found",
		})
		return
	}

	or, err := ctl.client.Officeroom.
		Query().
		Where(officeroom.IDEQ(int(obj.Officeroom))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "officeroom not found",
		})
		return
	}

	pn, err := ctl.client.Prename.
		Query().
		Where(prename.IDEQ(int(obj.Prename))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "prename not found",
		})
		return
	}
	
	u, err := ctl.client.Doctorinfo.
		Create().
		SetDoctorname(obj.Doctorname).
		SetDoctorsurname(obj.Doctorsurname).
		SetTelephonenumber(obj.Telephonenumber).
		SetLicensenumber(obj.Licensenumber).
		SetEdgesOfDepartment(dp).
		SetEdgesOfEducationlevel(el).
		SetEdgesOfOfficeroom(or).
		SetEdgesOfPrename(pn).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}


// GetDoctorinfo handles GET requests to retrieve a doctorinfo entity
// @Summary Get a doctorinfo entity by ID
// @Description get doctorinfo by ID
// @ID get-doctorinfo
// @Produce  json
// @Param id path int true "Doctorinfo ID"
// @Success 200 {object} ent.Doctorinfo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorinfos/{id} [get]
func (ctl *DoctorinfoController) GetDoctorinfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Doctorinfo.
		Query().
		WithEdgesOfDepartment().
		WithEdgesOfEducationlevel().
		WithEdgesOfOfficeroom().
		WithEdgesOfPrename().
		WithEdgesOfDoctor().
		Where(doctorinfo.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListDoctorinfo handles request to get a list of doctorinfo entities
// @Summary List doctorinfo entities
// @Description list doctorinfo entities
// @ID list-doctorinfo
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Doctorinfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorinfos [get]
func (ctl *DoctorinfoController) ListDoctorinfo(c *gin.Context) {
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

	doctorinfos, err := ctl.client.Doctorinfo.
		Query().
		WithEdgesOfDepartment().
		WithEdgesOfEducationlevel().
		WithEdgesOfOfficeroom().
		WithEdgesOfPrename().
		WithEdgesOfDoctor().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, doctorinfos)
}

// DeleteDoctorinfo handles DELETE requests to delete a doctorinfo entity
// @Summary Delete a doctorinfo entity by ID
// @Description get doctorinfo by ID
// @ID delete-doctorinfo
// @Produce  json
// @Param id path int true "Doctorinfo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorinfos/{id} [delete]
func (ctl *DoctorinfoController) DeleteDoctorinfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Doctorinfo.
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

// UpdateDoctorinfo handles PUT requests to update a doctorinfo entity
// @Summary Update a doctorinfo entity by ID
// @Description update doctorinfo by ID
// @ID update-doctorinfo
// @Accept   json
// @Produce  json
// @Param id path int true "Doctorinfo ID"
// @Param doctorinfo body ent.Doctorinfo true "Doctorinfo entity"
// @Success 200 {object} ent.Doctorinfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /doctorinfos/{id} [put]
func (ctl *DoctorinfoController) UpdateDoctorinfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Doctorinfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "doctorinfo binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Doctorinfo.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewDoctorinfoController creates and registers handles for the doctorinfo controller
func NewDoctorinfoController(router gin.IRouter, client *ent.Client) *DoctorinfoController {
	uc := &DoctorinfoController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDoctorinfoController registers routes to the main engine
func (ctl *DoctorinfoController) register() {
	doctorinfos := ctl.router.Group("/doctorinfos")

	doctorinfos.GET("", ctl.ListDoctorinfo)

	// CRUD
	doctorinfos.POST("", ctl.CreateDoctorinfo)
	doctorinfos.GET(":id", ctl.GetDoctorinfo)
	doctorinfos.PUT(":id", ctl.UpdateDoctorinfo)
	doctorinfos.DELETE(":id", ctl.DeleteDoctorinfo)
}
