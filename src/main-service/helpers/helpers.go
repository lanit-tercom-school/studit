package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

//Вывод ошибки в Get запросе
func LogErrorGet(url string, err error) {
	log.Printf("Get: %s Error %s", url, err)
}

func LogGet(url string, str string) {
	log.Printf("Get: %s %s", url, str)
}
func LogServer(str string) {
	log.Printf("Server: %s", str)
}
func LogErrorAuth(err error) {
	log.Printf("Auth: Error %s", err)
}
func LogErrorServer(err error) {
	log.Printf("Server: Error %s", err)
}
func LogAuth(str string) {
	log.Printf("Auth: %s", str)
}

//Get запрос с логами
func HttpGet(url string, o interface{}) (err error) {
	LogGet(url, "Sending")
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		LogErrorGet(url, err)
		return
	}
	LogGet(url, "Received "+resp.Status)
	if resp.StatusCode != 200 {
		err = errors.New("status code is not 200 Ok")
		LogErrorGet(url, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogErrorGet(url, err)
		return
	}
	err = json.Unmarshal(body, o)
	if err != nil {
		LogErrorGet(url, err)
		return
	}
	LogGet(url, "Success")
	return
}
