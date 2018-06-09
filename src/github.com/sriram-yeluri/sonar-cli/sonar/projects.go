package sonar

import (
	"fmt"
	"encoding/json"
	"github.com/sriram-yeluri/sonar-cli/glib"
)


func GetProjects(sonarURL string,user glib.AuthUser) (int){
	url := fmt.Sprintf("%s/api/components/search?qualifiers=TRK", sonarURL)

    req := glib.CreateHttpRequest("GET",url, user)
    _,respBody := glib.SendHttpRequest(req)

    var sonarComponents sonarComponentsStruct
    var sonarKeysList []string
    json.Unmarshal(respBody, &sonarComponents)
    for _, component := range sonarComponents.Components {
        sonarKeysList = append(sonarKeysList, component.Key)
        //sonarKeysList = append(sonarKeysList, component.Name)
        fmt.Println("Project Name : ", component.Name)
    }

    return len(sonarKeysList)

}

/*
@Function to create projects in sonarqube
@Param sonar base url
@Param user name and password from glib.AuthUSer
@return status of project creation

 */
func CreateProject(sonarURL string, user glib.AuthUser, proj ProjectStruct) (int){

	//TODO : Check if project exists before creating new project
	url := fmt.Sprintf("%s/api/projects/create", sonarURL)
	req := glib.CreateHttpRequest("POST",url, user)

	//Append POST data
	query := req.URL.Query()
	query.Add("project", proj.ProjectName)
	query.Add("name", proj.ProjectKey)
	req.URL.RawQuery = query.Encode()

	resp, _ := glib.SendHttpRequest(req)
	fmt.Println("Project Creation status :", resp.Status)

	return resp.StatusCode
}