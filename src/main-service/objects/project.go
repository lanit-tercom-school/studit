package objects

import (
	"errors"
	"main-service/conf"
	"time"

	gql "github.com/graphql-go/graphql"
)

type Project struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DateOfCreation time.Time `json:"created"`
	Logo           string    `json:"logo"`
	Tags           string    `json:"tags"`
	Status         string    `json:"status"` // 0 - проект еще не начался, идет набор, и т.д.
	// 1 - проект начался, ведутся лекции, разработка
	// 2 - проект завершен, активность закончена
}

var ProjectType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Project",
		Fields: gql.Fields{
			"Id": &gql.Field{
				Type: gql.String,
			},
			"Name": &gql.Field{
				Type: gql.String,
			},
			"Description": &gql.Field{
				Type: gql.String,
			},
			"DateOfCreation": &gql.Field{
				Type: gql.String,
			},
			"Logo": &gql.Field{
				Type: gql.String,
			},
			"Tags": &gql.Field{
				Type: gql.String,
			},
			"Status": &gql.Field{
				Type: gql.String,
			},
		},
	},
)

func ResolveGetProject(p gql.ResolveParams) (interface{}, error) {
	var id string
	id, ok := p.Args["Id"].(string)
	if !ok {
		err := errors.New("missed id")
		return nil, err
	}
	var project Project
	err := httpGet(conf.Configuration.DataServiceURL+"v1/project/"+id, &project)
	return project, err
}
