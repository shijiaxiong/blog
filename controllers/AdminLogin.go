package controllers

import (
	"github.com/astaxie/beego"
)

type AdminLoginController struct {
	beego.Controller
}

func (this *AdminLoginController) Get() {

	this.Data["PageTitle"] = "Admin sign in"
	this.Layout = "layout/layout.tpl"
	this.TplNames = "admin_login/login.tpl"
}

func (this *AdminLoginController) Post() {

	admin := this.Input().Get("admin")
	password := this.Input().Get("password")

	if admin == beego.AppConfig.String("admin") &&
		password == beego.AppConfig.String("adminpwd") {

		this.Ctx.SetCookie("admin", admin, 3600, "/")
		this.Ctx.SetCookie("adminpwd", password, 3600, "/")

		this.Data["Time"] = 3
		this.Data["Message"] = "登录成功"
		this.Data["Url"] = "/"

		this.Layout = "layout/layout.tpl"
		this.TplNames = "layout/message.tpl"
		return
	} else {

		this.Data["Time"] = 3
		this.Data["Message"] = "用户名或密码错误"
		this.Data["Url"] = "/adminlogin"

		this.Layout = "layout/layout.tpl"
		this.TplNames = "layout/message.tpl"
		return
	}
}
