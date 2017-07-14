package objects

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func logErrorGet(url string, err error) {
	log.Printf("Error GET: %s----%s", url, err)
}

func httpGet(url string, o interface{}) (err error) {
	log.Printf("Sending GET: %s", url)
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		logErrorGet(url, err)
		return
	}
	log.Printf("Received GET: %s----%s", url, resp.Status)
	if resp.StatusCode != 200 {
		err = errors.New("status code is not 200 Ok")
		logErrorGet(url, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logErrorGet(url, err)
		return
	}
	err = json.Unmarshal(body, o)
	if err != nil {
		logErrorGet(url, err)
		return
	}
	log.Printf("Success GET: %s", url)
	return
}
