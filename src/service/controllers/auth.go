package controllers

import (
	"encoding/json"
	_ "service/models"

	"github.com/astaxie/beego"
	"service/auth"
	"github.com/robbert229/jwt"
	"time"
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
	var v auth.Usr
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := auth.TryToLogin(v.Login, v.Password)
		if err != nil {
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

			sessionResponse := auth.SessionStruct{
				Token: token,
				UserId: user.Id,
				ExpiresIn: f.Format(time.UnixDate),
			}

			c.Data["json"] = sessionResponse
		}
	} else {
		c.Data["json"] = "Wrong request"
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
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(500) // TODO: change to 400?
			} else {
				c.Data["json"] = "OK"
			}

		} else {
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(500) // TODO: change to 400?
		}

	} else {
		c.Data["json"] = "Empty token"
		c.Ctx.Output.SetStatus(400)
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
// @Description logout current token
// @Param	token	query	string	false	"Token To Logout"
// @Failure 200 OK
// @Failure 403 Wrong token
// @Failure 400 Empty token
// @router / [get]
func (c *LogoutController) GetAll() {
	c.Logout()
}


type ControllerWithAuthorization struct {
	beego.Controller
}

// Наследовать для контроллеров, требующие валидации юзера
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

