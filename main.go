package main

import (
	_ "related-news/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLogger("file", `{"filename":"logs/app.log","level":6}`)
	}
	beego.Run()
}
