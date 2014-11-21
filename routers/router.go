package routers

import (
	"related-news/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/news/:id([0-9]+)", &controllers.NewsController{}, "get:Get"),
		beego.NSRouter("/news/:id([0-9]+)/append", &controllers.NewsController{},"get:Append"),
		beego.NSRouter("/news/analyze", &controllers.NewsController{},"get:Analyze"),
		beego.NSRouter("/news/analyze_dedup", &controllers.NewsController{},"get:DedupAnalyze"),
		beego.NSRouter("/news/init_news", &controllers.NewsController{},"get:InitNews"),
		beego.NSRouter("/news/len", &controllers.NewsController{},"get:Len"),
		beego.NSRouter("/news/add", &controllers.NewsController{},"get:Add"),

	)
	beego.AddNamespace(ns)
}
