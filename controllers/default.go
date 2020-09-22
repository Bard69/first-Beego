package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https://www.baidu.com"
	c.Data["Email"] = "k2942318217@163.com"
	c.TplName = "index.tpl"
}
