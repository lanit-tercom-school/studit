package controllers

import (
	"data-service/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// ProjectController operations for ProjectId
type ProjectController struct {
	beego.Controller
}

type GetAllProjectsResponce struct {
	TotalCount  int
	ProjectList []models.Project
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
// @Description create ProjectId
// @Param	body		body 	models.ProjectId	true		"body for ProjectId content"
// @Success 201 {int} models.ProjectId
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectController) Post() {
	var v models.Project
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddProject(&v); err == nil {
			c.Ctx.Output.SetStatus(HTTP_CREATED)
			c.Data["json"] = v
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
// @Description get ProjectId by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProjectId
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProjectController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProjectById(id)
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
// @Description get ProjectId
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProjectId
// @Failure 403
// @router / [get]
func (c *ProjectController) GetAll() {
	var limit int64 = 10
	var offset int64 = 0
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	l, err := models.GetAllProject(offset, limit)
	if err != nil {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	} else {
		count, err := models.GetProjectCount()
		if err != nil {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		} else {
			responce := GetAllProjectsResponce{
				ProjectList: l,
				TotalCount:  int(count),
			}
			c.Ctx.Output.SetStatus(HTTP_OK)
			c.Data["json"] = responce
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ProjectId
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ProjectId	true		"body for ProjectId content"
// @Success 200 {object} models.ProjectId
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProjectController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Project{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProjectById(&v); err == nil {
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
// @Description delete the ProjectId
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProjectController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteProject(id); err == nil {
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
