package credentials

import (
	"encoding/json"
	"fmt"
	"net/url"

	"gopkg.in/yaml.v2"
)

type Generic struct {
	*Summary
	Data map[string]string `json:"data" yaml:"data"`//this map should contain lowercase keys which map to secrets. ie. username, password, token, etc.
}

func (g *Generic) Init() {
	g.Summary = new(Summary)
	g.Summary.Init()
	g.secretType = "generic"
	g.Data = map[string]string {}
}

func (b *Generic)GetData() map[string]string {
	return b.Data
}

func (g *Generic) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (g *Generic) ToYamlString() (string, error) {

	yamlBytes, err := yaml.Marshal(g)
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(yamlBytes)), nil
}


//for generic secrets, we don't know which secret data is important, so we'll url encode it all and return it.
func (g *Generic) ToRawString() (string, error) {
	params := url.Values{}
	if len(g.Id) >0 {
		params.Add("id", g.Id)
	}
	if len(g.Name) >0 {
		params.Add("name", g.Name)
	}

	for k, v := range g.Data {
		params.Add(k, v)
	}

	return params.Encode(), nil
}

//func (g *Generic) ToTableString() (string, error) {
//	//nothing to print for a base tablestring
//	return "", nil
//}