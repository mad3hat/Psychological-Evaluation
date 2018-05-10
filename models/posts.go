package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Posts struct {
	Id        int       `orm:"column(id);auto"`
	Body      string    `orm:"column(body);null"`
	Respond   string    `orm:"column(respond);null"`
	UserId    int       `orm:"column(user_id);null"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now"`
}

func (t *Posts) TableName() string {
	return "posts"
}

func init() {
	orm.RegisterModel(new(Posts))
}

// AddPosts insert a new Posts into database and returns
// last inserted Id on success.
func AddPosts(m string) (err error) {
	o := orm.NewOrm()
	//o.Using("default")
	posting := new(Posts)
	posting.Body = m
	o.Insert(posting)
	return nil
}

// GetPostsById retrieves Posts by Id. Returns error if
// Id doesn't exist
func GetPostsById(id int) (v *Posts, err error) {
	o := orm.NewOrm()
	v = &Posts{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPosts retrieves all Posts matches certain condition. Returns empty list if
// no records exist
func GetAllPosts(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []Posts, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Posts))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Posts
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				//ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePosts updates Posts by Id and returns error if
// the record to be updated doesn't exist
func UpdatePostsById(m *Posts, res string) (err error) {
	o := orm.NewOrm()
	v := Posts{Id: m.Id}   //copy m to v
	// ascertain id exists in the database
	v.Respond = res
	if _, err = o.Update(&v, "Respond"); err != nil {
		fmt.Println("更新失败！")
	}
	return
}

// DeletePosts deletes Posts by Id and returns error if
// the record to be deleted doesn't exist
func DeletePosts(id int) (err error) {
	o := orm.NewOrm()
	v := Posts{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Posts{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
