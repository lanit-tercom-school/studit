package routers

import (
	"auth-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1/auth",
		beego.NSNamespace("/signin",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/signup",
			beego.NSInclude(
				&controllers.RegistrationController{},
			),
		),
		beego.NSNamespace("/reset",
			beego.NSInclude(
				&controllers.ResetPasswordController{},
			),
		),
		beego.NSNamespace("/change",
			beego.NSInclude(
				&controllers.ChangePasswordController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
