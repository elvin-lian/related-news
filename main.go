package main

import (
	"github.com/astaxie/beego"
	_ "github.com/elvin-lian/related-news/routers"
)

func main() {
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLogger("file", `{"filename":"logs/app.log"}`)
		beego.BeeLogger.DelLogger("console")
	}
	beego.Run()
}
