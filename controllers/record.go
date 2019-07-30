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

	// if name == "" && beginTimeStamp == 0 {
	// 	c.Data["json"] = models.PublicRecord
	// }
	// if name != "" && beginTimeStamp == 0 {
	// 	steps := make(map[int64]string)
	// 	for _, v := range models.PublicRecord {
	// 		steps[v.TimeStamp] = v.NameStep[name]
	// 	}
	// 	c.Data["json"] = steps
	// }
	// if name == "" && beginTimeStamp != 0 {
	// 	if endTimeStamp == 0 {
	// 		c.Data["json"] = "Missing endTimeStamp"
	// 	} else {
	// 		records := make([]models.Record, 0)
	// 		for _, v := range models.PublicRecord {
	// 			if v.TimeStamp >= beginTimeStamp && v.TimeStamp <= endTimeStamp {
	// 				records = append(records, v)
	// 			}
	// 		}
	// 		c.Data["json"] = records
	// 	}
	// }
	// if name != "" && beginTimeStamp != 0 {
	// 	steps := make(map[int64]string)
	// 	for _, v := range models.PublicRecord {
	// 		if v.TimeStamp >= beginTimeStamp && v.TimeStamp <= endTimeStamp {
	// 			steps[v.TimeStamp] = v.NameStep[name]
	// 		}
	// 	}
	// 	c.Data["json"] = steps
	// }
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
	//fmt.Println(r)
	//models.PublicRecord = append(models.PublicRecord, r)
	//fmt.Println(len(models.PublicRecord))
	//fmt.Println(models.PublicRecord[0])
	models.SubmitData(models.PostJSONToRecords(s))
	res := make(map[string]string)
	res["status_code"] = "200"
	c.Data["json"] = res
	c.ServeJSON()
}
