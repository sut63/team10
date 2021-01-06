package controllers

import (
	"context"
	_"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/typetreatment"
)

// TypetreatmentController defines the struct for the Typetreatment controller
type TypetreatmentController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateTypetreatment handles POST requests for adding Typetreatment entities
// @Summary Create Typetreatment
// @Description Create Typetreatment
// @ID create-Typetreatment
// @Accept   json
// @Produce  json
// @Param Typetreatment body ent.Typetreatment true "Typetreatment entity"
// @Success 200 {object} ent.Typetreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Typetreatments [post]
func (ctl *TypetreatmentController) CreateTypetreatment(c *gin.Context) {
	obj := ent.Typetreatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Typetreatment binding failed",
		})
		return
	}

	ttm, err := ctl.client.Typetreatment.
		Create().
		SetTypetreatment(obj.Typetreatment).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ttm)
}

// GetTypetreatment handles GET requests to retrieve a Typetreatment entity
// @Summary Get a Typetreatment entity by ID
// @Description get Typetreatment by ID
// @ID get-Typetreatment
// @Produce  json
// @Param id path int true "Typetreatment ID"
// @Success 200 {object} ent.Typetreatment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Typetreatments/{id} [get]
func (ctl *TypetreatmentController) GetTypetreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ttm, err := ctl.client.Typetreatment.
		Query().
		Where(typetreatment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ttm)
}

// ListTypetreatment handles request to get a list of Typetreatment entities
// @Summary List Typetreatment entities
// @Description list Typetreatment entities
// @ID list-Typetreatment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Typetreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Typetreatments [get]
func (ctl *TypetreatmentController) ListTypetreatment(c *gin.Context) {
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

	Typetreatment, err := ctl.client.Typetreatment.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, Typetreatment)
}

// UpdateTypetreatment handles PUT requests to update a Typetreatment entity
// @Summary Update a Typetreatment entity by ID
// @Description update Typetreatment by ID
// @ID update-Typetreatment
// @Accept   json
// @Produce  json
// @Param id path int true "Typetreatment ID"
// @Param Typetreatment body ent.Typetreatment true "Typetreatment entity"
// @Success 200 {object} ent.Typetreatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Typetreatments/{id} [put]
func (ctl *TypetreatmentController) UpdateTypetreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Typetreatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Typetreatment binding failed",
		})
		return
	}
	obj.ID = int(id)
	ttm, err := ctl.client.Typetreatment.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ttm)
}

// NewTypetreatmentController creates and registers handles for the Typetreatment controller
func NewTypetreatmentController(router gin.IRouter, client *ent.Client) *TypetreatmentController {
	uc := &TypetreatmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitTypetreatmentController registers routes to the main engine
func (ctl *TypetreatmentController) register() {
	Typetreatments := ctl.router.Group("/Typetreatments")

	Typetreatments.GET("", ctl.ListTypetreatment)

	// CRUD
	Typetreatments.POST("", ctl.CreateTypetreatment)
	Typetreatments.GET(":id", ctl.GetTypetreatment)
	Typetreatments.PUT(":id", ctl.UpdateTypetreatment)
	//Typetreatments.DELETE(":id", ctl.DeleteTypetreatment)
}
