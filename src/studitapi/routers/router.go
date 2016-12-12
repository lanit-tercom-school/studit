// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"studitapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),

		beego.NSNamespace("/task_for_test",
			beego.NSInclude(
				&controllers.TaskForTestController{},
			),
		),

		beego.NSNamespace("/variant",
			beego.NSInclude(
				&controllers.VariantController{},
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

		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),

		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),

		beego.NSNamespace("/tasks_tags_table",
			beego.NSInclude(
				&controllers.TasksTagsTableController{},
			),
		),

		beego.NSNamespace("/news",
			beego.NSInclude(
				&controllers.NewsController{},
			),
		),

		beego.NSNamespace("/news_tags",
			beego.NSInclude(
				&controllers.NewsTagsController{},
			),
		),

		beego.NSNamespace("/news_news_tags",
			beego.NSInclude(
				&controllers.NewsNewsTagsController{},
			),
		),

		beego.NSNamespace("/author",
			beego.NSInclude(
				&controllers.AuthorController{},
			),
		),

		beego.NSNamespace("/project_author",
			beego.NSInclude(
				&controllers.ProjectAuthorController{},
			),
		),

		beego.NSNamespace("/lesson",
			beego.NSInclude(
				&controllers.LessonController{},
			),
		),

		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),

		beego.NSNamespace("/project_user",
			beego.NSInclude(
				&controllers.ProjectUserController{},
			),
		),

		beego.NSNamespace("/contact_type",
			beego.NSInclude(
				&controllers.ContactTypeController{},
			),
		),

		beego.NSNamespace("/user_contact",
			beego.NSInclude(
				&controllers.UserContactController{},
			),
		),

		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),

		beego.NSNamespace("/user_course",
			beego.NSInclude(
				&controllers.UserCourseController{},
			),
		),

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/user_comments",
			beego.NSInclude(
				&controllers.UserCommentsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
