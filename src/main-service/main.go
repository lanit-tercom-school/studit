package main

import (
	"main-service/helpers"
	"main-service/root"
	"net/http"

	"main-service/handler"

	"log"

	"main-service/conf"

	gql "github.com/graphql-go/graphql"
)

var Schema gql.Schema

func init() {
	helpers.LogServer("Schema initialization")
	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query: root.RootQuery,
	})
	if err != nil {
		helpers.LogErrorServer(err)
	}
	helpers.LogServer("Schema initialization successfully")
	Schema = schema
}

func main() {
	helpers.LogServer("Start on " + conf.Configuration.HttpPort)
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})
	http.Handle("/graphql", h)
	err := http.ListenAndServe(":"+conf.Configuration.HttpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
