package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type AdminLoginController struct {
	beego.Controller
}

func (this *AdminLoginController) Get() {

	if this.Input().Get("exit") == "yes" {

		this.Ctx.SetCookie("admin", "", -1, "/")
		this.Ctx.SetCookie("adminpwd", "", -1, "/")

		this.Redirect("/adminlogin", 301)
		return
	}

	if isSingIn, _ := CheckAccount(this.Ctx); isSingIn {

		this.Redirect("/adminhome", 301)
		return
	}

	this.Data["PageTitle"] = "Admin sign in"
	this.Layout = "layout/layout.tpl"
	this.TplNames = "admin_login/login.tpl"
	return
}

func (this *AdminLoginController) Post() {

	admin := this.Input().Get("admin")
	password := this.Input().Get("password")

	if admin == beego.AppConfig.String("admin") &&
		password == beego.AppConfig.String("adminpwd") {

		this.Ctx.SetCookie("admin", admin, 0, "/")
		this.Ctx.SetCookie("adminpwd", password, 0, "/")

		this.Redirect("/adminhome", 301)
		return
	} else {

		this.Data["Msg"] = "用户名或密码错误"
		this.Layout = "layout/layout.tpl"
		this.TplNames = "admin_login/login.tpl"
		return
	}
}

func CheckAccount(ctx *context.Context) (bool, string) {

	ck, err := ctx.Request.Cookie("admin")

	if err != nil {
		return false, ""
	}

	admin := ck.Value

	ck, err = ctx.Request.Cookie("adminpwd")

	if err != nil {
		return false, ""
	}

	password := ck.Value

	return admin == beego.AppConfig.String("admin") &&
		password == beego.AppConfig.String("adminpwd"), admin
}
