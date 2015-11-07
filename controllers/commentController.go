package controllers
import (
    "github.com/astaxie/beego"
    "GoWeb/beego/models"
)

type CommentController struct {
    beego.Controller
}

func (this *CommentController) Post() {
    // 评论不需要登录
    tid := this.Input().Get("tid")
    content := this.Input().Get("content")
    err := models.AddComment(tid, content)
    if err != nil {
        beego.Error(err)
    }
    this.Redirect("/topic/view/" + tid, 302)
}

func (this *CommentController) Delete() {
    tid := this.Ctx.Input.Param("0")
    id := this.Ctx.Input.Param("1")
    err := models.DeleteComment(id)
    if err != nil {
        beego.Error(err)
    }
    this.Redirect("/topic/view/" + tid, 302)
}