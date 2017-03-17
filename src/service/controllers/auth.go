package controllers

import (
	"encoding/json"
	_ "service/models"

	"github.com/astaxie/beego"
	"service/auth"
	"github.com/vetcher/jwt"
	"time"
	"service/models"
)

var jwtManager = jwt.HmacSha256("Secret")

// Login, Logout, регистрация и восстановление пароля, получение token
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

func (c *AuthController) Login() {
	//c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	//c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var v auth.Usr
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := auth.TryToLogin(v.Login, v.Password)
		if err != nil {
			beego.Debug(err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(403)
		} else {
			// success, register new session
			claim := jwt.NewClaim()
			claim.Set("user_id", user.Id)
			f := time.Now().Add(time.Hour)
			claim.SetTime("exp", f)

			token, err := jwtManager.Encode(claim)
			if err != nil {
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(500)
			}

			sessionResponse := auth.UserAndToken{
				Token: token,
				UserId: user.Id,
				ExpiresIn: f.Format(time.UnixDate),
			}

			c.Data["json"] = sessionResponse
		}
	} else {
		c.Data["json"] = err.Error() // TODO: change to "Wrong request"
		c.Ctx.Output.SetStatus(403)
	}

	c.ServeJSON()
}


func (c *LogoutController) Logout() {
	userToken := c.GetString("token")
	if userToken != "" {
		if err := jwtManager.Validate(userToken); err == nil {
			claims, _ := jwtManager.Decode(userToken)
			_, err := claims.Get("user_id")
			if err != nil {
				c.Data["json"] = err.Error() // TODO: change on production
				c.Ctx.Output.SetStatus(500) // TODO: change to 400?
			} else {
				ban_up_to, err := claims.GetTime("exp")
				if err != nil {
					beego.Critical(err)
					c.Data["json"] = err.Error() // TODO: change on production
					c.Ctx.Output.SetStatus(500) // TODO: change to 400?
					c.ServeJSON()
					c.StopRun()
				}
				jwt.GlobalStorage.Ban(userToken, ban_up_to)
				c.Data["json"] = "OK"
			}

		} else {
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400) // TODO: change to 403?
		}

	} else {
		c.Data["json"] = "Empty token"
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}


type RegistrationController struct {
	beego.Controller
}

func (c *RegistrationController) Register() {
	var u models.User
	err := c.ParseForm(&u)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &u); err != nil {
		c.Data["json"] = err.Error() //TODO: change to "Invalid Form"
		c.Ctx.Output.SetStatus(400)
	}
	if u.Login == "" || u.Password == "" || u.Nickname == "" {
		c.Data["json"] = "Wrong Login/Password/Nickname"
		c.Ctx.Output.SetStatus(400)
	}
	beego.Debug(u)
	pass := auth.GenerateNewToken(15)
	err = auth.NewUser(pass, u)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(500)
	} else {
		// TODO: sent email with `pass`
		c.Data["json"] = pass // TODO:! CHANGE TO "OK" !
	}
	c.ServeJSON()
}

func (c *RegistrationController) Activate() {
	pass := c.GetString("pass")
	if pass != "" {
		err := auth.ActivateUser(pass)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)
		} else {
			c.Data["json"] = "Registred"
		}
	} else {
		c.Data["json"] = "Empty pass"
		c.Ctx.Output.SetStatus(400)
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
// @Param	body	body	models.User	true	"Никнейм, логин(email) и пароль обязательны" ""
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
// @Param	pass	query	string	false	"Pass to activate account"
// @Failure 200 OK
// @router / [get]
func (c *RegistrationController) Get() {
	c.Activate()
}

// TODO: убрать этот костыль
func (c *RegistrationController) GetOne() {
	c.Data["json"] = "Not Found"
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}


// TODO: добавить нормальные доки
// Post ...
// @Title Post
// @Description Запрос: auth.Usr, Ответ: auth.UserAndToken
// @Param	body		body 	auth.Usr	true ""
// @Success	200	{object} auth.UserAndToken
// @Failure	403	Invalid username or password
// @router / [post]
func (c *AuthController) Post() {
	c.Login()
}

// TODO: убрать этот костыль
func (c *AuthController) GetOne() {
	c.Data["json"] = "Not Found"
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

// TODO: убрать этот костыль
func (c *AuthController) GetAll() {
	c.Data["json"] = "Method Not Allowed"
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}

type LogoutController struct {
	beego.Controller
}

// TODO: убрать этот костыль
func (c *LogoutController) GetOne() {
	c.Data["json"] = "Not Found"
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description Осуществляет выход пользователя из системы
// @Param	token	query	string	false	"Token To Logout"
// @Failure 200 OK
// @Failure 403 Wrong token
// @Failure 400 Empty token
// @router / [get]
func (c *LogoutController) GetAll() {
	c.Logout()
}

func (c *ResetPasswordController) URLMapping() {
	c.Mapping("Get", c.Get)
}

type ResetPasswordController struct {
	beego.Controller
}

func (c *ResetPasswordController) ResetPasswordRequest() {
	login := c.GetString("login")
	u := models.User{
		Login: login,
	}
	pass := auth.GenerateNewToken(6)
	if err := auth.RequestToResetPassword(pass, u); err == nil {
		// TODO: sent email with `pass` to reset password
		c.Data["json"] = pass // TODO:! CHANGE TO "OK" !
	} else {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

type ResetPasswordActionJson struct {
	Pass	string	`json:"pass"`
	NewPassword	string	`json:"password"`
}

func (c *ResetPasswordController) ResetPasswordAction() {
	v := ResetPasswordActionJson{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(400)
	} else {

	}
	c.ServeJSON()
}

// TODO добавить верификацию запроса (проверка капчи) для предупреждения ddos
// Put ...
// @Title Reset Password Action
// @Description Сбросить пароль и установить новый
// @Success 200 "OK"
// @router / [put]
func (c *ResetPasswordController) Put() {
	c.ResetPasswordAction()
}

// TODO добавить верификацию запроса (проверка капчи) для предупреждения ddos
// Get ...
// @Title Reset Password Request
// @Description Осуществляет запрос на сброс пароля
// @Param	login	query	string	false	"Token To Logout"
// @Success 200 "OK"
// @router / [get]
func (c *ResetPasswordController) Get() {
	c.ResetPasswordRequest()
}

// TODO: встроить собственную проверку валидации, а не полагаться на Output.Status
type ControllerWithAuthorization struct {
	beego.Controller
}

// Наследовать для контроллеров, требующие валидации юзера
// В функции происходит валидация токена для маршрутов, которые этого требуют
// Внутри метода требуется проверка (нет, если метод+маршрут общедоступны), как прошла валидация токена
//	// Если проверка прошла успешно, то код ответа будет 200
//	if c.Ctx.Output.IsOk() {
//		// доступ разрешен
//	} else {
//		// доступ запрещен, обработка (например, 403 "Forbidden")
//	}
func (c *ControllerWithAuthorization) Prepare() {
	beego.Info("start validation")
	userToken := c.GetString("token")
	if userToken == "" {
		c.Data["json"] = "Wrong token (dev)" // TODO: change to `Unauthorized`
		c.Ctx.Output.SetStatus(400)
	} else {
		if jwtManager.Validate(userToken) == nil {
			claims, _ := jwtManager.Decode(userToken)
			userid, err := claims.Get("user_id")
			if err != nil {
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(500) // TODO: change to 400?
			} else {
				if int(userid.(float64)) < 0 {
					c.Data["json"] = err.Error()
					c.Ctx.Output.SetStatus(500) // TODO: change to 400?
				} else {
					beego.Info("success validation")
					c.Ctx.Output.SetStatus(200)
				}
			}
		} else {
			c.Data["json"] = "Wrong token (dev)" // TODO: change to `Unauthorized`
			c.Ctx.Output.SetStatus(400)
		}
	}
	beego.Info("exit prepare")
}

