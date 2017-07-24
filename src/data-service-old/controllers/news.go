package controllers

import (
	"data-service-old/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
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
	beego.Trace("Try to POST news")
	if c.CurrentUser.PermissionLevel == ADMIN {
		var v models.NewsJson
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if id, err := models.AddNews(&v); err == nil {
				beego.Trace("News with id", id, "created")
				c.Ctx.Output.SetStatus(HTTP_CREATED)
				c.Data["json"] = id
			} else {
				beego.Debug("Post news `AddNews` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			}
		} else {
			beego.Debug("Post news `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		}
	} else {
		beego.Debug("Access denied for `Post`")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
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
	beego.Trace("Get news with id", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		beego.Debug("GetOne `Atoi` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	}
	v, err := models.GetNewsById(id)
	if err != nil {
		beego.Debug("GetOne `GetNewsById` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		c.Data["json"] = err.Error()
	} else {
		beego.Trace("GetOne OK")
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
	var tags string

	beego.Trace("Parce request params for News")
	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		if v > 20 {
			limit = 20
		} else {
			limit = v
		}
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt("offset"); err == nil {
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
		tags = v
	}

	beego.Trace("Select from table")
	l, err := models.GetAllNews(sortBy, order, offset, limit, tags)
	if err != nil {
		beego.Debug("News GetAll `GetAllNews` error", err.Error())
		c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
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
	if c.CurrentUser.PermissionLevel == ADMIN {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Put `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		}
		v := models.NewsJson{Id: id}
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
			if err := models.UpdateNewsById(&v); err == nil {
				beego.Trace("Put news OK")
				c.Data["json"] = HTTP_OK_STR
			} else {
				beego.Debug("Put news `UpdateNewsById` error", err.Error())
				c.Data["json"] = err.Error()
				c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			}
		} else {
			beego.Debug("Put news `Unmarshal` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		}
	} else {
		beego.Debug("Access denied for `Put`")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR
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
	if c.CurrentUser.PermissionLevel == ADMIN {
		idStr := c.Ctx.Input.Param(":id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			beego.Debug("Delete `Atoi` error", err.Error())
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
			c.Data["json"] = err.Error()
		}
		if err := models.DeleteNews(id); err == nil {
			beego.Trace("Delete news OK")
			c.Data["json"] = HTTP_OK_STR
		} else {
			beego.Debug("Delete news `DeleteNews` error", err.Error())
			c.Data["json"] = err.Error()
			c.Ctx.Output.SetStatus(HTTP_BAD_REQUEST)
		}
	} else {
		beego.Debug("Access denied for `Delete`")
		c.Ctx.Output.SetStatus(HTTP_FORBIDDEN)
		c.Data["json"] = HTTP_FORBIDDEN_STR
	}
	c.ServeJSON()
}
