package main

import (
	"WechatSport-go/models"
	_ "WechatSport-go/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/", "static")
	port, _ := beego.AppConfig.Int("port")
	models.InitDatabase(beego.AppConfig.String("host"), port, beego.AppConfig.String("username"), beego.AppConfig.String("password"))
	beego.Run()
}
