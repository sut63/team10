package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/registrar"
)

// RegistrarController defines the struct for the registrar controller
type RegistrarController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRegistrar handles POST requests for adding registrar entities
// @Summary Create registrar
// @Description Create registrar
// @ID create-registrar
// @Accept   json
// @Produce  json
// @Param registrar body ent.Registrar true "Registrar entity"
// @Success 200 {object} ent.Registrar
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registrars [post]
func (ctl *RegistrarController) CreateRegistrar(c *gin.Context) {
	obj := ent.Registrar{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "registrar binding failed",
		})
		return
	}

	u, err := ctl.client.Registrar.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetRegistrar handles GET requests to retrieve a registrar entity
// @Summary Get a registrar entity by ID
// @Description get registrar by ID
// @ID get-registrar
// @Produce  json
// @Param id path int true "Registrar ID"
// @Success 200 {object} ent.Registrar
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registrars/{id} [get]
func (ctl *RegistrarController) GetRegistrar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Registrar.
		Query().
		Where(registrar.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListRegistrar handles request to get a list of registrar entities
// @Summary List registrar entities
// @Description list registrar entities
// @ID list-registrar
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Registrar
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registrars [get]
func (ctl *RegistrarController) ListRegistrar(c *gin.Context) {
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

	registrars, err := ctl.client.Registrar.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, registrars)
}

// DeleteRegistrar handles DELETE requests to delete a registrar entity
// @Summary Delete a registrar entity by ID
// @Description get registrar by ID
// @ID delete-registrar
// @Produce  json
// @Param id path int true "Registrar ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registrars/{id} [delete]
func (ctl *RegistrarController) DeleteRegistrar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Registrar.
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

// UpdateRegistrar handles PUT requests to update a registrar entity
// @Summary Update a registrar entity by ID
// @Description update registrar by ID
// @ID update-registrar
// @Accept   json
// @Produce  json
// @Param id path int true "Registrar ID"
// @Param registrar body ent.Registrar true "Registrar entity"
// @Success 200 {object} ent.Registrar
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /registrars/{id} [put]
func (ctl *RegistrarController) UpdateRegistrar(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Registrar{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "registrar binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Registrar.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewRegistrarController creates and registers handles for the registrar controller
func NewRegistrarController(router gin.IRouter, client *ent.Client) *RegistrarController {
	uc := &RegistrarController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRegistrarController registers routes to the main engine
func (ctl *RegistrarController) register() {
	registrars := ctl.router.Group("/registrars")

	registrars.GET("", ctl.ListRegistrar)

	// CRUD
	registrars.POST("", ctl.CreateRegistrar)
	registrars.GET(":id", ctl.GetRegistrar)
	registrars.PUT(":id", ctl.UpdateRegistrar)
	registrars.DELETE(":id", ctl.DeleteRegistrar)
}
