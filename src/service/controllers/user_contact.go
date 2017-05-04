package controllers

import (
	"service/models"
	"strconv"

	//"github.com/astaxie/beego"
	"encoding/json"
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
// @Param	body		body 	[]models.UserContactInput	true		"Тело запроса, должен быть массив, то что в примере в []"
// @Param   Bearer-token        header      string          true    "Токен"
// @Success 201 {int} models.UserContact
// @Failure 403 body is empty
// @router / [post]
func (c *UserContactController) Post() {
	if c.CurrentUser.UserId != -1 {
		cUser := models.User{ Id: c.CurrentUser.UserId, }
		v := []models.UserContactInput{}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			ContId := []string{}// Artem skazal aray
			for _,element := range v {
				if models.IsValidContactType(element.Type) {
					in := models.ContactTranslate(&element)
					in.UserId = &cUser
					if _, err := models.AddUserContact(&in); err == nil {
						c.Ctx.Output.SetStatus(201)
						ContId = append(ContId, element.Type+": "+element.Contact+" was added.")
					} else {
						ContId=append(ContId,  "Fail to add " +element.Type+": "+element.Contact)
					}
				} else {
					ContId = append(ContId, "Fail to add "+element.Type+": "+element.Contact)
				}
			}
			c.Data["json"] = ContId
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Unauthorized"
	}
	c.ServeJSON()
}
func (c *UserContactController) GetOne() {}


// GetAll ...
// @Title Get One
// @Description get UserContact
// @Param	id		path 	string	true		"АйДи пользователя, чьи контакты нужно получить"
// @Param   Bearer-token        header      string          true    "Токен"
// @Success 200 {object} models.UserContact
// @Failure 403 Forbidden
// @router /:id [get]
func (c *UserContactController) GetAll() {
	if c.CurrentUser.PermissionLevel!=-1 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		} else {
			//TODO add PROJECT_MASTER
			if _,err = models.GetUserById(id); err==nil {
				if c.CurrentUser.UserId == id || c.CurrentUser.PermissionLevel == 2 {
					v, err := models.GetAllUserContacts(id)
					if err != nil {
						c.Ctx.Output.SetStatus(400)
						c.Data["json"] = err.Error()
					} else {
						c.Ctx.Output.SetStatus(200)
						c.Data["json"] = v
					}
				} else {
					c.Ctx.Output.SetStatus(403)
					c.Data["json"] = "Forbidden"
				}
			} else {
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = "User Not Found"
			}
		}
	} else {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Unauthorized"
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
// @Param	id		path 	string	true		"АйДи контакта, который требуется удалить"
// @Param   Bearer-token        header      string          true    "Токен"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserContactController) Delete() {
	if c.CurrentUser.PermissionLevel != -1 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err == nil {
			v, err := models.GetUserContactById(id)
			if err != nil {
				c.Ctx.Output.SetStatus(400)
				c.Data["json"] = err.Error()
			} else {
				if v.UserId.Id == c.CurrentUser.UserId {
					if err := models.DeleteUserContact(v.Id); err == nil {
						c.Ctx.Output.SetStatus(200)
						c.Data["json"] = "OK"
					} else {
						c.Ctx.Output.SetStatus(400)
						c.Data["json"] = err.Error()
					}
				} else {
					c.Ctx.Output.SetStatus(403)
					c.Data["json"] = "Forbidden"
				}
			}
		} else {
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Unauthorized"
	}
	c.ServeJSON()
}
