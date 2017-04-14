package models

import (
	"errors"
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
)

type news struct {
	Id                  int         `orm:"column(id);pk;auto"`
	Title               string      `orm:"column(title)"`
	Description         string      `orm:"column(description)"`
	DateOfCreation      time.Time   `orm:"column(date_of_creation);type(datetime)"`
	LastEdit            time.Time   `orm:"column(last_edit);type(datetime)"`
	Tags                string      `orm:"column(tags)"`
	Image               string      `orm:"columt(image)"`
}

type NewsJson struct {
	Id              int         `json:"id"`
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	Created         time.Time   `json:"created"`
	Edited          time.Time   `json:"edited"`
	Tags            []string    `json:"tags"`
	Image           string      `json:"image"`
}

func (t *news) translate() NewsJson {
	return NewsJson{
		Id: t.Id,
		Title: t.Title,
		Description: t.Description,
		Created: t.DateOfCreation,
		Edited: t.LastEdit,
		Tags: strings.Split(t.Tags, ","),
		Image: t.Image,
	}
}

func (t *NewsJson) translate() news {
	return news{
		Id: t.Id,
		Title: t.Title,
		Description: t.Description,
		DateOfCreation: t.Created,
		LastEdit: t.Edited,
		Tags: strings.Join(t.Tags, ","),
		Image: t.Image,
	}
}

func (t *news) TableName() string {
	return "news"
}

func init() {
	orm.RegisterModel(new(news))
}

// AddNews insert a new News into database and returns
// last inserted Id on success.
func AddNews(m *NewsJson) (id int64, err error) {
	v := m.translate()
	v.Id = 0 // for auto inc
	v.DateOfCreation = time.Now()
	v.LastEdit = time.Now()
	o := orm.NewOrm()
	id, err = o.Insert(&v)
	return
}

// GetNewsById retrieves News by Id. Returns error if
// Id doesn't exist
func GetNewsById(id int) (m *NewsJson, err error) {
	o := orm.NewOrm()
	v := &news{Id: id}
	if err = o.Read(v); err == nil {
		m_temp := v.translate()  // need a temp variable
		m = &m_temp
		return m, nil
	}
	return nil, err
}

func TagInArrayOfStrings(tag string, tags []string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func TagInString(tag string, tags string) bool {
	temp_tags := strings.Split(tags, ",")
	for _, t := range temp_tags {
		if t == tag {
			return true
		}
	}
	return false
}

// GetAllNews retrieves all News matches certain condition. Returns empty list if
// no records exist
func GetAllNews(sortBy []string, order []string, offset int64, limit int64, tag string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(news))

	// Step 1: parse input args to SQL syntax

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
			return nil, errors.New("Error: Unused 'order' fields")
		}
	}

	// Step 2: Select items from table with params

	var l []news
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l); err == nil {
		if tag == "" {
			for _, v := range l {
				ml = append(ml, v.translate())
			}
			return ml, nil
		} else {
			for _, v := range l {
				r := v.translate()
				if TagInArrayOfStrings(tag, r.Tags) {
					ml = append(ml, r)
				}
			}
			return ml, nil
		}
	}
	return nil, err
}

// UpdateNews updates News by Id and returns error if
// the record to be updated doesn't exist
func UpdateNewsById(m *NewsJson) (err error) {
	m.Edited = time.Now()
	t := m.translate()
	o := orm.NewOrm()
	v := news{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(&t)
	}
	return
}

// DeleteNews deletes News by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNews(id int) (err error) {
	o := orm.NewOrm()
	v := news{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&news{Id: id})
	}
	return
}
