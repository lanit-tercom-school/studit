// @APIVersion 1.0.0
// @Title Auth-Service API
// @Description Auth-service provides authorization, registration, password change, etc.
package routers

import (
	"auth-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
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
