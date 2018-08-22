package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"LoveHome/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Post() {
	resp := make(map[string]interface{})
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &resp)
	if err != nil {
		resp["errno"] = 4001
		resp["errmsg"] = "查询数据错误"
	}
	/*
	{
		mobile: "4564",
		password: "123456",
		sms_code: "sdfasf"
	}
	 */
	beego.Debug(`resp["mobile"] = `, resp["mobile"])
	beego.Debug(`resp["password"] = `, resp["password"])
	beego.Debug(`resp["sms_code"] = `, resp["sms_code"])
	o := orm.NewOrm()
	user := models.User{}
	user.Mobile = resp["mobile"].(string)
	user.Password_hash = resp["password"].(string)
	user.Name = resp["mobile"].(string)

	id, err := o.Insert(&user)
	if err != nil {

	 resp["errno"] = 4002
	 resp["errmsg"] = "注册失败"
	 return
	}

	beego.Debug("reg success!, id = ", id)

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = "注册成功"

	c.SetSession("name", user.Name)

	defer c.RetData(resp)
}
func (c *UserController)RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}