package main

import (
	_ "BeegoWork/db_mysql"
	_ "BeegoWork/routers"
	"github.com/astaxie/beego"

)

func main() {
	beego.Run()
}

