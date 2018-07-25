package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostProject gql.Field
var DeleteProject gql.Field

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
			"GitHubUrl": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Status": &gql.ArgumentConfig{
				Type:         gql.String,
				DefaultValue: "opened",
			},
		},
		Resolve: objects.ResolvePostProject,
	}

	DeleteProject = gql.Field{
		Type: objects.MessageType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
		},
		Resolve: objects.ResolveDeleteProject,
	}
}
