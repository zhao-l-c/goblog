package controllers
import "github.com/astaxie/beego"


type LogoutController struct {
    beego.Controller
}

// goto login.html and clear cookies
func (this * LogoutController) Get() {
    this.Ctx.SetCookie("username", "", -1, "/")
    this.Ctx.SetCookie("password", "",-1, "/")
    this.Data["headerTitle"] = "Login"
    this.TplNames = "login.html"
}