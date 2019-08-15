package controllers

import (
	"WechatSport-go/models"

	"github.com/astaxie/beego"
)

//PeopleControllers 人列表的返回
type PeopleControllers struct {
	beego.Controller
}


//Get 返回人列表
func (c *PeopleControllers) Get() {
	
	c.Data["json"] = models.GetPeopleList()
	c.ServeJSON()
}
