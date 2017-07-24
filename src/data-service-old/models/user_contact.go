package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego"
)

type UserContact struct {
	Id      int          `orm:"column(id);pk;auto"              json:"id,omitempty"`
	Contact string       `orm:"column(contact)"                 json:"value"`
	Type    *ContactType `orm:"column(contact_type_id);rel(fk)" json:"type"`
	UserId  *User        `orm:"column(user_id);rel(fk)"         json:"-"`
}

//func (t *UserContact) TableName() string {
//	return "user_contacts"
//}

func init() {
	orm.RegisterModel(new(UserContact))
}

// GetUserContactById retrieves UserContact by Id. Returns error if
// Id doesn't exist
func GetUserContactById(id int) (v *UserContact, err error) {
	o := orm.NewOrm()
	v = &UserContact{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserContact retrieves all UserContact matches certain condition. Returns empty list if
// no records exist
func GetAllUserContact(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserContact))
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

	var l []UserContact
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

// UpdateUserContact updates UserContact by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserContactById(m *UserContact) (err error) {
	o := orm.NewOrm()
	v := UserContact{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserContact deletes UserContact by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserContact(id int) (err error) {
	o := orm.NewOrm()
	v := UserContact{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		o.Delete(&v)
	}
	return
}

func GetAllUserContacts(userId int) (ml []*UserContact, err error) {
	o := orm.NewOrm()
	var contacts []UserContact
	_, err = o.QueryTable(new(UserContact)).Filter("UserId", User{Id: userId}).RelatedSel().All(&contacts)
	if err != nil {
		return ml, err
	}
	for _, v := range contacts {
		ml = append(ml, &v)
	}
	return ml, nil
}
