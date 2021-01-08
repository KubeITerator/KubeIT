package helpers

import (
	"bufio"
	"errors"
	"regexp"
	"strings"
)

type YamlParser struct {
	kubitRegex   *regexp.Regexp
	defaultRegex *regexp.Regexp
}

func (yp *YamlParser) Init() error {
	var err error
	yp.kubitRegex, err = regexp.Compile("{{kubeit\\..*\\..*}}")
	yp.defaultRegex, err = regexp.Compile("`# Default: \".*\"`gm")
	return err
}

func (yp *YamlParser) ParseYaml(yaml string) (matches []ParsedParam, err error) {
	counter := 0
	scanner := bufio.NewScanner(strings.NewReader(yaml))

	match := ParsedParam{}

	for scanner.Scan() {
		line := scanner.Text()
		if loc := yp.kubitRegex.FindStringIndex(line); loc != nil {
			pTypeStrings := strings.Split(strings.TrimRight(strings.TrimLeft(line[loc[0]:loc[1]], "{{"), "}}"), ".")

			def := yp.defaultRegex.FindString(line)
			if def != "" {
				def = def[12 : len(def)-1]
			}

			var pType ParamType
			if pTypeStrings[0] == "kubeit" {
				pType = ParamType{pTypeStrings[1], pTypeStrings[2]}
			} else {
				return nil, errors.New("parsing failed: contained non kubeIT parameter")
			}
			match = ParsedParam{Line: counter, ParamType: pType, Loc: loc, Default: def}
			matches = append(matches, match)
		}
		counter++
	}

	return matches, nil

}
