package apistructs

type WorkflowParams []struct {
	Category string      `json:"category"`
	Name     string      `json:"name"`
	Value    interface{} `json:"value"`
}
