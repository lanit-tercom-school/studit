package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["auth-service/controllers:AuthController"] = append(beego.GlobalControllerRouter["auth-service/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["auth-service/controllers:ChangePasswordController"] = append(beego.GlobalControllerRouter["auth-service/controllers:ChangePasswordController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["auth-service/controllers:RegistrationController"] = append(beego.GlobalControllerRouter["auth-service/controllers:RegistrationController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["auth-service/controllers:RegistrationController"] = append(beego.GlobalControllerRouter["auth-service/controllers:RegistrationController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["auth-service/controllers:ResetPasswordController"] = append(beego.GlobalControllerRouter["auth-service/controllers:ResetPasswordController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["auth-service/controllers:ResetPasswordController"] = append(beego.GlobalControllerRouter["auth-service/controllers:ResetPasswordController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
