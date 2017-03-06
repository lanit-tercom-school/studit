package controllers

import (
	"encoding/json"
	"service/models"

	"github.com/astaxie/beego"
	"service/auth"
)

// AuthController operations for ContactType
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ContactType
// @Param	body		body 	auth.Usr	true		"body for Usr content"
// @Success 201 {int} models.ContactType
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Post() {
	var v auth.Usr
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := auth.TryToLogin(v.Login, v.Password)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			// TODO: handle session
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ContactType by id
// @Success 200 {object} models.ContactType
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AuthController) GetOne() {
	var response struct{
		Error string `json:"error"`
	}
	response.Error = "Method Not Allowed"
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ContactType
// @Success 200 {object} models.ContactType
// @Failure 403
// @router / [get]
func (c *AuthController) GetAll() {
	var response struct{
		Error string `json:"error"`
	}
	response.Error = "Method Not Allowed"
	c.Data["json"] = response
	c.Ctx.ResponseWriter.WriteHeader(405)
	c.ServeJSON()
}
/*
// Put ...
// @Title Put
// @Description update the ContactType
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ContactType	true		"body for ContactType content"
// @Success 200 {object} models.ContactType
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AuthController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ContactType{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateContactTypeById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the ContactType
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AuthController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteContactType(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}*/
