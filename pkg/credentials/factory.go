package credentials

import (
	"log"
)

func CreateSummary() (SummaryInterface){
	summary := new(Summary)
	summary.Init()
	return summary
}

func Create(secretType string) (GenericInterface, error) {

	//begin switch for credential types.
	switch secretType {

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