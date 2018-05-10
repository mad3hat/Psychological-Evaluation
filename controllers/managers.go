package controllers

import (
	"hello/models"

	"github.com/astaxie/beego"
)

// ManagersController operations for Managers
type ManagersController struct {
	beego.Controller
}

// URLMapping ...
func (c *ManagersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *ManagersController) Post() {
	inputs := c.Input()
	name := inputs.Get("name")
	password := inputs.Get("password")
	email := inputs.Get("email")
	err := models.AddManagers(name, email, password)
	if err == nil {
		c.Redirect("/v1/users", 302)
	} else {
		c.TplName = "error.html"
	}
}

// Get ...
// @router / [get]
func (c *ManagersController) Get() {
	u1, _ := models.GetResourceById(1)
	c.Data["link1"] = u1.Link
	c.Data["url1"] = u1.Url
	u2, _ := models.GetResourceById(2)
	c.Data["link2"] = u2.Link
	c.Data["url2"] = u2.Url
	u3, _ := models.GetResourceById(3)
	c.Data["link3"] = u3.Link
	c.Data["url3"] = u3.Url
	u4, _ := models.GetResourceById(4)
	c.Data["link4"] = u4.Link
	c.Data["url4"] = u4.Url
	u5, _ := models.GetResourceById(5)
	c.Data["link5"] = u5.Link
	c.Data["url5"] = u5.Url
	u6, _ := models.GetResourceById(6)
	c.Data["link6"] = u6.Link
	c.Data["url6"] = u6.Url
	c.TplName = "manage/add_manager.html"
}
