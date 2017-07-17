package objects

import (
	"errors"
	"main-service/conf"
	"time"

	"strconv"

	gql "github.com/graphql-go/graphql"
)

type User struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Description string `json:"description,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
}

type ProjectEnroll struct {
	Id               int
	Project          Project
	User             User
	EnrollingMessage string
	Time             time.Time
}
type ProjectUser struct {
	Id         int
	Project    Project
	User       User
	SignedDate time.Time
	Progress   int
}

var UserType = gql.NewObject(
	gql.ObjectConfig{
		Name: "User",
		Fields: gql.Fields{
			"Id": &gql.Field{
				Type: gql.String,
			},
			"Nickname": &gql.Field{
				Type: gql.String,
			},
			"Description": &gql.Field{
				Type: gql.String,
			},
			"Avatar": &gql.Field{
				Type: gql.String,
			},
			"Enrolls": &gql.Field{
				Type:    gql.NewList(ProjectType),
				Resolve: ResolveGetEnrollsForUser,
			},
			"Projects": &gql.Field{
				Type:    gql.NewList(ProjectType),
				Resolve: ResolveGetProjectsForUser,
			},
		},
	},
)

func ResolveGetUser(p gql.ResolveParams) (interface{}, error) {
	var id string
	id, ok := p.Args["Id"].(string)
	if !ok {
		err := errors.New("missed id")
		return nil, err
	}
	var user User
	err := httpGet(conf.Configuration.DataServiceURL+"v1/user/"+id, &user)
	return user, err
}

func ResolveGetEnrollsForUser(p gql.ResolveParams) (interface{}, error) {
	u := p.Source.(User)
	var projectEnrolls []ProjectEnroll
	var projects []Project
	err := httpGet(conf.Configuration.DataServiceURL+"v1/project_enroll/?query=User:"+strconv.Itoa(u.Id), &projectEnrolls)
	for _, v := range projectEnrolls {
		projects = append(projects, v.Project)
	}
	return projects, err
}

func ResolveGetProjectsForUser(p gql.ResolveParams) (interface{}, error) {
	u := p.Source.(User)
	var projectUsers []ProjectUser
	var projects []Project
	err := httpGet(conf.Configuration.DataServiceURL+"v1/project_user/?query=User:"+strconv.Itoa(u.Id), &projectUsers)
	for _, v := range projectUsers {
		projects = append(projects, v.Project)
	}
	return projects, err
}
