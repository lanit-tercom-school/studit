package controllers

import (
	"io"
	"main-service/models"
	"os"

	"time"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

// FileController --- контроллер для загрузка файлов
type FileController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *FileController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description Загрузить файл
// @Param   uploadfile      formData    file          true    "Файл"
// @Param   Bearer-token    header      string          true    "Токен доступа пользователя"
// @Success 201
// @Failure 403
// @router / [post]
func (c *FileController) Post() {
	if c.CurrentUser.PermissionLevel != models.VIEWER {
		beego.Trace("Uploading file ...")
		file, handler, err := c.Ctx.Request.FormFile("uploadfile")
		if err != nil {
			beego.Trace(err)
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			return
		}
		defer file.Close()
		filename := uuid.New().String()
		u, err := models.GetUserById(c.CurrentUser.UserId)
		if err != nil {
			beego.Trace(err)
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			return
		}
		t := models.File{
			User:           u,
			Name:           handler.Filename,
			Path:           "files/" + filename,
			DateOfCreation: time.Now(),
		}
		id, err := models.AddFile(&t)
		if err != nil {
			beego.Trace(err)
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			return
		}
		beego.Trace("Creating file with name " + filename)
		f, err := os.Create("files/" + filename)
		if err != nil {
			beego.Trace(err)
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		c.Data["json"] = id
		c.Ctx.Output.SetStatus(HTTP_CREATED)
	} else {
		beego.Trace("Can not upload file. Access is denied")
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description Получить информацию о файле по id
// @Param   id  path    string  true    "ID файла, о котором нужно узнать информацию"
// @Param   Bearer-token    header      string          true    "Токен"
// @Success 200
// @Failure 403
// @router /:id [get]
func (c *FileController) GetOne() {
	if c.CurrentUser.PermissionLevel != models.VIEWER {
		idStr, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			beego.Trace(err)
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		} else {
			f, err := models.GetFileById(idStr)
			if err == nil {
				if f.User.Id == c.CurrentUser.UserId || c.CurrentUser.PermissionLevel == models.ADMIN {
					c.Data["json"] = f
					c.Ctx.Output.SetStatus(HTTP_OK)
				} else {
					beego.Trace("Access is denied")
					c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				}
			} else {
				beego.Trace(err)
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			}
		}
	} else {
		beego.Trace("Access is denied")
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description Получить информацию о файлах пользователя
// @Param   Bearer-token    header      string          true    "Токен"
// @Param   Id query   int                              false   "Id пользователя (для администратора)"
// @Success 200
// @Failure 403
// @router / [get]
func (c *FileController) GetAll() {
	if c.CurrentUser.PermissionLevel != models.VIEWER {
		id, err := c.GetInt("Id")
		if err != nil {
			f, err := models.GetFilesByUserId(c.CurrentUser.UserId)
			if err == nil {
				c.Data["json"] = f
				c.Ctx.Output.SetStatus(HTTP_OK)
			} else {
				beego.Trace(err)
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			}
		} else {
			if c.CurrentUser.PermissionLevel == models.ADMIN {
				f, err := models.GetFilesByUserId(id)
				if err == nil {
					c.Data["json"] = f
					c.Ctx.Output.SetStatus(HTTP_OK)
				} else {
					beego.Trace(err)
					c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				}
			} else {
				beego.Trace("Access is denied")
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			}
		}
	} else {
		beego.Trace("Access is denied")
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description Удалить файл с сервера
// @Param   id              path    string      true        "ID файла, который нужно удалить"
// @Param   Bearer-token    header  string      true        "Токен доступа администратора или автора файла"
// @Success 200 "OK"
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FileController) Delete() {
	if c.CurrentUser.PermissionLevel != models.VIEWER {
		idStr, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			beego.Trace(err)
			c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
		} else {
			f, err := models.GetFileById(idStr)
			if err == nil {
				if f.User.Id == c.CurrentUser.UserId || c.CurrentUser.PermissionLevel == models.ADMIN {
					_, err := models.DeleteFile(f)
					if err != nil {
						beego.Trace(err)
						c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
					} else {
						os.Remove(f.Path)
						c.Ctx.Output.SetStatus(HTTP_OK)
						c.Data["json"] = "Success"
					}
				} else {
					beego.Trace("Access is denied")
					c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
				}
			} else {
				beego.Trace(err)
				c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
			}
		}
	} else {
		beego.Trace("Access is denied")
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
	}
	c.ServeJSON()
}
