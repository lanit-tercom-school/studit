package main

import (
	"data-service/auth"
	_ "data-service/routers"
	"net"
	"net/http"
	"net/rpc"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	dbconf, err := config.NewConfig("ini", "conf/database.conf")
	if err != nil {
		beego.Critical(err.Error())
		panic(err)
	}
	postgresStrConfig := "postgres://" + dbconf.String("login") + ":" +
		dbconf.String("password") + "@" + dbconf.String("host") + ":" + dbconf.String("port") + "/" +
		dbconf.String("database") + "?sslmode=" + dbconf.String("sslmode")
	orm.RegisterDataBase("default", "postgres", postgresStrConfig)
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
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
