package main

import (
	"flag"
	"log"
	"github.com/sriram-yeluri/sonar-cli/sonar"
	"github.com/sriram-yeluri/sonar-cli/utils"
)

func main() {
	debugMode := flag.Bool("debug", false, "run in debug mode")
	sonarURL := flag.String("url","http://localhost:9000", "Sonarqube URL")
	username := flag.String("username", "admin", "Sonarqube user name ")
	password := flag.String("password", "admin", "Sonarqube password ")
	getprojects := flag.Bool("getProjects", false, "Get list of all Sonarqube projects")
	projname := flag.String("projectName", "", "Project name")
	projkey := flag.String("projectKey", "", "project key")
	createProject := flag.Bool("createProject", false, "Create sonarqube project")
	flag.Parse()

	//Set credentials
	user := utils.AuthUser{Username: *username, Password: *password}
	project := sonar.ProjectStruct{ProjectName:*projname, ProjectKey:*projkey}

	if *debugMode {
		utils.DEBUG = true
	}

	if *createProject {
		sonar.CreateProject(*sonarURL, user, project)
	}else if *getprojects {
		sonar.GetProjects(*sonarURL, user)
	}else {
		flag.Usage()
		log.Fatal("Select a valid flag")
	}

	//sonar.CreateProject()
	//sonarURL := "http://localhost:9000"
	//count := sonar.GetProjects(*sonarURL, user)
	//fmt.Printf("Project Count = %d\n" , count)

}
