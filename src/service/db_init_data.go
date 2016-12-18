package main

import (
	_ "github.com/lanit-tercom-school/studit/service/routers"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5432/studit?sslmode=disable")
}

func main() {
	// Here add some scripts
}

