package controllers

import "github.com/astaxie/beego"

type ErrorController struct{
	beego.Controller
}

func (c *ErrorController) Error404()  {
	response := ErrorResponse{"Not Found"}
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

func (c *ErrorController) Error405()  {
	response := ErrorResponse{"Method Not Allowed"}
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}