package controllers

import (
	"main-service/models"
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
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description Записать пользователя на проект
// @Param   id              path    string  true    "ID проекта, на который нужно записаться"
// @Param   message         query   string  false   "Сопроводительный текст для мастеров"
// @Param   Bearer-token    header  string  true    "Токен доступа любого зарегистрированного пользователя"
// @Success 201 {int} "Created"
// @Failure 403 body is empty
// @router /:id [post]
func (c *UserEnrollOnProjectController) Post() {
	if c.CurrentUser.PermissionLevel == models.VIEWER {
		beego.Debug(c.Ctx.Input.IP(), "Access denied for `Post` new application form")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR

		// получить id проекта, на который пользователь хочет записаться
	} else if project_id, err := c.GetInt(":id"); err != nil {
		beego.Debug(c.Ctx.Input.IP(), "Not an int param. Should be int", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

		// проект, на который записывается пользователь
	} else if project, err := models.GetProjectById(project_id); err != nil {
		beego.Debug("Wrong project id", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)

		// пользователь, который записывается
	} else if user, err := models.GetUserById(c.CurrentUser.UserId); err != nil {
		beego.Critical("Corrupted claims", err.Error())
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = err.Error()

	} else {
		// записать пользователя
		beego.Trace("Good user_id")
		_, err := models.AddApplicationFromUserForProject(user, project, c.GetString("message"))
		if err != nil {
			beego.Critical("Corrupted claims", err.Error())
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = err.Error()

		} else {
			beego.Trace("New successful sign up on project")
			c.Ctx.Output.SetStatus(HTTP_CREATED)
			c.Data["json"] = HTTP_CREATED_STR
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
	beego.Trace("New GET for enrolled users")
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Debug("Can't parse", idStr, err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	}
	v, err := models.GetAllSignedUpOnProject(id)
	if err != nil {
		beego.Debug("GET list of signed up users error", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
	} else {
		var t []models.MainUserInfo
		for _, r := range v {
			t = append(t, models.MainUserInfo{
				Avatar:   r.Avatar,
				Nickname: r.Nickname,
				Id:       r.Id,
			})
		}
		beego.Trace("Success GET")
		c.Data["json"] = t
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description Получение информации по заявкам на проекты учителя
// @Param   Bearer-token    header  string  true    "Токен доступа мастера проекта"
// @Success 200 {object} []models.ObjectOfListOfEnrolledUsersOnProject "desc"
// @Failure 403
// @router / [get]
func (c *UserEnrollOnProjectController) GetAll() {
	beego.Trace("New GET for enrolled users")
	if c.CurrentUser.PermissionLevel >= models.LEADER {
		l, err := models.GetAllEnrolledOnProject(c.CurrentUser.UserId)
		if err != nil {
			beego.Debug("Something wrong", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		} else {
			beego.Trace("Good request")
			c.Data["json"] = l
		}
	} else {
		beego.Debug("Forbidden to get enrolled users")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR
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
	if c.CurrentUser.PermissionLevel == models.VIEWER {
		beego.Debug(c.Ctx.Input.IP(), "Access denied for `Delete` project_sign_up")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR

	} else {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Can't parse", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()

		} else if err := models.DeleteProjectSignUp(c.CurrentUser.UserId, id); err == nil {
			beego.Trace("Success sign out from project")
			c.Data["json"] = HTTP_OK_STR

		} else {
			beego.Debug("Can't delete from ProjectSingUp", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}
