package objects

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

//Вывод ошибки в Get запросе
func logErrorGet(url string, err error) {
	log.Printf("GET: %s Error %s", url, err)
}

func logGet(url string, str string) {
	log.Printf("Get: %s %s", url, str)
}

//Get запрос с логами
func httpGet(url string, o interface{}) (err error) {
	logGet(url, "Sending")
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		logErrorGet(url, err)
		return
	}
	logGet(url, "Received "+resp.Status)
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
	logGet(url, "Success")
	return
}
