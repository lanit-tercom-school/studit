package objects

import (
	"errors"
	"main-service/conf"
	"main-service/helpers"
	"time"

	gql "github.com/graphql-go/graphql"
)

//Project - используется для получения проекта с data-service
type Project struct {
	Id             int       `json:"Id"`
	Name           string    `json:"Name"`
	Description    string    `json:"Description"`
	DateOfCreation time.Time `json:"DateOfCreation"`
	Logo           string    `json:"Logo"`
	Tags           string    `json:"Tags"`
	Status         string    `json:"Status"`
}

//ProjectType - grqphql объект проекта
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

//ResolveGetProjectById - получение проета по Id с data-service
func ResolveGetProjectById(p gql.ResolveParams) (interface{}, error) {
	var id string
	id, ok := p.Args["Id"].(string)
	if !ok {
		err := errors.New("missed id")
		return nil, err
	}
	var project Project
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project/"+id, &project)
	return project, err
}

func ResolvePostProject(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	projectToGet := Project{}
	if c.PermissionLevel >= LEADER {
		helpers.LogAccesAllowed("PostProject")
		projectToSend := Project{
			DateOfCreation: time.Now(),
			Name:           helpers.InterfaceToString(p.Args["Name"]),
			Description:    helpers.InterfaceToString(p.Args["Description"]),
			Logo:           helpers.InterfaceToString(p.Args["Logo"]),
			Status:         helpers.InterfaceToString(p.Args["Status"]),
			Tags:           helpers.InterfaceToString(p.Args["Tags"]),
		}
		err := helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/project/", projectToSend, &projectToGet)
		return projectToGet, err
	}
	helpers.LogAccesDenied("PostProject")
	return nil, errors.New("Access is denied")
}

func ResolveGetProjectList(p gql.ResolveParams) (interface{}, error) {
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
	var project []Project
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project/?limit="+limit+"&offset="+offset, &project)
	return project, err
}
