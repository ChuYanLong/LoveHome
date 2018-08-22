package controllers

import (
	"github.com/astaxie/beego"
	"LoveHome/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (c *HouseIndexController) GetHouseIndex() {
	resp := make(map[string]interface{})


	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	c.RetData(resp)
}
func (c *HouseIndexController)RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}

