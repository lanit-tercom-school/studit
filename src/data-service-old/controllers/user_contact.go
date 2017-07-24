package controllers

import (
	"data-service-old/models"
	"strconv"

	//"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Доступ к контактам пользователей
type UserContactController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *UserContactController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create UserContact
// @Param   body            body    []models.UserContactInput   true    "Тело запроса, должен быть массив, то что в примере в []"
// @Param   Bearer-token    header  string                      true    "Токен"
// @Success 201 {int} models.UserContact
// @Failure 403 body is empty
// @router / [post]
func (c *UserContactController) Post() {
	if c.CurrentUser.UserId != VIEWER {
		cUser := models.User{Id: c.CurrentUser.UserId}
		v := []models.UserContactInput{}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			ContId := []string{} // Artem skazal array
			for _, element := range v {
				if models.IsValidContactType(element.Type) {
					in := models.ContactTranslate(&element)
					in.UserId = &cUser
					if _, err := models.AddUserContact(&in); err == nil {
						c.Ctx.Output.SetStatus(HTTP_CREATED)
						ContId = append(ContId, element.Type+": "+element.Contact+" was added.")
					} else {
						ContId = append(ContId, "Fail to add "+element.Type+": "+element.Contact)
					}
				} else {
					ContId = append(ContId, "Fail to add "+element.Type+": "+element.Contact)
				}
			}
			c.Data["json"] = ContId
			beego.Trace("Post user contact OK")
		} else {
			beego.Debug("Unmarshal error for user contact `Post`", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		}
	} else {
		beego.Debug("Access denied for user contact `Post`")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR
	}
	c.ServeJSON()
}
func (c *UserContactController) GetOne() {}

// GetAll ...
// @Title Get One
// @Description get UserContact
// @Param   id              path    string  true    "АйДи пользователя, чьи контакты нужно получить"
// @Param   Bearer-token    header  string  true    "Токен администратора или мастера проекта"
// @Success 200 {object} models.UserContact
// @Failure 403 Forbidden
// @router /:id [get]
func (c *UserContactController) GetAll() {
	if c.CurrentUser.PermissionLevel != VIEWER {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("GetAll user contact `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		} else {
			if _, err = models.GetUserById(id); err == nil {
				is_master, err := models.IsProjectMasterForUserById(id, c.CurrentUser.UserId)
				if err == nil {
					if c.CurrentUser.UserId == id || c.CurrentUser.PermissionLevel == 2 || is_master {
						v, err := models.GetAllUserContacts(id)
						if err != nil {
							beego.Debug("Error in `GetAllUserContacts` in user contact `GetAll`", err.Error())
							c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
							c.Data["json"] = err.Error()
						} else {
							beego.Trace("GetAll user contact OK")
							c.Data["json"] = v
						}
					} else {
						beego.Debug("Access denied for user contact `GetAll`")
						c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
						c.Data["json"] = HTTP_FORBIDDEN_STR
					}
				} else {
					beego.Debug("Error in `IsProjectMasterForUserById` in user contact `GetAll`", err.Error())
					c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
					c.Data["json"] = err.Error()
				}
			} else {
				beego.Debug("Does not exist user contact `GetAll`")
				c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
				c.Data["json"] = "User Not Found"
			}
		}
	} else {
		beego.Debug("Access denied for user contact `GetAll`")
		c.Ctx.Output.SetStatus(HTTP_UNAUTHORIZED)
		c.Data["json"] = HTTP_UNAUTHORIZED_STR
	}
	c.ServeJSON()
}

/*
// Put ...
// @Title Put
// @Description update the UserContact
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.UserContact	true		"body for UserContact content"
// @Success 200 {object} models.UserContact
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserContactController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.UserContact{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserContactById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
*/
// Delete ...
// @Title Delete
// @Description delete the UserContact
// @Param   id              path        string  true    "АйДи контакта, который требуется удалить"
// @Param   Bearer-token    header      string  true    "Токен"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserContactController) Delete() {
	if c.CurrentUser.PermissionLevel != VIEWER {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err == nil {
			v, err := models.GetUserContactById(id)
			if err != nil {
				beego.Debug("Delete user contact `GetUserNyId` error")
				c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
				c.Data["json"] = err.Error()
			} else {
				if v.UserId.Id == c.CurrentUser.UserId {
					if err := models.DeleteUserContact(v.Id); err == nil {
						beego.Trace("Delete user contact OK")
						c.Data["json"] = HTTP_OK_STR
					} else {
						beego.Debug("Delete user contact error", err.Error())
						c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
						c.Data["json"] = err.Error()
					}
				} else {
					beego.Debug("Access denied for user contact `Delete`")
					c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
					c.Data["json"] = HTTP_FORBIDDEN_STR
				}
			}
		} else {
			beego.Debug("Delete user contact `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		}
	} else {
		beego.Debug("Unauthorized user in user contact `Delete`")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR
	}
	c.ServeJSON()
}
