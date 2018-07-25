package controllers

import (
	"auth-service/models"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
)

// Изменить пароль пользователя
type ChangePasswordController struct {
	ControllerWithAuthorization
}

type ChangePasswordJson struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}

func TryToChangePassword(user_id int, password string) (user *models.User, err error) {
	// create default model
	user, err = models.GetUserById(user_id)

	if err != nil {
		return user, errors.New("Can't find User with this login (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Id < 1 {
		return user, errors.New("Bad user ID (dev)") // TODO: should be changed to "Invalid login or password"
	} else if user.Password != CustomStr(password).ToSHA1() {
		return user, errors.New("Invalid login or password")
	} else {
		return user, nil // all OK
	}
}

func ChangePasswordForUser(user *models.User) error {
	user.Password = CustomStr(user.Password).ToSHA1()
	return user.Update("password")
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
		c.Data["json"] = MakeMessageForSending(err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	} else {
		user, err := TryToChangePassword(c.CurrentUser.UserId, v.OldPassword)
		if err != nil {
			beego.Critical(c.Ctx.Input.IP(), "Change password in `TryToChangePassword` error:", err.Error())
			c.Data["json"] = MakeMessageForSending(err.Error())
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		} else {
			user.Password = v.NewPassword
			if len(user.Password) < 6 {
				c.Data["json"] = MakeMessageForSending(HTTP_NOT_ACCEPTABLE_STR)
				c.Ctx.Output.SetStatus(HTTP_NOT_ACCEPTABLE)
			} else {
				err := ChangePasswordForUser(user)
				if err != nil {
					beego.Critical(c.Ctx.Input.IP(), "Change password in `ChangePasswordForUser` error:", err.Error())
					c.Data["json"] = MakeMessageForSending(err.Error())
					c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				} else {
					beego.Trace("Password was changed")
					c.Data["json"] = MakeMessageForSending("Ok")
				}
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
		c.Data["json"] = MakeMessageForSending(HTTP_FORBIDDEN_STR)
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
	} else {
		c.ChangePassword()
	}
	c.ServeJSON()
}
