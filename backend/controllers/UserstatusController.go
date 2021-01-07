package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/userstatus"
)

// UserstatusController defines the struct for the userstatus controller
type UserstatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateUserstatus handles POST requests for adding userstatus entities
// @Summary Create userstatus
// @Description Create userstatus
// @ID create-userstatus
// @Accept   json
// @Produce  json
// @Param userstatus body ent.Userstatus true "Userstatus entity"
// @Success 200 {object} ent.Userstatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus [post]
func (ctl *UserstatusController) CreateUserstatus(c *gin.Context) {
	obj := ent.Userstatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "userstatus binding failed",
		})
		return
	}

	u, err := ctl.client.Userstatus.
		Create().
		SetUserstatus(obj.Userstatus).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetUserstatus handles GET requests to retrieve a userstatus entity
// @Summary Get a userstatus entity by ID
// @Description get userstatus by ID
// @ID get-userstatus
// @Produce  json
// @Param id path int true "Userstatus ID"
// @Success 200 {object} ent.Userstatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus/{id} [get]
func (ctl *UserstatusController) GetUserstatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Userstatus.
		Query().
		Where(userstatus.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListUserstatus handles request to get a list of userstatus entities
// @Summary List userstatus entities
// @Description list userstatus entities
// @ID list-userstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Userstatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus [get]
func (ctl *UserstatusController) ListUserstatus(c *gin.Context) {
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

	userstatus, err := ctl.client.Userstatus.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, userstatus)
}

// DeleteUserstatus handles DELETE requests to delete a userstatus entity
// @Summary Delete a userstatus entity by ID
// @Description get userstatus by ID
// @ID delete-userstatus
// @Produce  json
// @Param id path int true "Userstatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus/{id} [delete]
func (ctl *UserstatusController) DeleteUserstatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Userstatus.
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

// UpdateUserstatus handles PUT requests to update a userstatus entity
// @Summary Update a userstatus entity by ID
// @Description update userstatus by ID
// @ID update-userstatus
// @Accept   json
// @Produce  json
// @Param id path int true "Userstatus ID"
// @Param userstatus body ent.Userstatus true "Userstatus entity"
// @Success 200 {object} ent.Userstatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /userstatus/{id} [put]
func (ctl *UserstatusController) UpdateUserstatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Userstatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "userstatus binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Userstatus.
		UpdateOneID(int(id)).
		SetUserstatus(obj.Userstatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewUserstatusController creates and registers handles for the userstatus controller
func NewUserstatusController(router gin.IRouter, client *ent.Client) *UserstatusController {
	uc := &UserstatusController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *UserstatusController) register() {
	userstatus := ctl.router.Group("/userstatus")

	userstatus.GET("", ctl.ListUserstatus)

	// CRUD
	userstatus.POST("", ctl.CreateUserstatus)
	userstatus.GET(":id", ctl.GetUserstatus)
	userstatus.PUT(":id", ctl.UpdateUserstatus)
	userstatus.DELETE(":id", ctl.DeleteUserstatus)
}
