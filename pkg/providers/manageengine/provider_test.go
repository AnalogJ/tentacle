package manageengine_test

import (
	"testing"
	"github.com/analogj/tentacle/pkg/validator"
	"github.com/stretchr/testify/suite"
	"github.com/analogj/tentacle/pkg/credentials"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestManageEngineProvider_TestSuite(t *testing.T) {
	testSuite := new(validator.ProviderTestSuite)

	testSuite.ProviderConfig = map[string]interface{}{
		"type": "manageengine",
		"scheme": "http",
		"host": "demo.passwordmanagerpro.com",
		"token": "695B44C2-7EA5-4361-AADC-C241FB885B62",
	}
	testSuite.Get_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"resourceid": "913", "id": "990"},
		Data:      map[string]string{"username":"user", "password":"pT9T71^=q2FhO"},
		Metadata:  map[string]string{"resource_name":"ZTest", "type":"account", "resource_id":"913", "resource_type":"Windows"},
	}

	testSuite.GetById_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"resourceid": "913", "id": "989"},
		Data:      map[string]string{"password":"qc)V29W-%g", "username":"root"},
		Metadata:  map[string]string{"resource_type":"Windows", "resource_name":"ZTest", "type":"account", "resource_id":"913"},
	}

	testSuite.GetByPath_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"path": "wwwwwww/a"},
		Data:      map[string]string{"password":"b", "username":"a"},
		Metadata:  map[string]string{"type":"account", "resource_id":"2128", "resource_type":"Linux", "resource_name":"wwwwwww"},
	}


	testSuite.Get_UserPass_TestData = validator.GetRequestTestData{
		QueryData: map[string]string{"resourceid": "913", "id": "990"},
		Data:      map[string]string{"username":"user", "password":"pT9T71^=q2FhO"},
		Metadata:  map[string]string{"resource_name":"ZTest", "type":"account", "resource_id":"913", "resource_type":"Windows"},
	}


	var summaryList []credentials.SummaryInterface
	summary1 := new(credentials.Summary)
	summary1.Init()
	summary1.Id = "989"
	summary1.Name = "ZTest/root"
	summary1.Metadata = map[string]string {
		"resource_id": "913",
		"resource_name": "ZTest",
		"resource_type": "Windows",
		"type": "account",
	}

	summary2 := new(credentials.Summary)
	summary2.Init()
	summary2.Id = "990"
	summary2.Name = "ZTest/user"
	summary2.Metadata = map[string]string {
		"resource_id": "913",
		"resource_name": "ZTest",
		"resource_type": "Windows",
		"type": "account",
	}

	summaryList = append(summaryList, summary1)
	summaryList = append(summaryList, summary2)

	testSuite.List_TestData = validator.ListRequestTestData{
		QueryData: map[string]string{"resourceid": "913"},
		Results: summaryList,
	}

	suite.Run(t, testSuite)
}