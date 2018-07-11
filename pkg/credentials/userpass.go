package credentials

import (
	"fmt"
	"encoding/json"
)

type UserPass struct {
	*Base
	Username string
	Password string
}

func (u *UserPass) Init() {
	u.Base = new(Base)
	u.Base.Init()
	u.Type = "userpass"
}

func (u *UserPass) ToJsonString() (string, error) {

	jsonBytes, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return "", err
	}
	return 	fmt.Sprintf(string(jsonBytes)), nil
}

func (u *UserPass) ToRawString() (string, error) {
	//nothing to print for a base rawstring
	return fmt.Sprintf("username=%v\npassword=%v\n", u.Username, u.Password), nil
}

func (u *UserPass) ToTableString() (string, error) {
	//nothing to print for a base tablestring
	return "", nil
}