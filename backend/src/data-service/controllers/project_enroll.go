package controllers

import (
	"data-service/models"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// ProjectEnrollController operations for ProjectEnroll
type ProjectEnrollController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProjectEnrollController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ProjectEnroll
// @Param	body		body 	models.ProjectEnroll	true		"body for ProjectEnroll content"
// @Success 201 {int} models.ProjectEnroll
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectEnrollController) Post() {
	var v models.ProjectEnroll
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if id, err := models.AddProjectEnroll(&v); err == nil {
			c.Ctx.Output.SetStatus(HTTP_CREATED)
			p, err := models.GetProjectEnrollById(int(id))
			if err == nil {
				c.Data["json"] = p
			} else {
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				c.Data["json"] = MakeMessageForSending(err.Error())
			}
		} else {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
	} else {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ProjectEnroll by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProjectEnroll
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProjectEnrollController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProjectEnrollById(id)
	if err != nil {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	} else {
		c.Ctx.Output.SetStatus(HTTP_OK)
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ProjectEnroll
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProjectEnroll
// @Failure 403
// @router / [get]
func (c *ProjectEnrollController) GetAll() {
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
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				c.Data["json"] = MakeMessageForSending("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllProjectEnroll(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	} else {
		c.Ctx.Output.SetStatus(HTTP_OK)
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ProjectEnroll
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ProjectEnroll	true		"body for ProjectEnroll content"
// @Success 200 {object} models.ProjectEnroll
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProjectEnrollController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ProjectEnroll{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProjectEnrollById(&v); err == nil {
			c.Ctx.Output.SetStatus(HTTP_OK)
			c.Data["json"] = MakeMessageForSending(HTTP_OK_STR)
		} else {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
	} else {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the ProjectEnroll by db_id, user_id or project_id
// @Param	db_id	query	string	false	"The db_id you want to delete"
// @Param	project_id	query	string	false	"The project_id you want to delete"
// @Param	user_id	query	string false	"The user_id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403
// @router / [delete]
func (c *ProjectEnrollController) Delete() {

	if db_id, err := c.GetInt("db_id"); err == nil {

		if err := models.DeleteProjectEnroll(db_id); err == nil {
			c.Ctx.Output.SetStatus(HTTP_OK)
			c.Data["json"] = MakeMessageForSending(HTTP_OK_STR)
		} else if err.Error() == "<QuerySeter> no row found" {
			c.Ctx.Output.SetStatus(HTTP_NOT_FOUND)
			c.Data["json"] = MakeMessageForSending(HTTP_NOT_FOUND_STR)
		} else {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
		c.ServeJSON()
	}

	if project_id, err := c.GetInt("project_id"); err == nil {

		if err := models.DeleteProjectEnrollProjectId(project_id); err == nil {
			c.Ctx.Output.SetStatus(HTTP_OK)
			c.Data["json"] = MakeMessageForSending(HTTP_OK_STR)
		} else if err.Error() == "<QuerySeter> no row found" {
			c.Ctx.Output.SetStatus(HTTP_NOT_FOUND)
			c.Data["json"] = MakeMessageForSending(HTTP_NOT_FOUND_STR)
		} else {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
		c.ServeJSON()

	}

	if user_id, err := c.GetInt("user_id"); err == nil {

		if err := models.DeleteProjectEnrollUserId(user_id); err == nil {
			c.Ctx.Output.SetStatus(HTTP_OK)
			c.Data["json"] = MakeMessageForSending(HTTP_OK_STR)
		} else if err.Error() == "<QuerySeter> no row found" {
			c.Ctx.Output.SetStatus(HTTP_NOT_FOUND)
			c.Data["json"] = MakeMessageForSending(HTTP_NOT_FOUND_STR)
		} else {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
		c.ServeJSON()

	}
}
