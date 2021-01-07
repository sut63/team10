package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/prename"
)

// PrenameController defines the struct for the prename controller
type PrenameController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePrename handles POST requests for adding prename entities
// @Summary Create prename
// @Description Create prename
// @ID create-prename
// @Accept   json
// @Produce  json
// @Param prename body ent.Prename true "Prename entity"
// @Success 200 {object} ent.Prename
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Prenames [post]
func (ctl *PrenameController) CreatePrename(c *gin.Context) {
	obj := ent.Prename{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "prename binding failed",
		})
		return
	}

	u, err := ctl.client.Prename.
		Create().
		SetPrefix(obj.Prefix).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetPrename handles GET requests to retrieve a prename entity
// @Summary Get a prename entity by ID
// @Description get prename by ID
// @ID get-prename
// @Produce  json
// @Param id path int true "prename ID"
// @Success 200 {object} ent.Prename
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prenames/{id} [get]
func (ctl *PrenameController) GetPrename(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Prename.
		Query().
		Where(prename.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListPrename handles request to get a list of prename entities
// @Summary List prename entities
// @Description list prename entities
// @ID list-prename
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Prename
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prenames [get]
func (ctl *PrenameController) ListPrename(c *gin.Context) {
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

	prenames, err := ctl.client.Prename.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, prenames)
}

// DeletePrename handles DELETE requests to delete a prename entity
// @Summary Delete a prename entity by ID
// @Description get prename by ID
// @ID delete-prename
// @Produce  json
// @Param id path int true "Prename ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prenames/{id} [delete]
func (ctl *PrenameController) DeletePrename(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Prename.
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

// UpdatePrename handles PUT requests to update a prename entity
// @Summary Update a prename entity by ID
// @Description update prename by ID
// @ID update-prename
// @Accept   json
// @Produce  json
// @Param id path int true "Prename ID"
// @Param prename body ent.Prename true "Prename entity"
// @Success 200 {object} ent.Prename
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prenames/{id} [put]
func (ctl *PrenameController) UpdatePrename(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Prename{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "prename binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Prename.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewPrenameController creates and registers handles for the prename controller
func NewPrenameController(router gin.IRouter, client *ent.Client) *PrenameController {
	uc := &PrenameController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitPrenameController registers routes to the main engine
func (ctl *PrenameController) register() {
	prenames := ctl.router.Group("/prenames")

	prenames.GET("", ctl.ListPrename)

	// CRUD
	prenames.POST("", ctl.CreatePrename)
	prenames.GET(":id", ctl.GetPrename)
	prenames.PUT(":id", ctl.UpdatePrename)
	prenames.DELETE(":id", ctl.DeletePrename)
}
