package controllers

import (
	//"fmt"
	//"encoding/json"
	"errors"
	"hello/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// PostsController operations for Posts
type PostsController struct {
	beego.Controller
}

// URLMapping ...
func (c *PostsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Posts
// @Param	body		body 	models.Posts	true		"body for Posts content"
// @Success 201 {int} models.Posts
// @Failure 403 body is empty
// @router / [post]
func (c *PostsController) Post() {
	id := c.Input().Get("Id")
	intid, err := strconv.Atoi(id)
	if err == nil {
		err = models.DeletePosts(intid)
		if err == nil {
			c.Redirect("/v1/posts", 302)
		}
	}
	c.TplName = "manage/post_manage.html"
}

// GetOne ...
// @Title Get One
// @Description get Posts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Posts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PostsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetPostsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.TplName = "post.html"
	//c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Posts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Posts
// @Failure 403
// @router / [get]
func (c *PostsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
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

	// l, err := models.GetAllPosts(query, fields, sortby, order, offset, limit)
	// if err != nil {
	// 	c.Data["json"] = err.Error()
	// } else {
	// 	c.Data["json"] = l
	// }
	//c.TplName = "post.html"
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
		l, err := models.GetAllPosts(query, fields, sortby, order, offset, limit)
		if err == nil {
			c.Data["Posting"] = l
			c.TplName = "manage/post_manage.html"
		}
	}
}

// Put ...
// @Title Put
// @Description update the Posts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Posts	true		"body for Posts content"
// @Success 200 {object} models.Posts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PostsController) Put() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.Atoi(idStr)
	// v := models.Posts{Id: id}
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	// 	if err := models.UpdatePostsById(&v); err == nil {
	// 		c.Data["json"] = "OK"
	// 	} else {
	// 		c.Data["json"] = err.Error()
	// 	}
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.TplName = "post.html"
	//c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Posts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PostsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeletePosts(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.TplName = "post.html"
	//c.ServeJSON()
}

// PostsController operations for Posts
type PostingController struct {
	beego.Controller
}

// URLMapping ...
func (c *PostingController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *PostingController) Post() {
	inputs := c.Input()
	posting := inputs.Get("Posting")
	err := models.AddPosts(posting)
	if err == nil {
		c.Redirect("/v1/posting", 302)
	} else {
		c.TplName = "error.html"
	}
	c.TplName = "posting.html"
}

// Get ...
// @router / [get]
func (c *PostingController) Get() {
	sess := c.StartSession()
	name := sess.Get("name")
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64

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
		l, err := models.GetAllPosts(query, fields, sortby, order, offset, limit)
		if err == nil {
			c.Data["Post"] = l
		}
		c.TplName = "posting.html"
	}
}


// AddRespondController operations for Posts
type AddRespondController struct {
	beego.Controller
}

// URLMapping ...
func (c *AddRespondController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Post", c.Post)
}

// Get ...
// @router / [get]
func (c *AddRespondController) Get() {
	id := c.Input().Get("Id")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetPostsById(intid)
	c.Data["Body"] = v.Body
	c.Data["Id"] = v.Id
	c.Data["Respond"] = v.Respond

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
	c.TplName = "manage/add_respond.html"
}

// Post ...
// @router / [post]
func (c *AddRespondController) Post() {
	id := c.Input().Get("Id")
	respond := c.Input().Get("Respond")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetPostsById(intid)
	err := models.UpdatePostsById(v, respond)
	if err == nil {
		c.Redirect("/v1/posts", 302)
	}
	c.TplName = "manage/add_respond.html"
}
