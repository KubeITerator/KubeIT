package helpers

type ParsedParam struct {
	PType ParamType `json:"paramtype"`
	Line  int       `json:"linenumber"`
	Loc   []int     `json:"location"`
}

type ParamType struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

type MappingList struct {
	MappingList []Mapping `json:"mappinglist"`
}

type Mapping struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Default  string `json:"default"`
	Required bool   `json:"required"`
}

type Defaults struct {
	Default  string `json:"default"`
	Required bool   `json:"required"`
}

type CombinedMappings struct {
	ParsedParam
	Defaults
}

type ConfigMapData struct {
	mappings []CombinedMappings
	yaml     string
}
