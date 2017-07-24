package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["main-service/controllers:CommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LandingProjectsController"] = append(beego.GlobalControllerRouter["main-service/controllers:LandingProjectsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecommendedCourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticsController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestsTaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
