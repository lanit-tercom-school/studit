package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitComment(o orm.Ormer) []models.Comment {
	result := []models.Comment,
		{
			Text: "CommentCommentComment1",
		},
		{
			Text: "CommentCommentComment2",
		},
		{
			Text: "CommentCommentComment3",
		},
		{
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
