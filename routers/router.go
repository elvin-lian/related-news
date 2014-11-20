// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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
		beego.NSRouter("/news/init_news", &controllers.NewsController{},"get:InitNews"),
		beego.NSRouter("/news/len", &controllers.NewsController{},"get:Len"),
		beego.NSRouter("/news/add", &controllers.NewsController{},"get:Add"),

	)
	beego.AddNamespace(ns)
}
