package controllers

import (
	"context"
	"strconv"
	"fmt"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/unpaybill"
	_"github.com/team10/app/ent/treatment"
	"github.com/gin-gonic/gin"
)


// UnpaybillController defines the struct for the unpaybill controller
type UnpaybillController struct {
	client *ent.Client
	router gin.IRouter
}

// GetUnpaybill handles GET requests to retrieve a unpaybill entity
// @Summary Get a unpaybill entity by ID
// @Description get unpaybill by ID
// @ID get-unpaybill
// @Produce  json
// @Param id path int true "Unpaybill ID"
// @Success 200 {object} ent.Unpaybill
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills/{id} [get]
func (ctl *UnpaybillController) GetUnpaybill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Unpaybill.
		Query().
		WithEdgesOfTreatment().
		Where(unpaybill.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListUnpaybill handles request to get a list of unpaybill entities
// @Summary List unpaybill entities
// @Description list unpaybill entities
// @ID list-unpaybill
// @Produce json
// @Success 200 {array} ent.Unpaybill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills [get]
func (ctl *UnpaybillController) ListUnpaybill(c *gin.Context) {
	unpaybills, err := ctl.client.Unpaybill.
		Query().
		Where(unpaybill.StatusEQ("Unpay")).
		WithEdgesOfTreatment(func (t *ent.TreatmentQuery){
			t.QueryEdgesOfPatientrecord()
			t.WithEdgesOfPatientrecord()
		}).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, unpaybills)
}

// DeleteUnpaybill handles DELETE requests to delete a unpaybill entity
// @Summary Delete a unpaybill entity by ID
// @Description get unpaybill by ID
// @ID delete-unpaybill
// @Produce  json
// @Param id path int true "Unpaybill ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills/{id} [delete]
func (ctl *UnpaybillController) DeleteUnpaybill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Unpaybill.
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

// UpdateUnpaybill handles PUT requests to update a unpaybill entity
// @Summary Update a unpaybill entity by ID
// @Description update unpaybill by ID
// @ID update-unpaybill
// @Accept   json
// @Produce  json
// @Param id path int true "Unpaybill ID"
// @Param unpaybill body ent.Unpaybill true "Unpaybill entity"
// @Success 200 {object} ent.Unpaybill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills/{id} [put]
func (ctl *UnpaybillController) UpdateUnpaybill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Unpaybill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "unpaybill binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Unpaybill.
		UpdateOneID(int(id)).
		SetStatus("Paid").
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewUnpaybillController creates and registers handles for the unpaybill controller
func NewUnpaybillController(router gin.IRouter, client *ent.Client) *UnpaybillController {
	uc := &UnpaybillController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *UnpaybillController) register() {
	unpaybills := ctl.router.Group("/unpaybills")

	unpaybills.GET("", ctl.ListUnpaybill)
	unpaybills.GET(":id", ctl.GetUnpaybill)
	unpaybills.PUT(":id", ctl.UpdateUnpaybill)
	unpaybills.DELETE(":id", ctl.DeleteUnpaybill)
}