package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitStatistics(o orm.Ormer, courses []models.Course) []models.Statistics {
	result := []models.Statistics{
		models.Statistics{
			Hours:    12,
			CourseId: &courses[0],
		},
		{
			Hours:    15,
			CourseId: &courses[1],
		},
		{
			Hours:    15,
			CourseId: &courses[2],
		},
		{
			Hours:    18,
			CourseId: &courses[3],
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
