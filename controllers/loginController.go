package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
    this.Data["headerTitle"] = "Login"
    this.TplNames = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autologin := this.Input().Get("autologin") == "on"

	// valify
	if username == "admin" && password == "admin" {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
		this.Redirect("/", 301)
	} else {
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")
		this.Redirect("/login", 301)
	}
	return
}

/**
check username and password
*/
func checkLogin(ctx *context.Context) bool {
	username, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	password, err := ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	return username.Value == "admin" && password.Value == "admin"
}
