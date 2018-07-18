package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

type NewsData struct {
	message objects.Message
}

var PostNews gql.Field

var EditNews gql.Field

var EditNewsQueryType = gql.NewObject(
	gql.ObjectConfig{
		Name: "NewsQuery",
		Fields: gql.Fields{
			"ChangeTitle": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Id": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutNewsTitle,
			},
			"ChangeDescription": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Id": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutNewsDescription,
			},
			"ChangeImage": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Id": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutNewsImage,
			},
		},
	},
)

//You can check this queries in Altair
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
	EditNews = gql.Field{
		Type:    EditNewsQueryType,
		Resolve: func(p gql.ResolveParams) (interface{}, error) { return NewsData{}, nil },
	}
}
