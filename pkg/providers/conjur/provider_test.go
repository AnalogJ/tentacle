package conjur_test

import (
	"testing"
	"github.com/analogj/tentacle/pkg/validator"
	"github.com/stretchr/testify/suite"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestConjurProvider_TestSuite(t *testing.T) {
	testSuite := new(validator.ProviderTestSuite)

	testSuite.ProviderConfig = map[string]interface{}{
		"type": "conjur",
		"login": "admin",
		"api_key": "2t0wgzw3ap986c1fsnrw73az8z0r1zjcy802yzfd6e2wtqjnv2z94m2c", //junk password only used for testing. please don't change it :(
		"appliance_url": "https://eval.conjur.org",
		"account": "tentacle@mailinator.com",
	}
	testSuite.Get_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "eval/secret"},
		Data:      map[string]string{"text":"f2946778f7a86c419cf00eeb"},
		Metadata:  map[string]string{},
	}

	testSuite.GetById_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "eval/secret"},
		Data:      map[string]string{"text":"f2946778f7a86c419cf00eeb"},
		Metadata:  map[string]string{},
	}

	testSuite.Get_Text_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "eval/secret"},
		Data:      map[string]string{"text":"f2946778f7a86c419cf00eeb"},
		Metadata:  map[string]string{},
	}

	suite.Run(t, testSuite)
}