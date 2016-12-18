package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"] = append(beego.GlobalControllerRouter["github.com/lanit-tercom-school/studit/service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
