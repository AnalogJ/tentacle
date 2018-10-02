package lastpass_test

import (
	"testing"
	"github.com/analogj/tentacle/pkg/validator"
	"github.com/stretchr/testify/suite"
	"github.com/analogj/tentacle/pkg/credentials"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestLastpassProvider_TestSuite(t *testing.T) {
	testSuite := new(validator.ProviderTestSuite)

	testSuite.ProviderConfig = map[string]interface{}{
		"type": "lastpass",
		"username": "tentacle@mailinator.com",
		"password": "29IJLA4cqO$nC5#9jO@k", //junk password only used for testing. please don't change it :(
	}
	testSuite.Get_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "7133750796722550877"},
		Data: map[string]string{"username":"test", "password":"W95K04mHaL8PEBKgpCoo"},
		MetaData: map[string]string{"url":"http://www.example.com", "group":"", "notes":"test notes for example.com user/pass"},
	}

	testSuite.GetById_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "3219534729461706852"},
		Data: map[string]string{"username":"tentacle", "password":"Fxw9NNbmkMRV816vbAEt"},
		MetaData: map[string]string{"notes":"example.org test username password pair. ", "url":"https://www.example.org", "group":"Credit Cards"},
	}

	testSuite.Get_UserPass_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"id": "3219534729461706852"},
		Data: map[string]string{"password":"Fxw9NNbmkMRV816vbAEt", "username":"tentacle"},
		MetaData: map[string]string{"notes":"example.org test username password pair. ", "url":"https://www.example.org", "group":"Credit Cards"},
	}

	var summaryList []credentials.SummaryInterface
	summary1 := new(credentials.Summary)
	summary1.Init()
	summary1.Id = "7133750796722550877"
	summary1.Name = "example.com test"
	summary1.Metadata = map[string]string {
		"group": "",
		"notes": "test notes for example.com user/pass",
		"url": "http://www.example.com",
		"username": "test",
	}

	summary2 := new(credentials.Summary)
	summary2.Init()
	summary2.Id = "3219534729461706852"
	summary2.Name = "example.org"
	summary2.Metadata = map[string]string {
		"group": "Credit Cards",
		"notes": "example.org test username password pair. ",
		"url": "https://www.example.org",
		"username": "tentacle",
	}

	summaryList = append(summaryList, summary1)
	summaryList = append(summaryList, summary2)

	testSuite.List_TestData = validator.ListRequestTestData{
		QueryData: map[string]string{},
		Results: summaryList,
	}

	suite.Run(t, testSuite)
}