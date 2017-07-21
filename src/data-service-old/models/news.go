package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type News struct {
	Id          int       `orm:"column(id);pk;auto"`
	Title       string    `orm:"column(title)"`
	Description string    `orm:"column(description)"`
	Created     time.Time `orm:"column(created);auto_now_add"`
	Edited      time.Time `orm:"column(edited);auto_now"`
	Tags        []string  `orm:"-"`
	Image       string    `orm:"column(image)"`
}

type NewsSet struct {
	TotalCount    int64
	FilteredCount int64
	NewsList      []News
}

func init() {
	orm.RegisterModel(new(News))
}

// AddNews inserts a new News into database and returns last inserted id on success
func AddNews(news *News) (id int64, err error) {
	//If needed set news.Id = 0 for auto inc
	o := orm.NewOrm()
	id, err = o.Insert(&news)
	return
}

// GetNewsById retrieves News by Id and returns error if id doesn't exist
func GetNewsById(id int) (m *News, err error) {
	o := orm.NewOrm()
	v := &News{Id: id}
	if err = o.Read(v); err == nil {
		//m_temp := v.translate()  // need a temp variable
		//m = &m_temp
		//return m, nil
		return v, nil
	}
	return nil, err
}

func Check(sortExpressions []string) error {
	permittedKeysUsage := map[string]bool{
		"id":      false,
		"created": false,
		"edited":  false,
		"title":   false,
	}

	for _, v := range sortExpressions {
		used, ok := permittedKeysUsage[v]
		if !ok {
			return errors.New("Error: `" + v + "` is an invalid sortBy key")
		}
		if used {
			return errors.New("Error: sortBy key `" + v + "` used more than once")
		}
		permittedKeysUsage[v] = true
	}

	return nil
}

func GetOrderByClause(sortExpressions, orders []string) (string, error) {
	var result string

	if err := Check(sortExpressions); err != nil {
		return "", err
	}

	switch len(orders) {
	case 0:
		if len(sortExpressions) != 0 {
			return "", errors.New("Error: Insufficient number of orders")
		}
	// there is exactly one orders, all the sorted fields will be sorted by this orders
	case 1:
		for _, v := range sortExpressions {
			switch orders[0] {
			case "desc", "asc":
				result += v + " " + orders[0] + ", "
			default:
				return "", errors.New("Error: Invalid orders. Must be either [asc|desc]")
			}
		}
	// there is an orders for each sorted fields
	case len(sortExpressions):
		for i, v := range sortExpressions {
			switch orders[i] {
			case "desc", "asc":
				result += v + " " + orders[i] + ", "
			default:
				return "", errors.New("Error: Invalid orders. Must be either [asc|desc]")
			}
		}
	default:
		if len(sortExpressions) == 0 {
			return "", errors.New("Error: Unused 'orders' fields")
		}
		return "", errors.New("Error: 'sortby', 'orders' sizes mismatch or 'orders' size is not 1")
	}

	if len(result) > 0 {
		result = "ORDER BY " + result[:len(result)-2]
	}

	return result, nil
}

type NewsJSONSet struct {
	Arr string `orm:"type(json)"`
}

func (j NewsJSONSet) MarshalJSON() ([]byte, error) {
	return []byte(j.Arr), nil
}

// GetAllNews returns News that match certain conditions
func GetAllNews(sortExpressions, orders []string, offset, limit int, tags string) (interface{}, error) {
	orm.Debug = true
	o := orm.NewOrm()

	orderClause, err := GetOrderByClause(sortExpressions, orders)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var set NewsJSONSet
	err = o.Raw(`
SELECT row_to_json(t) arr
  FROM (SELECT COUNT(*) "TotalCount"
     , (SELECT COUNT(*) "FilteredCount"
          FROM news  WHERE tags @> string_to_array($1, ','))
     , (SELECT array_to_json(array_agg(row_to_json(d))) "NewsList"
          FROM (SELECT id "Id", title "Title", description "Description", created "Created"
              , edited "Edited", tags "Tags", image "Image"
                  FROM news
                 WHERE tags @> string_to_array($1, ',') `+orderClause+` ) d)
          FROM news) t`, tags).QueryRow(&set)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return set, nil
}

// UpdateNews updates News by Id and returns error if
// the record to be updated doesn't exist
func UpdateNewsById(m *News) (err error) {
	//m.Edited = time.Now()
	t := m //.translate()
	o := orm.NewOrm()
	v := News{Id: m.Id}
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
	v := News{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Delete(&News{Id: id})
	}
	return
}
