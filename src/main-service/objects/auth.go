package objects

import (
	"errors"
	"main-service/helpers"
	"strconv"
	"strings"
	"time"

	"main-service/conf"

	gql "github.com/graphql-go/graphql"
	"github.com/robbert229/jwt"
)

type SigninDataToSend struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type SigninDataToGet struct {
	Token            string    `json:"bearer_token"`
	User             User      `json:"user"`
	DataOfExpiration time.Time `json:"exp"`
	PermissionLevel  int       `json:"perm_lvl"`
}
type SignupDataToSend struct {
	Login    string `json:"login"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
type SignupDataToGet struct {
	ActivationCode string `json:"code"`
}

type PasswordDataToSend struct {
	Old string `json:"old"`
	New string `json:"new"`
}

var SignupDataType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Singup",
		Fields: gql.Fields{
			"ActivationCode": &gql.Field{
				Type: gql.String,
			},
		},
	},
)

var SigninDataType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Signin",
		Fields: gql.Fields{
			"Token": &gql.Field{
				Type: gql.String,
			},
			"User": &gql.Field{
				Type:    UserType,
				Resolve: ResolveGetUserBySigninData,
			},
			"DataOfExpiration": &gql.Field{
				Type: gql.String,
			},
			"PermissionLevel": &gql.Field{
				Type: gql.Int,
			},
		},
	},
)

func ResolveGetSignupDataByLoginPasswordNickname(p gql.ResolveParams) (interface{}, error) {
	var login, pass, nick string
	var ok bool
	login, ok = p.Args["Login"].(string)
	if !ok {
		err := errors.New("missed login")
		return nil, err
	}
	pass, ok = p.Args["Password"].(string)
	if !ok {
		err := errors.New("missed password")
		return nil, err
	}
	nick, ok = p.Args["Nickname"].(string)
	if !ok {
		err := errors.New("missed nickname")
		return nil, err
	}
	sendData := SignupDataToSend{
		Login:    login,
		Password: pass,
		Nickname: nick,
	}
	getData := SignupDataToGet{}

	err := helpers.HttpPost(conf.Configuration.AuthServiceURL+"v1/signup/", sendData, &getData)
	return getData, err
}
func ResolveGetActivationDataByCode(p gql.ResolveParams) (interface{}, error) {
	var code string
	var ok bool
	code, ok = p.Args["ActivationCode"].(string)
	if !ok {
		err := errors.New("missed code")
		return nil, err
	}
	getData := Message{}

	err := helpers.HttpGet(conf.Configuration.AuthServiceURL+"v1/signup/?pass="+code, &getData)
	return getData, err
}

func ResolveGetSigninDataByLoginAndPassword(p gql.ResolveParams) (interface{}, error) {
	var login, pass string
	var ok bool
	login, ok = p.Args["Login"].(string)
	if !ok {
		err := errors.New("missed login")
		return nil, err
	}
	pass, ok = p.Args["Password"].(string)
	if !ok {
		err := errors.New("missed password")
		return nil, err
	}
	sendData := SigninDataToSend{
		Login:    login,
		Password: pass,
	}
	getData := SigninDataToGet{}

	err := helpers.HttpPost(conf.Configuration.AuthServiceURL+"v1/signin", sendData, &getData)
	return getData, err
}

func ResolveGetUserBySigninData(p gql.ResolveParams) (interface{}, error) {
	id := p.Source.(SigninDataToGet).User.Id
	var user User
	err := helpers.HttpGet(conf.Configuration.DataServiceURL+"v1/user/"+strconv.Itoa(id), &user)
	return user, err
}
func ResolvePutNewPassword(p gql.ResolveParams) (interface{}, error) {
	token := helpers.InterfaceToString(p.Context.Value("Token"))
	new := helpers.InterfaceToString(p.Args["New"])
	old := helpers.InterfaceToString(p.Args["Old"])
	passwordDataToSend := PasswordDataToSend{
		New: new,
		Old: old,
	}
	message := Message{}
	err := helpers.HttpPutWithToken(conf.Configuration.AuthServiceURL+"v1/change/", token, passwordDataToSend, &message)
	return message, err
}

//********************************************
//Дальше идёт код необходимый для проверки токена
//********************************************
type CurrentClient struct {
	UserId          int
	PermissionLevel int
}

const (
	VIEWER             = iota - 1 // незарегистрированный пользователь
	USER               = 0        // обычный пользователь
	LEADER             = 1        // учитель, куратор, меет право создавать проекты
	ADMIN              = 2        // администратор, может всё
	MaxPermissionLevel = 2
)

var jwtManager jwt.Algorithm

func init() {
	helpers.LogAuth("Configuration")

	if conf.Configuration.JwtSecret == "" {
		helpers.LogErrorAuth(errors.New("Empty secure token"))
	} else {
		jwtManager = jwt.HmacSha256(conf.Configuration.JwtSecret)
		helpers.LogAuth("Configuration successfully")
	}
}

// Функция проверяет валидность токена и его полей и предоставляет Id пользователя и его уровень допуска
// для следующих методов контроллера
// `token` не должен быть пустой строкой
func newClient(tokenStr string) (client CurrentClient) {
	client = CurrentClient{
		UserId:          -1,
		PermissionLevel: VIEWER,
	}
	if !strings.HasPrefix(tokenStr, "Bearer ") {
		return
	}
	token := tokenStr[7:]
	if len(token) < 2 {
		helpers.LogErrorAuth(errors.New("Token too short"))
		return
	} else if err := jwtManager.Validate(token); err == nil {
		claims, _ := jwtManager.Decode(token) // err не нужна, т.к. проверяется во время .Validate()
		userId, err := claims.Get("user_id")
		userPermissionLevel, err1 := claims.Get("perm_lvl")
		if err != nil {
			helpers.LogErrorAuth(errors.New("Can't get user_id"))
		} else if int(userId.(float64)) < 0 {
			helpers.LogErrorAuth(errors.New("UserId below 0"))
		} else if err1 != nil {
			helpers.LogErrorAuth(errors.New("Can't get perm_lvl"))
		} else if int(userPermissionLevel.(float64)) < 0 {
			helpers.LogErrorAuth(errors.New("UserPermissionLevel below 0"))
		} else if int(userPermissionLevel.(float64)) > MaxPermissionLevel {
			helpers.LogErrorAuth(errors.New("UserPermissionLevel is " + userPermissionLevel.(string) + " that higher than " + strconv.Itoa(MaxPermissionLevel)))
		} else {
			client.UserId = int(userId.(float64))
			client.PermissionLevel = int(userPermissionLevel.(float64))
			helpers.LogAuth("Success for" + " UserId " + strconv.Itoa(client.UserId) + " PermissionLevel " + strconv.Itoa(client.PermissionLevel))
		}
	} else {
		helpers.LogErrorAuth(errors.New("Invalid token"))
	}
	return
}

// Валидация пользователя
// В функции происходит валидация токена
// Внутри метода требуется проверка (нет, если метод+маршрут общедоступны), какой уровень доступа
// Уровень доступа хранится в `c.PermissionLevel` и изменять его вне этой функции небезопасно
// При условии PermissionLevel == -1 не гарантируется правильный Id
func Auth(clientToken string) (c CurrentClient) {
	c = CurrentClient{}
	helpers.LogAuth("Start")
	if clientToken == "" {
		helpers.LogErrorAuth(errors.New("No token"))
		c.PermissionLevel = VIEWER
	} else {
		c = newClient(clientToken)
	}
	helpers.LogAuth("Exit")
	return
}
