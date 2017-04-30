package controllers

import (
	"encoding/json"
	"errors"
	"service/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

// Создание, изменение, удаление и просмотр проектов
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
// @Description Создать новый проект
// @Param   body            body        models.ProjectJson     true    "Тело запроса, см. пример"
// @Param   Bearer-token    header  string          true    "Токен доступа, пользователь должен быть не ниже куратора"
// @Success 201 {int} Created
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectController) Post() {
	beego.Trace("Try to POST project")
	if c.CurrentUser.PermissionLevel == 2 || c.CurrentUser.PermissionLevel == 1 {
		var v models.ProjectJson
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if id, err := models.AddProject(&v); err == nil {
				beego.Trace("Project with id", id, "created")
				v.Id = id
				user, err := models.GetUserById(c.CurrentUser.UserId)
				if err != nil {
					beego.Critical(c.Ctx.Input.IP(), "Claims corrupted", err.Error())
					c.Ctx.Output.SetStatus(500)
					c.Data["json"] = err.Error()

				} else {
					err := models.AddMasterToProject(user, &v)
					if err != nil {
						beego.Critical(c.Ctx.Input.IP(), "Can't add creator to project", err.Error())
						c.Ctx.Output.SetStatus(500)
						c.Data["json"] = err.Error()

					} else {
						beego.Trace("OK")
						c.Ctx.Output.SetStatus(201)
						c.Data["json"] = id
					}
				}
			} else {
				beego.Debug("Post project `AddProject` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(500)

			}
		} else {
			beego.Debug("Post project `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)

		}
	} else {
		beego.Debug("Access denied for `Post`")
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = "Forbbiden"

	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description Получить подробную информацию
// @Param   id  path    string  true    "ID проекта, информацию о котором нужно получить"
// @Success 200 {object} models.Project     Запрос прошел успешно
// @Failure 400 :id is wrong
// @router /:id [get]
func (c *ProjectController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	beego.Trace(c.Ctx.Input.IP(), "Get project with id", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Debug("GetOne `Atoi` error", err.Error())
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error() // TODO: change to "Wrong project id"
	} else {
		v, err := models.GetProjectById(int64(id))
		if err != nil {
			beego.Debug("GetOne `GetProjectById` error", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		} else {
			beego.Trace("GetOne OK")
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
// @Param       tag     query   string  false   "Получить проекты с тегом. Тег может быть только один."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer. Default 10"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} []models.Project Get array of projects filtered with specified filters (wtf this description)
// @Failure 403
// @router / [get]
func (c *ProjectController) GetAll() {
	var fields []string
	var sortBy []string
	var order []string
	var tag string
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
	// sortBy: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortBy = strings.Split(v, ",")
	}else{
		sortBy = []string{"Name"}
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}else{
		order  = []string{"asc"}
	}
	if v := c.GetString("tag"); v!= ""{
		tag = v
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

	beego.Trace(c.Ctx.Input.IP(), "Select from table")
	l, err := models.GetAllProject(query, fields, sortBy, order, offset, limit, tag)
	if err != nil {
		beego.Debug(c.Ctx.Input.IP(), "News GetAll `GetAllProject` error", err.Error())
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description Изменить/обновить проект
// @Param   id              path    string          true    "ID проекта, который нужно обновить"
// @Param   body            body    models.Project  true    "Тело запроса, см. пример"
// @Param   Bearer-token    header  string          true    "Токен доступа администратора или создателя проекта"
// @Success 200 "OK"
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProjectController) Put() {
	// TODO: добавить автора к проекту
	if c.CurrentUser.PermissionLevel == 2 || c.CurrentUser.PermissionLevel == 1 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Put `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
		v := models.Project{Id: int64(id)}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if err := models.UpdateProjectById(&v); err == nil {
				beego.Trace("Put project OK")
				c.Data["json"] = "OK"
			} else {
				beego.Debug("Put news `UpdateProjectById` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(400)
			}
		} else {
			beego.Debug("Put project `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		beego.Debug("Access denied for `Put`")
		c.Data["json"] = "You can't do it"
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

// TODO: удалится ли всё, что связано с проектом, если его удалить?
// Delete ...
// @Title Delete
// @Description delete the Project
// @Param   id              path    string      true        "ID проекта, который нужно удалить"
// @Param   Bearer-token    header  string      true        "Токен доступа администратора или автора проекта"
// @Success 200 "OK"
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProjectController) Delete() {
	// TODO: добавить проверку на автора
	if c.CurrentUser.PermissionLevel == 2 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Delete 'Atoi' error", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
		if err := models.DeleteProject(int64(id)); err == nil {
			beego.Trace("Delete OK")
			c.Data["json"] = "OK"
		} else {
			beego.Critical(c.Ctx.Input.IP(), "'DeleteProject' error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(500)
		}
	} else {
		beego.Debug("Access denied for `Delete`")
		c.Data["json"] = "Access denied for `Delete`" // TODO: change this
		c.Ctx.Output.SetStatus(403)
	}
	c.ServeJSON()
}
