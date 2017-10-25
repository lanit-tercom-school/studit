package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostFile gql.Field

func init() {
	PostFile = gql.Field{
		Type: objects.FileType,
		Resolve: objects.ResolvePostFile,
	}
}