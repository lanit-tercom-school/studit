package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ProjectUser struct {
	Id         int       `orm:"column(id);pk;auto"`
	Project    *Project  `orm:"column(project_id);rel(fk)"`
	User       *User     `orm:"column(user_id);rel(fk)"`
	SignedDate time.Time `orm:"column(signed_date);auto_now_add"`
	Progress   int       `orm:"column(progress)"`
}

func init() {
	orm.RegisterModel(new(ProjectUser))
}

// AddProjectUser insert a new ProjectUser into database and returns
// last inserted Id on success.
func AddProjectUser(m *ProjectUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProjectUserById retrieves ProjectUser by Id. Returns error if
// Id doesn't exist
func GetProjectUserById(id int) (v *ProjectUser, err error) {
	o := orm.NewOrm()
	v = &ProjectUser{Id: id}
	if err = o.QueryTable(v).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProjectUser retrieves all ProjectUser matches certain condition. Returns empty list if
// no records exist
func GetAllProjectUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectUser))
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

	var l []ProjectUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).RelatedSel().All(&l, fields...); err == nil {
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

// UpdateProjectUser updates ProjectUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectUserById(m *ProjectUser) (err error) {
	o := orm.NewOrm()
	v := ProjectUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteProjectUser deletes ProjectUser by Id and returns error if
//the recor to be deleted doesn't exist
func DeleteProjectUser(id int) (err error) {
	o := orm.NewOrm()
	v := ProjectUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProjectUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteProjectUserProjectId deletes ProjectUser by ProjectId and returns error if
// records to be deleted doesn't exist
func DeleteProjectUserProjectId(project_id int) (err error) {
	o := orm.NewOrm()
	project := Project{Id: project_id}
	v := ProjectUser{Project: &project}
	// ascertain id exists in the database
	if err = o.Read(&v, "Project"); err == nil {
		var num int64
		if num, err = o.Delete(&v, "Project"); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// DeleteProjectUserUserId deletes ProjectUser by UserId and returns error if
// records to be deleted doesn't exist
func DeleteProjectUserUserId(user_id int) (err error) {
	o := orm.NewOrm()
	user := User{Id: user_id}
	v := ProjectUser{User: &user}
	// ascertain id exists in the database
	if err = o.Read(&v, "User"); err == nil {
		var num int64
		if num, err = o.Delete(&v, "User"); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
