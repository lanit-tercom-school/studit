package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostProject gql.Field
var PostProjectEnroll gql.Field
var DeleteProjectEnroll gql.Field

func init() {
	PostProject = gql.Field{
		Type: objects.ProjectType,
		Args: gql.FieldConfigArgument{
			"Name": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"Description": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"Logo": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Tags": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Status": &gql.ArgumentConfig{
				Type:         gql.String,
				DefaultValue: "0",
			},
		},
		Resolve: objects.ResolvePostProject,
	}

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
