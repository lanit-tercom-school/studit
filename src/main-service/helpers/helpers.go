package helpers

import (
	"bytes"
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
func LogAccesAllowed(str string) {
	LogAuth("Access is allowed to " + str)
}
func LogAccesDenied(str string) {
	LogAuth("Access is denied to " + str)
}
func LogErrorPost(url string, err error) {
	log.Printf("Post: %s Error %s", url, err)
}

func LogPost(url string, str string) {
	log.Printf("Post: %s %s", url, str)
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
		err = errors.New("status code is not Ok")
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

//Post запрос с логами
func HttpPost(url string, send interface{}, get interface{}) (err error) {
	LogPost(url, "Sending")
	var resp *http.Response
	jsonToSend, err := json.Marshal(send)
	bodyToSend := bytes.NewBuffer(jsonToSend)

	if resp, err = http.Post(url, "application/json", bodyToSend); err != nil {
		LogErrorPost(url, err)
		return
	}
	LogPost(url, "Received "+resp.Status)
	if !(resp.StatusCode == 200 || resp.StatusCode == 201) {
		err = errors.New("status code is not Ok")
		LogErrorGet(url, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogErrorPost(url, err)
		return
	}

	err = json.Unmarshal(body, get)
	if err != nil {
		LogErrorPost(url, err)
		return
	}
	LogPost(url, "Success")
	return
}
