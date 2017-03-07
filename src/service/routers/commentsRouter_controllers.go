package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["service/controllers:AuthController"] = append(beego.GlobalControllerRouter["service/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:AuthorController"] = append(beego.GlobalControllerRouter["service/controllers:AuthorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CommentController"] = append(beego.GlobalControllerRouter["service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CommentController"] = append(beego.GlobalControllerRouter["service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CommentController"] = append(beego.GlobalControllerRouter["service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CommentController"] = append(beego.GlobalControllerRouter["service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CommentController"] = append(beego.GlobalControllerRouter["service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CourseController"] = append(beego.GlobalControllerRouter["service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CourseController"] = append(beego.GlobalControllerRouter["service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CourseController"] = append(beego.GlobalControllerRouter["service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CourseController"] = append(beego.GlobalControllerRouter["service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:CourseController"] = append(beego.GlobalControllerRouter["service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:LandingProjectsController"] = append(beego.GlobalControllerRouter["service/controllers:LandingProjectsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:LessonController"] = append(beego.GlobalControllerRouter["service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:LessonController"] = append(beego.GlobalControllerRouter["service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:LessonController"] = append(beego.GlobalControllerRouter["service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:LessonController"] = append(beego.GlobalControllerRouter["service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:LessonController"] = append(beego.GlobalControllerRouter["service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:LogoutController"] = append(beego.GlobalControllerRouter["service/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsNewsTagsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:NewsTagsController"] = append(beego.GlobalControllerRouter["service/controllers:NewsTagsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectAuthorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TagController"] = append(beego.GlobalControllerRouter["service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TagController"] = append(beego.GlobalControllerRouter["service/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TagController"] = append(beego.GlobalControllerRouter["service/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TagController"] = append(beego.GlobalControllerRouter["service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TagController"] = append(beego.GlobalControllerRouter["service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskController"] = append(beego.GlobalControllerRouter["service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskController"] = append(beego.GlobalControllerRouter["service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskController"] = append(beego.GlobalControllerRouter["service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskController"] = append(beego.GlobalControllerRouter["service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskController"] = append(beego.GlobalControllerRouter["service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TestController"] = append(beego.GlobalControllerRouter["service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TestController"] = append(beego.GlobalControllerRouter["service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TestController"] = append(beego.GlobalControllerRouter["service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TestController"] = append(beego.GlobalControllerRouter["service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:TestController"] = append(beego.GlobalControllerRouter["service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserController"] = append(beego.GlobalControllerRouter["service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserController"] = append(beego.GlobalControllerRouter["service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserController"] = append(beego.GlobalControllerRouter["service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserController"] = append(beego.GlobalControllerRouter["service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserController"] = append(beego.GlobalControllerRouter["service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VariantController"] = append(beego.GlobalControllerRouter["service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VariantController"] = append(beego.GlobalControllerRouter["service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VariantController"] = append(beego.GlobalControllerRouter["service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VariantController"] = append(beego.GlobalControllerRouter["service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VariantController"] = append(beego.GlobalControllerRouter["service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VideoController"] = append(beego.GlobalControllerRouter["service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VideoController"] = append(beego.GlobalControllerRouter["service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VideoController"] = append(beego.GlobalControllerRouter["service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VideoController"] = append(beego.GlobalControllerRouter["service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["service/controllers:VideoController"] = append(beego.GlobalControllerRouter["service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
