package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "service/routers"
	"github.com/vetcher/jwt"
	"time"
	"github.com/astaxie/beego/plugins/cors"
	"os"
)

func init() {
	jwt.GlobalStorage = jwt.NewStorage(time.Hour)
	postgresStrConfig := "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable"
	err := orm.RegisterDataBase("default", "postgres", postgresStrConfig)
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.SetStaticPath("/", "static")
		beego.SetStaticPath("/swagger", "swagger")
		beego.SetStaticPath("/logo", "logo")
		beego.SetStaticPath("/assets", "assets")
	}
	os.Mkdir("logs", 0777)
	beego.SetLogger("file", `{"filename":"logs/test.log"}`/*"\"logs/\ + time.Now().Format(\"2006-01-02 15_04") + ".log""*/)
	beego.Info("")
	beego.Run()
}
