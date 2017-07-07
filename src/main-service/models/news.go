package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type news struct {
	Id             int       `orm:"column(id);pk;auto"`
	Title          string    `orm:"column(title)"`
	Description    string    `orm:"column(description)"`
	DateOfCreation time.Time `orm:"column(date_of_creation);type(datetime)"`
	LastEdit       time.Time `orm:"column(last_edit);type(datetime)"`
	Tags           string    `orm:"column(tags)"`
	Image          string    `orm:"column(image)"`
}

type NewsJson struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Edited      time.Time `json:"edited"`
	Tags        []string  `json:"tags"`
	Image       string    `json:"image"`
}

type NewsSetJson struct {
	TotalCount   int        `json:"total_count"`
	NewsJsonList []NewsJson `json:"news_list"`
}

func (t *news) translate() NewsJson {
	return NewsJson{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		Created:     t.DateOfCreation,
		Edited:      t.LastEdit,
		Tags:        strings.Split(t.Tags, ","),
		Image:       t.Image,
	}
}

func (t *NewsJson) translate() news {
	return news{
		Id:             t.Id,
		Title:          t.Title,
		Description:    t.Description,
		DateOfCreation: t.Created,
		LastEdit:       t.Edited,
		Tags:           strings.Join(t.Tags, ","),
		Image:          t.Image,
	}
}

func (t *news) TableName() string {
	return "news"
}

func init() {
	orm.RegisterModel(new(news))
}

// AddNews inserts a new News into database and returns last inserted id on success
func AddNews(m *NewsJson) (id int64, err error) {
	v := m.translate()
	v.Id = 0 // for auto inc
	v.DateOfCreation = time.Now()
	v.LastEdit = time.Now()
	o := orm.NewOrm()
	id, err = o.Insert(&v)
	return
}

// GetNewsById retrieves News by Id and returns error if id doesn't exist
func GetNewsById(id int) (m *NewsJson, err error) {
	o := orm.NewOrm()
	v := &news{Id: id}
	if err = o.Read(v); err == nil {
		m_temp := v.translate() // need a temp variable
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

func TagInString(tag, tags string) bool {
	temp_tags := strings.Split(tags, ",")
	for _, t := range temp_tags {
		if t == tag {
			return true
		}
	}
	return false
}

func GetSortFields(sortBy, order []string) ([]string, error) {
	var result []string

	switch len(order) {
	case 0:
		if len(sortBy) != 0 {
			return nil, errors.New("Error: Insufficient number of orders")
		}
	// there is exactly one order, all the sorted fields will be sorted by this order
	case 1:
		for _, v := range sortBy {
			switch order[0] {
			case "desc":
				result = append(result, "-"+v)
			case "asc":
				result = append(result, v)
			default:
				return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
			}
		}
	// there is an order for each sorted fields
	case len(sortBy):
		for i, v := range sortBy {
			switch order[i] {
			case "desc":
				result = append(result, "-"+v)
			case "asc":
				result = append(result, v)
			default:
				return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
			}
		}
	default:
		if len(sortBy) == 0 {
			return nil, errors.New("Error: Unused 'order' fields")
		} else {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	}

	return result, nil
}

// GetAllNews returns News that match certain conditions
func GetAllNews(sortBy, order []string, offset, limit int, tag string) (
	interface{}, error) {
	orm.Debug = true

	sortFields, err := GetSortFields(sortBy, order)
	if err !=nil {
		return nil, err
	}

	query := orm.NewOrm().QueryTable(new(news))
	query = query.OrderBy(sortFields...)

	totalCount, err := query.Count()

	var newsList []news
	if _, err = query.Limit(limit, offset).All(&newsList); err != nil {
		return nil, err
	}

	var jsonNewsList []NewsJson

	if tag == "" {
		for _, v := range newsList {
			jsonNewsList = append(jsonNewsList, v.translate())
		}
	} else {
		for _, v := range newsList {
			r := v.translate()
			if TagInArrayOfStrings(tag, r.Tags) {
				jsonNewsList = append(jsonNewsList, r)
			}
		}
	}

	set := NewsSetJson{
		TotalCount:   int(totalCount),
		NewsJsonList: jsonNewsList}

	return set, nil
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
