package objects

import (
	"errors"
	"main-service/conf"

	gql "github.com/graphql-go/graphql"
)

type User struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Description string `json:"description,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
}


var UserType = gql.NewObject(
	gql.ObjectConfig{
		Name: "User",
		Fields: gql.Fields{
			"Id": &gql.Field{
				Type: gql.String,
			},
			"Nickname": &gql.Field{
				Type: gql.String,
			},
			"Description": &gql.Field{
				Type: gql.String,
			},
			"Avatar": &gql.Field{
				Type: gql.String,
			},
		},
	},
)

func ResolveGetUser(p gql.ResolveParams) (interface{}, error) {
	var id string
	id, ok := p.Args["Id"].(string)
	if !ok {
		err := errors.New("missed id")
		return nil, err
	}
	var user User
	err := httpGet(conf.Configuration.DataServiceURL+"v1/user/id/"+id, &user)
	return user, err
}
