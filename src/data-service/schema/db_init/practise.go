package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitPractise(o orm.Ormer, lessons []models.Lesson) []models.Practise {
	result := []models.Practise{
		models.Practise{
			LessonId:    &lessons[0],
			Description: "Упражнение к уроку 1 Курс 1",
		},
		{
			LessonId:    &lessons[1],
			Description: "Упражнение к уроку 1 Курс 1",
		},
		{
			LessonId:    &lessons[2],
			Description: "Упражнение к уроку 1 Курс 2",
		},
		{
			LessonId:    &lessons[3],
			Description: "Упражнение к уроку 1 Курс 4",
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
