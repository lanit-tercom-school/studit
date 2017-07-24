package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitUserContact(o orm.Ormer, users []models.User, contactTypes []models.ContactType) []models.UserContact {
	result := []models.UserContact{
		{
			UserId:  &users[0],
			Contact: "a@a",
			Type:    &contactTypes[3],
		},
		{
			UserId:  &users[1],
			Contact: "moder@moder.moder",
			Type:    &contactTypes[0],
		},
		{
			UserId:  &users[2],
			Contact: "egorka2003@maaail.ru",
			Type:    &contactTypes[0],
		},
		{
			UserId:  &users[0],
			Contact: "+7-123-456-78-90",
			Type:    &contactTypes[1],
		},
		{
			UserId:  &users[4],
			Contact: "slayer17",
			Type:    &contactTypes[4],
		},
		{
			UserId:  &users[3],
			Contact: "zagad0chnaya",
			Type:    &contactTypes[5],
		},
		{
			UserId:  &users[5],
			Contact: "nekotyanmimimi",
			Type:    &contactTypes[6],
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
