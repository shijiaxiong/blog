package controllers

import (
	"github.com/astaxie/beego"
)

type AdminHomeController struct {
	beego.Controller
}

func (this *AdminHomeController) Get() {

	if isSignIn, admin := CheckAccount(this.Ctx); isSignIn {

		this.Data["PageTitle"] = "Admin Home"
		this.Data["Admin"] = admin
		this.Data["IsHome"] = true
		this.Layout = "layout/admin_layout.tpl"
		this.TplNames = "admin_home/index.tpl"
	} else {

		this.Redirect("/adminlogin", 301)
		return
	}

}
