package controllers

import (
	"encoding/json"
	"errors"
	"service/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// LandingProjectsController oprations for Project
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
// @Failure 403
// @router / [get]
func (c *LandingProjectsController) GetAll() {
	l, err := models.GetLandingProjects()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router /:id [put]
func (c *LandingProjectsController) Put() {
	c.CustomAbort(405, "Method Not Allowed")
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
