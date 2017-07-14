package controllers

import (
	"data-service/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"

	"github.com/astaxie/beego"
)

// Операции с models.User, для некоторых требуется авторизация
type UserController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

type projectsGetter func(int) ([]*models.Project, error)

// Вызывает функцию с указанным Id пользователя и отсылает в канал полученных пользователей в сокращенном виде
// Используется для параллельного запроса к Masters, Enrolled и Users для проекта
// Функция должна соответствовать projectsGetter прототипу
func CallForProjectMainInfo(f projectsGetter, id int, c chan []models.MainProjectInfo, secondChan chan []models.MainProjectInfo) {
	users, err := f(id)
	if err != nil {
		c <- nil
		if secondChan != nil {
			secondChan <- nil
		}
	} else {
		var mainProjectInfo []models.MainProjectInfo
		for _, u := range users {
			mainProjectInfo = append(mainProjectInfo, models.MainProjectInfo{
				Id:   u.Id,
				Logo: u.Logo,
				Name: u.Name,
			})
		}
		c <- mainProjectInfo
		if secondChan != nil {
			secondChan <- mainProjectInfo
		}
	}
}

func CallForProjectApplications(c chan []models.ProjectApplications, projects_chan chan []models.MainProjectInfo) {
	projects := <-projects_chan
	if projects == nil || len(projects) == 0 {
		c <- nil
	} else {
		// Создаем массив, элементы которого уже инициализированы,
		// так возможна непоследовательная запись
		apps := make([]models.ProjectApplications, len(projects))
		// Группа ожидания синхронизации потоков
		var wg sync.WaitGroup
		for index, project := range projects {
			t := make(chan []interface{})
			wg.Add(1)
			go func() {
				// конкурентным способом получаем все заявки
				go models.GetAllEnrolledOnProjectWithoutAuthChecking(project.Id, t)
				apps[index].Project = project
				apps[index].Applications = <-t
				wg.Done()
			}()
		}
		wg.Wait()
		c <- apps
	}
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param   id  path    string  true    "ID пользователя, о котором нужно узнать информацию"
// @Param   cut query   bool    false   "Оставить только информацию о пользователе?"
// @Param   Bearer-token        header      string          true    "Токен"
// @Success 200 {object} models.User
// @Failure 400 :id is empty string
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		v, err := models.GetUserById(id)
		if err != nil {
			beego.Debug("GetOne user id not found", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		} else {
			c.Data["json"] = v
			/*beego.Trace("User founded, search for contacts")
			U := models.FullUserInfo{
				Id:          v.Id,
				Nickname:    v.Nickname,
				Description: v.Description,
				Avatar:      v.Avatar,
			}
			is_master, err := models.IsProjectMasterForUserById(id, c.CurrentUser.UserId)
			if err == nil {
				if contact, err := models.GetAllUserContacts(id); err != nil {
					beego.Critical("GetOne user Contacts GetAllUserContact error ", err.Error())
					//c.Data["json"] = err.Error()
					//c.Ctx.Output.SetStatus(500)
				} else {
					beego.Trace("Contacts founded")
					if c.CurrentUser.UserId == v.Id || c.CurrentUser.PermissionLevel == ADMIN || is_master {
						beego.Trace("Contacts pinned to response")
						U.Contact = contact
					}
				}
			} else {
				beego.Critical("Error in IsProjectMasterForUserById in User `GetOne(ID)` ", err.Error())
				//c.Data["json"] = err.Error()
				//c.Ctx.Output.SetStatus(500)
			}
			// TODO: refactor this govnocode
			enrolledChan := make(chan []models.MainProjectInfo)
			membersChan := make(chan []models.MainProjectInfo)
			mastersChan := make(chan []models.MainProjectInfo)
			applicationsChan := make(chan []models.ProjectApplications)
			chan_for_apps := make(chan []models.MainProjectInfo)
			beego.Trace("Search for projects")
			if cut_info, _ := c.GetBool("cut"); !cut_info {
				go CallForProjectMainInfo(models.GetProjectEnrollIdByUserId, id, enrolledChan, nil)
				go CallForProjectMainInfo(models.GetProjectUserIdByUserId, id, membersChan, nil)
				go CallForProjectMainInfo(models.GetProjectMasterIdByUserId, id, mastersChan, chan_for_apps)
				go CallForProjectApplications(applicationsChan, chan_for_apps)
			} else {
				go func() {
					enrolledChan <- nil
					mastersChan <- nil
					membersChan <- nil
					applicationsChan <- nil
				}()
			}
			beego.Trace("Ready to sent response")
			c.Data["json"] = models.AllInformationAboutUser{
				User:           U,
				EnrolledOn:     <-enrolledChan,
				MasterOf:       <-mastersChan,
				MemberOf:       <-membersChan,
				MyApplications: <-applicationsChan,
			}
			beego.Trace("Get user OK")*/
		}
	} else {
		beego.Debug("GetOne user `Atoi` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 400
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Param   Bearer-token        header      string          true    "Access token, Permission Level should be 2"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	if c.CurrentUser.PermissionLevel != VIEWER {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Put user `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		} else if c.CurrentUser.UserId == id {
			v := models.User{Id: id}
			if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
				v.Id = id
				if err := models.UpdateUserById(&v); err == nil {
					beego.Trace("Put user OK")
					c.Data["json"] = HTTP_OK_STR
				} else {
					beego.Debug("Put user `UpdateUserById` error", err.Error())
					c.Data["json"] = err.Error()
					c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
				}
			} else {
				beego.Debug("Put user `Unmarshal` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			}
		}
	} else {
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR
	}
	c.ServeJSON()
}
