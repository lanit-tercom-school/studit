package objects

import (
	"errors"
	"main-service/conf"
	"main-service/helpers"

	gql "github.com/graphql-go/graphql"
)

func init() {
	//Здесь прописанны поля, которые вызывали ошибку "typechecking loop involving"
	UserType.AddFieldConfig("Enrolls", &gql.Field{Type: gql.NewList(ProjectEnrollType), Resolve: ResolveGetEnrollsByUser})
	UserType.AddFieldConfig("ProjectOn", &gql.Field{Type: gql.NewList(ProjectOnType), Resolve: ResolveGetProjectOnByUser})
}

//User - используется для получения пользователя с data-service
type User struct {
	Id          int
	Nickname    string
	Description string
	Avatar      string
}

//UserType - grqphql объект пользователя
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
		},
	},
)

//ResolveGetUserById - Получение пользователя по Id с data-service
func ResolveGetUserById(p gql.ResolveParams) (interface{}, error) {
	var id string
	id, ok := p.Args["Id"].(string)
	if !ok {
		err := errors.New("missed id")
		return nil, err
	}
	var user User
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/user/"+id, &user)
	return user, err
}
