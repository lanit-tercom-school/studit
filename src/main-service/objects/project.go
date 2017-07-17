package objects

import (
	"errors"
	"main-service/conf"
	"strconv"
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

//ProjectEnroll - используется для получения заявок с data-service
type ProjectEnroll struct {
	Id             int       `json:"Id"`
	Project        Project   `json:"Project"`
	User           User      `json:"User"`
	Message        string    `json:"EnrollingMessage"`
	DateOfCreation time.Time `json:"Time"`
}

//ProjectUser - используется для получения связи пользователей и проектов с data-service
type ProjectUser struct {
	Id         int       `json:"Id"`
	Project    Project   `json:"Project"`
	User       User      `json:"User"`
	SignedDate time.Time `json:"SignedDate"`
	Progress   int       `json:"Progress"`
}

//ProjectOn - используется при получении поля ProjectOn пользователя
type ProjectOn struct {
	Project Project
	Enrolls []ProjectEnroll
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

//ProjectEnrollType - grqphql объект заявки на проект
var ProjectEnrollType = gql.NewObject(
	gql.ObjectConfig{
		Name: "ProjectEnroll",
		Fields: gql.Fields{
			"Project": &gql.Field{
				Type: ProjectType,
			},
			"User": &gql.Field{
				Type: UserType,
			},
			"Message": &gql.Field{
				Type: gql.String,
			},
			"DateOfCreation": &gql.Field{
				Type: gql.String,
			},
		},
	},
)

//ProjectOnType - grqphql объект связи пользователя и проекта
var ProjectOnType = gql.NewObject(
	gql.ObjectConfig{
		Name: "ProjectOnType",
		Fields: gql.Fields{
			"Project": &gql.Field{
				Type: ProjectType,
			},
			"Enrolls": &gql.Field{
				Type:    gql.NewList(ProjectEnrollType),
				Resolve: ResolveGetEnrollsByProjectOn,
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
	err := httpGet(conf.Configuration.DataServiceURL+"v1/project/"+id, &project)
	return project, err
}
func ResolveGetEnrollsByProjectOn(p gql.ResolveParams) (interface{}, error) {
	if p.Context.Value("CurrentUser").(CurrentClient).PermissionLevel >= LEADER {
		projectOn := p.Source.(ProjectOn)
		var projectEnrolls []ProjectEnroll
		err := httpGet(conf.Configuration.DataServiceURL+"v1/project_enroll/?query=Project:"+strconv.Itoa(projectOn.Project.Id), &projectEnrolls)
		return projectEnrolls, err
	}
	return nil, errors.New("Access is denied")
}
func ResolveGetEnrollsByUser(p gql.ResolveParams) (interface{}, error) {
	u := p.Source.(User)
	var projectEnrolls []ProjectEnroll
	err := httpGet(conf.Configuration.DataServiceURL+"v1/project_enroll/?query=User:"+strconv.Itoa(u.Id), &projectEnrolls)
	return projectEnrolls, err
}

func ResolveGetProjectOnByUser(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	u := p.Source.(User)
	if c.PermissionLevel == ADMIN || c.UserId == u.Id {
		var projectUsers []ProjectUser
		var projectOn ProjectOn
		var projectOns []ProjectOn
		err := httpGet(conf.Configuration.DataServiceURL+"v1/project_user/?query=User:"+strconv.Itoa(u.Id), &projectUsers)
		for _, v := range projectUsers {
			projectOn.Project = v.Project
			projectOns = append(projectOns, projectOn)
		}

		return projectOns, err
	}
	return nil, errors.New("Access is denied")
}
