package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	var err error
	this.Data["Topics"], err = models.GetAllTopic()

	if err != nil {
		beego.Error(err)
	}

	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
	this.Data["PageTitle"] = "Home"
	this.Layout = "layout/layout.tpl"
	this.TplNames = "home/index.tpl"
}
