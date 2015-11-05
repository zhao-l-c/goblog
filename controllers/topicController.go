package controllers

import (
    "github.com/astaxie/beego"
    "GoWeb/beego/models"
)

type TopicController struct {
    beego.Controller
}

func (this *TopicController) Get() {
    this.TplNames = "topic.html"
    this.Data["headerTitle"] = "Topics"
    this.Data["isLogin"] = checkLogin(this.Ctx)
    this.Data["isTopic"] = true
    // TODO
    // SetCommonData(this, "topic.html", "isTopic", "文章")
    var err error
    this.Data["Topics"], err = models.AllTopics(false)
    if err != nil {
        beego.Error(err)
    }
}

func (this *TopicController) Post() {
    if checkLogin(this.Ctx) {
        tid := this.Input().Get("id")
        title := this.Input().Get("title")
        content := this.Input().Get("content")
        cid := this.Input().Get("category")
        // inset one
        if len(tid) == 0 {
            err := models.AddTopic(title, content, cid)
            if err != nil {
                beego.Error(err)
            }
            // update one
        } else {
            err := models.UpdateTopic(tid, title, content, cid)
            if err != nil {
                beego.Error(err)
            }
        }
    }
    this.Redirect("/topic", 301)
}

func (this *TopicController) Add() {
    this.TplNames = "topicForm.html"
    this.Data["headerTitle"] = "Add Topic"
    this.Data["headerPrefix"] = "添加"
    var err error
    this.Data["Categories"], err = models.AllCategories(true)
    if err != nil {
        beego.Error(err)
    }
}

func (this *TopicController) Edit() {
    if checkLogin(this.Ctx) {
        id := this.Ctx.Input.Param("0")
        topic, err :=  models.GetTopic(id)
        if err != nil {
            beego.Error(err)
        }
        categories, err := models.AllCategories(true)
        if err != nil {
            beego.Error(err)
        }
        this.TplNames = "topicForm.html"
        this.Data["headerTitle"] = "Edit Topic"
        this.Data["headerPrefix"] = "修改"
        this.Data["Topic"] = topic
        this.Data["isTopic"] = true
        this.Data["isLogin"] = checkLogin(this.Ctx)
        this.Data["Categories"] = categories
    } else {
        this.Redirect("/login", 302)
    }
}

func (this *TopicController) View() {
    this.TplNames = "topicView.html"
    this.Data["headerTitle"] = "View Topics"
    this.Data["isLogin"] = checkLogin(this.Ctx)
    this.Data["isTopic"] = true

    id := this.Ctx.Input.Param("0")
    topic, err :=  models.GetTopic(id)
    if err != nil {
        beego.Error(err)
    }
    this.Data["Topic"] = topic
}

func (this *TopicController) Delete() {
    id := this.Ctx.Input.Param("0")
    err := models.DeleteTopic(id)
    if err != nil {
        beego.Error(err)
    }
    this.Redirect("/topic", 302)
}
