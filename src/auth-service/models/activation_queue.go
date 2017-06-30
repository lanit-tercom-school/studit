package models

import "github.com/astaxie/beego/orm"

type NotActivatedOnMainServiceUser struct {
	Id       int `orm:"column(id);pk;auto"`
	UserId   int `orm:"column(user_id)"`
	Nickname string `orm:"column(nickname)"`
}

func init() {
	orm.RegisterModel(new(NotActivatedOnMainServiceUser))
}

func GetAllPendingForActivationUsers() (users []NotActivatedOnMainServiceUser, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("wait_for_activation").All(&users)
	return
}

func (u *NotActivatedOnMainServiceUser) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (m *NotActivatedOnMainServiceUser) Delete() (err error) {
	o := orm.NewOrm()
	if err = o.Read(m); err == nil {
		_, err = o.Delete(m)
	}
	return
}