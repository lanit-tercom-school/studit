package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
)

func init() {
	err := orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit_file?sslmode=disable")
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

	// add projects

	/*project1 := m.Project{
		Id:             1,
		Name:           "Образовательный портал Studit",
		Description:    "Разработка образовательного портала для Lanit-Tercom School",
		DateOfCreation: time.Now(),
		Logo:           "http://85.143.214.42/files/1.jpg",
		Tags:           "studit,summerschool",
		Status:         1,
	}
	fastCheckErr(o.Insert(&project1))*/

	beego.Info("Initial data was successfully added to Database")
}
