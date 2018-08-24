package credentials

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Generic struct {
	*Base
	Data map[string]string `json:"data"`//this map should contain lowercase keys which map to secrets. ie. username, password, token, etc.
}

func (g *Generic) Init() {
	g.Base = new(Base)
	g.Base.Init()
	g.secretType = "generic"
	g.Data = map[string]string {}
}

func (g *Generic) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
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