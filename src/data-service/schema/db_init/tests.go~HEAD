package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitTest(o orm.Ormer, lessons []models.Lesson) []models.Test {
	result := []models.Test{
		models.Test{
			Title:    "Тест 1 урок 1 курс 1",
			LessonId: &lessons[0],
		},
		{
			Title:    "Тест 2 урок 1 курс 1",
			LessonId: &lessons[1],
		},
		{
			Title:    "Тест 1 урок 3 курс 1",
			LessonId: &lessons[2],
		},
		{
			Title:    "Тест 1 урок 1 курс 4",
			LessonId: &lessons[4],
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
