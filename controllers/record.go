package controllers

import (
	"WechatSport-go/models"
	"fmt"
	"time"
	"github.com/astaxie/beego"
)

//RecordController 1
type RecordController struct {
	beego.Controller
}



//Get is
func (c *RecordController) Get() {
	fmt.Println(time.Now().Unix())
	s := c.GetString("name")
	fmt.Println(s)
	s = c.GetString("date")
	fmt.Println(s)
	c.Data["json"] = s
	c.ServeJSON()
}

//Post is 7
func (c *RecordController) Post() {
	fmt.Println((string)(c.Ctx.Input.RequestBody[:]))
	models.PublicRecord = append(models.PublicRecord,*new(models.Record))
	fmt.Println(len(models.PublicRecord))
	c.Ctx.WriteString("")
}

