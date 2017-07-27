package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var GetFileById gql.Field
var GetFileList gql.Field

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
	GetFileList = gql.Field{
		Type: objects.FileListType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolveGetFileList,
	}
}
