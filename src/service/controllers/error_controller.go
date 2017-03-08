package controllers

import "github.com/astaxie/beego"

type ErrorController struct{
	beego.Controller
}

func (c *ErrorController) Error404()  {
	c.Data["json"] = "Not Found"
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.ServeJSON()
}

func (c *ErrorController) Error405()  {
	c.Data["json"] = "Method Not Allowed"
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}