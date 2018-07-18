package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id          int       `orm:"column(id);pk;auto"`
	Name        string    `orm:"column(name)"`
	Description string    `orm:"column(description)"`
	Created     time.Time `orm:"column(created);auto_now_add"`
	Logo        string    `orm:"column(logo)"`
	Tags        []string  `orm:"-"`
	Status      string    `orm:"column(status)"`
	GitHubUrl   string    `orm:"column(githuburl)"`
}

const (
	PROJECT_STATUS_OPENED  = "opened"
	PROJECT_STATUS_STARTED = "started"
	PROJECT_STATUS_ENDED   = "ended"
)

func init() {
	orm.RegisterModel(new(Project))
}

// AddProject insert a new Project into database and returns
// last inserted Id on success.
func AddProject(m *Project) (id int64, err error) {
	m.Status = "opened"
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProjectById retrieves Project by Id. Returns error if
// Id doesn't exist
func GetProjectById(id int) (v *Project, err error) {
	o := orm.NewOrm()
	v = &Project{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProject retrieves all Project matches certain condition. Returns empty list if
// no records exist
func GetAllProject(offset int64, limit int64) ([]Project, error) {
	list := []Project{}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Project))
	_, err := qs.Exclude("status", PROJECT_STATUS_ENDED).Offset(offset).Limit(limit).All(&list)
	if err != nil {
		beego.Trace(err.Error())
		return nil, err
	}
	return list, nil
}

// UpdateProject updates Project by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectById(m *Project) (err error) {
	o := orm.NewOrm()
	v := Project{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		v.Description = m.Description
		v.GitHubUrl = m.GitHubUrl
		v.Logo = m.Logo
		v.Name = m.Name
		v.Status = m.Status
		v.Tags = m.Tags
		if num, err = o.Update(&v); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProject deletes Project by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProject(id int) (err error) {
	o := orm.NewOrm()
	v := Project{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Project{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetProjectCount() (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Project))
	count, err = qs.Exclude("status", PROJECT_STATUS_ENDED).Count()
	return
}
