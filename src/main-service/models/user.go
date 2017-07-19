package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

const (
	VIEWER = iota - 1 // незарегистрированный пользователь
	USER   = 0        // обычный пользователь
	LEADER = 1        // учитель, куратор, меет право создавать проекты
	ADMIN  = 2        // администратор, может всё
)

type User struct {
	Id          int    `orm:"column(id);pk;auto"  json:"id"`
	Nickname    string `orm:"column(nickname)"    json:"nickname"`
	Description string `orm:"column(description)" json:"description,omitempty"`
	Avatar      string `orm:"column(avatar)"      json:"avatar,omitempty"`
}

type FullUserInfo struct {
	Id          int            `json:"id"`
	Nickname    string         `json:"nickname"`
	Description string         `json:"description"`
	Avatar      string         `json:"avatar"`
	Contact     []*UserContact `json:"contacts,omitempty"`
}

type AllInformationAboutUser struct {
	User           FullUserInfo          `json:"user"`
	MemberOf       []MainProjectInfo     `json:"member_of,omitempty"`
	MasterOf       []MainProjectInfo     `json:"master_of,omitempty"`
	EnrolledOn     []MainProjectInfo     `json:"enrolled_on,omitempty"`
	MyApplications []ProjectApplications `json:"my_applications,omitempty"`
}

type MainUserInfo struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// func (t *User) TableName() string {
// 	return "user"
// }

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
			Id:          n.Id,
			Nickname:    n.Nickname,
			Description: n.Description,
			Avatar:      n.Avatar,
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
		_, err = o.Delete(&v)
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
