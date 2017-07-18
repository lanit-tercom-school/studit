package queries

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

var AuthQuery gql.Field

type AuthData struct {
	Signin     objects.SigninDataToGet
	Signup     objects.SignupDataToGet
	Activation objects.ActivationDataToGet
}

var AuthQueryType = gql.NewObject(
	gql.ObjectConfig{
		Name: "AuthQuery",
		Fields: gql.Fields{
			"Signin": &gql.Field{
				Type: objects.SigninDataType,
				Args: gql.FieldConfigArgument{
					"Login": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Password": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolveGetSigninDataByLoginAndPassword,
			},
			"Signup": &gql.Field{
				Type: objects.SignupDataType,
				Args: gql.FieldConfigArgument{
					"Login": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Password": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Nickname": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolveGetSignupDataByLoginPasswordNickname,
			},
			"Activation": &gql.Field{
				Type: objects.ActivationDataType,
				Args: gql.FieldConfigArgument{
					"ActivationCode": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolveGetActivationDataByCode,
			},
		},
	},
)

func init() {
	AuthQuery = gql.Field{
		Type:    AuthQueryType,
		Resolve: func(p gql.ResolveParams) (interface{}, error) { return AuthData{}, nil },
	}
}
