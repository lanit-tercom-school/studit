package controllers

import (
	"errors"
	"service/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// UserContactController oprations for UserContact
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
// @Description Возвращает models.UserContact если юзер токена совпадает с владельцем контакта, другими словами список контакто доступен только владельцу
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	token		query 	string	false		"User token for access"
// @Failure 200 {object} models.UserContact
// @Failure 403 :id is empty
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
/*
// GetAll ...
// @Title Get All
// @Description get UserContact
// @Param	token	query	string	false	"user token"
// @Failure 200 {object} models.UserContact
// @Failure 403
// @router / [get]*/

func (c *UserContactController) GetAll() {
	userToken := c.GetString("token")
	if userToken != "" {
		sess := c.StartSession()
		userId := sess.Get(sessionName)
		if userId != nil {

		}
		var fields []string
		var sortby []string
		var order []string
		var query = make(map[string]string)
		var limit int64 = 10
		var offset int64

		// fields: col1,col2,entity.col3
		if v := c.GetString("fields"); v != "" {
			fields = strings.Split(v, ",")
		}
		// limit: 10 (default is 10)
		if v, err := c.GetInt64("limit"); err == nil {
			limit = v
		}
		// offset: 0 (default is 0)
		if v, err := c.GetInt64("offset"); err == nil {
			offset = v
		}
		// sortby: col1,col2
		if v := c.GetString("sortby"); v != "" {
			sortby = strings.Split(v, ",")
		}
		// order: desc,asc
		if v := c.GetString("order"); v != "" {
			order = strings.Split(v, ",")
		}
		// query: k:v,k:v
		if v := c.GetString("query"); v != "" {
			for _, cond := range strings.Split(v, ",") {
				kv := strings.SplitN(cond, ":", 2)
				if len(kv) != 2 {
					c.Data["json"] = errors.New("Error: invalid query key/value pair")
					c.ServeJSON()
					return
				}
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}

		l, err := models.GetAllUserContact(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = l
		}
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