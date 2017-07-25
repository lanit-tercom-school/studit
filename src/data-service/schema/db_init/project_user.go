package db_init

import (
	"data-service/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitProjectUser(o orm.Ormer, projects []models.Project, users []models.User) []models.ProjectUser {
	result := []models.ProjectUser{
		{
			User:     &users[0],
			Project:  &projects[0],
			SignedDate: time.Now(),
			Progress:   0,
		},
		{
			User:     &users[0],
			Project:  &projects[1],
			SignedDate: time.Now(),
			Progress:   0,
		},
		{
			User:     &users[0],
			Project:  &projects[2],
			SignedDate: time.Now(),
			Progress:   0,
		},
		{
			User:     &users[1],
			Project:  &projects[3],
			SignedDate: time.Now(),
			Progress:   0,
		},
		{
			User:     &users[1],
			Project:  &projects[1],
			SignedDate: time.Now(),
			Progress:   0,
		},
		{
			User:     &users[2],
			Project:  &projects[4],
			SignedDate: time.Now(),
			Progress:   0,
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
