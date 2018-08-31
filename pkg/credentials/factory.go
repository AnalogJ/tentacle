package credentials

import (
	"log"
)

func Create(secretType string) (Interface, error) {

	//begin switch for providers.
	switch secretType {

	//os specific providers
	case "base":
		secret := new(Base)
		secret.Init()
		return secret, nil

	case "generic":
		secret := new(Generic)
		secret.Init()
		return secret, nil

	case "ssh":
		secret := new(Ssh)
		secret.Init()
		return secret, nil

	case "text":
		secret := new(Text)
		secret.Init()
		return secret, nil

	case "userpass":
		secret := new(UserPass)
		secret.Init()
		return secret, nil
		//fall back error message

	default:
		log.Fatalf("%v type is not supported by tentacle", secretType)
		return nil, nil
	}
}