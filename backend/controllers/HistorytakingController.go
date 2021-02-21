package controllers

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/symptomseverity"
)

// HistorytakingController defines the struct for the historytaking controller
type HistorytakingController struct {
	client *ent.Client
	router gin.IRouter
}

// Historytaking defines the struct for the Historytaking entity
type Historytaking struct {
	Nurse           int
	Department      int
	Symptomseverity int
	Patientrecord   int
	Hight           string
	Weight          string
	Temp            string
	Pulse           string
	Respiration     string
	Bp              string
	Oxygen          string
	Symptom         string
	Datetime        string
}

// CreateHistorytaking handles POST requests for adding historytaking entities
// @Summary Create historytaking
// @Description Create historytaking
// @ID create-historytaking
// @Accept   json
// @Produce  json
// @Param historytaking body Historytaking true "Historytaking entity"
// @Success 200 {object} Historytaking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /historytaking [post]
func (ctl *HistorytakingController) CreateHistorytaking(c *gin.Context) {
	obj := Historytaking{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Historytaking binding failed",
		})
		return
	}

	if (int(obj.Symptomseverity) + int(obj.Department) + int(obj.Patientrecord) + len(obj.Hight) + len(obj.Weight) + len(obj.Temp) + len(obj.Pulse) + len(obj.Respiration) + len(obj.Bp) + len(obj.Oxygen) + len(obj.Symptom) ) <= 0 {
		c.JSON(400, gin.H{
			"error": "บันทึกข้อมูลไม่สำเร็จ",
		})
		return
	}

	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "โปรดเลือกพยาบาลที่ซักประวัติ",
		})
		return
	}

	ss, err := ctl.client.Symptomseverity.
		Query().
		Where(symptomseverity.IDEQ(int(obj.Symptomseverity))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "โปรดเลือกระดับอาการสำคัญ",
		})
		return
	}

	d, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "โปรดเลือกแผนก",
		})
		return
	}

	pt, err := ctl.client.Patientrecord.
		Query().
		Where(patientrecord.IDEQ(int(obj.Patientrecord))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "โปรดเลือกผู้ป่วย",
		})
		return
	}

	times := time.Now().Local()
	
	var h float32
	hights, err := strconv.ParseFloat(obj.Hight, 64);
		h = float32(hights)
		
	if len(obj.Hight) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกค่าส่วนสูง",
		})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ค่าส่วนสูงต้องเป็นตัวเลขจำนวนเต็มหรือตัวเลขทศนิยมเท่านั้น",
		})
		return
	}

	var w float32
	weights, err := strconv.ParseFloat(obj.Weight, 64);
		w = float32(weights)

	if len(obj.Weight) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกค่าน้ำหนัก",
		})
		return
	}
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ค่าน้ำหนักต้องเป็นตัวเลขจำนวนเต็มหรือตัวเลขทศนิยมเท่านั้น",
		})
		return
	}

	var t float32
	temps, err := strconv.ParseFloat(obj.Temp, 64); 
		t = float32(temps)

	if len(obj.Temp) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกค่าอุณหภูมิร่างกาย",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ค่าอุณหภูมิร่างกายต้องเป็นตัวเลขจำนวนเต็มหรือตัวเลขทศนิยมเท่านั้น",
		})
		return
	}
	
	pulses, err := strconv.Atoi(obj.Pulse)

	if len(obj.Pulse) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกค่าชีพจร",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ค่าชีพจรต้องเป็นตัวเลขจำนวนเต็มเท่านั้น",
		})
		return
	}

	respirations, err := strconv.Atoi(obj.Respiration)
	
	if len(obj.Respiration) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกค่าการหายใจ",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ค่าการหายใจต้องเป็นตัวเลขจำนวนเต็มเท่านั้น",
		})
		return
	}

	bps, err := strconv.Atoi(obj.Bp)

	if len(obj.Bp) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกค่าความดันโลหิต",
		})
		return
	}
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ค่าความดันโลหิตต้องเป็นตัวเลขจำนวนเต็มเท่านั้น",
		})
		return
	}

	oxygens, err := strconv.Atoi(obj.Oxygen); 

	if len(obj.Oxygen) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกค่าออกซิเจนในกระแสเลือด",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ค่าออกซิเจนในกระแสเลือดต้องเป็นตัวเลขจำนวนเต็มหรือตัวเลขทศนิยมเท่านั้น",
		})
		return
	}

	if len(obj.Symptom) <= 0 {
		c.JSON(400, gin.H{
			"error": "โปรดกรอกอาการสำคัญ",
		})
		return
	}
	
	ht, err := ctl.client.Historytaking.
		Create().
		SetEdgesOfNurse(n).
		SetEdgesOfDepartment(d).
		SetEdgesOfSymptomseverity(ss).
		SetEdgesOfPatientrecord(pt).
		SetHight(h).
		SetWeight(w).
		SetTemp(t).
		SetPulse(pulses).
		SetRespiration(respirations).
		SetBp(bps).
		SetOxygen(oxygens).
		SetSymptom(obj.Symptom).
		SetDatetime(times).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   ht,
	})
}

// GetHistorytaking handles GET requests to retrieve a historytaking entity
// @Summary Get a historytaking entity by ID
// @Description get historytaking by ID
// @ID get-historytaking
// @Produce  json
// @Param id path int true "Historytaking ID"
// @Success 200 {array} ent.Historytaking
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /historytaking/{id} [get]
func (ctl *HistorytakingController) GetHistorytaking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pr, err := ctl.client.Historytaking.
		Query().
		WithEdgesOfNurse().
		WithEdgesOfDepartment().
		WithEdgesOfSymptomseverity().
		WithEdgesOfPatientrecord().
		Where(historytaking.HasEdgesOfPatientrecordWith(patientrecord.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pr)
}

// ListHistorytaking handles request to get a list of historytaking entities
// @Summary List historytaking entities
// @Description list historytaking entities
// @ID list-historytaking
// @Produce json
// @Success 200 {array} ent.Historytaking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /historytaking [get]
func (ctl *HistorytakingController) ListHistorytaking(c *gin.Context) {

	historytaking, err := ctl.client.Historytaking.
		Query().
		WithEdgesOfNurse().
		WithEdgesOfDepartment().
		WithEdgesOfSymptomseverity().
		WithEdgesOfPatientrecord().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, historytaking)
}

// NewHistorytakingController creates and registers handles for the historytaking controller
func NewHistorytakingController(router gin.IRouter, client *ent.Client) *HistorytakingController {
	htc := &HistorytakingController{
		client: client,
		router: router,
	}
	htc.register()
	return htc
}

// InitHistorytakingController registers routes to the main engine
func (ctl *HistorytakingController) register() {
	historytaking := ctl.router.Group("/historytaking")

	historytaking.GET("", ctl.ListHistorytaking)

	// CRUD
	historytaking.POST("", ctl.CreateHistorytaking)
	historytaking.GET(":id", ctl.GetHistorytaking)
}
