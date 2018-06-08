package main

import (
	"fmt"
	"flag"
	"log"
	"github.com/sriram-yeluri/sonar-cli/sonar"
	"github.com/sriram-yeluri/sonar-cli/glib"
)

func main() {
	debugMode := flag.Bool("debug", false, "run in debug mode")
	sonarURL := flag.String("url","http://localhost:9000", "Sonarqube URL")
	username := flag.String("username", "admin", "Sonarqube user name ")
	password := flag.String("password", "admin", "Sonarqube password ")
	flag.Parse()

	//Set credentials
	user := glib.AuthUser{Username: *username, Password: *password}

	if *debugMode {
		glib.DEBUG = true
	}else {
		flag.Usage()
		log.Fatal("Select a valid flag")
	}

	//sonar.CreateProject()
	//sonarURL := "http://localhost:9000"
	count := sonar.GetProjects(*sonarURL, user)
	fmt.Printf("Project Count = %d\n" , count)

}
