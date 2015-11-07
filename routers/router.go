package routers

import (
	"GoWeb/beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// modole test
	beego.Router("/test", &controllers.TestController{})
	// home
	beego.Router("/", &controllers.MainController{})
	// login
	beego.Router("/login", &controllers.LoginController{})
	// logout
    beego.Router("/logout", &controllers.LogoutController{})
    //
    beego.Router("/category", &controllers.CategoryController{})
    //
    beego.Router("/topic", &controllers.TopicController{})
    //
    beego.Router("/comment", &controllers.CommentController{})
    // 使用自动路由
    beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.CommentController{})
}
