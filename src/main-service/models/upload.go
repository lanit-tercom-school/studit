package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type File struct {
	Id             int       `orm:"column(id);pk;auto"                         json:"id"`
	Name           string    `orm:"column(name)"                               json:"name"`
	Path           string    `orm:"column(path)"                               json:"path,omitempty"`
	DateOfCreation time.Time `orm:"column(date_of_creation);type(datetime)"    json:"created"`
	User           *User     `orm:"column(user_id);rel(fk)"                    json:"user"`
}

func (t *File) TableName() string {
	return "file"
}

func init() {
	orm.RegisterModel(new(File))
}

// AddFile insert a new File into database and returns
// last inserted Id on success.
func AddFile(m *File) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
