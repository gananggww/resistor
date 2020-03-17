package controllers

import (
	"resistor/services"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (o *ErrorController) Error404() {
	rs := services.Response{
		Msg:    "Page Not Found",
		Status: false,
	}
	o.Data["json"] = rs
	o.ServeJSON()
}

func (o *ErrorController) Error500() {
	rs := services.Response{
		Msg:    "Internal Server Error",
		Status: false,
	}
	o.Data["json"] = rs
	o.ServeJSON()
}

func (o *ErrorController) ErrorDb() {
	rs := services.Response{
		Msg:    "database is down",
		Status: false,
	}
	o.Data["json"] = rs
	o.ServeJSON()
}
