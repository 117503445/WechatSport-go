package controllers

import (
	"encoding/json"
	"fmt"

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
	o := s{Name: "qwe"}
	fmt.Printf("%s\n", o)

	v, err := json.Marshal(o)
	if err != nil {
		fmt.Printf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", v)
	c.Data["json"] = v
	c.ServeJSON()
}
