package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/paytype"
)

// PaytypeController defines the struct for the paytype controller
type PaytypeController struct {
	client *ent.Client
	router gin.IRouter
}

// GetPaytype handles GET requests to retrieve a paytype entity
// @Summary Get a paytype entity by ID
// @Description get paytype by ID
// @ID get-paytype
// @Produce  json
// @Param id path int true "Paytype ID"
// @Success 200 {object} ent.Paytype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytypes/{id} [get]
func (ctl *PaytypeController) GetPaytype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Paytype.
		Query().
		Where(paytype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}
// ListPaytype handles request to get a list of paytype entities
// @Summary List paytype entities
// @Description list paytype entities
// @ID list-paytype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Paytype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytypes [get]
func (ctl *PaytypeController) ListPaytype(c *gin.Context) {
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
	paytypes, err := ctl.client.Paytype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, paytypes)
}

// NewPaytypeController creates and registers handles for the paytype controller
func NewPaytypeController(router gin.IRouter, client *ent.Client) *PaytypeController {
	uc := &PaytypeController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *PaytypeController) register() {
	paytypes := ctl.router.Group("/paytypes")

	paytypes.GET("", ctl.ListPaytype)
	paytypes.GET(":id", ctl.GetPaytype)


}
