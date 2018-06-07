package glib

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func HttpRequest(user AuthUser){
	fmt.Println(user.Username)
	fmt.Println(user.Password)
}

//https://golang.org/pkg/net/http/
//To make a request with custom headers, use NewRequest and Client.Do. 
//func NewRequest(method, url string, body io.Reader) (*Request, error)

func CreateHttpRequest(url string, user AuthUser) (*http.Request) {
	//Todo - Add argument validations
	req, err := http.NewRequest("GET", url, nil)
    req.SetBasicAuth(user.Username, user.Password)
    req.Header.Add("accept", "application/json")
    req.Header.Add("content-type", "application/json")
	Error(err,"GetProjects::http NewRequest")
	
	return req
}

func SendHttpRequest(req *http.Request) ([]byte){
	
	//ToDo - Add aurgument validations	
	client := &http.Client{}
    resp, err := client.Do(req)
	Error(err,"GetProjects::client.Do")	
	respBody, err := ioutil.ReadAll(resp.Body)
	Error(err,"GetProjects::read response body")

	return respBody
}