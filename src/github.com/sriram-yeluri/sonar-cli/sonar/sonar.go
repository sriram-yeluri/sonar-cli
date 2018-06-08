package sonar

import (
	"fmt"
	"encoding/json"
	"github.com/sriram-yeluri/sonar-cli/glib"
)


func GetProjects(sonarURL string,user glib.AuthUser) (int){
	url := fmt.Sprintf("%s/api/components/search?qualifiers=TRK", sonarURL)

    req := glib.CreateHttpRequest("GET",url, user)
    respBody := glib.SendHttpRequest(req)

    var sonarComponents sonarComponentsStruct
    var sonarKeysList []string
    json.Unmarshal(respBody, &sonarComponents)
    for _, component := range sonarComponents.Components {
        sonarKeysList = append(sonarKeysList, component.Key)
    }
    return len(sonarKeysList)
}