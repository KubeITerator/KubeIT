package apistructs

type WorkflowParams []struct {
	Parameter string      `json:"param"`
	Value     interface{} `json:"value"`
}
