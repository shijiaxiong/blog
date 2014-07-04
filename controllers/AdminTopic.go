package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type AdminTopicController struct {
	beego.Controller
}

func (this *AdminTopicController) Get() {

	if isSignIn, admin := CheckAccount(this.Ctx); isSignIn {

		var err error
		this.Data["Topics"], err = models.GetAllTopic()

		if err != nil {
			beego.Error(err)
		}

		this.Data["IsTopic"] = true
		this.Data["PageTitle"] = "Topic Manager"
		this.Data["Admin"] = admin
		this.Layout = "layout/admin_layout.tpl"
		this.TplNames = "admin_topic/index.tpl"
	} else {
		this.Redirect("/adminlogin", 301)
		return
	}

}

func (this *AdminTopicController) Post() {

	op := this.Input().Get("op")

	switch op {

	case "add":

		title := this.Input().Get("title")
		content := this.Input().Get("content")
		category := this.Input().Get("category")

		if len(title) == 0 {
			break
		}

		if len(content) == 0 {
			break
		}

		if len(category) == 0 {
			break
		}

		err := models.AddTopic(title, content, category)

		if err != nil {
			beego.Error(err)
		}

		this.Redirect("./admintopic", 301)
		return

	case "save":

		id := this.Input().Get("id")
		if len(id) == 0 {
			this.Redirect("./admintopic", 301)
			return
		}

		title := this.Input().Get("title")
		content := this.Input().Get("content")
		category := this.Input().Get("category")

		err := models.UpdateTopic(id, title, content, category)

		if err != nil {
			beego.Error(err)
		}

		this.Redirect("./admintopic", 301)
		return
	}
}

func (this *AdminTopicController) Add() {

	if isSignIn, admin := CheckAccount(this.Ctx); isSignIn {
		var err error
		this.Data["Categories"], err = models.GetAllCategories()

		if err != nil {

			beego.Error(err)
		}

		this.Data["Admin"] = admin
		this.Data["IsTopic"] = true
		this.Data["PageTitle"] = "Add Topic"
		this.Layout = "layout/admin_layout.tpl"
		this.TplNames = "admin_topic/add.tpl"
	} else {
		this.Redirect("/adminhome", 301)
	}

}

func (this *AdminTopicController) Edit() {

	id := this.Input().Get("id")

	if len(id) == 0 {
		return
	}

	var err error
	this.Data["Topic"], err = models.GetOneTopic(id)

	if err != nil {
		beego.Error(err)
	}

	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}

	this.Data["PageTitle"] = "Eidt Topic"
	this.Data["IsTopic"] = true
	this.Layout = "layout/admin_layout.tpl"
	this.TplNames = "admin_topic/edit.tpl"
}
