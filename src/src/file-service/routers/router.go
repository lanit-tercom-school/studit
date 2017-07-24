// @APIVersion 1.0.0
// @Title File-Service API
// @Description File-service provides downloading, deleting files, and getting information about them
package routers

import (
	"file-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/files",
			beego.NSInclude(
				&controllers.FileController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
