package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var Auth gql.Field

func init() {
	Auth = gql.Field{
		Type: objects.AuthDataType,
		Args: gql.FieldConfigArgument{
			"Login": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Password": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolveGetAuthDataByLoginAndPassword,
	}
}
