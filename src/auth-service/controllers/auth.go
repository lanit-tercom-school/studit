package controllers

import (
	"github.com/astaxie/beego"
	"github.com/robbert229/jwt"
	"github.com/astaxie/beego/config"
	"time"
	"encoding/json"
	"auth-service/models"
	"errors"
)

type LoginResponse struct {
	Token           string              `json:"bearer_token"`
	User            *models.MainUserInfo `json:"user"`
	ExpiresIn       string              `json:"exp"`
	PermissionLevel int                 `json:"perm_lvl"`
}

type Usr struct {
	Login string    `json:"login"`
	Password string `json:"password"`
}

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
		beego.Critical("Empty secure token")
		panic("Empty secure token")
	}
	jwtManager = jwt.HmacSha256(r)
}


// basic http sign in with password and login. Func checks login+password combination with same combination in DB
func TryToLogin(login, password string) (user models.User, err error) {
	// create default model
	user = models.User{
		Login: login,
	}

	err = user.Read("login")
	if err != nil {
		return user, errors.New("Can't find User with this login (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Id < 1 {
		return user, errors.New("Bad user ID (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Password != CustomStr(password).ToSHA1() {
		return user, errors.New("Invalid login or password")
	} else {
		return user, nil  // all OK
	}
}

// Login, получение Bearer-token
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *AuthController) Login() {
	var v Usr
	// Парсим тело запроса
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Login error (403):", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
	} else {
		beego.Trace(v.Login, "Try to login")
		// Ищем в бд соответствия
		user, err := TryToLogin(v.Login, v.Password)
		if err != nil {
			beego.Debug("Login error (403):", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		} else {
			beego.Trace(user.Login, "Success login")
			// Правильный логин, выдаём новый токен
			claim := jwt.NewClaim()
			claim.Set("user_id", user.Id)
			claim.Set("perm_lvl", user.PermissionLevel)
			f := time.Now().Add(time.Hour * 3)
			claim.SetTime("exp", f)

			token, err := jwtManager.Encode(claim)
			if err != nil {
				beego.Debug("Encode error (500):", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			}
			// Прикрепляем токен, уровень, время истечения и базовую информацию о пользователе
			sessionResponse := LoginResponse{
				Token: token,
				User: models.GetMainUserInfo(user.Id),
				ExpiresIn: f.Format(time.UnixDate),
				PermissionLevel: user.PermissionLevel,
			}
			beego.Trace(user.Login, "Sent token")
			c.Data["json"] = sessionResponse
		}
	}
	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description Зайти в систему, получить Bearer-token
// @Param   body    body    auth.Usr    true    "Логин и пароль пользователя для входа"
// @Success	200	{object} auth.LoginResponse Description
// @Failure	403	Invalid username or password
// @router / [post]
func (c *AuthController) Post() {
	c.Login()
}