package controllers

import (
	"encoding/json"
	"service/models"
	"strconv"
	"github.com/astaxie/beego"
)

// Записаться и отписаться от проекта, получить список записанных
type UserEnrollOnProjectController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *UserEnrollOnProjectController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}
// Param   project_id      query   int     true    "ID проекта, на который нужно записаться"

// Post ...
// @Title Post
// @Description Записать пользователя на проект
// @Param   id              path    string  true    "ID проекта, на который нужно записаться"
// @Param   Bearer-token    header  string  true    "Токен доступа любого зарегистрированного пользователя"
// @Success 201 {int} "Created"
// @Failure 403 body is empty
// @router /:id [post]
func (c *UserEnrollOnProjectController) Post() {
	if c.CurrentUser.PermissionLevel == -1 {
		beego.Debug(c.Ctx.Input.IP(), "Access denied for `Post` new application form")
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = "Forbidden"

	} else {
		// получить id проекта, на который пользователь хочет записаться
		project_id, err := c.GetInt64(":id")
		if err != nil {
			beego.Debug(c.Ctx.Input.IP(), "Not an int param. Should be int", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()

		} else {
			// проект, на который записывается пользователь
			project, err := models.GetProjectById(project_id)
			if err != nil {
				beego.Debug("Wrong project id", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(400)

			} else {
				// пользователь, который записывается
				user, err := models.GetUserById(c.CurrentUser.UserId)
				if err != nil {
					beego.Critical("Corrupted claims", err.Error())
					c.Ctx.Output.SetStatus(500)
					c.Data["json"] = err.Error()

				} else {
					// записать пользователя
					beego.Trace("Good user_id")
					_, err := models.AddApplicationFromUserForProject(user, project)
					if err != nil {
						beego.Critical("Corrupted claims", err.Error())
						c.Ctx.Output.SetStatus(500)
						c.Data["json"] = err.Error()

					} else {
						beego.Trace("New successfull sign up on project")
						c.Ctx.Output.SetStatus(201)
						c.Data["json"] = "Created"
					}
				}
			}
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description Получить список записанных пользователей
// @Param   id  path    string  true    "ID проекта, на которых нужно узнать список записанных"
// @Success 200 []{int} список из ID пользователей
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserEnrollOnProjectController) GetOne() {
	beego.Trace("New GET for singed up users")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAllSignedUpOnProject(id)
	if err != nil {
		beego.Debug("GET list of signed up users error", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(500)
	} else {
		beego.Trace("Success GET")
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description Тестовый запрос для получения всех пар
// @Success 200 {object} models.ProjectEnroll
// @Failure 403
// @router / [get]
func (c *UserEnrollOnProjectController) GetAll() {
	l, err := models.GetAllProjectAuthor(make(map[string]string), []string{}, []string{}, []string{}, 0, 100)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
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

// wtf
func (c *UserEnrollOnProjectController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ProjectEnroll{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProjectAuthorById(&v); err == nil {
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
// @Description Отписаться от проекта
// @Param   id              path    string  true    "ID проекта, из которого нужно удалить заявку"
// @Param   Bearer-token    header  string  true    "Токен доступа любого зарегистрированного пользователя"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserEnrollOnProjectController) Delete() {
	beego.Trace("User want to sign out from project")
	if c.CurrentUser.PermissionLevel == -1 {
		beego.Debug(c.Ctx.Input.IP(), "Access denied for `Delete` project_sign_up")
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = "Forbidden"
	} else {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Can't parse", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		} else if err := models.DeleteProjectSignUp(c.CurrentUser.UserId, id); err == nil {
			beego.Trace("Success sign out from project")
			c.Data["json"] = "OK"
		} else {
			beego.Debug("Can't delete from ProjectSingUp", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}
