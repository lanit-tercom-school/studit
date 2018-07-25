package objects

import (
	"errors"
	"log"
	"main-service/conf"
	"main-service/helpers"
	"strconv"
	"time"

	gql "github.com/graphql-go/graphql"
)

func init() {
	//Здесь прописанны поля, которые вызывали ошибку "typechecking loop involving"
	ProjectType.AddFieldConfig("Users", &gql.Field{Type: gql.NewList(UserType), Resolve: ResolveGetUsersByProject})
}

//Project - используется для получения проекта с data-service
type Project struct {
	Id             int       `json:"Id"`
	Name           string    `json:"Name"`
	Description    string    `json:"Description"`
	DateOfCreation time.Time `json:"DateOfCreation"`
	Logo           string    `json:"Logo"`
	Tags           []string  `json:"Tags"`
	Status         string    `json:"Status"`
	GitHubUrl      string    `json:"GitHubUrl"`
}

type ProjectSet struct {
	TotalCount  int
	ProjectList []Project
}

//НЕ ЗАБЫТЬ УДАЛИТЬ!!!
func Log(text string) {
	log.Printf("Log: %v\n", text)
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
			"GitHubUrl": &gql.Field{
				Type: gql.String,
			},
		},
	},
)

var ProjectSetType = gql.NewObject(
	gql.ObjectConfig{
		Name: "ProjectSet",
		Fields: gql.Fields{
			"TotalCount": &gql.Field{
				Type: gql.String,
			},
			"ProjectList": &gql.Field{
				Type: gql.NewList(ProjectType),
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
			GitHubUrl:      helpers.InterfaceToString(p.Args["GitHubUrl"]),
			Tags:           helpers.InterfaceToArrayStrings(p.Args["Tags"]),
		}
		err := helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/project/", projectToSend, &projectToGet)
		user := User{
			Id: c.UserId,
		}
		project := Project{
			Id: projectToGet.Id,
		}
		projectUserToGet := ProjectUser{}
		projectUserToSend := ProjectUser{
			Project:    project,
			User:       user,
			SignedDate: time.Now(),
			Progress:   0,
		}
		helpers.HttpPost(conf.Configuration.DataServiceURL+"v1/project_user/", projectUserToSend, &projectUserToGet)
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
	var set ProjectSet
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/project/?limit="+limit+"&offset="+offset, &set)
	return set, err
}

func ResolveDeleteProject(p gql.ResolveParams) (interface{}, error) {
	c := p.Context.Value("CurrentUser").(CurrentClient)

	id, ok := p.Args["Id"].(int)
	if !ok {
		return nil, errors.New("Missed Id")
	}

	messageToGet := Message{}
	if c.PermissionLevel >= LEADER {
		helpers.LogAccesAllowed("DeleteProjectEnrolls")
		if err_del_enrolls := helpers.HttpDelete(conf.Configuration.DataServiceURL+"v1/project_enroll/?project_id="+strconv.Itoa(id), nil, &messageToGet); err_del_enrolls.Error() != "404" && err_del_enrolls != nil {
			return nil, err_del_enrolls
		}
		//helpers.LogDelete(conf.Configuration.DataServiceURL+"v1/project_enroll/?project_id="+strconv.Itoa(id), err_del_enrolls.Error())

		helpers.LogAccesAllowed("DeleteProjectUsers")
		if err_del_users := helpers.HttpDelete(conf.Configuration.DataServiceURL+"v1/project_user/?project_id="+strconv.Itoa(id), nil, &messageToGet); err_del_users.Error() != "404" && err_del_users != nil {
			return nil, err_del_users
		}
		//helpers.LogDelete(conf.Configuration.DataServiceURL+"v1/project_user/?project_id="+strconv.Itoa(id), err_del_users.Error())
		helpers.LogAccesAllowed("DeleteProject")
		err := helpers.HttpDelete(conf.Configuration.DataServiceURL+"v1/project/"+strconv.Itoa(id), nil, &messageToGet)

		return messageToGet, err
	}

	helpers.LogAccesDenied("DeleteProject")
	return nil, errors.New("Access is denied")
}
