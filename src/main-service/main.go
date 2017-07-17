package main

import (
	"main-service/root"
	"net/http"

	"github.com/graphql-go/handler"

	"log"

	gql "github.com/graphql-go/graphql"
)

var schema, _ = gql.NewSchema(gql.SchemaConfig{
	Query: root.RootQuery,
})

func main() {
	log.Print("Server started")
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", h)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
