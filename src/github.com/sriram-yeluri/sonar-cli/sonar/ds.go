package sonar

type sonarComponentsStruct struct {
    Paging struct {
        PageIndex int `json:"pageIndex"`
        PageSize  int `json:"pageSize"`
        Total     int `json:"total"`
    } `json:"paging"`
    Components []struct {
        Organization string `json:"organization"`
        ID           string `json:"id"`
        Key          string `json:"key"`
        Name         string `json:"name"`
        Qualifier    string `json:"qualifier"`
        Project      string `json:"project"`
    } `json:"components"`
}