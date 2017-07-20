package controllers

import (
	"data-service-old/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// ProjectUserController operations for ProjectUser
type ProjectUserController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *ProjectUserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description Добавление пользователя к проекту
// @Param       project_id     query     string   true           "ID проекта, на который надо записать пользователя"
// @Param       user_id        query     string   true           "ID пользователя, которого надо записать"
// @Param       Bearer-token   header   string   true           "Токен доступа администратора или куратора проекта"
// @Success 201 {int} models.ProjectUser "Добавленное поле ProjectUser"
// @Failure 403 access denied
// @router / [post]
func (c *ProjectUserController) Post() {
	// TODO: сделать проверку того, что куратор добавляет пользователя именно к своему проекту
	if c.CurrentUser.PermissionLevel < LEADER {
		beego.Debug(c.Ctx.Input.IP(), "Access denied for `Post` new user to project")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR

	} else if project_id, err := c.GetInt("project_id"); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Not an int param. Should be int")
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if user_id, err := c.GetInt("user_id"); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Not an int param. Should be int")
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if project, err := models.GetProjectById(project_id); err != nil {
		beego.Debug("Wrong project id", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)

	} else if user, err := models.GetUserById(user_id); err != nil {
		beego.Critical("Wrong user id", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	} else if err := models.IsTeacherInThisProject(c.CurrentUser.UserId, project_id); err != nil {
		beego.Critical("Teacher is not in this project", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	} else if err := models.AddUserToProject(user, project); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "`AddUserToProject` method error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = err.Error()

	} else {
		beego.Trace("New successfull sign up on project")
		c.Ctx.Output.SetStatus(HTTP_CREATED)
		c.Data["json"] = HTTP_CREATED_STR
	}
	c.ServeJSON()
}

// Get...
// @Title Get
// @Description Получение списка участников проекта
// @Param	id		path 	string	true		"ID проекта, список участников которого нужно узнать"
// @Success 200 {object} []int "Список пользователей"
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProjectUserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	beego.Trace(c.Ctx.Input.IP(), "Get project with id", idStr)
	if id, err := strconv.Atoi(idStr); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Get `Atoi` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if users, err := models.GetUsersByProjectId(id); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "`GetUsersByProjectId` method error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = err.Error()

	} else {
		var t []models.MainUserInfo
		for _, r := range users {
			t = append(t, models.MainUserInfo{
				Id:       r.Id,
				Nickname: r.Nickname,
				Avatar:   r.Avatar,
			})
		}
		beego.Trace("Success GET")
		c.Data["json"] = t
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description Тестовый запрос
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProjectUser
// @Failure 403
// @router / [get]
func (c *ProjectUserController) GetAll() {
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

	l, err := models.GetAllProjectUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ProjectUser
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ProjectUser	true		"body for ProjectUser content"
// @Success 200 {object} models.ProjectUser
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProjectUserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ProjectUser{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProjectUserById(&v); err == nil {
			c.Data["json"] = HTTP_OK_STR
		} else {
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		}
	} else {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Удаление пользователя из проекта
// @Param       user_id        query    string   true           "ID пользователя, которого надо удалить"
// @Param       project_id     query    string   true           "ID проекта, с которого надо удалить пользователя"
// @Param       Bearer-token   header   string   true           "Токен доступа пользователя (куратора/админа)"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router / [delete]
func (c *ProjectUserController) Delete() {
	// TODO: сделать проверку, что куратор удаляет пользователей именно из своего проекта
	user_id, err := c.GetInt("user_id")
	if err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Not an int param. Should be int")
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if project_id, err := c.GetInt("project_id"); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Not an int param. Should be int")
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if c.CurrentUser.PermissionLevel < LEADER && c.CurrentUser.UserId != user_id {
		beego.Debug(c.Ctx.Input.IP(), "Access denied for `Delete` user from project")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR

	} else if err := models.DeleteUserFromProject(user_id, project_id); err == nil {
		beego.Trace(c.Ctx.Input.IP(), "`Delete` user from project OK")
		c.Data["json"] = HTTP_OK_STR

	} else {
		beego.Debug(c.Ctx.Input.IP(), "Can't `Delete` user from project")
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
