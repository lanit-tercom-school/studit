package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type ProjectMaster struct {
	Id         int       `orm:"column(id);pk;auto"`
	ProjectId  *Project  `orm:"column(project_id);rel(fk)"`
	MasterId   *User     `orm:"column(master_id);rel(fk)"`
	SignedDate time.Time `orm:"column(signed_date);type(datetime)"`
}

func (t *ProjectMaster) TableName() string {
	return "project_master"
}

func init() {
	orm.RegisterModel(new(ProjectMaster))
}

// GetProjectMasterIdByUserId return array of projects
// where user is master
func GetProjectMasterIdByUserId(userId int) (projectid []int, err error){
	o := orm.NewOrm()
	var projects []ProjectMaster
	_, err = o.QueryTable(new(ProjectMaster)).Filter("MasterId", User{Id: userId}).RelatedSel().All(&projects)
	if err != nil {
		return projectid, err
	}
	for _, v := range projects {
		projectid = append(projectid, v.ProjectId.Id)
	}
	return projectid, nil
}

// IsProjectMasterUserById returns true
// if master is master of userproject
func IsProjectMasterForUserById(userId int, masterId int)(masterOfUser bool, err error) {
	userProjects, err := GetProjectUserIdByUserId(userId)
	if err != nil {
		return false, err
	}
	masterProjects, err := GetProjectMasterIdByUserId(masterId)
	if err != nil {
		return false, err
	}
	// finding intersection
	target := make(map[int]bool)
	for _, v := range masterProjects {
		target[v] = true
	}
	for _, v := range userProjects {
		if target[v] {
			return true, nil
		}
	}

	return false, nil
}

// AddProjectUser insert a new ProjectMaster into database and returns
// last inserted Id on success.
func AddMasterToProject(user *User, project *ProjectJson) (err error) {
	temp := project.translate()
	m := ProjectMaster{
		MasterId:   user,
		ProjectId:  &temp,
		SignedDate: time.Now(),
	}
	_, err = orm.NewOrm().Insert(&m)
	return
}

// Проверяет, содержится ли в списке пользователей пользователь с указанным ID
func IsUserInArray(user_id int, users []*User) bool {
	for _, x := range users {
		if x.Id == user_id{
			return true
		}
	}
	return false
}

func GetMastersOfTheProject(project_id int) (masters []*User, err error) {
	o := orm.NewOrm()
	var connections []ProjectMaster
	// выбираем всех пользователей, являющихся мастерами данного проекта
	_, err = o.QueryTable(new(ProjectMaster)).
			Filter("project_id", project_id).
			RelatedSel().
			All(&connections, "MasterId")
	// возвращаем только мастеров
	for _, x := range connections {
		masters = append(masters, x.MasterId)
	}
	return
}


// GetAllProjectUser retrieves all ProjectMaster matches certain condition. Returns empty list if
// no records exist
func GetAllProjectMaster() (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProjectMaster))
	var l []ProjectMaster
	if _, err = qs.All(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}


// DeleteProjectUser deletes ProjectMaster by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMasterFromProject(master_id int, project_id int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(ProjectMaster)).
		Filter("MasterId", master_id).
		Filter("ProjectId", project_id).
		Delete()
	return
}
