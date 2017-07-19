package db_init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"main-service/models"
)

func InitStatistics(o orm.Ormer, courses []models.Course ) ([]models.Statistics) {
	result := []models.Statistics{
		models.Statistics{
			Hours:    12,
			CourseId: &courses[0],
		},
		models.Statistics{
			Hours:    15,
			CourseId: &courses[1],
		},
		models.Statistics{
			Hours:    15,
			CourseId: &courses[2],
		},
		models.Statistics{
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