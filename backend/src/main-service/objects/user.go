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

/* ChangeUser получает из Data Service текущего пользователя
изменяет нужное поле ползователя */
func ChangeUser(p gql.ResolveParams, paramName string) (interface{}, error) {
	new := helpers.InterfaceToString(p.Args["New"])
	token := helpers.InterfaceToString(p.Context.Value("Token"))

	user := User{}
	id := strconv.Itoa(p.Context.Value("CurrentUser").(CurrentClient).UserId)
	url := conf.Configuration.DataServiceURL + "v1/user/" + id

	helpers.LogGet(url, "Getting user by id and token")
	err := helpers.HttpGetWithToken(url, token, &user)

	if err != nil {
		helpers.LogErrorGet(url, err)
		return nil, err
	}

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

// UpdateUserOnServer загружает пользователя в DataService
func UpdateUserOnServer(p gql.ResolveParams, user User) (interface{}, error) {
	token := helpers.InterfaceToString(p.Context.Value("Token"))

	id := p.Context.Value("CurrentUser").(CurrentClient).UserId

	url := conf.Configuration.DataServiceURL + "v1/user/" + strconv.Itoa(id)

	message := Message{}

	helpers.LogPut(url, "Putting user into data by token")
	err := helpers.HttpPutWithToken(url, token, user, &message)

	if err != nil {
		helpers.LogErrorPut(url, err)
		return nil, err
	}

	return message, err
}

//ResolvePutNewNickname - Смена NickName пользователя
func ResolvePutNewNickname(p gql.ResolveParams) (interface{}, error) {

	tempUser, err := ChangeUser(p, "Nickname")

	if err != nil {
		return nil, err
	}

	user, ok := tempUser.(User)

	if !ok {
		err = errors.New("missed user")
		return nil, err
	}

	message, err := UpdateUserOnServer(p, user)

	return message, err
}

//ResolvePutNewAvatar - Смена Avatar пользователя
func ResolvePutNewAvatar(p gql.ResolveParams) (interface{}, error) {

	tempUser, err := ChangeUser(p, "Avatar")

	if err != nil {
		return nil, err
	}

	user, ok := tempUser.(User)

	if !ok {
		err = errors.New("missed user")
		return nil, err
	}

	message, err := UpdateUserOnServer(p, user)

	return message, err
}

//ResolvePutNewDescription - Смена Description пользователя
func ResolvePutNewDescription(p gql.ResolveParams) (interface{}, error) {

	tempUser, err := ChangeUser(p, "Description")

	if err != nil {
		return nil, err
	}

	user, ok := tempUser.(User)

	if !ok {
		err = errors.New("missed user")
		return nil, err
	}

	message, err := UpdateUserOnServer(p, user)

	return message, err
}
