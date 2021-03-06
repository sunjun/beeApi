package main

import (
	"github.com/sunjun/videoapi/controllers"
	_ "github.com/sunjun/videoapi/docs"
	"github.com/sunjun/videoapi/models"
	_ "github.com/sunjun/videoapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	models.InitDB()
	controllers.InitServerQueue()
	controllers.InitClientQueue()
	beego.Run()
}
