package objects

import (
	"errors"
	"strconv"
	"time"

	"main-service/conf"
	"main-service/helpers"

	gql "github.com/graphql-go/graphql"
)

// структура новости, которую мы передаем дата-сервису
type ProjectNews struct {
	Id          int
	Project_id  int
	Title       string
	Description string
	Created     time.Time
	Edited      time.Time
	Image       string
}

// Функция, которая отправляет запрос дата-сервису для добавления новости к проекту
// может кидать ошибки:
// "Invalid project id" (если проекта с таким id не существует)
// "Access is denied" (если уровень доступа user-a меньше LEADER-а)4
// ошибку преобразования string к int-у
func ResolvePostProjectNews(p gql.ResolveParams) (interface{}, error) {
	helpers.LogAccesAllowed("PostNews")

	projectID, err := strconv.Atoi(helpers.InterfaceToString(p.Args["ProjectID"]))

	if err != nil {
		return nil, errors.New("Invalid project_id")
	}

	curClient := p.Context.Value("CurrentUser").(CurrentClient)

	if curClient.PermissionLevel < LEADER {
		return nil, errors.New("Access is denied")
	}

	newsToSend := ProjectNews{
		Title:       helpers.InterfaceToString(p.Args["Title"]),
		Project_id:  projectID,
		Description: helpers.InterfaceToString(p.Args["Description"]),
		Image:       helpers.InterfaceToString(p.Args["Image"]),
	}

	message := Message{}
	err = helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/project_news/", newsToSend, &message)
	return message, err
}
