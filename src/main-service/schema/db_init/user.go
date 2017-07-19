package db_init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"main-service/models"
	"main-service/auth"
	"fmt"
)

func InitUser(o orm.Ormer) ([]models.User) {
	result := []models.User{
		models.User{
			Nickname: "Admin",
			Avatar: InitAvatar(),
			Description: "Главный по тарелкам",
		},
		models.User{
			Nickname: "Moderator",
			Avatar: InitAvatar(),
			Description: "Главный по молоткам",
		},
		models.User{
			Nickname: "Егорка2003",
			Avatar: InitAvatar(),
			Description: "ЮЮЮ, ААА",
		},
		models.User{
			Nickname: "ZagadOchNayA",
			Avatar: InitAvatar(),
			Description: "Легко потерять, невозможно забить",
		},
		models.User{
			Nickname: "S1ayeR1",
			Avatar: InitAvatar(),
			Description: "bjklknufu",
		},
		models.User{
			Nickname: "NekoTyan",
			Avatar: InitAvatar(),
			Description: "^_^",
		},
		models.User{
			Nickname: "B",
			Avatar: InitAvatar(),
			Description: "BBB",
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

func InitAvatar() string {
	return fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath,
		auth.GenerateNewToken(6), auth.GenerateRandomColor(), "FFFFFF", auth.AvatarTemplateSize)
}