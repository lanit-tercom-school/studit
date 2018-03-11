package conf

import (
	"encoding/json"
	"main-service/helpers"
	"os"
)

type configuration struct {
	FileServiceURL string
	AuthServiceURL string
	DataServiceURL string
	JwtSecret      string
	HttpPort       string
	FilesURL	   string
}

//Configuration - object containing configuration information
var Configuration configuration

func init() {
	helpers.LogServer("Parsing conf.json")
	var err error
	Configuration = configuration{}
	file, err := os.Open("conf/conf.json")
	if err != nil {
		helpers.LogErrorServer(err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		helpers.LogErrorServer(err)
		panic(err)
	}
	helpers.LogServer("Parsing conf.json successfully")
}
