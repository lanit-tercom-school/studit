package objects

import (
	"errors"
	"main-service/conf"
	"main-service/helpers"
	"strconv"
	"time"

	gql "github.com/graphql-go/graphql"
)

//ProjectOn - используется при получении поля ProjectOn пользователя
type ProjectOn struct {
	Id      int
	Project Project
	Enrolls []ProjectEnroll
}

//ProjectUser - используется для получения связи пользователей и проектов с data-service
type ProjectUser struct {
	Id         int       `json:"Id"`
	Project    Project   `json:"Project"`
	User       User      `json:"User"`
	SignedDate time.Time `json:"SignedDate"`
	Progress   int       `json:"Progress"`
}

//ProjectOnType - grqphql объект связи пользователя и проекта
var ProjectOnType = gql.NewObject(
	gql.ObjectConfig{
		Name: "ProjectOnType",
		Fields: gql.Fields{
			"Id": &gql.Field{
				Type: gql.String,
			},
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

func ResolveGetProjectOnByUser(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	u := p.Source.(User)
	if c.PermissionLevel == ADMIN || c.UserId == u.Id {
		helpers.LogAccesAllowed("GetProjectOnByUser")
		var projectUsers []ProjectUser
		var projectOn ProjectOn
		var projectOns []ProjectOn
		err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project_user/?query=User:"+strconv.Itoa(u.Id), &projectUsers)
		for _, v := range projectUsers {
			projectOn.Project = v.Project
			projectOn.Id = v.Id
			projectOns = append(projectOns, projectOn)
		}

		return projectOns, err
	}
	helpers.LogAccesDenied("GetProjectOnByUser")
	return nil, errors.New("Access is denied")
}
func ResolveGetEnrollsByProjectOn(p gql.ResolveParams) (interface{}, error) {
	if p.Context.Value("CurrentUser").(CurrentClient).PermissionLevel >= LEADER {
		helpers.LogAccesAllowed("GetEnrollsByProjectOn")
		projectOn := p.Source.(ProjectOn)
		var projectEnrolls []ProjectEnroll
		err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project_enroll/?query=Project:"+strconv.Itoa(projectOn.Project.Id), &projectEnrolls)
		return projectEnrolls, err
	}
	helpers.LogAccesDenied("GetEnrollsByProjectOn")
	return nil, errors.New("Access is denied")
}
func ResolveGetEnrollsByUser(p gql.ResolveParams) (interface{}, error) {
	u := p.Source.(User)
	var projectEnrolls []ProjectEnroll
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project_enroll/?query=User:"+strconv.Itoa(u.Id), &projectEnrolls)
	return projectEnrolls, err
}

func ResolvePostProjectOn(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	user := User{
		Id: p.Args["User"].(int),
	}
	project := Project{
		Id: p.Args["Project"].(int),
	}
	var projectUsersOfRequesting []ProjectUser
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project_user/?query=User:"+strconv.Itoa(c.UserId), &projectUsersOfRequesting)
	if err != nil {
		return nil, err
	}
	var ok bool
	for _, v := range projectUsersOfRequesting {
		if v.User.Id == c.UserId {
			ok = true
			break
		}
	}
	projectUserToGet := ProjectUser{}
	if (c.PermissionLevel == LEADER && ok) || c.PermissionLevel == ADMIN {
		helpers.LogAccesAllowed("PostProjectOn")
		projectUserToSend := ProjectUser{
			Project:    project,
			User:       user,
			SignedDate: time.Now(),
			Progress:   0,
		}
		err := helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/project_user/", projectUserToSend, &projectUserToGet)
		projectOn := ProjectOn{
			Project: projectUserToGet.Project,
			Id:      projectUserToGet.Id,
		}

		return projectOn, err
	}
	helpers.LogAccesDenied("PostProjectOn")
	return nil, errors.New("Access is denied")
}

func ResolveDeleteProjectOn(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	id, ok := p.Args["Id"].(int)
	if !ok {
		return nil, errors.New("Missed Id")
	}

	projectUserToGet := ProjectUser{}
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project_user/"+strconv.Itoa(id), &projectUserToGet)
	if err != nil {
		return nil, err
	}
	messageToGet := Message{}
	var projectUsersOfRequesting []ProjectUser
	err = helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project_user/?query=User:"+strconv.Itoa(c.UserId), &projectUsersOfRequesting)
	if err != nil {
		return nil, err
	}
	for _, v := range projectUsersOfRequesting {
		if v.User.Id == c.UserId {
			ok = true
			break
		}
	}
	if c.UserId == projectUserToGet.User.Id || c.PermissionLevel == ADMIN || (c.PermissionLevel == LEADER && ok) {
		helpers.LogAccesAllowed("DeleteProjectOn")
		err := helpers.HttpDelete(conf.Configuration.DataServiceURL+"v1/project_user/"+strconv.Itoa(id), nil, &messageToGet)
		return messageToGet, err
	}

	helpers.LogAccesDenied("DeleteProjectOn")
	return nil, errors.New("Access is denied")
}
