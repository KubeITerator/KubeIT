package helpers

import (
	"bufio"
	"regexp"
	"strings"
)

type YamlParser struct {
	kubitRegex *regexp.Regexp
}

type Match struct {
	Type ParamType
	Line int
	Loc  []int
}

type ParamType struct {
	Category string
	Name     string
}

func (yp *YamlParser) Init() error {
	var err error
	yp.kubitRegex, err = regexp.Compile("{{kubeit\\..*\\..*}}")
	return err
}

func (yp *YamlParser) ParseYaml(yaml string) (matches []Match) {
	counter := 0
	scanner := bufio.NewScanner(strings.NewReader(yaml))

	match := Match{}

	for scanner.Scan() {
		if loc := yp.kubitRegex.FindStringIndex(scanner.Text()); loc != nil {

			pType := ParamType{scanner.Text()}
			match = Match{Line: counter, Type: scanner.Text()[loc[0]:loc[1]], Loc: loc}
			matches = append(matches, match)
		}
		counter++
	}

	return matches
}
