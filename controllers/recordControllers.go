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

//Get 获得记录,根据 name 和 beginTimeStamp endTimeStamp 进行筛选
func (c *RecordController) Get() {
	name := c.GetString("name")
	beginTimeStamp, _ := c.GetInt64("beginTimeStamp")
	endTimeStamp, _ := c.GetInt64("endTimeStamp")

	c.Data["json"] = models.GetRecords(name, beginTimeStamp, endTimeStamp)
	c.ServeJSON()
}

//Post 提交记录
func (c *RecordController) Post() {
	fmt.Println((string)(c.Ctx.Input.RequestBody[:]))
	var s []byte
	s = c.Ctx.Input.RequestBody
	var r models.Record
	json.Unmarshal(s, &r)
	models.SubmitData(models.PostJSONToRecords(s))
	res := make(map[string]string)
	res["status_code"] = "200"
	c.Data["json"] = res
	c.ServeJSON()
}
