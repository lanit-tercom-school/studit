package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	m "auth-service/models"
	auth "auth-service/controllers"
)

func init() {
	err := orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit_auth?sslmode=disable")
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func fastCheckErr(_ int64, err error) {
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	user1 := m.User{
		Id: 1,
		Login: "a@a",
		Password: auth.CustomStr("a").ToSHA1(),
		PermissionLevel: 2,
	}
	fastCheckErr(o.Insert(&user1))

	user2 := m.User{
		Id: 2,
		Login: "moder@moder.moder",
		Password: auth.CustomStr("moder").ToSHA1(),
		PermissionLevel: 1,
	}
	fastCheckErr(o.Insert(&user2))

	user3 := m.User{
		Id: 3,
		Login: "egorka2003@maaail.ru",
		Password: auth.CustomStr("пароль").ToSHA1(),
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user3))

	user4 := m.User{
		Id: 4,
		Login: "zagadka@maaail.ru",
		Password: auth.CustomStr("котикипёсики").ToSHA1(),
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user4))

	user5 := m.User{
		Id: 5,
		Login: "slayer342@bbk.ru",
		Password: auth.CustomStr("lala").ToSHA1(),
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user5))

	user6 := m.User{
		Id: 6,
		Login: "petrovich82@maaail.ru",
		Password: auth.CustomStr("pasuwaado").ToSHA1(),
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user6))

	user7 := m.User{
		Id: 7,
		Login: "b@b",
		Password: auth.CustomStr("b").ToSHA1(),
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user7))
}
