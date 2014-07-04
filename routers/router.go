package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/adminlogin", &controllers.AdminLoginController{})
	beego.Router("/adminhome", &controllers.AdminHomeController{})
	beego.Router("/admincategory", &controllers.AdminCategoryController{})
	beego.Router("/admintopic", &controllers.AdminTopicController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.AdminTopicController{})
}
