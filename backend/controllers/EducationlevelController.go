package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/educationlevel"
)

// EducationlevelController defines the struct for the educationlevel controller
type EducationlevelController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateEducationlevel handles POST requests for adding educationlevel entities
// @Summary Create educationlevel
// @Description Create educationlevel
// @ID create-educationlevel
// @Accept   json
// @Produce  json
// @Param educationlevel body ent.Educationlevel true "Educationlevel entity"
// @Success 200 {object} ent.Educationlevel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /educationlevels [post]
func (ctl *EducationlevelController) CreateEducationlevel(c *gin.Context) {
	obj := ent.Educationlevel{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "educationlevel binding failed",
		})
		return
	}

	u, err := ctl.client.Educationlevel.
		Create().
		SetLevel(obj.Level).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetEducationlevel handles GET requests to retrieve a educationlevel entity
// @Summary Get a educationlevel entity by ID
// @Description get educationlevel by ID
// @ID get-educationlevel
// @Produce  json
// @Param id path int true "Educationlevel ID"
// @Success 200 {object} ent.Educationlevel
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /educationlevels/{id} [get]
func (ctl *EducationlevelController) GetEducationlevel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Educationlevel.
		Query().
		Where(educationlevel.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListEducationlevel handles request to get a list of educationlevel entities
// @Summary List educationlevel entities
// @Description list educationlevel entities
// @ID list-educationlevel
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Educationlevel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /educationlevels [get]
func (ctl *EducationlevelController) ListEducationlevel(c *gin.Context) {
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

	educationlevels, err := ctl.client.Educationlevel.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, educationlevels)
}

// DeleteEducationlevel handles DELETE requests to delete a educationlevel entity
// @Summary Delete a educationlevel entity by ID
// @Description get educationlevel by ID
// @ID delete-educationlevel
// @Produce  json
// @Param id path int true "Educationlevel ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /educationlevels/{id} [delete]
func (ctl *EducationlevelController) DeleteEducationlevel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Educationlevel.
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

// UpdateEducationlevel handles PUT requests to update a educationlevel entity
// @Summary Update a educationlevel entity by ID
// @Description update educationlevel by ID
// @ID update-educationlevel
// @Accept   json
// @Produce  json
// @Param id path int true "Educationlevel ID"
// @Param educationlevel body ent.Educationlevel true "Educationlevel entity"
// @Success 200 {object} ent.Educationlevel
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /educationlevels/{id} [put]
func (ctl *EducationlevelController) UpdateEducationlevel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Educationlevel{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "educationlevel binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Educationlevel.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewEducationlevelController creates and registers handles for the educationlevel controller
func NewEducationlevelController(router gin.IRouter, client *ent.Client) *EducationlevelController {
	uc := &EducationlevelController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitEducationlevelController registers routes to the main engine
func (ctl *EducationlevelController) register() {
	educationlevels := ctl.router.Group("/educationlevels")

	educationlevels.GET("", ctl.ListEducationlevel)

	// CRUD
	educationlevels.POST("", ctl.CreateEducationlevel)
	educationlevels.GET(":id", ctl.GetEducationlevel)
	educationlevels.PUT(":id", ctl.UpdateEducationlevel)
	educationlevels.DELETE(":id", ctl.DeleteEducationlevel)
}
