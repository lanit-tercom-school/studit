package main

import (
	_ "file-service/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	dbconf, err := config.NewConfig("ini", "conf/database.conf")
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
	// "postgres://login:password@host:port/database?sslmode=disable"
	postgresStrConfig := "postgres://" + dbconf.String("login") + ":" +
		dbconf.String("password") + "@" + dbconf.String("host") + ":" + dbconf.String("port") + "/" +
		dbconf.String("database") + "?sslmode=" + dbconf.String("sslmode")
	err = orm.RegisterDataBase("default", "postgres", postgresStrConfig)
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/files"] = "files"
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
