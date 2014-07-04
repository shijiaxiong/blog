package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) View() {

	var err error
	this.Data["Topic"], err = models.GetOneTopicAndCategory(this.Input().Get("id"))

	if err != nil {
		beego.Error(err)
	}

	topic, _ := models.GetOneTopicAndCategory(this.Input().Get("id"))

	this.Data["CategoryTopics"], err = models.GetTopicByCategoryId(topic[0]["cid"], topic[0]["id"])

	if err != nil {
		beego.Error(err)
	}

	this.Data["Replys"], err = models.GetTopicReplyByTopicId(this.Input().Get("id"))

	if err != nil {
		beego.Error(err)
	}

	this.Data["PageTitle"] = "View Topic"
	this.Layout = "layout/layout.tpl"
	this.TplNames = "topic/view.tpl"
}

func (this *TopicController) Post() {

	err := models.AddReply(this.Input().Get("title"), this.Input().Get("content"), this.Input().Get("nickname"), this.Input().Get("tid"))

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view?id="+this.Input().Get("tid"), 301)
	return
}
