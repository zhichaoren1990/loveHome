package controllers

import (
	"github.com/astaxie/beego"
	"loveHome/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *SessionController) GetSessionData(){
	beego.Info("Session connect success")
	user := models.User{}
	resp := make(map[string]interface{})
	defer c.RetData(resp)




	name := c.GetSession("name")
	if name != nil{
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
		return
	}

	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

}


/**
	删除session
 */

func (c *SessionController) DeleteSessionData() {
	beego.Info("delete session connect")
	resp := make(map[string]interface{})
	defer  c.RetData(resp)
	c.DelSession("name")

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}
