package controllers

import (
	"context"
	"fmt"
	"strconv"
	
	"time"
	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/paytype"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/unpaybill"
	"github.com/team10/app/ent/patientrecord"
)

// BillController defines the struct for the bill controller
type BillController struct {
	client *ent.Client
	router gin.IRouter
}

// Bill defines the struct for the Bill entity
type Bill struct {
	Note        string
	Payercontact string
	Amount       string
	Paytype      int
	Financier    int
	Unpaybill    int
}

// CreateBill handles POST requests for adding bill entities
// @Summary Create bill
// @Description Create bill
// @ID create-bill
// @Accept   json
// @Produce  json
// @Param bill body Bill true "Bill entity"
// @Success 200 {object} Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [post]
func (ctl *BillController) CreateBill(c *gin.Context) {
	obj := Bill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bill binding failed",
		})
		return
	}
	pt, err := ctl.client.Paytype.
		Query().
		Where(paytype.IDEQ(int(obj.Paytype))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "bill not found",
		})
		return
	}
	f, err := ctl.client.Financier.
		Query().
		Where(financier.IDEQ(int(obj.Financier))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "bill not found",
		})
		return
	}
	ub, err := ctl.client.Unpaybill.
		Query().
		WithEdgesOfTreatment().
		Where(unpaybill.IDEQ(int(obj.Unpaybill))).
		Only(context.Background())
		if err != nil {
			c.JSON(400, gin.H{
				"error": "bill not found",
			})
			return
	}
	t := time.Now().Local()
	am, err := strconv.Atoi(obj.Amount)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Amount have to be number",
		})
		return
}
	u, err := ctl.client.Bill.
		Create().
		SetNote(obj.Note).
		SetPayercontact(obj.Payercontact).
		SetAmount(am).
		SetDate(t).
		SetEdgesOfPaytype(pt).
		SetEdgesOfOfficer(f).
		SetEdgesOfUnpaybill(ub).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   u,
	})
}

// GetBill handles GET requests to retrieve a bill entity
// @Summary Get a bill entity by ID
// @Description get bill by ID
// @ID get-bill
// @Produce  json
// @Param id path int true "Bill ID"
// @Success 200 {array} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills/{id} [get]
func (ctl *BillController) GetBill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Bill.
		Query().
		WithEdgesOfUnpaybill(func (q *ent.UnpaybillQuery){
			q.QueryEdgesOfTreatment()
			q.WithEdgesOfTreatment(func (t *ent.TreatmentQuery){
				t.QueryEdgesOfPatientrecord()
				t.WithEdgesOfPatientrecord()
			})
		}).
		WithEdgesOfPaytype().
		WithEdgesOfOfficer().
		Where(bill.HasEdgesOfUnpaybillWith(unpaybill.HasEdgesOfTreatmentWith(treatment.HasEdgesOfPatientrecordWith(patientrecord.IDEQ(int(id)))))).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListBill handles request to get a list of bill entities
// @Summary List bill entities
// @Description list bill entities
// @ID list-bill
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [get]
func (ctl *BillController) ListBill(c *gin.Context) {
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

	bills, err := ctl.client.Bill.
		Query().
		WithEdgesOfOfficer().
		WithEdgesOfPaytype().
		WithEdgesOfUnpaybill(func (q *ent.UnpaybillQuery){
			q.QueryEdgesOfTreatment()
			q.WithEdgesOfTreatment(func (t *ent.TreatmentQuery){
				t.QueryEdgesOfPatientrecord()
				t.WithEdgesOfPatientrecord()
			})
		}).
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bills)
}

// DeleteBill handles DELETE requests to delete a bill entity
// @Summary Delete a bill entity by ID
// @Description get bill by ID
// @ID delete-bill
// @Produce  json
// @Param id path int true "Bill ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills/{id} [delete]
func (ctl *BillController) DeleteBill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bill.
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

// NewBillController creates and registers handles for the bill controller
func NewBillController(router gin.IRouter, client *ent.Client) *BillController {
	uc := &BillController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *BillController) register() {
	bills := ctl.router.Group("/bills")

	bills.GET("", ctl.ListBill)
	bills.POST("", ctl.CreateBill)
	bills.GET(":id", ctl.GetBill)
	bills.DELETE(":id", ctl.DeleteBill)
}
