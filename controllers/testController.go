package controllers

import (
	"github.com/astaxie/beego"
)

/**
测试模板的使用
*/
type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.TplNames = "test.tpl"

	// 条件判断
	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false

	type User struct {
		Name string
		Age  int
	}

	// 注意属性需要大写
	user := &User{
		Name: "Mick",
		Age:  20,
	}
	c.Data["user"] = user

	// 循环
	nums := []int{1, 2, 3, 4, 5, 6}
	c.Data["nums"] = nums

	// 模板变量
	tplVar := "xxxyyyzzz"
	c.Data["tplVar"] = tplVar

	// str2html
	html := "<h2>wondlerfuls</h2>"
	c.Data["html"] = html
}
