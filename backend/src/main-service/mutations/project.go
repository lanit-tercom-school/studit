package mutations

import (
	"main-service/objects"

	gql "github.com/graphql-go/graphql"
)

type ProjectData struct {
	message objects.Message
}

var PostProject gql.Field

var EditProject gql.Field

var EditProjectQueryType = gql.NewObject(
	gql.ObjectConfig{
		Name: "ProjectQuery",
		Fields: gql.Fields{
			"ChangeName": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Id": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutProjectName,
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
				Resolve: objects.ResolvePutProjectDescription,
			},
			"ChangeLogo": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Id": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutProjectLogo,
			},
			"ChangeGitHubUrl": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Id": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutProjectGitHubUrl,
			},
			"ChangeStatus": &gql.Field{
				Type: objects.MessageType,
				Args: gql.FieldConfigArgument{
					"New": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"Id": &gql.ArgumentConfig{
						Type: gql.String,
					},
				},
				Resolve: objects.ResolvePutProjectStatus,
			},
		},
	},
)

func init() {
	PostProject = gql.Field{
		Type: objects.ProjectType,
		Args: gql.FieldConfigArgument{
			"Name": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"Description": &gql.ArgumentConfig{
				Type: gql.NewNonNull(gql.String),
			},
			"Logo": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Tags": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"GitHubUrl": &gql.ArgumentConfig{
				Type: gql.String,
			},
			"Status": &gql.ArgumentConfig{
				Type:         gql.String,
				DefaultValue: "opened",
			},
		},
		Resolve: objects.ResolvePostProject,
	}
	EditProject = gql.Field{
		Type:    EditProjectQueryType,
		Resolve: func(p gql.ResolveParams) (interface{}, error) { return ProjectData{}, nil },
	}
}
