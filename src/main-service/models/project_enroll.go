package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProjectEnroll struct {
	Id        int       `orm:"column(id);pk;auto"`
	UserId    *User     `orm:"column(user_id);rel(fk)"`
	ProjectId *Project  `orm:"column(project_id);rel(fk)"`
	Message   string    `orm:"column(enrolling_message)"`
	Time      time.Time `orm:"column(time);type(datetime)"`
}

func (t *ProjectEnroll) TableName() string {
	return "project_enroll"
}

func init() {
	orm.RegisterModel(new(ProjectEnroll))
}

// GetProjectEnrollIdByUserId returns an array of projects
// where user enrolls
func GetProjectEnrollIdByUserId(userId int) (projects []*Project, err error) {
	o := orm.NewOrm()
	var project_enrolled []ProjectEnroll
	_, err = o.QueryTable(new(ProjectEnroll)).Filter("UserId", User{Id: userId}).RelatedSel().All(&project_enrolled)
	if err != nil {
		return projects, err
	}
	for _, v := range project_enrolled {
		projects = append(projects, v.ProjectId)
	}
	return projects, nil
}

// AddProjectAuthor insert a new ProjectEnroll into database and returns
// last inserted Id on success.
func AddApplicationFromUserForProject(u *User, p *ProjectJson, message string) (id int64, err error) {
	temp := p.translate()
	m := ProjectEnroll{
		UserId:    u,
		ProjectId: &temp,
		Message:   message,
		Time:      time.Now(),
	}
	id, err = orm.NewOrm().Insert(&m)
	return
}

// Получить список всех пользователей, подавших заявку на проект
func GetAllSignedUpOnProject(project_id int) (ml []*User, err error) {
	o := orm.NewOrm()
	var singed_up []ProjectEnroll
	_, err = o.QueryTable(new(ProjectEnroll)).
		Filter("ProjectId", project_id).
		RelatedSel().
		All(&singed_up)
	if err != nil {
		return nil, err
	}
	for _, x := range singed_up {
		ml = append(ml, x.UserId)
	}
	return ml, nil
}

// Объект списка для пользователя, подавшего заявку
// для удобного просмотра списка пользователей, кто подал заявку
// туда добавляются контакты, текст и небольшое сообщение от самого пользователя
type ObjectOfListOfEnrolledUsersOnProject struct {
	User     MainUserInfo   `json:"user"`
	Contacts []*UserContact `json:"contacts"`
	Message  string         `json:"message"`
	Time     time.Time      `json:"date"`
}

// Для выдачи мастеру проекту, содержит информацию о проекте и список заявок
type ProjectApplication struct {
	Project Project `json:"project,omitempty"`
	User    User    `json:"user,omitempty"`
	Message string  `json:"message"`
}
type ProjectApplications struct {
	Project      MainProjectInfo `json:"project,omitempty"`
	Applications []interface{}   `json:"apps,omitempty"`
}

// Получить информацию о пользователях, подавших заявку на проект. Запрос исходит от конкретного пользователя.
func GetAllEnrolledOnProject(master_id int) (pa []ProjectApplication, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	var n int64
	n, err = o.Raw("Select  pe.enrolling_message,p.id as project_id,p.name as project_name, p.tags as project_tags,p.status as project_status,p.logo as project_logo,p.description as project_description, p.date_of_creation as project_date,u.id as user_id,u.nickname as user_name, u.avatar as user_avatar, u.description as user_description from public.project_user pu inner join public.project_enroll pe on pu.project_id=pe.project_id inner join public.user u on pe.user_id=u.id inner join public.project p on p.id=pe.project_id where pu.user_id=?", master_id).Values(&maps)
	if err != nil {
		beego.Debug("Something wrong with database request")
		return
	} else {
		pa = make([]ProjectApplication, n)
		for i, v := range maps {
			pa[i].User.Id, err = strconv.Atoi(v["user_id"].(string))
			if err != nil {
				beego.Debug("Error converting to int")
				return
			}
			pa[i].User.Nickname = v["user_name"].(string)
			pa[i].User.Avatar = v["user_avatar"].(string)
			pa[i].User.Description = v["user_description"].(string)
			pa[i].Project.Id, err = strconv.Atoi(v["project_id"].(string))
			if err != nil {
				beego.Debug("Error converting to int")
				return
			}
			pa[i].Project.Name = v["project_name"].(string)
			pa[i].Project.Description = v["project_description"].(string)
			pa[i].Project.Logo = v["project_logo"].(string)
			pa[i].Project.Status, err = strconv.Atoi(v["project_status"].(string))
			if err != nil {
				beego.Debug("Error converting to int")
				return
			}
			pa[i].Project.Tags = v["project_tags"].(string)
			pa[i].Message = v["enrolling_message"].(string)
			pa[i].Project.DateOfCreation, err = time.Parse(time.RFC3339, v["project_date"].(string))
			if err != nil {
				beego.Debug("Error converting to time")
				return
			}
		}
	}
	return
}

// Получает записанных без проверки на то, что пользователь является мастером проекта конкурентным способом
func GetAllEnrolledOnProjectWithoutAuthChecking(project_id int, c chan []interface{}) {
	o := orm.NewOrm()
	var wtf []ProjectEnroll
	_, err := o.QueryTable(new(ProjectEnroll)).Filter("ProjectId", project_id).RelatedSel().All(&wtf)
	if err != nil {
		c <- nil
	} else {
		var ml []interface{}
		for _, r := range wtf {
			contacts, err := GetAllUserContacts(r.UserId.Id)
			if err != nil {
				contacts = nil
			}
			ml = append(ml, ObjectOfListOfEnrolledUsersOnProject{
				User: MainUserInfo{
					Id:       r.UserId.Id,
					Avatar:   r.UserId.Avatar,
					Nickname: r.UserId.Nickname,
				},
				Message:  r.Message,
				Time:     r.Time,
				Contacts: contacts,
			})
		}
		c <- ml
	}
}

// UpdateProjectAuthor updates ProjectEnroll by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectAuthorById(m *ProjectEnroll) (err error) {
	o := orm.NewOrm()
	v := ProjectEnroll{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Update(m)
	}
	return
}

// DeleteProjectSignUp deletes ProjectEnroll by Project Id and returns error if
// the record to be deleted doesn't exist
func DeleteProjectSignUp(user_id, project_id int) (err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable(new(ProjectEnroll)).
		Filter("UserId", user_id).
		Filter("ProjectId", project_id).
		Delete()
	return
}
