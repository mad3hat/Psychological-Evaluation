package controllers

import (
	//"encoding/json"
	"errors"
	"strings"
	//"errors"
	"hello/models"
	"strconv"
	//"strings"
	"fmt"

	"github.com/astaxie/beego"
)

// QuestionsController operations for Questions
type QuestionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *QuestionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Questions
// @Param	body		body 	models.Questions	true		"body for Questions content"
// @Success 201 {int} models.Questions
// @Failure 403 body is empty
// @router / [post]
func (c *QuestionsController) Post() {
	id := c.Input().Get("Id")
	intid, err := strconv.Atoi(id)
	if err == nil {
		err = models.DeleteQuestions(intid)
		if err == nil {
			c.Redirect("/v1/questions", 302)
		}
	}
	c.TplName = "manage/question_manage.html"
	//c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Questions by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Questions
// @Failure 403 :id is empty
// @router /:id [get]
func (c *QuestionsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetQuestionsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.TplName = "question.html"
	//c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Questions
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Questions
// @Failure 403
// @router / [get]
func (c *QuestionsController) GetAll() {
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

	// l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
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
		l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
		if err == nil {
			// for i := 0; i < len(l); i++ {
			// 	c.Data[fmt.Sprintf("Question%d", i)] = l[i].Body
			// 	c.Data[fmt.Sprintf("A%d", i)] = l[i].A
			// 	c.Data[fmt.Sprintf("B%d", i)] = l[i].B
			// 	c.Data[fmt.Sprintf("C%d", i)] = l[i].C
			// 	c.Data[fmt.Sprintf("D%d", i)] = l[i].D
			// }
			c.Data["Questions"] = l
			c.TplName = "manage/question_manage.html"
		}
	}
	//c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Questions
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Questions	true		"body for Questions content"
// @Success 200 {object} models.Questions
// @Failure 403 :id is not int
// @router /:id [put]
func (c *QuestionsController) Put() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.Atoi(idStr)
	// v := models.Questions{Id: id}
	// if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
	// 	if err := models.UpdateQuestionsById(&v); err == nil {
	// 		c.Data["json"] = "OK"
	// 	} else {
	// 		c.Data["json"] = err.Error()
	// 	}
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.TplName = "question.html"
	//c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Questions
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *QuestionsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteQuestions(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.TplName = "question.html"
	//c.ServeJSON()
}

type AddQuestionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *AddQuestionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (c *AddQuestionsController) Post() {
	inputs := c.Input()
	question := inputs.Get("question")
	A := inputs.Get("A")
	B := inputs.Get("B")
	C := inputs.Get("C")
	D := inputs.Get("D")
	err := models.AddQuestions(question, A, B, C, D)
	if err == nil {
		c.Redirect("/v1/questions", 302)
	} else {
		c.TplName = "error.html"
	}
}

// Get ...
// @router / [get]
func (c *AddQuestionsController) Get() {
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
	c.TplName = "manage/add_question.html"
}

// Questions1Controller ...
type Questions1Controller struct {
	beego.Controller
}

// URLMapping ...
func (c *Questions1Controller) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Post", c.Post)
}

// Post ...
// @router / [post]
func (c *Questions1Controller) Post() {
	//q1 :=
}

// GetAll ...
// @router / [get]
func (c *Questions1Controller) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64

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
		l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
		if err == nil {
			for i := 0; i < len(l); i++ {
				c.Data[fmt.Sprintf("Question%d", i)] = l[i].Body
				c.Data[fmt.Sprintf("A%d", i)] = l[i].A
				c.Data[fmt.Sprintf("B%d", i)] = l[i].B
				c.Data[fmt.Sprintf("C%d", i)] = l[i].C
				c.Data[fmt.Sprintf("D%d", i)] = l[i].D
			}
			//c.Data["Questions"] = l
			c.TplName = "paper/question1.html"
		}
	}
}

// Questions2Controller ...
type Questions2Controller struct {
	beego.Controller
}

// URLMapping ...
func (c *Questions2Controller) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Post", c.Post)
}

// GetAll ...
// @router / [get]
func (c *Questions2Controller) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64

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
		l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
		if err == nil {
			for i := 0; i < len(l); i++ {
				c.Data[fmt.Sprintf("Question%d", i)] = l[i].Body
				c.Data[fmt.Sprintf("A%d", i)] = l[i].A
				c.Data[fmt.Sprintf("B%d", i)] = l[i].B
				c.Data[fmt.Sprintf("C%d", i)] = l[i].C
				c.Data[fmt.Sprintf("D%d", i)] = l[i].D
			}
			//c.Data["Questions"] = l
			c.TplName = "paper/question2.html"
		}
	}
}

// Questions3Controller ...
type Questions3Controller struct {
	beego.Controller
}

// URLMapping ...
func (c *Questions3Controller) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Post", c.Post)
}

// GetAll ...
// @router / [get]
func (c *Questions3Controller) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64

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
		l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
		if err == nil {
			for i := 0; i < len(l); i++ {
				c.Data[fmt.Sprintf("Question%d", i)] = l[i].Body
				c.Data[fmt.Sprintf("A%d", i)] = l[i].A
				c.Data[fmt.Sprintf("B%d", i)] = l[i].B
				c.Data[fmt.Sprintf("C%d", i)] = l[i].C
				c.Data[fmt.Sprintf("D%d", i)] = l[i].D
			}
			//c.Data["Questions"] = l
			c.TplName = "paper/question3.html"
		}
	}
}

// Questions4Controller ...
type Questions4Controller struct {
	beego.Controller
}

// URLMapping ...
func (c *Questions4Controller) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Post", c.Post)
}

// GetAll ...
// @router / [get]
func (c *Questions4Controller) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64

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
		l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
		if err == nil {
			for i := 0; i < len(l); i++ {
				c.Data[fmt.Sprintf("Question%d", i)] = l[i].Body
				c.Data[fmt.Sprintf("A%d", i)] = l[i].A
				c.Data[fmt.Sprintf("B%d", i)] = l[i].B
				c.Data[fmt.Sprintf("C%d", i)] = l[i].C
				c.Data[fmt.Sprintf("D%d", i)] = l[i].D
			}
			//c.Data["Questuions"] = l
			c.TplName = "paper/question4.html"
		}
	}
}

// Questions5Controller ...
type Questions5Controller struct {
	beego.Controller
}

// URLMapping ...
func (c *Questions5Controller) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Post", c.Post)
}

// GetAll ...
// @router / [get]
func (c *Questions5Controller) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64

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
		l, err := models.GetAllQuestions(query, fields, sortby, order, offset, limit)
		if err == nil {
			for i := 0; i < len(l); i++ {
				c.Data[fmt.Sprintf("Question%d", i)] = l[i].Body
				c.Data[fmt.Sprintf("A%d", i)] = l[i].A
				c.Data[fmt.Sprintf("B%d", i)] = l[i].B
				c.Data[fmt.Sprintf("C%d", i)] = l[i].C
				c.Data[fmt.Sprintf("D%d", i)] = l[i].D
			}
			//c.Data["Questions"] = l
			c.TplName = "paper/question5.html"
		}
	}
}

// ChangeQuestionController ...
type ChangeQuestionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ChangeQuestionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @router / [post]
func (this *ChangeQuestionController) Post() {
	id := this.Input().Get("Id")
	body := this.Input().Get("Body")
	a := this.Input().Get("A")
	b := this.Input().Get("B")
	c := this.Input().Get("C")
	d := this.Input().Get("D")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetQuestionsById(intid)
	err := models.UpdateQuestionsById(v, body, a, b, c, d)
	if err == nil {
		this.Redirect("/v1/questions", 302)
	}
	this.TplName = "manage/change_question.html"
}

// Get ...
// @router / [get]
func (c *ChangeQuestionController) Get() {
	id := c.Input().Get("Id")
	intid, _ := strconv.Atoi(id)
	v, _ := models.GetQuestionsById(intid)
	c.Data["Body"] = v.Body
	c.Data["A"] = v.A
	c.Data["B"] = v.B
	c.Data["C"] = v.C
	c.Data["D"] = v.D
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
	c.TplName = "manage/change_question.html"
}
