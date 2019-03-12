package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaController) GetArea()  {

	beego.Info("area connect success")
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	var areas []models.Area
	o := orm.NewOrm()
	num , err := o.QueryTable("area").All(&areas)
	if err != nil{
		beego.Error("area 读取数据错误")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if num ==0 {
		beego.Info("area 读取为空")

		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	resp["data"] = areas
	beego.Info("读取数据成功")

	beego.Info("data success area = ",resp,"num=",num)

}