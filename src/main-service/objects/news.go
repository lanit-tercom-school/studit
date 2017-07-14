package objects

import (
	"errors"
	"time"

	"main-service/conf"

	gql "github.com/graphql-go/graphql"
)

type News struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	DateOfCreation time.Time `json:"created"`
	LastEdit       time.Time `json:"edited"`
	Tags           []string  `json:"tags"`
	Image          string    `json:"image"`
}

var NewsType = gql.NewObject(
	gql.ObjectConfig{
		Name: "News",
		Fields: gql.Fields{
			"Id": &gql.Field{
				Type: gql.String,
			},
			"Title": &gql.Field{
				Type: gql.String,
			},
			"Description": &gql.Field{
				Type: gql.String,
			},
			"DateOfCreation": &gql.Field{
				Type: gql.String,
			},
			"LastEdit": &gql.Field{
				Type: gql.String,
			},
			"Tags": &gql.Field{
				Type: gql.NewList(gql.String),
			},
			"Image": &gql.Field{
				Type: gql.String,
			},
		},
	},
)

func ResolveGetNews(p gql.ResolveParams) (interface{}, error) {
	var id string
	id, ok := p.Args["Id"].(string)
	if !ok {
		err := errors.New("missed id")
		return nil, err
	}
	var news News
	err := httpGet(conf.Configuration.DataServiceURL+"v1/news/"+id, &news)
	return news, err
}
