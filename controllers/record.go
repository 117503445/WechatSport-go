package controllers

import (
	"WechatSport-go/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

//RecordController 记录 相关的控制器
type RecordController struct {
	beego.Controller
}

//Get 获得记录,根据 name 和 date 进行筛选
func (c *RecordController) Get() {
	steps := make(map[int64]string)
	name := c.GetString("name")
	if name != "" {
		for _, v := range models.PublicRecord {
			steps[v.TimeStamp] = v.NameStep[name]
		}
	}
	c.Data["json"] = steps
	c.ServeJSON()
}

//Post 提交记录
func (c *RecordController) Post() {
	fmt.Println((string)(c.Ctx.Input.RequestBody[:]))
	var s []byte
	s = c.Ctx.Input.RequestBody
	var r models.Record
	json.Unmarshal(s, &r)
	//fmt.Println(r)
	models.PublicRecord = append(models.PublicRecord, r)
	//fmt.Println(len(models.PublicRecord))
	//fmt.Println(models.PublicRecord[0])
	c.Ctx.WriteString("")
}
