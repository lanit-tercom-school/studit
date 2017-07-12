// @APIVersion 1.0.1
// @Title Lanit-Tercom School API
// @Description API для Lanit-Tercom School
// @TermsOfServiceUrl No terms
// @License No License
// @LicenseUrl No License
// @URL https://vk.com/ltschool
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
		beego.NSNamespace("/files",
			beego.NSInclude(
				&controllers.FileController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
