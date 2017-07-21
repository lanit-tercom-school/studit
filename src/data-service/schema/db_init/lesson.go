package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitLesson(o orm.Ormer, courses []models.Course) []models.Lesson {
	result := []models.Lesson{
		models.Lesson{
			Title:       "Урок1",
			CourseId:    &courses[0],
			Description: "Урок 1 курс 1",
			Rating:      5,
		},
		{
			Title:       "Урок2",
			CourseId:    &courses[0],
			Description: "Урок 2 курс 1",
			Rating:      5,
		},
		{
			Title:       "Урок3",
			CourseId:    &courses[0],
			Description: "Урок 3 курс 1",
			Rating:      3,
		},
		{
			Title:       "Урок1",
			CourseId:    &courses[1],
			Description: "Урок 1 курс 2",
			Rating:      4,
		},
		{
			Title:       "Урок1",
			CourseId:    &courses[3],
			Description: "Урок 1 курс 4",
			Rating:      5,
		},
	}

	param := make([]interface{}, len(result))
	for i, v := range result {
		param[i] = v
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

//return Handle(o, param).([]models.Lesson)

// func Handle(o orm.Ormer, result []interface{}) []interface{} {
// 	for i := 0; i < len(result); i++ {
// 		id, err := o.Insert(&result[i])
// 		if err != nil {
// 			beego.Critical(err.Error())
// 			panic(err)
// 		}
// 		result[i].(Id) = int(id)
// 	}

// 	return result
// }
