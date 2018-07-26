package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var PostProjectNews gql.Field

func init() {
	PostProjectNews = gql.Field{
		Type: objects.MessageType,
		Args: gql.FieldConfigArgument{
			"ProjectID": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.ID),
			},
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
		Resolve: objects.ResolvePostProjectNews,
	}
}
