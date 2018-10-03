package validator

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/analogj/tentacle/pkg/providers"
	"strings"
	"github.com/analogj/tentacle/pkg/credentials"
	"github.com/analogj/tentacle/pkg/constants"
)
// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ProviderTestSuite struct {
	suite.Suite
	Provider 				providers.Interface


	// the following values must be set in the calling provider test file.
	ProviderConfig        	map[string]interface{}
	Get_TestData          	GetRequestTestData
	GetById_TestData      	GetRequestTestData
	GetByPath_TestData		GetRequestTestData
	Get_Ssh_TestData      	GetRequestTestData
	Get_Text_TestData     	GetRequestTestData
	Get_UserPass_TestData 	GetRequestTestData

	List_TestData 		  	ListRequestTestData
}

type GetRequestTestData struct{
	//query data is submitted to the .Get function during the test
	QueryData map[string]string

	//data and metadata should be the expected values returned by the test
	Data     map[string]string
	Metadata map[string]string
}

type ListRequestTestData struct{
	//query data is submitted to the .Get function during the test
	QueryData map[string]string

	//results should be the expected values returned by the test
	Results []credentials.SummaryInterface
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ProviderTestSuite) SetupTest() {

	p, err := providers.Create(suite.ProviderConfig["type"].(string), suite.ProviderConfig)
	require.NoError(suite.T(), err)
	p.SetHttpClient(ProviderVcrSetup(suite.T()))
	suite.Provider = p

}

func (suite *ProviderTestSuite) TearDownTest() {

}

// Suite Code.

func (suite *ProviderTestSuite)TestProvider_New() {

	//setup

	//test

	//assert
	require.Implements(suite.T(), (*providers.Interface)(nil), suite.Provider, "should implement the provider interface")
	require.NotEmpty(suite.T(), suite.Provider.Capabilities(), "should have a valid list of capabilities, see pkg/providers/interface.go for more info")
}


func (suite *ProviderTestSuite)TestProvider_New_WithInvalidCredentialsShouldFail() {

	//setup

	//test
	_, cerr := providers.Create(suite.ProviderConfig["type"].(string), map[string]interface{}{"type": suite.ProviderConfig["type"].(string)})

	//assert
	require.Error(suite.T(), cerr)
}

func (suite *ProviderTestSuite)TestProvider_Authenticate() {

	//setup

	//test
	err := suite.Provider.Authenticate()

	//assert
	require.NoError(suite.T(), err)
}

func (suite *ProviderTestSuite)TestProvider_Get_WithEmptyConfigShouldFail() {

	//setup
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	_, err := suite.Provider.Get(map[string]string{})

	//assert
	require.Error(suite.T(), err, "should use validator functions to raise an error if empty query is passed into `.Get`")
}

func (suite *ProviderTestSuite)TestProvider_GetById() {

	//setup
	if enabled, ok := suite.Provider.Capabilities()[constants.CAP_GET_BY_ID]; !ok || !enabled {
		suite.T().Skip("skipping test, provider does not support it")
	}
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	cred, err := suite.Provider.Get(suite.GetById_TestData.QueryData)

	//assert
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), suite.GetById_TestData.QueryData["id"], cred.GetSecretId())
	require.Equal(suite.T(), suite.GetById_TestData.Data, cred.GetData())
	require.Equal(suite.T(), suite.GetById_TestData.Metadata, cred.GetMetadata())
}

func (suite *ProviderTestSuite)TestProvider_GetByPath() {

	//setup
	if enabled, ok := suite.Provider.Capabilities()[constants.CAP_GET_BY_PATH]; !ok || !enabled {
		suite.T().Skip("skipping test, provider does not support it")
	}
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	cred, err := suite.Provider.Get(suite.GetByPath_TestData.QueryData)

	//assert
	require.NoError(suite.T(), err)
	require.NotEmpty(suite.T(), suite.GetByPath_TestData.QueryData["path"])
	require.Equal(suite.T(), suite.GetByPath_TestData.Data, cred.GetData())
	require.Equal(suite.T(), suite.GetByPath_TestData.Metadata, cred.GetMetadata())
}

func (suite *ProviderTestSuite)TestProvider_Get_GenericCredential() {

	//setup
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	cred, err := suite.Provider.Get(suite.Get_TestData.QueryData)

	//assert
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), suite.Get_TestData.Data, cred.GetData())
	require.Equal(suite.T(), suite.Get_TestData.Metadata, cred.GetMetadata())
	require.NotEqual(suite.T(), "summary", cred.GetSecretType())

	for key, _ := range cred.GetData() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}

	for key, _ := range cred.GetMetadata() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}

}

func (suite *ProviderTestSuite)TestProvider_Get_SshCredential() {

	//setup
	if enabled, ok := suite.Provider.Capabilities()[constants.CAP_CRED_SSH]; !ok || !enabled {
		suite.T().Skip("skipping test, provider does not support it")
	}
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	cred, err := suite.Provider.Get(suite.Get_Ssh_TestData.QueryData)
	sshCred, castOk := cred.(*credentials.Ssh)

	//assert
	require.True(suite.T(), castOk, "should be able to cast credential interface to correct type (ssh)")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), suite.Get_Ssh_TestData.Data, sshCred.GetData())
	require.Equal(suite.T(), suite.Get_Ssh_TestData.Metadata, sshCred.GetMetadata())
	require.Equal(suite.T(), "ssh", sshCred.GetSecretType())

	require.NotEmpty(suite.T(), sshCred.Key())
	for key, _ := range cred.GetData() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}

	for key, _ := range cred.GetMetadata() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}
}

func (suite *ProviderTestSuite)TestProvider_Get_TextCredential() {

	//setup
	if enabled, ok := suite.Provider.Capabilities()[constants.CAP_CRED_TEXT]; !ok || !enabled {
		suite.T().Skip("skipping test, provider does not support it")
	}
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	cred, err := suite.Provider.Get(suite.Get_Text_TestData.QueryData)
	textCred, castOk := cred.(*credentials.Text)

	//assert
	require.True(suite.T(), castOk, "should be able to cast credential interface to correct type (text)")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), suite.Get_Text_TestData.Data, textCred.GetData())
	require.Equal(suite.T(), suite.Get_Text_TestData.Metadata, textCred.GetMetadata())
	require.Equal(suite.T(), "text", textCred.GetSecretType())

	require.NotEmpty(suite.T(), textCred.Text())
	for key, _ := range cred.GetData() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}

	for key, _ := range cred.GetMetadata() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}
}

func (suite *ProviderTestSuite)TestProvider_Get_UserPassCredential() {

	//setup
	if enabled, ok := suite.Provider.Capabilities()[constants.CAP_CRED_USERPASS]; !ok || !enabled {
		suite.T().Skip("skipping test, provider does not support it")
	}
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	cred, err := suite.Provider.Get(suite.Get_UserPass_TestData.QueryData)
	userPassCred, castOk := cred.(*credentials.UserPass)

	//assert
	require.True(suite.T(), castOk, "should be able to cast credential interface to correct type (ssh)")
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), suite.Get_UserPass_TestData.Data, userPassCred.GetData())
	require.Equal(suite.T(), suite.Get_UserPass_TestData.Metadata, userPassCred.GetMetadata())
	require.Equal(suite.T(), "userpass", userPassCred.GetSecretType())

	require.NotEmpty(suite.T(), userPassCred.Username())
	require.NotEmpty(suite.T(), userPassCred.Password())
	for key, _ := range cred.GetData() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}

	for key, _ := range cred.GetMetadata() {
		require.Equal(suite.T(), strings.ToLower(key), key) // all field name should be lowercase
	}
}

func (suite *ProviderTestSuite)TestProvider_List() {

	//setup
	if enabled, ok := suite.Provider.Capabilities()[constants.CAP_LIST]; !ok || !enabled {
		suite.T().Skip("skipping test, provider does not support it")
	}
	aerr := suite.Provider.Authenticate()
	require.NoError(suite.T(), aerr)

	//test
	creds, err := suite.Provider.List(suite.List_TestData.QueryData)

	//assert
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), suite.List_TestData.Results, creds)
}


//TODO:
// List with filter?
// Get by Path?
// to json yaml raw table etc