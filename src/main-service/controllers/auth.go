package controllers

import (
	"encoding/json"
	_ "main-service/models"

	"github.com/astaxie/beego"
	"main-service/auth"
	"github.com/robbert229/jwt"
	"time"
	"main-service/models"
	"github.com/astaxie/beego/config"
)


type CurrentClient struct {
	UserId          int
	PermissionLevel int
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
	// Если в конфиге нет, то генерирует непустой секрет. При перезапуске секрет изменится => старые токены протухнут
		r = auth.GenerateNewToken(15)
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
		beego.Debug("Token too short")
		return
	} else if err := jwtManager.Validate(token); err == nil {
		claims, _ := jwtManager.Decode(token) // err не нужна, т.к. проверяется во время .Validate()
		userId, err := claims.Get("user_id")
		userPermissionLevel, err1 := claims.Get("perm_lvl")
		if err != nil {
			beego.Critical("Can't get user_id", err.Error())

		} else if int(userId.(float64)) < 0 {
			beego.Critical("UserId below 0")

		} else if err1 != nil {
			beego.Critical("Can't get perm_lvl", err.Error())

		} else if int(userPermissionLevel.(float64)) < 0 {
			beego.Critical("UserPermissionLevel below 0")

		} else if int(userPermissionLevel.(float64)) > auth.MaxPermissionLevel {
			beego.Critical("UserPermissionLevel is", userPermissionLevel, "that higher than", auth.MaxPermissionLevel)

		} else {
			beego.Trace("Success validation for", userId, ", Permission level", userPermissionLevel)
			client.UserId = int(userId.(float64))
			client.PermissionLevel = int(userPermissionLevel.(float64))
		}
	} else {
		beego.Debug("Invalid token", err.Error())
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
		beego.Debug("Empty token")
		c.CurrentUser.PermissionLevel = models.VIEWER
		c.Data["json"] = HTTP_UNAUTHORIZED_STR
	} else {
		c.CurrentUser = newClient(clientToken)
	}
	beego.Trace("Exit AUTH Controller")
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
	var v auth.Usr
	// Парсим тело запроса
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Login error (403):", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
	} else {
		beego.Trace(v.Login, "Try to login")
		// Ищем в бд соответствия
		user, err := auth.TryToLogin(v.Login, v.Password)
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
			f := time.Now().Add(time.Hour)
			claim.SetTime("exp", f)

			token, err := jwtManager.Encode(claim)
			if err != nil {
				beego.Debug("Encode error (500):", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			}
			// Прикрепляем токен, уровень, время истечения и базовую информацию о пользователе
			sessionResponse := auth.LoginResponse{
				Token: token,
				User: models.MainUserInfo{
					Id: user.Id,
					Nickname: user.Nickname,
					Avatar: user.Avatar,
				},
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

// Регистрация нового пользователя
type RegistrationController struct {
	beego.Controller
}

type RegistrationUserModel struct {
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Nickname string    `json:"nickname"`
}

func (c *RegistrationController) Register() {
	var temp_u RegistrationUserModel

	beego.Trace("Start Register")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &temp_u); err != nil {
		beego.Debug("Can't Unmarshal:", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	} else {
		beego.Trace("Correct Unmarshal")

		if temp_u.Login == "" || temp_u.Password == "" || temp_u.Nickname == "" {
			beego.Debug("login:", temp_u.Login)
			beego.Debug("password:", temp_u.Password)
			beego.Debug("nickname:", temp_u.Nickname)
			c.Data["json"] = "Wrong Login/Password/Nickname"
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		} else {
			u := models.User{
				Login: temp_u.Login,
				Password: temp_u.Password,
				Nickname: temp_u.Nickname,
			}
			beego.Trace("Valid combo login & password & nickname")

			// Создает токен для подтверждения регистрации, который должен отправляться на email
			pass := auth.GenerateNewToken(15)
			beego.Trace(pass)
			err := auth.NewUser(pass, u)
			if err != nil {
				beego.Debug("Register error:" + err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			} else {
				// TODO: sent email with `pass`
				c.Data["json"] = struct {
					Code string `json:"code"`
				}{Code: pass} // TODO:! CHANGE TO "OK" !
				beego.Trace("Register OK")
			}
		}
	}
	c.ServeJSON()
}

func (c *RegistrationController) Activate() {
	pass := c.GetString("pass")
	if pass != "" {
		err := auth.ActivateUser(pass)
		if err != nil {
			beego.Debug("Activation error: " + err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		} else {
			beego.Trace("Activation OK")
			c.Data["json"] = "Registered"
		}
	} else {
		beego.Debug("Empty pass")
		c.Data["json"] = HTTP_BAD_REQUEST_STR
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	}
	c.ServeJSON()
}

func (c *RegistrationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// mock
// Post ...
// @Title Register
// @Description Проводит преварительную регистрацию пользователя, после требуется подтверждение
// @Param	body	body	controllers.RegistrationUserModel	true	"Никнейм, логин(email) и пароль обязательны" ""
// @Success	200 "OK"
// @Failure	400 "This user is already registered"
// @router / [post]
func (c *RegistrationController) Post() {
	c.Register()
}

// mock
// Get ...
// @Title Activate
// @Description Активирует аккаунт
// @Param	pass	query	string	false	"Код для активации аккаунта"
// @Failure 200 OK
// @router / [get]
func (c *RegistrationController) Get() {
	c.Activate()
}


func (c *ResetPasswordController) URLMapping() {
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
}

// Запрос на сброс пароля
type ResetPasswordController struct {
	beego.Controller
}

func (c *ResetPasswordController) ResetPasswordRequest() {
	login := c.GetString("login")
	u := models.User{
		Login: login,
	}
	beego.Info(login, "want to reset password")
	pass := auth.GenerateNewToken(6)
	if err := auth.RequestToResetPassword(pass, u); err == nil {
		// TODO: sent email with `pass` to reset password
		beego.Info("Pass for reset password was sent to", login)
		c.Data["json"] = pass // TODO:! CHANGE TO "OK" !
	} else {
		beego.Info("Wrong resetRequest for", login, ":", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	}
	c.ServeJSON()
}

type ResetPasswordActionJson struct {
	Login       string  `json:"login"`
	Pass        string  `json:"pass"` // activation token from email
	NewPassword string  `json:"password"`
}

func (c *ResetPasswordController) ResetPasswordAction() {
	v := ResetPasswordActionJson{}
	beego.Trace("New password reset")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		beego.Debug("Reset error:", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	} else {
		beego.Trace("Try to reset for", v.Login)
		err := auth.ResetPassword(v.Login, v.Pass, v.NewPassword)
		if err != nil {
			beego.Debug(err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		} else {
			beego.Trace("Password was reset for", v.Login)
			c.Data["json"] = HTTP_OK_STR
		}
	}
	c.ServeJSON()
}

// TODO добавить верификацию запроса (проверка капчи) для предупреждения ddos
// Put ...
// @Title Reset Password Action
// @Description Сбросить пароль и установить новый
// @Param   body    body    controllers.ResetPasswordActionJson true    "Логин пользователя, его пропуск и новый пароль"
// @Success 200 "OK"
// @router / [put]
func (c *ResetPasswordController) Put() {
	c.ResetPasswordAction()
}

// TODO добавить верификацию запроса (проверка капчи) для предупреждения ddos
// Get ...
// @Title Reset Password Request
// @Description Осуществляет запрос на сброс пароля
// @Param   login   query   string  true    "Логин пользователя, который хочет сбросить пароль"
// @Success 200 "OK"
// @router / [get]
func (c *ResetPasswordController) Get() {
	c.ResetPasswordRequest()
}

// Изменить пароль пользователя
type ChangePasswordController struct{
	ControllerWithAuthorization
}

type ChangePasswordJson struct {
	OldPassword string  `json:"old"`
	NewPassword string  `json:"new"`
}

// URLMapping ...
func (c *ChangePasswordController) URLMapping() {
	c.Mapping("Put", c.Put)
}

func (c *ChangePasswordController) ChangePassword() {
	v := ChangePasswordJson{}
	beego.Trace("New password change")
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		beego.Debug("Change error:", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	} else {
		user, err := auth.TryToChangePassword(c.CurrentUser.UserId, v.OldPassword)
		if err != nil {
			beego.Critical(c.Ctx.Input.IP(), "Change password in `TryToChangePassword` error:", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		} else {
			user.Password = v.NewPassword
			err := auth.ChangePasswordForUser(user)
			if err != nil {
				beego.Critical(c.Ctx.Input.IP(), "Change password in `ChangePasswordForUser` error:", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			} else {
				beego.Trace("Password was changed")
				c.Data["json"] = HTTP_OK_STR
			}
		}
	}
}

// Put ...
// @Title Запрос для смены пароля пользователя
// @Description Для изменения пароля
// @Param   body            body    controllers.ChangePasswordJson  true    "Старый и новый пароль"
// @Param   Bearer-token    header  string                          true    "Токен аутентификации"
// @Success 200 "OK"
// @router / [put]
func (c *ChangePasswordController) Put() {
	if c.CurrentUser.PermissionLevel < models.USER {
		beego.Debug("Unregistered user want to change password")
		c.Data["json"] = HTTP_FORBIDDEN_STR
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
	} else {
		c.ChangePassword()
	}
	c.ServeJSON()
}