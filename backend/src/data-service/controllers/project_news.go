package controllers

import (
	"data-service/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// ProjectEnrollController operations for ProjectEnroll
type ProjectNewsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProjectNewsController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create ProjectNews
// @Param	body		body 	models.ProjectNews	true		"body for ProjectNews content"
// @Success 201 {int} models.ProjectNews
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectNewsController) Post() {
	var v models.ProjectNews
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if id, err := models.AddProjectNews(&v); err == nil {
			c.Ctx.Output.SetStatus(HTTP_CREATED)
			_, err := models.GetProjectNewsById(int(id))
			if err == nil {
				c.Data["json"] = MessageType{"Project news was succesfully added"}
			} else {
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				c.Data["json"] = MakeMessageForSending(err.Error())
			}
		} else {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
	} else {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	}
	c.ServeJSON()
}
