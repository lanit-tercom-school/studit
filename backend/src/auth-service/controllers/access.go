package controllers

import (
	"auth-service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/robbert229/jwt"
)

type CurrentClient struct {
	UserId          int
	PermissionLevel int
}

const MaxPermissionLevel int = 2

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
		beego.Critical("400 Bad Request: Empty secure token")
		panic("400 Bad Request: Empty secure token")
	}
	jwtManager = jwt.HmacSha256(r)
}

// Функция проверяет валидность токена и его полей и предоставляет Id пользователя и его уровень допуска
// для следующих методов контроллера
// `token` не должен быть пустой строкой
func newClient(token string) (client CurrentClient) {
	client = CurrentClient{
		UserId:          -1,
		PermissionLevel: models.VIEWER,
	}
	if len(token) < 2 {
		beego.Debug("422 Unprocessable Entity: Token too short")
		return
	} else if err := jwtManager.Validate(token); err == nil {
		claims, _ := jwtManager.Decode(token) // err не нужна, т.к. проверяется во время .Validate()
		userId, err := claims.Get("user_id")
		userPermissionLevel, err1 := claims.Get("perm_lvl")
		if err != nil {
			beego.Critical("400 Bad Request:", err.Error())

		} else if int(userId.(float64)) < 0 {
			beego.Critical("404 Not Found: UserId below 0")

		} else if err1 != nil {
			beego.Critical("403 Forbidden: Can't get perm_lvl ", err1.Error())

		} else if int(userPermissionLevel.(float64)) < 0 {
			beego.Critical("403 Forbidden: UserPermissionLevel below 0")
		} else if int(userPermissionLevel.(float64)) > MaxPermissionLevel {
			beego.Critical("404 Not Found: UserPermissionLevel is", userPermissionLevel, "that higher than", MaxPermissionLevel)

		} else {
			beego.Trace("Success validation for", userId, ", Permission level", userPermissionLevel)
			client.UserId = int(userId.(float64))
			client.PermissionLevel = int(userPermissionLevel.(float64))
		}
	} else {
		beego.Debug("406 Not Acceptable: Invalid token", err.Error())
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
		beego.Debug("406 Not Acceptable: Empty token")
		c.CurrentUser.PermissionLevel = models.VIEWER
		c.Data["json"] = HTTP_UNAUTHORIZED_STR
	} else {
		c.CurrentUser = newClient(clientToken)
	}
	beego.Trace("Exit AUTH Controller")
}
