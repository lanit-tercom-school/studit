package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostProject gql.Field

func init() {
	PostProject = gql.Field{
		Type: objects.ProjectType,
		Args: gql.FieldConfigArgument{
			"Name": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Description": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Logo": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Tags": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Status": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolvePostProject,
	}
}
