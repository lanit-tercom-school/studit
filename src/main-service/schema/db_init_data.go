package main

import (
	_ "main-service/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	m "main-service/models"
	"time"
	"github.com/astaxie/beego"
	"fmt"
	 auth "main-service/auth"
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
		Logo: "http://85.143.214.42/files/1.jpg",
		Tags: "studit,summerschool",
		Status: 1,
	}
	fastCheckErr(o.Insert(&project1))

	project2 := m.Project{
		Id: 2,
		Name: "Модный фрилансер",
		Description: "Какие же стрелочки вокруг ноубука!",
		DateOfCreation: time.Now(),
		Logo: "http://85.143.214.42/files/2.jpg",
		Tags: "freelance",
		Status: 0,
	}
	fastCheckErr(o.Insert(&project2))

	project3 := m.Project{
		Id: 3,
		Name: "Оригинальное название",
		Description: "Click-bait описание",
		DateOfCreation: time.Now(),
		Logo: "http://85.143.214.42/files/3.jpg",
		Tags: "creative",
		Status: 0,
	}
	fastCheckErr(o.Insert(&project3))

	project4 := m.Project{
		Id: 4,
		Name: "TFS Mobile",
		Description: "Студенческий проект TFS Mobile",
		DateOfCreation: time.Now(),
		Logo: "http://www.carlthomasiv.com/wp-content/uploads/2012/08/tfs-logo2-318x235.jpg",
		Tags: "TFS,summerschool",
		Status: 0,
	}
	fastCheckErr(o.Insert(&project4))

	project5 := m.Project{
		Id: 5,
		Name: "Еще один проект",
		Description: "Описаниеописаниеописание",
		DateOfCreation: time.Now(),
		Logo: "https://www.glidetraining.com/wp-content/uploads/2014/03/Microsoft-Office-Project-2013.png",
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
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "Главный по тарелкам",
	}
	fastCheckErr(o.Insert(&user1))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user2 := m.User{
		Id:       2,
		Nickname: "Moderator",
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "Главный по молоткам",
	}
	fastCheckErr(o.Insert(&user2))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user3 := m.User{
		Id: 3,
		Nickname: "Егорка2003",
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "ЮЮЮ, ААА",
	}
	fastCheckErr(o.Insert(&user3))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user4 := m.User{
		Id: 4,
		Nickname: "ZagadOchNayA",
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "Легко потерять, невозможно забить",
	}
	fastCheckErr(o.Insert(&user4))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user5 := m.User{
		Id: 5,
		Nickname: "S1ayeR1",
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "bjklknufu",
	}
	fastCheckErr(o.Insert(&user5))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user6 := m.User{
		Id: 6,
		Nickname: "NekoTyan",
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "^_^",
	}
	fastCheckErr(o.Insert(&user6))

	avatar_seed = auth.GenerateNewToken(6)
	color_str = auth.GenerateRandomColor()
	user7 := m.User{
		Id: 7,
		Nickname: "B",
		Avatar: fmt.Sprintf("%s%s?colors=%s&colors=%s&size=%s", auth.AvatarTemplatePath, avatar_seed,
			color_str, "FFFFFF", auth.AvatarTemplateSize),
		Description: "BBB",
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


	news5 := m.NewsJson{
		Id: 5,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Яндeкc пpидумaл, кaк oбoйти блoкиpoвку пoпуляpныx pecуpcoв",
		Description: "Koмпaния Яндeкc нe cмoглa cмиpитьcя c бoльшим oттoкoм cвoиx пoльзoвaтeлeй," +
		 "cвязaнным c ввeдeниeм oчepeднoй пopции caнкций co cтopoны укpaинcкиx влacтeй. Moбильный"+
		"бpaузep для Aндpoид пoлучил cпeциaльнoe oбнoвлeниe, пpeдocтaвляющee вoзмoжнocть oбxoдa"+
		 "зaпpeтoв.",
		Tags: []string{"Other"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news5))

    news6 := m.NewsJson{
		Id: 6,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Лаборатория Касперского заявила о готовности раскрыть властям США программный код",
		Description: "Kaspersky Lab (Лаборатория Касперского) готова обнародовать американским "+
		"властям исходный код программ, чтобы избавить от подозрений по поводу работы компании, "+ 
		"сообщил агентству Ассошиэйтед Пресс исполнительный директор Kaspersky Lab Евгений "+
		"Касперский. Касперский отметил, что готов дать показания в Конгрессе США и доказать,"+
		"что компания не действует злонамеренно. Директор Kaspersky Lab подтвердил сообщение "+
		"NBC о том, что агенты ФБР провели беседы с несколькими его сотрудниками по вопросам "+
		"деятельности компании в США. Телеканал отмечает, что неизвестно, связано ли это с "+
		"расследованием вмешательства России в президентские выборы в США. По словам Касперского, "+
		"он не знает, чем именно интересовались сотрудники Бюро. Он подчеркнул, что компания "+
		"занимается исключительно вопросами кибербезопасности, но допустил, что власти некоторых "+
		"неназванных стран, но не России, пытались склонить ее к хакерской деятельности, которую "+
		"он назвал темной стороной",
		Tags: []string{"Other"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news6))


	news7 := m.NewsJson{
		Id: 7,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Стала известна неустранимая опасность от выключенного смартфона ",
		Description: "Ученые из Техасского университета в Остине (США) обнаружили новую опасность"+
		"от смартфонов, которую невозможно убрать даже их выключением. Соответствующее исследование"+
		"опубликовано в Journal of the Association for Consumer Research, кратко о нем сообщает"+
		"издание Medical Xpress. Специалисты полагают, что присутствие смартфона отвлекает человека,"+
	    "и, как следствие, приводит к снижению его когнитивных способностей, прежде всего возможности "+
		"быстро считать в уме. Данное наблюдение сохраняет свою силу и в случае, когда гаджет выключен."+
		"К подобным выводам авторы пришли, проведя тестирование когнитивных способностей различных "+
		"людей, у которых тем или иным способом был ограничен доступ к собственным смартфонам. "+
		"Наилучшие результаты продемонстрировали участники эксперимента, гаджеты которых находились "+
		"в других помещениях. Причиной наблюдаемой закономерности специалисты называют то, что часть "+
		"умственных усилий человека отводится на слежение за смартфоном. В случае если участника "+
		"эксперимента гаджет вообще не отвлекал (известно, что в помещении, в котором он находится, "+
		"смартфона нет), человек вообще не тратил сил на наблюдение за телефоном и лучше "+
		"концентрировался над поставленными задачами, чем остальные.",
		Tags: []string{"Other"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news7))

		news8 := m.NewsJson{
		Id: 8,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Умные девайсы следят за людьми",
		Description: " IoT-устройства — отдельные части «умного» дома, о котором все уже хорошо "+
		"знают. Мы покупаем IoT-устройства и даже не подозреваем, что эти «помощники» могут стать "+
		"инструментом тайной слежки или даже атаки. Рассказываем, как настроить «умные» гаджеты, "+
		"чтобы они приносили исключительно пользу. Предметы «умного» дома призваны помогать своим "+
		"занятым владельцам в быту: камеры следят за безопасностью в доме, датчики дыма или воды "+
		"оповещают владельца о состоянии дома и т.д. Использование потенциала Интернета вещей для "+
		"экономического и социального блага в ближайшие десятилетия будет одной из основных задач "+
		"общества, включая проблемы и возможности, вытекающие из этого явления. Повсеместное и "+
		"небезопасное использование IoT-устройств уже отмечено такой глобальной атакой, как "+
		"показательный случай с ботнетом Mirai. Он был использован для DDoS-атак, в результате "+
		"которых был выведен из строя сегмент сети Интернет в США, а также оказался недоступен "+
		"Twitter. Эта атака достигла рекордных показателей по объему генерируемого трафика.",
		Tags: []string{"Other"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news8))

		news9 := m.NewsJson{
		Id: 9,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Обновление Windows уничтожило компьютеры",
		Description: " Корпорация Microsoft представила тестовую сборку Windows 10 16212 для "+
		"участников программы Windows Insider. Обновление превращает в кирпичи как компьютеры, "+
		"так и мобильные устройства, предупреждает MSPowerUser. По данным источников издания, "+
		"после установки обновления устройство переходит в бесконечный цикл перезагрузки, и вернуть "+
		"его к жизни можно только через полный сброс и последующее восстановление. При этом "+
		"неизбежно происходит потеря важных пользовательских данных. Помимо участников программы "+
		"Windows Insider (причем речь идет о пользователях как с ранним доступом Fast Ring и "+
		"поздним Slow Ring), обновление получили обладатели обычных версий Windows, не принимающие "+
		"участие в тестировании. Всем, на чьи девайсы начало загружаться обновление, рекомендуется "+
		"отключить устройство от интернета. Если же апдейт загрузился и готов к установке, достаточно "+
		"поменять в настройках текущий год на 2050 - это позволит отсрочить применение обновления.",
		Tags: []string{"Other"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news9))

	news10 := m.NewsJson{
		Id: 10,
		Created: time.Now(),
		Edited: time.Now(),
		Title: "Россия заняла второе место в мире по дешевизне связи и Интернета,",
		Description: "Согласно исследованиям Всемирного экономического форума, "+
		"Россия занимает второе место в мире после Гонконга по дешевизне услуг связи и Интернета. "+
		"Об этом сегодня на коллегии Минсвязи РФ заявил министр связи Николай Никифоров. "+
		"«При этом надо учитывать, что по площади  Гонконг равен Москве, а по населению – половине "+
		"Москвы», — сказал Никифоров. По его словам, технология связи четверного поколения LTE уже "+
		"доступна на территории, где проживает 70% населения России. «Сейчас в России 120 тыс базовых "+
		"станций формата LTE. Это четверть всех базовых станций сотовой связи в РФ. Мы должны выйти "+
		"на уровень доступности современных средств связи – 97%.», — сказал Никифоров. Министр "+
		"сообщил, что в рамках устранения цифрового неравенства в России построено 35 тыс км "+
		"волоконно-оптических линий связи. Работы идут с 71 субъекте. Особо Никифоров отметил "+
		"прокладку подводной линии связи от Сахалина до Магадана и от Сахалина до Камчатки «Если "+
		"в 2012 было 32 населенных пункта с более 10 тыс населения, куда не приходила волоконно- "+
		"оптическая сеть. Сейчас осталось 4 таких пункта. И по каждому из них у нас есть понимание, "+ 
		"когда туда придет оптово-волоконная линия связи», — сказал Никифоров.",
		Tags: []string{"Other"},
		Image: "http://frank.jou.ufl.edu/wp-content/uploads/2015/05/ThinkstockPhotos-Website.jpg",
	}
	fastCheckErr(m.AddNews(&news10))
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
