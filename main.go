package main

import (
	_ "github.com/elvin-lian/related-news/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLogger("file", `{"filename":"logs/app.log"}`)
		beego.BeeLogger.DelLogger("console")
	}
	beego.Run()
}
