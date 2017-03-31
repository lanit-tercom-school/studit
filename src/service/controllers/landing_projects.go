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
	//c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetAll)
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

/*// Post ...
// @Title Post
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router / [post]
func (c *LandingProjectsController) Post() {
	var response struct{
		Error string `json:"error"`
	}
	response.Error = "Method Not Allowed"
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router /:id [get]
func (c *LandingProjectsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProjectById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
*/

func (c *LandingProjectsController) GetOne() {
	c.Data["json"] = "Not Found"
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Projects for Landing page
// @Success 200 {object} models.Project Get best projects for main/landing page
// @Failure 400 Bad Request
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
/*
// Put ...
// @Title Put
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router /:id [put]
func (c *LandingProjectsController) Put() {
	var response struct{
		Error string `json:"error"`
	}
	response.Error = "Method Not Allowed"
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Method Not Allowed
// @Success 405 Method Not Allowed
// @Failure 405 Method Not Allowed
// @router /:id [delete]
func (c *LandingProjectsController) Delete() {
	var response struct{
		Error string `json:"error"`
	}
	response.Error = "Method Not Allowed"
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}*/
