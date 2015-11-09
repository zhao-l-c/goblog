package main

import (
	"GoWeb/beego/models"
	_ "GoWeb/beego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
    "strings"
    "GoWeb/beego/controllers"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	// 添加模板函数处理附件字符串
    beego.AddFuncMap("getFileName", getFileName)

    // 第一种设置附件访问的方法：直接设置静态路径为文件所在路径
    // 设置/attachment（第一个参数）为相对路径attachment（第二个参数）
    // beego.SetStaticPath("/attachment", "attachment")

    // 第二种方法：直接使用controller处理
    beego.Router("/attachment/:all", &controllers.AttachmentController{})

	beego.Run()
}

func getFileName(in string) string {
    index := strings.Index(in, "$")
    return in[index + 1:]
}
