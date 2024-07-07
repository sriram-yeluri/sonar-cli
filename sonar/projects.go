package sonar

import (
	"fmt"
	"encoding/json"
	"github.com/sriram-yeluri/sonar-cli/utils"
	"log"
)

/*
func GetProjects(sonarURL string,user utils.AuthUser) (int){
	url := fmt.Sprintf("%s/api/components/search?qualifiers=TRK", sonarURL)

    req := utils.CreateHttpRequest("GET",url, user)
    _,respBody := utils.SendHttpRequest(req)

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
*/

func GetProjects(sonarURL string,user utils.AuthUser) (int){

	url := fmt.Sprintf("%s/api/components/search", sonarURL)
	var sonarComponents ComponentsStruct
	var sonarKeysList []string
	page := 1
	pageSize := 500
	pagesRemaining := 0

	req := utils.CreateHttpRequest("GET",url, user)

	query := req.URL.Query()
	query.Add("qualifiers","TRK")
	query.Add("p", fmt.Sprintf("%v",page))
	query.Add("ps",fmt.Sprintf("%v",pageSize))
	req.URL.RawQuery = query.Encode()
	_,respBody := utils.SendHttpRequest(req)

	json.Unmarshal(respBody, &sonarComponents)


	for _, component := range sonarComponents.Components {
		sonarKeysList = append(sonarKeysList, component.Key)
		//fmt.Println("Project Name : ", component.Name)
	}

	if sonarComponents.Paging.Total > 500 {
		pagesRemaining += (sonarComponents.Paging.Total/500) -1
		if (sonarComponents.Paging.Total%500) > 1{
			pagesRemaining ++
		}
		fmt.Println("Remaining Pages : ",pagesRemaining)
		for pagesRemaining > 0 {
			page ++
			query.Add("p", fmt.Sprintf("%v",page))
			req.URL.RawQuery = query.Encode()
			_,respBody := utils.SendHttpRequest(req)
			json.Unmarshal(respBody, &sonarComponents)
			for _, component := range sonarComponents.Components {
				sonarKeysList = append(sonarKeysList, component.Key)
				//fmt.Println("Project Name : ", component.Name)
			}
			pagesRemaining --
		}
	}
	fmt.Println("Total No. of Projects :", len(sonarKeysList))
	fmt.Println("Paging Total : ",sonarComponents.Paging.Total)
	return 0
}

/*
@Function to create projects in sonarqube
@Param sonar base url
@Param user name and password from glib.AuthUSer
@return status of project creation

 */
func CreateProject(sonarURL string, user utils.AuthUser, projectStruct ProjectStruct) (int){
	// Check if project exists before creating new project
	projStatus := SearchProject(sonarURL, user,projectStruct)

	if projStatus == 200 {
		url := fmt.Sprintf("%s/api/projects/create", sonarURL)
		req := utils.CreateHttpRequest("POST",url, user)

		//Append POST data
		query := req.URL.Query()
		query.Add("name", projectStruct.ProjectName)
		query.Add("project", projectStruct.ProjectKey)
		req.URL.RawQuery = query.Encode()

		resp, _ := utils.SendHttpRequest(req)

		if utils.DEBUG {
			fmt.Println("[DEBUG] fromn CreateProject function ")
			fmt.Println("[DEBUG] response : ", resp.StatusCode)
		}

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
func SearchProject(sonarURL string, user utils.AuthUser, projectStruct ProjectStruct) (int) {
	url := fmt.Sprintf("%s/api/projects/search", sonarURL)
	req := utils.CreateHttpRequest("GET",url, user)

	//Append Query data
	query := req.URL.Query()
	query.Add("projects", projectStruct.ProjectKey)
	req.URL.RawQuery = query.Encode()

	resp,respBody := utils.SendHttpRequest(req)

	var Search SearchProjectStruct
	json.Unmarshal(respBody, &Search)

	if utils.DEBUG {
		fmt.Println("[DEBUG] from SearchProject Function")
		fmt.Println("[DEBUG] Project Key : ", projectStruct.ProjectKey)
		fmt.Println("[DEBUG] Project Name : ", projectStruct.ProjectName)
		fmt.Println("[DEBUG] Search query response : ", resp.StatusCode)
	}

	if Search.Paging.Total > 0 {
		log.Fatal("Could not create Project, key already exists")
	}
	return resp.StatusCode
}