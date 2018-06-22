package main

import (
	_ "onRead/routers"
	"onRead/utils"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//初始化数据库
	utils.InitSql()

	beego.Run()
}
