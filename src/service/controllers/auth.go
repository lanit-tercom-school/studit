package controllers

import (
	"encoding/json"
	_ "service/models"

	"github.com/astaxie/beego"
	"service/auth"
)

var sessionName = beego.AppConfig.String("SessionName")
var sessionLifeTime, sessionLifeTimeErr = beego.AppConfig.Int64("SessionGCMaxLifetime")

// Login, Logout, регистрация и восстановление пароля, получение token
type AuthController struct {
	beego.Controller
}

func init() {
	if sessionLifeTimeErr != nil {
		beego.Critical(sessionLifeTimeErr.Error())
	}
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
	var v auth.Usr
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := auth.TryToLogin(v.Login, v.Password)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.ResponseWriter.WriteHeader(403)
		} else {
			// success, register new session
			sess := c.StartSession()
			//TODO: was in example, but didn't work
			// defer sess.SessionRelease()

			sessionResponse := auth.SessionStruct{
				Token: sess.SessionID(),
				UserId: user.Id,
				ExpiresIn: sessionLifeTime,
			}

			user_id := sess.Get(sessionName)
			if user_id == nil {
				c.SetSession(sessionName, user.Id)
			}

			c.Data["json"] = sessionResponse
		}
	} else {
		c.Data["json"] = err.Error()
		c.Ctx.ResponseWriter.WriteHeader(403)
	}
	c.ServeJSON()
}


func (c *LogoutController) Logout() {
	// TODO: тут какие-то костыли
	userToken := c.GetString("token")
	if userToken != "" {
		sess := c.StartSession()
		if userToken == sess.SessionID() {
			sess.Delete(sessionName)
			c.DestroySession()
			c.Data["json"] = SuccessResponse{"OK"}
		} else {
			c.Data["json"] = ErrorResponse{"Wrong token"}
		}
	} else {
		c.Data["json"] = ErrorResponse{"Empty token"}
	}
	c.ServeJSON()
}


// TODO: добавить нормальные доки
// Post ...
// @Title Post
// @Description login with username and password
// @Param	body		body 	auth.Usr	true ""
// @Failure	200	{object} auth.SessionStruct
// @Failure	403	Invalid username or password
// @router / [post]
func (c *AuthController) Post() {
	c.Login()
}

// TODO: убрать этот костыль
func (c *AuthController) GetOne() {
	response := ErrorResponse{"Not Found"}
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

// TODO: убрать этот костыль
func (c *AuthController) GetAll() {
	response := ErrorResponse{"Method Not Allowed"}
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}

type LogoutController struct {
	beego.Controller
}

// TODO: убрать этот костыль
func (c *LogoutController) GetOne() {
	response := ErrorResponse{"Not Found"}
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description logout current token
// @Param	token	query	string	false	"Token To Logout"
// @Failure 200 {object} controllers.SuccessResponse
// @Failure 400 {object} controllers.ErrorResponse
// @router / [get]
func (c *LogoutController) GetAll() {
	c.Logout()
}