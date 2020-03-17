package controllers

import (
	"encoding/json"
	"fmt"
	"resistor/models"
	"resistor/services"
	"runtime"

	"github.com/astaxie/beego"
)

//NewsController is a object controller in news
type NewsController struct {
	beego.Controller
}

//GetNews is a controller for a create news
func (o *NewsController) GetNews() {

	var perPage int
	var page int

	o.Ctx.Input.Bind(&perPage, "per_page")
	o.Ctx.Input.Bind(&page, "page")

	var resPage int = page

	if page <= 0 {
		page = 1
		resPage = 1
	}

	page = (page - 1) * perPage

	response, err := models.GetNews(perPage, page)

	count, _ := models.GetNewsCount()

	if err != nil {
		rs := services.Response{
			Msg:    fmt.Sprintf("%s", err),
			Status: false,
			Data:   nil,
		}
		o.Data["json"] = rs
		o.ServeJSON()
	}

	rs := services.ResponseP{
		Msg:    "Sukses",
		Status: true,
		Data:   response,
		Paginate: services.Paginate{
			Page:    resPage,
			PerPage: perPage,
			Count:   len(count),
		},
	}
	o.Data["json"] = rs
	o.ServeJSON()
}

//PostNews is a controller for a create news
func (o *NewsController) PostNews() {

	runtime.GOMAXPROCS(2)
	var ob models.News

	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	response, err := models.CreateNews(ob)

	if err != nil {
		rs := services.Response{
			Msg:    fmt.Sprintf("%s", err),
			Status: false,
			Data:   nil,
		}
		o.Data["json"] = rs
		o.ServeJSON()
	}

	es := services.ES{
		Id:      response.Id,
		Created: response.Created,
	}

	go es.CreateLogNews()

	rs := services.Response{
		Msg:    "Sukses",
		Status: true,
		Data:   response,
	}
	o.Data["json"] = rs
	o.ServeJSON()
}
