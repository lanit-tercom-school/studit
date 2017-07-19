package root

import (
	"main-service/mutations"

	gql "github.com/graphql-go/graphql"
)

var RootMutation = gql.NewObject(gql.ObjectConfig{
	Name: "RootMutation",
	Fields: gql.Fields{
		"PostNews":    &mutations.PostNews,
		"PostProject": &mutations.PostProject,
		"Auth":        &mutations.AuthQuery,
	},
})
