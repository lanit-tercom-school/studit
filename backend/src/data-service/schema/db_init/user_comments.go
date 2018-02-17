package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitUserComment(o orm.Ormer, users []models.User, comments []models.Comment) []models.UserComment {
	result := []models.UserComment{
		{
			Id:        1,
			UserId:    &users[0],
			CommentId: &comments[0],
		},
		{
			UserId:    &users[0],
			CommentId: &comments[1],
		},
		{
			UserId:    &users[1],
			CommentId: &comments[2],
		},
		{
			UserId:    &users[2],
			CommentId: &comments[3],
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
