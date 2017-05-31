package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"time"
)

type ProjectEnroll struct {
	Id          int         `orm:"column(id);pk;auto"`
	UserId      *User       `orm:"column(user_id);rel(fk)"`
	ProjectId   *Project    `orm:"column(project_id);rel(fk)"`
	Message     string      `orm:"column(enrolling_message)"`
	Time        time.Time   `orm:"column(time);type(datetime)"`
}

func (t *ProjectEnroll) TableName() string {
	return "project_enroll"
}

func init() {
	orm.RegisterModel(new(ProjectEnroll))
}

// GetProjectEnrollIdByUserId returns an array of projects
// where user enrolls
func GetProjectEnrollIdByUserId(userId int) (projects []*Project, err error){
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
		UserId: u,
		ProjectId: &temp,
		Message: message,
		Time: time.Now(),
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

// Получить информацию о пользователях, подавших заявку на проект. Запрос исходит от конкретного пользователя.
func GetAllEnrolledOnProject(project_id, master_id int) (ml []interface{}, err error) {
	if IsUserIsMasterForProject(master_id, project_id) {
		o := orm.NewOrm()
		var wtf []ProjectEnroll
		_, err := o.QueryTable(new(ProjectEnroll)).Filter("ProjectId", project_id).RelatedSel().All(&wtf)
		if err != nil {
			return nil, err
		} else {
			for _, r := range wtf {
				contacts, err := GetAllUserContacts(r.UserId.Id)
				if err != nil {
					contacts = nil
				}
				ml = append(ml, ObjectOfListOfEnrolledUsersOnProject{
					User: MainUserInfo{
						Id: r.UserId.Id,
						Avatar: r.UserId.Avatar,
						Nickname: r.UserId.Nickname,
					},
					Message: r.Message,
					Time: r.Time,
					Contacts: contacts,
				})
			}
			return ml, nil
		}
	} else {
		return nil, errors.New("Not a master of the project")
	}
	return
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
