package apistructs

type IteratedWorkflow struct {
	Files    []string `json:"files"`
	Splitter Splitter `json:"splitter"`
	Workload Workflow `json:"workload"`
}
type Params struct {
	Parameter string `json:"parameter"`
}
type Splitter struct {
	Type   string   `json:"type"`
	Params []Params `json:"params"`
}
type Workload struct {
	Name string `json:"name"`
	Exec string `json:"exec,omitempty"`
}
type Workflow struct {
	InitWorkload Workload `json:"init-workload,omitempty"`
	MainWorkload Workload `json:"main-workload"`
}
