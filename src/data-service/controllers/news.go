package controllers

import (
	"data-service/models"
	"encoding/json"
	//"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// NewsController operations for News
type NewsController struct {
	beego.Controller
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
// @Param	body		body 	models.News	true		"body for News content"
// @Success 201 {int} models.News
// @Failure 403 body is empty
// @router / [post]
func (c *NewsController) Post() {
	var v models.News
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddNews(&v); err == nil {
			c.Ctx.Output.SetStatus(HTTP_CREATED)
			c.Data["json"] = v
		} else {
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
	} else {
		c.Data["json"] = MakeMessageForSending(err.Error())
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get News by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.News
// @Failure 403 :id is empty
// @router /:id [get]
func (c *NewsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNewsById(id)
	if err != nil {
		c.Data["json"] = MakeMessageForSending(err.Error())
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get News
// @Param  sortCols       query  string  false  "Sorted-by columns. e.g. col1,col2 ..."
// @Param  orders         query  string  false  "Orders corresponding to each sorting columns in sortCols, if single value, apply to all sortCols. e.g. desc,asc ..."
// @Param  limit          query  string  false  "Limit the size of result set. Must be an integer"
// @Param  offset         query  string  false  "Start position of result set. Must be an integer"
// @Param  tags           query  string  false  "Tags, e.g. "World" or "World,Other"
// @Param  tagsOperation  query  string  false  "Tags operation, only &quot;and&quot; or &quot;or&quot;, default &quot;and&quot;
// @Success 200 {object} models.News
// @Failure 403
// @router / [get]
func (c *NewsController) GetAll() {
	var sortCols, orders []string
	var offset, limit int = 0, 10
	var tags string
	var tagsOperation string = "and"
	beego.Trace("Parce request params for News")

	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt("offset"); err == nil {
		offset = v
	}
	// sortCols: col1,col2
	if v := c.GetString("sortCols"); v != "" {
		sortCols = strings.Split(v, ",")
	}
	// orders: desc,asc
	if v := c.GetString("orders"); v != "" {
		orders = strings.Split(v, ",")
	}
	// tags: World,Other
	tags = c.GetString("tags")
	// tagsOperation: and,or
	if v := c.GetString("tagsOperation"); v != "" {
		tagsOperation = v
	}

	l, err := models.GetAllNews(sortCols, orders, offset, limit, tags, tagsOperation)
	if err != nil {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	} else {
		c.Ctx.Output.SetStatus(HTTP_OK)
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the News
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.News	true		"body for News content"
// @Success 200 {object} models.News
// @Failure 403 :id is not int
// @router /:id [put]
func (c *NewsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.News{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateNewsById(&v); err == nil {
			c.Ctx.Output.SetStatus(HTTP_OK)
			c.Data["json"] = MakeMessageForSending(HTTP_OK_STR)
		} else {
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			c.Data["json"] = MakeMessageForSending(err.Error())
		}
	} else {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the News
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NewsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteNews(id); err == nil {
		c.Ctx.Output.SetStatus(HTTP_OK)
		c.Data["json"] = MakeMessageForSending(HTTP_OK_STR)
	} else {
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		c.Data["json"] = MakeMessageForSending(err.Error())
	}
	c.ServeJSON()
}
