package controllers

import (
	"github.com/astaxie/beego"
    "GoWeb/beego/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplNames = "home.html"
	this.Data["headerTitle"] = "Home"
	this.Data["isLogin"] = checkLogin(this.Ctx)
    this.Data["isHome"] = true
    var err error

	this.Data["Topics"], err = models.AllTopics(true, "", "")
    if err != nil {
        beego.Error(err)
    }
    this.Data["Tags"], err = models.AllTags()
    if err != nil {
        beego.Error(err)
    }
}
