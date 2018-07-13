package objects

import (
	"errors"
	"log"
	"main-service/conf"
	"main-service/helpers"
	"strconv"

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

//temp function
func Log(text string) {
	log.Printf("Log: %v\n", text)
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

type UserDataToSend struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

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

func ChangeUser(p gql.ResolveParams, paramName string) (interface{}, error) {
	new := helpers.InterfaceToString(p.Args["New"])
	token := helpers.InterfaceToString(p.Context.Value("Token"))

	user := User{}
	id := strconv.Itoa(p.Context.Value("CurrentUser").(CurrentClient).UserId)

	err := helpers.HttpGetWithToken(conf.Configuration.DataServiceURL+"v1/user/"+id, token, &user)

	switch paramName {
	case "Nickname":
		user.Nickname = new
	case "Avatar":
		user.Avatar = new
	case "Description":
		user.Description = new
	default:
		err = errors.New("Invalid param")
		return nil, err
	}

	return user, err
}

//ResolvePutNewNickname - Смена NickName пользователя
func ResolvePutNewNickname(p gql.ResolveParams) (interface{}, error) {

	token := helpers.InterfaceToString(p.Context.Value("Token"))

	user, err := ChangeUser(p, "Nickname")

	message := Message{}
	err = helpers.HttpPutWithToken(conf.Configuration.DataServiceURL+"v1/user/"+strconv.Itoa(p.Context.Value("CurrentUser").(CurrentClient).UserId), token, user, &message)

	return message, err
}
