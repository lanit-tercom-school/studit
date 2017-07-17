package main

import (
	"main-service/root"
	"net/http"

	"main-service/handler"

	"log"

	gql "github.com/graphql-go/graphql"
)

var Schema gql.Schema

func init() {
	log.Println("Server: Schema initialization")
	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query: root.RootQuery,
	})
	if err != nil {
		log.Panicf("Schema: Error: %s", err)
	}
	log.Println("Server: Schema initialization successfully")
	Schema = schema
}

func main() {
	log.Print("Server: Started")
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
