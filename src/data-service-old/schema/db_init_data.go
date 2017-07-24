package main

import (
	"data-service-old/schema/db_init"
	_ "main-service/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	err := orm.RegisterDataBase("default", "postgres",
		"postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")

	projects := db_init.InitProject(o)
	users := db_init.InitUser(o)
	db_init.InitProjectUser(o, projects, users)
	contactTypes := db_init.InitContactType(o)
	db_init.InitUserContact(o, users, contactTypes)
	comments := db_init.InitComment(o)
	db_init.InitUserComment(o, users, comments)
	db_init.InitNews(o)
	courses := db_init.InitCourse(o)
	db_init.InitStatistics(o, courses)
	lessons := db_init.InitLesson(o, courses)
	db_init.InitRecommendedCourse(o, courses)
	tests := db_init.InitTest(o, lessons)
	testsTasks := db_init.InitTestsTask(o, tests)
	db_init.InitPractise(o, lessons)
	db_init.InitVideo(o, lessons)
	db_init.InitVariant(o, testsTasks)

	beego.Info("Initial data was successfully added to Database")
}
