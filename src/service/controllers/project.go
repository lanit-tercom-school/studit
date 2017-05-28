package controllers

import (
	"encoding/json"
	"errors"
	"service/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

// Создание, изменение, удаление и просмотр проектов
type ProjectController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *ProjectController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description Создать новый проект
// @Param   body            body        models.ProjectJson     true    "Тело запроса, см. пример"
// @Param   Bearer-token    header  string          true    "Токен доступа, пользователь должен быть не ниже куратора"
// @Success 201 {int} Created
// @Failure 403 body is empty
// @router / [post]
func (c *ProjectController) Post() {
	beego.Trace("Try to POST project")
	if c.CurrentUser.PermissionLevel == 2 || c.CurrentUser.PermissionLevel == 1 {
		var v models.ProjectJson
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if id, err := models.AddProject(&v); err == nil {
				beego.Trace("Project with id", id, "created")
				v.Id = id
				user, err := models.GetUserById(c.CurrentUser.UserId)
				if err != nil {
					beego.Critical(c.Ctx.Input.IP(), "Claims corrupted", err.Error())
					c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
					c.Data["json"] = err.Error()

				} else {
					err := models.AddMasterToProject(user, &v)
					if err != nil {
						beego.Critical(c.Ctx.Input.IP(), "Can't add creator to project", err.Error())
						c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
						c.Data["json"] = err.Error()

					} else {
						beego.Trace("OK")
						c.Ctx.Output.SetStatus(HTTP_CREATED)
						c.Data["json"] = id
					}
				}
			} else {
				beego.Debug("Post project `AddProject` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)

			}
		} else {
			beego.Debug("Post project `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)

		}
	} else {
		beego.Debug("Access denied for `Post`")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR

	}
	c.ServeJSON()
}


type usersGetter func(int) ([]*models.User, error)

// Вызывает функцию с указанным Id и отсылает в канал полученных пользователей в сокращенном виде
// Используется для параллельного запроса к Masters, Enrolled и Users для проекта
// Функция должна соответствовать usersGetter прототипу
func CallForPartUsers(f usersGetter, id int, c chan []models.MainUserInfo) {
	users, err := f(id)
	if err != nil {
		c <- nil
	} else {
		var partUsers []models.MainUserInfo
		for _, u := range users {
			partUsers = append(partUsers, models.MainUserInfo{
				Id: u.Id,
				Avatar: u.Avatar,
				Nickname: u.Nickname,
			})
		}
		c <- partUsers
	}
}

// GetOne ...
// @Title Get One
// @Description Получить подробную информацию о проекте
// @Param   id  path    string  true        "ID проекта, информацию о котором нужно получить"
// @Param   cut query   bool    false       "Оставить только информацию о проекте?"
// @Success 200 {object} models.AllInformationAboutProject     Запрос прошел успешно
// @Failure 400 :id is wrong
// @router /:id [get]
func (c *ProjectController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	beego.Trace(c.Ctx.Input.IP(), "Get project with id", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Debug("GetOne `Atoi` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error() // TODO: change to "Wrong project id"
	} else {
		v, err := models.GetProjectById(id)
		if err != nil {
			beego.Debug("GetOne `GetProjectById` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		} else {
			// Запускаем 3 параллельных запроса к мастерам, участникам и заявкам на проект
			// TODO: Исследовать на утечки памяти
			enrolledChan := make(chan []models.MainUserInfo)
			membersChan := make(chan []models.MainUserInfo)
			mastersChan := make(chan []models.MainUserInfo)
			if cut_info, _ := c.GetBool("cut"); !cut_info {
				go CallForPartUsers(models.GetAllSignedUpOnProject, id, enrolledChan)
				go CallForPartUsers(models.GetMastersOfTheProject, id, mastersChan)
				go CallForPartUsers(models.GetUsersByProjectId, id, membersChan)
			} else {
				go func() {
					enrolledChan <- nil
					mastersChan <- nil
					membersChan <- nil
				}()
			}
			beego.Trace("GetOne OK")
			c.Data["json"] = models.AllInformationAboutProject{
				Project: v,
				Enrolled: <-enrolledChan,
				Members: <-membersChan,
				Masters: <-mastersChan,
			}
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description Получить фильтрованный список проектов
// @Param   sortby  query   string  false   "Sorted-by fields. e.g. col1,col2 ..."
// @Param   order   query   string  false   "Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param   user    query   string  false   "Получить проекты, в которых участвует пользователь с заданным ID."
// @Param   master  query   string  false   "Получить проекты, автором которых является пользователь с заданным ID."
// @Param   tag     query   string  false   "Получить проекты с тегом. Тег может быть только один."
// @Param   status  query   string  false   "Получить проекты с заданным статусом ('завершен'/'еще не начат')"
// @Param   limit   query   string  false   "Наибольшее число объектов в ответе. Должно быть целым числом. Изначально равно 10"
// @Param   offset  query   string  false   "Смещение от начала. Должно быть целым числом"
// @Success 200 {object} []models.Project Get array of projects filtered with specified filters (wtf this description)
// @Failure 403
// @router / [get]
func (c *ProjectController) GetAll() {
	var sortBy []string
	var order []string
	var tag string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	var user   int64
	var master int64
	var status string


	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v, err := c.GetInt64("user"); err == nil {
		user = v
	}
	if v, err := c.GetInt64("master"); err == nil {
		master = v
	}
	// sortBy: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortBy = strings.Split(v, ",")
	}else{
		sortBy = []string{"Name"}
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}else{
		order  = []string{"asc"}
	}
	if v := c.GetString("tag"); v!= ""{
		tag = v
	}
	if v := c.GetString("status"); v != "" {
		if correctStatus(v) {
			status = v
		}
	}
	beego.Trace(status)
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
	beego.Trace("Select from table")
	l, err := models.GetAllProjects(query, sortBy, order, offset, limit, tag, user, master, status)
	if err != nil {
		beego.Debug(c.Ctx.Input.IP(), "News GetAll `GetAllProjects` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}


func correctStatus(status string) bool{
	if status == "0" || status == "1" || status == "2" {
		return true
	}
	return false
}

// Put ...
// @Title Put
// @Description Изменить/обновить проект
// @Param   id              path    string          true    "ID проекта, который нужно обновить"
// @Param   body            body    models.Project  true    "Тело запроса, см. пример"
// @Param   Bearer-token    header  string          true    "Токен доступа администратора или создателя проекта"
// @Success 200 "OK"
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProjectController) Put() {
	// TODO: добавить автора к проекту
	if c.CurrentUser.PermissionLevel == 2 || c.CurrentUser.PermissionLevel == 1 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Put `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		}
		v := models.ProjectJson{Id: id}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if err := models.UpdateProjectById(&v); err == nil {
				beego.Trace("Put project OK")
				c.Data["json"] = HTTP_OK_STR
			} else {
				beego.Debug("Put news `UpdateProjectById` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			}
		} else {
			beego.Debug("Put project `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		}
	} else {
		beego.Debug("Access denied for `Put`")
		c.Data["json"] = HTTP_FORBIDDEN_STR
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
	}
	c.ServeJSON()
}

// TODO: удалится ли всё, что связано с проектом, если его удалить?
// Delete ...
// @Title Delete
// @Description delete the Project
// @Param   id              path    string      true        "ID проекта, который нужно удалить"
// @Param   Bearer-token    header  string      true        "Токен доступа администратора или автора проекта"
// @Success 200 "OK"
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProjectController) Delete() {
	// TODO: добавить проверку на автора
	if c.CurrentUser.PermissionLevel == 2 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Delete 'Atoi' error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		}
		if err := models.DeleteProject(id); err == nil {
			beego.Trace("Delete OK")
			c.Data["json"] = "OK"
		} else {
			beego.Critical(c.Ctx.Input.IP(), "'DeleteProject' error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		}
	} else {
		beego.Debug("Access denied for `Delete`")
		c.Data["json"] = "Access denied for `Delete`" // TODO: change this
		c.Ctx.Output.SetStatus(403)
	}
	c.ServeJSON()
}
