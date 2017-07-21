package controllers

import (
	"auth-service/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
)

func init() {
	reset_passwords = make(map[string]struct {
		int
		string
	})
}

// 					   ["login"]struct{"userId", "pass"}
var reset_passwords map[string]struct {
	int
	string
}

func RequestToResetPassword(pass string, usr models.User) error {
	if user, err := models.GetUserByLogin(usr.Login); err == nil {
		reset_passwords[user.Login] = struct {
			int
			string
		}{user.Id, pass}
		return nil
	} else {
		return err
	}
}

func ResetPassword(login, pass, newPassword string) error {
	tuple, err := reset_passwords[login]
	if err && pass == tuple.string {
		u, err := models.GetUserById(tuple.int)
		if err == nil && u.Login == login {
			u.Password = CustomStr(newPassword).ToSHA1()
			u.Update()
			reset_passwords[login] = struct {
				int
				string
			}{tuple.int, GenerateNewToken(6)}
			return nil
		} else {
			if err != nil {
				return err
			} else {
				return errors.New("Wrong login")
			}
		}
	} else {
		return errors.New("Wrong login or pass")
	}
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
	pass := GenerateNewToken(6)
	if err := RequestToResetPassword(pass, u); err == nil {
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
	Login       string `json:"login"`
	Pass        string `json:"pass"` // activation token from email
	NewPassword string `json:"password"`
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
		err := ResetPassword(v.Login, v.Pass, v.NewPassword)
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
