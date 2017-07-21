package db_init

import (
	"data-service-old/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitContactType(o orm.Ormer) []models.ContactType {
	result := []models.ContactType{
		models.ContactType{
			Type: "email",
		},
		models.ContactType{
			Type: "mobile phone",
		},
		models.ContactType{
			Type: "phone",
		},
		models.ContactType{
			Type: "skype",
		},
		models.ContactType{
			Type: "vk.com",
		},
		models.ContactType{
			Type: "telegram",
		},
		models.ContactType{
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
