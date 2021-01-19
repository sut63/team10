package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/bloodtype"
)

// BloodtypeController defines the struct for the bloodtype controller
type BloodtypeController struct {
	client *ent.Client
	router gin.IRouter
}

// GetBloodtype handles GET requests to retrieve a bloodtype entity
// @Summary Get a bloodtype entity by ID
// @Description get bloodtype by ID
// @ID get-bloodtype
// @Produce  json
// @Param id path int true "Bloodtype ID"
// @Success 200 {object} ent.Bloodtype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bloodtype/{id} [get]
func (ctl *BloodtypeController) GetBloodtype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bt, err := ctl.client.Bloodtype.
		Query().
		Where(bloodtype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bt)
}

// ListBloodtype handles request to get a list of bloodtype entities
// @Summary List bloodtype entities
// @Description list bloodtype entities
// @ID list-bloodtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bloodtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bloodtype [get]
func (ctl *BloodtypeController) ListBloodtype(c *gin.Context) {
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

	bloodtype, err := ctl.client.Bloodtype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bloodtype)
}

// NewBloodtypeController creates and registers handles for the bloodtype controller
func NewBloodtypeController(router gin.IRouter, client *ent.Client) *BloodtypeController {
	bc := &BloodtypeController{
		client: client,
		router: router,
	}
	bc.register()
	return bc
}

// InitBloodtypeController registers routes to the main engine
func (ctl *BloodtypeController) register() {
	bloodtype := ctl.router.Group("/bloodtype")

	bloodtype.GET("", ctl.ListBloodtype)

	// CRUD
	bloodtype.GET(":id", ctl.GetBloodtype)
}
