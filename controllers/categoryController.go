package controllers
import (
    "github.com/astaxie/beego"
    "GoWeb/beego/models"
)

type CategoryController struct {
    beego.Controller
}

func (this *CategoryController) Get() {
    this.Data["headerTitle"] = "Categories"
    this.TplNames = "category.html"
    this.Data["isCategory"] = true
    this.Data["isLogin"] = checkLogin(this.Ctx)

    op := this.Input().Get("op")
    if op == "add" {
        name := this.Input().Get("name")
        err := models.AddCategory(name)
        if err != nil {
            beego.Error(err)
        }
    } else if op == "del" {
        id := this.Input().Get("id")
        err := models.DeleteCategory(id)
        if err != nil {
            beego.Error(err)
        }
    }
    var err error
    this.Data["Categories"], err = models.AllCategories(false)
    if err != nil {
        beego.Error(err)
    }
}
