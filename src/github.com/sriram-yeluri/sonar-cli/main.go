package main

import (
	"fmt"
	"github.com/sriram-yeluri/sonar-cli/sonar"
)

func main() {
	fmt.Println("Hello World")
	sonar.CreateProject()
	sonarURL := "http://localhost:9000"
	fmt.Printf("Project Count = %d\n" , sonar.GetProjects(sonarURL))
}
