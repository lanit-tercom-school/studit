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

type NewsJSONSet struct {
	Arr string `orm:"type(json)"`
}

func (set NewsJSONSet) MarshalJSON() ([]byte, error) {
	return []byte(set.Arr), nil
}

func init() {
	orm.RegisterModel(new(News))
}

// AddNews insert a new News into database and returns
// last inserted Id on success.
func AddNews(m *News) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNewsById retrieves News by Id. Returns error if
// Id doesn't exist
func GetNewsById(id int) (v *News, err error) {
	o := orm.NewOrm()
	v = &News{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNews retrieves all News matches certain condition. Returns empty list if no records exist
func GetAllNews(sortCols, orders []string, offset, limit int, tags, tagsOperator string) (interface{}, error) {
	orm.Debug = true
	o := orm.NewOrm()

	sqlTagsOperator, err := GetSqlTagsOperator(tagsOperator)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	orderClause, err := GetOrderByClause(sortCols, orders)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var set NewsJSONSet

	err = o.Raw(`
SELECT row_to_json(t) arr
  FROM (SELECT COUNT(*) "TotalCount"
     , (SELECT COUNT(*) "FilteredCount"
          FROM news
         WHERE tags `+sqlTagsOperator+` string_to_array($1, ','))
     , (SELECT array_to_json(array_agg(row_to_json(d))) "NewsList"
          FROM (SELECT id "Id", title "Title", description "Description", created "Created"
              , edited "Edited", tags "Tags", image "Image"
                  FROM news
                 WHERE tags `+sqlTagsOperator+` string_to_array($1, ',')
              `+orderClause+`ORDER BY created ASC
                OFFSET $2 LIMIT $3 ) d)
          FROM news ) t`, tags, offset, limit).QueryRow(&set)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return set, nil
}

func GetSqlTagsOperator(tagsOperator string) (string, error) {
	switch tagsOperator {
	case "and":
		return "@>", nil
	case "or":
		return "&&", nil
	default:
		return "", errors.New("Error: `" + tagsOperator + "` is an invalid tags operation. Must be either [and|or]")
	}
}

func GetOrderByClause(sortCols, orders []string) (string, error) {
	if err := Check(sortCols); err != nil {
		return "", err
	}

	var result string
	switch len(orders) {
	case 0:
		if len(sortCols) != 0 {
			return "", errors.New("Error: Insufficient number of orders")
		}
		// there is exactly one orders, all the sorted fields will be sorted by this orders
	case 1:
		for _, v := range sortCols {
			switch orders[0] {
			case "desc", "asc":
				result += v + " " + orders[0] + ", "
			default:
				return "", errors.New("Error: Invalid orders. Must be either [asc|desc]")
			}
		}
		// there is an orders for each sorted fields
	case len(sortCols):
		for i, v := range sortCols {
			switch orders[i] {
			case "desc", "asc":
				result += v + " " + orders[i] + ", "
			default:
				return "", errors.New("Error: Invalid orders. Must be either [asc|desc]")
			}
		}
	default:
		if len(sortCols) == 0 {
			return "", errors.New("Error: Unused 'orders' fields")
		}
		return "", errors.New("Error: 'sortCols' and 'orders' sizes mismatch or 'orders' size is not 1")
	}

	if len(result) > 0 {
		result = "ORDER BY " + result[:len(result)-2]
	}

	return result, nil
}

func Check(sortCols []string) error {
	permittedKeysUsage := map[string]bool{
		"id":      false,
		"created": false,
		"edited":  false,
		"title":   false,
	}

	for _, v := range sortCols {
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

// UpdateNews updates News by Id and returns error if
// the record to be updated doesn't exist
func UpdateNewsById(m *News) (err error) {
	o := orm.NewOrm()
	v := News{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
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
		var num int64
		if num, err = o.Delete(&News{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
