package controllers

import (
	"encoding/json"
	"errors"
	"service/models"
	"strconv"
	"strings"
)

// Операции с models.Project, для некоторых требуется авторизация
type ProjectController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *ProjectController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Project
// @Param	body		body 	models.Project	true		"body for Project content"
// @Param	token		body	string			false		"admin/moder token"
// @Success 201 OK
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectController) Post() {
	// TODO: обновить защиту когда будет лвлинг пользователей
	if c.Ctx.Output.IsOk() {
		var v models.Project
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if _, err := models.AddProject(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = v
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Project by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Project	Get project with specified id
// @Failure 400 :id is wrong
// @router /:id [get]
func (c *ProjectController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = err.Error() // TODO: change to "Wrong project id"
		c.Ctx.Output.SetStatus(400)
	} else {
		v, err := models.GetProjectById(id)
		if err != nil {
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)
		} else {
			c.Data["json"] = v
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Project
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer. Default 10"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.Project Get array of projects filtered with specified filters (wtf this description)
// @Failure 403
// @router / [get]
func (c *ProjectController) GetAll() {
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

	l, err := models.GetAllProject(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Project
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Project	true		"body for Project content"
// @Param	token		query	string			false		"admin/moder token"
// @Success 200 {object} models.Project Description
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProjectController) Put() {
	// TODO: обновить защиту когда будет лвлинг пользователей
	if c.Ctx.Output.IsOk() {
		idStr := c.Ctx.Input.Param(":id")
		id, _ := strconv.Atoi(idStr)
		v := models.Project{Id: id}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if err := models.UpdateProjectById(&v); err == nil {
				c.Data["json"] = "OK"
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Project
// @Param	id		path 	string	true		"The id you want to delete"
// @Param	token		query	string			false		"admin/moder token"
// @Success 200 OK
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProjectController) Delete() {
	// TODO: обновить защиту когда будет лвлинг пользователей
	if c.Ctx.Output.IsOk() {
		idStr := c.Ctx.Input.Param(":id")
		id, _ := strconv.Atoi(idStr)
		if err := models.DeleteProject(id); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}
