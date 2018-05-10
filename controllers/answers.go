package controllers

import (
	"strconv"
	"fmt"
	"github.com/astaxie/beego"
	"hello/models"
	
)

// AnswersController ...
type AnswersController struct{
	beego.Controller
}

// URLMapping ...
func (c *AnswersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *AnswersController) Post() {
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

	sum := 0
	for i := 0; i < 10; i++ {
		a := c.Input().Get(fmt.Sprintf("ques%d", i))
		b, _ := strconv.Atoi(a)
		sum = sum + b
	}
	if sum < 20 {
		c.TplName = "paper/answer1.html"
	} else if sum < 40 {
		c.TplName = "paper/answer2.html"
	} else {
		c.TplName = "paper/answer3.html"
	}
}
