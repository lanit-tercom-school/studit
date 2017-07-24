package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitVideo(o orm.Ormer, lessons []models.Lesson) []models.Video {
	result := []models.Video{
		{
			Id:       1,
			LessonId: &lessons[0],
			Link:     "/link_to_video1/",
		},
		{
			LessonId: &lessons[1],
			Link:     "/link_to_video2/",
		},
		{
			LessonId: &lessons[2],
			Link:     "/link_to_video3/",
		},
		{
			LessonId: &lessons[3],
			Link:     "/link_to_video4/",
		},
		{
			LessonId: &lessons[4],
			Link:     "/link_to_video5/",
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
