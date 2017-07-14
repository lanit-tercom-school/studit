package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	FileServiceURL string
	AuthServiceURL string
	DataServiceURL string
}

//Configuration - object containing configuration information
var Configuration configuration

func init() {
	var err error
	Configuration = configuration{}
	file, err := os.Open("conf/conf.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}
