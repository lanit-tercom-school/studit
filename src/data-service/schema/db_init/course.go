package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitCourse(o orm.Ormer) []models.Course {
	result := []models.Course{
		{
			Id:          1,
			Title:       "Курс 1",
			Description: "Описание курса",
			Logo:        "/files/course1.jpg",
			Rating:      4.8,
		},
		{
			Title:       "Курс 2",
			Description: "Описание курса",
			Logo:        "/files/course2.jpg",
			Rating:      4.5,
		},
		{
			Title:       "Курс 3",
			Description: "Описание курса",
			Logo:        "/files/course3.jpg",
			Rating:      4.6,
		},
		{
			Title:       "Курс 4",
			Description: "Описание курса",
			Logo:        "/files/course4.jpg",
			Rating:      3.9,
		},
	}

	for i := 0; i < len(result); i++ {
		id, err := o.Insert(&result[i])
		if err != nil {
			beego.Critical(err.Error())
			panic(err)
		}
		result[i].Id = int(id)
	}

	return result
}
