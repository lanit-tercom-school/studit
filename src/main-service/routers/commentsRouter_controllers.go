package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"] = append(beego.GlobalControllerRouter["main-service/controllers:ContactTypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:CourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:CourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LandingProjectsController"] = append(beego.GlobalControllerRouter["main-service/controllers:LandingProjectsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:LessonController"] = append(beego.GlobalControllerRouter["main-service/controllers:LessonController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:NewsController"] = append(beego.GlobalControllerRouter["main-service/controllers:NewsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:PractiseController"] = append(beego.GlobalControllerRouter["main-service/controllers:PractiseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectMasterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"] = append(beego.GlobalControllerRouter["main-service/controllers:ProjectUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"] = append(beego.GlobalControllerRouter["main-service/controllers:RecomendCoursesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:StatisticController"] = append(beego.GlobalControllerRouter["main-service/controllers:StatisticController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TagController"] = append(beego.GlobalControllerRouter["main-service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TagController"] = append(beego.GlobalControllerRouter["main-service/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TagController"] = append(beego.GlobalControllerRouter["main-service/controllers:TagController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TagController"] = append(beego.GlobalControllerRouter["main-service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TagController"] = append(beego.GlobalControllerRouter["main-service/controllers:TagController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TaskForTestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"] = append(beego.GlobalControllerRouter["main-service/controllers:TasksTagsTableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:TestController"] = append(beego.GlobalControllerRouter["main-service/controllers:TestController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCommentsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserContactController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserContactController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserCourseController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserCourseController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:id`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"] = append(beego.GlobalControllerRouter["main-service/controllers:UserEnrollOnProjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VariantController"] = append(beego.GlobalControllerRouter["main-service/controllers:VariantController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["main-service/controllers:VideoController"] = append(beego.GlobalControllerRouter["main-service/controllers:VideoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
