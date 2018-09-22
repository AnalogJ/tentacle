package credentials

import (
	"fmt"
	"encoding/json"
	"net/url"

	"gopkg.in/yaml.v2"
)

type UserPass struct {
	*Generic
}

func (u *UserPass) Init() {
	u.Generic = new(Generic)
	u.Generic.Init()
	u.secretType = "userpass"
}


func (u *UserPass)Username() string {

	if x, found := u.Data["username"]; found {
		return x
	} else {
		return ""
	}
}
func (u *UserPass)SetUsername(text string) {
	u.Data["username"] = text
}

func (u *UserPass)Password() string {

	if x, found := u.Data["password"]; found {
		return x
	} else {
		return ""
	}
}
func (u *UserPass)SetPassword(text string) {
	u.Data["password"] = text
}

func (u *UserPass) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (u *UserPass) ToYamlString() (string, error) {

	yamlBytes, err := yaml.Marshal(u)
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(yamlBytes)), nil
}

func (u *UserPass) ToRawString() (string, error) {
	params := url.Values{}
	for k, v := range u.Data {
		params.Add(k, v)
	}

	return params.Encode(), nil
}
func (u *UserPass) ToTableString() (string, error) {
	//nothing to print for a base tablestring
	return "", nil
}