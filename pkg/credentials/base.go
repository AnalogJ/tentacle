package credentials

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// The base class is used during list operations. It does not include any secret data.
type Base struct {
	Metadata map[string]string `json:"metadata"`
	Id string `json:"id"`//id should be a unique identifier used to differentiate between secrets/credentials. Can be just a name on some systems.
	Name string `json:"name"`// used to visually differentiate between credentails.
	Description string `json:"description"`
	secretType string `json:"secretType"`//this value will change depending on the secret type (base, ssh, text, userpass, generic, etc)
}

func (b *Base) Init() {
	//do nothing
	b.Metadata = map[string]string{}
	b.secretType = "base"
	b.Id = ""
	b.Name = ""
	b.Description = ""
}

func (b *Base)SecretType() string {
	return b.secretType
}


func (b *Base) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(b, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (b *Base) ToRawString() (string, error) {
	params := url.Values{}
	if len(b.Id) >0 {
		params.Add("id", b.Id)
	}
	if len(b.Name) >0 {
		params.Add("name", b.Name)
	}
	return params.Encode(), nil
}

func (b *Base) ToTableString() (string, error) {
	//nothing to print for a base tablestring
	return "", nil
}