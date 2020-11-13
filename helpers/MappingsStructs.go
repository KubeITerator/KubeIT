package helpers

type ParsedParam struct {
	Type ParamType `json:"paramtype"`
	Line int       `json:"linenumber"`
	Loc  []int     `json:"location"`
}

type ParamType struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

type MappingList struct {
	Mappings []Mapping
}

type Mapping struct {
	ParsedParam  ParsedParam
	ApiReference string
	Default      interface{}
}
