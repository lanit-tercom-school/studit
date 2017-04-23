package controllers

import (
	"service/models"

	"github.com/astaxie/beego"
)

// Контроллер для главной/landing страниц
type LandingProjectsController struct {
	beego.Controller
}

// URLMapping ...
func (c *LandingProjectsController) URLMapping() {
	c.Mapping("Get", c.GetAll)
	c.Mapping("GetAll", c.GetAll)
}
func (c *LandingProjectsController) GetOne() {
	c.Data["json"] = "Not Found"
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description Получить проекты
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