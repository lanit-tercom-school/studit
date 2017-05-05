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
	ProjectId  *Project  `orm:"column(project_id);rel(fk)"`
	UserId     *User     `orm:"column(user_id);rel(fk)"`
	SignedDate time.Time `orm:"column(signed_date);type(datetime)"`
	Progress   int       `orm:"column(progress)"`
}

func (t *ProjectUser) TableName() string {
	return "project_user"
}

func init() {
	orm.RegisterModel(new(ProjectUser))
}


// AddProjectUser insert a new ProjectUser into database and returns
// last inserted Id on success.
func AddUserToProject(m *ProjectUser) (err error) {
	o := orm.NewOrm()
	var p ProjectEnroll
	_, err = o.QueryTable(new(ProjectEnroll)).
		Filter("user_id", m.UserId).
		Filter("project_id", m.ProjectId).
		RelatedSel().
		All(&p)
	if err != nil {
		return err
	}
	_, err = o.QueryTable(new(ProjectEnroll)).Filter("id", p.Id).Delete()
	if err != nil {
		return err
	}
	_, err = o.Insert(m)
	return
}

func AddProjectUser(m *ProjectUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetUsersByProjectId(project_id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var users []ProjectUser
	_, err = o.QueryTable(new(ProjectUser)).
		Filter("project_id", project_id).
		RelatedSel().
		All(&users)
	if err != nil {
		return nil, err
	}
	for _, x := range users {
		ml = append(ml, x.UserId.Id)
	}
	return ml, nil
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

	var l []ProjectUser
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

// DeleteProjectUser deletes ProjectUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserFromProject(user_id int, project_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(ProjectUser)).
		Filter("UserId", user_id).
		Filter("ProjectId", project_id).
		Delete()
	return
}
