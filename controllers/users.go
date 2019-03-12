package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"loveHome/models"
)

type UsersController struct {
	beego.Controller
}

func (this *UsersController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *UsersController) Reg() {
	beego.Info("Reg is connect")
	resp := make(map[string]interface{})
	defer c.RetData(resp)
	json.Unmarshal(c.Ctx.Input.RequestBody,&resp)

	//执行入库操作
	//beego.Info(`mobile:`,resp["mobile"])
	//beego.Info(`password:`,resp["password"])
	//beego.Info(`sms_code:`,resp["sms_code"])

	o := orm.NewOrm()
	user := models.User{}
	user.Name = resp["mobile"].(string)
	user.Password = resp["password"].(string)
	user.Mobile = resp["mobile"].(string)

	id,err := o.Insert(&user)
	if err !=nil{
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		return
	}

	beego.Info("注册成功,id=",id)
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	c.SetSession("name",user.Name)




}