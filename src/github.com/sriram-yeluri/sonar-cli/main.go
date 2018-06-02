package main

import (
	"fmt"
	"github.com/sriram-yeluri/sonar-cli/sonar"
)

func main() {
	fmt.Println("Hello World")
	sonar.CreateProject()
	sonarURL := "https://localhost"
	sonar.GetProjects(sonarURL)
}
