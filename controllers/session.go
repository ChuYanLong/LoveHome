package controllers

import (
	"github.com/astaxie/beego"
	"LoveHome/models"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) GetSessionData() {
	resp := make(map[string]interface{})
	user := models.User{}
	//user.Name = "long"
	//resp["errno"] = 0
	//resp["errmsg"] = "OK"
	//resp["data"] = user
	name := c.GetSession("name")
	if name != nil{
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}else {
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
	}
	defer c.RetData(resp)
}

func (c *SessionController) DeleteSessionData()  {
 	resp := make(map[string]interface{})
 	defer c.RetData(resp)
 	c.DelSession("name")
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}

func (c *SessionController)RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}
