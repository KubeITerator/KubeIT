package helpers

import (
	"bufio"
	"errors"
	"regexp"
	"strings"
)

type YamlParser struct {
	kubitRegex *regexp.Regexp
}

type Match struct {
	Type ParamType `json:"paramtype"`
	Line int       `json:"linenumber"`
	Loc  []int     `json:"location"`
}

type ParamType struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

func (yp *YamlParser) Init() error {
	var err error
	yp.kubitRegex, err = regexp.Compile("{{kubeit\\..*\\..*}}")
	return err
}

func (yp *YamlParser) ParseYaml(yaml string) (matches []Match, err error) {
	counter := 0
	scanner := bufio.NewScanner(strings.NewReader(yaml))

	match := Match{}

	for scanner.Scan() {
		if loc := yp.kubitRegex.FindStringIndex(scanner.Text()); loc != nil {
			pTypeStrings := strings.Split(strings.TrimRight(strings.TrimLeft(scanner.Text()[loc[0]:loc[1]], "{{"), "}}"), ".")
			var pType ParamType
			if pTypeStrings[0] == "kubeit" {
				pType = ParamType{pTypeStrings[1], pTypeStrings[2]}
			} else {
				return nil, errors.New("parsing failed: contained non kubeIT parameter")
			}
			match = Match{Line: counter, Type: pType, Loc: loc}
			matches = append(matches, match)
		}
		counter++
	}

	return matches, nil
}
