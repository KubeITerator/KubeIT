package helpers

type ParsedParam struct {
	ParamType `json:",inline"`
	Line      int   `json:"linenumber"`
	Loc       []int `json:"location"`
}

type ParamType struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

type Mapping struct {
	ParamType `json:",inline"`
	Defaults  `json:",inline"`
}

type Defaults struct {
	Default  string `json:"default"`
	Required bool   `json:"required"`
}

type CombinedMappings struct {
	ParsedParam `json:",inline"`
	Defaults    `json:",inline"`
}

type ConfigMapData struct {
	mappings []CombinedMappings
	yaml     string
}
