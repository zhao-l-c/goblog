package controllers
import (
    "github.com/astaxie/beego"
)

/**
失败
*/
func SetCommonData(con beego.Controller, tplName, actNavBar, headerTitle string) {
    con.TplNames = tplName
    con.Data["isLogin"] = checkLogin(con.Ctx)
    con.Data[actNavBar] = true
    con.Data["headerTitle"] = headerTitle
}
