package sonar

type sonarComponentsStruct struct {
    Components []struct {
        Organization string `json:"organization"`
        ID           string `json:"id"`
        Key          string `json:"key"`
        Name         string `json:"name"`
        Qualifier    string `json:"qualifier"`
        Project      string `json:"project"`
    } `json:"components"`
}

type ProjectStruct struct {
    ProjectName string
    ProjectKey string
}