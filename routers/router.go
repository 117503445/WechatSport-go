package routers

import (
	"WechatSport-go/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api", &controllers.APIController{})
    beego.Router("/api/record", &controllers.RecordController{})
}
