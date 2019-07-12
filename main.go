package main

import (
	_ "WechatSport-go/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/", "static")
	beego.Run()
}
