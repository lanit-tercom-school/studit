package main

import (
	"main-service/auth"
	_ "main-service/routers"
	"main-service/sql"
	"net"
	"net/http"
	"net/rpc"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
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
	err = StartRPCService()
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
}

func StartRPCService() error {
	useractivator := new(auth.UserActivationService)
	err := rpc.Register(useractivator)
	if err != nil {
		return err
	} else {
		rpc.HandleHTTP()
		l, err := net.Listen("tcp", ":8088")
		if err != nil {
			return err
		} else {
			// Start RPC service
			go http.Serve(l, nil)
			return nil
		}
	}
}

func main() {
	sql.LoadSql("sql/queries/")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "Bearer-token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.SetStaticPath("/files", "files")
		beego.SetStaticPath("/swagger", "swagger")
	} else if beego.BConfig.RunMode == "prod" {
		beego.SetLevel(beego.LevelError)
	}
	os.Mkdir("logs", 0777)
	beego.SetLogger("file", `{"filename":"logs/main.log"}` /*"\"logs/\ + time.Now().Format(\"2006-01-02 15_04") + ".log""*/)
	beego.Info("")
	beego.Run()
}
