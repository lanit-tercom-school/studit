package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
)

type User struct {
    Id                  int     `orm:"column(id);pk;auto"                   json:"id"`
    Login               string  `orm:"column(login)"                        json:"login,omitempty"`
    Password            string  `orm:"column(password)"                     json:"-"`
    Nickname            string  `orm:"column(nickname)"                     json:"nickname"`
    Description         string  `orm:"column(description)"                  json:"description,omitempty"`
    Avatar              string  `orm:"column(avatar)"                       json:"avatar,omitempty"`
    // viewer - -1, registered user - 0, teacher - 1, admin 2, default is -1
    // Can't be higher than `auth.MaxPermissionLevel` !
    PermissionLevel     int     `orm:"column(permission_level);default(0)"  json:"permission_level,omitempty"`
}

type UserInfo struct {
	Id              int                	`json:"id"`
	Nickname        string                `json:"nickname"`
	Description     string                `json:"description"`
	Avatar          string                `json:"avatar"`
	PermissionLevel int               	 `json:"permission_level"`
}

func (t *User) TableName() string {
    return "user"
}

func init() {
    orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []User
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
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(n *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: n.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		//fields filter
		m := User{
			Id: n.Id,
			Login: v.Login,
			Password: v.Password,
			Nickname: n.Nickname,
			Description: n.Description,
			Avatar: n.Avatar,
			PermissionLevel: v.PermissionLevel,
		}
		_, err = o.Update(&m)
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

// return true if found, false if not
func (m *User) FindByLogin() bool {
	var anotherUser User
	err := orm.NewOrm().QueryTable("user").Filter("login", m.Login).One(&anotherUser)
	if err == orm.ErrMultiRows {
		panic(err)
	} else if err == orm.ErrNoRows {
		return false
	} else {
		return true
	}
}

// return true if found, false if not
func GetUserByLogin(login string) (*User, error) {
	var anotherUser User
	err := orm.NewOrm().QueryTable("user").Filter("login", login).One(&anotherUser)
	if err == orm.ErrMultiRows {
		panic(err)
	} else if err == orm.ErrNoRows {
		return nil, errors.New("Not found")
	} else {
		return &anotherUser, nil
	}
}
