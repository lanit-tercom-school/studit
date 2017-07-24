package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostProjectEnroll gql.Field
var DeleteProjectEnroll gql.Field

func init() {
	PostProjectEnroll = gql.Field{
		Type: objects.ProjectEnrollType,
		Args: gql.FieldConfigArgument{
			"User": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
			"Project": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
			"Message": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolvePostProjectEnroll,
	}
	DeleteProjectEnroll = gql.Field{
		Type: objects.MessageType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
		},
		Resolve: objects.ResolveDeleteProjectEnroll,
	}
}
