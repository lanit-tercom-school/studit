// @APIVersion 1.0.1
// @Title Data-Service API
// @Description The data service provides data transfer from the database
package routers

import (
	"data-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/news",
			beego.NSInclude(
				&controllers.NewsController{},
			),
		),
		beego.NSNamespace("/project/enroll",
			beego.NSInclude(
				&controllers.UserEnrollOnProjectController{},
			),
		),
		beego.NSNamespace("/project/id",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),
		beego.NSNamespace("/project/user",
			beego.NSInclude(
				&controllers.ProjectUserController{},
			),
		),
		beego.NSNamespace("/user/contact",
			beego.NSInclude(
				&controllers.UserContactController{},
			),
		),
		beego.NSNamespace("/user/id",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/main/projects",
			beego.NSInclude(
				&controllers.LandingProjectsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
