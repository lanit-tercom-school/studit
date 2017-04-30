package controllers

import (
	"encoding/json"
	"service/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

// Создание, изменение, удаление и просмотр новостей
type NewsController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *NewsController) URLMapping() {
    c.Mapping("Post", c.Post)
    c.Mapping("GetOne", c.GetOne)
    c.Mapping("GetAll", c.GetAll)
    c.Mapping("Put", c.Put)
    c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description Создать новую новость
// @Param   body                body        models.NewsJson true    "Тело запроса, см. пример, поля `id`, `created`, `edited` игнорируются"
// @Param   Bearer-token        header      string          true    "Токен доступа администратора"
// @Success 201 {int} ID созданой новости
// @Failure 400 body is empty
// @Failure 400 Forbidden
// @router / [post]
func (c *NewsController) Post() {
	beego.Trace(c.Ctx.Input.IP(), "Try to POST news")
	if c.CurrentUser.PermissionLevel == 2 {
		var v models.NewsJson
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if id, err := models.AddNews(&v); err == nil {
				beego.Trace(c.Ctx.Input.IP(), "News with id", id, "created")
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = id
			} else {
				beego.Debug(c.Ctx.Input.IP(), "Post news `AddNews` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(500)
			}
		} else {
			beego.Debug(c.Ctx.Input.IP(), "Post news `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)
		}
	} else {
        beego.Debug(c.Ctx.Input.IP(), "Access denied for `Post`")
		c.Ctx.Output.SetStatus(403)
        c.Data["json"] = "Forbidden"
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description Получить подробную новость
// @Param   id      path    string  true    "ID новости"
// @Success 200 {object} models.NewsJson    Успешный запрос
// @Failure 403 :id is empty
// @router /:id [get]
func (c *NewsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	beego.Trace(c.Ctx.Input.IP(), "Get news with id", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Debug(c.Ctx.Input.IP(), "GetOne `Atoi` error", err.Error())
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	}
	v, err := models.GetNewsById(id)
	if err != nil {
		beego.Debug(c.Ctx.Input.IP(), "GetOne `GetNewsById` error", err.Error())
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	} else {
		beego.Trace(c.Ctx.Input.IP(), "GetOne OK")
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description Получить список новостей
// @Param   sort_by     query   string  false   "Поле, по которому сортировать новости, напр. title, description, time"
// @Param   order       query   string  false   "Order corresponding to each sort_by field, if single value, apply to all sort_by fields. e.g. desc,asc ..., can be only `desc` or `asc`, default is asc"
// @Param   tag         query   string  false   "Получить новость с тегом. Тег может быть только один."
// @Param   limit       query   string  false   "Максимальное количество новостей. Должно быть числом. Не более 20"
// @Param   offset      query   string  false   "Отступ от начала. Должно быть числом."
// @Success 200 {object} []models.NewsJson  Успешный запрос
// @Failure 400 Error
// @router / [get]
func (c *NewsController) GetAll() {
	var sortBy []string
	var order []string
	var limit int64 = 10
	var offset int64
	var tag string
	beego.Trace(c.Ctx.Input.IP(), "Parce request params for News")
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		if limit > 20 {
			limit = 20
		} else {
			limit = v
		}
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortBy: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortBy = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// tags: Other
	if v := c.GetString("tag"); v != "" {
		tag = v
	}

	beego.Trace(c.Ctx.Input.IP(), "Select from table")
	l, err := models.GetAllNews(sortBy, order, offset, limit, tag)
	if err != nil {
		beego.Debug(c.Ctx.Input.IP(), "News GetAll `GetAllNews` error", err.Error())
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description Изменить новость
// @Param   id              path        string              true    "ID новости, которую нужно изменить"
// @Param   body            body        models.NewsJson     true    "Тело запроса, см. пример, поля id, created and edited fields ignores"
// @Param   Bearer-token    header      string              true    "Токен доступа администратора"
// @Success 200 OK
// @Failure 403 :id is not int
// @Failure 403 Forbidden
// @router /:id [put]
func (c *NewsController) Put() {
	if c.CurrentUser.PermissionLevel == 2 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug(c.Ctx.Input.IP(), "Put `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
		v := models.NewsJson{Id: id}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if err := models.UpdateNewsById(&v); err == nil {
				beego.Trace(c.Ctx.Input.IP(), "Put news OK")
				c.Data["json"] = "OK"
			} else {
				beego.Debug(c.Ctx.Input.IP(), "Put news `UpdateNewsById` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(400)
			}
		} else {
			beego.Debug(c.Ctx.Input.IP(), "Put news `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)
		}
	} else {
        beego.Debug(c.Ctx.Input.IP(), "Access denied for `Put`")
		c.Ctx.Output.SetStatus(400)
        c.Data["json"] = "Forbidden"
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Удалить новость
// @Param   id              path    string  true    "ID новости, которую нужно удалить"
// @Param   Bearer-token    header  string  true    "Токен доступа администратора"
// @Success 200 OK
// @Failure 403 id is empty
// @Failure 403 Forbidden
// @router /:id [delete]
func (c *NewsController) Delete() {
	if c.CurrentUser.PermissionLevel == 2 {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug(c.Ctx.Input.IP(), "Delete `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = err.Error()
		}
		if err := models.DeleteNews(id); err == nil {
			beego.Trace(c.Ctx.Input.IP(), "Delete news OK")
			c.Data["json"] = "OK"
		} else {
			beego.Debug(c.Ctx.Input.IP(), "Delete news `DeleteNews` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(400)
		}
	} else {
        beego.Debug(c.Ctx.Input.IP(), "Access denied for `Delete`")
		c.Ctx.Output.SetStatus(400)
        c.Data["json"] = "Forbidden"
	}
	c.ServeJSON()
}
