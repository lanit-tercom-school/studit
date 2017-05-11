package models

import (
	"errors"
	"fmt"
	//"reflect"
	"strings"
	"github.com/astaxie/beego/orm"
	"time"
)

// Модель для базы данных
type Project struct {
	Id          int64  `orm:"column(id);pk;auto"`
	Name        string `orm:"column(name)"`
	Description string `orm:"column(description)"`
	DateOfCreation time.Time   `orm:"column(date_of_creation);type(datetime)"`
	Logo        string `orm:"column(logo)"`
	Tags        string `orm:"column(tags)"`
	Status 	    string `orm:"column(status)"`
}

// Модель для общения с клиентами
type ProjectJson struct {
	Id              int64       `json:"id,omitempty"` //
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	DateOfCreation  time.Time   `json:"created"`
	Logo            string      `json:"logo"`
	Tags            []string    `json:"tags"`
	Status          string      `json:"status"`
}

// Обязательный превод от одной модели в другую
func (t *Project) translate()  ProjectJson{
	return ProjectJson{
		Id: t.Id,
		Name: t.Name,
		Description: t.Description,
		DateOfCreation: t.DateOfCreation,
		Logo: t.Logo,
		Tags: strings.Split(t.Tags, ","),
		Status: t.Status,
	}
}

func (t *ProjectJson) translate()  Project{
	return Project{
		Id: t.Id,
		Name: t.Name,
		Description: t.Description,
		DateOfCreation: t.DateOfCreation,
		Logo: t.Logo,
		Tags: strings.Join(t.Tags, ","),
		Status: t.Status,
	}
}


func (t *Project) TableName() string {
	return "project"
}

func init() {
	orm.RegisterModel(new(Project))
}

// AddProject insert a new Project into database and returns
// last inserted Id on success.
func AddProject(m *ProjectJson) (id int64, err error) {
	v := m.translate()
	v.Id = 0 // for auto inc
	v.DateOfCreation = time.Now()
	o := orm.NewOrm()
	id, err = o.Insert(&v)
	return
}

// GetProjectById retrieves Project by Id. Returns error if
// Id doesn't exist
func GetProjectById(id int64) (*ProjectJson, error) {
	o := orm.NewOrm()
	temp := &Project{Id: id}
	if err := o.Read(temp); err == nil {
		v := temp.translate()
		return &v, nil
	} else {
		return nil, err
	}
}

// GetAllProjects retrieves all Project matches certain condition. Returns empty list if
// no records exist
func GetAllProjects(query map[string]string, sortBy []string, order []string,
	offset int64, limit int64, tag string, user int64, master int64, status string) (ml []ProjectJson, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Project))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortBy) != 0 {
		if len(sortBy) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortBy {
				orderBy := ""
				if order[i] == "desc" {
					orderBy = "-" + v
				} else if order[i] == "asc" {
					orderBy = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderBy)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortBy) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortBy {
				orderBy := ""
				if order[0] == "desc" {
					orderBy = "-" + v
				} else if order[0] == "asc" {
					orderBy = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderBy)
			}
		} else if len(sortBy) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Project
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v.translate())
		}
		if tag != "" {
			FilterByTag(&ml, tag)
		}
		if user != 0 {
			FilterByUser(&ml, user)
		}
		if master != 0 {
			FilterByMaster(&ml, master)
		}
		if status != "" {
			FilterByStatus(&ml, status)
		}
		return ml, nil
	}
	return nil, err
}

func FilterByTag(ml *[]ProjectJson, tag string){
	for i := 0; i < len(*ml);{
		if( !TagInArrayOfStrings(tag, (*ml)[i].Tags)) {
			(*ml)[i] = (*ml)[len(*ml)-1]
			(*ml)= (*ml)[:len(*ml)-1]
			continue
		}

		i++
	}
}

func FilterByStatus(ml *[]ProjectJson, status string) {
	for i := 0; i < len(*ml);{
		if((*ml)[i].Status != status){
			(*ml)[i] = (*ml)[len(*ml)-1]
			(*ml)= (*ml)[:len(*ml)-1]
			continue
		}
		i++
	}
}

func FilterByUser(ml *[]ProjectJson, user int64){
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectUser))
	var l []ProjectUser
	qs.All(&l)
	userProjects := make(map[int64]bool)
	for _, v := range l {
		if(int(user) == (v.UserId).Id){
			userProjects[(v.ProjectId).Id] = true
		}
	}
	for i := 0; i < len(*ml);{
		if(!userProjects[(*ml)[i].Id]){
			(*ml)[i] = (*ml)[len(*ml)-1]
			(*ml)= (*ml)[:len(*ml)-1]
			continue
		}
		i++
	}
}

func FilterByMaster(ml * []ProjectJson, master int64){
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectMaster))
	var l []ProjectMaster
	qs.All(&l)
	createdProjects := make(map[int64]bool)
	for _, v := range l {
		if(int(master) == (v.MasterId).Id){
			createdProjects[(v.ProjectId).Id] = true
		}
	}

	for i := 0; i < len(*ml);{
		if(!createdProjects[(*ml)[i].Id]){
			(*ml)[i] = (*ml)[len(*ml)-1]
			(*ml)= (*ml)[:len(*ml)-1]
			continue
		}
		i++
	}
}

// UpdateProject updates Project by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectById(m *ProjectJson) (err error) {
	proj, err := GetProjectById(m.Id)
	t := m.translate()
	t.DateOfCreation = proj.DateOfCreation
	o := orm.NewOrm()
	v := Project{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(&t)
	}
	return
}

// DeleteProject deletes Project by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProject(id int64) (err error) {
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

// Return 3 Projects for landing page. Returns empty list if no Projects exists
// This is overloaded method for GetAllProjects with parameters
// ([], [], [], [], 0, 3)
func GetLandingProjects() (ml []ProjectJson, err error) {
	var query = make(map[string]string)
	return GetAllProjects(query, []string{}, []string{}, 0, 3, "", 0, 0, "")
}
