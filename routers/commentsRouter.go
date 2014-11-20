package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["related-news/controllers:NewsController"] = append(beego.GlobalControllerRouter["related-news/controllers:NewsController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["related-news/controllers:NewsController"] = append(beego.GlobalControllerRouter["related-news/controllers:NewsController"],
		beego.ControllerComments{
			"Append",
			`/:id/append`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["related-news/controllers:NewsController"] = append(beego.GlobalControllerRouter["related-news/controllers:NewsController"],
		beego.ControllerComments{
			"Analyze",
			`/analyze`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["related-news/controllers:NewsController"] = append(beego.GlobalControllerRouter["related-news/controllers:NewsController"],
		beego.ControllerComments{
			"InitNews",
			`/init_news`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["related-news/controllers:NewsController"] = append(beego.GlobalControllerRouter["related-news/controllers:NewsController"],
		beego.ControllerComments{
			"Len",
			`/len`,
			[]string{"get"},
			nil})

}
