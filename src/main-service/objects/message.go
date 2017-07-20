package objects

import gql "github.com/graphql-go/graphql"

type Message struct {
	Message string `json:"message"`
}

var MessageType = gql.NewObject(
	gql.ObjectConfig{
		Name: "Message",
		Fields: gql.Fields{
			"Message": &gql.Field{
				Type: gql.String,
			},
		},
	},
)
