package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type File struct {
	Id             int       `orm:"column(id);pk;auto"`
	Name           string    `orm:"column(name)"`
	Path           string    `orm:"column(path)"`
	DateOfCreation time.Time `orm:"column(date_of_creation);type(datetime)"`
	User           int       `orm:"column(user_id);"`
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

// GetFileById retrieves File by Id. Returns error if
// Id doesn't exist
func GetFileById(id int) (v *File, err error) {
	o := orm.NewOrm()
	v = &File{Id: id}
	if err = o.QueryTable("file").Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserById retrieves Files by user.
func GetFilesByUserId(userId int) (v []File, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("file").Filter("user_id", userId).RelatedSel().All(&v)
	if err == nil {
		return v, nil
	}
	return
}
func DeleteFile(m *File) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Delete(m)
	return
}
