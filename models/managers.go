package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Managers struct {
	Id        int       `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(255);null"`
	Email     string    `orm:"column(email);size(255)"`
	Password  string    `orm:"column(password);size(255)"`
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now"`
}

func (t *Managers) TableName() string {
	return "managers"
}

func init() {
	orm.RegisterModel(new(Managers))
}

// AddManagers insert a new Managers into database and returns
// last inserted Id on success.
func AddManagers(name string, email string, password string) (err error) {
	o := orm.NewOrm()
	//o.Using("default")
	manager := new(Managers)
	manager.Name = name
	manager.Email = email
	manager.Password = password
	o.Insert(manager)
	return nil
}

// GetManagersById retrieves Managers by Id. Returns error if
// Id doesn't exist
func GetManagersById(id int) (v *Managers, err error) {
	o := orm.NewOrm()
	v = &Managers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllManagers retrieves all Managers matches certain condition. Returns empty list if
// no records exist
func GetAllManagers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []Managers, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Managers))
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

	var l []Managers
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

// UpdateManagers updates Managers by Id and returns error if
// the record to be updated doesn't exist
func UpdateManagersById(m *Managers) (err error) {
	o := orm.NewOrm()
	v := Managers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteManagers deletes Managers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteManagers(id int) (err error) {
	o := orm.NewOrm()
	v := Managers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Managers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func FindManager(name string) (Managers, error) {
	o := orm.NewOrm()
	o.Using("default")
	manager := Managers{Name: name}
	err := o.Read(&manager, "name")

	return manager, err
}

func ValidateManager(manager Managers, password string) error {
	u, _ := FindManager(manager.Name)
	if u.Name == "" {
		return errors.New("用户名不能为空！")
	}
	fmt.Println(u.Password)
	if u.Password != password {
		return errors.New("用户名或密码错误！")
	}
	return nil
}
