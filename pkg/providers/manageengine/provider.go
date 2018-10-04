package manageengine

import (
	"github.com/analogj/tentacle/pkg/providers/base"
	"github.com/analogj/tentacle/pkg/credentials"
	manageengineapi "github.com/analogj/tentacle/pkg/providers/manageengine/api"
	"net/http"
	"fmt"
	"github.com/analogj/tentacle/pkg/errors"
	"strings"
	"github.com/analogj/tentacle/pkg/constants"
)

type provider struct {
	*base.Provider

	client *manageengineapi.Client
}

func New(alias string, config map[string]interface{}) (*provider, error) {
	p := new(provider)

	//validate the config and assign it to ProviderConfig
	p.Provider = new(base.Provider)
	p.ProviderConfig = config
	p.Alias = alias

	return p, p.ValidateRequireAllOf([]string{"host", "token"}, config)
}

func (p *provider) Capabilities() map[string]bool {
	return map[string]bool{
		constants.CAP_GET: true,
		constants.CAP_LIST: true,
		constants.CAP_GET_BY_ID: true,
		constants.CAP_GET_BY_PATH: true,
		constants.CAP_CRED_USERPASS: true,
	}
}

func (p *provider) Authenticate() error {

	var scheme string
	if schemeVal, ok := p.ProviderConfig["host"].(string); ok {
		scheme = schemeVal
	} else {
		scheme = "https"
	}

	p.client = manageengineapi.New(p.ProviderConfig["host"].(string), p.ProviderConfig["token"].(string), scheme, &http.Client{})

	if p.HttpClient != nil {
		p.client.HttpClient = p.HttpClient
	}

	return nil
}

func (p *provider) Get(queryData map[string]string) (credentials.GenericInterface, error) {

	//check if we have accountId && resourceid or path

	path, pathOk := queryData["path"]
	accountId, idOk := queryData["id"]
	resourceId, resourceIdOk := queryData["resourceid"]

	if !pathOk && !(resourceIdOk && idOk) {
		return nil, errors.InvalidArgumentsError("either path or id and resourceid must be specified")
	}

	if pathOk {
		parts := strings.Split(path, "/")
		if len(parts) != 2 {
			return nil, errors.InvalidArgumentsError("parts could not be seperated correctly")
		}

		resourceName := parts[0]
		accountName := parts[1]

		accountByName, err := p.client.GetAccountByName(resourceName, accountName)
		if err != nil {
			return nil, err
		}

		accountId = accountByName.Operation.Details.AccountId
		resourceId = accountByName.Operation.Details.ResourceId
	}

	accountPasswordInfo, err := p.client.GetAccountPassword(resourceId, accountId)
	if err != nil {
		return nil, err
	}

	// account details endpoint doesnt actually include much account information. Lets do a resource lookup, and find the account in there.
	resourceAccounts, err := p.client.GetResourceAccounts(resourceId)
	if err != nil {
		return nil, err
	}

	userPassSecret := new(credentials.UserPass)
	userPassSecret.Init()
	userPassSecret.Id = accountId
	userPassSecret.SetPassword(accountPasswordInfo.Operation.Details.Password)
	userPassSecret.Metadata["type"] = "account"
	userPassSecret.Metadata["resource_id"] = resourceId
	userPassSecret.Metadata["resource_type"] = resourceAccounts.Operation.Details.ResourceType
	userPassSecret.Metadata["resource_name"] = resourceAccounts.Operation.Details.ResourceName

	for _, account := range resourceAccounts.Operation.Details.Accounts{
		if account.Id == accountId {
			userPassSecret.Name = fmt.Sprintf("%s/%s", resourceAccounts.Operation.Details.ResourceName, account.Name)
			userPassSecret.SetUsername(account.Name)
			break
		}
	}

	return userPassSecret, nil
}

func (p *provider) List(queryData map[string]string) ([]credentials.SummaryInterface, error) {

	summarySecrets := []credentials.SummaryInterface{}

	resourceId, resourceIdOk := queryData["resourceid"]
	if  resourceIdOk {
		accountsList, err := p.client.GetResourceAccounts(resourceId)
		if err != nil {
			return nil, err
		}

		for _, val := range accountsList.Operation.Details.Accounts {
			summary := new(credentials.Summary)
			summary.Init()
			summary.Id = val.Id
			summary.Name = fmt.Sprintf("%s/%s",accountsList.Operation.Details.ResourceName, val.Name)
			summary.Metadata["type"] = "account"
			summary.Metadata["resource_id"] = accountsList.Operation.Details.ResourceId
			summary.Metadata["resource_type"] = accountsList.Operation.Details.ResourceType
			summary.Metadata["resource_name"] = accountsList.Operation.Details.ResourceName
			summarySecrets = append(summarySecrets, summary)
		}

	} else {
		//no resource id specified, lets list all resources.
		resourcesList, err := p.client.GetResources()
		if err != nil {
			return nil, err
		}

		for _, val := range resourcesList.Operation.Details {
			summary := new(credentials.Summary)
			summary.Init()
			summary.Id = val.Id
			summary.Name = val.Name
			summary.Description = val.Description
			summary.Metadata["type"] = "resource"
			summary.Metadata["resource_type"] = val.Type
			summary.Metadata["accounts"] = val.NumberOfAccounts

			summarySecrets = append(summarySecrets, summary)
		}

	}


	return summarySecrets, nil
}