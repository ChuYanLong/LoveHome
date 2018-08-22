package controllers

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"LoveHome/models"
)

type AreaController struct {
	beego.Controller
}


func (c *AreaController) GetArea() {
	beego.Info("connect success")
	resp := make(map[string]interface{})
	var areas []models.Area

	//从session拿数据
	//从mysql数据库获取到area的数据

	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	//err := o.Read(&area)
	if err != nil{
		//返回错误信息
		beego.Debug("数据错误")
		resp["errno"] = models.RECODE_DATAERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	if num == 0 {
		beego.Debug("没有查到数据")
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	//查询成功
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = &areas
	//无论是失败还是有没有数据,都要打包成json数据返回给前端,
	beego.Debug("query data success! resp = ", resp, "num = ", num)
	defer c.RetData(resp)


}
func (c *AreaController)RetData(resp map[string]interface{}) {
	c.Data["json"] = resp
	c.ServeJSON()
}
