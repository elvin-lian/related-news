package routers

import (
	"github.com/astaxie/beego"
	"github.com/elvin-lian/related-news/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/news/:id([0-9]+)", &controllers.NewsController{}, "get:Get"),
		beego.NSRouter("/news/:id([0-9]+)/append", &controllers.NewsController{}, "get:Append"),
		beego.NSRouter("/news/analyze", &controllers.NewsController{}, "get:Analyze"),
		beego.NSRouter("/news/init_news", &controllers.NewsController{}, "get:InitNews"),
		beego.NSRouter("/news/len", &controllers.NewsController{}, "get:Len"),
		beego.NSRouter("/news/add", &controllers.NewsController{}, "get:Add"),
		beego.NSRouter("/news/dedup_analyze", &controllers.NewsController{}, "get:DedupAnalyze"),
		beego.NSRouter("/news/dedup_len", &controllers.NewsController{}, "get:DedupLen"),
		beego.NSRouter("/news/dedup_add", &controllers.NewsController{}, "get:DedupAdd"),
		beego.NSRouter("/news/dedup_check", &controllers.NewsController{}, "get:DedupCheck"),
	)
	beego.AddNamespace(ns)
}
