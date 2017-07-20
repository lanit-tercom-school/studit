package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostProjectOn gql.Field
var DeleteProjectOn gql.Field

func init() {
	PostProjectOn = gql.Field{
		Type: objects.ProjectOnType,
		Args: gql.FieldConfigArgument{
			"User": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
			"Project": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
		},
		Resolve: objects.ResolvePostProjectOn,
	}
	DeleteProjectOn = gql.Field{
		Type: objects.MessageType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
		},
		Resolve: objects.ResolveDeleteProjectOn,
	}
}
