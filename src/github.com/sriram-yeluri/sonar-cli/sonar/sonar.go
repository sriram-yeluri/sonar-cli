package sonar

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/sriram-yeluri/sonar-cli/glib"
)

func CreateProject(){
	fmt.Println("Create Sonar Project")
}

func GetProjects(sonarURL string){
    url := fmt.Sprintf("%s/api/components/search?qualifiers=TRK", sonarURL)
    req, err := http.NewRequest("GET", url, nil)
    req.SetBasicAuth("******", "******")
    req.Header.Add("accept", "application/json")
    req.Header.Add("content-type", "application/json")
    if err != nil{
        log.Fatal("There was an error creating the request")
    }

    client := &http.Client{}

    resp, err := client.Do(req)
    glib.Error(err,"GetProjects::send http request")
	respBody, err := ioutil.ReadAll(resp.Body)
	glib.Error(err,"GetProjects::read response body")
    var sonarComponents sonarComponentsStruct
    var sonarKeysList []string
    json.Unmarshal(respBody, &sonarComponents)
    for _, component := range sonarComponents.Components {
        sonarKeysList = append(sonarKeysList, component.Key)
    }
    fmt.Println(sonarKeysList)
    fmt.Printf("No of components : %d", len(sonarKeysList))
}