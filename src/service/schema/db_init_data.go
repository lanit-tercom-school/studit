package main

import (
	_ "service/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	m "service/models"
	"time"
	"github.com/astaxie/beego"
	"fmt"
	"service/auth"
)

func init() {
	err := orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func fastCheckErr(_ int64, err error) {
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func main() {
	o := orm.NewOrm()
	o.Using("default")

	// add projects

	project1 := m.Project{
		Id: 1,
		Name: "Образовательный портал Studit",
		Description: "Разработка образовательного портала для Lanit-Tercom School",
		DateOfCreation: time.Now(),
		Logo: "/files/1.jpg",
		Tags: "studit,summerschool",
		Status: 1,
	}
	fastCheckErr(o.Insert(&project1))

	project2 := m.Project{
		Id: 2,
		Name: "Модный фрилансер",
		Description: "Какие же стрелочки вокруг ноубука!",
		DateOfCreation: time.Now(),
		Logo: "/files/2.jpg",
		Tags: "freelance",
		Status: 0,
	}
	fastCheckErr(o.Insert(&project2))

	project3 := m.Project{
		Id: 3,
		Name: "Оригинальное название",
		Description: "Click-bait описание",
		DateOfCreation: time.Now(),
		Logo: "/files/3.jpg",
		Tags: "creative",
		Status: 0,
	}
	fastCheckErr(o.Insert(&project3))

	project4 := m.Project{
		Id: 4,
		Name: "TFS Mobile",
		Description: "Студенческий проект TFS Mobile",
		DateOfCreation: time.Now(),
		Logo: "/files/4.jpg",
		Tags: "TFSMobile,summerschool",
		Status: 0,
	}
	fastCheckErr(o.Insert(&project4))

	project5 := m.Project{
		Id: 5,
		Name: "Еще один проект",
		Description: "Описаниеописаниеописание",
		DateOfCreation: time.Now(),
		Logo: "/files/5.jpg",
		Tags: "project",
		Status: 0,
	}
	fastCheckErr(o.Insert(&project5))

	// add users
	avatar_seed := auth.GenerateNewToken(6)
	color_str := auth.GenerateRandomColor()
	user1 := m.User{
		Id: 1,
		Nickname: "Admin",
		Login: "a@a",
		Password: auth.CustomStr("a").ToSHA1(),
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "Главный по тарелкам",
		PermissionLevel: 2,
	}
	fastCheckErr(o.Insert(&user1))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user2 := m.User{
		Id: 2,
		Nickname: "Moderator",
		Login: "moder@moder.moder",
		Password: auth.CustomStr("moder").ToSHA1(),
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "Главный по молоткам",
        	PermissionLevel: 1,
	}
	fastCheckErr(o.Insert(&user2))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user3 := m.User{
		Id: 3,
		Nickname: "Егорка2003",
		Login: "egorka2003@maaail.ru",
		Password: auth.CustomStr("пароль").ToSHA1(),
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "ЮЮЮ, ААА",
        	PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user3))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user4 := m.User{
		Id: 4,
		Nickname: "ZagadOchNayA",
		Login: "zagadka@maaail.ru",
		Password: auth.CustomStr("котикипёсики").ToSHA1(),
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "Легко потерять, невозможно забить",
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user4))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user5 := m.User{
		Id: 5,
		Nickname: "S1ayeR1",
		Login: "slayer342@bbk.ru",
		Password: auth.CustomStr("lala").ToSHA1(),
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "bjklknufu",
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user5))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user6 := m.User{
		Id: 6,
		Nickname: "NekoTyan",
		Login: "petrovich82@maaail.ru",
		Password: auth.CustomStr("pasuwaado").ToSHA1(),
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "^_^",
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user6))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user7 := m.User{
		Id: 7,
		Nickname: "B",
		Login: "b@b",
		Password: auth.CustomStr("b").ToSHA1(),
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "BBB",
		PermissionLevel: 0,
	}
	fastCheckErr(o.Insert(&user7))

	// add tags

	tag1 := m.Tag{
		Id: 100,
		Name: "Other",
	}
	fastCheckErr(o.Insert(&tag1))

	tag2 := m.Tag{
		Id: 101,
		Name: "New",
	}
	fastCheckErr(o.Insert(&tag2))

	tag3 := m.Tag{
		Id: 102,
		Name: "C/C++",
	}
	fastCheckErr(o.Insert(&tag3))

	tag4 := m.Tag{
		Id: 103,
		Name: "Golang",
	}
	fastCheckErr(o.Insert(&tag4))

	tag5 := m.Tag{
		Id: 104,
		Name: "JavaScript",
	}
	fastCheckErr(o.Insert(&tag5))

	tag6 := m.Tag{
		Id: 105,
		Name: "C#",
	}
	fastCheckErr(o.Insert(&tag6))

	tag7 := m.Tag{
		Id: 106,
		Name: "Front-end",
	}
	fastCheckErr(o.Insert(&tag7))

	tag8 := m.Tag{
		Id: 107,
		Name: "Back-end",
	}
	fastCheckErr(o.Insert(&tag8))

	//connect pojects and masters
	prj_mas_con1 := m.ProjectMaster{
		Id: 1,
		ProjectId: &project2,
		MasterId : &user1,
		SignedDate: time.Now(),
	}
	fastCheckErr(o.Insert(&prj_mas_con1))

	prj_mas_con2 := m.ProjectMaster{
		Id: 2,
		ProjectId: &project5,
		MasterId : &user2,
		SignedDate: time.Now(),
	}
	fastCheckErr(o.Insert(&prj_mas_con2))

	prj_mas_con3 := m.ProjectMaster{
		Id: 3,
		ProjectId: &project4,
		MasterId : &user3,
		SignedDate: time.Now(),
	}
	fastCheckErr(o.Insert(&prj_mas_con3))

	prj_mas_con4 := m.ProjectMaster{
		Id: 4,
		ProjectId: &project1,
		MasterId : &user4,
		SignedDate: time.Now(),
	}
	fastCheckErr(o.Insert(&prj_mas_con4))

	prj_mas_con5 := m.ProjectMaster{
		Id: 5,
		ProjectId: &project3,
		MasterId : &user5,
		SignedDate: time.Now(),
	}
	fastCheckErr(o.Insert(&prj_mas_con5))





	//fastCheckErr(o.Insert(&prj_mas_con2))

	// connect projects and users

	prj_usr_con1 := m.ProjectUser{
		Id: 1,
		UserId: &user1,
		ProjectId: &project1,
		SignedDate: time.Now(),
		Progress: 0,
	}
	fastCheckErr(o.Insert(&prj_usr_con1))

	prj_usr_con2 := m.ProjectUser{
		Id: 2,
		UserId: &user1,
		ProjectId: &project2,
		SignedDate: time.Now(),
		Progress: 0,
	}
	fastCheckErr(o.Insert(&prj_usr_con2))

	prj_usr_con3 := m.ProjectUser{
		Id: 3,
		UserId: &user1,
		ProjectId: &project3,
		SignedDate: time.Now(),
		Progress: 0,
	}
	fastCheckErr(o.Insert(&prj_usr_con3))

	prj_usr_con4 := m.ProjectUser{
		Id: 4,
		UserId: &user2,
		ProjectId: &project3,
		SignedDate: time.Now(),
		Progress: 0,
	}
	fastCheckErr(o.Insert(&prj_usr_con4))

	prj_usr_con5 := m.ProjectUser{
		Id: 5,
		UserId: &user2,
		ProjectId: &project2,
		SignedDate: time.Now(),
		Progress: 0,
	}
	fastCheckErr(o.Insert(&prj_usr_con5))

	prj_usr_con6 := m.ProjectUser{
		Id: 6,
		UserId: &user3,
		ProjectId: &project3,
		SignedDate: time.Now(),
		Progress: 0,
	}
	fastCheckErr(o.Insert(&prj_usr_con6))

	// add contact types
/*
	contactType1 := m.ContactType{
		Id: 1,
		Type: "email",
	}
	fastCheckErr(o.Insert(&contactType1))

	contactType2 := m.ContactType{
		Id: 2,
		Type: "mobile phone",
	}
	fastCheckErr(o.Insert(&contactType2))

	contactType3 := m.ContactType{
		Id: 3,
		Type: "phone",
	}
	fastCheckErr(o.Insert(&contactType3))

	contactType4 := m.ContactType{
		Id: 4,
		Type: "skype",
	}
	fastCheckErr(o.Insert(&contactType4))

	contactType5 := m.ContactType{
		Id: 5,
		Type: "vk.com",
	}
	fastCheckErr(o.Insert(&contactType5))

	contactType6 := m.ContactType{
		Id: 6,
		Type: "telegram",
	}
	fastCheckErr(o.Insert(&contactType6))

	contactType7 := m.ContactType{
		Id: 7,
		Type: "viber",
	}
	fastCheckErr(o.Insert(&contactType7))
*/
	// add user contacts

	user_contact1 := m.UserContact{
		Id: 1,
		UserId: &user1,
		Contact: "a@a",
		Type: "skype",
	}
	fastCheckErr(o.Insert(&user_contact1))

	user_contact2 := m.UserContact{
		Id: 2,
		UserId: &user2,
		Contact: "moder@moder.moder",
		Type: "email",
	}
	fastCheckErr(o.Insert(&user_contact2))

	user_contact3 := m.UserContact{
		Id: 3,
		UserId: &user3,
		Contact: "egorka2003@maaail.ru",
		Type: "email",
	}
	fastCheckErr(o.Insert(&user_contact3))

	user_contact4 := m.UserContact{
		Id: 4,
		UserId: &user1,
		Contact: "+7-123-456-78-90",
		Type: "mobile phone",
	}
	fastCheckErr(o.Insert(&user_contact4))

	user_contact5 := m.UserContact{
		Id: 5,
		UserId: &user5,
		Contact: "slayer17",
		Type: "vk.com",
	}
	fastCheckErr(o.Insert(&user_contact5))

	user_contact6 := m.UserContact{
		Id: 6,
		UserId: &user4,
		Contact: "zagad0chnaya",
		Type: "telegram",
	}
	fastCheckErr(o.Insert(&user_contact6))

	user_contact7 := m.UserContact{
		Id: 7,
		UserId: &user6,
		Contact: "nekotyanmimimi",
		Type: "viber",
	}
	fastCheckErr(o.Insert(&user_contact7))

	//add comments

	comment1 := m.Comment{
		Id:	1,
		Text: "CommentCommentComment1",
	}
	fastCheckErr(o.Insert(&comment1))

	comment2 := m.Comment{
		Id:	2,
		Text: "CommentCommentComment2",
	}
	fastCheckErr(o.Insert(&comment2))

	comment3 := m.Comment{
		Id:	3,
		Text: "CommentCommentComment3",
	}
	fastCheckErr(o.Insert(&comment3))

	comment4 := m.Comment{
		Id:	4,
		Text: "CommentCommentComment4",
	}
	fastCheckErr(o.Insert(&comment4))

	// add user comments

	usr_comment1 := m.UserComments{
		Id:		1,
		UserId:		&user1,
		CommentId:	&comment1,
		Date: 		time.Now(),
	}
	fastCheckErr(o.Insert(&usr_comment1))

	usr_comment2 := m.UserComments{
		Id:		2,
		UserId:		&user1,
		CommentId:	&comment2,
		Date: 		time.Now(),
	}
	fastCheckErr(o.Insert(&usr_comment2))

	usr_comment3 := m.UserComments{
		Id:		3,
		UserId:		&user2,
		CommentId:	&comment3,
		Date: 		time.Now(),
	}
	fastCheckErr(o.Insert(&usr_comment3))

	usr_comment4 := m.UserComments{
		Id:		4,
		UserId:		&user3,
		CommentId:	&comment4,
		Date: 		time.Now(),
	}
	fastCheckErr(o.Insert(&usr_comment4))

	// add news

	news1 := m.NewsJson{
		Id: 0,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Программисты признаются в своих грехах в знак протеста против собеседований" +
			" с кодингом «на бумаге» — в Твиттере появился новый флешмоб",
		Description: "Где проводят такие собеседования? Такой стиль собеседований широко используется в" +
			" IT-индустрии, в том числе в таких компаниях, как Google и Amazon. Кандидатам не дают никакого доступа" +
			" к справочным материалам и просят решить какую-либо техническую задачу, что, по мнению уже благополучно" +
			" работающих где-либо программистов, деморализует и не выявляет реальных навыков.",
		Tags: []string{"Other"},
		Image: "https://image.freepik.com/free-vector/programmer-working-on-the-computer_23-2147505689.jpg",
	}
	fastCheckErr(m.AddNews(&news1))

	news2 := m.NewsJson{
		Id: 0,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Недавний сбой в работе облачных сервисов Amazon был вызван опечаткой сотрудника компании",
		Description: "Крупный сбой в работе облачных сервисов Amazon, из-за которого 28 февраля 2017 года появились" +
			" проблемы в работе Slack, Trello, Coursera и многих других сайтов, произошёл из-за опечатки одного" +
			" из сотрудников компании. Об этом вчера сообщила Amazon.28 февраля 2017 года команда Amazon S3" +
			" занималась отладкой своей биллинговой системы." +
			" В ходе работ ей понадобилось перевести несколько серверов в автономный режим, но при вводе команды" +
			" была допущена ошибка: К сожалению, команда была введена неправильно и под отключение попало больше" +
			" серверов, чем предполагалось. Среди них были серверы, которые поддерживали работу подсистем S3." +
			" В частности, была отключена подсистема, управляющая метаданными и информацией о расположении всех" +
			" серверов S3 в регионе. Если вы это читаете, значит, мы дебажем на продакшене. От неё зависела" +
			" работа многих сервисов для поиска и хранения данных, в том" +
			" числе и Amazon Elastic Compute Cloud (EC2), который используется для доступа к вычислительным" +
			" мощностям. Представители компании пояснили, что для восстановления работоспособности сервисов нужно" +
			" было перезапустить некоторые системы и провести их проверку, что заняло довольно много времени." +
			" По их словам, S3 способен работать при отключении нескольких серверов, но массовая перезагрузка" +
			" стала проблемой. Сейчас Amazon уже внесла в S3 изменения, которые позволят ускорить процесс" +
			" восстановления систем. Кроме того, инженеры не смогут отключать серверы, которые задействованы" +
			" в системах «определённого уровня». Также компания пообещала исправить работу информационной" +
			" панели AWS Service Health Dashboard, которая не показывала информацию о сбоях, так как сама" +
			" зависела от упавшего сервиса S3.",
		Tags: []string{"Other", "World"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news2))

	news3 := m.NewsJson{
		Id: 3,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "LinkedIn официально отказалась выполнять требования российского законодательства",
		Description: "Профессиональная социальная сеть LinkedIn официально отказалась выполнять требования " +
			" Роскомнадзора и российского законодательства, сообщается на официальном сайте регулятора. В публикации " +
			" говорится:\"Роскомнадзор получил письмо от вице-президента по международной публичной политике LinkedIn" +
			" Corporation Пабло Л. Чавеса. В письме сообщается, что компания не готова устранить нарушения российского" +
			" законодательства. Компания отказалась исполнить требование о локализации баз с персональными данными" +
			" российских граждан на территории Российской Федерации, подтвердив таким образом свою незаинтересованность" +
			" в работе на российском рынке.\" Технически, отказ от выполнения требований переводит блокировку" +
			" LinkedIn в разряд «вечных» и является демаршем против политики российских властей и требований" +
			" регулирующих органов. Социальная сеть LinkedIn была заблокирована Роскомнадзором еще в ноябре" +
			" 2016 года. Также в российских версиях магазинов стало невозможно скачать приложение социальной" +
			" сети. Фактически, LinkedIn окончательно «сжег мосты»: компания предпочла, чтобы ее ресурс остался" +
			" заблокированным, чем выполнять требования Роскомнадзора и ФЗ-152. ",
		Tags: []string{"Other", "World"},
		Image: "https://image.freepik.com/free-vector/programmer-working-on-the-computer_23-2147505689.jpg",
	}
	fastCheckErr(m.AddNews(&news3))

	news4 := m.NewsJson{
		Id: 4,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Искусственный интеллект DeepMind AlphaGo тайно продолжает сокрушать людей-игроков в онлайн-режиме",
		Description: "Оказывается, что система искусственного интеллекта DeepMind AlphaGo, предназначенная для игры" +
			" в древнюю китайскую игру Го, не “сошла со сцены” после того, как в прошлом году она победила лучшего" +
			" в мире человека-игрока, шестикратного Чемпиона Мира Ли Седоля. Несколько месяцев назад на" +
			" онлайн-платформе Tygem игры Го появился игрок с ником “Master”, который начал крушить всех" +
			" подряд. “Master” дважды победил даже Ки Джи (Ke Jie), первого номера на Tygem, и выиграл в" +
			" 50 играх из 51 игры, в которых он принимал участие. Да и эта одна игра не была выиграна из-за" +
			" проблем с подключением к Интернету. В течение этих нескольких месяцев у многих людей возникли" +
			" подозрения, что позади аккаунта Master-а стоит система искусственного интеллекта. Но что это за" +
			" система и кто является ее разработчиками оставалось под покровом тайны до недавнего времени. И" +
			" только на прошлой неделе Демис Хассабис (Demis Hassabis), один из основателей проекта DeepMind," +
			" признался в своем твите, что за Master-ом скрывается очередная версия системы искусственного" +
			" интеллекта AlphaGo.",
		Tags: []string{"Other"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news4))

	//add courses

	course1 := m.Course{
		Id:		1,
		Title:		"Курс 1",
		Description:	"Описание курса",
		Logo:		"/files/course1.jpg",
		Rating:		4.8,
	}
	fastCheckErr(o.Insert(&course1))

	course2 := m.Course{
		Id:		2,
		Title:		"Курс 2",
		Description:	"Описание курса",
		Logo:		"/files/course2.jpg",
		Rating:		4.5,
	}
	fastCheckErr(o.Insert(&course2))

	course3 := m.Course{
		Id:		3,
		Title:		"Курс 3",
		Description:	"Описание курса",
		Logo:		"/files/course3.jpg",
		Rating:		4.6,
	}
	fastCheckErr(o.Insert(&course3))

	course4 := m.Course{
		Id:		4,
		Title:		"Курс 4",
		Description:	"Описание курса",
		Logo:		"/files/course4.jpg",
		Rating:		3.9,
	}
	fastCheckErr(o.Insert(&course4))

	//add statistic

	statistic1 := m.Statistic{
		Id:       1,
		Hours:    12,
		CourseId: &course1,
	}
	fastCheckErr(o.Insert(&statistic1))

	statistic2 := m.Statistic{
		Id:       2,
		Hours:    15,
		CourseId: &course2,
	}
	fastCheckErr(o.Insert(&statistic2))

	statistic3 := m.Statistic{
		Id:       3,
		Hours:    15,
		CourseId: &course3,
	}
	fastCheckErr(o.Insert(&statistic3))

	statistic4 := m.Statistic{
		Id:       4,
		Hours:    18,
		CourseId: &course4,
	}
	fastCheckErr(o.Insert(&statistic4))

	//add lessons

	lesson1 := m.Lesson{
		Id:          1,
		Title:       "Урок1",
		CourseId:    &course1,
		Description: "Урок 1 курс 1",
		Rating:      5,
	}
	fastCheckErr(o.Insert(&lesson1))

	lesson2 := m.Lesson{
		Id:          2,
		Title:       "Урок2",
		CourseId:    &course1,
		Description: "Урок 2 курс 1",
		Rating:      5,
	}
	fastCheckErr(o.Insert(&lesson2))

	lesson3 := m.Lesson{
		Id:          3,
		Title:       "Урок3",
		CourseId:    &course1,
		Description: "Урок 3 курс 1",
		Rating:      3,
	}
	fastCheckErr(o.Insert(&lesson3))

	lesson4 := m.Lesson{
		Id:          4,
		Title:       "Урок1",
		CourseId:    &course2,
		Description: "Урок 1 курс 2",
		Rating:      4,
	}
	fastCheckErr(o.Insert(&lesson4))

	lesson5 := m.Lesson{
		Id:          5,
		Title:       "Урок1",
		CourseId:    &course4,
		Description: "Урок 1 курс 4",
		Rating:      5,
	}
	fastCheckErr(o.Insert(&lesson5))

	//add recommend courses

	rec_course1 := m.RecomendCourses{
		Id:       1,
		CourseId: &course1,
		Link:     "/link_to_course1/",
	}
	fastCheckErr(o.Insert(&rec_course1))

	rec_course2 := m.RecomendCourses{
		Id:       2,
		CourseId: &course2,
		Link:     "/link_to_course2/",
	}
	fastCheckErr(o.Insert(&rec_course2))

	rec_course3 := m.RecomendCourses{
		Id:       3,
		CourseId: &course3,
		Link:     "/link_to_course3/",
	}
	fastCheckErr(o.Insert(&rec_course3))

	rec_course4 := m.RecomendCourses{
		Id:       4,
		CourseId: &course4,
		Link:     "/link_to_course4/",
	}
	fastCheckErr(o.Insert(&rec_course4))

	//add tests
/*
	test1 := m.Test{
		Id: 		1,
		Title:		"Тест 1 урок 1 курс 1",
		LessonId:	&lesson1,
	}
	fastCheckErr(o.Insert(&test1))

	test2 := m.Test{
		Id: 		2,
		Title:		"Тест 2 урок 1 курс 1",
		LessonId:	&lesson2,
	}
	fastCheckErr(o.Insert(test2))

	test3 := m.Test{
		Id: 		3,
		Title:		"Тест 1 урок 3 курс 1",
		LessonId:	&lesson3,
	}
	fastCheckErr(o.Insert(&test3))

	test4 := m.Test{
		Id: 		4,
		Title:		"Тест 1 урок 1 курс 4",
		LessonId:	&lesson5,
	}
	fastCheckErr(o.Insert(&test4))

	//add task for test

	task_fortest1 := m.TaskForTest{
		Id:		1,
		Question:	"Вопрос 1 Тест 1 Урок 1 Курс 1",
		TestId:		&test1,
	}
	fastCheckErr(o.Insert(&task_fortest1))

	task_fortest2 := m.TaskForTest{
		Id:		2,
		Question:	"Вопрос 2 Тест 1 Урок 1 Курс 1",
		TestId:		&test1,
	}
	fastCheckErr(o.Insert(&task_fortest2))

	task_fortest3 := m.TaskForTest{
		Id:		3,
		Question:	"Вопрос 1 Тест 2 Урок 1 Курс 1",
		TestId:		&test3,
	}
	fastCheckErr(o.Insert(&task_fortest3))

	task_fortest4 := m.TaskForTest{
		Id:		4,
		Question:	"Вопрос 1 Тест 1 Урок 1 Курс 4",
		TestId:		&test4,
	}
	fastCheckErr(o.Insert(&task_fortest4))

	//add Practises
*/
	practise1 := m.Practise{
		Id:		1,
		LessonId: 	&lesson1,
		Description: 	"Упражнение к уроку 1 Курс 1",
	}
	fastCheckErr(o.Insert(&practise1))

	practise2 := m.Practise{
		Id:		2,
		LessonId: 	&lesson2,
		Description: 	"Упражнение к уроку 1 Курс 1",
	}
	fastCheckErr(o.Insert(&practise2))

	practise3 := m.Practise{
		Id:		3,
		LessonId: 	&lesson3,
		Description: 	"Упражнение к уроку 1 Курс 2",
	}
	fastCheckErr(o.Insert(&practise3))

	practise4 := m.Practise{
		Id:		4,
		LessonId: 	&lesson4,
		Description: 	"Упражнение к уроку 1 Курс 4",
	}
	fastCheckErr(o.Insert(&practise4))

	//add Videos

	video1 := m.Video{
		Id:		1,
		LessonId:	&lesson1,
		Link:		"/link_to_video1/",
	}
	fastCheckErr(o.Insert(&video1))

	video2 := m.Video{
		Id:		2,
		LessonId:	&lesson2,
		Link:		"/link_to_video2/",
	}
	fastCheckErr(o.Insert(&video2))

	video3 := m.Video{
		Id:		3,
		LessonId:	&lesson3,
		Link:		"/link_to_video3/",
	}
	fastCheckErr(o.Insert(&video3))

	video4 := m.Video{
		Id:		4,
		LessonId:	&lesson4,
		Link:		"/link_to_video4/",
	}
	fastCheckErr(o.Insert(&video4))

	video5 := m.Video{
		Id:		5,
		LessonId:	&lesson5,
		Link:		"/link_to_video5/",
	}
	fastCheckErr(o.Insert(&video5))

	//add Variants
/*
	variant1 := m.Variant{
		Id :	400,
		Text:	"Вариант  ответа 1",
		CorrectAnswer: false,
		TaskForTestId: &task_fortest1,
	}
	fastCheckErr(o.Insert(&variant1))

	variant2 := m.Variant{
		Id :	401,
		Text:	"Вариант  ответа 2",
		CorrectAnswer: true,
		TaskForTestId: &task_fortest1,
	}
	fastCheckErr(o.Insert(&variant2))

	variant3 := m.Variant{
		Id :	402,
		Text:	"Вариант  ответа 1",
		CorrectAnswer: true,
		TaskForTestId: &task_fortest2,
	}
	fastCheckErr(o.Insert(&variant3))

	variant4 := m.Variant{
		Id :	403,
		Text:	"Вариант  ответа 3",
		CorrectAnswer: false,
		TaskForTestId: &task_fortest2,
	}
	fastCheckErr(o.Insert(&variant4))

	variant5 := m.Variant{
		Id :	404,
		Text:	"Вариант  ответа 1",
		CorrectAnswer: false,
		TaskForTestId: &task_fortest3,
	}
	fastCheckErr(o.Insert(&variant5))

	variant6 := m.Variant{
		Id :	405,
		Text:	"Вариант  ответа 2",
		CorrectAnswer: true,
		TaskForTestId: &task_fortest3,
	}
	fastCheckErr(o.Insert(&variant6))
*/

	beego.Info("Initial data was successfully added to Database")
}
