package sonar

import (
	"fmt"
	"encoding/json"
	"github.com/sriram-yeluri/sonar-cli/glib"
)

func CreateProject(){
	fmt.Println("Create Sonar Project")
	user := glib.AuthUser{"sriram","changeit"}
	glib.HttpRequest(user)
}


func GetProjects(sonarURL string) (int){
	url := fmt.Sprintf("%s/api/components/search?qualifiers=TRK", sonarURL)
	
    /*req, err := http.NewRequest("GET", url, nil)
    req.SetBasicAuth("admin", "admin")
    req.Header.Add("accept", "application/json")
    req.Header.Add("content-type", "application/json")
    glib.Error(err,"GetProjects::http NewRequest")
    */
    user := glib.AuthUser{"sriram","changeit"}
    req := glib.CreateHttpRequest(url, user)
    respBody := glib.SendHttpRequest(req)

    /*
    client := &http.Client{}
    resp, err := client.Do(req)
	glib.Error(err,"GetProjects::client.Do")
	
	respBody, err := ioutil.ReadAll(resp.Body)
	glib.Error(err,"GetProjects::read response body")
    */
    var sonarComponents sonarComponentsStruct
    var sonarKeysList []string
    json.Unmarshal(respBody, &sonarComponents)
    for _, component := range sonarComponents.Components {
        sonarKeysList = append(sonarKeysList, component.Key)
    }
    //fmt.Println(sonarKeysList)
    //fmt.Printf("No of components : %d\n", len(sonarKeysList))
    
    return len(sonarKeysList)
}