package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostNews gql.Field
var DeleteNews gql.Field

func init() {
	PostNews = gql.Field{
		Type: objects.NewsType,
		Args: gql.FieldConfigArgument{
			"Title": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"Description": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"Image": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolvePostNews,
	}

	DeleteNews = gql.Field{
		Type: objects.MessageType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.Int),
			},
		},
		Resolve: objects.ResolveDeleteNews,
	}
}
