package models

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// Модель для базы данных
type Project struct {
	Id          int       `orm:"column(id);pk;auto"           json:"id"`
	Name        string    `orm:"column(name)"                 json:"name"`
	Description string    `orm:"column(description)"          json:"description"`
	Created     time.Time `orm:"column(created);auto_now_add" json:"created"`
	Logo        string    `orm:"column(logo)"                 json:"logo"`
	Tags        []string  `orm:"-"                            json:"tags"`
	Status      string    `orm:"column(status)"               json:"status"`
}

// Вся информация о проекте
type AllInformationAboutProject struct {
	Project  *Project       `json:"project"`
	Enrolled []MainUserInfo `json:"enrolled"`
	Members  []MainUserInfo `json:"members"`
	Masters  []MainUserInfo `json:"masters"`
}

type MainProjectInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// func (t *Project) TableName() string {
// 	return "project"
// }

func init() {
	orm.RegisterModel(new(Project))
}

// AddProject insert a new Project into database and returns
// last inserted Id on success.
func AddProject(p *Project) (id int, err error) {
	p.Id = 0 // for auto inc
	p.Created = time.Now()
	p.Status = "opened"
	o := orm.NewOrm()
	id_, err := o.Insert(&p)
	// TODO add tags update after o.Insert()
	id = int(id_)
	return
}

// GetProjectById retrieves Project by Id. Returns error if
// Id doesn't exist
func GetProjectById(id int) (*Project, error) {
	o := orm.NewOrm()
	result := &Project{Id: id}
	err := o.Read(result)
	if err != nil {
		result = nil
	}
	return result, err
}

// GetAllProjects retrieves all Project matches certain condition. Returns empty list if
// no records exist
func GetAllProjects(query map[string]string, sortBy []string, order []string,
	offset int64, limit int64, tag string, user int64, master int64, status string) (ml []Project, err error) {
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
			ml = append(ml, v)
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

func FilterByTag(ml *[]Project, tag string) {
	//for i := 0; i < len(*ml); {
	//	if !TagInArrayOfStrings(tag, (*ml)[i].Tags) {
	//		(*ml)[i] = (*ml)[len(*ml)-1]
	//		*ml = (*ml)[:len(*ml)-1]
	//		continue
	//	}
	//	i++
	//}
}

func FilterByStatus(ml *[]Project, status string) {
	for i := 0; i < len(*ml); {
		if (*ml)[i].Status != status {
			// remove from slice, https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
			(*ml)[i] = (*ml)[len(*ml)-1]
			*ml = (*ml)[:len(*ml)-1]
			continue
		}
		i++
	}
}

func FilterByUser(ml *[]Project, user int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectUser))
	var l []ProjectUser
	qs.All(&l)
	userProjects := make(map[int]bool)
	for _, v := range l {
		if int(user) == (v.UserId).Id {
			userProjects[(v.ProjectId).Id] = true
		}
	}
	for i := 0; i < len(*ml); {
		if !userProjects[(*ml)[i].Id] {
			(*ml)[i] = (*ml)[len(*ml)-1]
			*ml = (*ml)[:len(*ml)-1]
			continue
		}
		i++
	}
}

func FilterByMaster(ml *[]Project, master int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectMaster))
	var l []ProjectMaster
	qs.All(&l)
	createdProjects := make(map[int]bool)
	for _, v := range l {
		if int(master) == (v.MasterId).Id {
			createdProjects[(v.ProjectId).Id] = true
		}
	}

	for i := 0; i < len(*ml); {
		if !createdProjects[(*ml)[i].Id] {
			(*ml)[i] = (*ml)[len(*ml)-1]
			*ml = (*ml)[:len(*ml)-1]
			continue
		}
		i++
	}
}

// UpdateProject updates Project by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectById(p *Project) (err error) {
	proj, err := GetProjectById(p.Id)
	p.Created = proj.Created
	o := orm.NewOrm()
	v := Project{Id: p.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(&p)
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
		_, err = o.Delete(&Project{Id: id})
	}
	return
}

// Return 3 Projects for landing page. Returns empty list if no Projects exists
// This is overloaded method for GetAllProjects with parameters
// ([], [], [], [], 0, 3)
func GetLandingProjects() (ml []Project, err error) {
	var query = make(map[string]string)
	return GetAllProjects(query, []string{}, []string{}, 0, 3, "", 0, 0, "")
}
