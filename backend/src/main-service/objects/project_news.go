package objects

import (
	"errors"
	"strconv"
	"time"

	"main-service/conf"
	"main-service/helpers"

	gql "github.com/graphql-go/graphql"
)

type ProjectNews struct {
	Id          int
	Project_id  int
	Title       string
	Description string
	Created     time.Time
	Edited      time.Time
	Image       string
}

func ResolvePostProjectNews(p gql.ResolveParams) (interface{}, error) {
	helpers.LogAccesAllowed("PostNews")

	projectID, err := strconv.Atoi(helpers.InterfaceToString(p.Args["ProjectID"]))

	if err != nil {
		return nil, errors.New("Invalid project_id")
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
