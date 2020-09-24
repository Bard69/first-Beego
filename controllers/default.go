package controllers

import (
	"BeegoWork/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1.获取name、age、sex
	//2.做数据对比
	//3.根据对比结果进行判断处理
	c.Data["Website"] = "https://www.baidu.com"
	c.Data["Email"] = "k2942318217@163.com"
	c.TplName = "index.tpl"
	//1.获取请求数据
	userName:=c.Ctx.Input.Query("user")
	password:=c.Ctx.Input.Query("psd")
	// 2.使用固定数据进行数据校验
	//admin 123456
	if userName!="admin"||password!="127657"{
		// 代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误"))
		return
}
	//校验正确的情况
	c.Ctx.ResponseWriter.Write([]byte("校验成功"))
}
//编写一个post方法，用于处理post类型的请求
//func (c *MainController) Post()  {
//
// for i:=0;i<10;i++{
//    fmt.Printf("第%d词打印\n")
// }
//}
//func (c *MainController) post()  {
////   接收(解析)Post请求的参数
// name:=c.Ctx.Request.FormValue("name")
// age:=c.Ctx.Request.FormValue("age")
// sex:=c.Ctx.Request.FormValue("sex")
// fmt.Println(name,age,sex)
////   2.进行数据校验
//if name !="admin"&&age!="18"{
// c.Ctx.WriteString("数据校验失败")
// return
//}
//c.Ctx.WriteString("数据校验成功")
//}
func (c *MainController) Post(){
	var person  models.Person
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil{
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	c.Ctx.WriteString("请求成功")
	fmt.Println("名字",person.Name)
	fmt.Println("年龄",person.Age)
	fmt.Println("性别",person.Sex)
}

