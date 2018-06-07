package main

import (
	"fmt"
	"flag"
	"log"
	"github.com/sriram-yeluri/sonar-cli/sonar"
	"github.com/sriram-yeluri/sonar-cli/glib"
)

func main() {
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	flag.Parse()
	if *debugMode {
		glib.DEBUG = true
	} else {
		flag.Usage()
		log.Fatal("Select a valid flag")
	}

	sonar.CreateProject()
	sonarURL := "http://localhost:9000"
	fmt.Printf("Project Count = %d\n" , sonar.GetProjects(sonarURL))
}
