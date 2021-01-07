package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/officeroom"
)

// OfficeroomController defines the struct for the office controller
type OfficeroomController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateOfficeroom handles POST requests for adding officeroom entities
// @Summary Create officeroom
// @Description Create officeroom
// @ID create-officeroom
// @Accept   json
// @Produce  json
// @Param officeroom body ent.Officeroom true "Officeroom entity"
// @Success 200 {object} ent.Officeroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officerooms [post]
func (ctl *OfficeroomController) CreateOfficeroom(c *gin.Context) {
	obj := ent.Officeroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "officeroom binding failed",
		})
		return
	}

	u, err := ctl.client.Officeroom.
		Create().
		SetRoomnumber(obj.Roomnumber).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetOfficeroom handles GET requests to retrieve a officeroom entity
// @Summary Get a officeroom entity by ID
// @Description get officeroom by ID
// @ID get-officeroom
// @Produce  json
// @Param id path int true "Officeroom ID"
// @Success 200 {object} ent.Officeroom
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officerooms/{id} [get]
func (ctl *OfficeroomController) GetOfficeroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Officeroom.
		Query().
		Where(officeroom.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListOfficeroom handles request to get a list of officeroom entities
// @Summary List officeroom entities
// @Description list officeroom entities
// @ID list-officeroom
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Officeroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officerooms [get]
func (ctl *OfficeroomController) ListOfficeroom(c *gin.Context) {
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

	officerooms, err := ctl.client.Officeroom.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, officerooms)
}

// DeleteOfficeroom handles DELETE requests to delete a officeroom entity
// @Summary Delete a officeroom entity by ID
// @Description get officeroom by ID
// @ID delete-officeroom
// @Produce  json
// @Param id path int true "Officeroom ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officerooms/{id} [delete]
func (ctl *OfficeroomController) DeleteOfficeroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Officeroom.
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

// UpdateOfficeroom handles PUT requests to update a officeroom entity
// @Summary Update a officeroom entity by ID
// @Description update officeroom by ID
// @ID update-officeroom
// @Accept   json
// @Produce  json
// @Param id path int true "Officeroom ID"
// @Param officeroom body ent.Officeroom true "Officeroom entity"
// @Success 200 {object} ent.Officeroom
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /officerooms/{id} [put]
func (ctl *OfficeroomController) UpdateOfficeroom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Officeroom{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "officeroom binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Officeroom.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewOfficeroomController creates and registers handles for the officeroom controller
func NewOfficeroomController(router gin.IRouter, client *ent.Client) *OfficeroomController {
	uc := &OfficeroomController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitOfficeroomController registers routes to the main engine
func (ctl *OfficeroomController) register() {
	officerooms := ctl.router.Group("/officerooms")

	officerooms.GET("", ctl.ListOfficeroom)

	// CRUD
	officerooms.POST("", ctl.CreateOfficeroom)
	officerooms.GET(":id", ctl.GetOfficeroom)
	officerooms.PUT(":id", ctl.UpdateOfficeroom)
	officerooms.DELETE(":id", ctl.DeleteOfficeroom)
}
