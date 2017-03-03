package controllers

import (
	"service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// LandingProjectsController operations for Project
type LandingProjectsController struct {
	beego.Controller
}

// URLMapping ...
func (c *LandingProjectsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router / [post]
func (c *LandingProjectsController) Post() {
	logs.Info("wtf post")
	c.CustomAbort(405, "Method Not Allowed")
}

// GetOne ...
// @Title Get One
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router /:id [get]
func (c *LandingProjectsController) GetOne() {
	c.CustomAbort(405, "Method Not Allowed")
}

// GetAll ...
// @Title Get All
// @Description get Projects for Landing page
// @Success 200 {object} models.Project
// @Failure 400 Bad Request
// @router / [get]
func (c *LandingProjectsController) GetAll() {
	l, err := models.GetLandingProjects()

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	logs.Info("wtf getall")
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router /:id [put]
func (c *LandingProjectsController) Put() {
	c.Data["json"] = "OK"
	c.CustomAbort(405, "Method Not Allowed")
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router /:id [delete]
func (c *LandingProjectsController) Delete() {
	c.CustomAbort(405, "Method Not Allowed")
}
