package sql

import "github.com/nleof/goyesql"

//map of sql queries
var QueriesMap goyesql.Queries

func init() {
	QueriesMap = goyesql.MustParseFile("sql/queries.sql")
}
