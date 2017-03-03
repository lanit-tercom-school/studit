package main

import (
	_ "service/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	m "service/models"
	"log"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
}

func fastCheckErr(_ int64, err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func main() {
	log.SetFlags(log.Ltime | log.Llongfile)
	o := orm.NewOrm()
	o.Using("default")

	// add projects

	project := m.Project{
		Id: 1,
		Name: "Образовательный портал Studit",
		Description: "Разработка образовательного портала для Lanit-Tercom School",
		Logo: "/logo/1.jpg",
	}
	fastCheckErr(o.Insert(&project))

	project = m.Project{
		Id: 2,
		Name: "Модный фрилансер",
		Description: "Какие же стрелочки вокруг ноубука!",
		Logo: "/logo/2.jpg",
	}
	fastCheckErr(o.Insert(&project))

	project = m.Project{
		Id: 3,
		Name: "Оригинальное название",
		Description: "Click-bait описание",
		Logo: "/logo/3.jpg",
	}
	fastCheckErr(o.Insert(&project))

	// add users

	user := m.User{
		Id: 1,
		Nickname: "Admin",
		Login: "admin@admin.admin",
		Password: "admin",
		Avatar: "/logo/1.jpg",
		Description: "Главный по тарелкам",
	}
	fastCheckErr(o.Insert(&project))

	user = m.User{
		Id: 2,
		Nickname: "Moderator",
		Login: "moder@moder.moder",
		Password: "moder",
		Avatar: "/logo/2.jpg",
		Description: "Главный по молоткам",
	}
	fastCheckErr(o.Insert(&user))

	user = m.User{
		Id: 3,
		Nickname: "Егорка2003",
		Login: "egorka2003@maaail.ru",
		Password: "пароль",
		Avatar: "/logo/3.jpg",
		Description: "ЮЮЮ, ААА",
	}
	fastCheckErr(o.Insert(&user))

	// add tags
	tag := m.Tag{
		Id: 100,
		Name: "Other",
	}
	fastCheckErr(o.Insert(&tag))

	tag = m.Tag{
		Id: 101,
		Name: "New",
	}
	fastCheckErr(o.Insert(&tag))

	tag = m.Tag{
		Id: 102,
		Name: "C/C++",
	}
	fastCheckErr(o.Insert(&tag))

	tag = m.Tag{
		Id: 103,
		Name: "Golang",
	}
	fastCheckErr(o.Insert(&tag))

	tag = m.Tag{
		Id: 104,
		Name: "JavaScript",
	}
	fastCheckErr(o.Insert(&tag))

	log.Print("Initial data was successfully added to Database")
}