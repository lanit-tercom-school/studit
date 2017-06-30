package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"net/http"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"encoding/json"
)

const (
	VIEWER = iota - 1   // незарегистрированный пользователь
	USER                // обычный пользователь
	LEADER              // учитель, куратор, меет право создавать проекты
	ADMIN               // администратор, может всё
)

type User struct {
	Id       int     `orm:"column(id);pk;auto"                   json:"id"`
	Login    string  `orm:"column(login)"                        json:"-"`
	Password string  `orm:"column(password)"                     json:"-"`
	// viewer - -1, registered user - 0, leader - 1, admin - 2, default is -1
	// Can't be higher than `auth.MaxPermissionLevel` !
	PermissionLevel int     `orm:"column(permission_level);default(0)"  json:"permission_level"`
}

func init() {
	orm.RegisterModel(new(User))
}

// return true if found, false if not
func (m *User) FindByLogin() bool {
	var anotherUser User
	err := orm.NewOrm().QueryTable("user").Filter("login", m.Login).One(&anotherUser)
	if err == orm.ErrMultiRows {
		panic(err)
	} else if err == orm.ErrNoRows {
		return false
	} else {
		return true
	}
}

// return true if found, false if not
func GetUserByLogin(login string) (*User, error) {
	var anotherUser User
	err := orm.NewOrm().QueryTable("user").Filter("login", login).One(&anotherUser)
	if err == orm.ErrMultiRows {
		panic(err)
	} else if err == orm.ErrNoRows {
		return nil, errors.New("Not found")
	} else {
		return &anotherUser, nil
	}
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (m *User) Insert() (int, error) {
	if id, err := orm.NewOrm().Insert(m); err != nil {
		return 0, err
	} else {
		return int(id), nil
	}
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() (err error) {
	o := orm.NewOrm()
	if err = o.Read(m); err == nil {
		_, err = o.Delete(m)
	}
	return
}

type MainUserInfo struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func GetMainUserInfo(id int) *MainUserInfo {
	resp, err := http.Get(fmt.Sprintf("localhost:8081/v1/user/%d?cut=true", id))
	if err != nil {
		beego.Critical("Get MainUserInfo error:", err.Error())
		return nil
	} else {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			beego.Critical("Get MainUserInfo error:", err.Error())
			return nil
		} else {
			var temp struct{
				User MainUserInfo `json:"user"`
			}
			err = json.Unmarshal(body, &temp)
			if err != nil {
				return nil
			} else {
				return &temp.User
			}
		}
	}
}