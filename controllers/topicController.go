package controllers

import (
    "github.com/astaxie/beego"
    "GoWeb/beego/models"
    "path"
    "GoWeb/beego/commons"
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
    this.Data["Topics"], err = models.AllTopics(true, "", "")
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
        tags := this.Input().Get("tags")

        // 可忽略第三个返回值error，不强制上传附件
        _, fh, _ := this.GetFile("attachment")
        var attachName string
        // 是否上传了附件，更新文章的时候使用，若为true则需要删除旧的附件
        var attached = false
        if fh != nil {
            attached = true
            attachName = fh.Filename
            beego.Info("upload file is:" + attachName)
            // path.Join使用相对路径，即当前beego执行文件所在路径
            uuid, err := commons.UUID()
            if err != nil {
                beego.Error(err)
                this.Redirect("/topic", 302)
            }
            attachName = uuid + "$" + attachName
            this.SaveToFile("attachment", path.Join("attachment", attachName))
        } else {
            beego.Error("no attachment...")
        }
        // inset one
        if len(tid) == 0 {
            err := models.AddTopic(title, content, tags, cid, attachName)
            if err != nil {
                beego.Error(err)
            }
        // update one
        } else {
            err := models.UpdateTopic(tid, title, content, tags, cid, attachName, attached)
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
    comments, err := models.QueryCommentsByTid(id)
    if err != nil {
        beego.Error(err)
    }
    this.Data["Topic"] = topic
    this.Data["Comments"] = comments
    this.Data["Tags"], err = models.AllTags()
    if err != nil {
        beego.Error(err)
    }
}

func (this *TopicController) Delete() {
    id := this.Ctx.Input.Param("0")
    err := models.DeleteTopic(id)
    if err != nil {
        beego.Error(err)
    }
    this.Redirect("/topic", 302)
}

func (this *TopicController) Category() {
    category := this.Ctx.Input.Param("0")
    topics, err := models.AllTopics(true, category, "")
    if err != nil {
        beego.Error(err)
    }
    this.TplNames = "home.html"
    this.Data["headerTitle"] = "Home"
    this.Data["isLogin"] = checkLogin(this.Ctx)
    this.Data["isHome"] = true
    this.Data["Topics"] = topics
}

func (this *TopicController) Tag() {
    tag := this.Ctx.Input.Param("0")
    topics, err := models.AllTopics(true, "", tag);
    if err != nil {
        beego.Error(err)
    }
    this.TplNames = "home.html"
    this.Data["headerTitle"] = "Home"
    this.Data["isLogin"] = checkLogin(this.Ctx)
    this.Data["isHome"] = true
    this.Data["Topics"] = topics
    this.Data["Tags"], err = models.AllTags()
    if err != nil {
        beego.Error(err)
    }
}