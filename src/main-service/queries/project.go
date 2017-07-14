package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var GetProjectById gql.Field

func init() {
	GetProjectById = gql.Field{
		Type: objects.ProjectType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.ID,
			},
		},
		Resolve: objects.ResolveGetProject,
	}
}
