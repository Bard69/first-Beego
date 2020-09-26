package routers

import (
	"BeegoWork/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.RegisterController{})
}
