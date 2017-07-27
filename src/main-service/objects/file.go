package objects

import (
	"errors"
	"time"

	"main-service/conf"
	"main-service/helpers"

	gql "github.com/graphql-go/graphql"
)

type File struct {
	Id             int      
	Name           string   
	Path           string    
	DateOfCreation time.Time 
	User           int     
}

var FileType = gql.NewObject(
	gql.ObjectConfig{
		Name: "File",
		Fields: gql.Fields{
			"Id": &gql.Field{
				Type: gql.String,
			},
			"Name": &gql.Field{
				Type: gql.String,
			},
			"Path": &gql.Field{
				Type: gql.String,
			},
			"DateOfCreation": &gql.Field{
				Type: gql.String,
			},
			"User": &gql.Field{
				Type: gql.Int,
			},
		},
	},
)

type FileList struct{
	Files     []File
}

var FileListType = gql.NewObject(
	gql.ObjectConfig{
		Name: "FileList",
		Fields: gql.Fields{
			"Files": &gql.Field{
				Type: gql.NewList(FileType),
			},
		},
	},
)


func ResolveGetFileById(p gql.ResolveParams) (interface{}, error) {
	var id string
	token:=helpers.InterfaceToString(p.Context.Value("Token"));
	id, ok := p.Args["Id"].(string)
	if !ok {
		err := errors.New("missed id")
		return nil, err
	}
	file:=File{};
	err := helpers.HttpGetWithToken(conf.Configuration.FileServiceURL+"v1/files/"+id,token, &file)
	return file, err
}

func ResolveGetFileList(p gql.ResolveParams) (interface{}, error) {
	var id string
	token:=helpers.InterfaceToString(p.Context.Value("Token"));
	id, ok := p.Args["Id"].(string)
	var files []File
	var err error
	if !ok {
		err = helpers.HttpGetWithToken(conf.Configuration.FileServiceURL+"v1/files/",token, &files)
	} else {
		err = helpers.HttpGetWithToken(conf.Configuration.FileServiceURL+"v1/files/?Id="+id,token, &files)
	}
	fileList:=FileList{
		Files: files,
	}
	return fileList, err
}

