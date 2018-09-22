package credentials

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Ssh struct {
	*Generic
}

func (s *Ssh) Init() {
	s.Generic = new(Generic)
	s.Generic.Init()
	s.secretType = "ssh"
}

func (s *Ssh)Key() string {

	if x, found := s.Data["key"]; found {
		return x
	} else {
		return ""
	}
}
func (s *Ssh)SetKey(key string) {
	s.Data["key"] = key
}


func (s *Ssh) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (s *Ssh) ToYamlString() (string, error) {

	yamlBytes, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(yamlBytes)), nil
}

//for Ssh secrets, we expect that the key is the primary data, so we'll return that when Raw is requested. All other data is available in the JSON response.
func (s *Ssh) ToRawString() (string, error) {
	return s.Key(), nil
}

func (s *Ssh) ToTableString() (string, error) {
	//TODO: print table string
	return "", nil
}