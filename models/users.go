package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id        int       `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(255);null"`
	Age       string    `orm:"column(age);size(64);null"`
	Sex       string    `orm:"column(sex);size(64);null"`
	Email     string    `orm:"column(email);size(255)"`
	Password  string    `orm:"column(password);size(255)"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now"`
}

func (t *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(name string, email string, password string) (err error) {
	o := orm.NewOrm()
	//o.Using("default")
	user := new(Users)
	user.Name = name
	user.Email = email
	user.Password = password
	o.Insert(user)
	return nil
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsers retrieves all Users matches certain condition. Returns empty list if
// no records exist
func GetAllUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []Users, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))
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

	var l []Users
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

// UpdateUsers updates Users by ID and returns error if
// the record to be updated doesn't exist
func UpdateUsersByID(m Users, name string, sex string, age string, email string, password string) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}   //copy m to v
	// ascertain id exists in the database
	v.Name = name
	v.Sex = sex
	v.Age = age
	v.Email = email
	v.Password = password
	if _, err = o.Update(&v, "Name", "Age", "Sex", "Email", "Password"); err != nil {
		fmt.Println("更新失败！")
	}
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id int) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		//var num int64
		_, err = o.Delete(&Users{Id: id})
		if err != nil {
			fmt.Println("删除失败！")
		}
	}
	return
}

func FindUser(name string) (Users, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Name: name}
	err := o.Read(&user, "name")

	return user, err
}

func ValidateUser(user Users, password string) error {
	u, _ := FindUser(user.Name)
	if u.Name == "" {
		return errors.New("用户名不能为空！")
	}
	fmt.Println(u.Password)
	if u.Password != password {
		return errors.New("用户名或密码错误！")
	}
	return nil
}

func IfExist(name string) error {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Name: name}
	err := o.Read(&user, "name")
	if err != nil {
		return nil
	}
	return errors.New("用户名已存在！")
}
