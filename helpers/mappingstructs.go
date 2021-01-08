package helpers

type ParsedParam struct {
	ParamType `json:",inline"`
	Line      int    `json:"linenumber"`
	Loc       []int  `json:"location"`
	Default   string `json:"default"`
}

type ParamType struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

type Template struct {
	Name    string        `json:"name"`
	Yaml    string        `json:"yaml"`
	PParams []ParsedParam `json:"params"`
}

// This can be used to store additional information.
type ConfigMapData struct {
	Templates []Template `json:"templates"`
}

type FinalMapping struct {
	ParsedParam `json:",inline"`
	FinalValue  string `json:"finalValue"`
}
