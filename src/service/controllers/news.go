package controllers

import (
	"encoding/json"
	"service/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

// NewsController operations for News
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
// @Description create News
// @Param	body		body 	models.NewsJson	true		"body for News content"
// @Param	token		query	string		true		"Access token"
// @Success 201 {int} models.NewsJson
// @Failure 403 body is empty
// @router / [post]
func (c *NewsController) Post() {
	beego.Trace(c.Ctx.Input.IP(), "Try to POST news")
	if c.Ctx.Output.IsOk() {
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
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get News by id
// @Param   id      path    string  true    "The key for static block"
// @Success 200 {object} models.NewsJson
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
// @Description Get bunch of news
// @Param   sort_by     query   string  false   "Sorted-by fields. e.g. title, description, time"
// @Param   order       query   string  false   "Order corresponding to each sort_by field, if single value, apply to all sort_by fields. e.g. desc,asc ..., can be only `desc` or `asc`, default is asc"
// @Param   tag         query   string  false   "Filter by one and only one tag. e.g. Other"
// @Param   limit       query   string  false   "Limit the size of result set. Must be an integer"
// @Param   offset      query   string  false   "Start position of result set. Must be an integer"
// @Success 200 {object} models.NewsJson
// @Failure 403
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
	// tag: Other
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
// @Description Update(edit) the News with id
// @Param   id      path    string              true        "The id you want to update"
// @Param   body    body    models.NewsJson     true        "Body for News content, id, created and edited fields ignores"
// @Param   token   query   string              true        "Access token"
// @Success 200 "OK"
// @Failure 403 :id is not int
// @router /:id [put]
func (c *NewsController) Put() {
	if c.Ctx.Output.IsOk() {
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
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Delete the News
// @Param	id		path 	string	true		"The id you want to delete"
// @Param	token	query	string	true		"Access token"
// @Success 200 {string} Delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NewsController) Delete() {
	if c.Ctx.Output.IsOk() {
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
		c.Ctx.Output.SetStatus(400)
	}
	c.ServeJSON()
}
