package db_init

import (
	"main-service/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitProject(o orm.Ormer) []models.Project {
	result := []models.Project{
		models.Project{
			Name:        "Образовательный портал Studit",
			Description: "Разработка образовательного портала для Lanit-Tercom School",
			Logo:        "http://85.143.214.42/files/1.jpg",
			Tags:        []string{"studit", "summerschool"},
			Status:      "started",
		},
		models.Project{
			Name:        "Модный фрилансер",
			Description: "Какие же стрелочки вокруг ноубука!",
			Logo:        "http://85.143.214.42/files/2.jpg",
			Tags:        []string{"freelance"},
			Status:      "opened",
		},
		models.Project{
			Name:        "Оригинальное название",
			Description: "Click-bait описание",
			Logo:        "http://85.143.214.42/files/3.jpg",
			Tags:        []string{"creative"},
			Status:      "opened",
		},
		models.Project{
			Name:        "TFS Mobile",
			Description: "Студенческий проект TFS Mobile",
			Logo:        "http://www.carlthomasiv.com/wp-content/uploads/2012/08/tfs-logo2-318x235.jpg",
			Tags:        []string{"TFS", "summerschool"},
			Status:      "opened",
		},
		models.Project{
			Name:        "Еще один проект",
			Description: "Описаниеописаниеописание",
			Logo:        "https://www.glidetraining.com/wp-content/uploads/2014/03/Microsoft-Office-Project-2013.png",
			Tags:        []string{"project"},
			Status:      "opened",
		},
	}

	for i := 0; i < len(result); i++ {
		id, err := o.Insert(&result[i])
		if err != nil {
			beego.Critical(err.Error())
			panic (err)
		}
		result[i].Id = int(id)
		o.Raw("UPDATE project SET tags = ? WHERE id = ?",
			"{\""+strings.Join(result[i].Tags[:], "\", \"")+"\"}", result[i].Id).Exec()
	}

	return result
}
