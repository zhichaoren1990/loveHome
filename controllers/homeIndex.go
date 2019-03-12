package controllers

import "github.com/astaxie/beego"

type HomeIndexController struct {
	beego.Controller
}

func (this *HomeIndexController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *HomeIndexController) GetHomeIndex(){
	beego.Info("HomeIndex connect success")
	 resp := make(map[string]interface{})
	 resp["errno"] = 4001
	 resp["errmsg"] = "查询失败"
	 c.RetData(resp)
}
