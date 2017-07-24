package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitVariant(o orm.Ormer, testsTasks []models.TestsTask) []models.Variant {
	result := []models.Variant{
		{
			Text:          "Вариант ответа 1",
			CorrectAnswer: false,
			TestsTaskId:   &testsTasks[0],
		},
		{
			Text:          "Вариант ответа 2",
			CorrectAnswer: true,
			TestsTaskId:   &testsTasks[0],
		},
		{
			Text:          "Вариант ответа 1",
			CorrectAnswer: true,
			TestsTaskId:   &testsTasks[1],
		},
		{
			Text:          "Вариант ответа 3",
			CorrectAnswer: false,
			TestsTaskId:   &testsTasks[1],
		},
		{
			Text:          "Вариант ответа 1",
			CorrectAnswer: false,
			TestsTaskId:   &testsTasks[2],
		},
		{
			Text:          "Вариант ответа 2",
			CorrectAnswer: true,
			TestsTaskId:   &testsTasks[2],
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
