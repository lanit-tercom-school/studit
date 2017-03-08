package controllers

import (
	"service/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Доступ к контактам пользователей
type UserContactController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserContactController) URLMapping() {
	//c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}
/*
// Post ...
// @Title Post
// @Description create UserContact
// @Param	body		body 	models.UserContact	true		"body for UserContact content"
// @Success 201 {int} models.UserContact
// @Failure 403 body is empty
// @router / [post]
func (c *UserContactController) Post() {
	var v models.UserContact
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUserContact(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
*/
// TODO: refactor this
// GetOne ...
// @Title Get One
// @Description Возвращает models.UserContact если юзер токена совпадает с владельцем контакта, другими словами, контакт доступен только владельцу
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	token		query 	string	false		"User token for access"
// @Failure 200 {object} models.UserContact
// @Failure 200 {object} controllers.ErrorResponse
// @router /:id [get]
func (c *UserContactController) GetOne() {
	userToken := c.GetString("token")
	if userToken != "" {
		idStr := c.Ctx.Input.Param(":id")
		id, _ := strconv.Atoi(idStr)
		v, err := models.GetUserContactById(id)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			sess := c.StartSession()
			userId := sess.Get(sessionName)
			if userId != v.UserId.Id {
				c.Data["json"] = ErrorResponse{"Forbidden (this contact is not yours) (dev)"} // TODO: change to `Forbidden`
			} else {
				// success
				c.Data["json"] = v
			}
		}
	} else {
		c.Data["json"] = ErrorResponse{"Wrong token (dev)"}  // TODO: change to `Unauthorized`
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get UserContact
// @Param	token	query	string	false	"user token"
// @Failure 200 {object} models.UserContact
// @Failure 403
// @router / [get]
func (c *UserContactController) GetAll() {
	userToken := c.GetString("token")
	if userToken != "" {
		sess := c.StartSession()
		userId := sess.Get(sessionName)
		if userId != nil {
			l, err := models.GetAllUserContacts(userId.(int))
			if err != nil {
				c.Data["json"] = err.Error()  // TODO: change err.Error()
			}
			c.Data["json"] = l
		}
	} else {
		c.Data["json"] = "Wrong token (dev)" // TODO: change this
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

// Delete ...
// @Title Delete
// @Description delete the UserContact
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserContactController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUserContact(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
*/