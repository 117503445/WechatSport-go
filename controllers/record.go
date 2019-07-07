package controllers

import (
	"WechatSport-go/models"
	"encoding/json"
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
	var s []byte
	s = c.Ctx.Input.RequestBody
	var r models.Record
	json.Unmarshal(s, &r)
	//fmt.Println(r)

	models.PublicRecord = append(models.PublicRecord, r)
	//fmt.Println(len(models.PublicRecord))
	fmt.Println(models.PublicRecord[0])
	c.Ctx.WriteString("")
}
