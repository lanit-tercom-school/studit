package controllers

import (
	"data-service/models"

	"github.com/astaxie/beego"
)

// Контроллер для главной/landing страниц
type LandingProjectsController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *LandingProjectsController) URLMapping() {
	c.Mapping("Get", c.GetAll)
}

// Get ...
// @Title Get
// @Description Получить проекты для главной страницы
// @Success 200 {object} models.Project Удачный запрос
// @Failure 400 Bad Request
// @router / [get]
func (c *LandingProjectsController) GetAll() {
	beego.Trace("Get landing projects")
	l, err := models.GetLandingProjects()

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}