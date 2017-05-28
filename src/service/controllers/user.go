package controllers

import (
	"encoding/json"
	"errors"
	"service/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

// Операции с models.User, для некоторых требуется авторизация
type UserController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param   Bearer-token        header      string          true    "Токен"
// @Success 200 {object} models.User
// @Failure 400 :id is empty string
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err== nil {
		v, err := models.GetUserById(id)
		if err != nil {
			beego.Debug("GetOne user id not found", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		} else {
			if contact, err := models.GetAllUserContacts(id); err != nil {
				beego.Critical("GetOne user Contacts GetAllUserContact error ", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(500)
			} else {
				ismaster, err :=models.IsProjectMasterForUserById(id, c.CurrentUser.UserId)
				if err ==nil {
					if c.CurrentUser.UserId == v.Id || c.CurrentUser.PermissionLevel == 2 || ismaster {
						c.Data["json"] = models.FullUserInfo{
							Id:              v.Id,
							Nickname:        v.Nickname,
							Description:     v.Description,
							Avatar:          v.Avatar,
							PermissionLevel: v.PermissionLevel,
							Contact:         contact,
						}
					} else {
						c.Data["json"] = models.User{
							Id:          v.Id,
							Nickname:    v.Nickname,
							Description: v.Description,
							Avatar:      v.Avatar,
						}
					}
				} else {
					beego.Critical("Error in IsProjectMasterForUserById in User `GetOne(ID)` ", err.Error())
					c.Data["json"] = err.Error()
					c.Ctx.Output.SetStatus(500)
				}
			}
		}
	} else {
		beego.Debug("GetOne user `Atoi` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 400
// @router / [get]
func (c *UserController) GetAll() {
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

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Param   Bearer-token        header      string          true    "Access token, Permission Level should be 2"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	if c.CurrentUser.PermissionLevel != -1 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Put user `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		} else if c.CurrentUser.UserId == id {
			v := models.User{Id: id}
			if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
				v.Id = id
				if err := models.UpdateUserById(&v); err == nil {
					beego.Trace("Put user OK")
					c.Data["json"] = HTTP_OK_STR
				} else {
					beego.Debug("Put user `UpdateUserById` error", err.Error())
					c.Data["json"] = err.Error()
					c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
				}
			} else {
				beego.Debug("Put user `Unmarshal` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			}
		}
	} else {
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Param	token		body	string			false		"admin/moder token"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	// TODO: обновить защиту когда будет лвлинг пользователей
	if c.Ctx.Output.IsOk() {
		idStr := c.Ctx.Input.Param(":id")
		id, _ := strconv.Atoi(idStr)
		if err := models.DeleteUser(id); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}
