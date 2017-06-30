package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"os"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
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
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "Bearer-token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "prod" {
		beego.SetLevel(beego.LevelError)
	}
	os.Mkdir("logs", 0777)
	beego.SetLogger("file", `{"filename":"logs/main.log"}`/*"\"logs/\ + time.Now().Format(\"2006-01-02 15_04") + ".log""*/)
	beego.Info("")
	beego.Run()
}
