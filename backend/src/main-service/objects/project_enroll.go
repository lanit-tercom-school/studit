package objects

import (
	"errors"
	"main-service/conf"
	"main-service/helpers"
	"strconv"
	"time"

	gql "github.com/graphql-go/graphql"
)

//ProjectEnroll - используется для получения заявок с data-service
type ProjectEnroll struct {
	Id             int       `json:"Id"`
	Project        Project   `json:"Project"`
	User           User      `json:"User"`
	Message        string    `json:"EnrollingMessage"`
	DateOfCreation time.Time `json:"Time"`
}

//ProjectEnrollType - grqphql объект заявки на проект
var ProjectEnrollType = gql.NewObject(
	gql.ObjectConfig{
		Name: "ProjectEnroll",
		Fields: gql.Fields{
			"Id": &gql.Field{
				Type: gql.String,
			},
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

func ResolvePostProjectEnroll(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	user := User{
		Id: p.Args["User"].(int),
	}
	project := Project{
		Id: p.Args["Project"].(int),
	}
	projectEnrollToGet := ProjectEnroll{}
	if c.UserId == user.Id || c.PermissionLevel == ADMIN {
		helpers.LogAccesAllowed("PostProjectEnroll")
		projectEnrollToSend := ProjectEnroll{
			Project:        project,
			User:           user,
			Message:        helpers.InterfaceToString(p.Args["Message"]),
			DateOfCreation: time.Now(),
		}
		err := helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/project_enroll/", projectEnrollToSend, &projectEnrollToGet)
		return projectEnrollToGet, err
	}
	helpers.LogAccesDenied("PostProjectEnroll")
	return nil, errors.New("Access is denied")
}

func ResolveDeleteProjectEnroll(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)
	id, ok := p.Args["Id"].(int)
	if !ok {
		return nil, errors.New("Missed Id")
	}
	projectEnrollToGet := ProjectEnroll{}
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project_enroll/"+strconv.Itoa(id), &projectEnrollToGet)

	if err != nil {
		return nil, err
	}
	messageToGet := Message{}
	if c.UserId == projectEnrollToGet.User.Id || c.PermissionLevel == LEADER {
		helpers.LogAccesAllowed("DeleteProjectEnroll")
		err := helpers.HttpDelete(conf.Configuration.DataServiceURL+"v1/project_enroll/?db_id="+strconv.Itoa(id), nil, &messageToGet)
		return messageToGet, err
	}

	helpers.LogAccesDenied("DeleteProjectEnroll")
	return nil, errors.New("Access is denied")
}
