package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	this.Data["Content"] = "value"
	this.Layout = "layout/layout.tpl"
	this.TplNames = "home/index.tpl"
}