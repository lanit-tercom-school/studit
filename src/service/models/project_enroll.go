package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ProjectSignUp struct {
	Id          int      `orm:"column(id);pk;auto"`
	UserId      *User    `orm:"column(user_id);rel(fk)"`
	ProjectId   *Project `orm:"column(project_id);rel(fk)"`
}

func (t *ProjectSignUp) TableName() string {
	return "project_sign_up"
}

func init() {
	orm.RegisterModel(new(ProjectSignUp))
}

// AddProjectAuthor insert a new ProjectSignUp into database and returns
// last inserted Id on success.
func AddApplicationFromUserForProject(m *ProjectSignUp) (id int64, err error) {
	id, err = orm.NewOrm().Insert(m)
	return
}

// GetProjectAuthorById retrieves ProjectSignUp by Id. Returns error if
// Id doesn't exist
func GetProjectAuthorById(id int) (v *ProjectSignUp, err error) {
	o := orm.NewOrm()
	v = &ProjectSignUp{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllSignedUpOnProject(project_id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var singed_up []ProjectSignUp
	_, err = o.QueryTable(new(ProjectSignUp)).
			Filter("ProjectId", project_id).
			RelatedSel().
			All(&singed_up)
	if err != nil {
		return nil, err
	}
	for _, x := range singed_up {
		ml = append(ml, x.UserId.Id)
	}
	return ml, nil
}

// GetAllProjectAuthor retrieves all ProjectSignUp matches certain condition. Returns empty list if
// no records exist
func GetAllProjectAuthor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectSignUp))
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

	var l []ProjectSignUp
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

// UpdateProjectAuthor updates ProjectSignUp by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectAuthorById(m *ProjectSignUp) (err error) {
	o := orm.NewOrm()
	v := ProjectSignUp{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProjectSignUp deletes ProjectSignUp by Project Id and returns error if
// the record to be deleted doesn't exist
func DeleteProjectSignUp(user_id, project_id int) (err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable(new(ProjectSignUp)).
			Filter("UserId", user_id).
			Filter("ProjectId", project_id).
			Delete()
	/*v := ProjectSignUp{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProjectSignUp{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}*/
	return
}
