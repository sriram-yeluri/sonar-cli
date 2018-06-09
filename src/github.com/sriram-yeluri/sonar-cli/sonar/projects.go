package sonar

import (
	"fmt"
	"encoding/json"
	"github.com/sriram-yeluri/sonar-cli/glib"
	"log"
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
        fmt.Println("Project Name : ", component.Name)
    }
    fmt.Println("Total No. of Projects :", len(sonarKeysList))
    return 0
}

/*
@Function to create projects in sonarqube
@Param sonar base url
@Param user name and password from glib.AuthUSer
@return status of project creation

 */
func CreateProject(sonarURL string, user glib.AuthUser, projectStruct ProjectStruct) (int){
	// Check if project exists before creating new project
	projStatus := SearchProject(sonarURL, user,projectStruct)

	if projStatus == 200 {
		url := fmt.Sprintf("%s/api/projects/create", sonarURL)
		req := glib.CreateHttpRequest("POST",url, user)

		//Append POST data
		query := req.URL.Query()
		query.Add("name", projectStruct.ProjectName)
		query.Add("project", projectStruct.ProjectKey)
		req.URL.RawQuery = query.Encode()

		resp, _ := glib.SendHttpRequest(req)
		if resp.StatusCode == 200 {
			fmt.Println("Project Created Successfully : ", projectStruct.ProjectName)
		}
		return resp.StatusCode
	}
	return 1
}

/*
@Function to search for projects in Sonarqube
@ Requires system administrator permission
 */
func SearchProject(sonarURL string, user glib.AuthUser, projectStruct ProjectStruct) (int) {
	url := fmt.Sprintf("%s/api/projects/search", sonarURL)
	req := glib.CreateHttpRequest("GET",url, user)

	//Append Query data
	query := req.URL.Query()
	query.Add("projects", projectStruct.ProjectKey)
	req.URL.RawQuery = query.Encode()

	resp,respBody := glib.SendHttpRequest(req)

	var Search SearchProjectStruct
	json.Unmarshal(respBody, &Search)


	if Search.Paging.Total > 0 {
		log.Fatal("Could not create Project, key already exists")
	}
	return resp.StatusCode
}