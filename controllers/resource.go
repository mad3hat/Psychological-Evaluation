package controllers

import (
	//"fmt"
	//"encoding/json"
	"errors"
	"strings"
	//"errors"
	"hello/models"
	"strconv"
	//"strings"

	"github.com/astaxie/beego"
)

// ResourceController operations for Resource
type ResourceController struct {
	beego.Controller
}

// URLMapping ...
func (c *ResourceController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Resource
// @Param	body		body 	models.Resource	true		"body for Resource content"
// @Success 201 {int} models.Resource
// @Failure 403 body is empty
// @router / [post]
func (c *ResourceController) Post() {
	id := c.Input().Get("Id")
	intid, err := strconv.Atoi(id)
	if err == nil {
		err = models.DeleteResource(intid)
		if err == nil {
			c.Redirect("/v1/resources", 302)
		}
	}
	c.TplName = "manage/resource_manage.html"
	//c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Resource by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Resource
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ResourceController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetResourceById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.TplName = "/manage/resource_manage.html"
	//c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Resource
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Resource
// @Failure 403
// @router / [get]
func (c *ResourceController) GetAll() {
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

	// l, err := models.GetAllResource(query, fields, sortby, order, offset, limit)
	// if err != nil {
	// 	c.Data["json"] = err.Error()
	// } else {
	// 	c.Data["json"] = l
	// }

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
		l, err := models.GetAllResource(query, fields, sortby, order, offset, limit)
		// if err != nil {
		// 	c.Data["json"] = err.Error()
		// } else {
		// 	c.Data["json"] = l
		// }
		if err == nil {
			// for i := 0; i < len(l); i++ {
			// 	c.Data[fmt.Sprintf("Link%d", i)] = l[i].Link 
			// 	c.Data[fmt.Sprintf("Url%d", i)] = l[i].Url
			// }
			c.Data["Resource"] = l
			c.TplName = "manage/resource_manage.html"
		}
	}
	//c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Resource
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Resource	true		"body for Resource content"
// @Success 200 {object} models.Resource
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ResourceController) Put() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.Atoi(idStr)
	// v := models.Resource{Id: id}
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	// 	if err := models.UpdateResourceById(&v); err == nil {
	// 		c.Data["json"] = "OK"
	// 	} else {
	// 		c.Data["json"] = err.Error()
	// 	}
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.TplName = "/manage/resource_manage.html"
	//c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Resource
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ResourceController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteResource(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.TplName = "/manage/resource_manage.html"
	//c.ServeJSON()
}

// AddResourceController operations for Resource
type AddResourceController struct {
	beego.Controller
}

// URLMapping ...
func (c *AddResourceController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *AddResourceController) Post() {
	inputs := c.Input()
	link := inputs.Get("link")
	url := inputs.Get("url")
	err := models.AddResource(link, url)
	if err == nil {
		c.Redirect("/v1/resource", 302)
	} else {
		c.TplName = "error.html"
	}
}

// Get ...
// @router / [get]
func (c *AddResourceController) Get() {
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
	c.TplName = "manage/add_resource.html"
}


// ChangeResourceController ...
type ChangeResourceController struct {
	beego.Controller
}

// URLMapping ...
func (c *ChangeResourceController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *ChangeResourceController) Post() {
	id := c.Input().Get("Id")
	link := c.Input().Get("Link")
	url := c.Input().Get("Url")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetResourceById(intid)
	err := models.UpdateResourceById(v, link, url)
	if err == nil {
		c.Redirect("/v1/resources", 302)
	}
	c.TplName = "manage/change_resource.html"
}

// Get ...
// @router / [get]
func (c *ChangeResourceController) Get() {
	id := c.Input().Get("Id")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetResourceById(intid)
	c.Data["Link"] = v.Link
	c.Data["Url"] = v.Url
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
	c.TplName = "manage/change_resource.html"
}