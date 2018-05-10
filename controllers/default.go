package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// URLMapping ...
func (c *MainController) URLMapping() {
	c.Mapping("*", c.Index)
}

// Index ...
// @router / [*]
func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
