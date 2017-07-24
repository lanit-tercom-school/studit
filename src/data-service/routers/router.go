// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
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

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/project_enroll",
			beego.NSInclude(
				&controllers.ProjectEnrollController{},
			),
		),

		beego.NSNamespace("/contact_type",
			beego.NSInclude(
				&controllers.ContactTypeController{},
			),
		),

		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),

		beego.NSNamespace("/user_contact",
			beego.NSInclude(
				&controllers.UserContactController{},
			),
		),

		beego.NSNamespace("/practise",
			beego.NSInclude(
				&controllers.PractiseController{},
			),
		),

		beego.NSNamespace("/video",
			beego.NSInclude(
				&controllers.VideoController{},
			),
		),

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),

		beego.NSNamespace("/task_for_test",
			beego.NSInclude(
				&controllers.TestsTaskController{},
			),
		),

		beego.NSNamespace("/Statistics",
			beego.NSInclude(
				&controllers.StatisticsController{},
			),
		),

		beego.NSNamespace("/user_course",
			beego.NSInclude(
				&controllers.UserCourseController{},
			),
		),

		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),

		beego.NSNamespace("/recomend_courses",
			beego.NSInclude(
				&controllers.RecommendedCourseController{},
			),
		),

		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),

		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),

		beego.NSNamespace("/variant",
			beego.NSInclude(
				&controllers.VariantController{},
			),
		),

		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),

		beego.NSNamespace("/project_user",
			beego.NSInclude(
				&controllers.ProjectUserController{},
			),
		),

		beego.NSNamespace("/user_comments",
			beego.NSInclude(
				&controllers.UserCommentController{},
			),
		),

		beego.NSNamespace("/lesson",
			beego.NSInclude(
				&controllers.LessonController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
