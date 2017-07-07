package sql

import (
	"errors"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/nleof/goyesql"
)

//map of sql queries
var QueriesMap goyesql.Queries

func LoadSql(path string) (err error) {
	QueriesMap = make(goyesql.Queries)
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		Queries, err := goyesql.ParseFile(path + f.Name())
		if err != nil {
			beego.Trace("Error parsing file " + f.Name())
			return err
		}
		for k, v := range Queries {
			if _, ok := QueriesMap[k]; ok {
				return errors.New("Error parsing SQLs. Sql name, " + string(k) + " already exists!")
			}
			QueriesMap[k] = v
		}
	}
	return
}
