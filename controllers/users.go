package controllers

import (
	//"fmt"
	"encoding/json"
	"errors"
	"hello/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// UsersController operations for Users
type UsersController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *UsersController) Post() {
	id := c.Input().Get("Id")
	intid, err := strconv.Atoi(id)
	if err == nil {
		err = models.DeleteUsers(intid)
		if err == nil {
			c.Redirect("/v1/users", 302)
		}
	}
	c.TplName = "manage/user_manage.html"
	//c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsersController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsersById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["Name"] = v.Name
		c.Data["Sex"] = v.Sex
		c.Data["Age"] = v.Age
		c.Data["Email"] = v.Email
		c.Data["Password"] = v.Password
	}
	c.TplName = "myinfo.html"
	c.TplName = "change_info.html"
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (c *UsersController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// // fields: col1,col2,entity.col3
	// if v := c.GetString("fields"); v != "" {
	// 	fields = strings.Split(v, ",")
	// }
	// // limit: 10 (default is 10)
	// if v, err := c.GetInt64("limit"); err == nil {
	// 	limit = v
	// }
	// // offset: 0 (default is 0)
	// if v, err := c.GetInt64("offset"); err == nil {
	// 	offset = v
	// }
	// // sortby: col1,col2
	// if v := c.GetString("sortby"); v != "" {
	// 	sortby = strings.Split(v, ",")
	// }
	// // order: desc,asc
	// if v := c.GetString("order"); v != "" {
	// 	order = strings.Split(v, ",")
	// }
	// // query: k:v,k:v
	// if v := c.GetString("query"); v != "" {
	// 	for _, cond := range strings.Split(v, ",") {
	// 		kv := strings.SplitN(cond, ":", 2)
	// 		if len(kv) != 2 {
	// 			c.Data["json"] = errors.New("Error: invalid query key/value pair")
	// 			c.ServeJSON()
	// 			return
	// 		}
	// 		k, v := kv[0], kv[1]
	// 		query[k] = v
	// 	}
	// }

	// l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	// // if err != nil {
	// // 	c.Data["json"] = err.Error()
	// // } else {
	// // 	c.Data["json"] = l
	// // }
	// if err == nil {
	// 	c.Data["Name1"] = l[0].Name
	// 	c.Data["Name2"] = l[1].Name
	// 	c.Data["Name3"] = l[2].Name
	// 	c.Data["Name4"] = l[3].Name
	// 	c.Data["Name5"] = l[4].Name
	// 	c.Data["Name6"] = l[5].Name
	// 	c.TplName = "manage/user_manage.html"
	// }

	//c.ServeJSON()
	sess := c.StartSession()
	name := sess.Get("name")
	if name == nil || name == "" {
		c.Redirect("/v1/login", 302)
	} else {
		// fields: col1,col2,entity.col3
		if v := c.GetString("fields"); v != "" {
			fields = strings.Split(v, ",")
		}
		// limit: 10 (default is 10)
		if v, err := c.GetInt64("limit"); err == nil {
			limit = v
		}
		// offset: 0 (default is 0)
		if v, err := c.GetInt64("offset"); err == nil {
			offset = v
		}
		// sortby: col1,col2
		if v := c.GetString("sortby"); v != "" {
			sortby = strings.Split(v, ",")
		}
		// order: desc,asc
		if v := c.GetString("order"); v != "" {
			order = strings.Split(v, ",")
		}
		// query: k:v,k:v
		if v := c.GetString("query"); v != "" {
			for _, cond := range strings.Split(v, ",") {
				kv := strings.SplitN(cond, ":", 2)
				if len(kv) != 2 {
					c.Data["json"] = errors.New("Error: invalid query key/value pair")
					c.ServeJSON()
					return
				}
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}

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
		l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
		// if err != nil {
		// 	c.Data["json"] = err.Error()
		// } else {
		// 	c.Data["json"] = l
		// }
		if err == nil {
			// for i := 0; i < len(l); i++ {
			// 	c.Data[fmt.Sprintf("Name%d", i)] = l[i].Name
			// 	c.Data[fmt.Sprintf("Age%d", i)] = l[i].Age
			// 	c.Data[fmt.Sprintf("Sex%d", i)] = l[i].Sex
			// 	c.Data[fmt.Sprintf("Email%d", i)] = l[i].Email
			// 	c.Data[fmt.Sprintf("Password%d", i)] = l[i].Password
			// }

			c.Data["Users"] = l
			c.TplName = "manage/user_manage.html"
		}
	}
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsersController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Users{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		// if err := models.UpdateUsersById(&v); err == nil {
		// 	c.Data["json"] = "OK"
		// } else {
		// 	c.Data["json"] = err.Error()
		// }
	} else {
		c.Data["json"] = err.Error()
	}
	c.TplName = "manage/user_manage.html"
	//c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsers(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.TplName = "manage/user_manage.html"
	//c.ServeJSON()
}
