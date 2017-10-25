package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"mime/multipart"
	"io"
)

type Message struct {
	Message string `json:"message"`
}

func InterfaceToString(i interface{}) (s string) {
	if i == nil {
		s = ""
	} else {
		s = i.(string)
	}
	return
}
func InterfaceToArrayStrings(i interface{}) (s []string) {
	if i == nil {
		s = make([]string, 0)
	} else {
		s = strings.Split(i.(string), ",")
	}
	return
}

func GetErrorMessageFromResponse(url string, resp *http.Response) (message string) {
	body, _ := ioutil.ReadAll(resp.Body)
	messageObj := Message{}
	json.Unmarshal(body, &messageObj)
	message = messageObj.Message
	return
}

//Вывод ошибки в Get запросе
func LogErrorGet(url string, err error) {
	log.Printf("Get: %s Error %s", url, err)
}
func LogGet(url string, str string) {
	log.Printf("Get: %s %s", url, str)
}

func LogErrorPut(url string, err error) {
	log.Printf("Put: %s Error %s", url, err)
}
func LogPut(url string, str string) {
	log.Printf("Put: %s %s", url, str)
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

func LogErrorDelete(url string, err error) {
	log.Printf("Delete: %s Error %s", url, err)
}
func LogDelete(url string, str string) {
	log.Printf("Delete: %s %s", url, str)
}

func LogServer(str string) {
	log.Printf("Server: %s", str)
}
func LogErrorServer(err error) {
	log.Printf("Server: Error %s", err)
}

func LogErrorAuth(err error) {
	log.Printf("Auth: Error %s", err)
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
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(GetErrorMessageFromResponse(url, resp))
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
	fmt.Println(o)

	LogGet(url, "Success")
	return
}

func HttpPutWithToken(url string, token string, send interface{}, get interface{}) (err error) {
	LogPut(url, "Sending")
	var resp *http.Response
	client := &http.Client{}
	jsonToSend, err := json.Marshal(send)
	bodyToSend := bytes.NewBuffer(jsonToSend)
	req, err := http.NewRequest("PUT", url, bodyToSend)
	if err != nil {
		LogErrorPut(url, err)
		return
	}
	req.Header.Set("Bearer-Token", token)
	resp, err = client.Do(req)
	if err != nil {
		LogErrorPut(url, err)
		return
	}
	LogPut(url, "Received "+resp.Status)
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(GetErrorMessageFromResponse(url, resp))
		LogErrorPut(url, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogErrorPut(url, err)
		return
	}
	err = json.Unmarshal(body, get)
	if err != nil {
		LogErrorPut(url, err)
		return
	}
	LogPut(url, "Success")
	return
}

func HttpGetWithToken(url string, token string,  get interface{}) (err error) {
	LogGet(url, "Sending")
	var resp *http.Response
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		LogErrorGet(url, err)
		return
	}
	req.Header.Set("Bearer-Token", token)
	resp, err = client.Do(req)
	if err != nil {
		LogErrorGet(url, err)
		return
	}
	LogGet(url, "Received "+resp.Status)
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(GetErrorMessageFromResponse(url, resp))
		LogErrorGet(url, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogErrorGet(url, err)
		return
	}
	err = json.Unmarshal(body, get)
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
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(GetErrorMessageFromResponse(url, resp))
		LogErrorPost(url, err)
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
func HttpDelete(url string, send interface{}, get interface{}) (err error) {
	LogDelete(url, "Sending")
	var resp *http.Response
	jsonToSend, err := json.Marshal(send)
	bodyToSend := bytes.NewBuffer(jsonToSend)
	c := &http.Client{}
	req, err := http.NewRequest("DELETE", url, bodyToSend)
	if err != nil {
		LogErrorDelete(url, err)
		return
	}
	resp, err = c.Do(req)
	if err != nil {
		LogErrorDelete(url, err)
		return
	}
	LogDelete(url, "Received "+resp.Status)
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(GetErrorMessageFromResponse(url, resp))
		LogErrorDelete(url, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogErrorDelete(url, err)
		return
	}
	err = json.Unmarshal(body, get)
	if err != nil {
		LogErrorDelete(url, err)
		return
	}
	LogDelete(url, "Success")
	return
}

func HttpPostWithTokenAndFile(url string, token string, file multipart.File, handler *multipart.FileHeader, get interface{}) (err error) {
	LogPost(url, "Sending")
	var resp *http.Response
	client := &http.Client{}
	req, err := newfileUploadRequest(url,"uploadfile",file,handler)
	if err != nil {
		LogErrorPost(url, err)
		return
	}
	req.Header.Set("Bearer-Token", token)
	resp, err = client.Do(req)
	if err != nil {
		LogErrorPost(url, err)
		return
	}
	LogPost(url, "Received "+resp.Status)
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(GetErrorMessageFromResponse(url, resp))
		LogErrorPost(url, err)
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

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(url string, paramName string, file multipart.File, handler *multipart.FileHeader) (*http.Request, error) {
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  part, err := writer.CreateFormFile(paramName, handler.Filename)
  if err != nil {
      return nil, err
  }
  _, err = io.Copy(part, file)
  err = writer.Close()
  if err != nil {
      return nil, err
  }
  req, err := http.NewRequest("POST", url, body)
  req.Header.Set("Content-Type", writer.FormDataContentType())
  return req, err
}