package main

import (
	"WechatSport-go/models"
	_ "WechatSport-go/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/", "static")
	port, err := beego.AppConfig.Int("port")
	fmt.Println(err)
	models.InitDatabase(beego.AppConfig.String("host"), port, beego.AppConfig.String("username"), beego.AppConfig.String("password"))
	beego.Run()
}
