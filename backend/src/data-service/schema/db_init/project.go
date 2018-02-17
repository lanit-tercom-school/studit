package db_init

import (
	"data-service/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitProject(o orm.Ormer) []models.Project {
	result := []models.Project{
		{
			Name:        "Образовательный портал Studit",
			Description: "Разработка образовательного портала для Lanit-Tercom School",
			Logo:        "http://www.calorizator.ru/sites/default/files/imagecache/product_512/product/juice-10.jpg",
			Tags:        []string{"studit", "summerschool"},
			Status:      "started",
			GitHubUrl:   "https://github.com/eyshella/test_for_studit",
		},
		{
			Name:        "Модный фрилансер",
			Description: "Какие же стрелочки вокруг ноубука!",
			Logo:        "http://findfood.ru/attaches/product/bezalkogolnyie-napitki/apelsinovyj-sok.jpg",
			Tags:        []string{"freelance"},
			Status:      "opened",
			GitHubUrl:   "https://github.com/eyshella/test_for_studit",
		},
		{
			Name:        "Оригинальное название",
			Description: "Click-bait описание",
			Logo:        "http://www.calorizator.ru/sites/default/files/imagecache/product_512/product/juice-35.jpg",
			Tags:        []string{"creative"},
			Status:      "opened",
			GitHubUrl:   "https://github.com/eyshella/test_for_studit",
		},
		{
			Name:        "TFS Mobile",
			Description: "Студенческий проект TFS Mobile",
			Logo:        "http://www.calorizator.ru/sites/default/files/imagecache/product_512/product/juice-20.jpg",
			Tags:        []string{"TFS", "summerschool"},
			Status:      "opened",
			GitHubUrl:   "https://github.com/eyshella/test_for_studit",
		},
		{
			Name:        "Еще один проект",
			Description: "Описаниеописаниеописание",
			Logo:        "http://www.calorizator.ru/sites/default/files/imagecache/recipes_full/recipe/57829.jpg",
			Tags:        []string{"project"},
			Status:      "opened",
			GitHubUrl:   "https://github.com/eyshella/test_for_studit",
		},
	}

	for i := 0; i < len(result); i++ {
		id, err := o.Insert(&result[i])
		if err != nil {
			beego.Critical(err.Error())
			panic(err)
		}
		result[i].Id = int(id)
		o.Raw("UPDATE project SET tags = ? WHERE id = ?",
			"{\""+strings.Join(result[i].Tags[:], "\", \"")+"\"}", result[i].Id).Exec()
	}

	return result
}
