package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type ErrorController struct{
	beego.Controller
}

func (c *ErrorController) Error404()  {
	if beego.BConfig.RunMode == "dev" {
		c.Redirect(c.Ctx.Input.Site()+":"+strconv.Itoa(c.Ctx.Input.Port()), 301)
	} else if beego.BConfig.RunMode == "prod" {
		c.Data["json"] = "Not Found"
		c.Ctx.ResponseWriter.WriteHeader(404)
		c.ServeJSON()
	}
}

func (c *ErrorController) Error405()  {
	c.Data["json"] = "Method Not Allowed"
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}