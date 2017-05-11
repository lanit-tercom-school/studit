package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ProjectEnroll struct {
	Id          int      `orm:"column(id);pk;auto"`
	UserId      *User    `orm:"column(user_id);rel(fk)"`
	ProjectId   *Project `orm:"column(project_id);rel(fk)"`
}

func (t *ProjectEnroll) TableName() string {
	return "project_enroll"
}

func init() {
	orm.RegisterModel(new(ProjectEnroll))
}

// GetProjectEnrollIdByUserId returns an array of projects
//where user enrolls
func GetProjectEnrollIdByUserId(userId int) (projectid []int64, err error){
	o := orm.NewOrm()
	var projects []ProjectEnroll
	_, err = o.QueryTable(new(ProjectEnroll)).Filter("UserId", User{Id: userId}).RelatedSel().All(&projects)
	if err != nil {
		return projectid, err
	}
	for _, v := range projects {
		projectid = append(projectid, v.ProjectId.Id)
	}
	return projectid, nil
}
// AddProjectAuthor insert a new ProjectEnroll into database and returns
// last inserted Id on success.
func AddApplicationFromUserForProject(u *User, p *ProjectJson) (id int64, err error) {
	temp := p.translate()
	m := ProjectEnroll{
		UserId: u,
		ProjectId: &temp,
	}
	id, err = orm.NewOrm().Insert(m)
	return
}

// GetProjectAuthorById retrieves ProjectEnroll by Id. Returns error if
// Id doesn't exist
func GetProjectAuthorById(id int) (v *ProjectEnroll, err error) {
	o := orm.NewOrm()
	v = &ProjectEnroll{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllSignedUpOnProject(project_id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var singed_up []ProjectEnroll
	_, err = o.QueryTable(new(ProjectEnroll)).
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

// GetAllProjectAuthor retrieves all ProjectEnroll matches certain condition. Returns empty list if
// no records exist
func GetAllProjectAuthor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectEnroll))
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

	var l []ProjectEnroll
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

// UpdateProjectAuthor updates ProjectEnroll by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectAuthorById(m *ProjectEnroll) (err error) {
	o := orm.NewOrm()
	v := ProjectEnroll{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProjectSignUp deletes ProjectEnroll by Project Id and returns error if
// the record to be deleted doesn't exist
func DeleteProjectSignUp(user_id, project_id int) (err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable(new(ProjectEnroll)).
			Filter("UserId", user_id).
			Filter("ProjectId", project_id).
			Delete()
	/*v := ProjectEnroll{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProjectEnroll{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}*/
	return
}
