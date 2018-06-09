package glib

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func HttpRequest(user AuthUser){
	fmt.Println(user.Username)
	fmt.Println(user.Password)
}

//https://golang.org/pkg/net/http/
//To make a request with custom headers, use NewRequest and Client.Do. 
//func NewRequest(method, url string, body io.Reader) (*Request, error)

func CreateHttpRequest(method string,url string, user AuthUser) (*http.Request) {
	//Todo - Add argument validations
	if user.Username == "" || user.Password == "" {
		log.Fatal("[Error] Missing user name or password ")
	}
	req, err := http.NewRequest(method, url, nil)
    req.SetBasicAuth(user.Username, user.Password)
    req.Header.Add("accept", "application/json")
    req.Header.Add("content-type", "application/json")
	Error(err,"GetProjects::http NewRequest")
	
	if DEBUG {
		fmt.Println("Debug from CreateHttpRequest function : ")
		fmt.Println("\nRequest URL :", req.URL)
		fmt.Println("\nRequest Header : ", req.Header)
	}
	return req
}

func SendHttpRequest(req *http.Request) (*http.Response, []byte){
	//ToDo - Add argument validations	
	client := &http.Client{}
    resp, err := client.Do(req)
	Error(err,"GetProjects::client.Do")
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	Error(err,"GetProjects::read response body")
	
	if DEBUG {
		fmt.Println("Debug from SendHttpRequest function : ")
		fmt.Println("\n ResponseBody : ", resp.Body)
		fmt.Println("\n Response Header : ", resp.Header)
		fmt.Println("\n Response Status : ", resp.Status)
	}
	return resp, respBody
}