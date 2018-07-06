package models

import (
	"errors"
	"fmt"
	"time"

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

type ProjectsJSONSet struct {
	Arr string `orm:"type(json)"`
}

func (set ProjectsJSONSet) MarshalJSON() ([]byte, error) {
	return []byte(set.Arr), nil
}

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
func GetAllProject(fields []string, sortby []string, order []string,
	offset int64, limit int64) (interface{}, error) {
	o := orm.NewOrm()
	orderSort, err := GetOrderBySort(sortby, order)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var set ProjectsJSONSet

	err = o.Raw(`
		SELECT row_to_json(t) arr
		  FROM (SELECT COUNT(*) "TotalCount"
			 , (SELECT COUNT(*) "FilteredCount"
				  FROM project)
			 , (SELECT array_to_json(array_agg(row_to_json(d))) "ProjectsList"
				  FROM (SELECT id "Id", created "Created", description "Description", logo "Logo"
					  , name "Name", status "Status",  tags "Tags", githuburl "Githuburl"
						  FROM project`+orderSort+`
						OFFSET $1 LIMIT $2 ) d)
				  FROM project) t`, offset, limit).QueryRow(&set)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return set, nil

}

func GetOrderBySort(sortby, order []string) (string, error) {
	if err := Check(sortby); err != nil {
		return "", err
	}

	var result string
	switch len(order) {
	case 0:
		if len(sortby) != 0 {
			return "", errors.New("Error: Insufficient number of order")
		}
		// there is exactly one orders, all the sorted fields will be sorted by this orders
	case 1:
		for _, v := range sortby {
			switch order[0] {
			case "desc", "asc":
				result += v + " " + order[0] + ", "
			default:
				return "", errors.New("Error: Invalid orders. Must be either [asc|desc]")
			}
		}
		// there is an orders for each sorted fields
	case len(sortby):
		for i, v := range sortby {
			switch order[i] {
			case "desc", "asc":
				result += v + " " + order[i] + ", "
			default:
				return "", errors.New("Error: Invalid orders. Must be either [asc|desc]")
			}
		}
	default:
		if len(sortby) == 0 {
			return "", errors.New("Error: Unused 'orders' fields")
		}
		return "", errors.New("Error: 'sortCols' and 'orders' sizes mismatch or 'orders' size is not 1")
	}

	if len(result) > 0 {
		result = "ORDER BY " + result[:len(result)-2]
	}

	return result, nil
}

// UpdateProject updates Project by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectById(m *Project) (err error) {
	o := orm.NewOrm()
	v := Project{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
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
