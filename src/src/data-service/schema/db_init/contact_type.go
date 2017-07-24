package db_init

import (
	"data-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitContactType(o orm.Ormer) []models.ContactType {
	result := []models.ContactType{
		{
			Type: "email",
		},
		{
			Type: "mobile phone",
		},
		{
			Type: "phone",
		},
		{
			Type: "skype",
		},
		{
			Type: "vk.com",
		},
		{
			Type: "telegram",
		},
		{
			Type: "viber",
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
