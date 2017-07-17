package objects

import (
	"errors"
	"log"
	"main-service/helpers"
	"strconv"
	"strings"

	"github.com/astaxie/beego/config"
	"github.com/robbert229/jwt"
)

//********************************************
//Дальше идёт код необходимый для аутенфикации
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
	auth_config, err := config.NewConfig("ini", "conf/auth.conf")
	var r string
	if err != nil {
		log.Panic(err)
	} else {
		// Пробует считать из конфига
		r = auth_config.String("jwt_secret")
	}
	if r == "" {
		helpers.LogErrorAuth(errors.New("Empty secure token"))
	} else {
		jwtManager = jwt.HmacSha256(r)
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
