package credentials

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// The base class is used during list operations. It does not include any secret data.
type Summary struct {
	Metadata map[string]string `json:"metadata"`
	Id string `json:"id"`//id should be a unique identifier used to differentiate between secrets/credentials. Can be just a name on some systems.
	Name string `json:"name"`// used to visually differentiate between credentails.
	Description string `json:"description"`
	secretType string `json:"secretType"`//this value will change depending on the secret type (base, ssh, text, userpass, generic, etc)
}

func (s *Summary) Init() {
	//do nothing
	s.Metadata = map[string]string{}
	s.secretType = "summary"
	s.Id = ""
	s.Name = ""
	s.Description = ""
}

func (s *Summary)GetSecretType() string {
	return s.secretType
}

func (s *Summary)GetSecretId() string {
	return s.Id
}

func (s *Summary)GetMetaData() map[string]string {
	return s.Metadata
}

func (s *Summary) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (s *Summary) ToRawString() (string, error) {
	params := url.Values{}
	if len(s.Id) >0 {
		params.Add("id", s.Id)
	}
	if len(s.Name) >0 {
		params.Add("name", s.Name)
	}
	return params.Encode(), nil
}

func (s *Summary) ToTableString() (string, error) {
	//nothing to print for a base tablestring
	return "", nil
}