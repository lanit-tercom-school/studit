package main

import (
	"main-service/helpers"
	"main-service/root"
	"net/http"

	"main-service/handler"

	"log"

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
	helpers.LogServer("Started")
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})
	http.Handle("/graphql", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
