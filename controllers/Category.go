package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {

	var err error
	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}

	this.Data["Topics"], err = models.GetTopicByCategoryId(this.Input().Get("id"), "")

	if err != nil {
		beego.Error(err)
	}
	this.Data["PageTitle"] = "Category List"
	this.Layout = "layout/layout.tpl"
	this.TplNames = "category/index.tpl"
}
