package router

import (
	"net/http"
	"strconv"

	"github.com/hel2o/eam/model"

	"github.com/hel2o/eam/db"

	"github.com/gin-gonic/gin"
)

type MSG struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count"`
	Date  interface{} `json:"data"`
}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
	return
}

func printPage(c *gin.Context) {
	c.HTML(http.StatusOK, "print.html", nil)
	return
}

func lookPage(c *gin.Context) {
	eam, err := db.GetEAMByID(c.Query("id"))
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	if eam.IsVirtual == "on" {
		eam.IsVirtual = "是"
	} else {
		eam.IsVirtual = "否"
	}
	c.HTML(http.StatusOK, "look.html", gin.H{"data": eam})
	return
}

func getEAMRouter(c *gin.Context) {
	limit, err := strconv.Atoi(c.PostForm("limit"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4000, Msg: err.Error()})
		return
	}
	page, err := strconv.Atoi(c.PostForm("page"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4001, Msg: err.Error()})
		return
	}
	ema, count, err := db.GetEAM(c.PostForm("search"), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4002, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MSG{Code: 0, Date: ema, Count: count})
}

func addEAMRouter(c *gin.Context) {
	var eam model.EAM
	err := c.BindJSON(&eam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4000, Msg: err.Error()})
		return
	}
	err = db.AddEAM(eam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4001, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MSG{Code: 3000})
}

func modifyEAMRouter(c *gin.Context) {
	var eam model.EAM
	err := c.BindJSON(&eam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4000, Msg: err.Error()})
		return
	}
	err = db.ModifyEAM(eam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4001, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MSG{Code: 3000})
}

func delEAMRouter(c *gin.Context) {
	var eam []model.EAM
	err := c.BindJSON(&eam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4000, Msg: err.Error()})
		return
	}
	var del int64
	del, err = db.DelEAM(eam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MSG{Code: 4001, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, MSG{Code: 3000, Count: del})
}
