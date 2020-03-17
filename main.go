package main

import (
	database "resistor/database"
	_ "resistor/routers"

	"github.com/astaxie/beego"
)

func init() {
	database.Run()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
