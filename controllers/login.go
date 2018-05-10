package controllers

import (
	"fmt"
	"hello/models"
	"strconv"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title Post
// @Description create Users
// @Param	name  password 	models.Users	true		"password for Users name"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {
	var user models.Users
	var manager models.Managers
	inputs := c.Input()
	user.Name = inputs.Get("name")
	manager.Name = inputs.Get("name")
	password := inputs.Get("password")
	err1 := models.ValidateUser(user, password)
	err2 := models.ValidateManager(manager, password)
	if err1 == nil {
		u, _ := models.FindUser(user.Name)
		c.SetSession("id", fmt.Sprintf("%d", u.Id))
		c.SetSession("name", fmt.Sprintf("%s", u.Name))
		c.Redirect("/v1/index", 302)
	} else {
		if err2 == nil {
			m, _ := models.FindManager(manager.Name)
			c.SetSession("id", fmt.Sprintf("%d", m.Id))
			c.SetSession("name", fmt.Sprintf("%s", m.Name))
			c.Redirect("/v1/users", 302)
		}
	}
}

// Get ...
// @router / [get]
func (c *LoginController) Get() {
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
	c.TplName = "login.html"
}

type LogoutController struct {
	beego.Controller
}

// URLMapping ...
func (c *LogoutController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get ...
// @router /:id [get]
func (c *LogoutController) Get() {
	c.DelSession("id")
	c.DelSession("name")
	c.Redirect("/v1/login", 302)
}

type SignupController struct {
	beego.Controller
}

// URLMapping ...
func (c *SignupController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
}

// Get ...
// @router / [get]
func (c *SignupController) Get() {
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
	c.TplName = "signup.html"
}

// Post ...
// @router / [post]
func (c *SignupController) Post() {
	inputs := c.Input()
	name := inputs.Get("name")
	Pwd := inputs.Get("password")
	email := inputs.Get("email")
	err := models.AddUsers(name, email, Pwd)
	if err == nil {
		c.Redirect("/v1/login", 302)
	} else {
		c.TplName = "error.html"
	}
}

// MyinfoController operations for Users
type MyinfoController struct {
	beego.Controller
}

// URLMapping ...
func (c *MyinfoController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get ...
// @router / [get]
func (c *MyinfoController) Get() {
	sess := c.StartSession()
	name := sess.Get("name")
	v, _ := models.FindUser(name.(string))
	c.Data["Name"] = v.Name
	c.Data["Sex"] = v.Sex
	c.Data["Age"] = v.Age
	c.Data["Email"] = v.Email

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
	c.TplName = "myinfo.html"
}

// ChangeinfoController operations for Users
type ChangeinfoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ChangeinfoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *ChangeinfoController) Post() {
	sess := c.StartSession()
	nam := sess.Get("name")
	v, _ := models.FindUser(nam.(string))
	inputs := c.Input()
	name := inputs.Get("Name")
	sex := inputs.Get("Sex")
	age := inputs.Get("Age")
	email := inputs.Get("Email")
	password := inputs.Get("Password")
	err := models.UpdateUsersByID(v, name, sex, age, email, password)
	if err != nil {
		fmt.Println("更新失败！")
	} else {
		c.Redirect("/v1/index", 302)
	}
	c.TplName = "change_info.html"
}

// Get ...
// @router / [get]
func (c *ChangeinfoController) Get() {
	sess := c.StartSession()
	name := sess.Get("name")
	v, _ := models.FindUser(name.(string))
	c.Data["Name"] = v.Name
	c.Data["Sex"] = v.Sex
	c.Data["Age"] = v.Age
	c.Data["Email"] = v.Email
	c.Data["Password"] = v.Password

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
	c.TplName = "change_info.html"
}


// ChangePassController operations for Users
type ChangePassController struct {
	beego.Controller
}

// URLMapping ...
func (c *ChangePassController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *ChangePassController) Post() {
	id := c.Input().Get("Id")
	pass := c.Input().Get("Password")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetUsersById(intid)
	err := models.UpdateUsersByID(*v, v.Name, v.Sex, v.Age, v.Email, pass)
	if err == nil {
		c.Redirect("/v1/users", 302)
	}
	c.TplName = "manage/change_pass.html"
}

// Get ...
// @router / [get]
func (c *ChangePassController) Get() {
	id := c.Input().Get("Id")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetUsersById(intid)
	c.Data["Name"] = v.Name
	c.Data["Id"] = v.Id

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
	c.TplName = "manage/change_pass.html"
}