package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "service/routers"
)

func init() {
	// TODO: delete this or refactor
	/*mngr_config := session.ManagerConfig{
		CookieLifeTime: 3600,
		CookieName: "omfg_this_is_cookie",
		EnableSetCookie: false,
		EnableSidInHttpHeader: true,
		EnableSidInUrlQuery: false,
		Gclifetime: 3600,
		Maxlifetime: 3600,
		SessionNameInHttpHeader: "Session",
		Secure: false,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", &mngr_config)*/
	/*if err != nil {
		beego.Critical(err.Error())
	}*/
	err := orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
	if err != nil {
		beego.Critical(err.Error())
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
