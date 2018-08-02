package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ProjectNews struct {
	Id          int       `orm:"column(id);pk;auto"`
	Project_id  int       `orm:"column(project_id)"`
	Title       string    `orm:"column(title)"`
	Description string    `orm:"column(description)"`
	Created     time.Time `orm:"column(created);auto_now_add"`
	Edited      time.Time `orm:"column(edited);auto_now_add"`
	Image       string    `orm:"column(image)"`
}

func init() {
	orm.RegisterModel(new(ProjectNews))
}

// AddProjectNews добавляет новость проекта в БД
// throws error if couldn't insert project news into data base
func AddProjectNews(m *ProjectNews) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProjectNewsById достает новость проекта по id новости из БД
// throws error if news with this ID don't exist
func GetProjectNewsById(id int) (v *ProjectNews, err error) {
	o := orm.NewOrm()
	v = &ProjectNews{Id: id}
	if err = o.QueryTable(v).Filter("Id", v.Id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
