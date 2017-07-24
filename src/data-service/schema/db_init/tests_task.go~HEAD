package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitTestsTask(o orm.Ormer, tests []models.Test) []models.TestsTask {
	result := []models.TestsTask{
		models.TestsTask{
			Question: "Вопрос 1 Тест 1 Урок 1 Курс 1",
			TestId:   &tests[0],
		},
		{
			Question: "Вопрос 2 Тест 1 Урок 1 Курс 1",
			TestId:   &tests[0],
		},
		{
			Question: "Вопрос 1 Тест 2 Урок 1 Курс 1",
			TestId:   &tests[2],
		},
		{
			Question: "Вопрос 1 Тест 1 Урок 1 Курс 4",
			TestId:   &tests[3],
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
