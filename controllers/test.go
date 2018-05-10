package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

// URLMapping ...
func (c *TestController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
}

// Get ...
// @router / [get]
func (c *TestController) Get() {
	c.TplName = "test.html"
}

// Post ...
// @router / [post]
func (c *TestController) Post() {

}