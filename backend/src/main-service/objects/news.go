package objects

import (
	"errors"
	"strconv"
	"time"

	"main-service/conf"
	"main-service/helpers"

	gql "github.com/graphql-go/graphql"
)

type News struct {
	Id          int
	Title       string
	Description string
	Created     time.Time
	Edited      time.Time
	Tags        []string
	Image       string
}

type NewsSet struct {
	TotalCount    int
	FilteredCount int
	NewsList      []News
}

//
//func (set NewsJSONSet) MarshalJSON() ([]byte, error) {
//	return []byte(set.Arr), nil
//}

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
			"Created": &gql.Field{
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

var NewsSetType = gql.NewObject(
	gql.ObjectConfig{
		Name: "NewsSet",
		Fields: gql.Fields{
			"TotalCount": &gql.Field{
				Type: gql.String,
			},
			"FilteredCount": &gql.Field{
				Type: gql.String,
			},
			"NewsList": &gql.Field{
				Type: gql.NewList(NewsType),
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
			Title:       helpers.InterfaceToString(p.Args["Title"]),
			Description: helpers.InterfaceToString(p.Args["Description"]),
			Image:       helpers.InterfaceToString(p.Args["Image"]),
		}
		err := helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/news/", newsToSend, &newsToGet)
		return newsToGet, err
	}
	helpers.LogAccesDenied("PostNews")
	return nil, errors.New("Access is denied")
}

func ResolveGetNewsList(p gql.ResolveParams) (interface{}, error) {
	var limit, offset string
	limit, ok := p.Args["Limit"].(string)
	if !ok {
		err := errors.New("missed Limit")
		return nil, err
	}
	offset, ok = p.Args["Offset"].(string)
	if !ok {
		err := errors.New("missed Offset")
		return nil, err
	}
	var set NewsSet
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/news/?limit="+limit+"&offset="+offset, &set)
	return set, err
}

//Sends "news" object to server
func UpdateNewsOnServer(p gql.ResolveParams, news News) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	if c.PermissionLevel >= LEADER {
		helpers.LogAccesAllowed("UpadateNewsOnServer")

		token := helpers.InterfaceToString(p.Context.Value("Token"))
		url := conf.Configuration.DataServiceURL + "v1/news/" + strconv.Itoa(news.Id)
		message := Message{}
		err := helpers.HttpPutWithToken(url, token, news, &message)

		if err != nil {
			helpers.LogErrorPut(url, err)
			return nil, err
		}

		return message, err
	}

	helpers.LogAccesDenied("UpadateNewsOnServer")
	return nil, errors.New("Access is denied")
}

//Changes one of selected fields in news object and sends it on server
func ChangeNewsField(p gql.ResolveParams, fieldName string) (interface{}, error) {
	helpers.LogPut("ChangeNewsField", "function")
	tempNews, err := ResolveGetNews(p)

	if err != nil {
		err = errors.New("do not get news")
		return nil, err
	}

	news, ok := tempNews.(News)

	if !ok {
		err = errors.New("missed news")
		return nil, err
	}

	new := helpers.InterfaceToString(p.Args["New"])

	switch fieldName {
	case "Title":
		news.Title = new
	case "Description":
		news.Description = new
	case "Image":
		news.Image = new
	default:
		err = errors.New("Invalid field")
		return nil, err
	}

	return UpdateNewsOnServer(p, news)
}

func ResolvePutNewsTitle(p gql.ResolveParams) (interface{}, error) {
	return ChangeNewsField(p, "Title")
}

func ResolvePutNewsDescription(p gql.ResolveParams) (interface{}, error) {
	return ChangeNewsField(p, "Description")
}

func ResolvePutNewsImage(p gql.ResolveParams) (interface{}, error) {
	return ChangeNewsField(p, "Image")
}
