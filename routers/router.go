package routers

import (
	"resistor/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/v1/query", &controllers.QueryController{}, "get:AutocompleteQuery")
	beego.ErrorController(&controllers.ErrorController{})

	ns := beego.NewNamespace("/api/",
		beego.NSRouter("/news", &controllers.NewsController{}, "get:GetNews;post:PostNews"),
	)

	//register namespace
	beego.AddNamespace(ns)
}
