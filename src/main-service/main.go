package main

import (
	"main-service/root"
	"net/http"

	"github.com/graphql-go/handler"

	"log"

	gql "github.com/graphql-go/graphql"
)

var Schema gql.Schema

func init() {
	log.Println("Initialization schema")
	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query: root.RootQuery,
	})
	if err != nil {
		log.Panicf("Error: %s", err)
	}
	Schema = schema
}

func main() {
	log.Print("Server started")
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
