package db_init

import (
	"data-service-old/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitComment(o orm.Ormer) []models.Comment {
	result := []models.Comment{
		models.Comment{
			Text: "CommentCommentComment1",
		},
		models.Comment{
			Text: "CommentCommentComment2",
		},
		models.Comment{
			Text: "CommentCommentComment3",
		},
		models.Comment{
			Text: "CommentCommentComment4",
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
