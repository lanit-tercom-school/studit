package main

import (
	_ "service/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	m "service/models"
	"fmt"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
}

func main() {
	o := orm.NewOrm()
	o.Using("default")
	project1 := m.Project{
		Id: 1,
		Name: "Образовательный портал Lanit-Tercom",
		Description: "Разработка образовательного портала для Lanit-Tercom School",
		Logo: "/logo/1.jpg",
	}
	_, err := o.Insert(&project1)
	if err != nil {
		fmt.Println("panic " + err.Error())
	}

	project2 := m.Project{
		Id: 2,
		Name: "Модный фрилансер",
		Description: "Какие же стрелочки вокруг ноубука!",
		Logo: "/logo/2.jpg",
	}
	_, err = o.Insert(&project2)
	if err != nil {
		fmt.Println("panic " + err.Error())
	}

	project3 := m.Project{
		Id: 3,
		Name: "Оригинальное название",
		Description: "Click-bait описание",
		Logo: "/logo/3.jpg",
	}
	_, err = o.Insert(&project3)
	if err != nil {
		fmt.Println("panic " + err.Error())
	}
}