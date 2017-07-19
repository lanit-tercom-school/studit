package db_init

import (
	"main-service/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitRecommendedCourse(o orm.Ormer, courses []models.Course) []models.RecommendedCourse {
	result := []models.RecommendedCourse{
		models.RecommendedCourse{
			CourseId: &courses[0],
			Link:     "/link_to_course1/",
		},
		models.RecommendedCourse{
			CourseId: &courses[1],
			Link:     "/link_to_course2/",
		},
		models.RecommendedCourse{
			CourseId: &courses[2],
			Link:     "/link_to_course3/",
		},
		models.RecommendedCourse{
			CourseId: &courses[3],
			Link:     "/link_to_course4/",
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
