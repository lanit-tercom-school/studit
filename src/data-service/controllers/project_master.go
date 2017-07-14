package controllers

import (
	"data-service/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Записаться и отписаться от проекта, получить список записанных
type ProjectMasterController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *ProjectMasterController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description Добавить мастера к проекту
// @Param   user_id         query   string  true    "ID пользователя, корого нужно сделать мастером"
// @Param   project_id      query   string  true    "ID проекта, на который нужно добавить мастера"
// @Param   Bearer-token    header  string  true    "Токен доступа мастера проекта или админа"
// @Success 201 {int} "Created"
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectMasterController) Post() {
	if c.CurrentUser.PermissionLevel == VIEWER {
		beego.Debug("Access denied for `Post` new master for project")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR

		// получить id проекта, на который нужно добавить мастера
	} else if project_id, err := c.GetInt("project_id"); err != nil {
		beego.Debug("Not an int param. Should be int", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

		// Получаем список мастеров проекта
	} else if masters_of_this_project, err := models.GetMastersOfTheProject(project_id); err != nil {
		beego.Debug("Wrong project id", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if c.CurrentUser.PermissionLevel !=ADMIN && !models.IsUserInArray(c.CurrentUser.UserId, masters_of_this_project) {
		beego.Debug("Request from not master of this project")
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = "You are not a master of this project"

		// проект, на который записывается пользователь
	} else if project, err := models.GetProjectById(project_id); err != nil {
		beego.Debug("Wrong project id", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)

		// пользователь, которого назначают мастером
	} else if user_id, err := c.GetInt("user_id"); err != nil {
		beego.Debug("Not an int param. Should be int", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if user, err := models.GetUserById(user_id); err != nil {
		beego.Debug("Wrong user id", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else {
		// записать пользователя
		beego.Trace("Good user_id and project_id and Bearer-token")
		err := models.AddMasterToProject(user, project)
		if err != nil {
			beego.Critical("Can't add master", err.Error())
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = err.Error()

		} else {
			beego.Trace("Welcome, master")
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
func (c *ProjectMasterController) GetOne() {
	beego.Trace("New GET for project masters")
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Debug("Not an int param. Should be int", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else if v, err := models.GetMastersOfTheProject(id); err != nil {
		beego.Debug("GET masters of project error", err.Error())
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)

	} else {
		var t []models.MainUserInfo
		for _, r := range v {
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
// @Description Тестовый запрос для получения всех пар
// @Success 200 {object} []models.ProjectMaster
// @Failure 403
// @router / [get]
func (c *ProjectMasterController) GetAll() {
	l, err := models.GetAllProjectMaster()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Отобрать статус мастера. Может сделать только другой мастер или админ, самого себя удалить нельзя
// @Param   user_id         query   string  true    "ID пользователя, у которого нужно отобрать статус мастера"
// @Param   project_id      query   string  true    "ID проекта, на котором пользователь является мастером"
// @Param   Bearer-token    header  string  true    "Токен доступа мастера или админа"
// @Success 200 "OK"
// @Failure 403 id is empty
// @router / [delete]
func (c *ProjectMasterController) Delete() {
	beego.Trace("User want to rank down master")
	if c.CurrentUser.PermissionLevel == VIEWER {
		beego.Debug("Access denied for `Delete` master from project")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = "Forbidden"

	} else if user_id, err := c.GetInt("user_id"); err != nil {
		beego.Debug("Param `user_id` not int", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

		// Нельзя удалить самого себя
	} else if c.CurrentUser.UserId == user_id {
		beego.Debug(user_id, "want to delete himself")
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = "You can't delete yourself"

	} else if project_id, err := c.GetInt("project_id"); err != nil {
		beego.Debug("Param `user_id` not int", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

		// Получаем список мастеров
	} else if masters, err := models.GetMastersOfTheProject(project_id); err != nil {
		beego.Debug("Wrong `project_id`", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

		// Не мастера проекта и не админы не допускаются к дальнейшим действиям
	} else if c.CurrentUser.PermissionLevel != ADMIN && !models.IsUserInArray(c.CurrentUser.UserId, masters) {
		beego.Debug("Request from not master of this project")
		beego.Trace(project_id)
		beego.Trace(masters)
		beego.Trace(c.CurrentUser.UserId)
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = "You are not a master of this project"

		// Само удаление мастера с проекта
	} else if err := models.DeleteMasterFromProject(user_id, project_id); err != nil {
		beego.Debug("Can't delete from ProjectSingUp", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()

	} else {
		beego.Trace("Success `Delete` master from project")
		c.Data["json"] = HTTP_OK_STR
	}
	c.ServeJSON()
}
