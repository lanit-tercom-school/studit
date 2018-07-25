package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/robbert229/jwt"
)

const (
	VIEWER = iota - 1 // незарегистрированный пользователь
	USER   = 0        // обычный пользователь
	LEADER = 1        // учитель, куратор, меет право создавать проекты
	ADMIN  = 2        // администратор, может всё
)

type CurrentClient struct {
	UserId          int
	PermissionLevel int
}

const MaxPermissionLevel int = 2

var jwtManager jwt.Algorithm

func init() {
	auth_config, err := config.NewConfig("ini", "conf/auth.conf")
	var r string
	if err != nil {
		beego.Critical(err)
	} else {
		// Пробует считать из конфига
		r = auth_config.String("jwt_secret")
	}
	if r == "" {
		beego.Critical(string(HTTP_BAD_REQUEST) + HTTP_BAD_REQUEST_STR + ":" + "Empty secure token") //400
		panic(string(HTTP_BAD_REQUEST) + HTTP_BAD_REQUEST_STR + ":" + "Empty secure token")          //400
	}
	jwtManager = jwt.HmacSha256(r)
}

// Функция проверяет валидность токена и его полей и предоставляет Id пользователя и его уровень допуска
// для следующих методов контроллера
// `token` не должен быть пустой строкой
func newClient(token string) (client CurrentClient) {
	client = CurrentClient{
		UserId:          -1,
		PermissionLevel: VIEWER,
	}
	if len(token) < 2 {
		beego.Debug(string(HTTP_BAD_REQUEST) + HTTP_BAD_REQUEST_STR + ":" + "Token too short")
		return
	} else if err := jwtManager.Validate(token); err == nil {
		claims, _ := jwtManager.Decode(token) // err не нужна, т.к. проверяется во время .Validate()
		userId, err := claims.Get("user_id")
		userPermissionLevel, err1 := claims.Get("perm_lvl")
		if err != nil {
			beego.Critical(string(HTTP_BAD_REQUEST)+HTTP_BAD_REQUEST_STR+":"+"Can't get user_id", err.Error())

		} else if int(userId.(float64)) < 0 {
			beego.Critical(string(HTTP_NOT_FOUND) + HTTP_NOT_FOUND_STR + ":" + "UserId below 0")

		} else if err1 != nil {
			beego.Critical(string(HTTP_FORBIDDEN)+HTTP_FORBIDDEN_STR+":"+"Can't get perm_lvl", err.Error())

		} else if int(userPermissionLevel.(float64)) < 0 {
			beego.Critical(string(HTTP_FORBIDDEN) + HTTP_FORBIDDEN_STR + ":" + "UserPermissionLevel below 0")

		} else if int(userPermissionLevel.(float64)) > MaxPermissionLevel {
			beego.Critical(string(HTTP_FORBIDDEN)+HTTP_FORBIDDEN_STR+":"+"UserPermissionLevel is", userPermissionLevel, "that higher than", MaxPermissionLevel)

		} else {
			beego.Trace("Success validation for", userId, ", Permission level", userPermissionLevel)
			client.UserId = int(userId.(float64))
			client.PermissionLevel = int(userPermissionLevel.(float64))
		}
	} else {
		beego.Debug(string(HTTP_NOT_ACCEPTABLE)+HTTP_NOT_ACCEPTABLE_STR+":"+"Invalid token", err.Error())
	}
	return
}

// Контроллер с функцией проверки авторизации перед вызовом основных методов Get, Post, etc...
type ControllerWithAuthorization struct {
	beego.Controller
	CurrentUser CurrentClient
}

// Наследовать для контроллеров, требующие валидации юзера
// В функции происходит валидация токена для контроллеров, которые этого требуют
// Внутри метода требуется проверка (нет, если метод+маршрут общедоступны), какой уровень доступа
// Уровень доступа хранится в `c.CurrentUser.PermissionLevel` и изменять его вне этой функции небезопасно
// При условии PermissionLevel == -1 не гарантируется правильный Id
func (c *ControllerWithAuthorization) Prepare() {
	beego.Trace("Start validation")
	clientToken := c.Ctx.Input.Header("Bearer-token")
	if clientToken == "" {
		beego.Debug(string(HTTP_BAD_REQUEST) + HTTP_BAD_REQUEST_STR + ":" + "Empty token")
		c.CurrentUser.PermissionLevel = VIEWER
		c.Data["json"] = HTTP_UNAUTHORIZED_STR
	} else {
		c.CurrentUser = newClient(clientToken)
	}
	beego.Trace("Exit AUTH Controller")
}
