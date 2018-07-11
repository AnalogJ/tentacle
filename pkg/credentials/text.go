package credentials

import (
	"encoding/json"
	"fmt"
)

type Text struct {
	*Base
	Data string
}

func (t *Text) Init() {
	t.Base = new(Base)
	t.Base.Init()
	t.Type = "text"
}

func (t *Text) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (t *Text) ToRawString() (string, error) {
	//nothing to print for a base rawstring
	return t.Data, nil
}

func (t *Text) ToTableString() (string, error) {
	//TODO: print table string
	return "", nil
}