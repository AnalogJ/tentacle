package credentials

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Text struct {
	*Generic
}

func (t *Text) Init() {
	t.Generic = new(Generic)
	t.Generic.Init()
	t.secretType = "text"
}

func (t *Text)Text() string {

	if x, found := t.Data["text"]; found {
		return x
	} else {
		return ""
	}
}
func (t *Text)SetText(text string) {
	t.Data["text"] = text
}

func (t *Text) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (t *Text) ToYamlString() (string, error) {

	yamlBytes, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(yamlBytes)), nil
}

func (t *Text) ToRawString() (string, error) {
	//nothing to print for a base rawstring
	return t.Text(), nil
}

func (t *Text) ToTableString() (string, error) {
	//TODO: print table string
	return "", nil
}