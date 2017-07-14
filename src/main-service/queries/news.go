package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var GetNewsById gql.Field

func init() {
	GetNewsById = gql.Field{
		Type: objects.NewsType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.ID,
			},
		},
		Resolve: objects.ResolveGetNews,
	}
}
