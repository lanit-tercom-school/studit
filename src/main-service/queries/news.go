package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var GetNewsById gql.Field
var GetNewsList gql.Field

func init() {
	GetNewsById = gql.Field{
		Type: objects.NewsType,
		Args: gql.FieldConfigArgument{
			"Id": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolveGetNews,
	}
	GetNewsList = gql.Field{
		Type: gql.NewList(objects.NewsType),
		Args: gql.FieldConfigArgument{
			"Limit": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Offset": &gql.ArgumentConfig{
				Type: gql.String,
			},
		},
		Resolve: objects.ResolveGetNewsList,
	}
}
