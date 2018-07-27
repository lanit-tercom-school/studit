package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var UserQuery gql.Field

type UserData struct {
	Nickname objects.Message
}

var UserQueryType = gql.NewObject(
	gql.ObjectConfig{
		Name: "UserQuery",
		Fields: gql.Fields{
			"ChangeNickname": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutNewNickname,
			},
			"ChangeAvatar": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutNewAvatar,
			},
			"ChangeDescription": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutNewDescription,
			},
			"ChangeContact": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"TypeId": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutNewContact,
			},
		},
	},
)

func init() {
	UserQuery = gql.Field{
		Type:    UserQueryType,
		Resolve: func(p gql.ResolveParams) (interface{}, error) { return UserData{}, nil },
	}
}
