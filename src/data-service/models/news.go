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

func (j NewsJSONSet) MarshalJSON() ([]byte, error) {
	return []byte(j.Arr), nil
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

func GetOrderByClause(sortCols, orders []string) (string, error) {
	var result string

	if err := Check(sortCols); err != nil {
		return "", err
	}

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
		return "", errors.New("Error: 'sortby', 'orders' sizes mismatch or 'orders' size is not 1")
	}

	if len(result) > 0 {
		result = "ORDER BY " + result[:len(result)-2]
	}

	return result, nil
}

func GetSqlTagsOperation(tagsOperation string) (string, error) {

	switch tagsOperation {
	case "and":
		return "@>", nil
	case "or":
		return "&&", nil
	default:
		return "", errors.New("Error: `" + string(tagsOperation) + "` is an invalid tags operation. Must be either [and|or]")
	}
}

// GetAllNews retrieves all News matches certain condition. Returns empty list if
// no records exist
func GetAllNews(query map[string]string, fields, sortCols, orders []string, offset, limit int, tags, tagsOperation string) (interface{}, error) {
	orm.Debug = true
	o := orm.NewOrm()

	orderClause, err := GetOrderByClause(sortCols, orders)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sqlTagsOperation, err := GetSqlTagsOperation(tagsOperation)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var set NewsJSONSet

	err = o.Raw(`
SELECT row_to_json(t) arr
  FROM (SELECT COUNT(*) "TotalCount"
     , (SELECT COUNT(*) "FilteredCount"
          FROM news  WHERE tags `+sqlTagsOperation+` string_to_array($1, ','))
     , (SELECT array_to_json(array_agg(row_to_json(d))) "NewsList"
          FROM (SELECT id "Id", title "Title", description "Description", created "Created"
              , edited "Edited", tags "Tags", image "Image"
                  FROM news
                 WHERE tags `+sqlTagsOperation+` string_to_array($1, ',')
              `+orderClause+`
                OFFSET $2 LIMIT $3 ) d)
          FROM news) t`, tags, offset, limit).QueryRow(&set)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return set, nil

	//qs := o.QueryTable(new(News))
	//// query k=v
	//for k, v := range query {
	//	// rewrite dot-notation to Object__Attribute
	//	k = strings.Replace(k, ".", "__", -1)
	//	if strings.Contains(k, "isnull") {
	//		qs = qs.Filter(k, (v == "true" || v == "1"))
	//	} else {
	//		qs = qs.Filter(k, v)
	//	}
	//}
	//// order by:
	//var sortFields []string
	//if len(sortCols) != 0 {
	//	if len(sortCols) == len(order) {
	//		// 1) for each sort field, there is an associated order
	//		for i, v := range sortCols {
	//			orderby := ""
	//			if order[i] == "desc" {
	//				orderby = "-" + v
	//			} else if order[i] == "asc" {
	//				orderby = v
	//			} else {
	//				return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
	//			}
	//			sortFields = append(sortFields, orderby)
	//		}
	//		qs = qs.OrderBy(sortFields...)
	//	} else if len(sortCols) != len(order) && len(order) == 1 {
	//		// 2) there is exactly one order, all the sorted fields will be sorted by this order
	//		for _, v := range sortCols {
	//			orderby := ""
	//			if order[0] == "desc" {
	//				orderby = "-" + v
	//			} else if order[0] == "asc" {
	//				orderby = v
	//			} else {
	//				return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
	//			}
	//			sortFields = append(sortFields, orderby)
	//		}
	//	} else if len(sortCols) != len(order) && len(order) != 1 {
	//		return nil, errors.New("Error: 'sortCols', 'order' sizes mismatch or 'order' size is not 1")
	//	}
	//} else {
	//	if len(order) != 0 {
	//		return nil, errors.New("Error: unused 'order' fields")
	//	}
	//}
	//
	//var l []News
	//qs = qs.OrderBy(sortFields...)
	//if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
	//	if len(fields) == 0 {
	//		for _, v := range l {
	//			ml = append(ml, v)
	//		}
	//	} else {
	//		// trim unused fields
	//		for _, v := range l {
	//			m := make(map[string]interface{})
	//			val := reflect.ValueOf(v)
	//			for _, fname := range fields {
	//				m[fname] = val.FieldByName(fname).Interface()
	//			}
	//			ml = append(ml, m)
	//		}
	//	}
	//	return ml, nil
	//}
	//return nil, err
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
