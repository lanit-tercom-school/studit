package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var GetProjectById gql.Field
var GetProjectList gql.Field

func init() {
	GetProjectById = gql.Field{
		Type: objects.ProjectType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.ID,
			},
		},
		Resolve: objects.ResolveGetProjectById,
	}
	GetProjectList = gql.Field{
		Type: objects.ProjectsSetType,
		Args: gql.FieldConfigArgument{
			"Limit": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Offset": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolveGetProjectList,
	}

}
