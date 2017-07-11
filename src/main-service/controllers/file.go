package controllers

import (
	"io"
	"main-service/models"
	"os"

	"time"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
)

// UploadController --- контроллер для загрузка файлов
type FileController struct {
	ControllerWithAuthorization
}

// URLMapping ...
func (c *FileController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description Загрузить файл в свагере не работает :(
// @Param   uploadfile            form	               true    "Файл"
// @Param   Bearer-token    header      string          true    "Токен доступа пользователя"
// @Success 201
// @Failure 403
// @router / [post]
func (c *FileController) Post() {
	//if c.CurrentUser.PermissionLevel != models.VIEWER {
	beego.Trace("Uploading file ...")
	file, handler, err := c.Ctx.Request.FormFile("uploadfile")
	if err != nil {
		beego.Trace(err)
		return
	}
	defer file.Close()
	filename := uuid.New().String()
	u, err := models.GetUserById(1)
	if err != nil {
		beego.Trace(err)
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
		return
	}
	beego.Trace("Creating file with name " + filename)
	f, err := os.Create("files/" + filename)
	if err != nil {
		beego.Trace(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	c.Data["json"] = id
	c.Ctx.Output.SetStatus(HTTP_CREATED)
	/*} else {
		beego.Trace("Can not upload file. You are only Viewer")
		c.Ctx.Output.SetStatus(HTTP_INTERNAL_SERVER_ERROR)
	}*/
	c.ServeJSON()
}
