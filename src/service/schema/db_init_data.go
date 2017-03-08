package main

import (
	_ "service/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	m "service/models"
	"log"
	"time"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
}

func fastCheckErr(_ int64, err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	o := orm.NewOrm()
	o.Using("default")

	// add projects

	project1 := m.Project{
		Id: 1,
		Name: "Образовательный портал Studit",
		Description: "Разработка образовательного портала для Lanit-Tercom School",
		Logo: "/logo/1.jpg",
	}
	fastCheckErr(o.Insert(&project1))

	project2 := m.Project{
		Id: 2,
		Name: "Модный фрилансер",
		Description: "Какие же стрелочки вокруг ноубука!",
		Logo: "/logo/2.jpg",
	}
	fastCheckErr(o.Insert(&project2))

	project3 := m.Project{
		Id: 3,
		Name: "Оригинальное название",
		Description: "Click-bait описание",
		Logo: "/logo/3.jpg",
	}
	fastCheckErr(o.Insert(&project3))

	// add users

	user1 := m.User{
		Id: 1,
		Nickname: "Admin",
		Login: "admin@admin.admin",
		Password: "admin",
		Avatar: "/logo/1.jpg",
		Description: "Главный по тарелкам",
	}
	fastCheckErr(o.Insert(&user1))

	user2 := m.User{
		Id: 2,
		Nickname: "Moderator",
		Login: "moder@moder.moder",
		Password: "moder",
		Avatar: "/logo/2.jpg",
		Description: "Главный по молоткам",
	}
	fastCheckErr(o.Insert(&user2))

	user3 := m.User{
		Id: 3,
		Nickname: "Егорка2003",
		Login: "egorka2003@maaail.ru",
		Password: "пароль",
		Avatar: "/logo/3.jpg",
		Description: "ЮЮЮ, ААА",
	}
	fastCheckErr(o.Insert(&user3))

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

	// add author permission

	author1 := m.Author{
		Id: 1,
		UserId: &user1,
	}
	fastCheckErr(o.Insert(&author1))

	// connect projects and authors

	prj_aut_con1 := m.ProjectAuthor{
		Id: 1,
		AuthorId: &author1,
		ProjectId: &project1,
	}
	fastCheckErr(o.Insert(&prj_aut_con1))

	prj_aut_con2 := m.ProjectAuthor{
		Id: 2,
		AuthorId: &author1,
		ProjectId: &project2,
	}
	fastCheckErr(o.Insert(&prj_aut_con2))

	prj_aut_con3 := m.ProjectAuthor{
		Id: 3,
		AuthorId: &author1,
		ProjectId: &project3,
	}
	fastCheckErr(o.Insert(&prj_aut_con3))

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

	// add user contacts

	user_contact1 := m.UserContact{
		Id: 1,
		UserId: &user1,
		Contact: "admin@admin.admin",
		ContactTypeId: &contactType1,
	}
	fastCheckErr(o.Insert(&user_contact1))

	user_contact2 := m.UserContact{
		Id: 2,
		UserId: &user2,
		Contact: "moder@moder.moder",
		ContactTypeId: &contactType1,
	}
	fastCheckErr(o.Insert(&user_contact2))

	user_contact3 := m.UserContact{
		Id: 3,
		UserId: &user3,
		Contact: "egorka2003@maaail.ru",
		ContactTypeId: &contactType1,
	}
	fastCheckErr(o.Insert(&user_contact3))

	user_contact4 := m.UserContact{
		Id: 4,
		UserId: &user1,
		Contact: "+7-123-456-78-90",
		ContactTypeId: &contactType2,
	}
	fastCheckErr(o.Insert(&user_contact4))

	// add news

	news1 := m.News{
		Id: 1,
		Date: time.Now(),
		Title: "Программисты признаются в своих грехах в знак протеста против собеседований" +
			" с кодингом «на бумаге» — в Твиттере появился новый флешмоб",
		Description: "Где проводят такие собеседования?	Такой стиль собеседований широко используется в" +
			" IT-индустрии, в том числе в таких компаниях, как Google и Amazon. Кандидатам не дают никакого доступа" +
			" к справочным материалам и просят решить какую-либо техническую задачу, что, по мнению уже благополучно" +
			" работающих где-либо программистов, деморализует и не выявляет реальных навыков.",
	}
	fastCheckErr(o.Insert(&news1))

	news2 := m.News{
		Id: 2,
		Date: time.Now(),
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
	}
	fastCheckErr(o.Insert(&news2))

	// add tags for news (only)

	n_tag1 := m.NewsTags{
		Id: 100,
		Text: "Other",
	}
	fastCheckErr(o.Insert(&n_tag1))

	n_tag2 := m.NewsTags{
		Id: 101,
		Text: "Blog",
	}
	fastCheckErr(o.Insert(&n_tag2))

	n_tag3 := m.NewsTags{
		Id: 102,
		Text: "World",
	}
	fastCheckErr(o.Insert(&n_tag3))

	n_tag4 := m.NewsTags{
		Id: 103,
		Text: "Technologies",
	}
	fastCheckErr(o.Insert(&n_tag4))

	// connect news tags with news

	news_tag_con1 := m.NewsNewsTags{
		Id: 1,
		NewsId: &news1,
		NewsTagsId: &n_tag1,
	}
	fastCheckErr(o.Insert(&news_tag_con1))

	news_tag_con2 := m.NewsNewsTags{
		Id: 2,
		NewsId: &news2,
		NewsTagsId: &n_tag1,
	}
	fastCheckErr(o.Insert(&news_tag_con2))

	news_tag_con3 := m.NewsNewsTags{
		Id: 3,
		NewsId: &news2,
		NewsTagsId: &n_tag3,
	}
	fastCheckErr(o.Insert(&news_tag_con3))

	log.Print("Initial data was successfully added to Database")
}