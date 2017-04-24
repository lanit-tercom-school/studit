package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ProjectMaster struct {
	Id         int       `orm:"column(id);pk;auto"`
	ProjectId  *Project  `orm:"column(project_id);rel(fk)"`
	MasterId   *User     `orm:"column(master_id);rel(fk)"`
	SignedDate time.Time `orm:"column(signed_date);type(datetime)"`
}

func (t *ProjectMaster) TableName() string {
	return "project_user"
}

func init() {
	orm.RegisterModel(new(ProjectMaster))
}

// AddProjectUser insert a new ProjectMaster into database and returns
// last inserted Id on success.
func AddMasterToProject(m *ProjectMaster) (err error) {
	_, err = orm.NewOrm().Insert(m)
	return
}

func IsUserInArray(user_id int, users []*User) bool {
	for _, x := range users {
		if x.Id == user_id{
			return true
		}
	}
	return false
}

func GetMastersOfTheProject(project_id int) (masters []*User, err error) {
	o := orm.NewOrm()
	var connections []ProjectMaster
	// выбираем всех пользователей, являющихся мастерами данного проекта
	_, err = o.QueryTable(new(ProjectMaster)).
			Filter("project_id", project_id).
			RelatedSel().
			All(&connections, "MasterId")
	// возвращаем только мастеров
	for _, x := range connections {
		masters = append(masters, x.MasterId)
	}
}

func AddProjectUser(m *ProjectMaster) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetUsersByProjectId(project_id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var users []ProjectMaster
	_, err = o.QueryTable(new(ProjectMaster)).
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


// GetAllProjectUser retrieves all ProjectMaster matches certain condition. Returns empty list if
// no records exist
func GetAllProjectUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectMaster))
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

	var l []ProjectMaster
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

// UpdateProjectUser updates ProjectMaster by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectUserById(m *ProjectMaster) (err error) {
	o := orm.NewOrm()
	v := ProjectMaster{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProjectUser deletes ProjectMaster by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserFromProject(user_id int, project_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(ProjectMaster)).
		Filter("UserId", user_id).
		Filter("ProjectId", project_id).
		Delete()
	return
}
