package controllers

import (
	"github.com/astaxie/beego"
)

//APIController is
type APIController struct {
	beego.Controller
}
type s struct {
	Name string
}

//Get is
func (c *APIController) Get() {
	c.Data["json"] = "API works"
	c.ServeJSON()
}
