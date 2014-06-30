package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type AdminCategoryController struct {
	beego.Controller
}

func (this *AdminCategoryController) Get() {

	if isSignIn, admin := CheckAccount(this.Ctx); isSignIn {

		op := this.Input().Get("op")

		switch op {

		case "add":

			title := this.Input().Get("title")

			if len(title) == 0 {
				break
			}

			err := models.AddCategory(title)

			if err != nil {

				beego.Error(err)
			}

			this.Redirect("./admincategory", 301)
			return

		case "del":

			id := this.Input().Get("id")

			if len(id) == 0 {
				break
			}

			err := models.DelCategory(id)

			if err != nil {
				beego.Error(err)
			}

			this.Redirect("./admincategory", 301)
			return
		}

		this.Data["PageTitle"] = "CategoryManager"
		this.Data["Admin"] = admin
		this.Data["IsCategory"] = true

		var err error
		this.Data["Categories"], err = models.GetAllCategories()

		if err != nil {

			beego.Error(err)
		}

		this.Layout = "layout/admin_layout.tpl"
		this.TplNames = "admin_category/index.tpl"
	} else {

		this.Redirect("/adminlogin", 301)
		return
	}
}
