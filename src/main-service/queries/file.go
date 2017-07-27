package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var GetFileById gql.Field

func init() {
	GetFileById = gql.Field{
		Type: objects.FileType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolveGetFileById,
	}
}
