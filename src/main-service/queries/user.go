package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var GetUserById gql.Field

func init() {
	GetUserById = gql.Field{
		Type: objects.UserType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.ID,
			},
		},
		Resolve: objects.ResolveGetUser,
	}
}
