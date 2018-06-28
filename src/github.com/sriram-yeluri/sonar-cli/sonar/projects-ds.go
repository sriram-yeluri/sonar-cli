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


type SearchProjectStruct struct {
    Paging struct {
        PageIndex int `json:"pageIndex"`
        PageSize  int `json:"pageSize"`
        Total     int `json:"total"`
    }
}
/*
type Components []struct {
    Organization string `json:"organization"`
    ID           string `json:"id"`
    Key          string `json:"key"`
    Name         string `json:"name"`
    Qualifier    string `json:"qualifier"`
    Project      string `json:"project"`
}
*/

type ComponentsStruct struct {
	Paging struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Total     int `json:"total"`
	}
	Components []struct {
		Organization string `json:"organization"`
		ID           string `json:"id"`
		Key          string `json:"key"`
		Name         string `json:"name"`
		Qualifier    string `json:"qualifier"`
		Project      string `json:"project"`
	}
}