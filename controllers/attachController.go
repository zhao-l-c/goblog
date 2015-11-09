package controllers
import (
    "github.com/astaxie/beego"
    "net/url"
    "os"
    "io"
)

type AttachmentController struct {
    beego.Controller
}

func (this *AttachmentController) Get() {
    // url即附件名称
    url, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1: ])
    if err != nil {
        this.Ctx.WriteString(err.Error())
        return
    }
    f, err := os.Open(url)
    if err != nil {
        this.Ctx.WriteString(err.Error())
        return
    }
    defer f.Close()

    _, err = io.Copy(this.Ctx.ResponseWriter, f)
    if err != nil {
        this.Ctx.WriteString(err.Error())
        return
    }
}
