package objects

import (
	"errors"
	"time"

	"main-service/conf"
	"main-service/helpers"

	gql "github.com/graphql-go/graphql"
)

type News struct {
	Id             int
	Title          string
	Description    string
	DateOfCreation time.Time
	LastEdit       time.Time
	Tags           string
	Image          string
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
				Type: gql.String,
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
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/news/"+id, &news)
	return news, err
}

func ResolvePostNews(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	newsToGet := News{}
	if c.PermissionLevel >= LEADER {
		helpers.LogAccesAllowed("PostNews")
		newsToSend := News{
			DateOfCreation: time.Now(),
			Title:          helpers.InterfaceToString(p.Args["Title"]),
			Description:    helpers.InterfaceToString(p.Args["Description"]),
			Image:          helpers.InterfaceToString(p.Args["Image"]),
			LastEdit:       time.Now(),
		}
		err := helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/news/", newsToSend, &newsToGet)
		return newsToGet, err
	}
	helpers.LogAccesDenied("PostNews")
	return nil, errors.New("Access is denied")
}
