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

type Mapping struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Default  string `json:"default"`
	Required bool   `json:"required"`
}
