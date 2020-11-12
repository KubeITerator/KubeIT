package apistructs

type IteratedWorkflow struct {
	File         string   `json:"files"`
	SharedVolume Volume   `json:"sharedVolume,omitempty"`
	Splitter     Splitter `json:"splitter"`
	Workload     Workload `json:"workload"`
}
type Params struct {
	Parameter string `json:"parameter"`
}
type Splitter struct {
	Type   string   `json:"type"`
	Params []Params `json:"params"`
}
type WorkflowTemplate struct {
	Name       string `json:"name"`
	Entrypoint string `json:"entrypoint"`
	Exec       string `json:"exec,omitempty"`
}
type Workload struct {
	InitWorkload WorkflowTemplate `json:"initWorkload,omitempty"`
	MainWorkload WorkflowTemplate `json:"mainWorkload"`
}

type Volume struct {
	Name         string `json:"name"`
	Size         string `json:"size"`
	StorageClass string `json:"storageClass"`
}
