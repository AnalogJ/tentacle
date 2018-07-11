package credentials

import (
	"encoding/json"
	"fmt"
)

type Ssh struct {
	*Text
}

func (s *Ssh) Init() {
	s.Text = new(Text)
	s.Text.Init()
	s.Type = "ssh"
}

func (s *Ssh) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (s *Ssh) ToRawString() (string, error) {
	//nothing to print for a base rawstring
	return s.Data, nil
}

func (s *Ssh) ToTableString() (string, error) {
	//TODO: print table string
	return "", nil
}