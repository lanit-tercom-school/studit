package main

import (
	_ "service/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
<<<<<<< HEAD:src/service/main.go
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
=======
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@127.0.0.1:5432/studit?sslmode=disable")
>>>>>>> origin/feature/project:src/studitapi/main.go
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

